GlobalXtreme Go Storage Module
======

## Install Private Module Go-Lang

```bash
export GOPRIVATE=github.com/globalxtreme/*

go get github.com/globalxtreme/go-storage
```

## Setup .env

Copy from **.env.example**
```shell
# Enter public storage host and port (contact developer)
PUBLIC_STORAGE_HOST=<public-storage-host:port>

# Enter your client id (contact developer for generate your credential)
PUBLIC_STORAGE_CLIENT_ID=<your-client-id>

# Enter your client secret (contact developer for generate your credential)
PUBLIC_STORAGE_CLIENT_SECRET=<your-client-id>

# Enter public storage domain and your service name (must be the same as what is registered in public storage)
PUBLIC_STORAGE_GATEWAY_BASE="http://localhost:8000/link/${SERVICE}"
```

## Source Code
I assume you are using **github.com/globalxtreme/go-backend-service**

```go
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	publicStorage, cleanup := gxstorage.NewPublicStorageClient()
	defer cleanup()

	var r *http.Request
	file, handler, _ := r.FormFile("your-param-file")
	
	// Upload from endpoint multipart
	upload, err := publicStorage.UploadFile(gxstorage.PublicStorageUpload{
		File:  file,
		Path:  "images/testings/",
		Name:  handler.Filename,
		Title: "Your title or description",
	})
	if err != nil {
		log.Panicf("Unable to upload file public storage: %v", +err)
	}
	fmt.Println(upload)

	// Move file from path
	move, err := publicStorage.MoveFile(gxstorage.PublicStorageMove{
		File: "storages/app/PDFs/testings/hasil.pdf",
		Name: "hasil.pdf",
		Path: "PDFs/testings/",
	})
	if err != nil {
		log.Panicf("Unable to move file public storage: %v", +err)
	}
	fmt.Println(move)

	// Delete file with path
	del, err := publicStorage.Delete("http://localhost:8080/link/dev-test/pdfs/testings/hasil.pdf")
	if err != nil {
		log.Panicf("%v", err)
	}
	fmt.Println(del)
}

```