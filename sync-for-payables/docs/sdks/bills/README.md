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

The *Create bill* endpoint creates a new [bill](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-bills-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating a bill.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Create(ctx, operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codatsyncpayables.Float64(3442.89),
            Currency: codatsyncpayables.String("USD"),
            CurrencyRate: codatsyncpayables.Float64(5488.49),
            DueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            ID: codatsyncpayables.String("dc1ac600-dec0-401a-8802-e2ec09ff8f0f"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("6ff3477c-13e9-402c-9412-5b0960a66815"),
                        Name: codatsyncpayables.String("Harriet Gottlieb"),
                    },
                    Description: codatsyncpayables.String("deserunt"),
                    DiscountAmount: codatsyncpayables.Float64(9650.95),
                    DiscountPercentage: codatsyncpayables.Float64(6095.37),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "23c5949f-83f3-450c-b876-ffb901c6ecbb",
                        Name: codatsyncpayables.String("Silvia Considine"),
                    },
                    Quantity: 7581.94,
                    SubTotal: codatsyncpayables.Float64(9928.87),
                    TaxAmount: codatsyncpayables.Float64(4598.75),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(5000.21),
                        ID: codatsyncpayables.String("9ffafeda-53e5-4ae6-a0ac-184c2b9c247c"),
                        Name: codatsyncpayables.String("Byron Farrell"),
                    },
                    TotalAmount: codatsyncpayables.Float64(6751.26),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0e1942f3-2e55-4055-b56f-5d56d0bd0af2",
                                Name: codatsyncpayables.String("Mrs. Jerald Waelchi"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b4f62cba-3f89-441a-abc0-b80a6924d3b2",
                                Name: codatsyncpayables.String("Earnest Wisoky"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("quos"),
                            ID: "f895010f-5dd3-4d6f-a180-4e54c82f168a",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "3c8873e4-8438-40b1-b6b8-ca275a60a04c",
                            Name: codatsyncpayables.String("Sheryl Hermiston"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "99171b51-c1bd-4b1c-b4b8-88ebdfc4ccca",
                            Name: codatsyncpayables.String("Tracy Reinger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fc0b2dce-1087-43e4-ab00-6d678878ba85",
                            Name: codatsyncpayables.String("Joe Orn"),
                        },
                    },
                    UnitAmount: 1681.42,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("quas"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(9580.68),
                        TotalAmount: codatsyncpayables.Float64(9016.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("fa9c95f2-eac5-4565-9307-cfee81206e28"),
                            Name: codatsyncpayables.String("Peggy Windler"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(2666.8),
                        ID: codatsyncpayables.String("1c480d3f-2132-4af0-b102-d514f4cc6f18"),
                        Note: codatsyncpayables.String("expedita"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("excepturi"),
                        TotalAmount: codatsyncpayables.Float64(3966.1),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(6273.41),
                        TotalAmount: codatsyncpayables.Float64(4087.74),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("a4f77a87-ee3e-44be-b52c-65b34418e3bb"),
                            Name: codatsyncpayables.String("Joshua Schiller"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(4834.59),
                        ID: codatsyncpayables.String("5e0e8419-d8f8-44f1-84f3-e07edcc4aa5f"),
                        Note: codatsyncpayables.String("sequi"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("laborum"),
                        TotalAmount: codatsyncpayables.Float64(7489.73),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(162.52),
                        TotalAmount: codatsyncpayables.Float64(3699.41),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("a972e056-7282-427b-ad30-9470bf7a4fa8"),
                            Name: codatsyncpayables.String("Jody Wolff"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(6750.58),
                        ID: codatsyncpayables.String("6fae54eb-f60c-4321-b023-b75d2367fe1a"),
                        Note: codatsyncpayables.String("accusantium"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("minus"),
                        TotalAmount: codatsyncpayables.Float64(5509.94),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(4857.95),
                        TotalAmount: codatsyncpayables.Float64(5886.62),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("f0a396d9-0c36-44b7-815d-fbace188b1c4"),
                            Name: codatsyncpayables.String("Cornelius Corwin"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(4283.78),
                        ID: codatsyncpayables.String("ce611fee-b1c7-4cbd-b6ee-c74378ba2531"),
                        Note: codatsyncpayables.String("iusto"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("magnam"),
                        TotalAmount: codatsyncpayables.Float64(4877.99),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("c915ad2c-af5d-4d67-a3dc-0f5ae2f3a6b7"),
                    PurchaseOrderNumber: codatsyncpayables.String("doloremque"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("08787561-43f5-4a6c-98b5-5554080d40bc"),
                    PurchaseOrderNumber: codatsyncpayables.String("similique"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("cc6cbd6b-5f3e-4c90-9304-f926bad25538"),
                    PurchaseOrderNumber: codatsyncpayables.String("quasi"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("9b474b0e-d20e-4562-88ff-f639a910abdc"),
                    PurchaseOrderNumber: codatsyncpayables.String("fuga"),
                },
            },
            Reference: codatsyncpayables.String("tempore"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusUnknown,
            SubTotal: 3830.66,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "nisi": map[string]interface{}{
                        "provident": "laboriosam",
                        "accusamus": "ab",
                    },
                    "itaque": map[string]interface{}{
                        "eaque": "alias",
                        "qui": "consequuntur",
                        "vitae": "quidem",
                        "sequi": "amet",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "5d89acb3-ecfd-4a8d-8c54-9ef03004978a",
                SupplierName: codatsyncpayables.String("ex"),
            },
            TaxAmount: 1017.7,
            TotalAmount: 9535.64,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 842.88,
                    Name: "Ms. Darnell Denesik",
                },
                shared.BillWithholdingTax{
                    Amount: 5199.85,
                    Name: "Jared Koepp DVM",
                },
                shared.BillWithholdingTax{
                    Amount: 9686.89,
                    Name: "Jared Blick",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(681578),
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
    2. [Push operation status endpoint](https://docs.codat.io/sync-for-payables-api#/operations/get-push-operation).

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
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Delete(ctx, operations.DeleteBillRequest{
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperation != nil {
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

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support downloading a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

The *Get bill* endpoint returns a single bill for a given `billId`.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support getting a specific bill.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support getting a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetBillAttachmentRequest](../../models/operations/getbillattachmentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.GetBillAttachmentResponse](../../models/operations/getbillattachmentresponse.md), error**


## GetCreateUpdateModel

﻿The *Get create/update bill model* endpoint returns the expected data for the request payload when creating and updating a [bill](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company and integration.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating and updating a bill.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBillModelRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCreateUpdateBillModelRequest](../../models/operations/getcreateupdatebillmodelrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GetCreateUpdateBillModelResponse](../../models/operations/getcreateupdatebillmodelresponse.md), error**


## List

The *List bills* endpoint returns a list of [bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.List(ctx, operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncpayables.String("-modifiedDate"),
        Page: codatsyncpayables.Int(1),
        PageSize: codatsyncpayables.Int(100),
        Query: codatsyncpayables.String("sed"),
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

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support listing bill attachments.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

    if res.Attachments != nil {
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

The *Update bill* endpoint updates an existing [bill](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-bills-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating a bill.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Update(ctx, operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codatsyncpayables.Float64(1975.19),
            Currency: codatsyncpayables.String("EUR"),
            CurrencyRate: codatsyncpayables.Float64(5283.15),
            DueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            ID: codatsyncpayables.String("a97ff334-cddf-4857-a9e6-1876c6ab21d2"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("fc94d6fe-cd79-4939-8066-a6d2d0003553"),
                        Name: codatsyncpayables.String("Billie Schmitt"),
                    },
                    Description: codatsyncpayables.String("doloremque"),
                    DiscountAmount: codatsyncpayables.Float64(5168.33),
                    DiscountPercentage: codatsyncpayables.Float64(4352.66),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "fa21e915-2cb3-4119-967b-8e3c8db03408",
                        Name: codatsyncpayables.String("Leslie Sporer"),
                    },
                    Quantity: 3078.74,
                    SubTotal: codatsyncpayables.Float64(9382.44),
                    TaxAmount: codatsyncpayables.Float64(9454.31),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(8267.97),
                        ID: codatsyncpayables.String("455906d1-263d-448e-935c-2c9e81f30be3"),
                        Name: codatsyncpayables.String("Leroy Fisher Sr."),
                    },
                    TotalAmount: codatsyncpayables.Float64(8488.6),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "21657650-6641-4870-99d2-1f9ad030c4ec",
                                Name: codatsyncpayables.String("Carl Breitenberg V"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "36429068-b850-42a5-9e7f-73bc845e320a",
                                Name: codatsyncpayables.String("Christine Mueller"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("rerum"),
                            ID: "adf947c9-a867-4bc4-a426-665816ddca8e",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "1fcb4c59-3ec1-42cd-aad0-ec7afedbd80d",
                            Name: codatsyncpayables.String("Barry Funk"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "7f9390c5-8880-4983-9abf-9ef3ffdd9f7f",
                            Name: codatsyncpayables.String("Lillie Moen"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4d35724c-db0f-44d2-8118-7d56844eded8",
                            Name: codatsyncpayables.String("Ms. Madeline Miller"),
                        },
                    },
                    UnitAmount: 9090.93,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("628bdfc2-032b-46c8-b992-3b7e13584f7a"),
                        Name: codatsyncpayables.String("Carl Davis"),
                    },
                    Description: codatsyncpayables.String("praesentium"),
                    DiscountAmount: codatsyncpayables.Float64(6155.97),
                    DiscountPercentage: codatsyncpayables.Float64(1120.14),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f82ce115-7172-4305-b77d-cfa89df975e3",
                        Name: codatsyncpayables.String("Gertrude Kautzer"),
                    },
                    Quantity: 444.43,
                    SubTotal: codatsyncpayables.Float64(5968.2),
                    TaxAmount: codatsyncpayables.Float64(1458.41),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(9320.57),
                        ID: codatsyncpayables.String("9c3ddc5f-111d-4ea1-826d-541a4d190feb"),
                        Name: codatsyncpayables.String("Evelyn Kuhlman MD"),
                    },
                    TotalAmount: codatsyncpayables.Float64(8119.03),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c0dbbddb-4847-408f-b4e3-91e6bc158c4c",
                                Name: codatsyncpayables.String("Lila Hermiston"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "99ea3422-60e9-4b20-8ce7-8a1bd8fb7a0a",
                                Name: codatsyncpayables.String("Julie Homenick"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "723d4097-fa30-4e9a-b725-b29122030d83",
                                Name: codatsyncpayables.String("Dan Nolan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7799d22e-8c1f-4849-b825-fdc42c876c2c",
                                Name: codatsyncpayables.String("Marcella Windler"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("eligendi"),
                            ID: "fc1c7623-0f84-41fb-9bd2-3fdb14db6be5",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "85998e22-ae20-4da1-afc2-b271a289c57e",
                            Name: codatsyncpayables.String("Maurice Friesen"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "439d2224-6569-4462-8070-84f7ab37cef0",
                            Name: codatsyncpayables.String("Lois Crooks V"),
                        },
                    },
                    UnitAmount: 2907.61,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("db55410a-dc66-49af-90a2-6c7cdc981f06"),
                        Name: codatsyncpayables.String("Dr. Austin Little"),
                    },
                    Description: codatsyncpayables.String("libero"),
                    DiscountAmount: codatsyncpayables.Float64(7225),
                    DiscountPercentage: codatsyncpayables.Float64(2253.83),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3cfaa348-c31b-4f40-bee4-fcf0c42b78f1",
                        Name: codatsyncpayables.String("Sue Cronin"),
                    },
                    Quantity: 5681.62,
                    SubTotal: codatsyncpayables.Float64(5493.48),
                    TaxAmount: codatsyncpayables.Float64(6679.77),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(217.01),
                        ID: codatsyncpayables.String("dc766324-ccb0-46c8-8a12-d02529270b8d"),
                        Name: codatsyncpayables.String("Courtney Conroy"),
                    },
                    TotalAmount: codatsyncpayables.Float64(8506.28),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "95b8bcf2-4db9-4596-9335-2f74533994d7",
                                Name: codatsyncpayables.String("Gustavo Ullrich"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6e9389f5-abb7-4f66-a550-a28382ac483a",
                                Name: codatsyncpayables.String("Rufus Conn II"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bba65016-4e06-4f5b-b6ae-591bc8bdef36",
                                Name: codatsyncpayables.String("Denise Reilly"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("quod"),
                            ID: "205fda84-0774-4a68-a9a3-5d086b6f66fe",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "20e9f443-b425-47b9-92c8-dbda6a61efa2",
                            Name: codatsyncpayables.String("Sheryl Littel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fd0a9eba-47f7-4d3e-b049-640d6a1831c8",
                            Name: codatsyncpayables.String("Harriet Smitham"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "96fdf1ad-837a-4e80-81c1-9c95ba998678",
                            Name: codatsyncpayables.String("Doug Ernser"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "96991af3-88ce-4036-9444-8c7977a0ef2f",
                            Name: codatsyncpayables.String("Mr. Edna Howe"),
                        },
                    },
                    UnitAmount: 8905.05,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("feef9341-52ed-47e2-93f4-c157deaa7170"),
                        Name: codatsyncpayables.String("Randall Greenholt"),
                    },
                    Description: codatsyncpayables.String("quo"),
                    DiscountAmount: codatsyncpayables.Float64(8042.1),
                    DiscountPercentage: codatsyncpayables.Float64(9670.06),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "667aaf9b-bad1-485f-a431-d6bf5c838fbb",
                        Name: codatsyncpayables.String("Miss Jermaine Cole"),
                    },
                    Quantity: 4062.72,
                    SubTotal: codatsyncpayables.Float64(4836.26),
                    TaxAmount: codatsyncpayables.Float64(9630.94),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(7949.27),
                        ID: codatsyncpayables.String("4b425e99-e623-44c9-b7b7-9dfeb77a5c38"),
                        Name: codatsyncpayables.String("Theodore Reichert"),
                    },
                    TotalAmount: codatsyncpayables.Float64(5851.09),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e506ef89-0a54-4b47-9f16-f56d385a3c4a",
                                Name: codatsyncpayables.String("Miss Ben Ferry"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("provident"),
                            ID: "e26ced8f-9fdb-4941-8f63-bbf817837b01",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "dd788624-189e-4b44-873f-5033f19dbf12",
                            Name: codatsyncpayables.String("Della Trantow II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "eab9cd7e-5224-4a6a-8e12-3b7847ec59e1",
                            Name: codatsyncpayables.String("Milton Kirlin"),
                        },
                    },
                    UnitAmount: 7675.55,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("optio"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(7293.86),
                        TotalAmount: codatsyncpayables.Float64(3805.83),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("d7696ff3-c574-4750-9357-e44f51f8b084"),
                            Name: codatsyncpayables.String("Norman Casper"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(1182.36),
                        ID: codatsyncpayables.String("93a24546-7f94-4874-82d5-cc4972233e66"),
                        Note: codatsyncpayables.String("cum"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("rem"),
                        TotalAmount: codatsyncpayables.Float64(9774.72),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(8309.09),
                        TotalAmount: codatsyncpayables.Float64(129.17),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("0b979ef2-0387-4320-990c-cc1096400313"),
                            Name: codatsyncpayables.String("Nathan Toy II"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(9622.8),
                        ID: codatsyncpayables.String("65fe72dc-4077-4d0c-83f4-08efc15ceb4d"),
                        Note: codatsyncpayables.String("ea"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("quasi"),
                        TotalAmount: codatsyncpayables.Float64(8834.26),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(296),
                        TotalAmount: codatsyncpayables.Float64(9807.05),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("75aedf2a-cab5-48b9-91c9-26ddb589461e"),
                            Name: codatsyncpayables.String("Miss Elaine Considine"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(4128.97),
                        ID: codatsyncpayables.String("d9502f0e-a930-4b69-b7ac-2f72f8850090"),
                        Note: codatsyncpayables.String("tempora"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("quasi"),
                        TotalAmount: codatsyncpayables.Float64(829.15),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(5585.23),
                        TotalAmount: codatsyncpayables.Float64(1728.78),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("07888ec6-6183-4bfe-9659-eb40ec16faf7"),
                            Name: codatsyncpayables.String("Olivia Auer"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(1808.11),
                        ID: codatsyncpayables.String("a4da37cb-aaf4-4452-8484-2c9b2ad32daf"),
                        Note: codatsyncpayables.String("necessitatibus"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("quasi"),
                        TotalAmount: codatsyncpayables.Float64(6493.73),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("8f444457-3fec-4d47-b53f-63c8209379aa"),
                    PurchaseOrderNumber: codatsyncpayables.String("ex"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("9cd5fbcf-79da-418a-b822-bf95894e6861"),
                    PurchaseOrderNumber: codatsyncpayables.String("animi"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("db55f9e5-d751-4c9f-a8f7-502bfdc34508"),
                    PurchaseOrderNumber: codatsyncpayables.String("modi"),
                },
            },
            Reference: codatsyncpayables.String("illo"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusUnknown,
            SubTotal: 4742.67,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolore": map[string]interface{}{
                        "nostrum": "ex",
                        "amet": "voluptate",
                    },
                    "molestias": map[string]interface{}{
                        "ipsum": "hic",
                        "quidem": "odit",
                        "molestiae": "accusamus",
                        "quia": "inventore",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "f862657b-36fc-46b9-b587-ce525c67641a",
                SupplierName: codatsyncpayables.String("totam"),
            },
            TaxAmount: 2443.32,
            TotalAmount: 968.03,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 8897.63,
                    Name: "Brenda Greenholt",
                },
            },
        },
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        TimeoutInMinutes: codatsyncpayables.Int(766005),
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

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/sync-for-payables-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support uploading a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.UploadAttachment(ctx, operations.UploadBillAttachmentRequest{
        RequestBody: &operations.UploadBillAttachmentRequestBody{
            Content: []byte("quia"),
            RequestBody: "beatae",
        },
        BillID: "EILBDVJVNUAGVKRQ",
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

