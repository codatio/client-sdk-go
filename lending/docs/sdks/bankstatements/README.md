# BankStatements
(*BankStatements*)

## Overview

Retrieve banking data from linked bank accounts.

### Available Operations

* [EndUploadSession](#enduploadsession) - End upload session
* [GetUploadConfiguration](#getuploadconfiguration) - Get upload configuration
* [SetUploadConfiguration](#setuploadconfiguration) - Set upload configuration
* [StartUploadSession](#startuploadsession) - Start upload session
* [UploadBankStatementData](#uploadbankstatementdata) - Upload data

## EndUploadSession

Use the *End upload session* endpoint to finalize a bank statement upload session. Include a `status` in the request body to indicate if you want to cancel the processing of the dataset or trigger the ingestion and enrichment of the data.

A session is a one-time process that enables you to upload bank statements to Codat. It will time out after 90 minutes if no data is uploaded.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankStatements.EndUploadSession(ctx, operations.EndBankStatementUploadSessionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DatasetID: "d58db2ab-edef-4881-8b6b-a7e79d33244c",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.EndBankStatementUploadSessionRequest](../../pkg/models/operations/endbankstatementuploadsessionrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.EndBankStatementUploadSessionResponse](../../pkg/models/operations/endbankstatementuploadsessionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetUploadConfiguration

Use the *Get upload configuration* endpoint to view the existing bank statement upload configuration for the specified data connection.

With this configuration, you set the source of the data you plan to upload, the ID of the account in third-party banking platform, and a provider ID, if required. This lets us determine the expected format of the data and any source-specific requirements.

When you use the [*Upload data*](https://docs.codat.io/lending-api#/operations/upload-bank-statement-data) endpoint next, you must upload the data for the account you configured. 

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankStatements.GetUploadConfiguration(ctx, operations.GetBankStatementUploadConfigurationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankStatementUploadConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetBankStatementUploadConfigurationRequest](../../pkg/models/operations/getbankstatementuploadconfigurationrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.GetBankStatementUploadConfigurationResponse](../../pkg/models/operations/getbankstatementuploadconfigurationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## SetUploadConfiguration

Use the *Set upload configuration* endpoint to create bank statement upload configuration for the specified data connection. 

With this configuration, you set the source of the data you plan to upload, the ID of the account in third-party banking platform, and a provider ID, if required. This lets us determine the expected format of the data and any source-specific requirements.

Each data connection can only have one configuration for each company and external account ID combination. You will receive a Bad Request response if you try to set it again. 

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankStatements.SetUploadConfiguration(ctx, operations.SetBankStatementUploadConfigurationRequest{
        BankStatementUploadConfiguration: &shared.BankStatementUploadConfiguration{
            AccountID: lending.String("abc123-ABC"),
            Source: shared.SourceCodat.ToPointer(),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankStatementUploadConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.SetBankStatementUploadConfigurationRequest](../../pkg/models/operations/setbankstatementuploadconfigurationrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.SetBankStatementUploadConfigurationResponse](../../pkg/models/operations/setbankstatementuploadconfigurationresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## StartUploadSession

Use the *Start upload session* endpoint to initiate a bank statement upload session for a given company.

A session is a one-time process that enables you to upload bank statements to Codat. It will time out after 90 minutes if no data is uploaded. 

You can only have one active session per data type at a time. You can complete or cancel a session using the [*End upload session*](https://docs.codat.io/lending-api#/operations/end-bank-statement-upload-session) endpoint.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankStatements.StartUploadSession(ctx, operations.StartBankStatementUploadSessionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.StartBankStatementUploadSessionRequest](../../pkg/models/operations/startbankstatementuploadsessionrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.StartBankStatementUploadSessionResponse](../../pkg/models/operations/startbankstatementuploadsessionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## UploadBankStatementData

During an active session, use the **Upload data* endpoint to uploads a page of bank accounts or bank transactions data to the session.

Make sure you created configuration for the account using the [*Set upload configuration*](https://docs.codat.io/lending-api#/operations/set-bank-statement-upload-configuration) endpoint before attempting an upload. 

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankStatements.UploadBankStatementData(ctx, operations.UploadBankStatementDataRequest{
        RequestBody: operations.CreateUploadBankStatementDataRequestBodyAny(
            map[string]any{
                "currency": "Lek",
                "id": "<id>",
                "identifiers": map[string]any{
                    "type": "Investment",
                },
                "modifiedDate": "2022-10-23T00:00:00Z",
                "name": "<value>",
                "sourceModifiedDate": "2022-10-23T00:00:00Z",
                "type": "Credit",
            },
        ),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DatasetID: "15fc5ae4-359b-4dee-ba7c-3a22596ededc",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UploadBankStatementDataRequest](../../pkg/models/operations/uploadbankstatementdatarequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UploadBankStatementDataResponse](../../pkg/models/operations/uploadbankstatementdataresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |