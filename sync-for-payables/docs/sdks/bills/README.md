# Bills
(*Bills*)

## Overview

Bills

### Available Operations

* [Create](#create) - Create bill
* [Delete](#delete) - Delete bill
* [DeleteAttachment](#deleteattachment) - Delete bill attachment
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/types"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Create(ctx, operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: types.MustNewDecimalFromString("4865.89"),
            Currency: syncforpayables.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("6384.24"),
            DueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            ID: syncforpayables.String("<ID>"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("South"),
                    },
                    Description: syncforpayables.String("Vision-oriented responsive function"),
                    DiscountAmount: types.MustNewDecimalFromString("9510.62"),
                    DiscountPercentage: types.MustNewDecimalFromString("8915.1"),
                    IsDirectCost: syncforpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: syncforpayables.String("deposit"),
                    },
                    Quantity: types.MustNewDecimalFromString("3015.1"),
                    SubTotal: types.MustNewDecimalFromString("899.64"),
                    TaxAmount: types.MustNewDecimalFromString("7150.4"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("7926.2"),
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("Gasoline Screen mobile"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6562.56"),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: syncforpayables.String("Durham after"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: syncforpayables.String("Fay - Durgan"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforpayables.String("Fiat"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: syncforpayables.String("Grocery Borders Northwest"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("6519.85"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayables.Bool(false),
            },
            ModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Note: syncforpayables.String("animated Minivan"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("808.14"),
                        TotalAmount: types.MustNewDecimalFromString("2006.64"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("<ID>"),
                            Name: syncforpayables.String("West"),
                        },
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("3379.66"),
                        ID: syncforpayables.String("<ID>"),
                        Note: syncforpayables.String("Towels"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("payment 1080p"),
                        TotalAmount: types.MustNewDecimalFromString("2597.72"),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: syncforpayables.String("<ID>"),
                    PurchaseOrderNumber: syncforpayables.String("silver Indiana"),
                },
            },
            Reference: syncforpayables.String("Toyota Neptunium round"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusOpen,
            SubTotal: types.MustNewDecimalFromString("9057.54"),
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "in": map[string]interface{}{
                        "pariatur": "Dollar",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: syncforpayables.String("South"),
            },
            TaxAmount: types.MustNewDecimalFromString("4914.27"),
            TotalAmount: types.MustNewDecimalFromString("2663.33"),
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: types.MustNewDecimalFromString("4904.2"),
                    Name: "markets",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforpayables.Int(924484),
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

The *Delete bill* endpoint allows you to delete a specified bill from an accounting platform. 

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are itemized records of goods received or services provided to the SMB.

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
> This functionality is currently supported for our QuickBooks Online, Xero and Oracle NetSuite integrations.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Delete(ctx, operations.DeleteBillRequest{
        BillID: "7110701885",
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


## DeleteAttachment

The *Delete bill attachment* endpoint allows you to delete a specified bill attachment from an accounting platform.  

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices
that represent the SMB's financial obligations to their supplier for a
purchase of goods or services. 

### Process  

1. Pass the `{billId}` and `{attachmentId}` to the *Delete bill attachment* endpoint and store the `pushOperationKey` returned. 

2. Check the status of the delete operation by checking the status of push operation either via 

1. [Push operation webhook](https://docs.codat.io/introduction/webhookscore-rules-types#push-operation-status-has-changed) (advised), 

2. [Push operation status endpoint](https://docs.codat.io/sync-for-payables-api#/operations/get-push-operation). A `Success` status indicates that the bill attachment object was deleted from the accounting platform. 

3. (Optional) Check that the bill attachment was deleted from the accounting platform. 

>**Supported Integrations**
>
>This functionality is currently only supported for our QuickBooks Online integration. 

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.DeleteAttachment(ctx, operations.DeleteBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteBillAttachmentRequest](../../models/operations/deletebillattachmentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.DeleteBillAttachmentResponse](../../models/operations/deletebillattachmentresponse.md), error**


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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.DownloadAttachment(ctx, operations.DownloadBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "EILBDVJVNUAGVKRQ",
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Get(ctx, operations.GetBillRequest{
        BillID: "7110701885",
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.List(ctx, operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforpayables.String("-modifiedDate"),
        Page: syncforpayables.Int(1),
        PageSize: syncforpayables.Int(100),
        Query: syncforpayables.String("Northeast Metal Canada"),
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.ListAttachments(ctx, operations.ListBillAttachmentsRequest{
        BillID: "EILBDVJVNUAGVKRQ",
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/types"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Update(ctx, operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: types.MustNewDecimalFromString("8574.78"),
            Currency: syncforpayables.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("5971.29"),
            DueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            ID: syncforpayables.String("<ID>"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("dock Quality redundant"),
                    },
                    Description: syncforpayables.String("Visionary bi-directional analyzer"),
                    DiscountAmount: types.MustNewDecimalFromString("2782.81"),
                    DiscountPercentage: types.MustNewDecimalFromString("8965.01"),
                    IsDirectCost: syncforpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: syncforpayables.String("withdrawal extend"),
                    },
                    Quantity: types.MustNewDecimalFromString("2494.4"),
                    SubTotal: types.MustNewDecimalFromString("3668.07"),
                    TaxAmount: types.MustNewDecimalFromString("1395.79"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("6447.13"),
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("syndicate East Baht"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6298.17"),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: syncforpayables.String("guestbook driver users"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: syncforpayables.String("Schroeder - Nienow"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforpayables.String("Wooden Internal"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: syncforpayables.String("Dodge brightly"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7115.64"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayables.Bool(false),
            },
            ModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Note: syncforpayables.String("frictionless haptic"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("8238.97"),
                        TotalAmount: types.MustNewDecimalFromString("5143.61"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("<ID>"),
                            Name: syncforpayables.String("Diesel Avon"),
                        },
                        Currency: syncforpayables.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("8941.07"),
                        ID: syncforpayables.String("<ID>"),
                        Note: syncforpayables.String("male hack"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("absolve West"),
                        TotalAmount: types.MustNewDecimalFromString("182.83"),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: syncforpayables.String("<ID>"),
                    PurchaseOrderNumber: syncforpayables.String("quisquam"),
                },
            },
            Reference: syncforpayables.String("deliverables Ergonomic Money"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusPaid,
            SubTotal: types.MustNewDecimalFromString("4173.15"),
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "repellat": map[string]interface{}{
                        "voluptatibus": "Latvian",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: syncforpayables.String("Chips Omnigender tremendously"),
            },
            TaxAmount: types.MustNewDecimalFromString("9620.25"),
            TotalAmount: types.MustNewDecimalFromString("6952.8"),
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: types.MustNewDecimalFromString("3189.59"),
                    Name: "Benz",
                },
            },
        },
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: syncforpayables.Bool(false),
        TimeoutInMinutes: syncforpayables.Int(523103),
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.UploadAttachment(ctx, operations.UploadBillAttachmentRequest{
        RequestBody: &operations.UploadBillAttachmentRequestBody{
            Content: []byte("v/ghW&IC$x"),
            RequestBody: "Elegant Producer Electric",
        },
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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

