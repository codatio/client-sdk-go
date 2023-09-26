# LoanWriteback.Transfers

### Available Operations

* [Create](#create) - Create transfer
* [GetCreateModel](#getcreatemodel) - Get create transfer model

## Create

The *Create transfer* endpoint creates a new [transfer](https://docs.codat.io/lending-api#/schemas/Transfer) for a given company's connection.

[Transfers](https://docs.codat.io/lending-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create transfer model](https://docs.codat.io/lending-api#/operations/get-create-transfers-model).

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
                DataType: shared.DataTypeInvoices.ToPointer(),
                ID: "ac366c8d-d6b1-4442-9074-74778a7bd466",
            },
            Date: lending.String("2022-10-23T00:00:00.000Z"),
            DepositedRecordRefs: []shared.RecordRef{
                shared.RecordRef{
                    DataType: lending.String("journalEntry"),
                    ID: lending.String("8c10ab3c-dca4-4251-904e-523c7e0bc717"),
                },
            },
            Description: lending.String("totam"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("e4796f2a-70c6-4882-82aa-482562f222e9"),
                    Name: lending.String("Carl Koch"),
                },
                Amount: types.MustNewDecimalFromString("829.71"),
                Currency: lending.String("USD"),
            },
            ID: lending.String("cbe61e6b-7b95-4bc0-ab3c-20c4f3789fd8"),
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(false),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "a": map[string]interface{}{
                        "error": "sint",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: lending.String("dd2efd12-1aa6-4f1e-a74b-db04f1575608"),
                    Name: lending.String("Rosemarie Jacobs"),
                },
                Amount: types.MustNewDecimalFromString("6802.7"),
                Currency: lending.String("GBP"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "9f1d1705-1339-4d08-886a-1840394c2607",
                    Name: lending.String("Elisa Mosciski"),
                },
            },
        },
        AllowSyncOnPushComplete: lending.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: lending.Int(374323),
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

The *Get create transfer model* endpoint returns the expected data for the request payload when creating a [transfer](https://docs.codat.io/lending-api#/schemas/Transfer) for a given company and integration.

[Transfers](https://docs.codat.io/lending-api#/schemas/Transfer) record the movement of money between two bank accounts, or between a bank account and a nominal account.

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

