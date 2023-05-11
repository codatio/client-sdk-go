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
    res, err := s.Invoices.Create(ctx, operations.CreateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(3856.2),
            AdditionalTaxPercentage: codataccounting.Float64(6970.56),
            AmountDue: 4400.63,
            Currency: codataccounting.String("praesentium"),
            CurrencyRate: codataccounting.Float64(5207.16),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("error"),
                ID: "0a3fd3c8-1da1-40f8-823d-f931da3edb51",
            },
            DiscountPercentage: codataccounting.Float64(9961.99),
            DueDate: codataccounting.String("est"),
            ID: codataccounting.String("d94acc94-3513-4772-ad15-321b832a56d6"),
            InvoiceNumber: codataccounting.String("sint"),
            IssueDate: "architecto",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0ff60eb9-a665-48e6-9a4b-843d382dbec7"),
                        Name: codataccounting.String("Jacquelyn Jast"),
                    },
                    Description: codataccounting.String("autem"),
                    DiscountAmount: codataccounting.Float64(163.03),
                    DiscountPercentage: codataccounting.Float64(3914.99),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "59468ce3-04d8-4849-bf82-14c337f96bb0",
                        Name: codataccounting.String("Cecil Mohr"),
                    },
                    Quantity: 4625.93,
                    SubTotal: codataccounting.Float64(1436.68),
                    TaxAmount: codataccounting.Float64(8176.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7291.09),
                        ID: codataccounting.String("1344ba9f-78a5-4c0e-97aa-b62e97261fb0"),
                        Name: codataccounting.String("Jerome Lowe"),
                    },
                    TotalAmount: codataccounting.Float64(4514.48),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "51996b5b-4b50-4eef-b12b-7a7ab0344b17",
                                Name: codataccounting.String("Nancy Johnson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "deebef89-7f3d-4d0c-8d33-f11b3e4e080a",
                                Name: codataccounting.String("Terry Bednar IV"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6ec759e0-2f37-402c-9c8e-2d30ead3104f",
                                Name: codataccounting.String("Theodore Gerlach IV"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("rerum"),
                            ID: "f375b442-8282-41fd-b2f6-9e59267c71cc",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3cd4258d-0358-4a82-8808-fe2751a2047c",
                            Name: codataccounting.String("Marjorie Funk"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "43f9619b-b7d4-40d5-a11f-a436e6259233",
                            Name: codataccounting.String("Luther Hane"),
                        },
                    },
                    UnitAmount: 8428.91,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("237397c7-85b5-4db4-b500-183febdf676b"),
                        Name: codataccounting.String("Andrea Bashirian"),
                    },
                    Description: codataccounting.String("deserunt"),
                    DiscountAmount: codataccounting.Float64(7054.18),
                    DiscountPercentage: codataccounting.Float64(4503.12),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "50052a56-47ed-4c43-9ed8-c4320f41240d",
                        Name: codataccounting.String("Eva Lebsack"),
                    },
                    Quantity: 7842.87,
                    SubTotal: codataccounting.Float64(3778.95),
                    TaxAmount: codataccounting.Float64(5909.98),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2042.83),
                        ID: codataccounting.String("b94c3b9d-2488-4d79-9aa4-2fc405669f69"),
                        Name: codataccounting.String("Frank Batz"),
                    },
                    TotalAmount: codataccounting.Float64(1466.93),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "24945081-9d7c-43b1-b418-44060e00310d",
                                Name: codataccounting.String("Norma Frami"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("error"),
                            ID: "01f5afd2-a6c4-4484-aae9-d89253c8962f",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "96bf51e4-652d-43c3-83d6-1778af491247",
                            Name: codataccounting.String("Anne Hamill"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1909e910-44a5-4de5-9ac7-706670cf1cf5",
                            Name: codataccounting.String("Jeffery Dibbert II"),
                        },
                    },
                    UnitAmount: 1280.87,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("51e66bb4-2689-47d9-9a2d-335670e93ee6"),
                        Name: codataccounting.String("Dominick Hammes"),
                    },
                    Description: codataccounting.String("dolorem"),
                    DiscountAmount: codataccounting.Float64(3687.85),
                    DiscountPercentage: codataccounting.Float64(5295.29),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "aaeacae3-23a3-41bf-bba1-cc97716c802c",
                        Name: codataccounting.String("Dr. Kirk Veum"),
                    },
                    Quantity: 8218.96,
                    SubTotal: codataccounting.Float64(6105.84),
                    TaxAmount: codataccounting.Float64(8294.02),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2051.2),
                        ID: codataccounting.String("23f1aa63-ed9c-4f1c-856b-cba51ef2454a"),
                        Name: codataccounting.String("Marlene Wolf"),
                    },
                    TotalAmount: codataccounting.Float64(9973.54),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "16cdd544-4a75-4628-b3c7-dd9efaf43dc6",
                                Name: codataccounting.String("Victoria Jenkins DVM"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nesciunt"),
                            ID: "138f30df-3db0-422f-aa56-5fb8f652ebb9",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "83838790-243b-4293-9ab3-0e917f50fda0",
                            Name: codataccounting.String("Francis Lesch MD"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "55a292b0-bc3b-4b74-8664-eb1d03388b0d",
                            Name: codataccounting.String("Ms. Jeannette Price"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fee74b6f-eb94-457c-beda-f39d16fbf76f",
                            Name: codataccounting.String("Andrew Kerluke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "303e3023-b93e-4343-96cf-55b4313553cc",
                            Name: codataccounting.String("Justin Schaefer II"),
                        },
                    },
                    UnitAmount: 8084.57,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("labore"),
            Note: codataccounting.String("culpa"),
            PaidOnDate: codataccounting.String("illum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("minus"),
                        Currency: codataccounting.String("sint"),
                        CurrencyRate: codataccounting.Float64(5786.36),
                        TotalAmount: codataccounting.Float64(286.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4c5195b8-648c-4efa-b8f1-e2d3b901e095"),
                            Name: codataccounting.String("Beulah Pouros"),
                        },
                        Currency: codataccounting.String("minus"),
                        CurrencyRate: codataccounting.Float64(6917.42),
                        ID: codataccounting.String("b19f713d-95a4-4169-8138-7271e18ea9e4"),
                        Note: codataccounting.String("veniam"),
                        PaidOnDate: codataccounting.String("illo"),
                        Reference: codataccounting.String("illo"),
                        TotalAmount: codataccounting.Float64(5361.81),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quisquam"),
                        Currency: codataccounting.String("fugit"),
                        CurrencyRate: codataccounting.Float64(7587.36),
                        TotalAmount: codataccounting.Float64(7784.03),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("57fbd60b-1a78-4ed2-9a9d-4eea85658c2d"),
                            Name: codataccounting.String("Dixie Hackett"),
                        },
                        Currency: codataccounting.String("quas"),
                        CurrencyRate: codataccounting.Float64(6955.11),
                        ID: codataccounting.String("e4f278fd-9667-4e46-851d-2ffaa58dcef2"),
                        Note: codataccounting.String("ratione"),
                        PaidOnDate: codataccounting.String("quaerat"),
                        Reference: codataccounting.String("minus"),
                        TotalAmount: codataccounting.Float64(6170.85),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nostrum"),
                        Currency: codataccounting.String("veniam"),
                        CurrencyRate: codataccounting.Float64(7312.33),
                        TotalAmount: codataccounting.Float64(6180.63),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("bdf2190a-bd9b-4bcc-a725-ec2659ce0280"),
                            Name: codataccounting.String("Jesus Abernathy"),
                        },
                        Currency: codataccounting.String("excepturi"),
                        CurrencyRate: codataccounting.Float64(9071.7),
                        ID: codataccounting.String("f68e45c8-addf-4ac7-9450-0430c6632b43"),
                        Note: codataccounting.String("provident"),
                        PaidOnDate: codataccounting.String("inventore"),
                        Reference: codataccounting.String("sapiente"),
                        TotalAmount: codataccounting.Float64(8387.98),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sapiente"),
                        Currency: codataccounting.String("ipsa"),
                        CurrencyRate: codataccounting.Float64(1140.37),
                        TotalAmount: codataccounting.Float64(7509.59),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("3e91e8f7-bc69-4d46-8a77-eceb26d10f1e"),
                            Name: codataccounting.String("Alan Howell DDS"),
                        },
                        Currency: codataccounting.String("nihil"),
                        CurrencyRate: codataccounting.Float64(7867.65),
                        ID: codataccounting.String("0f0f873f-9d5c-425f-93e0-b4a4a4253c30"),
                        Note: codataccounting.String("consequuntur"),
                        PaidOnDate: codataccounting.String("ullam"),
                        Reference: codataccounting.String("molestiae"),
                        TotalAmount: codataccounting.Float64(1037.45),
                    },
                },
            },
            SalesOrderRefs: []string{
                "maiores",
            },
            SourceModifiedDate: codataccounting.String("labore"),
            Status: shared.InvoiceStatusEnumUnknown,
            SubTotal: codataccounting.Float64(8074.3),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "recusandae": map[string]interface{}{
                        "pariatur": "porro",
                        "enim": "tempora",
                    },
                    "voluptatum": map[string]interface{}{
                        "itaque": "sit",
                        "excepturi": "recusandae",
                        "numquam": "architecto",
                    },
                },
            },
            TotalAmount: 6845.98,
            TotalDiscount: codataccounting.Float64(4808.29),
            TotalTaxAmount: 6330.56,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 816.73,
                    Name: "Mr. Leticia Nitzsche",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(289108),
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
    res, err := s.Invoices.DownloadAttachment(ctx, operations.DownloadInvoiceAttachmentRequest{
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
    res, err := s.Invoices.List(ctx, operations.ListInvoicesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("harum"),
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

List invoice attachments

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
    res, err := s.Invoices.Update(ctx, operations.UpdateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(6791.83),
            AdditionalTaxPercentage: codataccounting.Float64(5932.05),
            AmountDue: 8445.45,
            Currency: codataccounting.String("ipsam"),
            CurrencyRate: codataccounting.Float64(5928.98),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("omnis"),
                ID: "88192cfd-0c77-4c53-a7e7-d4ee6e8b90ba",
            },
            DiscountPercentage: codataccounting.Float64(7821.55),
            DueDate: codataccounting.String("consectetur"),
            ID: codataccounting.String("84e23967-03fe-4c31-8508-24d189a36a6b"),
            InvoiceNumber: codataccounting.String("sunt"),
            IssueDate: "facere",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7eb707aa-60c8-4fe4-ae61-77db9db3b70f"),
                        Name: codataccounting.String("Jonathon Quigley"),
                    },
                    Description: codataccounting.String("ducimus"),
                    DiscountAmount: codataccounting.Float64(451.76),
                    DiscountPercentage: codataccounting.Float64(9196.52),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "e770e360-97ef-47c2-86e6-1b0d308714c2",
                        Name: codataccounting.String("Genevieve Erdman"),
                    },
                    Quantity: 5444.06,
                    SubTotal: codataccounting.Float64(4105.74),
                    TaxAmount: codataccounting.Float64(1937.94),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4604.15),
                        ID: codataccounting.String("ca85c3fe-6557-44db-af94-a7c98f13af28"),
                        Name: codataccounting.String("Pete Crona"),
                    },
                    TotalAmount: codataccounting.Float64(1646.09),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f4f3ded3-56d7-4e14-b21c-d98196d55af6",
                                Name: codataccounting.String("Hubert Casper"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b79ae336-81c2-43c3-9a7c-0e17cb12c5ba",
                                Name: codataccounting.String("Steve Herzog"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "22cd5cba-6fbf-4ec9-b2af-6813d65bfece",
                                Name: codataccounting.String("Fred Shields"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("provident"),
                            ID: "16f7fc7d-da70-4ec6-8e60-75894d61c14c",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0227e37c-0d97-47f1-a549-1abe9751b106",
                            Name: codataccounting.String("Clarence Dicki I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "69815aae-99fc-4de9-a729-c9d4f2d8a446",
                            Name: codataccounting.String("Mary Schamberger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0db73a2f-93f4-467d-80d8-da56122026ab",
                            Name: codataccounting.String("Emmett Davis"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "485c1976-af98-40da-ba08-9fc44db27453",
                            Name: codataccounting.String("Rochelle Hermiston"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7c6d0cbc-fdcd-4334-b6f6-23bcecab50ae",
                            Name: codataccounting.String("Dr. Alvin Weber"),
                        },
                    },
                    UnitAmount: 5565.17,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("cum"),
            Note: codataccounting.String("sint"),
            PaidOnDate: codataccounting.String("laborum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nisi"),
                        Currency: codataccounting.String("id"),
                        CurrencyRate: codataccounting.Float64(8611.66),
                        TotalAmount: codataccounting.Float64(202.23),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5486e7b4-13cb-4e2d-976d-c1c43d40f61d"),
                            Name: codataccounting.String("Mrs. Melinda Borer"),
                        },
                        Currency: codataccounting.String("cumque"),
                        CurrencyRate: codataccounting.Float64(7322.16),
                        ID: codataccounting.String("e5ee4f72-1184-4077-af32-e3b49dbe0f23"),
                        Note: codataccounting.String("harum"),
                        PaidOnDate: codataccounting.String("voluptate"),
                        Reference: codataccounting.String("distinctio"),
                        TotalAmount: codataccounting.Float64(3820.07),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("fugiat"),
                        Currency: codataccounting.String("perspiciatis"),
                        CurrencyRate: codataccounting.Float64(5855.5),
                        TotalAmount: codataccounting.Float64(2959.58),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8d6eded4-7768-40fc-ba17-a82e5e82fd28"),
                            Name: codataccounting.String("Albert Auer MD"),
                        },
                        Currency: codataccounting.String("iusto"),
                        CurrencyRate: codataccounting.Float64(8934.34),
                        ID: codataccounting.String("91392ab4-4cb1-4835-808f-461ce53e9144"),
                        Note: codataccounting.String("perspiciatis"),
                        PaidOnDate: codataccounting.String("rem"),
                        Reference: codataccounting.String("animi"),
                        TotalAmount: codataccounting.Float64(6129.79),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("libero"),
                        Currency: codataccounting.String("deserunt"),
                        CurrencyRate: codataccounting.Float64(2957.26),
                        TotalAmount: codataccounting.Float64(3918.99),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0addfde4-10c3-47da-a918-2a49d9625d3c"),
                            Name: codataccounting.String("Toby Wisoky V"),
                        },
                        Currency: codataccounting.String("blanditiis"),
                        CurrencyRate: codataccounting.Float64(9171.68),
                        ID: codataccounting.String("ea445279-2bcd-4440-aa98-becce0486de0"),
                        Note: codataccounting.String("repellendus"),
                        PaidOnDate: codataccounting.String("ipsam"),
                        Reference: codataccounting.String("aliquid"),
                        TotalAmount: codataccounting.Float64(8330.92),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nihil"),
                        Currency: codataccounting.String("non"),
                        CurrencyRate: codataccounting.Float64(7171.48),
                        TotalAmount: codataccounting.Float64(121.81),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("05503e8d-c626-4ff7-bc65-675f5b70e3e4"),
                            Name: codataccounting.String("Darrin Sawayn"),
                        },
                        Currency: codataccounting.String("dolorum"),
                        CurrencyRate: codataccounting.Float64(5828.22),
                        ID: codataccounting.String("1ec52624-d000-414e-b45c-ea11ac53ebb6"),
                        Note: codataccounting.String("nostrum"),
                        PaidOnDate: codataccounting.String("corrupti"),
                        Reference: codataccounting.String("odio"),
                        TotalAmount: codataccounting.Float64(9462.07),
                    },
                },
            },
            SalesOrderRefs: []string{
                "eius",
            },
            SourceModifiedDate: codataccounting.String("voluptatem"),
            Status: shared.InvoiceStatusEnumDraft,
            SubTotal: codataccounting.Float64(647.33),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "placeat": map[string]interface{}{
                        "cum": "sint",
                        "est": "quod",
                    },
                    "voluptates": map[string]interface{}{
                        "non": "quae",
                        "perferendis": "mollitia",
                        "voluptates": "provident",
                        "doloribus": "unde",
                    },
                },
            },
            TotalAmount: 1671.44,
            TotalDiscount: codataccounting.Float64(7933.34),
            TotalTaxAmount: 6618.61,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 1012.53,
                    Name: "Jose D'Amore",
                },
                shared.WithholdingTaxitems{
                    Amount: 1005.96,
                    Name: "Justin Gusikowski IV",
                },
                shared.WithholdingTaxitems{
                    Amount: 8059.91,
                    Name: "Dr. Tami O'Reilly",
                },
                shared.WithholdingTaxitems{
                    Amount: 8418.47,
                    Name: "Erma Barton",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(329849),
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
    res, err := s.Invoices.UploadAttachment(ctx, operations.UploadInvoiceAttachmentRequest{
        RequestBody: &operations.UploadInvoiceAttachmentRequestBody{
            Content: []byte("facere"),
            RequestBody: "excepturi",
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
