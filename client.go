package go_storage

import (
	"context"
	"errors"
	"github.com/globalxtreme/go-storage/RPC/gRPC/Storage"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type PublicStorageClient struct {
	gRPCClient
	credential Storage.PublicStorageCredential

	PublicStorage Storage.PublicStorageClient
}

type PublicStorageUpload struct {
	File  multipart.File
	Path  string
	Name  string
	Title string
}

type PublicStorageMove struct {
	File  string
	Path  string
	Name  string
	Title string
}

func NewPublicStorageClient(timeout ...time.Duration) (*PublicStorageClient, context.CancelFunc) {
	clientId := os.Getenv("PUBLIC_STORAGE_CLIENT_ID")
	clientSecret := os.Getenv("PUBLIC_STORAGE_CLIENT_SECRET")
	address := os.Getenv("PUBLIC_STORAGE_HOST")

	if len(clientId) == 0 || len(clientSecret) == 0 || len(address) == 0 {
		log.Panicf("Please setup your public storage Client ID and Screct and Host!")
	}

	client := PublicStorageClient{}
	cleanup := client.RPCDialClient(address, timeout...)

	client.PublicStorage = Storage.NewPublicStorageClient(client.Conn)
	client.credential = Storage.PublicStorageCredential{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}

	return &client, cleanup
}

func (rpc *PublicStorageClient) UploadFile(store PublicStorageUpload) (*Storage.PublicStorageResponse, error) {
	if len(store.Name) == 0 {
		return nil, errors.New("Please enter your filename [Name]")
	}

	filename := generateRandomName() + filepath.Ext(store.Name)

	stream, err := rpc.PublicStorage.Store(rpc.Ctx)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	for {
		n, err := store.File.Read(buf)
		if err == io.EOF {
			res, err := stream.CloseAndRecv()
			if err != nil {
				return nil, err
			}

			return res, nil
		}
		if err != nil {
			return nil, err
		}

		req := Storage.PublicStorageStoreRequest{
			Content:    buf[:n],
			Path:       store.Path,
			Filename:   filename,
			Title:      store.Title,
			Credential: &rpc.credential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func (rpc *PublicStorageClient) MoveFile(store PublicStorageMove) (*Storage.PublicStorageResponse, error) {
	var filename string
	if len(store.Name) == 0 {
		filename = generateRandomName() + filepath.Ext(store.File)
	} else {
		filename = store.Name
	}

	file, err := os.Open(store.File)
	if err != nil {
		return nil, err
	}

	stream, err := rpc.PublicStorage.Store(rpc.Ctx)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			res, err := stream.CloseAndRecv()
			if err != nil {
				return nil, err
			}

			return res, nil
		}
		if err != nil {
			return nil, err
		}

		req := Storage.PublicStorageStoreRequest{
			Content:    buf[:n],
			Path:       store.Path,
			Filename:   filename,
			Title:      store.Title,
			Credential: &rpc.credential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func (rpc *PublicStorageClient) Delete(path string) (*Storage.PublicStorageResponse, error) {
	res, err := rpc.PublicStorage.Delete(rpc.Ctx, &Storage.PublicStorageDeleteRequest{
		Path:       path,
		Credential: &rpc.credential,
	})
	if err != nil {
		log.Panicf("Delete file invalid: %v", err)
	}

	return res, nil
}

type gRPCClient struct {
	Ctx    context.Context
	Conn   *grpc.ClientConn
	Cancel context.CancelFunc
}

func (client *gRPCClient) RPCDialClient(host string, timeout ...time.Duration) context.CancelFunc {
	dialTimeout := 5 * time.Second
	if len(timeout) > 0 {
		dialTimeout = timeout[0]
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)

	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicf("Did not connect to %s: %v", host, err)
	}

	client.Ctx = ctx
	client.Conn = conn
	client.Cancel = cancel

	cleanup := func() {
		client.Cancel()
		client.Conn.Close()
	}

	return cleanup
}

func generateRandomName() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, 20)
	for i := 0; i < 20; i++ {
		randomBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomBytes) + strconv.FormatInt(time.Now().UnixNano(), 10)
}
