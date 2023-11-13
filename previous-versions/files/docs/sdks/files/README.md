# Files
(*Files*)

## Overview

Endpoints to manage uploaded files.

### Available Operations

* [DownloadFiles](#downloadfiles) - Download all files for a company
* [ListFiles](#listfiles) - List all files uploaded by a company
* [UploadFiles](#uploadfiles) - Upload files for a company

## DownloadFiles

The *Download files* endpoint downloads all files that have  been uploaded by to SMB to Codat. A `date` may be specified to download any files uploaded on the date provided.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
)

func main() {
    s := files.New(
        files.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Date: files.String("2022-10-23T00:00:00.000Z"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DownloadFilesRequest](../../pkg/models/operations/downloadfilesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.DownloadFilesResponse](../../pkg/models/operations/downloadfilesresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.Schema                    | 400,401,402,404,429,500,503         | application/json                    |
| sdkerrors.DownloadFilesErrorMessage | 403                                 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |

## ListFiles

﻿The *List files* endpoint returns a list of all files uploaded to Codat by the SMB. 

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
)

func main() {
    s := files.New(
        files.WithSecurity(shared.Security{
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListFilesRequest](../../pkg/models/operations/listfilesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.ListFilesResponse](../../pkg/models/operations/listfilesresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.Schema                | 401,402,404,429,500,503         | application/json                |
| sdkerrors.ListFilesErrorMessage | 403                             | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## UploadFiles

The *Upload files* endpoint uploads multiple files provided by the SMB to Codat. This may include personal identity documents, pitch decks, contracts, or files with accounting and banking data.

Uploaded files must meet the following requirements:

- Up to 20 files can be uploaded at a time.
- PDF, XLS, XLSX, XLSB, CSV, DOC, DOCX, PPT, PPTX, JPEG, JPG, and PNG files can be uploaded.
- Each file can be up to 10MB in size.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
)

func main() {
    s := files.New(
        files.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.UploadFiles(ctx, operations.UploadFilesRequest{
        RequestBody: &operations.UploadFilesRequestBody{
            Content: []byte("0x6261bDdB39"),
            FileName: "syndicate_central_defect.wav",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UploadFilesRequest](../../pkg/models/operations/uploadfilesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.UploadFilesResponse](../../pkg/models/operations/uploadfilesresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.Schema                  | 400,401,402,404,429,500,503       | application/json                  |
| sdkerrors.UploadFilesErrorMessage | 403                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |
