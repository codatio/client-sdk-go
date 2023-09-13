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
    res, err := s.AccountingCreditNotes.CreateAccountingCreditNote(ctx, operations.CreateAccountingCreditNoteRequest{
        AccountingCreditNote: &shared.AccountingCreditNote{
            AdditionalTaxAmount: types.MustNewDecimalFromString("3834.41"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("4776.65"),
            AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codatsynccommerce.String("placeat"),
            Currency: codatsynccommerce.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("4799.77"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("excepturi"),
                ID: "6ed151a0-5dfc-42dd-b7cc-78ca1ba928fc",
            },
            DiscountPercentage: *types.MustNewDecimalFromString("5218.48"),
            ID: codatsynccommerce.String("16742cb7-3920-4592-9396-fea7596eb10f"),
            IssueDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("aa2352c5-9559-407a-bf1a-3a2fa9467739"),
                        Name: codatsynccommerce.String("Beatrice Brown"),
                    },
                    Description: codatsynccommerce.String("enim"),
                    DiscountAmount: types.MustNewDecimalFromString("1381.83"),
                    DiscountPercentage: types.MustNewDecimalFromString("7783.46"),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3f5ad019-da1f-4fe7-8f09-7b0074f15471",
                        Name: codatsynccommerce.String("Bill Thompson"),
                    },
                    Quantity: *types.MustNewDecimalFromString("641.47"),
                    SubTotal: types.MustNewDecimalFromString("2168.22"),
                    TaxAmount: types.MustNewDecimalFromString("6924.72"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("5651.89"),
                        ID: codatsynccommerce.String("9d488e1e-91e4-450a-92ab-d44269802d50"),
                        Name: codatsynccommerce.String("Sonya Marks"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("7351.94"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "4f63c969-e9a3-4efa-b7df-b14cd66ae395",
                                Name: codatsynccommerce.String("Toby Pouros"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("id"),
                            ID: "88f3a669-9707-44ba-8469-b6e214195989",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "fa563e25-16fe-44c8-b711-e5b7fd2ed028",
                            Name: codatsynccommerce.String("Victor Casper"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("transfer"),
                            ID: codatsynccommerce.String("c692601f-b576-4b0d-9f0d-30c5fbb25870"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "53202c73-d5fe-49b9-8c28-909b3fe49a8d",
                            Name: codatsynccommerce.String("Loren Renner"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("5542.42"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("dolorem"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("1861.93"),
                        TotalAmount: types.MustNewDecimalFromString("2187.49"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("f9b77f3a-4100-4674-abf6-9280d1ba77a8"),
                            Name: codatsynccommerce.String("Terence Rau"),
                        },
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("4560.15"),
                        ID: codatsynccommerce.String("ae4203ce-5e6a-495d-8a0d-446ce2af7a73"),
                        Note: codatsynccommerce.String("quisquam"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("amet"),
                        TotalAmount: types.MustNewDecimalFromString("7308.56"),
                    },
                },
            },
            RemainingCredit: *types.MustNewDecimalFromString("8802.98"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: *types.MustNewDecimalFromString("2133.12"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sapiente": map[string]interface{}{
                        "totam": "nihil",
                    },
                },
            },
            TotalAmount: *types.MustNewDecimalFromString("256.62"),
            TotalDiscount: *types.MustNewDecimalFromString("7115.84"),
            TotalTaxAmount: *types.MustNewDecimalFromString("2074.7"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: *types.MustNewDecimalFromString("1536.94"),
                    Name: "Kelli Hintz",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(214880),
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

