# Invoices

## Overview

Invoices

### Available Operations

* [DownloadInvoicePdf](#downloadinvoicepdf) - Get invoice as PDF
* [CreateInvoice](#createinvoice) - Create invoice
* [DeleteInvoice](#deleteinvoice) - Delete invoice
* [DownloadInvoiceAttachment](#downloadinvoiceattachment) - Download invoice attachment
* [GetCreateUpdateInvoicesModel](#getcreateupdateinvoicesmodel) - Get create/update invoice model
* [GetInvoice](#getinvoice) - Get invoice
* [GetInvoiceAttachment](#getinvoiceattachment) - Get invoice attachment
* [GetInvoiceAttachments](#getinvoiceattachments) - Get invoice attachments
* [ListInvoices](#listinvoices) - List invoices
* [UpdateInvoice](#updateinvoice) - Update invoice
* [UploadInvoiceAttachment](#uploadinvoiceattachment) - Push invoice attachment

## DownloadInvoicePdf

Get invoice as PDF

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DownloadInvoicePdfRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.DownloadInvoicePdf(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## CreateInvoice

Posts a new invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support creating invoices.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(4724.55),
            AdditionalTaxPercentage: codataccounting.Float64(6309.47),
            AmountDue: 6663.31,
            Currency: codataccounting.String("iure"),
            CurrencyRate: codataccounting.Float64(141.26),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("placeat"),
                ID: "8fe46e61-77db-49db-bb70-ffbb6970ee77",
            },
            DiscountPercentage: codataccounting.Float64(580.86),
            DueDate: codataccounting.String("eveniet"),
            ID: codataccounting.String("36097ef7-c206-4e61-b0d3-08714c20a3d9"),
            InvoiceNumber: codataccounting.String("corrupti"),
            IssueDate: "ea",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7ca85c3f-e655-474d-baf9-4a7c98f13af2"),
                        Name: codataccounting.String("Mack Rempel"),
                    },
                    Description: codataccounting.String("sapiente"),
                    DiscountAmount: codataccounting.Float64(1646.09),
                    DiscountPercentage: codataccounting.Float64(7103.52),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f4f3ded3-56d7-4e14-b21c-d98196d55af6",
                        Name: codataccounting.String("Hubert Casper"),
                    },
                    Quantity: 7070.73,
                    SubTotal: codataccounting.Float64(4962.61),
                    TaxAmount: codataccounting.Float64(6175.35),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6847.46),
                        ID: codataccounting.String("e33681c2-3c39-4a7c-8e17-cb12c5ba825f"),
                        Name: codataccounting.String("Nicholas Dare"),
                    },
                    TotalAmount: codataccounting.Float64(3186.14),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ba6fbfec-932a-4f68-93d6-5bfecec2dd69",
                                Name: codataccounting.String("Eileen Ziemann"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c7dda70e-c60e-4607-9894-d61c14cd9022",
                                Name: codataccounting.String("Sophie Doyle"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0d977f1a-5491-4abe-9751-b106d23e03e6",
                                Name: codataccounting.String("Dwight Casper"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ae99fcde-9e72-49c9-94f2-d8a44640ca60",
                                Name: codataccounting.String("Geoffrey Kovacek"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consequuntur"),
                            ID: "f93f467d-c0d8-4da5-a122-026ab8f27748",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1976af98-0da7-4a08-9fc4-4db274530e5c",
                            Name: codataccounting.String("Jared Schinner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cbcfdcd3-34b6-4f62-bbce-cab50aee5e0d",
                            Name: codataccounting.String("Clifton Rippin"),
                        },
                    },
                    UnitAmount: 9594.77,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nisi"),
            Note: codataccounting.String("id"),
            PaidOnDate: codataccounting.String("nulla"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ullam"),
                        Currency: codataccounting.String("incidunt"),
                        CurrencyRate: codataccounting.Float64(5382.58),
                        TotalAmount: codataccounting.Float64(3905.07),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e7b413cb-e2d1-476d-81c4-3d40f61d1711"),
                            Name: codataccounting.String("Joy Runolfsdottir"),
                        },
                        Currency: codataccounting.String("enim"),
                        CurrencyRate: codataccounting.Float64(9324.15),
                        ID: codataccounting.String("e4f72118-4077-42f3-ae3b-49dbe0f23b7b"),
                        Note: codataccounting.String("suscipit"),
                        PaidOnDate: codataccounting.String("fugiat"),
                        Reference: codataccounting.String("perspiciatis"),
                        TotalAmount: codataccounting.Float64(5855.5),
                    },
                },
            },
            SalesOrderRefs: []string{
                "quas",
                "assumenda",
            },
            SourceModifiedDate: codataccounting.String("aliquid"),
            Status: shared.InvoiceStatusEnumVoid,
            SubTotal: codataccounting.Float64(8165.56),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "fugiat": map[string]interface{}{
                        "voluptate": "odio",
                        "voluptas": "deleniti",
                    },
                    "eaque": map[string]interface{}{
                        "minus": "iure",
                        "laborum": "ab",
                        "iure": "deserunt",
                        "blanditiis": "dolores",
                    },
                    "necessitatibus": map[string]interface{}{
                        "vero": "totam",
                        "eos": "delectus",
                    },
                    "illum": map[string]interface{}{
                        "praesentium": "fugiat",
                    },
                },
            },
            TotalAmount: 1061.18,
            TotalDiscount: codataccounting.Float64(187.83),
            TotalTaxAmount: 3114.84,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 6869.46,
                    Name: "Mr. Henrietta Marquardt",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(185041),
    }

    res, err := s.Invoices.CreateInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateInvoiceResponse != nil {
        // handle response
    }
}
```

## DeleteInvoice

Deletes an invoice from the accounting package for a given company.

> **Supported Integrations**
> 
> This functionality is currently only supported for our QuickBooks Online integration. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DeleteInvoiceRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.DeleteInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## DownloadInvoiceAttachment

Download invoice attachments

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DownloadInvoiceAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.DownloadInvoiceAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetCreateUpdateInvoicesModel

Get create/update invoice model. Returns the expected data for the request payload.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateUpdateInvoicesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Invoices.GetCreateUpdateInvoicesModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetInvoice

Get invoice

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetInvoiceRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.GetInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```

## GetInvoiceAttachment

Get invoice attachment

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetInvoiceAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.GetInvoiceAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetInvoiceAttachments

Get invoice attachments

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetInvoiceAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.GetInvoiceAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## ListInvoices

Gets the latest invoices for a company, with pagination

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListInvoicesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dolorum"),
    }

    res, err := s.Invoices.ListInvoices(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoices != nil {
        // handle response
    }
}
```

## UpdateInvoice

Posts an updated invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support updating invoices.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(7197.2),
            AdditionalTaxPercentage: codataccounting.Float64(2761.77),
            AmountDue: 3084.29,
            Currency: codataccounting.String("eligendi"),
            CurrencyRate: codataccounting.Float64(7102.56),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("architecto"),
                ID: "835008f4-61ce-453e-9144-98a9ba460add",
            },
            DiscountPercentage: codataccounting.Float64(9692.94),
            DueDate: codataccounting.String("temporibus"),
            ID: codataccounting.String("e410c37d-aa91-482a-89d9-625d3caffc19"),
            InvoiceNumber: codataccounting.String("blanditiis"),
            IssueDate: "voluptates",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a4452792-bcd4-440e-a98b-ecce0486de0d"),
                        Name: codataccounting.String("Jeanne Skiles"),
                    },
                    Description: codataccounting.String("distinctio"),
                    DiscountAmount: codataccounting.Float64(121.81),
                    DiscountPercentage: codataccounting.Float64(612.49),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5503e8dc-626f-4f77-8656-75f5b70e3e4c",
                        Name: codataccounting.String("Gregg Russel"),
                    },
                    Quantity: 5828.22,
                    SubTotal: codataccounting.Float64(696.5),
                    TaxAmount: codataccounting.Float64(8969.21),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8038.15),
                        ID: codataccounting.String("52624d00-014e-4f45-8ea1-1ac53ebb6587"),
                        Name: codataccounting.String("Mrs. Jeffery Gerlach II"),
                    },
                    TotalAmount: codataccounting.Float64(8117.74),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b9acee40-0ae9-4f92-8af1-b025f1d14718",
                                Name: codataccounting.String("Charlie Wolf DVM"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ad0c06c5-d954-472c-9d14-fc43b70bca88",
                                Name: codataccounting.String("Miss Ernesto Kulas"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consectetur"),
                            ID: "351c3dd1-eb8f-47f7-9f4f-23f1c0a586c3",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7d7b67fe-ef5e-4142-995b-1dbeceff7c4b",
                            Name: codataccounting.String("Dolores Jerde"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "78275eea-7681-4746-8063-f799b7956c0b",
                            Name: codataccounting.String("Miss Mindy O'Keefe"),
                        },
                    },
                    UnitAmount: 1799.04,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0a40e7c4-ae64-4064-a726-57b01a07c08f"),
                        Name: codataccounting.String("Jimmy Morar DDS"),
                    },
                    Description: codataccounting.String("dolores"),
                    DiscountAmount: codataccounting.Float64(3486.63),
                    DiscountPercentage: codataccounting.Float64(4420.79),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "930d6f09-3a3e-4fa4-ad36-6dfa1011a091",
                        Name: codataccounting.String("Rodney Turcotte"),
                    },
                    Quantity: 7336.81,
                    SubTotal: codataccounting.Float64(3503.06),
                    TaxAmount: codataccounting.Float64(2447.13),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5203.56),
                        ID: codataccounting.String("62de1a9d-14fe-472e-921f-90303dfc3383"),
                        Name: codataccounting.String("Javier Wisozk"),
                    },
                    TotalAmount: codataccounting.Float64(6258.87),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d1d32090-fc15-47ac-9fe1-961ce9be41c8",
                                Name: codataccounting.String("Belinda Stoltenberg"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "d9719d07-b200-4a58-bfd2-967df8fd882a",
                                Name: codataccounting.String("Miss Tomas Hodkiewicz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("commodi"),
                            ID: "20cd9c5a-fdd0-44c3-b525-12beae1d87ec",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "fdcea8e7-a883-4116-a2cd-a6d77c1d8606",
                            Name: codataccounting.String("Kathy Douglas"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "227866db-8a74-49e3-9845-11cc75e4f0c0",
                            Name: codataccounting.String("Ellen Rodriguez"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b758cc94-562f-4006-9685-fcd1a173d84b",
                            Name: codataccounting.String("Elias Connelly"),
                        },
                    },
                    UnitAmount: 1277.99,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9834afb0-735c-4b62-85d4-a29aaa1e1691"),
                        Name: codataccounting.String("Alma Ziemann"),
                    },
                    Description: codataccounting.String("aspernatur"),
                    DiscountAmount: codataccounting.Float64(9264.79),
                    DiscountPercentage: codataccounting.Float64(9197.03),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "209505bf-03a9-43e9-8480-ca37fb107890",
                        Name: codataccounting.String("Theresa O'Connell"),
                    },
                    Quantity: 2286.72,
                    SubTotal: codataccounting.Float64(1995.11),
                    TaxAmount: codataccounting.Float64(933.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4630.78),
                        ID: codataccounting.String("2e2dd79e-c74b-4a7e-88dd-b36fd1ccc341"),
                        Name: codataccounting.String("Willard Johnston"),
                    },
                    TotalAmount: codataccounting.Float64(2225.2),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "74f0a740-fb4a-4b44-9c3a-09e763995d80",
                                Name: codataccounting.String("Garry Reichel"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "94455ebc-550a-41c4-a6b5-9c8366fdcc13",
                                Name: codataccounting.String("Ida Lemke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("et"),
                            ID: "b855e889-d9ef-4932-a900-0a13ad812420",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "fd234118-98e7-4387-9efb-e8baebabb794",
                            Name: codataccounting.String("Rita Keeling"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "351bb976-3172-40b7-ba5a-5365a79f1527",
                            Name: codataccounting.String("Miss Winifred Barton PhD"),
                        },
                    },
                    UnitAmount: 2078.87,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("61fed8dc-5eff-4b45-be90-89e871fdb4d6"),
                        Name: codataccounting.String("Fernando Rippin"),
                    },
                    Description: codataccounting.String("perspiciatis"),
                    DiscountAmount: codataccounting.Float64(7911.41),
                    DiscountPercentage: codataccounting.Float64(5639.57),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "85e43734-a5d7-42d9-add7-85be5e7afe55",
                        Name: codataccounting.String("Kristina Kozey"),
                    },
                    Quantity: 3860.49,
                    SubTotal: codataccounting.Float64(1265.14),
                    TaxAmount: codataccounting.Float64(5110.98),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1221.01),
                        ID: codataccounting.String("f44e3a23-394a-468c-880d-30ff72164d0a"),
                        Name: codataccounting.String("Keith Zulauf"),
                    },
                    TotalAmount: codataccounting.Float64(8512.2),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "6553b89e-0009-4c66-92de-7b3562201a6a",
                                Name: codataccounting.String("Bennie Gislason"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7b1a5b90-8d4e-4304-91aa-35d4a839f03b",
                                Name: codataccounting.String("Willis Krajcik"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "918f0313-9845-407e-8e39-c7e23ecb0604",
                                Name: codataccounting.String("Melanie Corkery"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ipsum"),
                            ID: "a3d6c657-e9de-48f7-b002-d1986aa99d3a",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "32329e45-837e-48f2-ad6b-b10e255fdc48",
                            Name: codataccounting.String("Krystal Hoeger"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "08675cbf-1868-456a-be82-cdf9d0fc282c",
                            Name: codataccounting.String("Eileen Kassulke"),
                        },
                    },
                    UnitAmount: 1988.04,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("optio"),
            Note: codataccounting.String("ratione"),
            PaidOnDate: codataccounting.String("a"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nostrum"),
                        Currency: codataccounting.String("totam"),
                        CurrencyRate: codataccounting.Float64(5904.76),
                        TotalAmount: codataccounting.Float64(7165.27),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ea5d264e-41e2-4ca8-8822-e513f6d9d2ad"),
                            Name: codataccounting.String("Nellie Schmeler V"),
                        },
                        Currency: codataccounting.String("occaecati"),
                        CurrencyRate: codataccounting.Float64(328.36),
                        ID: codataccounting.String("77c10b68-7921-463e-a7d4-8860543c0a30"),
                        Note: codataccounting.String("labore"),
                        PaidOnDate: codataccounting.String("excepturi"),
                        Reference: codataccounting.String("quisquam"),
                        TotalAmount: codataccounting.Float64(2173.4),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quod"),
                        Currency: codataccounting.String("voluptatibus"),
                        CurrencyRate: codataccounting.Float64(3769.3),
                        TotalAmount: codataccounting.Float64(7836.87),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0276e7b2-1bad-490d-a743-fd6c2a10e6c2"),
                            Name: codataccounting.String("Erik Lindgren"),
                        },
                        Currency: codataccounting.String("eos"),
                        CurrencyRate: codataccounting.Float64(3181.5),
                        ID: codataccounting.String("6a5b0922-7fcc-4479-96c9-77bbc57f3892"),
                        Note: codataccounting.String("quos"),
                        PaidOnDate: codataccounting.String("est"),
                        Reference: codataccounting.String("blanditiis"),
                        TotalAmount: codataccounting.Float64(3890.14),
                    },
                },
            },
            SalesOrderRefs: []string{
                "eaque",
            },
            SourceModifiedDate: codataccounting.String("quo"),
            Status: shared.InvoiceStatusEnumDraft,
            SubTotal: codataccounting.Float64(5387.08),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "eum": map[string]interface{}{
                        "facere": "ea",
                        "sequi": "voluptates",
                    },
                    "tempora": map[string]interface{}{
                        "officia": "exercitationem",
                        "laboriosam": "quos",
                        "aliquam": "vel",
                    },
                    "numquam": map[string]interface{}{
                        "odio": "omnis",
                        "quo": "maiores",
                    },
                    "maxime": map[string]interface{}{
                        "quisquam": "eaque",
                        "officiis": "corporis",
                    },
                },
            },
            TotalAmount: 507.41,
            TotalDiscount: codataccounting.Float64(2229.84),
            TotalTaxAmount: 9822.17,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4264.61,
                    Name: "Johnny Bins PhD",
                },
                shared.WithholdingTaxitems{
                    Amount: 5621.97,
                    Name: "Al Ledner",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(185898),
    }

    res, err := s.Invoices.UpdateInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateInvoiceResponse != nil {
        // handle response
    }
}
```

## UploadInvoiceAttachment

Push invoice attachment

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UploadInvoiceAttachmentRequest{
        RequestBody: &operations.UploadInvoiceAttachmentRequestBody{
            Content: []byte("totam"),
            RequestBody: "recusandae",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Invoices.UploadInvoiceAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
