# Files

## Overview

Endpoints to manage uploaded files.

### Available Operations

* [DownloadFiles](#downloadfiles) - Download all files for a company
* [ListFiles](#listfiles) - List all files uploaded by a company
* [UploadFiles](#uploadfiles) - Upload files for a company

## DownloadFiles

You can specify a date to download specific files for.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/files"
	"github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Date: codatfiles.String("provident"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## ListFiles

Returns an array of files that have been uploaded for a given company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/files"
	"github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.ListFiles(ctx, operations.ListFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Files != nil {
        // handle response
    }
}
```

## UploadFiles

Upload files

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/files"
	"github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.UploadFiles(ctx, operations.UploadFilesRequest{
        RequestBody: &operations.UploadFilesRequestBody{
            Content: []byte("distinctio"),
            RequestBody: "quibusdam",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```