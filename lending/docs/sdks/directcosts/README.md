# DirectCosts
(*LoanWriteback.DirectCosts*)

## Overview

### Available Operations

* [Create](#create) - Create direct cost
* [GetCreateModel](#getcreatemodel) - Get create direct cost model

## Create

The *Create direct cost* endpoint creates a new [direct cost](https://docs.codat.io/lending-api#/schemas/DirectCost) for a given company's connection.

[Direct costs](https://docs.codat.io/lending-api#/schemas/DirectCost) are the expenses associated with a business' operations. For example, purchases of raw materials that are paid off at the point of the purchase and service fees are considered direct costs.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/lending-api#/operations/get-create-directCosts-model).

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
    res, err := s.LoanWriteback.DirectCosts.Create(ctx, operations.CreateDirectCostRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostPrototype: &shared.DirectCostPrototype{
            ContactRef: &shared.ContactRef{
                DataType: shared.ContactRefDataTypeSuppliers.ToPointer(),
                ID: "80000001-1671793885",
            },
            Currency: "USD",
            IssueDate: "2023-03-21T10:19:52.223Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: lending.String("8000000D-1671793811"),
                        Name: lending.String("Purchases - Hardware for Resale"),
                    },
                    Description: lending.String("test description line 1"),
                    DiscountAmount: types.MustNewDecimalFromString("0"),
                    DiscountPercentage: types.MustNewDecimalFromString("0"),
                    ItemRef: &shared.PropertieItemRef{
                        ID: "80000001-1674566705",
                        Name: lending.String("item test"),
                    },
                    Quantity: types.MustNewDecimalFromString("1"),
                    SubTotal: types.MustNewDecimalFromString("99"),
                    TaxAmount: types.MustNewDecimalFromString("360"),
                    TotalAmount: types.MustNewDecimalFromString("70"),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.RecordRef{
                            DataType: lending.String("invoice"),
                        },
                        RecordRefs: []shared.TrackingRecordRef{
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeTrackingCategories.ToPointer(),
                            },
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeTrackingCategories.ToPointer(),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "80000001-1674553252",
                            Name: lending.String("Class 1"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7"),
                },
            },
            Note: lending.String("directCost 21/03 09.20"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: lending.String("2023-01-29T10:19:52.223Z"),
                        Currency: lending.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("0"),
                        TotalAmount: types.MustNewDecimalFromString("88"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: lending.String("80000028-1671794219"),
                            Name: lending.String("Bank Account 1"),
                        },
                        Currency: lending.String("GBP"),
                        Note: lending.String("payment allocations note"),
                        PaidOnDate: lending.String("2023-01-28T10:19:52.223Z"),
                        Reference: lending.String("payment allocations reference"),
                        TotalAmount: types.MustNewDecimalFromString("54"),
                    },
                },
            },
            Reference: lending.String("test ref"),
            SubTotal: types.MustNewDecimalFromString("362"),
            TaxAmount: types.MustNewDecimalFromString("4"),
            TotalAmount: types.MustNewDecimalFromString("366"),
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateDirectCostRequest](../../pkg/models/operations/createdirectcostrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateDirectCostResponse](../../pkg/models/operations/createdirectcostresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetCreateModel

The *Get create direct cost model* endpoint returns the expected data for the request payload when creating a [direct cost](https://docs.codat.io/lending-api#/schemas/DirectCost) for a given company and integration.

[Direct costs](https://docs.codat.io/lending-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCreateDirectCostsModelRequest](../../pkg/models/operations/getcreatedirectcostsmodelrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCreateDirectCostsModelResponse](../../pkg/models/operations/getcreatedirectcostsmodelresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |