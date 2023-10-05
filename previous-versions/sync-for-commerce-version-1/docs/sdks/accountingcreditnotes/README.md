# AccountingCreditNotes
(*AccountingCreditNotes*)

## Overview

Credit notes

### Available Operations

* [CreateAccountingCreditNote](#createaccountingcreditnote) - Create credit note

## CreateAccountingCreditNote

The *Create credit note* endpoint creates a new [credit note](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating an account.


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
    res, err := s.AccountingCreditNotes.CreateAccountingCreditNote(ctx, operations.CreateAccountingCreditNoteRequest{
        AccountingCreditNote: &shared.AccountingCreditNote{
            AdditionalTaxAmount: types.MustNewDecimalFromString("6733.79"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("612.72"),
            AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: syncforcommerceversion1.String("lavender Planner"),
            Currency: syncforcommerceversion1.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("2314.32"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: syncforcommerceversion1.String("Swaniawski - Okuneva"),
                ID: "<ID>",
            },
            DiscountPercentage: types.MustNewDecimalFromString("1210.61"),
            ID: syncforcommerceversion1.String("<ID>"),
            IssueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("candela Metal policy"),
                    },
                    Description: syncforcommerceversion1.String("Universal 4th generation model"),
                    DiscountAmount: types.MustNewDecimalFromString("6593.55"),
                    DiscountPercentage: types.MustNewDecimalFromString("3629.12"),
                    IsDirectIncome: syncforcommerceversion1.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: syncforcommerceversion1.String("Titusville Car"),
                    },
                    Quantity: types.MustNewDecimalFromString("9339.43"),
                    SubTotal: types.MustNewDecimalFromString("9776.2"),
                    TaxAmount: types.MustNewDecimalFromString("4570.33"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("2840.32"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("Nissan Shirt"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("3862.17"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "<ID>",
                                Name: syncforcommerceversion1.String("system"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: syncforcommerceversion1.String("Labadie and Sons"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforcommerceversion1.String("Mann second siemens"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("transfer"),
                            ID: syncforcommerceversion1.String("<ID>"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "<ID>",
                            Name: syncforcommerceversion1.String("scalable"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("9063.02"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("Conroy fuzzy Mobility"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("5905.56"),
                        TotalAmount: types.MustNewDecimalFromString("8276.36"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("<ID>"),
                            Name: syncforcommerceversion1.String("synthesizing becquerel Operations"),
                        },
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("8697.42"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Note: syncforcommerceversion1.String("Convertible"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("Grenada networks Fantastic"),
                        TotalAmount: types.MustNewDecimalFromString("5575.36"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("4029.48"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: types.MustNewDecimalFromString("8342.79"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quos": map[string]interface{}{
                        "nihil": "Concrete",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("1868.28"),
            TotalDiscount: types.MustNewDecimalFromString("4123.14"),
            TotalTaxAmount: types.MustNewDecimalFromString("8369.46"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("879.63"),
                    Name: "male Bedfordshire architectures",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(814888),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateAccountingCreditNoteRequest](../../models/operations/createaccountingcreditnoterequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.CreateAccountingCreditNoteResponse](../../models/operations/createaccountingcreditnoteresponse.md), error**

