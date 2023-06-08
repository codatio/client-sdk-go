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
            AmountDue: codataccounting.Float64(8237.18),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(2327.72),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("4ec1b781-b36a-4080-88d1-00efada200ef"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("22eb2164-cf9a-4b83-a6c7-23ffda9e06be"),
                        Name: codataccounting.String("Kyle Leffler"),
                    },
                    Description: codataccounting.String("quisquam"),
                    DiscountAmount: codataccounting.Float64(1202.77),
                    DiscountPercentage: codataccounting.Float64(9730.17),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c0e115c8-0bff-4918-944e-c42defcce8f1",
                        Name: codataccounting.String("Javier Koch"),
                    },
                    Quantity: 2027.96,
                    SubTotal: codataccounting.Float64(8955.13),
                    TaxAmount: codataccounting.Float64(4235.88),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2086.83),
                        ID: codataccounting.String("562a7b40-8f05-4e3d-88fd-af313a1f5fd9"),
                        Name: codataccounting.String("Rose Hills"),
                    },
                    TotalAmount: codataccounting.Float64(20.64),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "36f25ea9-44f3-4b75-ac11-f6c37a512624",
                                Name: codataccounting.String("Maxine Ernser"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bc05a23a-45ce-4fc5-bde1-0a0ce2169e51",
                                Name: codataccounting.String("Dorothy Boehm"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6dc5e347-6279-49bf-bbe6-949fb2bb4eca",
                                Name: codataccounting.String("Ben Satterfield"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nemo"),
                            ID: "db3adebd-5dae-4a4c-906a-8aa94c02644c",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e9d9a457-8adc-41ac-a00d-ec001ac802e2",
                            Name: codataccounting.String("Edmund Bednar"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8f0f816f-f347-47c1-be90-2c14125b0960",
                            Name: codataccounting.String("Brent Jacobs II"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1a472af9-23c5-4949-b83f-350cf876ffb9",
                            Name: codataccounting.String("Mildred Schinner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cbb4e243-cf78-49ff-afed-a53e5ae6e0ac",
                            Name: codataccounting.String("Mattie Gibson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b9c247c8-8373-4a40-a194-2f32e5505575",
                            Name: codataccounting.String("Essie Hayes"),
                        },
                    },
                    UnitAmount: 3827.64,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d0bd0af2-dfe1-43db-8f62-cba3f8941aeb"),
                        Name: codataccounting.String("Anthony Rath MD"),
                    },
                    Description: codataccounting.String("aliquid"),
                    DiscountAmount: codataccounting.Float64(5704.23),
                    DiscountPercentage: codataccounting.Float64(1674.35),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4d3b2ecf-cc8f-4895-810f-5dd3d6fa1804",
                        Name: codataccounting.String("Derek Haag"),
                    },
                    Quantity: 1624.5,
                    SubTotal: codataccounting.Float64(9824.45),
                    TaxAmount: codataccounting.Float64(815.81),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4004.48),
                        ID: codataccounting.String("8a363c88-73e4-4843-80b1-f6b8ca275a60"),
                        Name: codataccounting.String("Charles Grimes"),
                    },
                    TotalAmount: codataccounting.Float64(6137.02),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "cc699171-b51c-41bd-b1cf-4b888ebdfc4c",
                                Name: codataccounting.String("Lowell Olson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bc7fc0b2-dce1-4087-be42-b006d678878b",
                                Name: codataccounting.String("Dave Hegmann MD"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("enim"),
                            ID: "8208c54f-efa9-4c95-b2ea-c5565d307cfe",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1206e281-3fa4-4a41-8480-d3f2132af031",
                            Name: codataccounting.String("Nicole Stokes II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4cc6f18b-f962-41a6-a4f7-7a87ee3e4be7",
                            Name: codataccounting.String("Janice Russel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b34418e3-bb91-4c8d-975e-0e8419d8f84f",
                            Name: codataccounting.String("Suzanne Grant"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e07edcc4-aa5f-43ca-bd90-5a972e056728",
                            Name: codataccounting.String("Kathy Kris"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d309470b-f7a4-4fa8-bcf5-35a6fae54ebf",
                            Name: codataccounting.String("Sandra Rowe Sr."),
                        },
                    },
                    UnitAmount: 9730.96,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("sed"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(3455.14),
                        TotalAmount: codataccounting.Float64(8369.91),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("2367fe1a-0cc8-4df7-9f0a-396d90c364b7"),
                            Name: codataccounting.String("Ralph Hamill"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6496.57),
                        ID: codataccounting.String("ce188b1c-4ee2-4c8c-ace6-11feeb1c7cbd"),
                        Note: codataccounting.String("cum"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("saepe"),
                        TotalAmount: codataccounting.Float64(9347.82),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("74378ba2-5317-4747-9c91-5ad2caf5dd67"),
                    PurchaseOrderNumber: codataccounting.String("explicabo"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("3dc0f5ae-2f3a-46b7-8087-8756143f5a6c"),
                    PurchaseOrderNumber: codataccounting.String("provident"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("8b555540-80d4-40bc-acc6-cbd6b5f3ec90"),
                    PurchaseOrderNumber: codataccounting.String("provident"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("304f926b-ad25-4538-99b4-74b0ed20e562"),
                    PurchaseOrderNumber: codataccounting.String("aliquam"),
                },
            },
            Reference: codataccounting.String("blanditiis"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusDraft,
            SubTotal: 9906.52,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "nesciunt": map[string]interface{}{
                        "animi": "provident",
                        "beatae": "ipsa",
                        "mollitia": "nam",
                    },
                    "assumenda": map[string]interface{}{
                        "fuga": "tempore",
                        "commodi": "fugit",
                        "suscipit": "voluptate",
                        "nisi": "aliquid",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "96e1ec00-221b-4335-989a-cb3ecfda8d0c",
                SupplierName: codataccounting.String("veniam"),
            },
            TaxAmount: 3075.32,
            TotalAmount: 5911.43,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 9380.94,
                    Name: "Mrs. Edna Abbott",
                },
                shared.BillWithholdingTax{
                    Amount: 4620.97,
                    Name: "Dr. Luke Jaskolski",
                },
                shared.BillWithholdingTax{
                    Amount: 842.88,
                    Name: "Ms. Darnell Denesik",
                },
                shared.BillWithholdingTax{
                    Amount: 5199.85,
                    Name: "Jared Koepp DVM",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(968689),
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
        BillID: "EILBDVJVNUAGVKRQ",
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
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
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
        BillID: "EILBDVJVNUAGVKRQ",
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
        Query: codataccounting.String("maxime"),
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
        BillID: "7110701885",
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
            AmountDue: codataccounting.Float64(665.96),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(2359.7),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("2a3c80a9-7ff3-434c-9df8-57a9e61876c6"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("21d29dfc-94d6-4fec-9799-390066a6d2d0"),
                        Name: codataccounting.String("Linda Frami"),
                    },
                    Description: codataccounting.String("velit"),
                    DiscountAmount: codataccounting.Float64(1887.32),
                    DiscountPercentage: codataccounting.Float64(5553.86),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "cec086fa-21e9-4152-8b31-19167b8e3c8d",
                        Name: codataccounting.String("Kenneth Friesen IV"),
                    },
                    Quantity: 8310.67,
                    SubTotal: codataccounting.Float64(4159.53),
                    TaxAmount: codataccounting.Float64(8436.79),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2312.55),
                        ID: codataccounting.String("64ffd455-906d-4126-bd48-e935c2c9e81f"),
                        Name: codataccounting.String("Elizabeth Roberts"),
                    },
                    TotalAmount: codataccounting.Float64(9258.47),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "3202d721-6576-4506-a418-70d9d21f9ad0",
                                Name: codataccounting.String("Sharon Ruecker"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "cc11a083-6429-4068-b850-2a55e7f73bc8",
                                Name: codataccounting.String("Holly Torphy"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptatem"),
                            ID: "a319f4ba-df94-47c9-a867-bc4242666581",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "dca8ef51-fcb4-4c59-bec1-2cdaad0ec7af",
                            Name: codataccounting.String("Irving Rohan"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "df448a47-f939-40c5-8880-983dabf9ef3f",
                            Name: codataccounting.String("Sammy Simonis"),
                        },
                    },
                    UnitAmount: 4877.65,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f079af4d-3572-44cd-b0f4-d281187d5684"),
                        Name: codataccounting.String("Eloise Stoltenberg"),
                    },
                    Description: codataccounting.String("praesentium"),
                    DiscountAmount: codataccounting.Float64(3154.65),
                    DiscountPercentage: codataccounting.Float64(6581.25),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9065e628-bdfc-4203-ab6c-879923b7e135",
                        Name: codataccounting.String("Mario Zboncak"),
                    },
                    Quantity: 9301.27,
                    SubTotal: codataccounting.Float64(932.54),
                    TaxAmount: codataccounting.Float64(1812.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7733.55),
                        ID: codataccounting.String("6891f82c-e115-4717-a305-377dcfa89df9"),
                        Name: codataccounting.String("Annette Volkman"),
                    },
                    TotalAmount: codataccounting.Float64(4131.35),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "86092e9c-3ddc-45f1-91de-a1026d541a4d",
                                Name: codataccounting.String("Guadalupe Abernathy"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b21780bc-cc0d-4bbd-9b48-4708fb4e391e",
                                Name: codataccounting.String("Mrs. Susie Schowalter"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("porro"),
                            ID: "4c4e5459-9ea3-4422-a0e9-b200ce78a1bd",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b7a0a116-ce72-43d4-897f-a30e9af725b2",
                            Name: codataccounting.String("Willie Denesik I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d83f5aeb-7799-4d22-a8c1-f8493825fdc4",
                            Name: codataccounting.String("Jacquelyn Lueilwitz"),
                        },
                    },
                    UnitAmount: 8054.21,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2c2dfb4c-fc1c-4762-b0f8-41fb1bd23fdb"),
                        Name: codataccounting.String("Valerie Streich"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(9248.25),
                    DiscountPercentage: codataccounting.Float64(3422.26),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a685998e-22ae-420d-a16f-c2b271a289c5",
                        Name: codataccounting.String("Kellie Ledner"),
                    },
                    Quantity: 9261.42,
                    SubTotal: codataccounting.Float64(6050.43),
                    TaxAmount: codataccounting.Float64(585.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3029.05),
                        ID: codataccounting.String("39d22246-5694-4624-8708-4f7ab37cef02"),
                        Name: codataccounting.String("Ms. Kathryn Hayes"),
                    },
                    TotalAmount: codataccounting.Float64(8527.37),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "55410adc-669a-4f90-a26c-7cdc981f0689",
                                Name: codataccounting.String("Scott Schumm"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b33cfaa3-48c3-41bf-807e-e4fcf0c42b78",
                                Name: codataccounting.String("Gerald Herman"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6398a0dc-7663-424c-8b06-c8ca12d02529",
                                Name: codataccounting.String("Viola Abernathy"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("possimus"),
                            ID: "5722dd89-5b8b-4cf2-8db9-59693352f745",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "994d78de-3b6e-4938-9f5a-bb7f662550a2",
                            Name: codataccounting.String("Travis Leannon"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "483afd23-15bb-4a65-8164-e06f5bf6ae59",
                            Name: codataccounting.String("Mamie Schaefer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "def3612b-63c2-405f-9a84-0774a68a9a35",
                            Name: codataccounting.String("Daniel MacGyver"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6f66fef0-20e9-4f44-bb42-57b992c8dbda",
                            Name: codataccounting.String("Dr. Harriet Jones"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a2198258-fd0a-49eb-a47f-7d3ef049640d",
                            Name: codataccounting.String("Molly Carroll"),
                        },
                    },
                    UnitAmount: 1145.88,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("atque"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(9559.13),
                        TotalAmount: codataccounting.Float64(3228.49),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("96fdf1ad-837a-4e80-81c1-9c95ba998678"),
                            Name: codataccounting.String("Doug Ernser"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4187.87),
                        ID: codataccounting.String("991af388-ce03-4614-848c-7977a0ef2f53"),
                        Note: codataccounting.String("laboriosam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("sed"),
                        TotalAmount: codataccounting.Float64(5301.99),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(9324.44),
                        TotalAmount: codataccounting.Float64(9103.96),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f934152e-d7e2-453f-8c15-7deaa7170f44"),
                            Name: codataccounting.String("Jodi Sanford"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(3875.67),
                        ID: codataccounting.String("7aaf9bba-d185-4fe4-b1d6-bf5c838fbb8c"),
                        Note: codataccounting.String("aspernatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("impedit"),
                        TotalAmount: codataccounting.Float64(7226.58),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("7fc4b425-e99e-4623-8c9f-7b79dfeb77a5"),
                    PurchaseOrderNumber: codataccounting.String("quod"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("38d4baf9-1e50-46ef-890a-54b475f16f56"),
                    PurchaseOrderNumber: codataccounting.String("possimus"),
                },
            },
            Reference: codataccounting.String("nesciunt"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusPartiallyPaid,
            SubTotal: 6823.27,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "maxime": map[string]interface{}{
                        "laborum": "eligendi",
                        "autem": "adipisci",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "1b99e26c-ed8f-49fd-b941-0f63bbf81783",
                SupplierName: codataccounting.String("molestiae"),
            },
            TaxAmount: 7221.51,
            TotalAmount: 457.28,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 6506.78,
                    Name: "Josh Sporer",
                },
            },
        },
        BillID: "7110701885",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(424553),
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
            Content: []byte("eos"),
            RequestBody: "labore",
        },
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
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
