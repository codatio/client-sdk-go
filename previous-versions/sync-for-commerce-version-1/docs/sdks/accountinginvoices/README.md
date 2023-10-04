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
            AdditionalTaxAmount: types.MustNewDecimalFromString("9907.57"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("9015.63"),
            AmountDue: types.MustNewDecimalFromString("195.45"),
            Currency: syncforcommerceversion1.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("1021.57"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: syncforcommerceversion1.String("Runte Inc"),
                ID: "<ID>",
            },
            DiscountPercentage: types.MustNewDecimalFromString("7432.38"),
            DueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            ID: syncforcommerceversion1.String("<ID>"),
            InvoiceNumber: syncforcommerceversion1.String("Manors"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("as"),
                    },
                    Description: syncforcommerceversion1.String("Visionary discrete task-force"),
                    DiscountAmount: types.MustNewDecimalFromString("1010.92"),
                    DiscountPercentage: types.MustNewDecimalFromString("6455.29"),
                    IsDirectIncome: syncforcommerceversion1.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: syncforcommerceversion1.String("Shoes Tennessee"),
                    },
                    Quantity: types.MustNewDecimalFromString("8154.23"),
                    SubTotal: types.MustNewDecimalFromString("7362.43"),
                    TaxAmount: types.MustNewDecimalFromString("6235.41"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("9166.8"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Name: syncforcommerceversion1.String("yellow Chair"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("1049.23"),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "<ID>",
                                Name: syncforcommerceversion1.String("rural Bulgarian Producer"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: syncforcommerceversion1.String("Grimes, Yost and Champlin"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforcommerceversion1.String("Organized UDP"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("journalEntry"),
                            ID: syncforcommerceversion1.String("<ID>"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "<ID>",
                            Name: syncforcommerceversion1.String("Garden"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("5528.53"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("Home Applications Fermium"),
            PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingInvoicePaymentAllocation{
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("4747.86"),
                        TotalAmount: types.MustNewDecimalFromString("8682.37"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("<ID>"),
                            Name: syncforcommerceversion1.String("West Avon Herzegovina"),
                        },
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("7606.26"),
                        ID: syncforcommerceversion1.String("<ID>"),
                        Note: syncforcommerceversion1.String("Berkshire"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("Palm scam"),
                        TotalAmount: types.MustNewDecimalFromString("2223.49"),
                    },
                },
            },
            SalesOrderRefs: []shared.AccountingInvoiceSalesOrderReference{
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: shared.DataTypeInvoices.ToPointer(),
                    ID: syncforcommerceversion1.String("<ID>"),
                },
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusUnknown,
            SubTotal: types.MustNewDecimalFromString("7559.48"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "explicabo": map[string]interface{}{
                        "fugiat": "definition",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("2429.55"),
            TotalDiscount: types.MustNewDecimalFromString("6497.09"),
            TotalTaxAmount: types.MustNewDecimalFromString("337.26"),
            WithholdingTax: []shared.AccountingInvoiceWithholdingTax{
                shared.AccountingInvoiceWithholdingTax{
                    Amount: types.MustNewDecimalFromString("676.3"),
                    Name: "whether Division so",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(658558),
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

