package gxstorage

import (
	"context"
	"errors"
	"github.com/globalxtreme/go-storage/v2/grpc/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	PublicStorageRPCClient     storage.PublicStorageClient
	PublicStorageRPCCredential storage.PublicStorageCredential
	PublicStorageRPCConf       publicStorageConf
	PublicStorageRPCActive     bool
)

type publicStorageConf struct {
	Timeout          time.Duration
	KeepalivePing    time.Duration
	KeepaliveTimeout time.Duration
}

type PublicStorageUpload struct {
	File        multipart.File
	FileHandler *multipart.FileHeader
	Path        string
	Name        string
	Title       string
	MimeType    *string
}

type PublicStorageMove struct {
	File     string
	Path     string
	Name     string
	Title    string
	MimeType *string
}

func InitPublicStorageRPC() func() {
	clientId := os.Getenv("PUBLIC_STORAGE_CLIENT_ID")
	clientSecret := os.Getenv("PUBLIC_STORAGE_CLIENT_SECRET")
	address := os.Getenv("PUBLIC_STORAGE_HOST")

	toInt := func(val string) int64 {
		parseVal, _ := strconv.Atoi(val)
		return int64(parseVal)
	}

	var timeout time.Duration
	if timeoutENV := os.Getenv("PUBLIC_STORAGE_TIMEOUT"); timeoutENV != "" {
		timeout = time.Duration(toInt(timeoutENV)) * time.Second
	} else {
		timeout = 5 * time.Second
	}

	var keepalivePing time.Duration
	if keepalivePingENV := os.Getenv("PUBLIC_STORAGE_KEEPALIVE_PING"); keepalivePingENV != "" {
		keepalivePing = time.Duration(toInt(keepalivePingENV)) * time.Second
	} else {
		keepalivePing = 60 * time.Second
	}

	var keepaliveTimeout time.Duration
	if keepaliveTimeoutENV := os.Getenv("PUBLIC_STORAGE_KEEPALIVE_TIMEOUT"); keepaliveTimeoutENV != "" {
		keepaliveTimeout = time.Duration(toInt(keepaliveTimeoutENV)) * time.Second
	} else {
		keepaliveTimeout = 20 * time.Second
	}

	if len(clientId) == 0 || len(clientSecret) == 0 || len(address) == 0 {
		log.Panicf("Please setup your public storage Client ID and Screct and Host!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	keepaliveParam := keepalive.ClientParameters{
		Time:                keepalivePing,
		Timeout:             keepaliveTimeout,
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

	PublicStorageRPCClient = storage.NewPublicStorageClient(conn)
	PublicStorageRPCCredential = storage.PublicStorageCredential{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
	PublicStorageRPCConf = publicStorageConf{
		Timeout:          timeout,
		KeepalivePing:    keepalivePing,
		KeepaliveTimeout: keepaliveTimeout,
	}
	PublicStorageRPCActive = true

	cleanup := func() {
		cancel()
		conn.Close()
	}

	return cleanup
}

func UploadFile(store PublicStorageUpload) (*storage.PublicStorageResponse, error) {
	if len(store.Name) == 0 {
		return nil, errors.New("Please enter your filename [Name]")
	}

	mimeType := GetMimeType(store.File, store.FileHandler, store.MimeType)
	filename := generateRandomName() + filepath.Ext(store.Name)

	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
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

		req := storage.PublicStorageStoreRequest{
			Content:    buf[:n],
			Path:       store.Path,
			Filename:   filename,
			Title:      store.Title,
			MimeType:   mimeType,
			Credential: &PublicStorageRPCCredential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func MoveFile(store PublicStorageMove) (*storage.PublicStorageResponse, error) {
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

	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
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

		req := storage.PublicStorageStoreRequest{
			Content:    buf[:n],
			Path:       store.Path,
			Filename:   filename,
			Title:      store.Title,
			MimeType:   GetMimeTypeByPath(store.File, store.MimeType),
			Credential: &PublicStorageRPCCredential,
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func Delete(path string) (*storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
	defer cancel()

	res, err := PublicStorageRPCClient.Delete(ctx, &storage.PublicStorageDeleteRequest{
		Path:       path,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("Delete file invalid: %v", err)
	}

	return res, nil
}

func GetMimeType(file multipart.File, handler *multipart.FileHeader, mimeType *string) string {
	if mimeType == nil || *mimeType == "" {
		buf := make([]byte, 512)
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Panicf("Unable to reading file: %v", err)
		}

		mimeTypeSystem := http.DetectContentType(buf[:n])
		if mimeTypeSystem == "application/zip" {
			ext := strings.ToLower(filepath.Ext(handler.Filename))
			switch ext {
			case ".xlsx":
				return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
			case ".docx":
				return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
			}
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Panicf("Unable to reset file pointer: %v", err)
		}

		return mimeTypeSystem
	}

	return *mimeType
}

func GetMimeTypeByPath(path string, mimeType *string) string {
	if mimeType == nil || *mimeType == "" {
		file, err := os.Open(path)
		if err != nil {
			log.Panicf("Unable to open file: %v", err)
		}
		defer file.Close()

		buf := make([]byte, 512)
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Panicf("Unable to reading file: %v", err)
		}

		mimeTypeSystem := http.DetectContentType(buf[:n])
		if mimeTypeSystem == "application/zip" {
			ext := strings.ToLower(filepath.Ext(path))
			switch ext {
			case ".xlsx":
				return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
			case ".docx":
				return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
			}
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Panicf("Unable to reset file pointer: %v", err)
		}

		return mimeTypeSystem
	}

	return *mimeType
}

func generateRandomName() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, 20)
	for i := 0; i < 20; i++ {
		randomBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomBytes) + strconv.FormatInt(time.Now().UnixNano(), 10)
}
