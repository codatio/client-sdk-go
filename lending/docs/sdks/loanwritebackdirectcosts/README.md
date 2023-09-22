# LoanWriteback.DirectCosts

### Available Operations

* [Create](#create) - Create direct cost
* [GetCreateModel](#getcreatemodel) - Get create direct cost model

## Create

The *Create direct cost* endpoint creates a new [direct cost](https://docs.codat.io/accounting-api#/schemas/DirectCost) for a given company's connection.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are the expenses associated with a business' operations. For example, purchases of raw materials that are paid off at the point of the purchase and service fees are considered direct costs.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/accounting-api#/operations/get-create-directCosts-model).

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
                DataType: lending.String("magni"),
                ID: "d502a94b-b4f6-43c9-a9e9-a3efa77dfb14",
            },
            Currency: "EUR",
            CurrencyRate: types.MustNewDecimalFromString("8137.98"),
            ID: lending.String("66ae395e-fb9b-4a88-b3a6-6997074ba446"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: lending.String("b6e21419-5989-40af-a563-e2516fe4c8b7"),
                        Name: lending.String("Diane VonRueden"),
                    },
                    Description: lending.String("nihil"),
                    DiscountAmount: types.MustNewDecimalFromString("9988.48"),
                    DiscountPercentage: types.MustNewDecimalFromString("8411.4"),
                    ItemRef: &shared.ItemRef{
                        ID: "2ed02892-1cdd-4c69-a601-fb576b0d5f0d",
                        Name: lending.String("Jennifer Runolfsdottir"),
                    },
                    Quantity: types.MustNewDecimalFromString("7299.91"),
                    SubTotal: types.MustNewDecimalFromString("7499.99"),
                    TaxAmount: types.MustNewDecimalFromString("1716.29"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("3394.04"),
                        ID: lending.String("87053202-c73d-45fe-9b90-c28909b3fe49"),
                        Name: lending.String("Casey Stracke"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("7301.22"),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.RecordRef{
                            DataType: lending.String("transfer"),
                            ID: lending.String("48633323-f9b7-47f3-a410-0674ebf69280"),
                        },
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: lending.String("transfer"),
                                ID: lending.String("1ba77a89-ebf7-437a-a420-3ce5e6a95d8a"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0d446ce2-af7a-473c-b3be-453f870b326b",
                            Name: lending.String("Joanna Kohler"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("1864.58"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(false),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            Note: lending.String("maxime"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.AccountingPaymentAllocationAllocation{
                        AllocatedOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Currency: lending.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("1175.31"),
                        TotalAmount: types.MustNewDecimalFromString("6748.48"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: lending.String("8422bb67-9d23-4227-95bf-0cbb1e31b8b9"),
                            Name: lending.String("Dixie Durgan"),
                        },
                        Currency: lending.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("6772.63"),
                        ID: lending.String("1108e0ad-cf4b-4921-879f-ce953f73ef7f"),
                        Note: lending.String("distinctio"),
                        PaidOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Reference: lending.String("odio"),
                        TotalAmount: types.MustNewDecimalFromString("6304.48"),
                    },
                },
            },
            Reference: lending.String("facilis"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("4981.4"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolore": map[string]interface{}{
                        "quibusdam": "illum",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("1943.42"),
            TotalAmount: types.MustNewDecimalFromString("6178.77"),
        },
        AllowSyncOnPushComplete: lending.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: lending.Bool(false),
        TimeoutInMinutes: lending.Int(773326),
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

The *Get create direct cost model* endpoint returns the expected data for the request payload when creating a [direct cost](https://docs.codat.io/accounting-api#/schemas/DirectCost) for a given company and integration.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

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

