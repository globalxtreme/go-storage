package gxstorage

import (
	"context"
	"errors"
	"github.com/globalxtreme/go-storage/v2/grpc/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/durationpb"
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
	File          multipart.File
	FileHandler   *multipart.FileHeader
	Path          string
	Name          string
	Title         string
	MimeType      *string
	OwnerId       string
	OwnerType     string
	Reference     string
	ReferenceType string
	CreatedBy     string
	CreatedByName string
	SavedUntil    time.Duration
	WithWatermark bool
	External      bool
}

type PublicStorageMove struct {
	File          string
	Path          string
	Name          string
	Title         string
	MimeType      *string
	OwnerId       string
	OwnerType     string
	Reference     string
	ReferenceType string
	CreatedBy     string
	CreatedByName string
	SavedUntil    time.Duration
	WithWatermark bool
	External      bool
}

type PublicStorageMoveCopyToAnotherService struct {
	File       string
	ToPath     string
	ToClientID string
}

type PublicStorageMoveCopyFromAnotherService struct {
	File         string
	ToPath       string
	FromClientID string
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
	filename := GenerateFilename(store.Name)

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
			Content:       buf[:n],
			Path:          store.Path,
			Filename:      filename,
			OriginalName:  store.Name,
			Title:         store.Title,
			MimeType:      mimeType,
			OwnerId:       store.OwnerId,
			OwnerType:     store.OwnerType,
			CreatedBy:     store.CreatedBy,
			Reference:     store.Reference,
			ReferenceType: store.ReferenceType,
			CreatedByName: store.CreatedByName,
			Credential:    &PublicStorageRPCCredential,
			WithWatermark: store.WithWatermark,
			External:      store.External,
		}

		if store.SavedUntil > 0 {
			req.SavedUntil = durationpb.New(store.SavedUntil)
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func MoveFile(store PublicStorageMove) (*storage.PublicStorageResponse, error) {
	mimeType := GetMimeTypeByPath(store.File, store.MimeType)
	filename := GenerateFilename(store.Name)

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
			Content:       buf[:n],
			Path:          store.Path,
			Filename:      filename,
			OriginalName:  store.Name,
			Title:         store.Title,
			MimeType:      mimeType,
			OwnerId:       store.OwnerId,
			OwnerType:     store.OwnerType,
			Reference:     store.Reference,
			ReferenceType: store.ReferenceType,
			CreatedBy:     store.CreatedBy,
			CreatedByName: store.CreatedByName,
			Credential:    &PublicStorageRPCCredential,
			WithWatermark: store.WithWatermark,
			External:      store.External,
		}

		if store.SavedUntil > 0 {
			req.SavedUntil = durationpb.New(store.SavedUntil)
		}

		if err := stream.Send(&req); err != nil {
			return nil, err
		}
	}
}

func MoveToAnotherService(form PublicStorageMoveCopyToAnotherService) (*storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
	defer cancel()

	res, err := PublicStorageRPCClient.MoveToAnotherService(ctx, &storage.PublicStorageMoveCopyRequest{
		File:       form.File,
		ClientID:   form.ToClientID,
		ToPath:     form.ToPath,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("move file to another service invalid: %v", err)
	}

	return res, nil
}

func CopyToAnotherService(form PublicStorageMoveCopyToAnotherService) (*storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
	defer cancel()

	res, err := PublicStorageRPCClient.CopyToAnotherService(ctx, &storage.PublicStorageMoveCopyRequest{
		File:       form.File,
		ClientID:   form.ToClientID,
		ToPath:     form.ToPath,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("Copy file to another service invalid: %v", err)
	}

	return res, nil
}

func MoveFromAnotherService(form PublicStorageMoveCopyFromAnotherService) (*storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
	defer cancel()

	res, err := PublicStorageRPCClient.MoveFromAnotherService(ctx, &storage.PublicStorageMoveCopyRequest{
		File:       form.File,
		ClientID:   form.FromClientID,
		ToPath:     form.ToPath,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("move file from another service invalid: %v", err)
	}

	return res, nil
}

func CopyFromAnotherService(form PublicStorageMoveCopyFromAnotherService) (*storage.PublicStorageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PublicStorageRPCConf.Timeout)
	defer cancel()

	res, err := PublicStorageRPCClient.CopyFromAnotherService(ctx, &storage.PublicStorageMoveCopyRequest{
		File:       form.File,
		ClientID:   form.FromClientID,
		ToPath:     form.ToPath,
		Credential: &PublicStorageRPCCredential,
	})
	if err != nil {
		log.Panicf("Copy file from another service invalid: %v", err)
	}

	return res, nil
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

func GenerateFilename(filename string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

	randomBytes := make([]byte, 36)
	for i := 0; i < 36; i++ {
		randomBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomBytes) + strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(filename)
}
