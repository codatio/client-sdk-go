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
            AmountDue: codataccounting.Float64(4620.97),
            Currency: codataccounting.String("atque"),
            CurrencyRate: codataccounting.Float64(6383.63),
            DueDate: codataccounting.String("ex"),
            ID: codataccounting.String("1fa1cf20-688f-477c-9ffc-71dca163f2a3"),
            IssueDate: "eligendi",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0a97ff33-4cdd-4f85-ba9e-61876c6ab21d"),
                        Name: codataccounting.String("Ada Stark"),
                    },
                    Description: codataccounting.String("unde"),
                    DiscountAmount: codataccounting.Float64(2814.54),
                    DiscountPercentage: codataccounting.Float64(8145.85),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "6fecd799-3900-466a-ad2d-000355338cec",
                        Name: codataccounting.String("Lena Kerluke"),
                    },
                    Quantity: 1440.58,
                    SubTotal: codataccounting.Float64(842.07),
                    TaxAmount: codataccounting.Float64(8996.52),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6119.7),
                        ID: codataccounting.String("152cb311-9167-4b8e-bc8d-b03408d6d364"),
                        Name: codataccounting.String("Irvin Shields"),
                    },
                    TotalAmount: codataccounting.Float64(3539.04),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "06d1263d-48e9-435c-ac9e-81f30be3e432",
                                Name: codataccounting.String("Wanda Stanton"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "16576506-6418-470d-9d21-f9ad030c4ecc",
                                Name: codataccounting.String("Ms. Teresa Ondricka"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6429068b-8502-4a55-a7f7-3bc845e320a3",
                                Name: codataccounting.String("Natasha Wiza"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolorum"),
                            ID: "df947c9a-867b-4c42-8266-65816ddca8ef",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "fcb4c593-ec12-4cda-ad0e-c7afedbd80df",
                            Name: codataccounting.String("Marjorie Lockman"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f9390c58-8809-483d-abf9-ef3ffdd9f7f0",
                            Name: codataccounting.String("Guadalupe Murazik"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d35724cd-b0f4-4d28-9187-d56844eded85",
                            Name: codataccounting.String("Wade Berge"),
                        },
                    },
                    UnitAmount: 9090.93,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("628bdfc2-032b-46c8-b992-3b7e13584f7a"),
                        Name: codataccounting.String("Carl Davis"),
                    },
                    Description: codataccounting.String("praesentium"),
                    DiscountAmount: codataccounting.Float64(6155.97),
                    DiscountPercentage: codataccounting.Float64(1120.14),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f82ce115-7172-4305-b77d-cfa89df975e3",
                        Name: codataccounting.String("Gertrude Kautzer"),
                    },
                    Quantity: 444.43,
                    SubTotal: codataccounting.Float64(5968.2),
                    TaxAmount: codataccounting.Float64(1458.41),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9320.57),
                        ID: codataccounting.String("9c3ddc5f-111d-4ea1-826d-541a4d190feb"),
                        Name: codataccounting.String("Evelyn Kuhlman MD"),
                    },
                    TotalAmount: codataccounting.Float64(8119.03),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c0dbbddb-4847-408f-b4e3-91e6bc158c4c",
                                Name: codataccounting.String("Lila Hermiston"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "99ea3422-60e9-4b20-8ce7-8a1bd8fb7a0a",
                                Name: codataccounting.String("Julie Homenick"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "723d4097-fa30-4e9a-b725-b29122030d83",
                                Name: codataccounting.String("Dan Nolan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7799d22e-8c1f-4849-b825-fdc42c876c2c",
                                Name: codataccounting.String("Marcella Windler"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eligendi"),
                            ID: "fc1c7623-0f84-41fb-9bd2-3fdb14db6be5",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "85998e22-ae20-4da1-afc2-b271a289c57e",
                            Name: codataccounting.String("Maurice Friesen"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "439d2224-6569-4462-8070-84f7ab37cef0",
                            Name: codataccounting.String("Lois Crooks V"),
                        },
                    },
                    UnitAmount: 2907.61,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("db55410a-dc66-49af-90a2-6c7cdc981f06"),
                        Name: codataccounting.String("Dr. Austin Little"),
                    },
                    Description: codataccounting.String("libero"),
                    DiscountAmount: codataccounting.Float64(7225),
                    DiscountPercentage: codataccounting.Float64(2253.83),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3cfaa348-c31b-4f40-bee4-fcf0c42b78f1",
                        Name: codataccounting.String("Sue Cronin"),
                    },
                    Quantity: 5681.62,
                    SubTotal: codataccounting.Float64(5493.48),
                    TaxAmount: codataccounting.Float64(6679.77),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(217.01),
                        ID: codataccounting.String("dc766324-ccb0-46c8-8a12-d02529270b8d"),
                        Name: codataccounting.String("Courtney Conroy"),
                    },
                    TotalAmount: codataccounting.Float64(8506.28),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "95b8bcf2-4db9-4596-9335-2f74533994d7",
                                Name: codataccounting.String("Gustavo Ullrich"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6e9389f5-abb7-4f66-a550-a28382ac483a",
                                Name: codataccounting.String("Rufus Conn II"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bba65016-4e06-4f5b-b6ae-591bc8bdef36",
                                Name: codataccounting.String("Denise Reilly"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quod"),
                            ID: "205fda84-0774-4a68-a9a3-5d086b6f66fe",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "20e9f443-b425-47b9-92c8-dbda6a61efa2",
                            Name: codataccounting.String("Sheryl Littel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fd0a9eba-47f7-4d3e-b049-640d6a1831c8",
                            Name: codataccounting.String("Harriet Smitham"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "96fdf1ad-837a-4e80-81c1-9c95ba998678",
                            Name: codataccounting.String("Doug Ernser"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "96991af3-88ce-4036-9444-8c7977a0ef2f",
                            Name: codataccounting.String("Mr. Edna Howe"),
                        },
                    },
                    UnitAmount: 8905.05,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("a"),
            Note: codataccounting.String("itaque"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("repellat"),
                        Currency: codataccounting.String("cupiditate"),
                        CurrencyRate: codataccounting.Float64(2372.08),
                        TotalAmount: codataccounting.Float64(3035.49),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("152ed7e2-53f4-4c15-bdea-a7170f445acc"),
                            Name: codataccounting.String("Rick Howell"),
                        },
                        Currency: codataccounting.String("officia"),
                        CurrencyRate: codataccounting.Float64(9949.02),
                        ID: codataccounting.String("9bbad185-fe43-41d6-bf5c-838fbb8c20cb"),
                        Note: codataccounting.String("ex"),
                        PaidOnDate: codataccounting.String("odio"),
                        Reference: codataccounting.String("delectus"),
                        TotalAmount: codataccounting.Float64(7949.27),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("ut"),
                        Currency: codataccounting.String("distinctio"),
                        CurrencyRate: codataccounting.Float64(2630.79),
                        TotalAmount: codataccounting.Float64(1764.18),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5e99e623-4c9f-47b7-9dfe-b77a5c38d4ba"),
                            Name: codataccounting.String("Kent Blanda"),
                        },
                        Currency: codataccounting.String("eaque"),
                        CurrencyRate: codataccounting.Float64(4064.62),
                        ID: codataccounting.String("ef890a54-b475-4f16-b56d-385a3c4ac631"),
                        Note: codataccounting.String("rerum"),
                        PaidOnDate: codataccounting.String("occaecati"),
                        Reference: codataccounting.String("provident"),
                        TotalAmount: codataccounting.Float64(8974.34),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("fugit"),
                        Currency: codataccounting.String("autem"),
                        CurrencyRate: codataccounting.Float64(7637.09),
                        TotalAmount: codataccounting.Float64(9107.26),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d8f9fdb9-410f-463b-bf81-7837b01afdd7"),
                            Name: codataccounting.String("Jordan Kassulke"),
                        },
                        Currency: codataccounting.String("sunt"),
                        CurrencyRate: codataccounting.Float64(5015.91),
                        ID: codataccounting.String("9eb44873-f503-43f1-9dbf-125ce4152eab"),
                        Note: codataccounting.String("error"),
                        PaidOnDate: codataccounting.String("placeat"),
                        Reference: codataccounting.String("temporibus"),
                        TotalAmount: codataccounting.Float64(4547.61),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("earum"),
                        Currency: codataccounting.String("minima"),
                        CurrencyRate: codataccounting.Float64(1419.86),
                        TotalAmount: codataccounting.Float64(1383.06),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4a6a0e12-3b78-447e-859e-1f67f3c4cce4"),
                            Name: codataccounting.String("Ricardo Stanton"),
                        },
                        Currency: codataccounting.String("sint"),
                        CurrencyRate: codataccounting.Float64(4001.45),
                        ID: codataccounting.String("ff3c5747-5013-457e-84f5-1f8b084c3197"),
                        Note: codataccounting.String("officiis"),
                        PaidOnDate: codataccounting.String("dicta"),
                        Reference: codataccounting.String("excepturi"),
                        TotalAmount: codataccounting.Float64(2336.18),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("245467f9-4874-4c2d-9cc4-972233e66bd8"),
                    PurchaseOrderNumber: codataccounting.String("voluptatibus"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("e5d00b97-9ef2-4038-b320-590ccc109640"),
                    PurchaseOrderNumber: codataccounting.String("voluptatem"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("313b3e50-44f6-45fe-b2dc-4077d0cc3f40"),
                    PurchaseOrderNumber: codataccounting.String("corrupti"),
                },
            },
            Reference: codataccounting.String("itaque"),
            SourceModifiedDate: codataccounting.String("earum"),
            Status: shared.BillStatusVoid,
            SubTotal: 1199.28,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "impedit": map[string]interface{}{
                        "cum": "dolore",
                        "illum": "ea",
                        "officiis": "quasi",
                        "accusamus": "animi",
                    },
                    "necessitatibus": map[string]interface{}{
                        "maiores": "odio",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "5aedf2ac-ab58-4b99-9c92-6ddb589461e7",
                SupplierName: codataccounting.String("numquam"),
            },
            TaxAmount: 1446.21,
            TotalAmount: 788.02,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 7486.56,
                    Name: "Rafael Schuster",
                },
                shared.BillWithholdingTax{
                    Amount: 263.89,
                    Name: "Elsa Adams",
                },
                shared.BillWithholdingTax{
                    Amount: 6238.68,
                    Name: "Margaret Rau",
                },
                shared.BillWithholdingTax{
                    Amount: 9787.97,
                    Name: "Lee Runte",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(460597),
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
        BillID: "explicabo",
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
        BillID: "delectus",
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
        BillID: "quos",
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
        BillID: "deleniti",
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
        Query: codataccounting.String("enim"),
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
        BillID: "sit",
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
            AmountDue: codataccounting.Float64(279.55),
            Currency: codataccounting.String("natus"),
            CurrencyRate: codataccounting.Float64(293.46),
            DueDate: codataccounting.String("tempora"),
            ID: codataccounting.String("91160820-7888-4ec6-a183-bfe9659eb40e"),
            IssueDate: "quod",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6faf75b0-b532-4a4d-a37c-baaf4452c484"),
                        Name: codataccounting.String("Sheri McGlynn"),
                    },
                    Description: codataccounting.String("est"),
                    DiscountAmount: codataccounting.Float64(8687.09),
                    DiscountPercentage: codataccounting.Float64(2163.79),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "2dafe81a-88f4-4444-973f-ecd47353f63c",
                        Name: codataccounting.String("Aaron Becker"),
                    },
                    Quantity: 4722.9,
                    SubTotal: codataccounting.Float64(6040.27),
                    TaxAmount: codataccounting.Float64(6463.21),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6621.84),
                        ID: codataccounting.String("69cd5fbc-f79d-4a18-a782-2bf95894e686"),
                        Name: codataccounting.String("Lynda Schuppe"),
                    },
                    TotalAmount: codataccounting.Float64(3129.53),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9e5d751c-9fe8-4f75-82bf-dc3450841f17",
                                Name: codataccounting.String("Eleanor Gibson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "379f3fb2-7e21-4f86-a657-b36fc6b9f587",
                                Name: codataccounting.String("Bert Hand"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c67641a8-312e-4504-bb4c-21ccb423abcd",
                                Name: codataccounting.String("Arturo Bosco"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "abdd88e7-1f6c-4482-92d7-771e7fd07400",
                                Name: codataccounting.String("Tomas Zieme"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("consequuntur"),
                            ID: "9de1dd70-97b5-4da0-8c57-fa6c78a216e1",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "afeca619-1498-4140-b64f-f8ae170ef03b",
                            Name: codataccounting.String("Josefina Durgan"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "aa868555-9667-432a-a5dc-b6682cb70f8c",
                            Name: codataccounting.String("Kristopher Herman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6e91b9a9-f748-446e-ac33-09db0536d9e7",
                            Name: codataccounting.String("Mr. Vicky Parker"),
                        },
                    },
                    UnitAmount: 9386.72,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("quis"),
            Note: codataccounting.String("dolorem"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("sed"),
                        Currency: codataccounting.String("quo"),
                        CurrencyRate: codataccounting.Float64(917.36),
                        TotalAmount: codataccounting.Float64(761.45),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a25a8bf9-2f97-4428-ad9a-9f8bf8221125"),
                            Name: codataccounting.String("Marion Medhurst"),
                        },
                        Currency: codataccounting.String("blanditiis"),
                        CurrencyRate: codataccounting.Float64(2326.02),
                        ID: codataccounting.String("87f7a79c-d72c-4d24-84da-21729f2ac41e"),
                        Note: codataccounting.String("tenetur"),
                        PaidOnDate: codataccounting.String("exercitationem"),
                        Reference: codataccounting.String("nihil"),
                        TotalAmount: codataccounting.Float64(1547.23),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("quis"),
                        Currency: codataccounting.String("maiores"),
                        CurrencyRate: codataccounting.Float64(1168.07),
                        TotalAmount: codataccounting.Float64(1158.49),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("69ac1e41-d8a2-43c2-be34-f2dfa4a197f6"),
                            Name: codataccounting.String("Clay Monahan"),
                        },
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(3687.61),
                        ID: codataccounting.String("1fe17120-9985-43e9-b543-d854439ee224"),
                        Note: codataccounting.String("non"),
                        PaidOnDate: codataccounting.String("laboriosam"),
                        Reference: codataccounting.String("accusantium"),
                        TotalAmount: codataccounting.Float64(2736.38),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("aliquam"),
                        Currency: codataccounting.String("dolorem"),
                        CurrencyRate: codataccounting.Float64(7122.09),
                        TotalAmount: codataccounting.Float64(7711.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("154188c2-f56e-485d-a783-2eabd617c3b0"),
                            Name: codataccounting.String("Leon Braun"),
                        },
                        Currency: codataccounting.String("ut"),
                        CurrencyRate: codataccounting.Float64(6942.92),
                        ID: codataccounting.String("f01bad87-06d4-4608-abfb-dc41ff5d4e2a"),
                        Note: codataccounting.String("saepe"),
                        PaidOnDate: codataccounting.String("labore"),
                        Reference: codataccounting.String("doloribus"),
                        TotalAmount: codataccounting.Float64(7040.54),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("cb35d176-38f1-4edb-b835-9ecc5cb860f8"),
                    PurchaseOrderNumber: codataccounting.String("impedit"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("d580ba73-810e-44fe-8447-297cd3b1dd3b"),
                    PurchaseOrderNumber: codataccounting.String("libero"),
                },
            },
            Reference: codataccounting.String("impedit"),
            SourceModifiedDate: codataccounting.String("repudiandae"),
            Status: shared.BillStatusOpen,
            SubTotal: 2814.16,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "harum": map[string]interface{}{
                        "aliquid": "blanditiis",
                        "labore": "repudiandae",
                    },
                    "reiciendis": map[string]interface{}{
                        "exercitationem": "voluptatem",
                        "beatae": "qui",
                        "laboriosam": "temporibus",
                        "in": "veritatis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "cffbd0eb-74b8-4421-953b-44bd3c43159d",
                SupplierName: codataccounting.String("adipisci"),
            },
            TaxAmount: 2439.38,
            TotalAmount: 8815.49,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 5913.67,
                    Name: "Mr. Carmen Schmidt Jr.",
                },
                shared.BillWithholdingTax{
                    Amount: 2305.3,
                    Name: "Isaac Hyatt",
                },
            },
        },
        BillID: "fuga",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(266461),
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
            Content: []byte("et"),
            RequestBody: "eveniet",
        },
        BillID: "aliquid",
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
