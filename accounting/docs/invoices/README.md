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
            AdditionalTaxAmount: codataccounting.Float64(4832.71),
            AdditionalTaxPercentage: codataccounting.Float64(2115.84),
            AmountDue: 8116.96,
            Currency: codataccounting.String("dignissimos"),
            CurrencyRate: codataccounting.Float64(8445.24),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("placeat"),
                ID: "9efaf43d-c623-4620-b313-8f30df3db022",
            },
            DiscountPercentage: codataccounting.Float64(9381.13),
            DueDate: codataccounting.String("similique"),
            ID: codataccounting.String("a565fb8f-652e-4bb9-9383-838790243b29"),
            InvoiceNumber: codataccounting.String("ratione"),
            IssueDate: "facere",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b30e917f-50fd-4a04-88b1-bb55a292b0bc"),
                        Name: codataccounting.String("Pam Quitzon"),
                    },
                    Description: codataccounting.String("aliquam"),
                    DiscountAmount: codataccounting.Float64(3855.01),
                    DiscountPercentage: codataccounting.Float64(3916.12),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4eb1d033-88b0-4d1b-b17a-fee74b6feb94",
                        Name: codataccounting.String("Colleen Schmidt"),
                    },
                    Quantity: 8490.45,
                    SubTotal: codataccounting.Float64(6734.93),
                    TaxAmount: codataccounting.Float64(9682.72),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2309.03),
                        ID: codataccounting.String("9d16fbf7-6fd1-462b-b03e-3023b93e3431"),
                        Name: codataccounting.String("Nichole Williamson"),
                    },
                    TotalAmount: codataccounting.Float64(7436.12),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "313553cc-f1c2-404c-8adc-c9904c5195b8",
                                Name: codataccounting.String("Leslie Langosh"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fa78f1e2-d3b9-401e-8952-bbb4cbb19f71",
                                Name: codataccounting.String("Meredith Mante"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolore"),
                            ID: "169c1387-271e-418e-a9e4-5118c2cc57fb",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0b1a78ed-29a9-4d4e-aa85-658c2d4f4c88",
                            Name: codataccounting.String("Pat Grant"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8fd9667e-46c5-41d2-bfaa-58dcef234c95",
                            Name: codataccounting.String("Kelli Mosciski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f2190abd-9bbc-4c27-a5ec-2659ce028084",
                            Name: codataccounting.String("Sadie Huel"),
                        },
                    },
                    UnitAmount: 9461.61,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("68e45c8a-ddfa-4c75-8500-430c6632b439"),
                        Name: codataccounting.String("Jeannie Smitham Sr."),
                    },
                    Description: codataccounting.String("nobis"),
                    DiscountAmount: codataccounting.Float64(2362.69),
                    DiscountPercentage: codataccounting.Float64(9294.26),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "91e8f7bc-69d4-460a-b7ec-eb26d10f1ef2",
                        Name: codataccounting.String("Crystal Bernier"),
                    },
                    Quantity: 7867.65,
                    SubTotal: codataccounting.Float64(461.37),
                    TaxAmount: codataccounting.Float64(9654.91),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(163.7),
                        ID: codataccounting.String("f873f9d5-c25f-4d3e-8b4a-4a4253c30257"),
                        Name: codataccounting.String("Rebecca Wunsch"),
                    },
                    TotalAmount: codataccounting.Float64(8074.3),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e7dc548b-e09e-441a-ba21-5ca12a4ba9d5",
                                Name: codataccounting.String("Fredrick Ledner V"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2cfd0c77-c53e-47e7-94ee-6e8b90bac384",
                                Name: codataccounting.String("Shawn Fahey"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("in"),
                            ID: "03fec31c-5082-44d1-89a3-6a6b2d27eb70",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "a60c8fe4-6e61-477d-b9db-3b70ffbb6970",
                            Name: codataccounting.String("Phil Kuhlman DVM"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6097ef7c-206e-461b-8d30-8714c20a3d98",
                            Name: codataccounting.String("Annie Koelpin"),
                        },
                    },
                    UnitAmount: 5213.33,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5c3fe655-74db-4af9-8a7c-98f13af28db2"),
                        Name: codataccounting.String("Timmy Cruickshank"),
                    },
                    Description: codataccounting.String("eius"),
                    DiscountAmount: codataccounting.Float64(9761.21),
                    DiscountPercentage: codataccounting.Float64(2074.59),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "ded356d7-e14b-421c-9981-96d55af69a1c",
                        Name: codataccounting.String("Shelley Kunze"),
                    },
                    Quantity: 9064.68,
                    SubTotal: codataccounting.Float64(2021.61),
                    TaxAmount: codataccounting.Float64(2477.82),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4269.64),
                        ID: codataccounting.String("81c23c39-a7c0-4e17-8b12-c5ba825fe22c"),
                        Name: codataccounting.String("Lloyd Runolfsdottir"),
                    },
                    TotalAmount: codataccounting.Float64(4047.58),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "bfec932a-f681-43d6-9bfe-cec2dd6916f7",
                                Name: codataccounting.String("Roosevelt Koss"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a70ec60e-6075-4894-961c-14cd90227e37",
                                Name: codataccounting.String("Gary Schultz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7f1a5491-abe9-4751-b106-d23e03e69815",
                                Name: codataccounting.String("Rex Towne"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fcde9e72-9c9d-44f2-98a4-4640ca60db73",
                                Name: codataccounting.String("Bobby Will"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("maiores"),
                            ID: "467dc0d8-da56-4122-826a-b8f277485c19",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "af980da7-a089-4fc4-8db2-74530e5cc7c6",
                            Name: codataccounting.String("Richard Schmeler"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "dcd334b6-f623-4bce-8ab5-0aee5e0da8b9",
                            Name: codataccounting.String("Darrin Hudson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "05486e7b-413c-4be2-9176-dc1c43d40f61",
                            Name: codataccounting.String("Mr. Dennis Kilback"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7cbe5ee4-f721-4184-8772-f32e3b49dbe0",
                            Name: codataccounting.String("Billy Franey"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b6d9948d-6ede-4d47-b680-fc7a17a82e5e",
                            Name: codataccounting.String("Craig Windler"),
                        },
                    },
                    UnitAmount: 5097.39,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("fugiat"),
            Note: codataccounting.String("beatae"),
            PaidOnDate: codataccounting.String("perferendis"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("aperiam"),
                        Currency: codataccounting.String("harum"),
                        CurrencyRate: codataccounting.Float64(4794.64),
                        TotalAmount: codataccounting.Float64(8934.34),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("91392ab4-4cb1-4835-808f-461ce53e9144"),
                            Name: codataccounting.String("Wallace O'Keefe"),
                        },
                        Currency: codataccounting.String("deserunt"),
                        CurrencyRate: codataccounting.Float64(2957.26),
                        ID: codataccounting.String("60addfde-410c-437d-aa91-82a49d9625d3"),
                        Note: codataccounting.String("eligendi"),
                        PaidOnDate: codataccounting.String("laborum"),
                        Reference: codataccounting.String("delectus"),
                        TotalAmount: codataccounting.Float64(9680.39),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("minus"),
                        Currency: codataccounting.String("inventore"),
                        CurrencyRate: codataccounting.Float64(5825.21),
                        TotalAmount: codataccounting.Float64(5027.27),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("eea44527-92bc-4d44-8ea9-8becce0486de"),
                            Name: codataccounting.String("Dianna Hintz"),
                        },
                        Currency: codataccounting.String("nihil"),
                        CurrencyRate: codataccounting.Float64(2493.35),
                        ID: codataccounting.String("b005503e-8dc6-426f-b77c-65675f5b70e3"),
                        Note: codataccounting.String("vero"),
                        PaidOnDate: codataccounting.String("eius"),
                        Reference: codataccounting.String("optio"),
                        TotalAmount: codataccounting.Float64(9589.07),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("impedit"),
                    ID: codataccounting.String("6a91ec52-624d-4000-94ef-45cea11ac53e"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("quidem"),
                    ID: codataccounting.String("b6587f34-0414-4c5b-9ace-e400ae9f92ca"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("earum"),
                    ID: codataccounting.String("1b025f1d-1471-48c6-ba2f-ad0c06c5d954"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("in"),
                    ID: codataccounting.String("2cdd14fc-43b7-40bc-a88f-a70c43351c3d"),
                },
            },
            SourceModifiedDate: codataccounting.String("nulla"),
            Status: shared.InvoiceStatusUnknown,
            SubTotal: codataccounting.Float64(8962.8),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "deleniti": map[string]interface{}{
                        "odio": "tenetur",
                        "quam": "nemo",
                        "sapiente": "magnam",
                        "hic": "aspernatur",
                    },
                    "ipsum": map[string]interface{}{
                        "quasi": "cumque",
                        "eaque": "error",
                        "corporis": "totam",
                        "commodi": "maxime",
                    },
                    "non": map[string]interface{}{
                        "repudiandae": "odio",
                        "temporibus": "reprehenderit",
                        "soluta": "voluptas",
                    },
                },
            },
            TotalAmount: 4838.96,
            TotalDiscount: codataccounting.Float64(9698.61),
            TotalTaxAmount: 8864.96,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9451.88,
                    Name: "Gretchen Bergnaum",
                },
                shared.WithholdingTaxitems{
                    Amount: 8642.28,
                    Name: "Dr. Dustin Reilly",
                },
                shared.WithholdingTaxitems{
                    Amount: 9097.17,
                    Name: "Merle Yost",
                },
                shared.WithholdingTaxitems{
                    Amount: 7789.75,
                    Name: "Hannah Bergnaum",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(928900),
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
        InvoiceID: "occaecati",
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
        InvoiceID: "odit",
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
        InvoiceID: "ducimus",
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
        InvoiceID: "corrupti",
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
        InvoiceID: "consequuntur",
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
        Query: codataccounting.String("voluptate"),
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
        InvoiceID: "ipsam",
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
            AdditionalTaxAmount: codataccounting.Float64(8902.78),
            AdditionalTaxPercentage: codataccounting.Float64(9179.07),
            AmountDue: 6424.25,
            Currency: codataccounting.String("esse"),
            CurrencyRate: codataccounting.Float64(3851.06),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("laudantium"),
                ID: "17468063-f799-4b79-96c0-b0fa0bb20a40",
            },
            DiscountPercentage: codataccounting.Float64(9171.02),
            DueDate: codataccounting.String("reprehenderit"),
            ID: codataccounting.String("c4ae6406-4272-4657-b01a-07c08fd3921c"),
            InvoiceNumber: codataccounting.String("dolores"),
            IssueDate: "exercitationem",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("930d6f09-3a3e-4fa4-ad36-6dfa1011a091"),
                        Name: codataccounting.String("Rodney Turcotte"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(3503.06),
                    DiscountPercentage: codataccounting.Float64(2447.13),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "862de1a9-d14f-4e72-a521-f90303dfc338",
                        Name: codataccounting.String("Faye Kozey"),
                    },
                    Quantity: 9437.98,
                    SubTotal: codataccounting.Float64(6258.87),
                    TaxAmount: codataccounting.Float64(4112.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8325.3),
                        ID: codataccounting.String("1d32090f-c157-4ac9-be19-61ce9be41c86"),
                        Name: codataccounting.String("Drew Schultz"),
                    },
                    TotalAmount: codataccounting.Float64(6057.59),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "19d07b20-0a58-4ffd-a967-df8fd882a8e6",
                                Name: codataccounting.String("Hannah Thompson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0cd9c5af-dd04-4c37-9251-2beae1d87ecc",
                                Name: codataccounting.String("Lucia Stoltenberg"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("animi"),
                            ID: "8e7a8831-1662-4cda-ad77-c1d86066237d",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "27866db8-a749-4e39-8451-1cc75e4f0c00",
                            Name: codataccounting.String("Patty Harber"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "58cc9456-2f00-4696-85fc-d1a173d84bbe",
                            Name: codataccounting.String("Miss Emma White"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "34afb073-5cb6-4285-94a2-9aaa1e169156",
                            Name: codataccounting.String("Adrian Schuster"),
                        },
                    },
                    UnitAmount: 9197.03,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("209505bf-03a9-43e9-8480-ca37fb107890"),
                        Name: codataccounting.String("Theresa O'Connell"),
                    },
                    Description: codataccounting.String("amet"),
                    DiscountAmount: codataccounting.Float64(1995.11),
                    DiscountPercentage: codataccounting.Float64(933.23),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "72e2dd79-ec74-4ba7-a88d-db36fd1ccc34",
                        Name: codataccounting.String("Francis Macejkovic"),
                    },
                    Quantity: 4958.43,
                    SubTotal: codataccounting.Float64(2225.2),
                    TaxAmount: codataccounting.Float64(2974.63),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4391.23),
                        ID: codataccounting.String("4f0a740f-b4ab-4441-83a0-9e763995d808"),
                        Name: codataccounting.String("Preston Thompson"),
                    },
                    TotalAmount: codataccounting.Float64(2965.41),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "55ebc550-a1c4-426b-99c8-366fdcc13558",
                                Name: codataccounting.String("Rosalie Borer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "55e889d9-ef93-42e9-800a-13ad8124208e",
                                Name: codataccounting.String("Wilfred Deckow"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dicta"),
                            ID: "1898e738-79ef-4be8-baeb-abb794536e90",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1bb97631-720b-477a-9a53-65a79f15271f",
                            Name: codataccounting.String("Dr. Janet Sanford"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1fed8dc5-effb-4453-a908-9e871fdb4d69",
                            Name: codataccounting.String("Bridget Schumm"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c985e437-34a5-4d72-99ed-d785be5e7afe",
                            Name: codataccounting.String("Cathy D'Amore"),
                        },
                    },
                    UnitAmount: 7352.56,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("mollitia"),
            Note: codataccounting.String("laboriosam"),
            PaidOnDate: codataccounting.String("explicabo"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sunt"),
                        Currency: codataccounting.String("repellat"),
                        CurrencyRate: codataccounting.Float64(3036.95),
                        TotalAmount: codataccounting.Float64(2670.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e3a23394-a68c-4c80-930f-f72164d0a91f"),
                            Name: codataccounting.String("Eduardo Stark"),
                        },
                        Currency: codataccounting.String("minima"),
                        CurrencyRate: codataccounting.Float64(3654.98),
                        ID: codataccounting.String("3b89e000-9c66-492d-a7b3-562201a6aab4"),
                        Note: codataccounting.String("deserunt"),
                        PaidOnDate: codataccounting.String("voluptates"),
                        Reference: codataccounting.String("in"),
                        TotalAmount: codataccounting.Float64(6889.51),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("vitae"),
                        Currency: codataccounting.String("fuga"),
                        CurrencyRate: codataccounting.Float64(3711.71),
                        TotalAmount: codataccounting.Float64(6917.11),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("908d4e30-491a-4a35-94a8-39f03bab77b9"),
                            Name: codataccounting.String("Mrs. Mattie Wilderman I"),
                        },
                        Currency: codataccounting.String("iste"),
                        CurrencyRate: codataccounting.Float64(5540.29),
                        ID: codataccounting.String("4507e0e3-9c7e-423e-8b06-04652e23a3d6"),
                        Note: codataccounting.String("quo"),
                        PaidOnDate: codataccounting.String("voluptas"),
                        Reference: codataccounting.String("quis"),
                        TotalAmount: codataccounting.Float64(4438.01),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("vero"),
                        Currency: codataccounting.String("unde"),
                        CurrencyRate: codataccounting.Float64(8413.46),
                        TotalAmount: codataccounting.Float64(8936.05),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8f7f002d-1986-4aa9-9d3a-1d32329e4583"),
                            Name: codataccounting.String("Alexis Lindgren"),
                        },
                        Currency: codataccounting.String("officia"),
                        CurrencyRate: codataccounting.Float64(8553.77),
                        ID: codataccounting.String("6bb10e25-5fdc-4480-96e3-308675cbf186"),
                        Note: codataccounting.String("molestias"),
                        PaidOnDate: codataccounting.String("nostrum"),
                        Reference: codataccounting.String("vel"),
                        TotalAmount: codataccounting.Float64(6468.22),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("officiis"),
                    ID: codataccounting.String("82cdf9d0-fc28-42c6-a6af-3c3f5589bea5"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("possimus"),
                    ID: codataccounting.String("264e41e2-ca84-4822-a513-f6d9d2ad37c3"),
                },
            },
            SourceModifiedDate: codataccounting.String("eaque"),
            Status: shared.InvoiceStatusPartiallyPaid,
            SubTotal: codataccounting.Float64(5821.15),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "iure": map[string]interface{}{
                        "nobis": "quae",
                        "sit": "rerum",
                    },
                },
            },
            TotalAmount: 3867.85,
            TotalDiscount: codataccounting.Float64(5362.23),
            TotalTaxAmount: 4770.94,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 1523.59,
                    Name: "Lucy Fahey",
                },
                shared.WithholdingTaxitems{
                    Amount: 4686.34,
                    Name: "Kyle Ledner",
                },
                shared.WithholdingTaxitems{
                    Amount: 608.42,
                    Name: "Eleanor Feeney MD",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "adipisci",
        TimeoutInMinutes: codataccounting.Int(23984),
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
            Content: []byte("labore"),
            RequestBody: "excepturi",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "quisquam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
