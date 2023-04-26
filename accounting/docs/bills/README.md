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

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating a bill.

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
    req := operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(7068.72),
            Currency: codataccounting.String("non"),
            CurrencyRate: codataccounting.Float64(6495.34),
            DueDate: codataccounting.String("assumenda"),
            ID: codataccounting.String("ebd5daea-4c50-46a8-aa94-c02644cf5e9d"),
            IssueDate: "error",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4578adc1-ac60-40de-8001-ac802e2ec09f"),
                        Name: codataccounting.String("Dr. Armando Wunsch"),
                    },
                    Description: codataccounting.String("dicta"),
                    DiscountAmount: codataccounting.Float64(3805.95),
                    DiscountPercentage: codataccounting.Float64(9382.57),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f3477c13-e902-4c14-925b-0960a668151a",
                        Name: codataccounting.String("Constance Dach"),
                    },
                    Quantity: 6095.37,
                    SubTotal: codataccounting.Float64(1512.3),
                    TaxAmount: codataccounting.Float64(2009.5),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8054.63),
                        ID: codataccounting.String("5949f83f-350c-4f87-affb-901c6ecbb4e2"),
                        Name: codataccounting.String("Lillian Rosenbaum"),
                    },
                    TotalAmount: codataccounting.Float64(5000.21),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ffafeda5-3e5a-4e6e-8ac1-84c2b9c247c8",
                                Name: codataccounting.String("Allen Kozey"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "40e1942f-32e5-4505-9756-f5d56d0bd0af",
                                Name: codataccounting.String("Elena Zieme I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "db4f62cb-a3f8-4941-aebc-0b80a6924d3b",
                                Name: codataccounting.String("Eloise Rowe"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quo"),
                            ID: "8f895010-f5dd-43d6-ba18-04e54c82f168",
                        },
                        IsBilledTo: shared.BilledToTypeEnumCustomer,
                        IsRebilledTo: shared.BilledToTypeEnumUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "63c8873e-4843-480b-9f6b-8ca275a60a04",
                            Name: codataccounting.String("Jay Morar"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "699171b5-1c1b-4db1-8f4b-888ebdfc4ccc",
                            Name: codataccounting.String("Marshall McClure"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7fc0b2dc-e108-473e-82b0-06d678878ba8",
                            Name: codataccounting.String("Kay Bradtke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8208c54f-efa9-4c95-b2ea-c5565d307cfe",
                            Name: codataccounting.String("Hugh Carroll III"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e2813fa4-a41c-4480-93f2-132af03102d5",
                            Name: codataccounting.String("Danielle Willms"),
                        },
                    },
                    UnitAmount: 7505.37,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6f18bf96-21a6-4a4f-b7a8-7ee3e4be752c"),
                        Name: codataccounting.String("Beatrice Purdy"),
                    },
                    Description: codataccounting.String("quaerat"),
                    DiscountAmount: codataccounting.Float64(1039.88),
                    DiscountPercentage: codataccounting.Float64(5069.66),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "e3bb91c8-d975-4e0e-8419-d8f84f144f3e",
                        Name: codataccounting.String("Joy Toy"),
                    },
                    Quantity: 7690.47,
                    SubTotal: codataccounting.Float64(3028.78),
                    TaxAmount: codataccounting.Float64(6778.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6448.27),
                        ID: codataccounting.String("5f3cabd9-05a9-472e-8567-28227b2d3094"),
                        Name: codataccounting.String("Sharon Raynor"),
                    },
                    TotalAmount: codataccounting.Float64(6671.69),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "fa87cf53-5a6f-4ae5-8ebf-60c321f023b7",
                                Name: codataccounting.String("Paulette Dibbert"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7fe1a0cc-8df7-49f0-a396-d90c364b7c15",
                                Name: codataccounting.String("Emilio Rau"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("accusamus"),
                            ID: "188b1c4e-e2c8-4c6c-a611-feeb1c7cbdb6",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "c74378ba-2531-4774-bdc9-15ad2caf5dd6",
                            Name: codataccounting.String("Judith Feest"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f5ae2f3a-6b70-4087-8756-143f5a6c98b5",
                            Name: codataccounting.String("Holly Harber V"),
                        },
                    },
                    UnitAmount: 408.74,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d40bcacc-6cbd-46b5-b3ec-909304f926ba"),
                        Name: codataccounting.String("Wayne Hintz"),
                    },
                    Description: codataccounting.String("voluptatum"),
                    DiscountAmount: codataccounting.Float64(987.59),
                    DiscountPercentage: codataccounting.Float64(6225.66),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b474b0ed-20e5-4624-8fff-639a910abdca",
                        Name: codataccounting.String("Edgar Corkery"),
                    },
                    Quantity: 3931.22,
                    SubTotal: codataccounting.Float64(3971.6),
                    TaxAmount: codataccounting.Float64(5897.12),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3857.6),
                        ID: codataccounting.String("e1ec0022-1b33-45d8-9acb-3ecfda8d0c54"),
                        Name: codataccounting.String("Mrs. Hugo Wehner Jr."),
                    },
                    TotalAmount: codataccounting.Float64(2737.32),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "78a61fa1-cf20-4688-b77c-1ffc71dca163",
                                Name: codataccounting.String("Aaron O'Kon"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "80a97ff3-34cd-4df8-97a9-e61876c6ab21",
                                Name: codataccounting.String("Victor Mosciski"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c94d6fec-d799-4390-866a-6d2d00035533",
                                Name: codataccounting.String("Wilbert Walsh IV"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("iure"),
                            ID: "fa21e915-2cb3-4119-967b-8e3c8db03408",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d364ffd4-5590-46d1-a63d-48e935c2c9e8",
                            Name: codataccounting.String("Miss Jeannie Emmerich"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e43202d7-2165-4765-8664-1870d9d21f9a",
                            Name: codataccounting.String("Miss Michael Ferry"),
                        },
                    },
                    UnitAmount: 8907.65,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("cumque"),
            Note: codataccounting.String("maxime"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("beatae"),
                        Currency: codataccounting.String("id"),
                        CurrencyRate: codataccounting.Float64(90.6),
                        TotalAmount: codataccounting.Float64(5515.76),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("36429068-b850-42a5-9e7f-73bc845e320a"),
                            Name: codataccounting.String("Christine Mueller"),
                        },
                        Currency: codataccounting.String("rerum"),
                        CurrencyRate: codataccounting.Float64(6786.95),
                        ID: codataccounting.String("df947c9a-867b-4c42-8266-65816ddca8ef"),
                        Note: codataccounting.String("nemo"),
                        PaidOnDate: codataccounting.String("illo"),
                        Reference: codataccounting.String("doloribus"),
                        TotalAmount: codataccounting.Float64(7666.7),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("4c593ec1-2cda-4ad0-ac7a-fedbd80df448"),
                    PurchaseOrderNumber: codataccounting.String("similique"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("47f9390c-5888-4098-bdab-f9ef3ffdd9f7"),
                    PurchaseOrderNumber: codataccounting.String("voluptatibus"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("079af4d3-5724-4cdb-8f4d-281187d56844"),
                    PurchaseOrderNumber: codataccounting.String("accusamus"),
                },
            },
            Reference: codataccounting.String("nulla"),
            SourceModifiedDate: codataccounting.String("repudiandae"),
            Status: shared.BillStatusEnumDraft,
            SubTotal: 5057.99,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "animi": map[string]interface{}{
                        "quae": "eum",
                        "nostrum": "eveniet",
                        "laboriosam": "ratione",
                    },
                    "blanditiis": map[string]interface{}{
                        "illum": "reiciendis",
                        "placeat": "dolores",
                        "consequatur": "nesciunt",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "2b6c8799-23b7-4e13-984f-7ae12c6891f8",
                SupplierName: codataccounting.String("aspernatur"),
            },
            TaxAmount: 7552.4,
            TotalAmount: 9178.77,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 810.53,
                    Name: "Heidi Bode",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(209602),
    }

    res, err := s.Bills.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillResponse != nil {
        // handle response
    }
}
```

## Delete

Deletes a bill from the accounting package for a given company.

> **Supported Integrations**
> 
> This functionality is currently only supported for our Oracle NetSuite and QuickBooks Online integrations. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    req := operations.DeleteBillRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.Delete(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## DownloadAttachment

Download bill attachment

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
    req := operations.DownloadBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.DownloadAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Get

Get bill

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
    req := operations.GetBillRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Bills.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bill != nil {
        // handle response
    }
}
```

## GetAttachment

Get bill attachment

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
    req := operations.GetBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.GetAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

Get create/update bill model.

 > **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating and updating a bill.

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
    req := operations.GetCreateUpdateBillsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.GetCreateUpdateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

Gets the latest bills for a company, with pagination

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
    req := operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("eaque"),
    }

    res, err := s.Bills.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bills != nil {
        // handle response
    }
}
```

## ListAttachments

List bill attachments

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
    req := operations.ListBillAttachmentsRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.ListAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## Update

Posts an updated bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support updating a bill.

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
    req := operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(3474.6),
            Currency: codataccounting.String("amet"),
            CurrencyRate: codataccounting.Float64(4541.65),
            DueDate: codataccounting.String("voluptate"),
            ID: codataccounting.String("dcfa89df-975e-4356-a860-92e9c3ddc5f1"),
            IssueDate: "vitae",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("dea1026d-541a-44d1-90fe-b21780bccc0d"),
                        Name: codataccounting.String("Kelvin Shanahan"),
                    },
                    Description: codataccounting.String("magnam"),
                    DiscountAmount: codataccounting.Float64(5123.49),
                    DiscountPercentage: codataccounting.Float64(2726.35),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "708fb4e3-91e6-4bc1-98c4-c4e54599ea34",
                        Name: codataccounting.String("Dr. Rose Hoeger"),
                    },
                    Quantity: 7495.37,
                    SubTotal: codataccounting.Float64(1861.3),
                    TaxAmount: codataccounting.Float64(374.77),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(92.48),
                        ID: codataccounting.String("ce78a1bd-8fb7-4a0a-916c-e723d4097fa3"),
                        Name: codataccounting.String("Alyssa Morar"),
                    },
                    TotalAmount: codataccounting.Float64(4974.8),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5b291220-30d8-43f5-aeb7-799d22e8c1f8",
                                Name: codataccounting.String("Erika Farrell III"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("delectus"),
                            ID: "dc42c876-c2c2-4dfb-8cfc-1c76230f841f",
                        },
                        IsBilledTo: shared.BilledToTypeEnumCustomer,
                        IsRebilledTo: shared.BilledToTypeEnumUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "bd23fdb1-4db6-4be5-a685-998e22ae20da",
                            Name: codataccounting.String("Lucy Wilkinson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "271a289c-57e8-454e-9043-9d2224656946",
                            Name: codataccounting.String("Juanita Batz IV"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4f7ab37c-ef02-4225-994d-b55410adc669",
                            Name: codataccounting.String("Miss Cary McKenzie"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6c7cdc98-1f06-4898-9d6b-b33cfaa348c3",
                            Name: codataccounting.String("Antoinette Wolf IV"),
                        },
                    },
                    UnitAmount: 9334.56,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("necessitatibus"),
            Note: codataccounting.String("numquam"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("eligendi"),
                        Currency: codataccounting.String("sapiente"),
                        CurrencyRate: codataccounting.Float64(0.73),
                        TotalAmount: codataccounting.Float64(7704.67),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("42b78f15-6263-498a-8dc7-66324ccb06c8"),
                            Name: codataccounting.String("Rex Bode"),
                        },
                        Currency: codataccounting.String("sit"),
                        CurrencyRate: codataccounting.Float64(1720.42),
                        ID: codataccounting.String("529270b8-d572-42dd-895b-8bcf24db9596"),
                        Note: codataccounting.String("provident"),
                        PaidOnDate: codataccounting.String("amet"),
                        Reference: codataccounting.String("dolor"),
                        TotalAmount: codataccounting.Float64(3440.1),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("qui"),
                        Currency: codataccounting.String("tenetur"),
                        CurrencyRate: codataccounting.Float64(4773.52),
                        TotalAmount: codataccounting.Float64(2925.71),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("533994d7-8de3-4b6e-9389-f5abb7f66255"),
                            Name: codataccounting.String("Monique Denesik"),
                        },
                        Currency: codataccounting.String("totam"),
                        CurrencyRate: codataccounting.Float64(1835.04),
                        ID: codataccounting.String("ac483afd-2315-4bba-a501-64e06f5bf6ae"),
                        Note: codataccounting.String("nemo"),
                        PaidOnDate: codataccounting.String("molestias"),
                        Reference: codataccounting.String("architecto"),
                        TotalAmount: codataccounting.Float64(7112.75),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("quisquam"),
                        Currency: codataccounting.String("praesentium"),
                        CurrencyRate: codataccounting.Float64(7080.75),
                        TotalAmount: codataccounting.Float64(8264.78),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ef3612b6-3c20-45fd-a840-774a68a9a35d"),
                            Name: codataccounting.String("Leona Hodkiewicz"),
                        },
                        Currency: codataccounting.String("maiores"),
                        CurrencyRate: codataccounting.Float64(4120.24),
                        ID: codataccounting.String("6fef020e-9f44-43b4-a57b-992c8dbda6a6"),
                        Note: codataccounting.String("dicta"),
                        PaidOnDate: codataccounting.String("recusandae"),
                        Reference: codataccounting.String("sapiente"),
                        TotalAmount: codataccounting.Float64(6617.64),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("odit"),
                        Currency: codataccounting.String("inventore"),
                        CurrencyRate: codataccounting.Float64(6122.38),
                        TotalAmount: codataccounting.Float64(5421.87),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("258fd0a9-eba4-47f7-93ef-049640d6a183"),
                            Name: codataccounting.String("Rosalie Lindgren"),
                        },
                        Currency: codataccounting.String("temporibus"),
                        CurrencyRate: codataccounting.Float64(9559.13),
                        ID: codataccounting.String("596fdf1a-d837-4ae8-8c1c-19c95ba99867"),
                        Note: codataccounting.String("voluptatum"),
                        PaidOnDate: codataccounting.String("sapiente"),
                        Reference: codataccounting.String("deserunt"),
                        TotalAmount: codataccounting.Float64(2212.4),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("696991af-388c-4e03-a144-48c7977a0ef2"),
                    PurchaseOrderNumber: codataccounting.String("delectus"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("536028ef-eef9-4341-92ed-7e253f4c157d"),
                    PurchaseOrderNumber: codataccounting.String("voluptates"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("aa7170f4-45ac-4cf6-a7aa-f9bbad185fe4"),
                    PurchaseOrderNumber: codataccounting.String("nesciunt"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("1d6bf5c8-38fb-4b8c-a0cb-67fc4b425e99"),
                    PurchaseOrderNumber: codataccounting.String("debitis"),
                },
            },
            Reference: codataccounting.String("laboriosam"),
            SourceModifiedDate: codataccounting.String("eos"),
            Status: shared.BillStatusEnumOpen,
            SubTotal: 2796.79,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "occaecati": map[string]interface{}{
                        "voluptate": "tempore",
                        "in": "omnis",
                        "possimus": "tenetur",
                        "recusandae": "expedita",
                    },
                    "iusto": map[string]interface{}{
                        "harum": "ad",
                        "quod": "ratione",
                    },
                    "totam": map[string]interface{}{
                        "dolore": "nam",
                        "officia": "maiores",
                        "cupiditate": "illo",
                        "saepe": "enim",
                    },
                    "eaque": map[string]interface{}{
                        "eveniet": "delectus",
                        "deleniti": "provident",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "0a54b475-f16f-456d-b85a-3c4ac631b99e",
                SupplierName: codataccounting.String("fugit"),
            },
            TaxAmount: 4188.92,
            TotalAmount: 7637.09,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 8528.73,
                    Name: "Tommie Mraz",
                },
                shared.BillWithholdingTax{
                    Amount: 7063.71,
                    Name: "Dr. Marcus Bosco",
                },
                shared.BillWithholdingTax{
                    Amount: 2458.49,
                    Name: "Rudolph Weimann IV",
                },
                shared.BillWithholdingTax{
                    Amount: 5422.42,
                    Name: "Mr. Nellie Reichert",
                },
            },
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(990454),
    }

    res, err := s.Bills.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillResponse != nil {
        // handle response
    }
}
```

## UploadAttachment

Upload bill attachment

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
    req := operations.UploadBillAttachmentRequest{
        RequestBody: &operations.UploadBillAttachmentRequestBody{
            Content: []byte("at"),
            RequestBody: "quibusdam",
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.UploadAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
