# Invoices

## Overview

Invoices

### Available Operations

* [Create](#create) - Create invoice
* [Delete](#delete) - Delete invoice
* [DownloadAttachment](#downloadattachment) - Download invoice attachment
* [DownloadPdf](#downloadpdf) - Get invoice as PDF
* [Get](#get) - Get invoice
* [GetAttachment](#getattachment) - Get invoice attachment
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update invoice model
* [List](#list) - List invoices
* [ListAttachments](#listattachments) - List invoice attachments
* [Update](#update) - Update invoice
* [UploadAttachment](#uploadattachment) - Push invoice attachment

## Create

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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.Create(ctx, operations.CreateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(6499.01),
            AdditionalTaxPercentage: codataccounting.Float64(3702.19),
            AmountDue: 4227.22,
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(427.39),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("atque"),
                ID: "91500970-19a4-48f8-8ece-7bf904e01105",
            },
            DiscountPercentage: codataccounting.Float64(8370.8),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("8908162c-6beb-468a-8f65-7b7d03a1480f"),
            InvoiceNumber: codataccounting.String("blanditiis"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("30f069d8-1061-48d9-be15-2297510da803"),
                        Name: codataccounting.String("Janice Cronin"),
                    },
                    Description: codataccounting.String("nobis"),
                    DiscountAmount: codataccounting.Float64(7615.63),
                    DiscountPercentage: codataccounting.Float64(3870.67),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1c2a702b-b97e-4e10-ada2-de35f8e01bf3",
                        Name: codataccounting.String("Silvia Murazik"),
                    },
                    Quantity: 2836.19,
                    SubTotal: codataccounting.Float64(3304.68),
                    TaxAmount: codataccounting.Float64(2609.08),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(619.54),
                        ID: codataccounting.String("2ac1704b-f1cc-49fc-a1aa-e5eb5f0c492b"),
                        Name: codataccounting.String("Jackie Graham"),
                    },
                    TotalAmount: codataccounting.Float64(406.34),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "a2267aae-e79e-43c7-9ad3-1becb83d2378",
                                Name: codataccounting.String("Rogelio Doyle"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c23d9450-a986-4a49-9bac-707f06b28ecc",
                                Name: codataccounting.String("Raul Gerlach I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "86f62c96-9c4c-4c6b-b889-0a3fd3c81da1",
                                Name: codataccounting.String("Tabitha Lesch"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("non"),
                            ID: "df931da3-edb5-41fa-994a-cc9435137726",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "5321b832-a56d-4691-80ff-60eb9a6658e6",
                            Name: codataccounting.String("Grant Fritsch"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3d382dbe-c75c-468c-a065-9468ce304d88",
                            Name: codataccounting.String("Shelly Quitzon"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "214c337f-96bb-40c6-9e37-2db1344ba9f7",
                            Name: codataccounting.String("Ernesto Heaney PhD"),
                        },
                    },
                    UnitAmount: 8162.72,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7aab62e9-7261-4fb0-858d-27b51996b5b4"),
                        Name: codataccounting.String("Jon Bashirian"),
                    },
                    Description: codataccounting.String("sapiente"),
                    DiscountAmount: codataccounting.Float64(4822.43),
                    DiscountPercentage: codataccounting.Float64(988.3),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "2b7a7ab0-344b-4171-8688-deebef897f3d",
                        Name: codataccounting.String("Mark Schmidt"),
                    },
                    Quantity: 2232.35,
                    SubTotal: codataccounting.Float64(2259.45),
                    TaxAmount: codataccounting.Float64(9483.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(979.03),
                        ID: codataccounting.String("1b3e4e08-0aa1-4041-86ec-759e02f3702c"),
                        Name: codataccounting.String("Blanca Langworth"),
                    },
                    TotalAmount: codataccounting.Float64(8570.75),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0ead3104-fa44-4707-bf37-5b44282821fd",
                                Name: codataccounting.String("Aaron Weimann"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eveniet"),
                            ID: "59267c71-cc8d-43cd-8258-d0358a82c808",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "2751a204-7c04-449e-943f-9619bb7d40d5",
                            Name: codataccounting.String("Eric Bergnaum"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "36e62592-33f9-45c9-9237-397c785b5db4",
                            Name: codataccounting.String("Mr. Wesley Ankunding"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3febdf67-6b72-406d-ab75-0052a5647edc",
                            Name: codataccounting.String("Annie Morissette"),
                        },
                    },
                    UnitAmount: 5481.43,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c4320f41-240d-4448-bac6-93b94c3b9d24"),
                        Name: codataccounting.String("Guy Spinka"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(6383.23),
                    DiscountPercentage: codataccounting.Float64(6864.21),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "42fc4056-69f6-49a0-86d2-1249450819d7",
                        Name: codataccounting.String("Miss Nathan Ritchie"),
                    },
                    Quantity: 745.3,
                    SubTotal: codataccounting.Float64(5385.16),
                    TaxAmount: codataccounting.Float64(2986.57),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2793.53),
                        ID: codataccounting.String("060e0031-0d02-43dc-901f-5afd2a6c4484"),
                        Name: codataccounting.String("Genevieve Waters"),
                    },
                    TotalAmount: codataccounting.Float64(5056.63),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "253c8962-f489-46bf-91e4-652d3c343d61",
                                Name: codataccounting.String("Stella Littel"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "49124772-5e62-4190-9e91-044a5de59ac7",
                                Name: codataccounting.String("Jessica Hyatt"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0cf1cf59-3260-4525-9e66-bb426897d99a",
                                Name: codataccounting.String("Christie Frami"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("suscipit"),
                            ID: "70e93ee6-cf59-4f35-8aae-acae323a31bf",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "a1cc9771-6c80-42cc-9e0c-7d9d323f1aa6",
                            Name: codataccounting.String("Cecelia Stoltenberg"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1c856bcb-a51e-4f24-94a4-7facf116cdd5",
                            Name: codataccounting.String("Jamie Funk"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "562873c7-dd9e-4faf-83dc-623620f3138f",
                            Name: codataccounting.String("Michelle Stroman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "db022faa-565f-4b8f-a52e-bb9d38383879",
                            Name: codataccounting.String("Beverly Green"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "293dab30-e917-4f50-bda0-4c8b1bb55a29",
                            Name: codataccounting.String("Karla Armstrong"),
                        },
                    },
                    UnitAmount: 2482.32,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("bb744664-eb1d-4033-88b0-d1bb17afee74"),
                        Name: codataccounting.String("Hector Willms"),
                    },
                    Description: codataccounting.String("occaecati"),
                    DiscountAmount: codataccounting.Float64(2897.27),
                    DiscountPercentage: codataccounting.Float64(3697.22),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "7c7edaf3-9d16-4fbf-b6fd-162b303e3023",
                        Name: codataccounting.String("Andy Erdman"),
                    },
                    Quantity: 2724.93,
                    SubTotal: codataccounting.Float64(1973.88),
                    TaxAmount: codataccounting.Float64(893.2),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3986.87),
                        ID: codataccounting.String("cf55b431-3553-4ccf-9c20-4c4adcc9904c"),
                        Name: codataccounting.String("Debra Medhurst"),
                    },
                    TotalAmount: codataccounting.Float64(5405.93),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "48cefa78-f1e2-4d3b-901e-0952bbb4cbb1",
                                Name: codataccounting.String("Mrs. Courtney Kuhic"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95a4169c-1387-4271-a18e-a9e45118c2cc",
                                Name: codataccounting.String("Dora White"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("aliquid"),
                            ID: "0b1a78ed-29a9-4d4e-aa85-658c2d4f4c88",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "4f278fd9-667e-446c-91d2-ffaa58dcef23",
                            Name: codataccounting.String("Alexandra Morissette"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9bdf2190-abd9-4bbc-8272-5ec2659ce028",
                            Name: codataccounting.String("Miss Priscilla Gerhold"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9ef68e45-c8ad-4dfa-8754-500430c6632b",
                            Name: codataccounting.String("Dr. Carmen McKenzie"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f01c3e91-e8f7-4bc6-9d46-0a77eceb26d1",
                            Name: codataccounting.String("Lorene Bosco"),
                        },
                    },
                    UnitAmount: 1818.77,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("neque"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(461.37),
                        TotalAmount: codataccounting.Float64(9654.91),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0f873f9d-5c25-4fd3-a0b4-a4a4253c3025"),
                            Name: codataccounting.String("Julie Bergstrom"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(8074.3),
                        ID: codataccounting.String("7e7dc548-be09-4e41-a7a2-15ca12a4ba9d"),
                        Note: codataccounting.String("ipsam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("omnis"),
                        TotalAmount: codataccounting.Float64(5229.85),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5952.02),
                        TotalAmount: codataccounting.Float64(1465.4),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("cfd0c77c-53e7-4e7d-8ee6-e8b90bac384e"),
                            Name: codataccounting.String("Cindy Marquardt"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1935.64),
                        ID: codataccounting.String("fec31c50-824d-4189-a36a-6b2d27eb707a"),
                        Note: codataccounting.String("est"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("aut"),
                        TotalAmount: codataccounting.Float64(8090.72),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8853.36),
                        TotalAmount: codataccounting.Float64(2924.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6e6177db-9db3-4b70-bfbb-6970ee770e36"),
                            Name: codataccounting.String("Violet Kshlerin"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8024.72),
                        ID: codataccounting.String("206e61b0-d308-4714-820a-3d98637ca85c"),
                        Note: codataccounting.String("adipisci"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("repudiandae"),
                        TotalAmount: codataccounting.Float64(4193.51),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4780.21),
                        TotalAmount: codataccounting.Float64(2820.79),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("dbaf94a7-c98f-413a-b28d-b2cf2bf4f3de"),
                            Name: codataccounting.String("Vincent Hansen"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8883.28),
                        ID: codataccounting.String("14b21cd9-8196-4d55-af69-a1c4b79ae336"),
                        Note: codataccounting.String("praesentium"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("eligendi"),
                        TotalAmount: codataccounting.Float64(1805.44),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("eligendi"),
                    ID: codataccounting.String("39a7c0e1-7cb1-42c5-ba82-5fe22cd5cba6"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusPaid,
            SubTotal: codataccounting.Float64(9961.28),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quisquam": map[string]interface{}{
                        "amet": "consequuntur",
                        "fuga": "a",
                        "aliquid": "voluptatum",
                    },
                    "sunt": map[string]interface{}{
                        "illum": "ea",
                    },
                    "veniam": map[string]interface{}{
                        "delectus": "earum",
                        "placeat": "saepe",
                        "quod": "odit",
                    },
                    "assumenda": map[string]interface{}{
                        "ea": "provident",
                        "inventore": "ea",
                        "repellat": "quam",
                        "delectus": "minus",
                    },
                },
            },
            TotalAmount: 4677.01,
            TotalDiscount: codataccounting.Float64(8447.75),
            TotalTaxAmount: 8284.89,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4773.62,
                    Name: "Kate Schowalter PhD",
                },
                shared.WithholdingTaxitems{
                    Amount: 3877.68,
                    Name: "Melinda Heaney",
                },
                shared.WithholdingTaxitems{
                    Amount: 2559.53,
                    Name: "Milton Bogisich I",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(787873),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateInvoiceRequest](../../models/operations/createinvoicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.CreateInvoiceResponse](../../models/operations/createinvoiceresponse.md), error**


## Delete

﻿The *Delete invoice* endpoint allows you to delete a specified invoice from an accounting platform.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

### Process
1. Pass the `{invoiceId}` to the *Delete invoice* endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete operation by checking the status of push operation either via
    1. [Push operation webhook](https://docs.codat.io/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
    2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation).

   A `Success` status indicates that the invoice object was deleted from the accounting platform.
3. (Optional) Check that the invoice was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting an invoice from an accounting platform might cause related objects to be modified. For example, if you delete a paid invoice from QuickBooks Online, the invoice is deleted but the payment against that invoice is not. The payment is converted to a payment on account.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Deleted | 
|-------------|--------------|
| QuickBooks Online | Yes    |     

> **Supported Integrations**
> 
> This functionality is currently only supported for our QuickBooks Online integration. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.
> We're increasing support for object deletion across various accounting platforms and data types. You can check our [Accounting API Public Product Roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) for the latest status.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.Delete(ctx, operations.DeleteInvoiceRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "facere",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteInvoiceRequest](../../models/operations/deleteinvoicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.DeleteInvoiceResponse](../../models/operations/deleteinvoiceresponse.md), error**


## DownloadAttachment

The *Download invoice attachment* endpoint downloads a specific attachment for a given `invoiceId` and `attachmentId`.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support downloading an invoice attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.DownloadAttachment(ctx, operations.DownloadInvoiceAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DownloadInvoiceAttachmentRequest](../../models/operations/downloadinvoiceattachmentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.DownloadInvoiceAttachmentResponse](../../models/operations/downloadinvoiceattachmentresponse.md), error**


## DownloadPdf

﻿Download invoice as a pdf.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.DownloadPdf(ctx, operations.DownloadInvoicePdfRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        InvoiceID: "aut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DownloadInvoicePdfRequest](../../models/operations/downloadinvoicepdfrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.DownloadInvoicePdfResponse](../../models/operations/downloadinvoicepdfresponse.md), error**


## Get

The *Get invoice* endpoint returns a single invoice for a given invoiceId.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support getting a specific invoice.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.Get(ctx, operations.GetInvoiceRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        InvoiceID: "aspernatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetInvoiceRequest](../../models/operations/getinvoicerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.GetInvoiceResponse](../../models/operations/getinvoiceresponse.md), error**


## GetAttachment

The *Get invoice attachment* endpoint returns a specific attachment for a given `invoiceId` and `attachmentId`.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support getting an invoice attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.GetAttachment(ctx, operations.GetInvoiceAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "odit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetInvoiceAttachmentRequest](../../models/operations/getinvoiceattachmentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.GetInvoiceAttachmentResponse](../../models/operations/getinvoiceattachmentresponse.md), error**


## GetCreateUpdateModel

The *Get create/update invoice model* endpoint returns the expected data for the request payload when creating and updating an [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) for a given company and integration.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support creating and updating an invoice.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.GetCreateUpdateModel(ctx, operations.GetCreateUpdateInvoicesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCreateUpdateInvoicesModelRequest](../../models/operations/getcreateupdateinvoicesmodelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetCreateUpdateInvoicesModelResponse](../../models/operations/getcreateupdateinvoicesmodelresponse.md), error**


## List

The *List invoices* endpoint returns a list of [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) for a given company's connection.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.List(ctx, operations.ListInvoicesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("molestiae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListInvoicesRequest](../../models/operations/listinvoicesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.ListInvoicesResponse](../../models/operations/listinvoicesresponse.md), error**


## ListAttachments

The *List invoice attachments* endpoint returns a list of attachments available to download for given `invoiceId`.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support listing invoice attachments.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.ListAttachments(ctx, operations.ListInvoiceAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "recusandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListInvoiceAttachmentsRequest](../../models/operations/listinvoiceattachmentsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.ListInvoiceAttachmentsResponse](../../models/operations/listinvoiceattachmentsresponse.md), error**


## Update

The *Update invoice* endpoint updates an existing [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) for a given company's connection.

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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.Update(ctx, operations.UpdateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(2046.33),
            AdditionalTaxPercentage: codataccounting.Float64(4704),
            AmountDue: 7513.92,
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(8156.07),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("iste"),
                ID: "77f1a549-1abe-4975-9b10-6d23e03e6981",
            },
            DiscountPercentage: codataccounting.Float64(3304.02),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("ae99fcde-9e72-49c9-94f2-d8a44640ca60"),
            InvoiceNumber: codataccounting.String("illum"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3a2f93f4-67dc-40d8-9a56-122026ab8f27"),
                        Name: codataccounting.String("Eleanor Lang"),
                    },
                    Description: codataccounting.String("et"),
                    DiscountAmount: codataccounting.Float64(5825.69),
                    DiscountPercentage: codataccounting.Float64(4501.13),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "6af980da-7a08-49fc-84db-274530e5cc7c",
                        Name: codataccounting.String("Doreen Ankunding"),
                    },
                    Quantity: 7667.05,
                    SubTotal: codataccounting.Float64(9649.39),
                    TaxAmount: codataccounting.Float64(8335.9),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7785.2),
                        ID: codataccounting.String("d334b6f6-23bc-4eca-b50a-ee5e0da8b9af"),
                        Name: codataccounting.String("Mrs. Alberta Stoltenberg"),
                    },
                    TotalAmount: codataccounting.Float64(5382.58),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e7b413cb-e2d1-476d-81c4-3d40f61d1711",
                                Name: codataccounting.String("Joy Runolfsdottir"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5ee4f721-1840-4772-b32e-3b49dbe0f23b",
                                Name: codataccounting.String("Lola Homenick"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("cupiditate"),
                            ID: "48d6eded-4776-480f-87a1-7a82e5e82fd2",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1040a7e9-1392-4ab4-8cb1-835008f461ce",
                            Name: codataccounting.String("Robin Treutel I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "98a9ba46-0add-4fde-810c-37daa9182a49",
                            Name: codataccounting.String("Sergio Jacobi"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d3caffc1-98ee-4a44-9279-2bcd440ea98b",
                            Name: codataccounting.String("Guadalupe Sawayn II"),
                        },
                    },
                    UnitAmount: 5297.99,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6de0d56d-73b0-4055-83e8-dc626ff77c65"),
                        Name: codataccounting.String("Courtney Hamill"),
                    },
                    Description: codataccounting.String("harum"),
                    DiscountAmount: codataccounting.Float64(4443.84),
                    DiscountPercentage: codataccounting.Float64(256.53),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "e3e4cfcc-6a91-4ec5-a624-d00014ef45ce",
                        Name: codataccounting.String("Samuel Bruen"),
                    },
                    Quantity: 3285.77,
                    SubTotal: codataccounting.Float64(2073.91),
                    TaxAmount: codataccounting.Float64(8924.85),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6954.27),
                        ID: codataccounting.String("b6587f34-0414-4c5b-9ace-e400ae9f92ca"),
                        Name: codataccounting.String("Mr. Joe Rogahn"),
                    },
                    TotalAmount: codataccounting.Float64(9546.52),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d14718c6-fa2f-4ad0-806c-5d95472cdd14",
                                Name: codataccounting.String("Sherman Gerlach"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("in"),
                            ID: "0bca88fa-70c4-4335-9c3d-d1eb8f7f75f4",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3f1c0a58-6c3a-4e7d-bb67-feef5e142d95",
                            Name: codataccounting.String("Ralph Stracke"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "eff7c4b1-56e9-4278-a75e-ea7681746806",
                            Name: codataccounting.String("Mandy Kutch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b7956c0b-0fa0-4bb2-8a40-e7c4ae640642",
                            Name: codataccounting.String("Janice Jones"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b01a07c0-8fd3-4921-8257-930d6f093a3e",
                            Name: codataccounting.String("Lorenzo Gutmann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "366dfa10-11a0-491b-bec8-b53862de1a9d",
                            Name: codataccounting.String("Alicia Ziemann"),
                        },
                    },
                    UnitAmount: 1519.16,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("ipsam"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(593.03),
                        TotalAmount: codataccounting.Float64(2423.78),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("03dfc338-397f-4ffa-ad1d-32090fc157ac"),
                            Name: codataccounting.String("Ms. Aubrey Thiel"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(7529.61),
                        ID: codataccounting.String("e9be41c8-69dd-47d9-b19d-07b200a58ffd"),
                        Note: codataccounting.String("consequuntur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("suscipit"),
                        TotalAmount: codataccounting.Float64(4825.35),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("delectus"),
                    ID: codataccounting.String("8fd882a8-e60b-4e62-8cd9-c5afdd04c375"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("aspernatur"),
                    ID: codataccounting.String("512beae1-d87e-4cc5-bdce-a8e7a8831166"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("ratione"),
                    ID: codataccounting.String("cda6d77c-1d86-4066-a37d-4227866db8a7"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("labore"),
                    ID: codataccounting.String("9e398451-1cc7-45e4-b0c0-04b5bb758cc9"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusSubmitted,
            SubTotal: codataccounting.Float64(3765.76),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "delectus": map[string]interface{}{
                        "consequatur": "suscipit",
                    },
                },
            },
            TotalAmount: 5771.02,
            TotalDiscount: codataccounting.Float64(4067.2),
            TotalTaxAmount: 5272.66,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9529.5,
                    Name: "Owen Boyer IV",
                },
                shared.WithholdingTaxitems{
                    Amount: 2334.63,
                    Name: "Dwight Fritsch",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "saepe",
        TimeoutInMinutes: codataccounting.Int(139567),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateInvoiceRequest](../../models/operations/updateinvoicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.UpdateInvoiceResponse](../../models/operations/updateinvoiceresponse.md), error**


## UploadAttachment

The *Upload invoice attachment* endpoint uploads an attachment and assigns it against a specific `invoiceId`.

[Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) are itemized records of goods sold or services provided to a customer.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/accounting-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support uploading an invoice attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.UploadAttachment(ctx, operations.UploadInvoiceAttachmentRequest{
        RequestBody: &operations.UploadInvoiceAttachmentRequestBody{
            Content: []byte("modi"),
            RequestBody: "tenetur",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "explicabo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UploadInvoiceAttachmentRequest](../../models/operations/uploadinvoiceattachmentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.UploadInvoiceAttachmentResponse](../../models/operations/uploadinvoiceattachmentresponse.md), error**

