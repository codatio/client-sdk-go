# AccountingInvoices
(*AccountingInvoices*)

## Overview

Invoices

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
    res, err := s.AccountingInvoices.CreateAccountingInvoice(ctx, operations.CreateAccountingInvoiceRequest{
        AccountingInvoice: &shared.AccountingInvoice{
            AmountDue: types.MustNewDecimalFromString("9907.57"),
            Currency: syncforcommerceversion1.String("EUR"),
            CustomerRef: &shared.AccountingCustomerRef{
                ID: "<ID>",
            },
            DueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{},
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                    },
                    Quantity: types.MustNewDecimalFromString("1021.57"),
                    TaxRateRef: &shared.TaxRateRef{},
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "<ID>",
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "<ID>",
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("journalEntry"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "<ID>",
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7432.38"),
                },
            },
            Metadata: &shared.Metadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingInvoicePaymentAllocation{
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("USD"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{},
                        Currency: syncforcommerceversion1.String("EUR"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                    },
                },
            },
            SalesOrderRefs: []shared.AccountingInvoiceSalesOrderReference{
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: shared.DataTypeInvoices.ToPointer(),
                },
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusPartiallyPaid,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "key": map[string]interface{}{
                        "key": "string",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("1416.23"),
            TotalTaxAmount: types.MustNewDecimalFromString("9069.87"),
            WithholdingTax: []shared.AccountingInvoiceWithholdingTax{
                shared.AccountingInvoiceWithholdingTax{
                    Amount: types.MustNewDecimalFromString("598.23"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAccountingInvoiceRequest](../../models/operations/createaccountinginvoicerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.CreateAccountingInvoiceResponse](../../models/operations/createaccountinginvoiceresponse.md), error**

