# Transfers
(*LoanWriteback.Transfers*)

## Overview

### Available Operations

* [Create](#create) - Create transfer
* [GetCreateModel](#getcreatemodel) - Get create transfer model

## Create

The *Create transfer* endpoint creates a new [transfer](https://docs.codat.io/lending-api#/schemas/Transfer) for a given company's connection.

[Transfers](https://docs.codat.io/lending-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create transfer model](https://docs.codat.io/lending-api#/operations/get-create-transfers-model).

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v6"
	"context"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/types"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Transfers.Create(ctx, operations.CreateTransferRequest{
        AccountingTransfer: &shared.AccountingTransfer{
            ContactRef: &shared.ContactRef{
                DataType: shared.ContactRefDataTypeCustomers.ToPointer(),
                ID: "80000028-167239230944",
            },
            Date: lending.String("2023-01-26T11:51:18.104Z"),
            DepositedRecordRefs: []shared.RecordRef{
                shared.RecordRef{
                    DataType: lending.String("invoice"),
                },
            },
            Description: lending.String("test transfers push 20230126 12.08"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("80000028-1671794219"),
                },
                Amount: types.MustNewDecimalFromString("12"),
                Currency: lending.String("USD"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(true),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00Z"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00Z"),
            Status: shared.AccountingTransferStatusUnknown.ToPointer(),
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("80000004-1671793811"),
                },
                Amount: types.MustNewDecimalFromString("12"),
                Currency: lending.String("EUR"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "80000001-1674553252",
                    Name: lending.String("Class 1"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreateTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateTransferRequest](../../pkg/models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateTransferResponse](../../pkg/models/operations/createtransferresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetCreateModel

The *Get create transfer model* endpoint returns the expected data for the request payload when creating a [transfer](https://docs.codat.io/lending-api#/schemas/Transfer) for a given company and integration.

[Transfers](https://docs.codat.io/lending-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v6"
	"context"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Transfers.GetCreateModel(ctx, operations.GetCreateTransfersModelRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCreateTransfersModelRequest](../../pkg/models/operations/getcreatetransfersmodelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetCreateTransfersModelResponse](../../pkg/models/operations/getcreatetransfersmodelresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |