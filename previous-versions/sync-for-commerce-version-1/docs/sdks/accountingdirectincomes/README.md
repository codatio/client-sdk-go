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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingDirectIncomes.CreateAccountingDirectIncome(ctx, operations.CreateAccountingDirectIncomeRequest{
        AccountingDirectIncome: &shared.AccountingDirectIncome{
            ContactRef: &shared.AccountingDirectIncomeContactRef{
                DataType: codatsynccommerce.String("tempora"),
                ID: "5626d436-813f-416d-9f5f-ce6c556146c3",
            },
            Currency: "EUR",
            CurrencyRate: types.MustNewDecimalFromString("1324.87"),
            ID: codatsynccommerce.String("50fb008c-42e1-441a-ac36-6c8dd6b14429"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("7474778a-7bd4-466d-a8c1-0ab3cdca4251"),
                        Name: codatsynccommerce.String("William Goodwin"),
                    },
                    Description: codatsynccommerce.String("aspernatur"),
                    DiscountAmount: types.MustNewDecimalFromString("1970.54"),
                    DiscountPercentage: types.MustNewDecimalFromString("7791.92"),
                    ItemRef: &shared.DirectIncomeLineItemItemReference{
                        ID: "7e0bc717-8e47-496f-aa70-c688282aa482",
                        Name: codatsynccommerce.String("Sue Corkery"),
                    },
                    Quantity: *types.MustNewDecimalFromString("1871.31"),
                    SubTotal: types.MustNewDecimalFromString("1294.12"),
                    TaxAmount: types.MustNewDecimalFromString("9039.84"),
                    TaxRateRef: &shared.DirectIncomeLineItemTaxRateReference{
                        EffectiveTaxRate: types.MustNewDecimalFromString("5789.22"),
                        ID: codatsynccommerce.String("817ee17c-be61-4e6b-bb95-bc0ab3c20c4f"),
                        Name: codatsynccommerce.String("Joy Labadie"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("8577.23"),
                    TrackingCategoryRefs: []shared.DirectIncomeLineItemTrackingCategoryRefs{
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "871f99dd-2efd-4121-aa6f-1e674bdb04f1",
                            Name: codatsynccommerce.String("Delores Hermiston IV"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("1852.32"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("ex"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("6802.7"),
                        TotalAmount: types.MustNewDecimalFromString("996.15"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("9f1d1705-1339-4d08-886a-1840394c2607"),
                            Name: codatsynccommerce.String("Elisa Mosciski"),
                        },
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("9903.45"),
                        ID: codatsynccommerce.String("0642dac7-af51-45cc-813a-a63aae8d6786"),
                        Note: codatsynccommerce.String("labore"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("facilis"),
                        TotalAmount: types.MustNewDecimalFromString("7382.27"),
                    },
                },
            },
            Reference: codatsynccommerce.String("commodi"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SubTotal: *types.MustNewDecimalFromString("3605.45"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "reiciendis": map[string]interface{}{
                        "assumenda": "nemo",
                    },
                },
            },
            TaxAmount: *types.MustNewDecimalFromString("9249.67"),
            TotalAmount: *types.MustNewDecimalFromString("3975.33"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(46007),
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

