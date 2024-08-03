package gxstorage

import (
	"context"
	"errors"
	"github.com/globalxtreme/go-storage/RPC/gRPC/Storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	PublicStorageRPCClient     Storage.PublicStorageClient
	PublicStorageRPCCredential Storage.PublicStorageCredential
	PublicStorageRPCActive     bool
)

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

func InitPublicStorageRPC() func() {
	clientId := os.Getenv("PUBLIC_STORAGE_CLIENT_ID")
	clientSecret := os.Getenv("PUBLIC_STORAGE_CLIENT_SECRET")
	address := os.Getenv("PUBLIC_STORAGE_HOST")

	if len(clientId) == 0 || len(clientSecret) == 0 || len(address) == 0 {
		log.Panicf("Please setup your public storage Client ID and Screct and Host!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	keepaliveParam := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}

	conn, err := grpc.DialContext(ctx, address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepaliveParam),
	)
	if err != nil {
		log.Panicf("Did not connect to %s: %v", address, err)
	}

	PublicStorageRPCClient = Storage.NewPublicStorageClient(conn)
	PublicStorageRPCCredential = Storage.PublicStorageCredential{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
	PublicStorageRPCActive = true

	cleanup := func() {
		cancel()
		conn.Close()
	}

	return cleanup
}

func UploadFile(store PublicStorageUpload) (*Storage.PublicStorageResponse, error) {
	if len(store.Name) == 0 {
		return nil, errors.New("Please enter your filename [Name]")
	}

	filename := generateRandomName() + filepath.Ext(store.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := PublicStorageRPCClient.Store(ctx)
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
			Credential: &PublicStorageRPCCredential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func MoveFile(store PublicStorageMove) (*Storage.PublicStorageResponse, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := PublicStorageRPCClient.Store(ctx)
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
			Credential: &PublicStorageRPCCredential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func Delete(path string) (*Storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := PublicStorageRPCClient.Delete(ctx, &Storage.PublicStorageDeleteRequest{
		Path:       path,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("Delete file invalid: %v", err)
	}

	return res, nil
}

func generateRandomName() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, 20)
	for i := 0; i < 20; i++ {
		randomBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomBytes) + strconv.FormatInt(time.Now().UnixNano(), 10)
}
