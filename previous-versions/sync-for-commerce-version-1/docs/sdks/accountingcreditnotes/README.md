# AccountingCreditNotes
(*AccountingCreditNotes*)

## Overview

Retrieve standardized Credit notes from linked accounting software.

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
            AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            CreditNoteNumber: syncforcommerceversion1.String("09/03 17.15"),
            Currency: syncforcommerceversion1.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("1"),
            CustomerRef: &shared.AccountingCustomerRef{
                ID: "80000002-1674552702",
            },
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            IssueDate: syncforcommerceversion1.String("2023-03-09T02:21:26.726327+00:00"),
            LineItems: []shared.CreditNoteLineItem{

            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            Note: syncforcommerceversion1.String("credit note 20230309 17.15"),
            PaymentAllocations: []shared.PaymentAllocationItems{
                shared.PaymentAllocationItems{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
                        Currency: syncforcommerceversion1.String("USD"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        Currency: syncforcommerceversion1.String("EUR"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            Status: shared.CreditNoteStatusSubmitted,
            SubTotal: types.MustNewDecimalFromString("10.2"),
            TotalAmount: types.MustNewDecimalFromString("1.25"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
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

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
