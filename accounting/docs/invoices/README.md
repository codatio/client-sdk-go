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
            AdditionalTaxAmount: codataccounting.Float64(50.37),
            AdditionalTaxPercentage: codataccounting.Float64(162.33),
            AmountDue: 632.79,
            Currency: codataccounting.String("tempora"),
            CurrencyRate: codataccounting.Float64(9194.06),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("hic"),
                ID: "45cea11a-c53e-4bb6-987f-340414c5b9ac",
            },
            DiscountPercentage: codataccounting.Float64(9131.37),
            DueDate: codataccounting.String("debitis"),
            ID: codataccounting.String("400ae9f9-2caf-41b0-a5f1-d14718c6fa2f"),
            InvoiceNumber: codataccounting.String("mollitia"),
            IssueDate: "quibusdam",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c06c5d95-472c-4dd1-8fc4-3b70bca88fa7"),
                        Name: codataccounting.String("Roxanne Green"),
                    },
                    Description: codataccounting.String("exercitationem"),
                    DiscountAmount: codataccounting.Float64(817.93),
                    DiscountPercentage: codataccounting.Float64(7886.47),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3dd1eb8f-7f75-4f4f-a3f1-c0a586c3ae7d",
                        Name: codataccounting.String("Patty Hoeger"),
                    },
                    Quantity: 8864.96,
                    SubTotal: codataccounting.Float64(9148.03),
                    TaxAmount: codataccounting.Float64(9451.88),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3338.68),
                        ID: codataccounting.String("e142d95b-1dbe-4cef-b7c4-b156e9278275"),
                        Name: codataccounting.String("Grady Nitzsche"),
                    },
                    TotalAmount: codataccounting.Float64(5122.23),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7468063f-799b-4795-ac0b-0fa0bb20a40e",
                                Name: codataccounting.String("Krista Goyette"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("aliquid"),
                            ID: "40642726-57b0-41a0-bc08-fd3921c25793",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6f093a3e-fa46-4d36-adfa-1011a091b3ec",
                            Name: codataccounting.String("Garry Heller"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2de1a9d1-4fe7-42e5-a1f9-0303dfc33839",
                            Name: codataccounting.String("Lela Welch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6d1d3209-0fc1-457a-89fe-1961ce9be41c",
                            Name: codataccounting.String("Charlie Monahan"),
                        },
                    },
                    UnitAmount: 4955.97,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("quibusdam"),
            Note: codataccounting.String("omnis"),
            PaidOnDate: codataccounting.String("molestiae"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("unde"),
                        Currency: codataccounting.String("repellendus"),
                        CurrencyRate: codataccounting.Float64(78.52),
                        TotalAmount: codataccounting.Float64(4916.69),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b200a58f-fd29-467d-b8fd-882a8e60be62"),
                            Name: codataccounting.String("Lynne Schultz"),
                        },
                        Currency: codataccounting.String("veniam"),
                        CurrencyRate: codataccounting.Float64(6380.64),
                        ID: codataccounting.String("fdd04c37-5251-42be-ae1d-87ecc5fdcea8"),
                        Note: codataccounting.String("eveniet"),
                        PaidOnDate: codataccounting.String("molestiae"),
                        Reference: codataccounting.String("laborum"),
                        TotalAmount: codataccounting.Float64(5326.99),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("adipisci"),
                    ID: codataccounting.String("11662cda-6d77-4c1d-8606-6237d4227866"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("quibusdam"),
                    ID: codataccounting.String("b8a749e3-9845-411c-875e-4f0c004b5bb7"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("ipsam"),
                    ID: codataccounting.String("8cc94562-f006-4968-9fcd-1a173d84bbe2"),
                },
            },
            SourceModifiedDate: codataccounting.String("modi"),
            Status: shared.InvoiceStatusVoid,
            SubTotal: codataccounting.Float64(1277.99),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "praesentium": map[string]interface{}{
                        "magnam": "mollitia",
                    },
                    "doloribus": map[string]interface{}{
                        "doloremque": "odio",
                        "ratione": "corporis",
                        "eligendi": "expedita",
                    },
                    "laboriosam": map[string]interface{}{
                        "molestias": "corporis",
                    },
                },
            },
            TotalAmount: 8154.52,
            TotalDiscount: codataccounting.Float64(2900.44),
            TotalTaxAmount: 6668.63,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 6035.51,
                    Name: "Dr. Wilbur Orn III",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(607549),
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
        InvoiceID: "veritatis",
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
        InvoiceID: "nemo",
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
        InvoiceID: "nisi",
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
        InvoiceID: "repellat",
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
        InvoiceID: "voluptate",
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
        Query: codataccounting.String("possimus"),
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
        InvoiceID: "aspernatur",
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
            AdditionalTaxAmount: codataccounting.Float64(9264.79),
            AdditionalTaxPercentage: codataccounting.Float64(9197.03),
            AmountDue: 1804.63,
            Currency: codataccounting.String("perferendis"),
            CurrencyRate: codataccounting.Float64(6247.3),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("ullam"),
                ID: "05bf03a9-3e94-4480-8a37-fb10789032ac",
            },
            DiscountPercentage: codataccounting.Float64(2261.31),
            DueDate: codataccounting.String("amet"),
            ID: codataccounting.String("3172e2dd-79ec-474b-a7e8-8ddb36fd1ccc"),
            InvoiceNumber: codataccounting.String("nesciunt"),
            IssueDate: "labore",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c8657347-4f0a-4740-bb4a-b441c3a09e76"),
                        Name: codataccounting.String("Katrina Monahan"),
                    },
                    Description: codataccounting.String("rem"),
                    DiscountAmount: codataccounting.Float64(313.05),
                    DiscountPercentage: codataccounting.Float64(5480.08),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "bbe79445-5ebc-4550-a1c4-26b59c8366fd",
                        Name: codataccounting.String("Kim Braun"),
                    },
                    Quantity: 3640.73,
                    SubTotal: codataccounting.Float64(5323.1),
                    TaxAmount: codataccounting.Float64(1563.13),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7881.7),
                        ID: codataccounting.String("1b855e88-9d9e-4f93-ae90-00a13ad81242"),
                        Name: codataccounting.String("Jennie Veum"),
                    },
                    TotalAmount: codataccounting.Float64(1827.3),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "411898e7-3879-4efb-a8ba-ebabb794536e",
                                Name: codataccounting.String("Jeffrey Frami MD"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quidem"),
                            ID: "97631720-b77a-45a5-b65a-79f15271f01c",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "361fed8d-c5ef-4fb4-93e9-089e871fdb4d",
                            Name: codataccounting.String("Katrina Krajcik"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9c985e43-734a-45d7-ad9e-dd785be5e7af",
                            Name: codataccounting.String("Derrick Hane"),
                        },
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
                    UnitAmount: 98.84,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ipsum"),
            Note: codataccounting.String("quidem"),
            PaidOnDate: codataccounting.String("dolorum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("molestiae"),
                        Currency: codataccounting.String("reprehenderit"),
                        CurrencyRate: codataccounting.Float64(7232.85),
                        TotalAmount: codataccounting.Float64(5807.71),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("18f03139-8450-47e0-a39c-7e23ecb06046"),
                            Name: codataccounting.String("Denise Walter"),
                        },
                        Currency: codataccounting.String("similique"),
                        CurrencyRate: codataccounting.Float64(2248.7),
                        ID: codataccounting.String("d6c657e9-de8f-47f0-82d1-986aa99d3a1d"),
                        Note: codataccounting.String("ratione"),
                        PaidOnDate: codataccounting.String("odit"),
                        Reference: codataccounting.String("amet"),
                        TotalAmount: codataccounting.Float64(1716.48),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("provident"),
                        Currency: codataccounting.String("repudiandae"),
                        CurrencyRate: codataccounting.Float64(2826.23),
                        TotalAmount: codataccounting.Float64(3547.05),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("837e8f2a-d6bb-410e-a55f-dc480d6e3308"),
                            Name: codataccounting.String("Stella Herman"),
                        },
                        Currency: codataccounting.String("sapiente"),
                        CurrencyRate: codataccounting.Float64(632.04),
                        ID: codataccounting.String("86856a7e-82cd-4f9d-8fc2-82c666af3c3f"),
                        Note: codataccounting.String("quis"),
                        PaidOnDate: codataccounting.String("nostrum"),
                        Reference: codataccounting.String("totam"),
                        TotalAmount: codataccounting.Float64(5904.76),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("distinctio"),
                        Currency: codataccounting.String("accusamus"),
                        CurrencyRate: codataccounting.Float64(6668.05),
                        TotalAmount: codataccounting.Float64(3319.44),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d264e41e-2ca8-4482-ae51-3f6d9d2ad37c"),
                            Name: codataccounting.String("Shirley Mueller III"),
                        },
                        Currency: codataccounting.String("esse"),
                        CurrencyRate: codataccounting.Float64(7504.07),
                        ID: codataccounting.String("10b68792-163e-467d-8886-0543c0a3049c"),
                        Note: codataccounting.String("ipsum"),
                        PaidOnDate: codataccounting.String("quod"),
                        Reference: codataccounting.String("voluptatibus"),
                        TotalAmount: codataccounting.Float64(3769.3),
                    },
                },
            },
            SalesOrderRefs: []shared.SalesOrderRef{
                shared.SalesOrderRef{
                    DataType: codataccounting.String("sit"),
                    ID: codataccounting.String("276e7b21-bad9-40d2-b43f-d6c2a10e6c29"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("quam"),
                    ID: codataccounting.String("8ec256a5-b092-427f-8c47-996c977bbc57"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("sapiente"),
                    ID: codataccounting.String("38928a86-00c5-48d6-bd63-e4aa56846457"),
                },
                shared.SalesOrderRef{
                    DataType: codataccounting.String("omnis"),
                    ID: codataccounting.String("cfc6c0e5-03f5-4683-9f1d-8ed87b28e8af"),
                },
            },
            SourceModifiedDate: codataccounting.String("culpa"),
            Status: shared.InvoiceStatusPaid,
            SubTotal: codataccounting.Float64(7973.76),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "laudantium": map[string]interface{}{
                        "necessitatibus": "consequuntur",
                        "aliquam": "dicta",
                    },
                    "earum": map[string]interface{}{
                        "dolorem": "quidem",
                        "consequuntur": "ratione",
                    },
                    "ut": map[string]interface{}{
                        "dolore": "inventore",
                    },
                },
            },
            TotalAmount: 4725.74,
            TotalDiscount: codataccounting.Float64(8462.5),
            TotalTaxAmount: 1103.98,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 8934.16,
                    Name: "Miss Jeannie Hudson",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "cupiditate",
        TimeoutInMinutes: codataccounting.Int(633887),
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
            Content: []byte("recusandae"),
            RequestBody: "labore",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        InvoiceID: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
