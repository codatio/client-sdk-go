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
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
            AdditionalTaxAmount: codataccounting.Float64(7162.96),
            AdditionalTaxPercentage: codataccounting.Float64(938.38),
            AmountDue: 1709.49,
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(3147.59),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("nam"),
                ID: "a825fe22-cd5c-4ba6-bbfe-c932af6813d6",
            },
            DiscountPercentage: codataccounting.Float64(3347.36),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("fecec2dd-6916-4f7f-87dd-a70ec60e6075"),
            InvoiceNumber: codataccounting.String("praesentium"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d61c14cd-9022-47e3-bc0d-977f1a5491ab"),
                        Name: codataccounting.String("Eduardo Kulas DDS"),
                    },
                    Description: codataccounting.String("vitae"),
                    DiscountAmount: codataccounting.Float64(387.35),
                    DiscountPercentage: codataccounting.Float64(4168.03),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "d23e03e6-9815-4aae-99fc-de9e729c9d4f",
                        Name: codataccounting.String("Freda Leannon"),
                    },
                    Quantity: 2746.11,
                    SubTotal: codataccounting.Float64(3852.3),
                    TaxAmount: codataccounting.Float64(3114.5),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(16.63),
                        ID: codataccounting.String("ca60db73-a2f9-43f4-a7dc-0d8da5612202"),
                        Name: codataccounting.String("Janie Roberts"),
                    },
                    TotalAmount: codataccounting.Float64(1810.16),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7485c197-6af9-480d-a7a0-89fc44db2745",
                                Name: codataccounting.String("Brenda Torp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c7c6d0cb-cfdc-4d33-8b6f-623bcecab50a",
                                Name: codataccounting.String("Clay Hauck PhD"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("officia"),
                            ID: "8b9af6ad-0548-46e7-b413-cbe2d176dc1c",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d40f61d1-7115-47cb-a5ee-4f7211840772",
                            Name: codataccounting.String("Chad Connelly"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "49dbe0f2-3b7b-46d9-948d-6eded477680f",
                            Name: codataccounting.String("Ms. Brad Pacocha"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "82e5e82f-d28d-4104-8a7e-91392ab44cb1",
                            Name: codataccounting.String("Mr. Jeff Hilll"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f461ce53-e914-4498-a9ba-460addfde410",
                            Name: codataccounting.String("Jeff Kuhic"),
                        },
                    },
                    UnitAmount: 6589.17,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9182a49d-9625-4d3c-affc-198eea445279"),
                        Name: codataccounting.String("Bridget Schneider"),
                    },
                    Description: codataccounting.String("ut"),
                    DiscountAmount: codataccounting.Float64(588.2),
                    DiscountPercentage: codataccounting.Float64(9020.01),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a98becce-0486-4de0-956d-73b005503e8d",
                        Name: codataccounting.String("Raul Christiansen"),
                    },
                    Quantity: 9794.26,
                    SubTotal: codataccounting.Float64(4671.12),
                    TaxAmount: codataccounting.Float64(4892.61),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7730.17),
                        ID: codataccounting.String("65675f5b-70e3-4e4c-bcc6-a91ec52624d0"),
                        Name: codataccounting.String("Margaret Berge"),
                    },
                    TotalAmount: codataccounting.Float64(9411.74),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5cea11ac-53eb-4b65-87f3-40414c5b9ace",
                                Name: codataccounting.String("Miss Francis Beier"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9f92caf1-b025-4f1d-9471-8c6fa2fad0c0",
                                Name: codataccounting.String("Francis Harris"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("exercitationem"),
                            ID: "472cdd14-fc43-4b70-bca8-8fa70c43351c",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d1eb8f7f-75f4-4f23-b1c0-a586c3ae7d7b",
                            Name: codataccounting.String("Glenda Wisozk"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5e142d95-b1db-4ece-bf7c-4b156e927827",
                            Name: codataccounting.String("Raquel VonRueden"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "68174680-63f7-499b-b956-c0b0fa0bb20a",
                            Name: codataccounting.String("Jennifer Von"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4ae64064-2726-457b-81a0-7c08fd3921c2",
                            Name: codataccounting.String("Delores Moore PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6f093a3e-fa46-4d36-adfa-1011a091b3ec",
                            Name: codataccounting.String("Garry Heller"),
                        },
                    },
                    UnitAmount: 3904.23,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("nulla"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8706.71),
                        TotalAmount: codataccounting.Float64(1150.28),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4fe72e52-1f90-4303-9fc3-38397fffa6d1"),
                            Name: codataccounting.String("Ms. Earl Collier DVM"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(956.45),
                        ID: codataccounting.String("57ac9fe1-961c-4e9b-a41c-869dd7d9719d"),
                        Note: codataccounting.String("consequatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("harum"),
                        TotalAmount: codataccounting.Float64(1382.61),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("eaque"),
                    ID: codataccounting.String("a58ffd29-67df-48fd-882a-8e60be620cd9"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusDraft,
            SubTotal: codataccounting.Float64(6380.64),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "possimus": map[string]interface{}{
                        "doloremque": "ut",
                        "eligendi": "nesciunt",
                        "voluptate": "corporis",
                        "aspernatur": "veniam",
                    },
                    "quasi": map[string]interface{}{
                        "harum": "earum",
                    },
                    "mollitia": map[string]interface{}{
                        "quasi": "vero",
                        "atque": "voluptate",
                        "itaque": "quisquam",
                        "minus": "corporis",
                    },
                    "delectus": map[string]interface{}{
                        "quod": "saepe",
                        "animi": "deleniti",
                        "eveniet": "molestiae",
                        "laborum": "voluptatum",
                    },
                },
            },
            TotalAmount: 5002.48,
            TotalDiscount: codataccounting.Float64(2381.22),
            TotalTaxAmount: 921.66,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4150.3,
                    Name: "Norma Sawayn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(428284),
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
        InvoiceID: "quibusdam",
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
        InvoiceID: "iusto",
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
        InvoiceID: "voluptate",
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
        InvoiceID: "cumque",
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
        InvoiceID: "sunt",
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
        Query: codataccounting.String("fugiat"),
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
        InvoiceID: "rem",
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
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
            AdditionalTaxAmount: codataccounting.Float64(3878.48),
            AdditionalTaxPercentage: codataccounting.Float64(394.9),
            AmountDue: 3920.09,
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(1404.57),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("nesciunt"),
                ID: "7d422786-6db8-4a74-9e39-84511cc75e4f",
            },
            DiscountPercentage: codataccounting.Float64(5.69),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("004b5bb7-58cc-4945-a2f0-069685fcd1a1"),
            InvoiceNumber: codataccounting.String("nihil"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("84bbe24f-2983-44af-b073-5cb6285d4a29"),
                        Name: codataccounting.String("Dr. Wilbur Orn III"),
                    },
                    Description: codataccounting.String("omnis"),
                    DiscountAmount: codataccounting.Float64(866.05),
                    DiscountPercentage: codataccounting.Float64(3678.44),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "6f7d2ee2-0950-45bf-83a9-3e94480ca37f",
                        Name: codataccounting.String("Andrew Bednar"),
                    },
                    Quantity: 6183.21,
                    SubTotal: codataccounting.Float64(542.66),
                    TaxAmount: codataccounting.Float64(1965.04),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1438.47),
                        ID: codataccounting.String("ac333172-e2dd-479e-874b-a7e88ddb36fd"),
                        Name: codataccounting.String("Vicky Ruecker"),
                    },
                    TotalAmount: codataccounting.Float64(2904.79),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c8657347-4f0a-4740-bb4a-b441c3a09e76",
                                Name: codataccounting.String("Katrina Monahan"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("rem"),
                            ID: "08bbe794-455e-4bc5-90a1-c426b59c8366",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "cc135582-c1b8-455e-889d-9ef932e9000a",
                            Name: codataccounting.String("Peggy Nolan"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "24208efd-2341-4189-8e73-879efbe8baeb",
                            Name: codataccounting.String("Pete Rice"),
                        },
                    },
                    UnitAmount: 2814.76,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("536e9035-1bb9-4763-9720-b77a5a5365a7"),
                        Name: codataccounting.String("Timmy Bernier"),
                    },
                    Description: codataccounting.String("voluptate"),
                    DiscountAmount: codataccounting.Float64(705.55),
                    DiscountPercentage: codataccounting.Float64(9961.97),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "01c0d361-fed8-4dc5-affb-453e9089e871",
                        Name: codataccounting.String("Jody Rogahn"),
                    },
                    Quantity: 3763.5,
                    SubTotal: codataccounting.Float64(5968.08),
                    TaxAmount: codataccounting.Float64(4751.53),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7361.58),
                        ID: codataccounting.String("dd9c985e-4373-44a5-972d-9edd785be5e7"),
                        Name: codataccounting.String("Terrell Tillman"),
                    },
                    TotalAmount: codataccounting.Float64(1729.85),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7ba6281f-44e3-4a23-b94a-68cc80d30ff7",
                                Name: codataccounting.String("Debra Jacobi"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0a91fe9d-9655-43b8-9e00-09c6692de7b3",
                                Name: codataccounting.String("Loretta Collier Jr."),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a6aab4ae-7b1a-45b9-88d4-e30491aa35d4",
                                Name: codataccounting.String("Alfredo Fahey"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consequatur"),
                            ID: "3bab77b9-18f0-4313-9845-07e0e39c7e23",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b0604652-e23a-43d6-8657-e9de8f7f002d",
                            Name: codataccounting.String("Toni Legros"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "99d3a1d3-2329-4e45-837e-8f2ad6bb10e2",
                            Name: codataccounting.String("Roberta Wisozk"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "480d6e33-0867-45cb-b186-856a7e82cdf9",
                            Name: codataccounting.String("Donald Wiegand"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "82c666af-3c3f-4558-9bea-5d264e41e2ca",
                            Name: codataccounting.String("Francis Lockman"),
                        },
                    },
                    UnitAmount: 8864.94,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("513f6d9d-2ad3-47c3-8990-77c10b687921"),
                        Name: codataccounting.String("Cindy Vandervort"),
                    },
                    Description: codataccounting.String("repellendus"),
                    DiscountAmount: codataccounting.Float64(2517.76),
                    DiscountPercentage: codataccounting.Float64(5242.33),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "860543c0-a304-49c3-8f6c-0276e7b21bad",
                        Name: codataccounting.String("Robert Stanton"),
                    },
                    Quantity: 2834.63,
                    SubTotal: codataccounting.Float64(2437.27),
                    TaxAmount: codataccounting.Float64(9841.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8573.88),
                        ID: codataccounting.String("6c2a10e6-c297-48ec-a56a-5b09227fcc47"),
                        Name: codataccounting.String("Wendell Huels"),
                    },
                    TotalAmount: codataccounting.Float64(4585.85),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "bbc57f38-928a-4860-8c58-d67d63e4aa56",
                                Name: codataccounting.String("Jay Kassulke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "79cfc6c0-e503-4f56-831f-1d8ed87b28e8",
                                Name: codataccounting.String("Darnell Nader"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("excepturi"),
                            ID: "86e241e4-3b23-4424-97d1-3e3f62aa9ae4",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8ab4a9c4-92c5-4e8b-a5d4-aa4a508bd380",
                            Name: codataccounting.String("Aaron Marks"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "dd71bdda-a30b-47b9-9449-ae69c088d418",
                            Name: codataccounting.String("Ms. Wm Kohler II"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f423d543-935f-4377-ac5c-9b7e93b6a3c5",
                            Name: codataccounting.String("Mrs. Grace Botsford"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7c34cab0-ecb8-412a-a614-8944a8e90850",
                            Name: codataccounting.String("Annette Reilly"),
                        },
                    },
                    UnitAmount: 3468.95,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("38253343-fb0a-44e6-aea4-7578d171e294"),
                        Name: codataccounting.String("Christy Bernhard"),
                    },
                    Description: codataccounting.String("optio"),
                    DiscountAmount: codataccounting.Float64(3993.59),
                    DiscountPercentage: codataccounting.Float64(4504.05),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9b6b2f25-359b-4855-9015-b62c8b83a38a",
                        Name: codataccounting.String("Angelo Langosh"),
                    },
                    Quantity: 1119.71,
                    SubTotal: codataccounting.Float64(2686.2),
                    TaxAmount: codataccounting.Float64(2857.2),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1609.42),
                        ID: codataccounting.String("00c2caeb-1ae1-4ecf-8c34-946bba7a05a8"),
                        Name: codataccounting.String("Oscar O'Connell"),
                    },
                    TotalAmount: codataccounting.Float64(7724.02),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b3688cca-3632-4727-a0e9-66e97e054103",
                                Name: codataccounting.String("Michele Koelpin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8ff24911-45fa-4b9e-99a4-af336664eaa6",
                                Name: codataccounting.String("Timmy Daniel"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quasi"),
                            ID: "4e8c1b35-2acc-4eda-8c52-27814eca016b",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1ea1342d-4104-4a25-af71-de57a11d614a",
                            Name: codataccounting.String("Shannon Carter"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2ea48673-d522-4b82-8a90-30660f024c79",
                            Name: codataccounting.String("Alexander Rosenbaum"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4c2b3a32-c488-4ade-a2f6-aa558a65e208",
                            Name: codataccounting.String("Susan Boyer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a34bb87d-4f62-4127-a607-d1606294514c",
                            Name: codataccounting.String("Freda Reichel"),
                        },
                    },
                    UnitAmount: 6279.31,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("sapiente"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1578.59),
                        TotalAmount: codataccounting.Float64(7272.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e8787034-93f4-49aa-8465-a3283279b719"),
                            Name: codataccounting.String("Raymond Rosenbaum"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4429.32),
                        ID: codataccounting.String("3d86e3b3-5e49-4a31-b577-8ce54cacb0e3"),
                        Note: codataccounting.String("vero"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("sint"),
                        TotalAmount: codataccounting.Float64(4960.42),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2892.81),
                        TotalAmount: codataccounting.Float64(3596.49),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("bacf63b2-1518-46ab-9e3a-022614315d15"),
                            Name: codataccounting.String("Lena Cummings"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(3778.77),
                        ID: codataccounting.String("1afc7186-ff20-4b7a-b3df-40ca0d7657c1"),
                        Note: codataccounting.String("ex"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("ab"),
                        TotalAmount: codataccounting.Float64(7166.01),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(132.23),
                        TotalAmount: codataccounting.Float64(3450.21),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5271b251-1dd6-406d-91b2-8272bc9c3221"),
                            Name: codataccounting.String("Becky Kunde V"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(276.89),
                        ID: codataccounting.String("fcbb2b93-c15f-4670-bd17-84831653eeb3"),
                        Note: codataccounting.String("tempore"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("vero"),
                        TotalAmount: codataccounting.Float64(1407.83),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("dicta"),
                    ID: codataccounting.String("c3109983-663c-466d-8bb7-df6cb09c8b40"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("praesentium"),
                    ID: codataccounting.String("e0713774-de4f-4ee1-81d9-780a10c47b95"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusDraft,
            SubTotal: codataccounting.Float64(370.79),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "vel": map[string]interface{}{
                        "blanditiis": "soluta",
                        "quia": "similique",
                        "ipsam": "a",
                        "alias": "perferendis",
                    },
                    "aspernatur": map[string]interface{}{
                        "sit": "esse",
                    },
                    "accusamus": map[string]interface{}{
                        "quae": "dolore",
                        "molestias": "maiores",
                    },
                    "cupiditate": map[string]interface{}{
                        "alias": "sit",
                    },
                },
            },
            TotalAmount: 6122.66,
            TotalDiscount: codataccounting.Float64(9351.45),
            TotalTaxAmount: 8479.43,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5771.16,
                    Name: "Phyllis Koch",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "quidem",
        TimeoutInMinutes: codataccounting.Int(306065),
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
            Content: []byte("fuga"),
            RequestBody: "itaque",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "iste",
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

