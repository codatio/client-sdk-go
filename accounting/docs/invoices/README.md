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
            AdditionalTaxAmount: codataccounting.Float64(4948.15),
            AdditionalTaxPercentage: codataccounting.Float64(777.6),
            AmountDue: 2672.67,
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(1644.32),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("perferendis"),
                ID: "a3d98637-ca85-4c3f-a655-74dbaf94a7c9",
            },
            DiscountPercentage: codataccounting.Float64(5258.6),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("13af28db-2cf2-4bf4-b3de-d356d7e14b21"),
            InvoiceNumber: codataccounting.String("eligendi"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8196d55a-f69a-41c4-b79a-e33681c23c39"),
                        Name: codataccounting.String("Dr. Adrian Sauer III"),
                    },
                    Description: codataccounting.String("optio"),
                    DiscountAmount: codataccounting.Float64(7162.96),
                    DiscountPercentage: codataccounting.Float64(938.38),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "2c5ba825-fe22-4cd5-8ba6-fbfec932af68",
                        Name: codataccounting.String("Sylvia Stanton"),
                    },
                    Quantity: 7210.93,
                    SubTotal: codataccounting.Float64(9644.52),
                    TaxAmount: codataccounting.Float64(9353.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8111.17),
                        ID: codataccounting.String("ec2dd691-6f7f-4c7d-9a70-ec60e6075894"),
                        Name: codataccounting.String("Milton Bogisich I"),
                    },
                    TotalAmount: codataccounting.Float64(7878.73),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "90227e37-c0d9-477f-9a54-91abe9751b10",
                                Name: codataccounting.String("Hope Denesik"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "03e69815-aae9-49fc-9e9e-729c9d4f2d8a",
                                Name: codataccounting.String("Hazel Howe DDS"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a60db73a-2f93-4f46-bdc0-d8da56122026",
                                Name: codataccounting.String("Gerardo Macejkovic"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "77485c19-76af-4980-9a7a-089fc44db274",
                                Name: codataccounting.String("Gladys Beier"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("minus"),
                            ID: "c7c6d0cb-cfdc-4d33-8b6f-623bcecab50a",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "5e0da8b9-af6a-4d05-886e-7b413cbe2d17",
                            Name: codataccounting.String("Miss Krystal Schroeder"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d40f61d1-7115-47cb-a5ee-4f7211840772",
                            Name: codataccounting.String("Chad Connelly"),
                        },
                    },
                    UnitAmount: 7148.8,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("49dbe0f2-3b7b-46d9-948d-6eded477680f"),
                        Name: codataccounting.String("Ms. Brad Pacocha"),
                    },
                    Description: codataccounting.String("blanditiis"),
                    DiscountAmount: codataccounting.Float64(1744.64),
                    DiscountPercentage: codataccounting.Float64(8960.53),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5e82fd28-d104-40a7-a913-92ab44cb1835",
                        Name: codataccounting.String("Sarah Kuvalis"),
                    },
                    Quantity: 3800.12,
                    SubTotal: codataccounting.Float64(780.67),
                    TaxAmount: codataccounting.Float64(7896.36),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8808.82),
                        ID: codataccounting.String("53e91449-8a9b-4a46-8add-fde410c37daa"),
                        Name: codataccounting.String("Lawrence Leuschke"),
                    },
                    TotalAmount: codataccounting.Float64(2651.48),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d9625d3c-affc-4198-aea4-452792bcd440",
                                Name: codataccounting.String("Rex Mayert"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ecce0486-de0d-456d-b3b0-05503e8dc626",
                                Name: codataccounting.String("Emilio Konopelski"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "65675f5b-70e3-4e4c-bcc6-a91ec52624d0",
                                Name: codataccounting.String("Margaret Berge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("hic"),
                            ID: "45cea11a-c53e-4bb6-987f-340414c5b9ac",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "400ae9f9-2caf-41b0-a5f1-d14718c6fa2f",
                            Name: codataccounting.String("Marty Altenwerth III"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5d95472c-dd14-4fc4-bb70-bca88fa70c43",
                            Name: codataccounting.String("Annette Bode"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "dd1eb8f7-f75f-44f2-bf1c-0a586c3ae7d7",
                            Name: codataccounting.String("Shane Kub"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ef5e142d-95b1-4dbe-8eff-7c4b156e9278",
                            Name: codataccounting.String("Colleen Hilpert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a7681746-8063-4f79-9b79-56c0b0fa0bb2",
                            Name: codataccounting.String("Dr. Molly Greenfelder"),
                        },
                    },
                    UnitAmount: 7753.56,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4ae64064-2726-457b-81a0-7c08fd3921c2"),
                        Name: codataccounting.String("Delores Moore PhD"),
                    },
                    Description: codataccounting.String("eum"),
                    DiscountAmount: codataccounting.Float64(9807.15),
                    DiscountPercentage: codataccounting.Float64(248.34),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "93a3efa4-6d36-46df-a101-1a091b3ec8b5",
                        Name: codataccounting.String("Deanna Hudson"),
                    },
                    Quantity: 9003.52,
                    SubTotal: codataccounting.Float64(723.21),
                    TaxAmount: codataccounting.Float64(6823.09),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5648.16),
                        ID: codataccounting.String("d14fe72e-521f-4903-83df-c338397fffa6"),
                        Name: codataccounting.String("Justin Stroman"),
                    },
                    TotalAmount: codataccounting.Float64(102.26),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0fc157ac-9fe1-4961-8e9b-e41c869dd7d9",
                                Name: codataccounting.String("Pamela Mitchell IV"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b200a58f-fd29-467d-b8fd-882a8e60be62",
                                Name: codataccounting.String("Lynne Schultz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5afdd04c-3752-4512-beae-1d87ecc5fdce",
                                Name: codataccounting.String("Alberto Veum"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptatum"),
                            ID: "8311662c-da6d-477c-9d86-066237d42278",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "db8a749e-3984-4511-8c75-e4f0c004b5bb",
                            Name: codataccounting.String("Roberta Luettgen"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4562f006-9685-4fcd-9a17-3d84bbe24f29",
                            Name: codataccounting.String("Earl Gusikowski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b0735cb6-285d-44a2-9aaa-1e169156f7d2",
                            Name: codataccounting.String("Miss Phil Davis"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "05bf03a9-3e94-4480-8a37-fb10789032ac",
                            Name: codataccounting.String("Ms. Wendy Dooley"),
                        },
                    },
                    UnitAmount: 9039.96,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("assumenda"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7923.49),
                        TotalAmount: codataccounting.Float64(4469.67),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4ba7e88d-db36-4fd1-8cc3-41c86573474f"),
                            Name: codataccounting.String("Molly Kling DVM"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2791.66),
                        ID: codataccounting.String("ab441c3a-09e7-4639-95d8-08bbe794455e"),
                        Note: codataccounting.String("nobis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("nemo"),
                        TotalAmount: codataccounting.Float64(3290.47),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1055.18),
                        TotalAmount: codataccounting.Float64(7693.16),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("426b59c8-366f-4dcc-9355-82c1b855e889"),
                            Name: codataccounting.String("Arturo Towne"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1451.48),
                        ID: codataccounting.String("e9000a13-ad81-4242-88ef-d23411898e73"),
                        Note: codataccounting.String("praesentium"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("unde"),
                        TotalAmount: codataccounting.Float64(9076.15),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("facilis"),
                    ID: codataccounting.String("e8baebab-b794-4536-a903-51bb97631720"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("quidem"),
                    ID: codataccounting.String("77a5a536-5a79-4f15-a71f-01c0d361fed8"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("nulla"),
                    ID: codataccounting.String("c5effb45-3e90-489e-871f-db4d697bdd9c"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("molestias"),
                    ID: codataccounting.String("85e43734-a5d7-42d9-add7-85be5e7afe55"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusPartiallyPaid,
            SubTotal: codataccounting.Float64(4720.94),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "mollitia": map[string]interface{}{
                        "explicabo": "laudantium",
                        "sunt": "repellat",
                    },
                    "aliquam": map[string]interface{}{
                        "eveniet": "dolorem",
                        "laborum": "eos",
                    },
                    "dolor": map[string]interface{}{
                        "unde": "eius",
                    },
                },
            },
            TotalAmount: 6602.4,
            TotalDiscount: codataccounting.Float64(4293.32),
            TotalTaxAmount: 5230.55,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 7593.83,
                    Name: "Paul Stark DVM",
                },
                shared.WithholdingTaxitems{
                    Amount: 9401.37,
                    Name: "Jacqueline Blick",
                },
                shared.WithholdingTaxitems{
                    Amount: 8368.03,
                    Name: "Dr. Rosie McClure",
                },
                shared.WithholdingTaxitems{
                    Amount: 5901,
                    Name: "Jackie Johnston",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(237032),
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
        InvoiceID: "tempore",
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
        InvoiceID: "praesentium",
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
        InvoiceID: "necessitatibus",
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
        InvoiceID: "perferendis",
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
        Query: codataccounting.String("consequatur"),
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
        InvoiceID: "doloremque",
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
            AdditionalTaxAmount: codataccounting.Float64(5696.51),
            AdditionalTaxPercentage: codataccounting.Float64(7989.88),
            AmountDue: 4285.11,
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(6211.79),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("explicabo"),
                ID: "de7b3562-201a-46aa-b4ae-7b1a5b908d4e",
            },
            DiscountPercentage: codataccounting.Float64(2357.58),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("491aa35d-4a83-49f0-bbab-77b918f03139"),
            InvoiceNumber: codataccounting.String("quos"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("07e0e39c-7e23-4ecb-8604-652e23a3d6c6"),
                        Name: codataccounting.String("Delores Swift"),
                    },
                    Description: codataccounting.String("debitis"),
                    DiscountAmount: codataccounting.Float64(5242.12),
                    DiscountPercentage: codataccounting.Float64(9387.2),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "7f002d19-86aa-499d-ba1d-32329e45837e",
                        Name: codataccounting.String("Dominick Denesik"),
                    },
                    Quantity: 3883.38,
                    SubTotal: codataccounting.Float64(7358.71),
                    TaxAmount: codataccounting.Float64(6883.86),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(771.66),
                        ID: codataccounting.String("0e255fdc-480d-46e3-b086-75cbf186856a"),
                        Name: codataccounting.String("Gretchen Lesch"),
                    },
                    TotalAmount: codataccounting.Float64(8162.54),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9d0fc282-c666-4af3-83f5-589bea5d264e",
                                Name: codataccounting.String("Doris Tromp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a84822e5-13f6-4d9d-aad3-7c3099077c10",
                                Name: codataccounting.String("Lester Leuschke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2163e67d-4886-4054-bc0a-3049c3cf6c02",
                                Name: codataccounting.String("Katie Tromp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "21bad90d-2743-4fd6-82a1-0e6c2978ec25",
                                Name: codataccounting.String("Brandi Harris V"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consequuntur"),
                            ID: "27fcc479-96c9-477b-bc57-f38928a8600c",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d67d63e4-aa56-4846-8579-cfc6c0e503f5",
                            Name: codataccounting.String("Dr. Maxine Dickinson PhD"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ed87b28e-8afa-4bc9-86e2-41e43b234241",
                            Name: codataccounting.String("Angel Brown"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3f62aa9a-e4ae-48ab-8a9c-492c5e8ba5d4",
                            Name: codataccounting.String("Oliver Gusikowski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "08bd380c-29aa-48dd-b1bd-daa30b7b9144",
                            Name: codataccounting.String("Luke Turner"),
                        },
                    },
                    UnitAmount: 7652.45,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("088d418b-b718-404f-823d-543935f377ac"),
                        Name: codataccounting.String("Leticia Marks"),
                    },
                    Description: codataccounting.String("officiis"),
                    DiscountAmount: codataccounting.Float64(6126.88),
                    DiscountPercentage: codataccounting.Float64(2441.7),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b6a3c523-105e-47c3-8cab-0ecb812a6614",
                        Name: codataccounting.String("Austin Goyette"),
                    },
                    Quantity: 5308.56,
                    SubTotal: codataccounting.Float64(9273.16),
                    TaxAmount: codataccounting.Float64(5759.77),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(419.17),
                        ID: codataccounting.String("85075bc2-5382-4533-83fb-0a4e66ea4757"),
                        Name: codataccounting.String("Laurence Bergstrom DVM"),
                    },
                    TotalAmount: codataccounting.Float64(1857.6),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "41818fc6-79b6-4b2f-a535-9b855d015b62",
                                Name: codataccounting.String("Julian Rippin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a38a8a88-c144-4200-82ca-eb1ae1ecf8c3",
                                Name: codataccounting.String("Shelly Gleason"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ba7a05a8-b4a9-4ec5-b368-8cca36327276",
                                Name: codataccounting.String("Jana McCullough"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("earum"),
                            ID: "97e05410-3347-4d78-bf24-91145fab9e59",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "af336664-eaa6-4bf2-bf14-e8c1b352acce",
                            Name: codataccounting.String("Lorenzo Schoen"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "27814eca-016b-4c41-aa13-42d4104a25ef",
                            Name: codataccounting.String("Joyce Steuber"),
                        },
                    },
                    UnitAmount: 4979.58,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("dicta"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2578.31),
                        TotalAmount: codataccounting.Float64(6458.59),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4317692e-a486-473d-922b-828a9030660f"),
                            Name: codataccounting.String("Denise Hagenes"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(6929.74),
                        ID: codataccounting.String("4cc64c2b-3a32-4c48-8ade-62f6aa558a65"),
                        Note: codataccounting.String("eveniet"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("aperiam"),
                        TotalAmount: codataccounting.Float64(5340.39),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(965.66),
                        TotalAmount: codataccounting.Float64(3850.79),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ca34bb87-d4f6-4212-ba60-7d1606294514"),
                            Name: codataccounting.String("Tony Stracke"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6279.31),
                        ID: codataccounting.String("9f38bd2b-e878-4703-893f-49aa8465a328"),
                        Note: codataccounting.String("adipisci"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("iusto"),
                        TotalAmount: codataccounting.Float64(6184.81),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1091.5),
                        TotalAmount: codataccounting.Float64(6243.12),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d1cea673-d86e-43b3-9e49-a3135778ce54"),
                            Name: codataccounting.String("Alfonso Sawayn PhD"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(8753.05),
                        ID: codataccounting.String("a975045b-acf6-43b2-9518-6ab5e3a02261"),
                        Note: codataccounting.String("tempora"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("ab"),
                        TotalAmount: codataccounting.Float64(3225.74),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(3435.3),
                        TotalAmount: codataccounting.Float64(4058.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8299e61a-fc71-486f-b20b-7a73df40ca0d"),
                            Name: codataccounting.String("Vicki Hammes"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4029.47),
                        ID: codataccounting.String("41bbf055-271b-4251-9dd6-06dd1b28272b"),
                        Note: codataccounting.String("minus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("placeat"),
                        TotalAmount: codataccounting.Float64(1884.9),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("sunt"),
                    ID: codataccounting.String("1697b188-0fcb-4b2b-93c1-5f670bd17848"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.InvoiceStatusUnknown,
            SubTotal: codataccounting.Float64(4271.84),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolor": map[string]interface{}{
                        "necessitatibus": "harum",
                        "amet": "tempore",
                        "ex": "vero",
                        "odit": "quaerat",
                    },
                    "dicta": map[string]interface{}{
                        "adipisci": "quasi",
                        "alias": "occaecati",
                        "perspiciatis": "deleniti",
                        "dolor": "eum",
                    },
                },
            },
            TotalAmount: 4314.81,
            TotalDiscount: codataccounting.Float64(2091.69),
            TotalTaxAmount: 8086.29,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4157.82,
                    Name: "Noel Renner",
                },
                shared.WithholdingTaxitems{
                    Amount: 8616.31,
                    Name: "Gilbert Schaden V",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "nobis",
        TimeoutInMinutes: codataccounting.Int(520852),
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
            Content: []byte("distinctio"),
            RequestBody: "modi",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "aperiam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
