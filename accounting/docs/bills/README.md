# Bills

## Overview

Bills

### Available Operations

* [Create](#create) - Create bill
* [Delete](#delete) - Delete bill
* [DownloadAttachment](#downloadattachment) - Download bill attachment
* [Get](#get) - Get bill
* [GetAttachment](#getattachment) - Get bill attachment
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bill model
* [List](#list) - List bills
* [ListAttachments](#listattachments) - List bill attachments
* [Update](#update) - Update bill
* [UploadAttachment](#uploadattachment) - Upload bill attachment

## Create

Posts a new bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) to see which integrations support this endpoint.


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
    res, err := s.Bills.Create(ctx, operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(4486.01),
            Currency: codataccounting.String("culpa"),
            CurrencyRate: codataccounting.Float64(424.89),
            DueDate: codataccounting.String("fuga"),
            ID: codataccounting.String("116ce723-d409-47fa-b0e9-af725b291220"),
            IssueDate: "amet",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d83f5aeb-7799-4d22-a8c1-f8493825fdc4"),
                        Name: codataccounting.String("Jacquelyn Lueilwitz"),
                    },
                    Description: codataccounting.String("maxime"),
                    DiscountAmount: codataccounting.Float64(1514.86),
                    DiscountPercentage: codataccounting.Float64(7915.38),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "2dfb4cfc-1c76-4230-b841-fb1bd23fdb14",
                        Name: codataccounting.String("Malcolm Johnson"),
                    },
                    Quantity: 3422.26,
                    SubTotal: codataccounting.Float64(6422.79),
                    TaxAmount: codataccounting.Float64(3755.88),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5163.63),
                        ID: codataccounting.String("5998e22a-e20d-4a16-bc2b-271a289c57e8"),
                        Name: codataccounting.String("Ellen Walter II"),
                    },
                    TotalAmount: codataccounting.Float64(2199.4),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d2224656-9462-4407-884f-7ab37cef0222",
                                Name: codataccounting.String("Jean Mayert"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b55410ad-c669-4af9-8a26-c7cdc981f068",
                                Name: codataccounting.String("Johnnie Berge"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bb33cfaa-348c-431b-b407-ee4fcf0c42b7",
                                Name: codataccounting.String("Abel Buckridge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consequuntur"),
                            ID: "6398a0dc-7663-424c-8b06-c8ca12d02529",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0b8d5722-dd89-45b8-bcf2-4db959693352",
                            Name: codataccounting.String("Clinton Greenholt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "994d78de-3b6e-4938-9f5a-bb7f662550a2",
                            Name: codataccounting.String("Travis Leannon"),
                        },
                    },
                    UnitAmount: 7547.84,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("incidunt"),
            Note: codataccounting.String("deleniti"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("est"),
                        Currency: codataccounting.String("reiciendis"),
                        CurrencyRate: codataccounting.Float64(8204.62),
                        TotalAmount: codataccounting.Float64(1378.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("315bba65-0164-4e06-b5bf-6ae591bc8bde"),
                            Name: codataccounting.String("Mr. Dale Jenkins"),
                        },
                        Currency: codataccounting.String("ex"),
                        CurrencyRate: codataccounting.Float64(2033.96),
                        ID: codataccounting.String("c205fda8-4077-44a6-8a9a-35d086b6f66f"),
                        Note: codataccounting.String("accusamus"),
                        PaidOnDate: codataccounting.String("reiciendis"),
                        Reference: codataccounting.String("consequatur"),
                        TotalAmount: codataccounting.Float64(1487.14),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("e9f443b4-257b-4992-88db-da6a61efa219"),
                    PurchaseOrderNumber: codataccounting.String("atque"),
                },
            },
            Reference: codataccounting.String("explicabo"),
            SourceModifiedDate: codataccounting.String("ullam"),
            Status: shared.BillStatusPaid,
            SubTotal: 9874.35,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "aut": map[string]interface{}{
                        "iste": "eveniet",
                        "nam": "animi",
                        "labore": "voluptate",
                    },
                    "voluptatibus": map[string]interface{}{
                        "nulla": "dolorem",
                        "voluptates": "a",
                    },
                    "perferendis": map[string]interface{}{
                        "excepturi": "aliquid",
                        "dolore": "voluptatem",
                    },
                    "illum": map[string]interface{}{
                        "culpa": "dicta",
                        "atque": "ratione",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "1c87adf5-96fd-4f1a-9837-ae80c1c19c95",
                SupplierName: codataccounting.String("tempore"),
            },
            TaxAmount: 6379.69,
            TotalAmount: 6105.63,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 5589.92,
                    Name: "Dora Lemke",
                },
                shared.BillWithholdingTax{
                    Amount: 2212.4,
                    Name: "Brad Mayer",
                },
                shared.BillWithholdingTax{
                    Amount: 5865.56,
                    Name: "Kristine Wilderman",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(561858),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillResponse != nil {
        // handle response
    }
}
```

## Delete

﻿The _Delete Bills_ endpoint allows you to delete a specified Bill from an accounting platform. 

### Process 
1. Pass the `{billId}` to the _Delete Bills_ endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete operation by checking the status of push operation either via
    1. [Push operation webhook](/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
    2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation).

   A `Success` status indicates that the Bill object was deleted from the accounting platform.
3. (Optional) Check that the Bill was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting a Bill from an accounting platform might cause related objects to be modified. For example, if you delete a paid Bill in QuickBooks Online, the bill is deleted but the bill payment against that bill is not. The bill payment is converted to a payment on account.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Delete | Details                                                                                                      |  
|-------------|-------------|--------------------------------------------------------------------------------------------------------------|
| QuickBooks Online | No          | -                                                                                                            |
| Oracle NetSuite   | No          | When deleting a Bill that's already linked to a Bill payment, you must delete the linked Bill payment first. |

> **Supported Integrations**
> 
> This functionality is currently only supported for our QuickBooks Online abd Oracle NetSuite integrations. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    res, err := s.Bills.Delete(ctx, operations.DeleteBillRequest{
        BillID: "quod",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

﻿Download bill attachment.

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
    res, err := s.Bills.DownloadAttachment(ctx, operations.DownloadBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "repudiandae",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

﻿Get a bill.

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
    res, err := s.Bills.Get(ctx, operations.GetBillRequest{
        BillID: "eaque",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bill != nil {
        // handle response
    }
}
```

## GetAttachment

﻿Get bill attachment.

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
    res, err := s.Bills.GetAttachment(ctx, operations.GetBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "consectetur",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

﻿Get create/update bill model.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating and updating a bill.

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
    res, err := s.Bills.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBillsModelRequest{
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

﻿Gets the latest bills for a company, with pagination.

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
    res, err := s.Bills.List(ctx, operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("autem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bills != nil {
        // handle response
    }
}
```

## ListAttachments

﻿List bill attachments.

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
    res, err := s.Bills.ListAttachments(ctx, operations.ListBillAttachmentsRequest{
        BillID: "vitae",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

﻿Posts an updated bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support updating a bill.

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
    res, err := s.Bills.Update(ctx, operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(2569.75),
            Currency: codataccounting.String("incidunt"),
            CurrencyRate: codataccounting.Float64(2669.76),
            DueDate: codataccounting.String("quos"),
            ID: codataccounting.String("c7977a0e-f2f5-4360-a8ef-eef934152ed7"),
            IssueDate: "itaque",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("53f4c157-deaa-4717-8f44-5accf667aaf9"),
                        Name: codataccounting.String("Garry Nicolas IV"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(9594.2),
                    DiscountPercentage: codataccounting.Float64(9133.93),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "431d6bf5-c838-4fbb-8c20-cb67fc4b425e",
                        Name: codataccounting.String("Sergio Tremblay"),
                    },
                    Quantity: 2277.06,
                    SubTotal: codataccounting.Float64(2796.79),
                    TaxAmount: codataccounting.Float64(7835.39),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5816.8),
                        ID: codataccounting.String("f7b79dfe-b77a-45c3-8d4b-af91e506ef89"),
                        Name: codataccounting.String("Lee Hegmann"),
                    },
                    TotalAmount: codataccounting.Float64(2574.88),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5f16f56d-385a-43c4-ac63-1b99e26ced8f",
                                Name: codataccounting.String("Darrin Skiles"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "410f63bb-f817-4837-b01a-fdd788624189",
                                Name: codataccounting.String("Rudy Gorczany"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ducimus"),
                            ID: "3f5033f1-9dbf-4125-8e41-52eab9cd7e52",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "a6a0e123-b784-47ec-99e1-f67f3c4cce4b",
                            Name: codataccounting.String("Angel Kris"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ff3c5747-5013-457e-84f5-1f8b084c3197",
                            Name: codataccounting.String("Samuel Mann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "245467f9-4874-4c2d-9cc4-972233e66bd8",
                            Name: codataccounting.String("Frankie Herzog Jr."),
                        },
                    },
                    UnitAmount: 7277.89,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("excepturi"),
            Note: codataccounting.String("odio"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("officiis"),
                        Currency: codataccounting.String("delectus"),
                        CurrencyRate: codataccounting.Float64(1692.29),
                        TotalAmount: codataccounting.Float64(232.36),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("38732059-0ccc-4109-a400-313b3e5044f6"),
                            Name: codataccounting.String("Shawna Trantow"),
                        },
                        Currency: codataccounting.String("assumenda"),
                        CurrencyRate: codataccounting.Float64(7637.47),
                        ID: codataccounting.String("4077d0cc-3f40-48ef-815c-eb4d6e1eae0f"),
                        Note: codataccounting.String("odio"),
                        PaidOnDate: codataccounting.String("veniam"),
                        Reference: codataccounting.String("fuga"),
                        TotalAmount: codataccounting.Float64(9290.12),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("possimus"),
                        Currency: codataccounting.String("tenetur"),
                        CurrencyRate: codataccounting.Float64(1499.41),
                        TotalAmount: codataccounting.Float64(6481.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("cab58b99-1c92-46dd-b589-461e7421cbe6"),
                            Name: codataccounting.String("Mr. Rene Harris"),
                        },
                        Currency: codataccounting.String("consequatur"),
                        CurrencyRate: codataccounting.Float64(8776.19),
                        ID: codataccounting.String("a930b69f-7ac2-4f72-b885-009049116082"),
                        Note: codataccounting.String("perferendis"),
                        PaidOnDate: codataccounting.String("esse"),
                        Reference: codataccounting.String("quas"),
                        TotalAmount: codataccounting.Float64(5007.68),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("laudantium"),
                        Currency: codataccounting.String("voluptates"),
                        CurrencyRate: codataccounting.Float64(7939.09),
                        TotalAmount: codataccounting.Float64(4196.69),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6183bfe9-659e-4b40-ac16-faf75b0b532a"),
                            Name: codataccounting.String("Antonia Muller"),
                        },
                        Currency: codataccounting.String("eligendi"),
                        CurrencyRate: codataccounting.Float64(7363.13),
                        ID: codataccounting.String("aaf4452c-4842-4c9b-aad3-2dafe81a88f4"),
                        Note: codataccounting.String("incidunt"),
                        PaidOnDate: codataccounting.String("labore"),
                        Reference: codataccounting.String("ut"),
                        TotalAmount: codataccounting.Float64(3134.2),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("3fecd473-53f6-43c8-a093-79aa69cd5fbc"),
                    PurchaseOrderNumber: codataccounting.String("sapiente"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("79da18a7-822b-4f95-894e-6861adb55f9e"),
                    PurchaseOrderNumber: codataccounting.String("ipsam"),
                },
            },
            Reference: codataccounting.String("fugiat"),
            SourceModifiedDate: codataccounting.String("odio"),
            Status: shared.BillStatusPartiallyPaid,
            SubTotal: 819.17,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "occaecati": map[string]interface{}{
                        "necessitatibus": "rem",
                        "a": "nihil",
                        "veniam": "aut",
                        "magni": "rerum",
                    },
                    "voluptatibus": map[string]interface{}{
                        "quod": "non",
                        "dolore": "enim",
                        "alias": "blanditiis",
                        "modi": "illo",
                    },
                    "a": map[string]interface{}{
                        "molestiae": "autem",
                    },
                    "dolore": map[string]interface{}{
                        "nostrum": "ex",
                        "amet": "voluptate",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "9f3fb27e-21f8-4626-97b3-6fc6b9f587ce",
                SupplierName: codataccounting.String("ad"),
            },
            TaxAmount: 1676.13,
            TotalAmount: 3457.46,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 4005.1,
                    Name: "Miss Terri Gerhold",
                },
                shared.BillWithholdingTax{
                    Amount: 2443.32,
                    Name: "Jane Towne II",
                },
                shared.BillWithholdingTax{
                    Amount: 4496.94,
                    Name: "Mario Runolfsson DDS",
                },
                shared.BillWithholdingTax{
                    Amount: 7678.8,
                    Name: "Jesus Corkery",
                },
            },
        },
        BillID: "facilis",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(792555),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillResponse != nil {
        // handle response
    }
}
```

## UploadAttachment

﻿Upload bill attachment.

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
    res, err := s.Bills.UploadAttachment(ctx, operations.UploadBillAttachmentRequest{
        RequestBody: &operations.UploadBillAttachmentRequestBody{
            Content: []byte("vero"),
            RequestBody: "impedit",
        },
        BillID: "omnis",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
