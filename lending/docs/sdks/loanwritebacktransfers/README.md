# LoanWriteback.Transfers

### Available Operations

* [Create](#create) - Create transfer
* [GetCreateModel](#getcreatemodel) - Get create transfer model

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
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/types"
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
            ContactRef: &shared.AccountingTransferContactRef{
                DataType: lending.String("velit"),
                ID: "66c8dd6b-1442-4907-8747-78a7bd466d28",
            },
            Date: lending.String("2022-10-23T00:00:00.000Z"),
            DepositedRecordRefs: []shared.RecordRef{
                shared.RecordRef{
                    DataType: lending.String("journalEntry"),
                    ID: lending.String("0ab3cdca-4251-4904-a523-c7e0bc7178e4"),
                },
            },
            Description: lending.String("odio"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("96f2a70c-6882-482a-a482-562f222e9817"),
                    Name: lending.String("Sheldon Boehm"),
                },
                Amount: types.MustNewDecimalFromString("7241.68"),
                Currency: lending.String("EUR"),
            },
            ID: lending.String("61e6b7b9-5bc0-4ab3-820c-4f3789fd871f"),
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(false),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "pariatur": map[string]interface{}{
                        "possimus": "quia",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("efd121aa-6f1e-4674-bdb0-4f15756082d6"),
                    Name: lending.String("Miss Percy Parisian"),
                },
                Amount: types.MustNewDecimalFromString("984.78"),
                Currency: lending.String("EUR"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "17051339-d080-486a-9840-394c26071f93",
                    Name: lending.String("Ms. Glen Zboncak"),
                },
            },
        },
        AllowSyncOnPushComplete: lending.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: lending.Bool(false),
        TimeoutInMinutes: lending.Int(162954),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateTransferResponse](../../models/operations/createtransferresponse.md), error**


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
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCreateTransfersModelRequest](../../models/operations/getcreatetransfersmodelrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetCreateTransfersModelResponse](../../models/operations/getcreatetransfersmodelresponse.md), error**

