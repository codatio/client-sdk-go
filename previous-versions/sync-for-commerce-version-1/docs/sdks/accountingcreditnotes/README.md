# AccountingCreditNotes

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
            AdditionalTaxAmount: types.MustNewDecimalFromString("3927.85"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("9255.97"),
            AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: syncforcommerceversion1.String("ab"),
            Currency: syncforcommerceversion1.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("871.29"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: syncforcommerceversion1.String("deserunt"),
                ID: "05dfc2dd-f7cc-478c-a1ba-928fc816742c",
            },
            DiscountPercentage: types.MustNewDecimalFromString("7369.18"),
            ID: syncforcommerceversion1.String("73920592-9396-4fea-b596-eb10faaa2352"),
            IssueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("5955907a-ff1a-43a2-ba94-67739251aa52"),
                        Name: syncforcommerceversion1.String("Jimmy Wiegand"),
                    },
                    Description: syncforcommerceversion1.String("possimus"),
                    DiscountAmount: types.MustNewDecimalFromString("135.71"),
                    DiscountPercentage: types.MustNewDecimalFromString("971.01"),
                    IsDirectIncome: syncforcommerceversion1.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9da1ffe7-8f09-47b0-874f-15471b5e6e13",
                        Name: syncforcommerceversion1.String("Virgil Mante"),
                    },
                    Quantity: types.MustNewDecimalFromString("5089.69"),
                    SubTotal: types.MustNewDecimalFromString("5232.48"),
                    TaxAmount: types.MustNewDecimalFromString("9167.23"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("939.4"),
                        ID: syncforcommerceversion1.String("e91e450a-d2ab-4d44-a698-02d502a94bb4"),
                        Name: syncforcommerceversion1.String("Andre Franey"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("3960.98"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "9e9a3efa-77df-4b14-8d66-ae395efb9ba8",
                                Name: syncforcommerceversion1.String("Timmy Feeney"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: syncforcommerceversion1.String("vel"),
                            ID: "997074ba-4469-4b6e-a141-959890afa563",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "516fe4c8-b711-4e5b-bfd2-ed028921cddc",
                            Name: syncforcommerceversion1.String("Miriam Connelly Jr."),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("transfer"),
                            ID: syncforcommerceversion1.String("b576b0d5-f0d3-40c5-bbb2-587053202c73"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "d5fe9b90-c289-409b-bfe4-9a8d9cbf4863",
                            Name: syncforcommerceversion1.String("Rosa Dibbert"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("5695.74"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("voluptate"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("2274.14"),
                        TotalAmount: types.MustNewDecimalFromString("6805.45"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("4100674e-bf69-4280-91ba-77a89ebf737a"),
                            Name: syncforcommerceversion1.String("Mrs. Ray Collins"),
                        },
                        Currency: syncforcommerceversion1.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("3200.17"),
                        ID: syncforcommerceversion1.String("e6a95d8a-0d44-46ce-aaf7-a73cf3be453f"),
                        Note: syncforcommerceversion1.String("totam"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("sit"),
                        TotalAmount: types.MustNewDecimalFromString("7115.84"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("2074.7"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusSubmitted,
            SubTotal: types.MustNewDecimalFromString("7304.42"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptas": map[string]interface{}{
                        "deserunt": "quam",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("2148.8"),
            TotalDiscount: types.MustNewDecimalFromString("2776.28"),
            TotalTaxAmount: types.MustNewDecimalFromString("1864.58"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("5867.84"),
                    Name: "Miss Jody Rogahn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(276894),
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

