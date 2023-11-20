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
// Enter public storage host and port (contact developer)
PUBLIC_STORAGE_HOST=<public-storage-host:port>

// Enter your client id (contact developer for generate your credential)
PUBLIC_STORAGE_CLIENT_ID=<your-client-id>

// Enter your client secret (contact developer for generate your credential)
PUBLIC_STORAGE_CLIENT_SECRET=<your-client-id>

// Enter public storage domain and your service name (must be the same as what is registered in public storage)
PUBLIC_STORAGE_GATEWAY_BASE="http://localhost:8000/link/${SERVICE}"
```