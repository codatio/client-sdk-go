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
            AmountDue: codataccounting.Float64(7674.66),
            Currency: codataccounting.String("doloremque"),
            CurrencyRate: codataccounting.Float64(5168.33),
            DueDate: codataccounting.String("iure"),
            ID: codataccounting.String("fa21e915-2cb3-4119-967b-8e3c8db03408"),
            IssueDate: "repellendus",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d364ffd4-5590-46d1-a63d-48e935c2c9e8"),
                        Name: codataccounting.String("Miss Jeannie Emmerich"),
                    },
                    Description: codataccounting.String("sequi"),
                    DiscountAmount: codataccounting.Float64(9258.47),
                    DiscountPercentage: codataccounting.Float64(2863.29),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3202d721-6576-4506-a418-70d9d21f9ad0",
                        Name: codataccounting.String("Sharon Ruecker"),
                    },
                    Quantity: 7639.37,
                    SubTotal: codataccounting.Float64(8061.24),
                    TaxAmount: codataccounting.Float64(922.64),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1073.01),
                        ID: codataccounting.String("a0836429-068b-4850-aa55-e7f73bc845e3"),
                        Name: codataccounting.String("Helen O'Reilly V"),
                    },
                    TotalAmount: codataccounting.Float64(9738.94),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "badf947c-9a86-47bc-8242-6665816ddca8",
                                Name: codataccounting.String("Dr. Emilio Hilll"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b4c593ec-12cd-4aad-8ec7-afedbd80df44",
                                Name: codataccounting.String("Otis Greenholt"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("iste"),
                            ID: "390c5888-0983-4dab-b9ef-3ffdd9f7f079",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "4d35724c-db0f-44d2-8118-7d56844eded8",
                            Name: codataccounting.String("Ms. Madeline Miller"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "628bdfc2-032b-46c8-b992-3b7e13584f7a",
                            Name: codataccounting.String("Carl Davis"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "891f82ce-1157-4172-b053-77dcfa89df97",
                            Name: codataccounting.String("Tasha Dickinson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "686092e9-c3dd-4c5f-911d-ea1026d541a4",
                            Name: codataccounting.String("Dr. Terry Mohr"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b21780bc-cc0d-4bbd-9b48-4708fb4e391e",
                            Name: codataccounting.String("Mrs. Susie Schowalter"),
                        },
                    },
                    UnitAmount: 7803.7,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4c4e5459-9ea3-4422-a0e9-b200ce78a1bd"),
                        Name: codataccounting.String("Cary Predovic"),
                    },
                    Description: codataccounting.String("doloremque"),
                    DiscountAmount: codataccounting.Float64(6813.36),
                    DiscountPercentage: codataccounting.Float64(1175.46),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "16ce723d-4097-4fa3-8e9a-f725b2912203",
                        Name: codataccounting.String("Hope Lemke"),
                    },
                    Quantity: 3564.85,
                    SubTotal: codataccounting.Float64(6442.99),
                    TaxAmount: codataccounting.Float64(9319.53),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7143),
                        ID: codataccounting.String("7799d22e-8c1f-4849-b825-fdc42c876c2c"),
                        Name: codataccounting.String("Marcella Windler"),
                    },
                    TotalAmount: codataccounting.Float64(7579.62),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c1c76230-f841-4fb1-bd23-fdb14db6be5a",
                                Name: codataccounting.String("Lena Herzog"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8e22ae20-da16-4fc2-b271-a289c57e854e",
                                Name: codataccounting.String("Jeffrey Gutmann"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "d2224656-9462-4407-884f-7ab37cef0222",
                                Name: codataccounting.String("Jean Mayert"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b55410ad-c669-4af9-8a26-c7cdc981f068",
                                Name: codataccounting.String("Johnnie Berge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("libero"),
                            ID: "b33cfaa3-48c3-41bf-807e-e4fcf0c42b78",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "5626398a-0dc7-4663-a4cc-b06c8ca12d02",
                            Name: codataccounting.String("Judy Mertz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b8d5722d-d895-4b8b-8f24-db959693352f",
                            Name: codataccounting.String("Joanne Hermiston"),
                        },
                    },
                    UnitAmount: 5831.38,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("occaecati"),
            Note: codataccounting.String("numquam"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("molestiae"),
                        Currency: codataccounting.String("quas"),
                        CurrencyRate: codataccounting.Float64(8341.77),
                        TotalAmount: codataccounting.Float64(9065.24),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("3b6e9389-f5ab-4b7f-a625-50a28382ac48"),
                            Name: codataccounting.String("Monique Wisoky"),
                        },
                        Currency: codataccounting.String("consectetur"),
                        CurrencyRate: codataccounting.Float64(809.98),
                        ID: codataccounting.String("5bba6501-64e0-46f5-bf6a-e591bc8bdef3"),
                        Note: codataccounting.String("commodi"),
                        PaidOnDate: codataccounting.String("vitae"),
                        Reference: codataccounting.String("fugit"),
                        TotalAmount: codataccounting.Float64(7240.73),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("ex"),
                        Currency: codataccounting.String("neque"),
                        CurrencyRate: codataccounting.Float64(7977.12),
                        TotalAmount: codataccounting.Float64(1761.04),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("05fda840-774a-468a-9a35-d086b6f66fef"),
                            Name: codataccounting.String("Tammy Bartoletti"),
                        },
                        Currency: codataccounting.String("maiores"),
                        CurrencyRate: codataccounting.Float64(3114.49),
                        ID: codataccounting.String("43b4257b-992c-48db-9a6a-61efa2198258"),
                        Note: codataccounting.String("doloribus"),
                        PaidOnDate: codataccounting.String("pariatur"),
                        Reference: codataccounting.String("aut"),
                        TotalAmount: codataccounting.Float64(6302.86),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("iste"),
                        Currency: codataccounting.String("eveniet"),
                        CurrencyRate: codataccounting.Float64(7236.23),
                        TotalAmount: codataccounting.Float64(6585.44),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("47f7d3ef-0496-440d-aa18-31c87adf596f"),
                            Name: codataccounting.String("Winston Bergstrom"),
                        },
                        Currency: codataccounting.String("praesentium"),
                        CurrencyRate: codataccounting.Float64(2053.9),
                        ID: codataccounting.String("7ae80c1c-19c9-45ba-9986-78fa3f696991"),
                        Note: codataccounting.String("fuga"),
                        PaidOnDate: codataccounting.String("a"),
                        Reference: codataccounting.String("dolor"),
                        TotalAmount: codataccounting.Float64(5280.82),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("molestias"),
                        Currency: codataccounting.String("quod"),
                        CurrencyRate: codataccounting.Float64(9203.89),
                        TotalAmount: codataccounting.Float64(502.91),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("3614448c-7977-4a0e-b2f5-36028efeef93"),
                            Name: codataccounting.String("Kathleen Harris"),
                        },
                        Currency: codataccounting.String("possimus"),
                        CurrencyRate: codataccounting.Float64(4917.84),
                        ID: codataccounting.String("e253f4c1-57de-4aa7-970f-445accf667aa"),
                        Note: codataccounting.String("repellat"),
                        PaidOnDate: codataccounting.String("cupiditate"),
                        Reference: codataccounting.String("soluta"),
                        TotalAmount: codataccounting.Float64(7332.26),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("d185fe43-1d6b-4f5c-838f-bb8c20cb67fc"),
                    PurchaseOrderNumber: codataccounting.String("ut"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("b425e99e-6234-4c9f-bb79-dfeb77a5c38d"),
                    PurchaseOrderNumber: codataccounting.String("dolore"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("baf91e50-6ef8-490a-94b4-75f16f56d385"),
                    PurchaseOrderNumber: codataccounting.String("fuga"),
                },
            },
            Reference: codataccounting.String("sequi"),
            SourceModifiedDate: codataccounting.String("maxime"),
            Status: shared.BillStatusOpen,
            SubTotal: 6714.28,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "autem": map[string]interface{}{
                        "sunt": "rerum",
                    },
                    "occaecati": map[string]interface{}{
                        "necessitatibus": "fugit",
                        "autem": "optio",
                        "eveniet": "fugiat",
                    },
                    "blanditiis": map[string]interface{}{
                        "natus": "sapiente",
                        "repellendus": "facilis",
                        "molestias": "dolore",
                        "et": "accusantium",
                    },
                    "maiores": map[string]interface{}{
                        "velit": "tempore",
                        "expedita": "hic",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "817837b0-1afd-4d78-8624-189eb44873f5",
                SupplierName: codataccounting.String("accusantium"),
            },
            TaxAmount: 1902.6,
            TotalAmount: 2358.13,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 1030.53,
                    Name: "Rufus Reynolds Sr.",
                },
                shared.BillWithholdingTax{
                    Amount: 3661.17,
                    Name: "Mrs. Darrel Grant",
                },
                shared.BillWithholdingTax{
                    Amount: 9298.49,
                    Name: "Damon Mueller",
                },
                shared.BillWithholdingTax{
                    Amount: 4547.61,
                    Name: "Warren Conroy",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(625378),
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
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
        Query: codataccounting.String("vel"),
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
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AmountDue: codataccounting.Float64(6798.35),
            Currency: codataccounting.String("alias"),
            CurrencyRate: codataccounting.Float64(9303.98),
            DueDate: codataccounting.String("ab"),
            ID: codataccounting.String("23b7847e-c59e-41f6-bf3c-4cce4b6d7696"),
            IssueDate: "repellat",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3c574750-1357-4e44-b51f-8b084c3197e1"),
                        Name: codataccounting.String("Glenn Nolan"),
                    },
                    Description: codataccounting.String("corporis"),
                    DiscountAmount: codataccounting.Float64(3088.66),
                    DiscountPercentage: codataccounting.Float64(3816.39),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "7f94874c-2d5c-4c49-b223-3e66bd8fe5d0",
                        Name: codataccounting.String("Jeannette Mante"),
                    },
                    Quantity: 8874.22,
                    SubTotal: codataccounting.Float64(9615.76),
                    TaxAmount: codataccounting.Float64(1692.29),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(232.36),
                        ID: codataccounting.String("38732059-0ccc-4109-a400-313b3e5044f6"),
                        Name: codataccounting.String("Shawna Trantow"),
                    },
                    TotalAmount: codataccounting.Float64(8248.61),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "4077d0cc-3f40-48ef-815c-eb4d6e1eae0f",
                                Name: codataccounting.String("Sally Paucek"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f2acab58-b991-4c92-addb-589461e7421c",
                                Name: codataccounting.String("Rolando Jerde"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "502f0ea9-30b6-49f7-ac2f-72f885009049",
                                Name: codataccounting.String("Ms. Carolyn Jacobson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "07888ec6-6183-4bfe-9659-eb40ec16faf7",
                                Name: codataccounting.String("Olivia Auer"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ratione"),
                            ID: "2a4da37c-baaf-4445-ac48-42c9b2ad32da",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "81a88f44-4457-43fe-8d47-353f63c82093",
                            Name: codataccounting.String("Geneva Oberbrunner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cd5fbcf7-9da1-48a7-822b-f95894e6861a",
                            Name: codataccounting.String("Marco Hermann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9e5d751c-9fe8-4f75-82bf-dc3450841f17",
                            Name: codataccounting.String("Eleanor Gibson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "379f3fb2-7e21-4f86-a657-b36fc6b9f587",
                            Name: codataccounting.String("Bert Hand"),
                        },
                    },
                    UnitAmount: 7934.02,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("67641a83-12e5-4047-b4c2-1ccb423abcdc"),
                        Name: codataccounting.String("Carl Weimann"),
                    },
                    Description: codataccounting.String("distinctio"),
                    DiscountAmount: codataccounting.Float64(8546.46),
                    DiscountPercentage: codataccounting.Float64(8621.67),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "88e71f6c-4825-42d7-b71e-7fd074009ef8",
                        Name: codataccounting.String("Carlos Morissette"),
                    },
                    Quantity: 717.41,
                    SubTotal: codataccounting.Float64(8135.82),
                    TaxAmount: codataccounting.Float64(8536.25),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4986.39),
                        ID: codataccounting.String("097b5da0-8c57-4fa6-878a-216e19bafeca"),
                        Name: codataccounting.String("Mrs. Kathleen McDermott"),
                    },
                    TotalAmount: codataccounting.Float64(5410.46),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "40b64ff8-ae17-40ef-83b5-f37e4aa86855",
                                Name: codataccounting.String("Cora Jenkins"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nesciunt"),
                            ID: "2aa5dcb6-682c-4b70-b8cf-d5fb6e91b9a9",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "4846e2c3-309d-4b05-b6d9-e75ca006f539",
                            Name: codataccounting.String("Miss Krista Bosco"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a8bf92f9-7428-4ad9-a9f8-bf8221125359",
                            Name: codataccounting.String("Joey Labadie"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7f7a79cd-72cd-4248-8da2-1729f2ac41ef",
                            Name: codataccounting.String("Jackie Crist"),
                        },
                    },
                    UnitAmount: 1168.07,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("169ac1e4-1d8a-423c-a3e3-4f2dfa4a197f"),
                        Name: codataccounting.String("Betsy Walter"),
                    },
                    Description: codataccounting.String("aspernatur"),
                    DiscountAmount: codataccounting.Float64(911.36),
                    DiscountPercentage: codataccounting.Float64(3687.61),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1fe17120-9985-43e9-b543-d854439ee224",
                        Name: codataccounting.String("Kristen Bashirian"),
                    },
                    Quantity: 2107.1,
                    SubTotal: codataccounting.Float64(7122.09),
                    TaxAmount: codataccounting.Float64(7711.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1027.34),
                        ID: codataccounting.String("54188c2f-56e8-45da-b832-eabd617c3b0d"),
                        Name: codataccounting.String("Heather Nader"),
                    },
                    TotalAmount: codataccounting.Float64(6942.92),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "01bad870-6d46-4082-bfbd-c41ff5d4e2ae",
                                Name: codataccounting.String("Ora Purdy"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b35d1763-8f1e-4db7-8359-ecc5cb860f8c",
                                Name: codataccounting.String("Miss Dan Lakin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "73810e4f-e444-4729-bcd3-b1dd3bbce247",
                                Name: codataccounting.String("Karl Jacobs"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "eff50126-d71c-4ffb-90eb-74b8421953b4",
                                Name: codataccounting.String("Melody Stoltenberg"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("labore"),
                            ID: "3159d33e-5953-4c00-9139-863aa41e6c31",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "2f1fcb51-c9a4-41ff-be9c-bd795ee65e07",
                            Name: codataccounting.String("Johnnie Schamberger"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f616ea5c-7164-4193-8b90-f2e09d19d2fc",
                            Name: codataccounting.String("Faith Mosciski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e105944b-935d-4237-a72f-90849d6aed4a",
                            Name: codataccounting.String("Earnest Rogahn"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "37cd9222-c9ff-4574-91aa-bfa2e761f0ca",
                            Name: codataccounting.String("Meredith Green"),
                        },
                    },
                    UnitAmount: 8757.66,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f1031e68-99f0-4c20-81e2-2cd55cc0584a"),
                        Name: codataccounting.String("Lena Goyette"),
                    },
                    Description: codataccounting.String("voluptas"),
                    DiscountAmount: codataccounting.Float64(8175.09),
                    DiscountPercentage: codataccounting.Float64(6079.37),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "71fc820c-65b0-437b-b8e0-cc885187e4de",
                        Name: codataccounting.String("Leslie Pacocha"),
                    },
                    Quantity: 5461.33,
                    SubTotal: codataccounting.Float64(8037.63),
                    TaxAmount: codataccounting.Float64(3236.14),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8450.13),
                        ID: codataccounting.String("ddb46aa1-cfd6-4d82-8da0-131911296466"),
                        Name: codataccounting.String("Dr. Veronica Runte"),
                    },
                    TotalAmount: codataccounting.Float64(932.99),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "29042f56-9b7a-4ff0-aa22-16cbe071bc16",
                                Name: codataccounting.String("Rochelle Cormier"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a3b084da-9925-47d0-8f40-847a742d8449",
                                Name: codataccounting.String("Chelsea Reynolds"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ecf6b99b-c635-462e-bfdf-55c294c060b0",
                                Name: codataccounting.String("Janie Bogisich"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7764eef6-d0c6-4d6e-99c7-3dd634571509",
                                Name: codataccounting.String("Nelson Walker"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("doloremque"),
                            ID: "d3c5a1f9-c242-4c7b-a6a1-f30c73df5b67",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "890f42a4-bb43-48d8-9b26-0591d745e3c2",
                            Name: codataccounting.String("Lorraine Marquardt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3f567e0e-2527-465b-9d62-fcdace1f0121",
                            Name: codataccounting.String("Bernadette Waelchi"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "39e8f25c-d0d1-49d9-99f4-39e39266cbd9",
                            Name: codataccounting.String("Marta Kreiger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2b241136-95d1-4e66-98fc-c4596217c297",
                            Name: codataccounting.String("Beth Klocko"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "34254038-bfb5-4971-a981-90557389cedb",
                            Name: codataccounting.String("Spencer Kirlin"),
                        },
                    },
                    UnitAmount: 6340.91,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("dolor"),
            Note: codataccounting.String("occaecati"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("unde"),
                        Currency: codataccounting.String("labore"),
                        CurrencyRate: codataccounting.Float64(8658.06),
                        TotalAmount: codataccounting.Float64(4269.42),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6bc2ae48-0632-4b99-94b6-fa2206369828"),
                            Name: codataccounting.String("Vivian Dietrich"),
                        },
                        Currency: codataccounting.String("ab"),
                        CurrencyRate: codataccounting.Float64(32.41),
                        ID: codataccounting.String("006bef49-21ec-4205-bb74-9366ac8ee0f2"),
                        Note: codataccounting.String("libero"),
                        PaidOnDate: codataccounting.String("sapiente"),
                        Reference: codataccounting.String("veritatis"),
                        TotalAmount: codataccounting.Float64(5927.91),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("veniam"),
                        Currency: codataccounting.String("quos"),
                        CurrencyRate: codataccounting.Float64(5216.94),
                        TotalAmount: codataccounting.Float64(8135.44),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("40d03f3d-eba2-497b-a3e9-0bc40df868fd"),
                            Name: codataccounting.String("Mrs. Norma Greenholt"),
                        },
                        Currency: codataccounting.String("libero"),
                        CurrencyRate: codataccounting.Float64(2213.92),
                        ID: codataccounting.String("31d492f4-f127-4fb0-a0bf-1f8217978d0a"),
                        Note: codataccounting.String("eligendi"),
                        PaidOnDate: codataccounting.String("impedit"),
                        Reference: codataccounting.String("officia"),
                        TotalAmount: codataccounting.Float64(4850.97),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("aeb7b702-1a52-4046-b64e-99fb0e67e094"),
                    PurchaseOrderNumber: codataccounting.String("hic"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("dfed5540-ef53-4a34-a1b8-fe99731adc05"),
                    PurchaseOrderNumber: codataccounting.String("fugiat"),
                },
            },
            Reference: codataccounting.String("quas"),
            SourceModifiedDate: codataccounting.String("quis"),
            Status: shared.BillStatusVoid,
            SubTotal: 9080.3,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "illum": map[string]interface{}{
                        "rerum": "voluptate",
                        "perferendis": "maiores",
                        "harum": "ratione",
                        "molestias": "odio",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "4290d336-561e-4ca1-aef8-9451bd76eeeb",
                SupplierName: codataccounting.String("nemo"),
            },
            TaxAmount: 841.54,
            TotalAmount: 5623.39,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 2657.73,
                    Name: "Dana Bradtke",
                },
                shared.BillWithholdingTax{
                    Amount: 8325.89,
                    Name: "Mr. Vivian Harvey",
                },
                shared.BillWithholdingTax{
                    Amount: 437.97,
                    Name: "Janis Greenholt",
                },
                shared.BillWithholdingTax{
                    Amount: 7217.73,
                    Name: "Dr. Louise Will",
                },
            },
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(296127),
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
            Content: []byte("corrupti"),
            RequestBody: "exercitationem",
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
