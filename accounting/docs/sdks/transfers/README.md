# Transfers

## Overview

Transfers

### Available Operations

* [Create](#create) - Create transfer
* [Get](#get) - Get transfer
* [GetCreateModel](#getcreatemodel) - Get create transfer model
* [List](#list) - List transfers

## Create

The *Create transfer* endpoint creates a new [transfer](https://docs.codat.io/accounting-api#/schemas/Transfer) for a given company's connection.

[Transfers](https://docs.codat.io/accounting-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create transfer model](https://docs.codat.io/accounting-api#/operations/get-create-transfers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.Create(ctx, operations.CreateTransferRequest{
        Transfer: &shared.Transfer{
            ContactRef: &shared.TransferContactRef{
                DataType: codataccounting.String("laudantium"),
                ID: "1ddf7e08-8f74-4ef5-8c92-16e8926313bb",
            },
            Date: codataccounting.String("2022-10-23T00:00:00.000Z"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("quo"),
                    ID: codataccounting.String("2c8d2701-096b-466a-96e3-e1d9d3b66033"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("quaerat"),
                    ID: codataccounting.String("a11aa1d5-d224-47de-9b3d-46170e768a96"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("nobis"),
                    ID: codataccounting.String("b3987883-98eb-4a1b-bf71-43356f6349a1"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("voluptas"),
                    ID: codataccounting.String("4249b211-ce46-4b95-9652-b158ca9142f0"),
                },
            },
            Description: codataccounting.String("veniam"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("2632b31c-ad69-42ff-8874-5005e9d3d934"),
                    Name: codataccounting.String("Kevin Franecki"),
                },
                Amount: codataccounting.Float64(3566.2),
                Currency: codataccounting.String("EUR"),
            },
            ID: codataccounting.String("388664f6-9855-430a-ae2a-ed6aaf863c28"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aut": map[string]interface{}{
                        "vel": "perspiciatis",
                        "id": "amet",
                        "pariatur": "iste",
                        "voluptatem": "voluptas",
                    },
                    "hic": map[string]interface{}{
                        "officiis": "soluta",
                        "at": "nostrum",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("ad7ec739-4f25-4f63-8b37-30714e6be8c3"),
                    Name: codataccounting.String("Michael Mertz"),
                },
                Amount: codataccounting.Float64(3122.98),
                Currency: codataccounting.String("EUR"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "42ac299a-6e5e-47ae-b134-02e945f53743",
                    Name: codataccounting.String("Moses Schulist Jr."),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateTransferResponse](../../models/operations/createtransferresponse.md), error**


## Get

The *Get transfer* endpoint returns a single transfer for a given transferId.

[Transfers](https://docs.codat.io/accounting-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support getting a specific transfer.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.Get(ctx, operations.GetTransferRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransferID: "sint",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetTransferRequest](../../models/operations/gettransferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.GetTransferResponse](../../models/operations/gettransferresponse.md), error**


## GetCreateModel

The *Get create transfer model* endpoint returns the expected data for the request payload when creating a [transfer](https://docs.codat.io/accounting-api#/schemas/Transfer) for a given company and integration.

[Transfers](https://docs.codat.io/accounting-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating a transfer.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.GetCreateModel(ctx, operations.GetCreateTransfersModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCreateTransfersModelRequest](../../models/operations/getcreatetransfersmodelrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetCreateTransfersModelResponse](../../models/operations/getcreatetransfersmodelresponse.md), error**


## List

The *List transfers* endpoint returns a list of [transfers](https://docs.codat.io/accounting-api#/schemas/Transfer) for a given company's connection.

[Transfers](https://docs.codat.io/accounting-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.List(ctx, operations.ListTransfersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("praesentium"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListTransfersRequest](../../models/operations/listtransfersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListTransfersResponse](../../models/operations/listtransfersresponse.md), error**

