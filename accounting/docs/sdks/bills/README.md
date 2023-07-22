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

The *Create bill* endpoint creates a new [bill](https://docs.codat.io/accounting-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
            AmountDue: codataccounting.Float64(9453.2),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(3712.95),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("a9e61876-c6ab-421d-a9df-c94d6fecd799"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0066a6d2-d000-4355-b38c-ec086fa21e91"),
                        Name: codataccounting.String("Kathryn Runolfsdottir"),
                    },
                    Description: codataccounting.String("beatae"),
                    DiscountAmount: codataccounting.Float64(1234.95),
                    DiscountPercentage: codataccounting.Float64(5658.45),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "167b8e3c-8db0-4340-8d6d-364ffd455906",
                        Name: codataccounting.String("Keith Crist"),
                    },
                    Quantity: 8436.59,
                    SubTotal: codataccounting.Float64(2552.64),
                    TaxAmount: codataccounting.Float64(5231.09),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8846.22),
                        ID: codataccounting.String("935c2c9e-81f3-40be-be43-202d72165765"),
                        Name: codataccounting.String("Kristin Howell IV"),
                    },
                    TotalAmount: codataccounting.Float64(4935.79),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d9d21f9a-d030-4c4e-8c11-a0836429068b",
                                Name: codataccounting.String("Pedro Armstrong"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quaerat"),
                            ID: "5e7f73bc-845e-4320-a319-f4badf947c9a",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7bc42426-6658-416d-9ca8-ef51fcb4c593",
                            Name: codataccounting.String("Devin Boyle"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "aad0ec7a-fedb-4d80-9f44-8a47f9390c58",
                            Name: codataccounting.String("Willard Barrows"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3dabf9ef-3ffd-4d9f-bf07-9af4d35724cd",
                            Name: codataccounting.String("Jeffrey Wisoky"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "281187d5-6844-4ede-985a-9065e628bdfc",
                            Name: codataccounting.String("Elizabeth Douglas"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6c879923-b7e1-4358-8f7a-e12c6891f82c",
                            Name: codataccounting.String("Keith Bode"),
                        },
                    },
                    UnitAmount: 823.96,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("72305377-dcfa-489d-b975-e356686092e9"),
                        Name: codataccounting.String("Norman Skiles"),
                    },
                    Description: codataccounting.String("minima"),
                    DiscountAmount: codataccounting.Float64(9519.01),
                    DiscountPercentage: codataccounting.Float64(1048.34),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "11dea102-6d54-41a4-9190-feb21780bccc",
                        Name: codataccounting.String("Muriel Reichel"),
                    },
                    Quantity: 8504.06,
                    SubTotal: codataccounting.Float64(7468.34),
                    TaxAmount: codataccounting.Float64(2973.25),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5123.49),
                        ID: codataccounting.String("4708fb4e-391e-46bc-958c-4c4e54599ea3"),
                        Name: codataccounting.String("Nicole Christiansen DVM"),
                    },
                    TotalAmount: codataccounting.Float64(5757.53),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "200ce78a-1bd8-4fb7-a0a1-16ce723d4097",
                                Name: codataccounting.String("Dr. Doug Dibbert"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "af725b29-1220-430d-83f5-aeb7799d22e8",
                                Name: codataccounting.String("Roger Zulauf"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "93825fdc-42c8-476c-ac2d-fb4cfc1c7623",
                                Name: codataccounting.String("Johanna Lueilwitz DVM"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nam"),
                            ID: "1bd23fdb-14db-46be-9a68-5998e22ae20d",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6fc2b271-a289-4c57-a854-e90439d22246",
                            Name: codataccounting.String("Kristin McDermott"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "407084f7-ab37-4cef-8222-5194db55410a",
                            Name: codataccounting.String("Garrett Hoeger"),
                        },
                    },
                    UnitAmount: 6668.17,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f90a26c7-cdc9-481f-8689-81d6bb33cfaa"),
                        Name: codataccounting.String("Clara Larson"),
                    },
                    Description: codataccounting.String("veritatis"),
                    DiscountAmount: codataccounting.Float64(7217.23),
                    DiscountPercentage: codataccounting.Float64(9747.75),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "407ee4fc-f0c4-42b7-8f15-626398a0dc76",
                        Name: codataccounting.String("Rosa Considine"),
                    },
                    Quantity: 8031.44,
                    SubTotal: codataccounting.Float64(7133.71),
                    TaxAmount: codataccounting.Float64(371.81),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4086.77),
                        ID: codataccounting.String("c8ca12d0-2529-4270-b8d5-722dd895b8bc"),
                        Name: codataccounting.String("Ernest Grimes"),
                    },
                    TotalAmount: codataccounting.Float64(5854.45),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9693352f-7453-4399-8d78-de3b6e9389f5",
                                Name: codataccounting.String("Gerardo Ritchie"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "662550a2-8382-4ac4-83af-d2315bba6501",
                                Name: codataccounting.String("Ms. Eva Upton"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("veniam"),
                            ID: "bf6ae591-bc8b-4def-b612-b63c205fda84",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "74a68a9a-35d0-486b-af66-fef020e9f443",
                            Name: codataccounting.String("Randall Daniel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "992c8dbd-a6a6-41ef-a219-8258fd0a9eba",
                            Name: codataccounting.String("Allison Wiza"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3ef04964-0d6a-4183-9c87-adf596fdf1ad",
                            Name: codataccounting.String("Tony Konopelski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "80c1c19c-95ba-4998-a78f-a3f696991af3",
                            Name: codataccounting.String("Daryl Schmitt I"),
                        },
                    },
                    UnitAmount: 4209.1,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("numquam"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(7914.54),
                        TotalAmount: codataccounting.Float64(4524.81),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("977a0ef2-f536-4028-afee-f934152ed7e2"),
                            Name: codataccounting.String("Ethel Windler"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(3534.93),
                        ID: codataccounting.String("7deaa717-0f44-45ac-8f66-7aaf9bbad185"),
                        Note: codataccounting.String("sapiente"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("ut"),
                        TotalAmount: codataccounting.Float64(2010.05),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(3843.54),
                        TotalAmount: codataccounting.Float64(6963.24),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f5c838fb-b8c2-40cb-a7fc-4b425e99e623"),
                            Name: codataccounting.String("Robyn McCullough"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(4473.23),
                        ID: codataccounting.String("9dfeb77a-5c38-4d4b-af91-e506ef890a54"),
                        Note: codataccounting.String("nam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("odio"),
                        TotalAmount: codataccounting.Float64(3427.72),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("16f56d38-5a3c-44ac-a31b-99e26ced8f9f"),
                    PurchaseOrderNumber: codataccounting.String("repellendus"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("b9410f63-bbf8-4178-b7b0-1afdd7886241"),
                    PurchaseOrderNumber: codataccounting.String("blanditiis"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("9eb44873-f503-43f1-9dbf-125ce4152eab"),
                    PurchaseOrderNumber: codataccounting.String("error"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("cd7e5224-a6a0-4e12-bb78-47ec59e1f67f"),
                    PurchaseOrderNumber: codataccounting.String("amet"),
                },
            },
            Reference: codataccounting.String("cumque"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusVoid,
            SubTotal: 7763.34,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempora": map[string]interface{}{
                        "suscipit": "illum",
                        "iusto": "aliquid",
                        "sint": "aliquid",
                    },
                    "repellat": map[string]interface{}{
                        "consectetur": "eligendi",
                        "ullam": "nihil",
                        "eius": "dignissimos",
                        "corporis": "perferendis",
                    },
                    "architecto": map[string]interface{}{
                        "corporis": "nihil",
                    },
                    "officiis": map[string]interface{}{
                        "magnam": "maiores",
                        "ipsam": "dicta",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "f8b084c3-197e-4193-a245-467f94874c2d",
                SupplierName: codataccounting.String("ipsam"),
            },
            TaxAmount: 7503.43,
            TotalAmount: 7841.2,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 5961.33,
                    Name: "Lois Dibbert",
                },
                shared.BillWithholdingTax{
                    Amount: 9041.93,
                    Name: "Beth Ritchie",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(977472),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateBillRequest](../../models/operations/createbillrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.CreateBillResponse](../../models/operations/createbillresponse.md), error**


## Delete

﻿The *Delete bill* endpoint allows you to delete a specified bill from an accounting platform. 

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are itemized records of goods received or services provided to the SMB.

### Process 
1. Pass the `{billId}` to the *Delete bill* endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete operation by checking the status of push operation either via
    1. [Push operation webhook](https://docs.codat.io/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
    2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation).

   A `Success` status indicates that the bill object was deleted from the accounting platform.
3. (Optional) Check that the bill was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting a bill from an accounting platform might cause related objects to be modified. For example, if you delete a paid bill in QuickBooks Online, the bill is deleted but the bill payment against that bill is not. The bill payment is converted to a payment on account.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Delete | Details                                                                                                      |  
|-------------|-------------|--------------------------------------------------------------------------------------------------------------|
| QuickBooks Online | No          | -                                                                                                            |
| Oracle NetSuite   | No          | When deleting a bill that's already linked to a bill payment, you must delete the linked bill payment first. |

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
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteBillRequest](../../models/operations/deletebillrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.DeleteBillResponse](../../models/operations/deletebillresponse.md), error**


## DownloadAttachment

The *Download bill attachment* endpoint downloads a specific attachment for a given `billId` and `attachmentId`.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support downloading a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DownloadBillAttachmentRequest](../../models/operations/downloadbillattachmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.DownloadBillAttachmentResponse](../../models/operations/downloadbillattachmentresponse.md), error**


## Get

The *Get bill* endpoint returns a single bill for a given billId.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support getting a specific bill.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        BillID: "EILBDVJVNUAGVKRQ",
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

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetBillRequest](../../models/operations/getbillrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |


### Response

**[*operations.GetBillResponse](../../models/operations/getbillresponse.md), error**


## GetAttachment

The *Get bill attachment* endpoint returns a specific attachment for a given `billId` and `attachmentId`.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support getting a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
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

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetBillAttachmentRequest](../../models/operations/getbillattachmentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.GetBillAttachmentResponse](../../models/operations/getbillattachmentresponse.md), error**


## GetCreateUpdateModel

﻿The *Get create/update bill model* endpoint returns the expected data for the request payload when creating and updating a [bill](https://docs.codat.io/accounting-api#/schemas/Bill) for a given company and integration.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating and updating a bill.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCreateUpdateBillsModelRequest](../../models/operations/getcreateupdatebillsmodelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCreateUpdateBillsModelResponse](../../models/operations/getcreateupdatebillsmodelresponse.md), error**


## List

The *List bills* endpoint returns a list of [bills](https://docs.codat.io/accounting-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        Query: codataccounting.String("voluptatem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bills != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListBillsRequest](../../models/operations/listbillsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |


### Response

**[*operations.ListBillsResponse](../../models/operations/listbillsresponse.md), error**


## ListAttachments

The *List bill attachments* endpoint returns a list of attachments available to download for a given `billId`.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support listing bill attachments.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListBillAttachmentsRequest](../../models/operations/listbillattachmentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.ListBillAttachmentsResponse](../../models/operations/listbillattachmentsresponse.md), error**


## Update

The *Update bill* endpoint updates an existing [bill](https://docs.codat.io/accounting-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
            AmountDue: codataccounting.Float64(5684.19),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(6091.61),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("f2038732-0590-4ccc-9096-400313b3e504"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("65fe72dc-4077-4d0c-83f4-08efc15ceb4d"),
                        Name: codataccounting.String("Cecelia Boyer"),
                    },
                    Description: codataccounting.String("necessitatibus"),
                    DiscountAmount: codataccounting.Float64(296),
                    DiscountPercentage: codataccounting.Float64(9807.05),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "75aedf2a-cab5-48b9-91c9-26ddb589461e",
                        Name: codataccounting.String("Miss Elaine Considine"),
                    },
                    Quantity: 9245.59,
                    SubTotal: codataccounting.Float64(4128.97),
                    TaxAmount: codataccounting.Float64(8203.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5890.98),
                        ID: codataccounting.String("502f0ea9-30b6-49f7-ac2f-72f885009049"),
                        Name: codataccounting.String("Ms. Carolyn Jacobson"),
                    },
                    TotalAmount: codataccounting.Float64(170.4),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "888ec661-83bf-4e96-99eb-40ec16faf75b",
                                Name: codataccounting.String("Juana Herman"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a4da37cb-aaf4-4452-8484-2c9b2ad32daf",
                                Name: codataccounting.String("Bob Boyle"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("blanditiis"),
                            ID: "f4444573-fecd-4473-93f6-3c8209379aa6",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d5fbcf79-da18-4a78-a2bf-95894e6861ad",
                            Name: codataccounting.String("Pedro Haley"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5d751c9f-e8f7-4502-bfdc-3450841f1764",
                            Name: codataccounting.String("Bernice Jaskolski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9f3fb27e-21f8-4626-97b3-6fc6b9f587ce",
                            Name: codataccounting.String("Sara Hegmann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7641a831-2e50-447b-8c21-ccb423abcdc9",
                            Name: codataccounting.String("Lana Pfannerstill"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "dd88e71f-6c48-4252-9777-1e7fd074009e",
                            Name: codataccounting.String("Jaime Schumm"),
                        },
                    },
                    UnitAmount: 8304.73,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e1dd7097-b5da-408c-97fa-6c78a216e19b"),
                        Name: codataccounting.String("Randal Walker"),
                    },
                    Description: codataccounting.String("laboriosam"),
                    DiscountAmount: codataccounting.Float64(717.34),
                    DiscountPercentage: codataccounting.Float64(5842.92),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1498140b-64ff-48ae-970e-f03b5f37e4aa",
                        Name: codataccounting.String("Ricardo Lynch"),
                    },
                    Quantity: 3573.88,
                    SubTotal: codataccounting.Float64(6141.75),
                    TaxAmount: codataccounting.Float64(4116.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4196),
                        ID: codataccounting.String("732aa5dc-b668-42cb-b0f8-cfd5fb6e91b9"),
                        Name: codataccounting.String("Ross Wilderman"),
                    },
                    TotalAmount: codataccounting.Float64(5110.54),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "6e2c3309-db05-436d-9e75-ca006f5392c1",
                                Name: codataccounting.String("Jodi Crona"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8bf92f97-428a-4d9a-9f8b-f8221125359d",
                                Name: codataccounting.String("Guy Feest"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("repellat"),
                            ID: "7a79cd72-cd24-484d-a217-29f2ac41ef57",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f1169ac1-e41d-48a2-bc23-e34f2dfa4a19",
                            Name: codataccounting.String("Elsa Kerluke"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "22151fe1-7120-4998-93e9-f543d854439e",
                            Name: codataccounting.String("Randy Collier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "60443bc1-5418-48c2-b56e-85da7832eabd",
                            Name: codataccounting.String("Frances Kuhlman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b0d51a44-bf01-4bad-8706-d46082bfbdc4",
                            Name: codataccounting.String("Lucia Wintheiser"),
                        },
                    },
                    UnitAmount: 2975.85,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e2ae4fb5-cb35-4d17-a38f-1edb78359ecc"),
                        Name: codataccounting.String("Blanca Prohaska"),
                    },
                    Description: codataccounting.String("doloremque"),
                    DiscountAmount: codataccounting.Float64(9640.52),
                    DiscountPercentage: codataccounting.Float64(5582.01),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "cd580ba7-3810-4e4f-a444-7297cd3b1dd3",
                        Name: codataccounting.String("Randolph Russel"),
                    },
                    Quantity: 2814.16,
                    SubTotal: codataccounting.Float64(4724.44),
                    TaxAmount: codataccounting.Float64(6906.54),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4559.58),
                        ID: codataccounting.String("684eff50-126d-471c-bfbd-0eb74b842195"),
                        Name: codataccounting.String("Melody Grady"),
                    },
                    TotalAmount: codataccounting.Float64(8614.06),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c43159d3-3e59-453c-8011-39863aa41e6c",
                                Name: codataccounting.String("Rebecca Schmitt DVM"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dicta"),
                            ID: "fcb51c9a-41ff-4be9-8bd7-95ee65e076cc",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "bf616ea5-c716-4419-b4b9-0f2e09d19d2f",
                            Name: codataccounting.String("Nicholas Wisoky"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e105944b-935d-4237-a72f-90849d6aed4a",
                            Name: codataccounting.String("Earnest Rogahn"),
                        },
                    },
                    UnitAmount: 2048.77,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7cd9222c-9ff5-4749-9aab-fa2e761f0ca4"),
                        Name: codataccounting.String("Francisco Hauck"),
                    },
                    Description: codataccounting.String("reiciendis"),
                    DiscountAmount: codataccounting.Float64(1170.53),
                    DiscountPercentage: codataccounting.Float64(234.1),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "31e6899f-0c20-401e-a2cd-55cc0584a184",
                        Name: codataccounting.String("Christian Hirthe"),
                    },
                    Quantity: 4525.15,
                    SubTotal: codataccounting.Float64(1150.77),
                    TaxAmount: codataccounting.Float64(9714.36),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7665.01),
                        ID: codataccounting.String("820c65b0-37bb-48e0-8c88-5187e4de04af"),
                        Name: codataccounting.String("Naomi Schneider"),
                    },
                    TotalAmount: codataccounting.Float64(8135.45),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b46aa1cf-d6d8-428d-a013-191129646645",
                                Name: codataccounting.String("Raymond Sporer DVM"),
                            },
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
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("molestiae"),
                            ID: "764eef6d-0c6d-46ed-9c73-dd634571509a",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "870d3c5a-1f9c-4242-87b6-6a1f30c73df5",
                            Name: codataccounting.String("Ms. Angel Kreiger"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0f42a4bb-438d-485b-a605-91d745e3c205",
                            Name: codataccounting.String("Sylvester Mante"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f567e0e2-5276-45b1-962f-cdace1f01216",
                            Name: codataccounting.String("Cornelius Crooks"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9e8f25cd-0d19-4d95-9f43-9e39266cbd95",
                            Name: codataccounting.String("Clinton Oberbrunner"),
                        },
                    },
                    UnitAmount: 6980.88,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("magnam"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(3911.05),
                        TotalAmount: codataccounting.Float64(5979.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5d1e6698-fcc4-4596-a17c-297767633425"),
                            Name: codataccounting.String("Mary Fisher"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7169.75),
                        ID: codataccounting.String("5971e981-9055-4738-9ced-bac7fda39594"),
                        Note: codataccounting.String("pariatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("laboriosam"),
                        TotalAmount: codataccounting.Float64(7444.74),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("2ae48063-2b99-454b-afa2-206369828553"),
                    PurchaseOrderNumber: codataccounting.String("optio"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("b10006be-f492-41ec-a053-b749366ac8ee"),
                    PurchaseOrderNumber: codataccounting.String("voluptatem"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("f2bf1958-8d40-4d03-b3de-ba297be3e90b"),
                    PurchaseOrderNumber: codataccounting.String("nobis"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("40df868f-d524-405c-b331-d492f4f127fb"),
                    PurchaseOrderNumber: codataccounting.String("aperiam"),
                },
            },
            Reference: codataccounting.String("saepe"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusVoid,
            SubTotal: 9438.65,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "delectus": map[string]interface{}{
                        "fugit": "inventore",
                        "reprehenderit": "sint",
                        "dignissimos": "voluptatum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "d0acca77-aeb7-4b70-a1a5-2046b64e99fb",
                SupplierName: codataccounting.String("doloremque"),
            },
            TaxAmount: 8871.99,
            TotalAmount: 3942.08,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 8997.35,
                    Name: "Daisy Graham",
                },
                shared.BillWithholdingTax{
                    Amount: 9504.65,
                    Name: "Kristopher Herman",
                },
            },
        },
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(924506),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateBillRequest](../../models/operations/updatebillrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.UpdateBillResponse](../../models/operations/updatebillresponse.md), error**


## UploadAttachment

The *Upload bill attachment* endpoint uploads an attachment and assigns it against a specific `billId`.

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/accounting-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support uploading a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
            Content: []byte("a"),
            RequestBody: "exercitationem",
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

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UploadBillAttachmentRequest](../../models/operations/uploadbillattachmentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.UploadBillAttachmentResponse](../../models/operations/uploadbillattachmentresponse.md), error**

