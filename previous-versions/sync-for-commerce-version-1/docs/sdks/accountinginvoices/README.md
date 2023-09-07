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
            AdditionalTaxAmount: codatsynccommerce.Float64(7386.83),
            AdditionalTaxPercentage: codatsynccommerce.Float64(2326.27),
            AmountDue: 4490.83,
            Currency: codatsynccommerce.String("USD"),
            CurrencyRate: codatsynccommerce.Float64(9372.85),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("facere"),
                ID: "4f6fbee4-1f33-4317-be35-b60eb1ea4265",
            },
            DiscountPercentage: codatsynccommerce.Float64(3742.96),
            DueDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            ID: codatsynccommerce.String("ba3c2874-4ed5-43b8-8f3a-8d8f5c0b2f2f"),
            InvoiceNumber: codatsynccommerce.String("facilis"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("b194a276-b269-416f-a1f0-8f4294e3698f"),
                        Name: codatsynccommerce.String("Rhonda Klocko"),
                    },
                    Description: codatsynccommerce.String("sit"),
                    DiscountAmount: codatsynccommerce.Float64(2484.13),
                    DiscountPercentage: codatsynccommerce.Float64(8880.44),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "8b445e80-ca55-4efd-a0e4-57e1858b6a89",
                        Name: codatsynccommerce.String("Rudolph Trantow"),
                    },
                    Quantity: 3416.98,
                    SubTotal: codatsynccommerce.Float64(6390.28),
                    TaxAmount: codatsynccommerce.Float64(6762.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsynccommerce.Float64(5483.61),
                        ID: codatsynccommerce.String("e4824d0a-b407-4508-8e51-862065e904f3"),
                        Name: codatsynccommerce.String("Gerald Bradtke"),
                    },
                    TotalAmount: codatsynccommerce.Float64(6952.7),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "8abf603a-79f9-4dfe-8ab7-da8a50ce187f",
                                Name: codatsynccommerce.String("Sam Powlowski IV"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("amet"),
                            ID: "d689eee9-526f-48d9-86e8-81ead4f0e101",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "63f94e29-e973-4e92-aa57-a15be3e06080",
                            Name: codatsynccommerce.String("Tricia Denesik"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("transfer"),
                            ID: codatsynccommerce.String("3ab8845f-0597-4a60-bf2a-54a31e94764a"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "3e865e79-56f9-4251-a5a9-da660ff57bfa",
                            Name: codatsynccommerce.String("Irving Gleichner"),
                        },
                    },
                    UnitAmount: 8897.94,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("cumque"),
            PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingInvoicePaymentAllocation{
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(3354.98),
                        TotalAmount: codatsynccommerce.Float64(820.57),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("2c103264-8dc2-4f61-9199-ebfd0e9fe6c6"),
                            Name: codatsynccommerce.String("Denise Runolfsdottir"),
                        },
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(8987.6),
                        ID: codatsynccommerce.String("d0117996-312f-4de0-8771-778ff61d0174"),
                        Note: codatsynccommerce.String("esse"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("consectetur"),
                        TotalAmount: codatsynccommerce.Float64(3998.12),
                    },
                },
            },
            SalesOrderRefs: []shared.AccountingInvoiceSalesOrderReference{
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: codatsynccommerce.String("ipsa"),
                    ID: codatsynccommerce.String("a15db6a6-6065-49a1-adea-ab5851d6c645"),
                },
            },
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusUnknown,
            SubTotal: codatsynccommerce.Float64(5615.77),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "cum": map[string]interface{}{
                        "aliquid": "beatae",
                    },
                },
            },
            TotalAmount: 5308.6,
            TotalDiscount: codatsynccommerce.Float64(6063.08),
            TotalTaxAmount: 852.33,
            WithholdingTax: []shared.AccountingInvoiceWithholdingTax{
                shared.AccountingInvoiceWithholdingTax{
                    Amount: 7032.18,
                    Name: "Trevor Bartell",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(103298),
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

