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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"log"
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
            AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Currency: syncforcommerceversion1.String("GBP"),
            CustomerRef: &shared.AccountingCustomerRef{
                ID: "<ID>",
            },
            DiscountPercentage: types.MustNewDecimalFromString("3961.39"),
            IssueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{},
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                    },
                    Quantity: types.MustNewDecimalFromString("1740.95"),
                    TaxRateRef: &shared.TaxRateRef{},
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefItems{
                            shared.TrackingCategoryRefItems{
                                ID: "<ID>",
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.CreditNoteLineItemAccountingProjectReference{
                            ID: "<ID>",
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("accountTransaction"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefItems{
                        shared.TrackingCategoryRefItems{
                            ID: "<ID>",
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("6472.07"),
                },
            },
            Metadata: &shared.Metadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.PaymentAllocationItems{
                shared.PaymentAllocationItems{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("EUR"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{},
                        Currency: syncforcommerceversion1.String("USD"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("3693.65"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusVoid,
            SubTotal: types.MustNewDecimalFromString("1915.04"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "key": map[string]interface{}{
                        "key": "string",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("5893.9"),
            TotalDiscount: types.MustNewDecimalFromString("579.23"),
            TotalTaxAmount: types.MustNewDecimalFromString("3881.42"),
            WithholdingTax: []shared.WithholdingTaxItems{
                shared.WithholdingTaxItems{
                    Amount: types.MustNewDecimalFromString("7369.44"),
                    Name: "string",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingCreditNoteRequest](../../pkg/models/operations/createaccountingcreditnoterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.CreateAccountingCreditNoteResponse](../../pkg/models/operations/createaccountingcreditnoteresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
