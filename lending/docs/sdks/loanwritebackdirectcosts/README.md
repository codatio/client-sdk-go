# LoanWritebackDirectCosts
(*LoanWriteback.DirectCosts*)

### Available Operations

* [Create](#create) - Create direct cost
* [GetCreateModel](#getcreatemodel) - Get create direct cost model

## Create

The *Create direct cost* endpoint creates a new [direct cost](https://docs.codat.io/lending-api#/schemas/DirectCost) for a given company's connection.

[Direct costs](https://docs.codat.io/lending-api#/schemas/DirectCost) are the expenses associated with a business' operations. For example, purchases of raw materials that are paid off at the point of the purchase and service fees are considered direct costs.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/lending-api#/operations/get-create-directCosts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating an account.


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
    res, err := s.LoanWriteback.DirectCosts.Create(ctx, operations.CreateDirectCostRequest{
        AccountingDirectCost: &shared.AccountingDirectCost{
            ContactRef: &shared.ContactRef{
                DataType: shared.DataTypeInvoices.ToPointer(),
                ID: "<ID>",
            },
            Currency: "USD",
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{},
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                    },
                    Quantity: types.MustNewDecimalFromString("6384.24"),
                    TaxRateRef: &shared.TaxRateRef{},
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.RecordRef{
                            DataType: lending.String("transfer"),
                        },
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: lending.String("invoice"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("2884.08"),
                },
            },
            Metadata: &shared.Metadata{},
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.AccountingPaymentAllocationAllocation{
                        AllocatedOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Currency: lending.String("GBP"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{},
                        Currency: lending.String("EUR"),
                        PaidOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                    },
                },
            },
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("9510.62"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "key": map[string]interface{}{
                        "key": "string",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("8915.1"),
            TotalAmount: types.MustNewDecimalFromString("0.86"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateDirectCostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateDirectCostRequest](../../models/operations/createdirectcostrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.CreateDirectCostResponse](../../models/operations/createdirectcostresponse.md), error**


## GetCreateModel

The *Get create direct cost model* endpoint returns the expected data for the request payload when creating a [direct cost](https://docs.codat.io/lending-api#/schemas/DirectCost) for a given company and integration.

[Direct costs](https://docs.codat.io/lending-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating a direct cost.


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
    res, err := s.LoanWriteback.DirectCosts.GetCreateModel(ctx, operations.GetCreateDirectCostsModelRequest{
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
| `request`                                                                                                  | [operations.GetCreateDirectCostsModelRequest](../../models/operations/getcreatedirectcostsmodelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCreateDirectCostsModelResponse](../../models/operations/getcreatedirectcostsmodelresponse.md), error**

