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
                DataType: codatsynccommerce.String("nisi"),
                ID: "f48b656b-cdb3-45ff-ae4b-27537a8cd9e7",
            },
            Currency: "GBP",
            CurrencyRate: codatsynccommerce.Float64(1180.41),
            ID: codatsynccommerce.String("9c177d52-5f77-4b11-8eeb-52ff785fc378"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("d4c98e0c-2bb8-49eb-b5da-d636c600503d"),
                        Name: codatsynccommerce.String("Willis Rippin Sr."),
                    },
                    Description: codatsynccommerce.String("laudantium"),
                    DiscountAmount: codatsynccommerce.Float64(407.1),
                    DiscountPercentage: codatsynccommerce.Float64(9384.12),
                    ItemRef: &shared.DirectIncomeLineItemItemReference{
                        ID: "739ae9e0-57eb-4809-a281-0331f3981d4c",
                        Name: codatsynccommerce.String("Donna Aufderhar"),
                    },
                    Quantity: 467.91,
                    SubTotal: codatsynccommerce.Float64(4894.59),
                    TaxAmount: codatsynccommerce.Float64(9980.26),
                    TaxRateRef: &shared.DirectIncomeLineItemTaxRateReference{
                        EffectiveTaxRate: codatsynccommerce.Float64(2436.78),
                        ID: codatsynccommerce.String("c93c73b9-da3f-42ce-9a7e-23f2257411fa"),
                        Name: codatsynccommerce.String("Kyle Reichel"),
                    },
                    TotalAmount: codatsynccommerce.Float64(2883),
                    TrackingCategoryRefs: []shared.DirectIncomeLineItemTrackingCategoryRefs{
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "e472e802-857a-45b4-8463-a7d575f1400e",
                            Name: codatsynccommerce.String("Gertrude Gerhold"),
                        },
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "7334ec1b-781b-436a-8808-8d100efada20",
                            Name: codatsynccommerce.String("Mrs. Olive Weissnat"),
                        },
                    },
                    UnitAmount: 1858.97,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("eb2164cf-9ab8-4366-8723-ffda9e06bee4"),
                        Name: codatsynccommerce.String("Howard Hermann DVM"),
                    },
                    Description: codatsynccommerce.String("eligendi"),
                    DiscountAmount: codatsynccommerce.Float64(620.35),
                    DiscountPercentage: codatsynccommerce.Float64(8850.22),
                    ItemRef: &shared.DirectIncomeLineItemItemReference{
                        ID: "115c80bf-f918-4544-ac42-defcce8f1977",
                        Name: codatsynccommerce.String("Lydia Douglas"),
                    },
                    Quantity: 2086.83,
                    SubTotal: codatsynccommerce.Float64(3577.58),
                    TaxAmount: codatsynccommerce.Float64(3753.5),
                    TaxRateRef: &shared.DirectIncomeLineItemTaxRateReference{
                        EffectiveTaxRate: codatsynccommerce.Float64(1636.84),
                        ID: codatsynccommerce.String("a7b408f0-5e3d-448f-9af3-13a1f5fd9425"),
                        Name: codatsynccommerce.String("Lucas Abbott"),
                    },
                    TotalAmount: codatsynccommerce.Float64(4124.33),
                    TrackingCategoryRefs: []shared.DirectIncomeLineItemTrackingCategoryRefs{
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "25ea944f-3b75-46c1-9f6c-37a512624383",
                            Name: codatsynccommerce.String("Kristy Quigley II"),
                        },
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "a23a45ce-fc5f-4de1-8a0c-e2169e510019",
                            Name: codatsynccommerce.String("Elmer Spinka"),
                        },
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "e3476279-9bfb-4be6-949f-b2bb4ecae6c3",
                            Name: codatsynccommerce.String("Maurice Stanton"),
                        },
                        shared.DirectIncomeLineItemTrackingCategoryRefs{
                            ID: "adebd5da-ea4c-4506-a8aa-94c02644cf5e",
                            Name: codatsynccommerce.String("Drew Mraz"),
                        },
                    },
                    UnitAmount: 3442.89,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("corrupti"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(986.1),
                        TotalAmount: codatsynccommerce.Float64(6472.18),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("c600dec0-01ac-4802-a2ec-09ff8f0f816f"),
                            Name: codatsynccommerce.String("Lee Gibson"),
                        },
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(946.97),
                        ID: codatsynccommerce.String("3e902c14-125b-4096-8a66-8151a472af92"),
                        Note: codatsynccommerce.String("nesciunt"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("quis"),
                        TotalAmount: codatsynccommerce.Float64(5861.08),
                    },
                },
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(9804.1),
                        TotalAmount: codatsynccommerce.Float64(5120.81),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("3f350cf8-76ff-4b90-9c6e-cbb4e243cf78"),
                            Name: codatsynccommerce.String("Jerald Wilkinson"),
                        },
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(8615.91),
                        ID: codatsynccommerce.String("a53e5ae6-e0ac-4184-82b9-c247c88373a4"),
                        Note: codatsynccommerce.String("perferendis"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("architecto"),
                        TotalAmount: codatsynccommerce.Float64(5646.27),
                    },
                },
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(9829.91),
                        TotalAmount: codatsynccommerce.Float64(2050.11),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("2e550557-56f5-4d56-90bd-0af2dfe13db4"),
                            Name: codatsynccommerce.String("Elmer Champlin"),
                        },
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(2425.31),
                        ID: codatsynccommerce.String("f8941aeb-c0b8-40a6-924d-3b2ecfcc8f89"),
                        Note: codatsynccommerce.String("corporis"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("illo"),
                        TotalAmount: codatsynccommerce.Float64(142.51),
                    },
                },
            },
            Reference: codatsynccommerce.String("doloribus"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 8698.48,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "neque": map[string]interface{}{
                        "vel": "sapiente",
                        "mollitia": "quae",
                        "quos": "aperiam",
                        "non": "voluptates",
                    },
                    "ad": map[string]interface{}{
                        "quisquam": "quas",
                        "consequuntur": "maiores",
                    },
                    "inventore": map[string]interface{}{
                        "laudantium": "est",
                        "dolor": "aliquid",
                    },
                    "consectetur": map[string]interface{}{
                        "rem": "voluptatum",
                        "ducimus": "adipisci",
                        "recusandae": "tempora",
                        "blanditiis": "numquam",
                    },
                },
            },
            TaxAmount: 1963.74,
            TotalAmount: 5323.2,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(27078),
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

