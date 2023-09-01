# AccountingInvoices

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
    res, err := s.AccountingInvoices.CreateAccountingInvoice(ctx, operations.CreateAccountingInvoiceRequest{
        AccountingInvoice: &shared.AccountingInvoice{
            AdditionalTaxAmount: codatsynccommerce.Float64(7031.89),
            AdditionalTaxPercentage: codatsynccommerce.Float64(841.78),
            AmountDue: 9492.8,
            Currency: codatsynccommerce.String("USD"),
            CurrencyRate: codatsynccommerce.Float64(6940.88),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("totam"),
                ID: "ca275a60-a04c-4495-8c69-9171b51c1bdb",
            },
            DiscountPercentage: codatsynccommerce.Float64(1053.72),
            DueDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            ID: codatsynccommerce.String("f4b888eb-dfc4-4ccc-a99b-c7fc0b2dce10"),
            InvoiceNumber: codatsynccommerce.String("laudantium"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("e42b006d-6788-478b-a858-1a58208c54fe"),
                        Name: codatsynccommerce.String("Gerard Mraz"),
                    },
                    Description: codatsynccommerce.String("quaerat"),
                    DiscountAmount: codatsynccommerce.Float64(9853.79),
                    DiscountPercentage: codatsynccommerce.Float64(1560.98),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "eac5565d-307c-4fee-8120-6e2813fa4a41",
                        Name: codatsynccommerce.String("Dr. Herbert Legros"),
                    },
                    Quantity: 9998.54,
                    SubTotal: codatsynccommerce.Float64(1323.05),
                    TaxAmount: codatsynccommerce.Float64(802.84),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsynccommerce.Float64(1932.36),
                        ID: codatsynccommerce.String("2af03102-d514-4f4c-86f1-8bf9621a6a4f"),
                        Name: codatsynccommerce.String("Tamara O'Kon"),
                    },
                    TotalAmount: codatsynccommerce.Float64(9085.39),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "3e4be752-c65b-4344-98e3-bb91c8d975e0",
                                Name: codatsynccommerce.String("Miss Dwight Goldner"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "8f84f144-f3e0-47ed-8c4a-a5f3cabd905a",
                                Name: codatsynccommerce.String("Mitchell Crona II"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "6728227b-2d30-4947-8bf7-a4fa87cf535a",
                                Name: codatsynccommerce.String("Ora Okuneva"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "4ebf60c3-21f0-423b-b5d2-367fe1a0cc8d",
                                Name: codatsynccommerce.String("Darren McKenzie V"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("nesciunt"),
                            ID: "96d90c36-4b7c-415d-bbac-e188b1c4ee2c",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "6ce611fe-eb1c-47cb-9b6e-ec74378ba253",
                            Name: codatsynccommerce.String("Heidi Koch"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("transfer"),
                            ID: codatsynccommerce.String("c915ad2c-af5d-4d67-a3dc-0f5ae2f3a6b7"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "08787561-43f5-4a6c-98b5-5554080d40bc",
                            Name: codatsynccommerce.String("Gregg Russel"),
                        },
                    },
                    UnitAmount: 7437.95,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("laboriosam"),
            PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingInvoicePaymentAllocation{
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(8988.26),
                        TotalAmount: codatsynccommerce.Float64(8037.92),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("909304f9-26ba-4d25-9381-9b474b0ed20e"),
                            Name: codatsynccommerce.String("Terri Davis"),
                        },
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(9806.49),
                        ID: codatsynccommerce.String("f639a910-abdc-4ab6-a676-696e1ec00221"),
                        Note: codatsynccommerce.String("quidem"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("amet"),
                        TotalAmount: codatsynccommerce.Float64(3466.08),
                    },
                },
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(6016.26),
                        TotalAmount: codatsynccommerce.Float64(6294.61),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("cb3ecfda-8d0c-4549-af03-004978a61fa1"),
                            Name: codatsynccommerce.String("Ms. Darnell Denesik"),
                        },
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(9456.37),
                        ID: codatsynccommerce.String("77c1ffc7-1dca-4163-b2a3-c80a97ff334c"),
                        Note: codatsynccommerce.String("illum"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("tenetur"),
                        TotalAmount: codatsynccommerce.Float64(5620.66),
                    },
                },
            },
            SalesOrderRefs: []shared.AccountingInvoiceSalesOrderReference{
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: codatsynccommerce.String("esse"),
                    ID: codatsynccommerce.String("a9e61876-c6ab-421d-a9df-c94d6fecd799"),
                },
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: codatsynccommerce.String("dolor"),
                    ID: codatsynccommerce.String("90066a6d-2d00-4035-9338-cec086fa21e9"),
                },
            },
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusSubmitted,
            SubTotal: codatsynccommerce.Float64(1631.81),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quidem": map[string]interface{}{
                        "beatae": "sunt",
                    },
                    "molestias": map[string]interface{}{
                        "autem": "ducimus",
                    },
                    "libero": map[string]interface{}{
                        "necessitatibus": "ipsum",
                        "impedit": "quos",
                        "illum": "distinctio",
                    },
                    "voluptatem": map[string]interface{}{
                        "quaerat": "consequatur",
                    },
                },
            },
            TotalAmount: 5154.33,
            TotalDiscount: codatsynccommerce.Float64(8310.67),
            TotalTaxAmount: 4159.53,
            WithholdingTax: []shared.AccountingInvoiceWithholdingTax{
                shared.AccountingInvoiceWithholdingTax{
                    Amount: 2312.55,
                    Name: "Michele Wehner",
                },
                shared.AccountingInvoiceWithholdingTax{
                    Amount: 2941.81,
                    Name: "Ms. Samantha Metz",
                },
                shared.AccountingInvoiceWithholdingTax{
                    Amount: 1168.67,
                    Name: "Gertrude Doyle",
                },
                shared.AccountingInvoiceWithholdingTax{
                    Amount: 5231.09,
                    Name: "Alejandro DuBuque",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(175803),
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

