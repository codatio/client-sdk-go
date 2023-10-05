# AccountingDirectIncomes
(*AccountingDirectIncomes*)

## Overview

Direct incomes

### Available Operations

* [CreateAccountingDirectIncome](#createaccountingdirectincome) - Create direct income

## CreateAccountingDirectIncome

The *Create direct income* endpoint creates a new [direct income](https://docs.codat.io/accounting-api#/schemas/DirectIncome) for a given company's connection.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale. For example, cash sales of items to a customer, referral commissions, and service fee refunds are considered direct incomes.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingDirectIncomes.CreateAccountingDirectIncome(ctx, operations.CreateAccountingDirectIncomeRequest{
        AccountingDirectIncome: &shared.AccountingDirectIncome{
            ContactRef: &shared.AccountingDirectIncomeContactRef{
                DataType: shared.DataTypeInvoices.ToPointer(),
                ID: "<ID>",
            },
            Currency: "GBP",
            CurrencyRate: types.MustNewDecimalFromString("6548.38"),
            ID: syncforcommerceversion1.String("<ID>"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("Meitnerium watt Assistant"),
                    },
                    Description: syncforcommerceversion1.String("Organized directional moratorium"),
                    DiscountAmount: types.MustNewDecimalFromString("6259.31"),
                    DiscountPercentage: types.MustNewDecimalFromString("3251.14"),
                    ItemRef: &shared.DirectIncomeLineItemItemReference{
                        ID: "<ID>",
                        Name: syncforcommerceversion1.String("Supervisor virtual sadly"),
                    },
                    Quantity: types.MustNewDecimalFromString("7605.88"),
                    SubTotal: types.MustNewDecimalFromString("2695.32"),
                    TaxAmount: types.MustNewDecimalFromString("3396.83"),
                    TaxRateRef: &shared.DirectIncomeLineItemTaxRateReference{
                        EffectiveTaxRate: types.MustNewDecimalFromString("2606.73"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("Pensacola Som Northwest"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6880.22"),
                    TrackingCategoryRefs: []shared.DirectIncomeLineItemTrackingCategoryRefs{
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "<ID>",
                            Name: syncforcommerceversion1.String("Pickup"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("8819.46"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("magni"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("2449.75"),
                        TotalAmount: types.MustNewDecimalFromString("6540.35"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("<ID>"),
                            Name: syncforcommerceversion1.String("Officer mobile Infrastructure"),
                        },
                        Currency: syncforcommerceversion1.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("4022.79"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Note: syncforcommerceversion1.String("Tools Dynamic Industrial"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("parallelism"),
                        TotalAmount: types.MustNewDecimalFromString("9002.17"),
                    },
                },
            },
            Reference: syncforcommerceversion1.String("Man Cotton virtual"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("4300.87"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "eos": map[string]interface{}{
                        "facere": "TLS",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("356.06"),
            TotalAmount: types.MustNewDecimalFromString("9132.9"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(156594),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingDirectIncomeRequest](../../models/operations/createaccountingdirectincomerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.CreateAccountingDirectIncomeResponse](../../models/operations/createaccountingdirectincomeresponse.md), error**

