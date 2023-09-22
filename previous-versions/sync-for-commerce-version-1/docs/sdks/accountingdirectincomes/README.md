# AccountingDirectIncomes

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
                DataType: syncforcommerceversion1.String("nostrum"),
                ID: "fce6c556-146c-43e2-90fb-008c42e141aa",
            },
            Currency: "EUR",
            CurrencyRate: types.MustNewDecimalFromString("2453.67"),
            ID: syncforcommerceversion1.String("66c8dd6b-1442-4907-8747-78a7bd466d28"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("10ab3cdc-a425-4190-8e52-3c7e0bc7178e"),
                        Name: syncforcommerceversion1.String("Courtney Mayert"),
                    },
                    Description: syncforcommerceversion1.String("dolores"),
                    DiscountAmount: types.MustNewDecimalFromString("6455.7"),
                    DiscountPercentage: types.MustNewDecimalFromString("4752.89"),
                    ItemRef: &shared.DirectIncomeLineItemItemReference{
                        ID: "0c688282-aa48-4256-af22-2e9817ee17cb",
                        Name: syncforcommerceversion1.String("Gene Botsford"),
                    },
                    Quantity: types.MustNewDecimalFromString("6900.25"),
                    SubTotal: types.MustNewDecimalFromString("4732.21"),
                    TaxAmount: types.MustNewDecimalFromString("6996.22"),
                    TaxRateRef: &shared.DirectIncomeLineItemTaxRateReference{
                        EffectiveTaxRate: types.MustNewDecimalFromString("5801.97"),
                        ID: syncforcommerceversion1.String("5bc0ab3c-20c4-4f37-89fd-871f99dd2efd"),
                        Name: syncforcommerceversion1.String("Marilyn Botsford"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("3984.34"),
                    TrackingCategoryRefs: []shared.DirectIncomeLineItemTrackingCategoryRefs{
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "f1e674bd-b04f-4157-9608-2d68ea19f1d1",
                            Name: syncforcommerceversion1.String("Mrs. Cynthia Hansen"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("6144.65"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("accusantium"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("5130.75"),
                        TotalAmount: types.MustNewDecimalFromString("4287.96"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("a1840394-c260-471f-93f5-f0642dac7af5"),
                            Name: syncforcommerceversion1.String("Darlene Sawayn"),
                        },
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("2414.18"),
                        ID: syncforcommerceversion1.String("aa63aae8-d678-464d-bb67-5fd5e60b375e"),
                        Note: syncforcommerceversion1.String("facere"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("doloribus"),
                        TotalAmount: types.MustNewDecimalFromString("3817.6"),
                    },
                },
            },
            Reference: syncforcommerceversion1.String("reiciendis"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            SubTotal: types.MustNewDecimalFromString("9049.49"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "necessitatibus": map[string]interface{}{
                        "dolore": "sunt",
                    },
                },
            },
            TaxAmount: types.MustNewDecimalFromString("9920.12"),
            TotalAmount: types.MustNewDecimalFromString("2415.45"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(249420),
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

