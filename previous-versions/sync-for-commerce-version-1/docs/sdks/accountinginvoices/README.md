# AccountingInvoices
(*AccountingInvoices*)

## Overview

Retrieve standardized Invoices from linked accounting software.

### Available Operations

* [CreateAccountingInvoice](#createaccountinginvoice) - Create invoice

## CreateAccountingInvoice

The *Create invoice* endpoint creates a new [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) for a given company's connection.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support creating an account.


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
    res, err := s.AccountingInvoices.CreateAccountingInvoice(ctx, operations.CreateAccountingInvoiceRequest{
        AccountingInvoice: &shared.AccountingInvoice{
            AdditionalTaxAmount: types.MustNewDecimalFromString("0"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("0"),
            AmountDue: types.MustNewDecimalFromString("87326532"),
            Currency: syncforcommerceversion1.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("1"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: syncforcommerceversion1.String("Test Customer 1"),
                ID: "80000002-1674552702",
            },
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            DueDate: syncforcommerceversion1.String("2023-05-24T11:09:01.438Z"),
            InvoiceNumber: syncforcommerceversion1.String("18/04 15.26"),
            IssueDate: "2023-04-18T11:09:01.438Z",
            LineItems: []shared.InvoiceLineItem{

            },
            ModifiedDate: syncforcommerceversion1.String("2023-02-14T11:09:01.438Z"),
            Note: syncforcommerceversion1.String("invoice push 20230418 15.26"),
            PaidOnDate: syncforcommerceversion1.String("2023-02-10T11:09:01.438Z"),
            PaymentAllocations: []shared.AccountingPaymentAllocation{
                shared.AccountingPaymentAllocation{
                    Allocation: shared.AccountingInvoiceAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2023-02-14T11:09:01.438Z"),
                        Currency: syncforcommerceversion1.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("1"),
                        TotalAmount: types.MustNewDecimalFromString("725"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("string"),
                            Name: syncforcommerceversion1.String("string"),
                        },
                        Currency: syncforcommerceversion1.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("1"),
                        ID: syncforcommerceversion1.String("80000004-1789341990"),
                        Note: syncforcommerceversion1.String("string"),
                        PaidOnDate: syncforcommerceversion1.String("2023-02-14T11:09:01.438Z"),
                        Reference: syncforcommerceversion1.String("string"),
                        TotalAmount: types.MustNewDecimalFromString("725"),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderReference{
                shared.SalesOrderReference{},
            },
            SourceModifiedDate: syncforcommerceversion1.String("2023-02-14T11:09:01.438Z"),
            Status: shared.InvoiceStatusSubmitted,
            SubTotal: types.MustNewDecimalFromString("30"),
            TotalAmount: types.MustNewDecimalFromString("30"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTax{
                shared.WithholdingTax{
                    Amount: types.MustNewDecimalFromString("0"),
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
    if res.AccountingCreateInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAccountingInvoiceRequest](../../pkg/models/operations/createaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAccountingInvoiceResponse](../../pkg/models/operations/createaccountinginvoiceresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
