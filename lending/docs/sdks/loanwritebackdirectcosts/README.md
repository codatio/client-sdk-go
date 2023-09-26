# LoanWriteback.DirectCosts

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
                ID: "02d502a9-4bb4-4f63-8969-e9a3efa77dfb",
            },
            Currency: "GBP",
            CurrencyRate: types.MustNewDecimalFromString("2974.37"),
            ID: lending.String("cd66ae39-5efb-49ba-88f3-a66997074ba4"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: lending.String("69b6e214-1959-4890-afa5-63e2516fe4c8"),
                        Name: lending.String("Dr. Arnold Bradtke"),
                    },
                    Description: lending.String("expedita"),
                    DiscountAmount: types.MustNewDecimalFromString("4692.49"),
                    DiscountPercentage: types.MustNewDecimalFromString("9988.48"),
                    ItemRef: &shared.ItemRef{
                        ID: "d2ed0289-21cd-4dc6-9260-1fb576b0d5f0",
                        Name: lending.String("Vincent Anderson"),
                    },
                    Quantity: types.MustNewDecimalFromString("9441.24"),
                    SubTotal: types.MustNewDecimalFromString("7299.91"),
                    TaxAmount: types.MustNewDecimalFromString("7499.99"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("1716.29"),
                        ID: lending.String("58705320-2c73-4d5f-a9b9-0c28909b3fe4"),
                        Name: lending.String("Omar Leuschke"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("7508.44"),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.RecordRef{
                            DataType: lending.String("accountTransaction"),
                            ID: lending.String("f4863332-3f9b-477f-ba41-00674ebf6928"),
                        },
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: lending.String("journalEntry"),
                                ID: lending.String("d1ba77a8-9ebf-4737-ae42-03ce5e6a95d8"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a0d446ce-2af7-4a73-8f3b-e453f870b326",
                            Name: lending.String("Glen Oberbrunner"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("2776.28"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: lending.Bool(false),
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            Note: lending.String("cupiditate"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.AccountingPaymentAllocationAllocation{
                        AllocatedOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Currency: lending.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("7470.8"),
                        TotalAmount: types.MustNewDecimalFromString("1175.31"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: lending.String("a8422bb6-79d2-4322-b15b-f0cbb1e31b8b"),
                            Name: lending.String("Kevin Willms"),
                        },
                        Currency: lending.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("2408.29"),
                        ID: lending.String("a1108e0a-dcf4-4b92-9879-fce953f73ef7"),
                        Note: lending.String("hic"),
                        PaidOnDate: lending.String("2022-10-23T00:00:00.000Z"),
                        Reference: lending.String("quod"),
                        TotalAmount: types.MustNewDecimalFromString("4861.6"),
                    },
                },
            },
            Reference: lending.String("similique"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("8742.88"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ducimus": map[string]interface{}{
                        "dolore": "quibusdam",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("8489.44"),
            TotalAmount: types.MustNewDecimalFromString("1943.42"),
        },
        AllowSyncOnPushComplete: lending.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: lending.Int(617877),
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

