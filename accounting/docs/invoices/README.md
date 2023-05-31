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

Posts a new invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) to see which integrations support this endpoint.


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
            AdditionalTaxAmount: codataccounting.Float64(7733.49),
            AdditionalTaxPercentage: codataccounting.Float64(6927.98),
            AmountDue: 289.46,
            Currency: codataccounting.String("voluptas"),
            CurrencyRate: codataccounting.Float64(447.24),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("numquam"),
                ID: "652e23a3-d6c6-457e-9de8-f7f002d1986a",
            },
            DiscountPercentage: codataccounting.Float64(6416.3),
            DueDate: codataccounting.String("natus"),
            ID: codataccounting.String("9d3a1d32-329e-4458-b7e8-f2ad6bb10e25"),
            InvoiceNumber: codataccounting.String("ipsam"),
            IssueDate: "reiciendis",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c480d6e3-3086-475c-bf18-6856a7e82cdf"),
                        Name: codataccounting.String("Sammy Barrows"),
                    },
                    Description: codataccounting.String("fugit"),
                    DiscountAmount: codataccounting.Float64(5376.88),
                    DiscountPercentage: codataccounting.Float64(1666.02),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c666af3c-3f55-489b-aa5d-264e41e2ca84",
                        Name: codataccounting.String("Bobby Crooks"),
                    },
                    Quantity: 684,
                    SubTotal: codataccounting.Float64(2081.12),
                    TaxAmount: codataccounting.Float64(9914.72),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4176.65),
                        ID: codataccounting.String("d9d2ad37-c309-4907-bc10-b68792163e67"),
                        Name: codataccounting.String("Kyle Ledner"),
                    },
                    TotalAmount: codataccounting.Float64(608.42),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "43c0a304-9c3c-4f6c-8276-e7b21bad90d2",
                                Name: codataccounting.String("April Frami"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6c2a10e6-c297-48ec-a56a-5b09227fcc47",
                                Name: codataccounting.String("Wendell Huels"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("esse"),
                            ID: "7bbc57f3-8928-4a86-80c5-8d67d63e4aa5",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "464579cf-c6c0-4e50-bf56-831f1d8ed87b",
                            Name: codataccounting.String("Christy Walter"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "abc986e2-41e4-43b2-b424-17d13e3f62aa",
                            Name: codataccounting.String("Shannon Walker"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e8ab4a9c-492c-45e8-ba5d-4aa4a508bd38",
                            Name: codataccounting.String("Kara Cremin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a8dd71bd-daa3-40b7-b914-49ae69c088d4",
                            Name: codataccounting.String("Cassandra Rice"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1804f423-d543-4935-b377-ac5c9b7e93b6",
                            Name: codataccounting.String("Earl Ruecker"),
                        },
                    },
                    UnitAmount: 2263.68,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("105e7c34-cab0-4ecb-812a-66148944a8e9"),
                        Name: codataccounting.String("Ms. Jennie Hartmann"),
                    },
                    Description: codataccounting.String("nam"),
                    DiscountAmount: codataccounting.Float64(7904.63),
                    DiscountPercentage: codataccounting.Float64(1817.55),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "53825334-3fb0-4a4e-a6ea-47578d171e29",
                        Name: codataccounting.String("Ms. Christine Leannon"),
                    },
                    Quantity: 7630.13,
                    SubTotal: codataccounting.Float64(3993.59),
                    TaxAmount: codataccounting.Float64(4504.05),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5874.59),
                        ID: codataccounting.String("b6b2f253-59b8-455d-815b-62c8b83a38a8"),
                        Name: codataccounting.String("Dwayne MacGyver I"),
                    },
                    TotalAmount: codataccounting.Float64(2857.2),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "00c2caeb-1ae1-4ecf-8c34-946bba7a05a8",
                                Name: codataccounting.String("Oscar O'Connell"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("impedit"),
                            ID: "5b3688cc-a363-4272-b60e-966e97e05410",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "47d78ff2-4911-445f-ab9e-59a4af336664",
                            Name: codataccounting.String("Gerard O'Conner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2ff14e8c-1b35-42ac-8eda-cc5227814eca",
                            Name: codataccounting.String("Alice Kautzer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "41ea1342-d410-44a2-9ef7-1de57a11d614",
                            Name: codataccounting.String("Ms. Eddie Frami"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92ea4867-3d52-42b8-a8a9-030660f024c7",
                            Name: codataccounting.String("Abraham Goyette"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "64c2b3a3-2c48-48ad-a62f-6aa558a65e20",
                            Name: codataccounting.String("Mrs. Marvin Armstrong"),
                        },
                    },
                    UnitAmount: 6853.65,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("34bb87d4-f621-427a-a07d-1606294514c3"),
                        Name: codataccounting.String("Preston McCullough"),
                    },
                    Description: codataccounting.String("omnis"),
                    DiscountAmount: codataccounting.Float64(9600.37),
                    DiscountPercentage: codataccounting.Float64(2360.34),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "8bd2be87-8703-4493-b49a-a8465a328327",
                        Name: codataccounting.String("Miss Cesar Konopelski"),
                    },
                    Quantity: 707.2,
                    SubTotal: codataccounting.Float64(7573.22),
                    TaxAmount: codataccounting.Float64(8840.57),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6422.68),
                        ID: codataccounting.String("673d86e3-b35e-449a-b135-778ce54cacb0"),
                        Name: codataccounting.String("Chris Terry"),
                    },
                    TotalAmount: codataccounting.Float64(4960.42),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "045bacf6-3b21-4518-aab5-e3a022614315",
                                Name: codataccounting.String("Dennis Heathcote"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "299e61af-c718-46ff-a0b7-a73df40ca0d7",
                                Name: codataccounting.String("Erin Kihn III"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eius"),
                            ID: "1bbf0552-71b2-4511-9d60-6dd1b28272bc",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3221697b-1880-4fcb-b2b9-3c15f670bd17",
                            Name: codataccounting.String("Francis Labadie III"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3eeb3b6e-241c-4310-9983-663c66dcbb7d",
                            Name: codataccounting.String("Gilbert Schaden V"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c8b408e0-7137-474d-a4fe-e101d9780a10",
                            Name: codataccounting.String("Frederick Kub"),
                        },
                    },
                    UnitAmount: 3444.01,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("040d6c8b-2a5f-4002-a07e-4048f90009ed"),
                        Name: codataccounting.String("Violet Baumbach"),
                    },
                    Description: codataccounting.String("quos"),
                    DiscountAmount: codataccounting.Float64(8818.91),
                    DiscountPercentage: codataccounting.Float64(6977.83),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4ae9d641-61e9-4150-8323-b2c09b924771",
                        Name: codataccounting.String("Herman Howe"),
                    },
                    Quantity: 9129.86,
                    SubTotal: codataccounting.Float64(3652.88),
                    TaxAmount: codataccounting.Float64(7053.17),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4498.47),
                        ID: codataccounting.String("ec762664-9d84-4eb9-a4cf-d2276e0b88fb"),
                        Name: codataccounting.String("Brad Stoltenberg"),
                    },
                    TotalAmount: codataccounting.Float64(6278.38),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b6e8dbf8-12f8-43b1-8a6a-9ffc561929cc",
                                Name: codataccounting.String("Tracy Hilll MD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "1395918d-a1d4-48e7-8e3c-f8e1143da930",
                                Name: codataccounting.String("Robin Cummerata"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("aut"),
                            ID: "8af22184-439b-43de-8756-ccce470cd214",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6e6152cf-01d0-4d8c-ba4b-9a5bf935dfe9",
                            Name: codataccounting.String("Debbie Zieme"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1e9c097e-da62-4344-ae1a-9237e9984c80",
                            Name: codataccounting.String("Alexander Koss"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "891923c1-8ca8-4d69-8568-9214fa20207e",
                            Name: codataccounting.String("Marta Murphy I"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8cd7f1bc-2cab-4af7-bc2c-cba4bef0df68",
                            Name: codataccounting.String("Shaun Swift"),
                        },
                    },
                    UnitAmount: 1625.48,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("earum"),
            Note: codataccounting.String("necessitatibus"),
            PaidOnDate: codataccounting.String("quam"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("expedita"),
                        Currency: codataccounting.String("itaque"),
                        CurrencyRate: codataccounting.Float64(40.87),
                        TotalAmount: codataccounting.Float64(4114.07),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("9fb36add-7040-480e-8a3f-c73a5a034b11"),
                            Name: codataccounting.String("Bobbie Mayer"),
                        },
                        Currency: codataccounting.String("ratione"),
                        CurrencyRate: codataccounting.Float64(6253.92),
                        ID: codataccounting.String("fa6987a4-72b7-409a-953e-22301068539c"),
                        Note: codataccounting.String("saepe"),
                        PaidOnDate: codataccounting.String("ipsa"),
                        Reference: codataccounting.String("perspiciatis"),
                        TotalAmount: codataccounting.Float64(2348.29),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("repellendus"),
                    ID: codataccounting.String("10acd15d-8cc3-406b-b86b-3d37bd204a1f"),
                },
            },
            SourceModifiedDate: codataccounting.String("ipsum"),
            Status: shared.InvoiceStatusDraft,
            SubTotal: codataccounting.Float64(378.08),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "rerum": map[string]interface{}{
                        "ex": "voluptatibus",
                    },
                    "voluptas": map[string]interface{}{
                        "odio": "dolorum",
                        "eius": "praesentium",
                    },
                    "corporis": map[string]interface{}{
                        "provident": "quod",
                    },
                },
            },
            TotalAmount: 2245.24,
            TotalDiscount: codataccounting.Float64(2424.79),
            TotalTaxAmount: 4619.68,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5749.23,
                    Name: "Christina Luettgen",
                },
                shared.WithholdingTaxitems{
                    Amount: 5522.12,
                    Name: "Vicki Reilly I",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(784316),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateInvoiceResponse != nil {
        // handle response
    }
}
```

## Delete

﻿The _Delete Invoices_ endpoint allows you to delete a specified Invoice from an accounting platform.

### Process
1. Pass the `{invoiceId}` to the _Delete Invoices_ endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete operation by checking the status of push operation either via
    1. [Push operation webhook](/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
    2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation).

   A `Success` status indicates that the Invoice object was deleted from the accounting platform.
3. (Optional) Check that the Invoice was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting an Invoice from an accounting platform might cause related objects to be modified. For example, if you delete a paid invoice from QuickBooks Online, the invoice is deleted but the payment against that invoice is not. The payment is converted to a payment on account.

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
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## DownloadAttachment

﻿Download invoice attachment.

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
    res, err := s.Invoices.DownloadAttachment(ctx, operations.DownloadInvoicesAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

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
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Get

﻿Get an invoice.

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
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```

## GetAttachment

﻿Get invoice attachment.

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
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

﻿Get create/update invoice model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support creating and updating invoices.

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

## List

﻿Gets the latest invoices for a company, with pagination.

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
        Query: codataccounting.String("nihil"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoices != nil {
        // handle response
    }
}
```

## ListAttachments

﻿List invoice attachments

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
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## Update

﻿Posts an updated invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support updating invoices.
operationId: update-invoice

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
            AdditionalTaxAmount: codataccounting.Float64(9937.18),
            AdditionalTaxPercentage: codataccounting.Float64(8210.46),
            AmountDue: 1600.93,
            Currency: codataccounting.String("odit"),
            CurrencyRate: codataccounting.Float64(3147.32),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("debitis"),
                ID: "47871a88-ed72-4a2d-8af4-158ac2d0f0f5",
            },
            DiscountPercentage: codataccounting.Float64(5522.56),
            DueDate: codataccounting.String("optio"),
            ID: codataccounting.String("3b87b470-40d0-4d98-a9d8-2c5e306f5576"),
            InvoiceNumber: codataccounting.String("maiores"),
            IssueDate: "nemo",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("deb0286d-0bc4-43b1-8ab3-78f2fcff81dd"),
                        Name: codataccounting.String("Ms. Lance Thiel"),
                    },
                    Description: codataccounting.String("repellat"),
                    DiscountAmount: codataccounting.Float64(4819.23),
                    DiscountPercentage: codataccounting.Float64(2892.95),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "ef54c921-6e89-4263-93bb-6fc2c8d27010",
                        Name: codataccounting.String("Duane Pouros"),
                    },
                    Quantity: 6641.93,
                    SubTotal: codataccounting.Float64(8554.62),
                    TaxAmount: codataccounting.Float64(3899.32),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8980.86),
                        ID: codataccounting.String("3e1d9d3b-6603-434a-91aa-1d5d2247de9b"),
                        Name: codataccounting.String("Meredith Greenfelder IV"),
                    },
                    TotalAmount: codataccounting.Float64(514.52),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "768a96bb-3987-4883-98eb-a1bbf7143356",
                                Name: codataccounting.String("Marc Fay"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a164249b-211c-4e46-b951-652b158ca914",
                                Name: codataccounting.String("Johanna Bartoletti"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "632b31ca-d692-4ffc-8745-005e9d3d934e",
                                Name: codataccounting.String("Sheila Jerde"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c388664f-6985-4530-a2e2-aed6aaf863c2",
                                Name: codataccounting.String("Clint Beatty DDS"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("vel"),
                            ID: "9a3d906f-6ebd-45ad-bec7-394f25f634b3",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0714e6be-8c3e-409c-a4d3-42ac299a6e5e",
                            Name: codataccounting.String("Lee Weber I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "02e945f5-3743-4efd-a119-8221f9b1f7d9",
                            Name: codataccounting.String("Amos Wilkinson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9682acee-fb04-4f8c-912c-aabea708ed57",
                            Name: codataccounting.String("Wallace Schultz"),
                        },
                    },
                    UnitAmount: 3304.22,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d460599d-5c33-4495-b6d5-5209e9a2253b"),
                        Name: codataccounting.String("Elena Kreiger"),
                    },
                    Description: codataccounting.String("praesentium"),
                    DiscountAmount: codataccounting.Float64(5156.7),
                    DiscountPercentage: codataccounting.Float64(3967.41),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "eeae5fd4-b39f-48a1-8906-78f13c686d83",
                        Name: codataccounting.String("Amos Roob"),
                    },
                    Quantity: 756.85,
                    SubTotal: codataccounting.Float64(4594.79),
                    TaxAmount: codataccounting.Float64(3137.17),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9721.89),
                        ID: codataccounting.String("fa906ae5-59b7-42eb-a746-030fe18376c2"),
                        Name: codataccounting.String("Merle Strosin"),
                    },
                    TotalAmount: codataccounting.Float64(4381.93),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "790ed0c1-6a7b-4a47-8404-489f6770ef04",
                                Name: codataccounting.String("Miss Brian McCullough"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ba25ee6c-75af-48a6-8a7a-e346e0979e5a",
                                Name: codataccounting.String("Miss Darrel Keeling"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("culpa"),
                            ID: "ca645de9-8675-451a-9cce-61ec2c79a39a",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "a4d5a65b-4d55-462d-9b7d-9e2d6fcf5576",
                            Name: codataccounting.String("Bobbie Stoltenberg"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5c3a8902-82a5-41f4-9cf6-796ed3d724c1",
                            Name: codataccounting.String("Jerald Hand PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "98cce3f7-1660-40da-8e3a-a61c6fe09d85",
                            Name: codataccounting.String("Shelia Hand"),
                        },
                    },
                    UnitAmount: 2018.2,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2c8c7c3c-710e-4167-bd90-5cb4bedef3c1"),
                        Name: codataccounting.String("Pearl Ruecker"),
                    },
                    Description: codataccounting.String("aperiam"),
                    DiscountAmount: codataccounting.Float64(5631.81),
                    DiscountPercentage: codataccounting.Float64(6128.99),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5528250d-cbbc-4d3b-921b-88c1ee5e7a06",
                        Name: codataccounting.String("Dr. Sherry Marks"),
                    },
                    Quantity: 5137.75,
                    SubTotal: codataccounting.Float64(9667.54),
                    TaxAmount: codataccounting.Float64(6470.68),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(140.4),
                        ID: codataccounting.String("b7d17649-26b0-4cf5-a6cb-6ebabe5e0b99"),
                        Name: codataccounting.String("Mrs. Lee Rogahn"),
                    },
                    TotalAmount: codataccounting.Float64(5558.1),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "6a87bb7a-ecbe-4569-970c-b069907f9894",
                                Name: codataccounting.String("Gloria Gottlieb V"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9f01f344-2c61-4be1-b3ba-cde532b6526f",
                                Name: codataccounting.String("Sam Cole"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3fe2859c-e322-4231-be66-64c41d2fba5c",
                                Name: codataccounting.String("Shannon Bahringer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b8d291be-b810-4a2a-a874-9479edd4fcf7",
                                Name: codataccounting.String("Dan Bechtelar"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("blanditiis"),
                            ID: "7f08f392-7107-46a2-8b40-c8f08bff1081",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8f86996c-8e22-4be0-a3cf-47893bd23f86",
                            Name: codataccounting.String("Mary Beatty"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c7834273-caa9-4118-b38f-1b61a331a54d",
                            Name: codataccounting.String("Stephen Bartell"),
                        },
                    },
                    UnitAmount: 2674.45,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f92fed93-9ba8-4f71-a299-2c20ee1228ac"),
                        Name: codataccounting.String("Alberta Shanahan"),
                    },
                    Description: codataccounting.String("dolore"),
                    DiscountAmount: codataccounting.Float64(4824),
                    DiscountPercentage: codataccounting.Float64(8600.36),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "240bc11e-a482-4824-8cc6-a2f0f5b9d3cb",
                        Name: codataccounting.String("Doris Murray"),
                    },
                    Quantity: 5367.04,
                    SubTotal: codataccounting.Float64(4891.64),
                    TaxAmount: codataccounting.Float64(8210.12),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2214.66),
                        ID: codataccounting.String("100e8e2b-9b0d-4746-92a7-c7d1ea0e79fa"),
                        Name: codataccounting.String("Robin Rau"),
                    },
                    TotalAmount: codataccounting.Float64(9824.77),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "79f650b1-e707-4e7e-8396-713bacce072a",
                                Name: codataccounting.String("Ms. Taylor Jacobson IV"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("facere"),
                            ID: "279c10c1-8516-4fd7-8be2-621272628fa5",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "962867e7-2b3a-4650-a4b1-57f9bb6ef72a",
                            Name: codataccounting.String("Brenda Legros PhD"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9b661a7d-ef16-48b6-8cb2-822b4a9850ed",
                            Name: codataccounting.String("Ora Greenholt DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9c4ae551-40e7-4572-ae00-3c2f02941925",
                            Name: codataccounting.String("Kay Runolfsdottir"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "41c999f4-69f6-4f1c-b1a3-f023c669e6a6",
                            Name: codataccounting.String("Mr. Vicki Bartoletti"),
                        },
                    },
                    UnitAmount: 7126.9,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("id"),
            Note: codataccounting.String("consequatur"),
            PaidOnDate: codataccounting.String("quis"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("unde"),
                        Currency: codataccounting.String("quos"),
                        CurrencyRate: codataccounting.Float64(5264.74),
                        TotalAmount: codataccounting.Float64(8122.68),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6720c310-3f1a-440c-8f3e-c8688fd8ec6f"),
                            Name: codataccounting.String("Mr. Michael Feest"),
                        },
                        Currency: codataccounting.String("voluptatibus"),
                        CurrencyRate: codataccounting.Float64(42.61),
                        ID: codataccounting.String("aaaeee00-4eba-47bf-8732-be509c508713"),
                        Note: codataccounting.String("quae"),
                        PaidOnDate: codataccounting.String("a"),
                        Reference: codataccounting.String("eaque"),
                        TotalAmount: codataccounting.Float64(3967.4),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("maiores"),
                        Currency: codataccounting.String("voluptatem"),
                        CurrencyRate: codataccounting.Float64(7324.85),
                        TotalAmount: codataccounting.Float64(7767.95),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e55a8687-143c-4979-85ff-797a5da664b7"),
                            Name: codataccounting.String("Mitchell Kuhn"),
                        },
                        Currency: codataccounting.String("nihil"),
                        CurrencyRate: codataccounting.Float64(2955.37),
                        ID: codataccounting.String("baaa2832-bb65-4862-92a3-1f9b14aa6bde"),
                        Note: codataccounting.String("quo"),
                        PaidOnDate: codataccounting.String("reprehenderit"),
                        Reference: codataccounting.String("repellat"),
                        TotalAmount: codataccounting.Float64(2617.34),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("dolore"),
                    ID: codataccounting.String("232e9a5d-ee1a-4cd7-aa89-981b58fe682e"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("sunt"),
                    ID: codataccounting.String("c2dbe23d-58e8-4247-9122-c9f67678fa27"),
                },
            },
            SourceModifiedDate: codataccounting.String("occaecati"),
            Status: shared.InvoiceStatusSubmitted,
            SubTotal: codataccounting.Float64(5416.5),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "commodi": map[string]interface{}{
                        "dolor": "voluptas",
                        "dolor": "facere",
                    },
                },
            },
            TotalAmount: 6680.48,
            TotalDiscount: codataccounting.Float64(152.97),
            TotalTaxAmount: 4809.17,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 426.64,
                    Name: "Brett Yundt",
                },
                shared.WithholdingTaxitems{
                    Amount: 7200.66,
                    Name: "Ruben Green III",
                },
                shared.WithholdingTaxitems{
                    Amount: 2407.81,
                    Name: "Angel Larkin",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(701054),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateInvoiceResponse != nil {
        // handle response
    }
}
```

## UploadAttachment

﻿Upload invoice attachment.

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
            Content: []byte("rem"),
            RequestBody: "unde",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
