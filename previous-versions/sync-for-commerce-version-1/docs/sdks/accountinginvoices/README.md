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
            AdditionalTaxAmount: types.MustNewDecimalFromString("2282.63"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("1059.06"),
            AmountDue: types.MustNewDecimalFromString("4895.09"),
            Currency: syncforcommerceversion1.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("8915.23"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: syncforcommerceversion1.String("consectetur"),
                ID: "5b60eb1e-a426-4555-ba3c-28744ed53b88",
            },
            DiscountPercentage: types.MustNewDecimalFromString("9425.84"),
            DueDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            ID: syncforcommerceversion1.String("a8d8f5c0-b2f2-4fb7-b194-a276b26916fe"),
            InvoiceNumber: syncforcommerceversion1.String("illo"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("08f4294e-3698-4f44-bf60-3e8b445e80ca"),
                        Name: syncforcommerceversion1.String("Lorraine Walsh"),
                    },
                    Description: syncforcommerceversion1.String("magni"),
                    DiscountAmount: types.MustNewDecimalFromString("486.9"),
                    DiscountPercentage: types.MustNewDecimalFromString("9014.83"),
                    IsDirectIncome: syncforcommerceversion1.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "457e1858-b6a8-49fb-a3a5-aa8e4824d0ab",
                        Name: syncforcommerceversion1.String("Barbara Koelpin IV"),
                    },
                    Quantity: types.MustNewDecimalFromString("5580.65"),
                    SubTotal: types.MustNewDecimalFromString("9221.12"),
                    TaxAmount: types.MustNewDecimalFromString("3611.51"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("894.94"),
                        ID: syncforcommerceversion1.String("862065e9-04f3-4b11-94b8-abf603a79f9d"),
                        Name: syncforcommerceversion1.String("Noah Armstrong"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("4406.66"),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "da8a50ce-187f-486b-8173-d689eee9526f",
                                Name: syncforcommerceversion1.String("Wilfred Mueller"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: syncforcommerceversion1.String("repudiandae"),
                            ID: "881ead4f-0e10-4125-a3f9-4e29e973e922",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "7a15be3e-0608-407e-ab6e-3ab8845f0597",
                            Name: syncforcommerceversion1.String("Shane Abshire"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: syncforcommerceversion1.String("journalEntry"),
                            ID: syncforcommerceversion1.String("a54a31e9-4764-4a3e-865e-7956f9251a5a"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "9da660ff-57bf-4aad-8f9e-fc1b4512c103",
                            Name: syncforcommerceversion1.String("Agnes Gibson"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7730.84"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Note: syncforcommerceversion1.String("sapiente"),
            PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.AccountingInvoicePaymentAllocation{
                shared.AccountingInvoicePaymentAllocation{
                    Allocation: shared.AccountingInvoicePaymentAllocationAllocation{
                        AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforcommerceversion1.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("1070.04"),
                        TotalAmount: types.MustNewDecimalFromString("5834.04"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforcommerceversion1.String("9ebfd0e9-fe6c-4632-8a3a-ed0117996312"),
                            Name: syncforcommerceversion1.String("Mrs. Orville Treutel"),
                        },
                        Currency: syncforcommerceversion1.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("1158.34"),
                        ID: syncforcommerceversion1.String("778ff61d-0174-4763-a0a1-5db6a660659a"),
                        Note: syncforcommerceversion1.String("ab"),
                        PaidOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforcommerceversion1.String("possimus"),
                        TotalAmount: types.MustNewDecimalFromString("9139.92"),
                    },
                },
            },
            SalesOrderRefs: []shared.AccountingInvoiceSalesOrderReference{
                shared.AccountingInvoiceSalesOrderReference{
                    DataType: syncforcommerceversion1.String("mollitia"),
                    ID: syncforcommerceversion1.String("ab5851d6-c645-4b08-b618-91baa0fe1ade"),
                },
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusUnknown,
            SubTotal: types.MustNewDecimalFromString("5349.17"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "earum": map[string]interface{}{
                        "ex": "sapiente",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("5241.84"),
            TotalDiscount: types.MustNewDecimalFromString("7963.2"),
            TotalTaxAmount: types.MustNewDecimalFromString("3651"),
            WithholdingTax: []shared.AccountingInvoiceWithholdingTax{
                shared.AccountingInvoiceWithholdingTax{
                    Amount: types.MustNewDecimalFromString("9920.74"),
                    Name: "Marion Aufderhar",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(770675),
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

