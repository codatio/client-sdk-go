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
        Date: codatfiles.String("2022-10-23T00:00:00.000Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DownloadFilesRequest](../../models/operations/downloadfilesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.DownloadFilesResponse](../../models/operations/downloadfilesresponse.md), error**


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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListFilesRequest](../../models/operations/listfilesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |


### Response

**[*operations.ListFilesResponse](../../models/operations/listfilesresponse.md), error**


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

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UploadFilesRequest](../../models/operations/uploadfilesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.UploadFilesResponse](../../models/operations/uploadfilesresponse.md), error**
