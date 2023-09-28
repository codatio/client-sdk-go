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
                ID: "502a94bb-4f63-4c96-9e9a-3efa77dfb14c",
            },
            Currency: "EUR",
            CurrencyRate: types.MustNewDecimalFromString("4118.2"),
            ID: lending.String("6ae395ef-b9ba-488f-ba66-997074ba4469"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: lending.String("6e214195-9890-4afa-963e-2516fe4c8b71"),
                        Name: lending.String("Elvira Herman"),
                    },
                    Description: lending.String("repellat"),
                    DiscountAmount: types.MustNewDecimalFromString("8411.4"),
                    DiscountPercentage: types.MustNewDecimalFromString("1494.48"),
                    ItemRef: &shared.ItemRef{
                        ID: "ed028921-cddc-4692-a01f-b576b0d5f0d3",
                        Name: lending.String("Erma Hessel"),
                    },
                    Quantity: types.MustNewDecimalFromString("7499.99"),
                    SubTotal: types.MustNewDecimalFromString("1716.29"),
                    TaxAmount: types.MustNewDecimalFromString("3394.04"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("5210.37"),
                        ID: lending.String("7053202c-73d5-4fe9-b90c-28909b3fe49a"),
                        Name: lending.String("Ervin McLaughlin"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("9644.9"),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.RecordRef{
                            DataType: lending.String("invoice"),
                            ID: lending.String("8633323f-9b77-4f3a-8100-674ebf69280d"),
                        },
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: lending.String("journalEntry"),
                                ID: lending.String("ba77a89e-bf73-47ae-8203-ce5e6a95d8a0"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d446ce2a-f7a7-43cf-bbe4-53f870b326b5",
                            Name: lending.String("Darryl Emard"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("5867.84"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(false),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            Note: lending.String("pariatur"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.AccountingPaymentAllocationAllocation{
                        AllocatedOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Currency: lending.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("6748.48"),
                        TotalAmount: types.MustNewDecimalFromString("5173.79"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: lending.String("422bb679-d232-4271-9bf0-cbb1e31b8b90"),
                            Name: lending.String("Mike Greenholt"),
                        },
                        Currency: lending.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("1002.94"),
                        ID: lending.String("108e0adc-f4b9-4218-b9fc-e953f73ef7fb"),
                        Note: lending.String("quod"),
                        PaidOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Reference: lending.String("similique"),
                        TotalAmount: types.MustNewDecimalFromString("7085.48"),
                    },
                },
            },
            Reference: lending.String("vero"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("2930.2"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quibusdam": map[string]interface{}{
                        "illum": "sequi",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("6178.77"),
            TotalAmount: types.MustNewDecimalFromString("7733.26"),
        },
        AllowSyncOnPushComplete: lending.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: lending.Int(13236),
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

