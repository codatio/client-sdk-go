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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
            AmountDue: codataccounting.Float64(1017.7),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(6730.1),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("cf20688f-77c1-4ffc-b1dc-a163f2a3c80a"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ff334cdd-f857-4a9e-a187-6c6ab21d29df"),
                        Name: codataccounting.String("Austin Grady"),
                    },
                    Description: codataccounting.String("doloribus"),
                    DiscountAmount: codataccounting.Float64(9260.27),
                    DiscountPercentage: codataccounting.Float64(7874.52),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "d7993900-66a6-4d2d-8003-55338cec086f",
                        Name: codataccounting.String("Billy Boehm"),
                    },
                    Quantity: 1164.63,
                    SubTotal: codataccounting.Float64(3690.99),
                    TaxAmount: codataccounting.Float64(1631.81),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7652.71),
                        ID: codataccounting.String("b3119167-b8e3-4c8d-b034-08d6d364ffd4"),
                        Name: codataccounting.String("Ms. Samantha Metz"),
                    },
                    TotalAmount: codataccounting.Float64(1168.67),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "63d48e93-5c2c-49e8-9f30-be3e43202d72",
                                Name: codataccounting.String("Beth Hand"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("minima"),
                            ID: "06641870-d9d2-41f9-ad03-0c4ecc11a083",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "29068b85-02a5-45e7-b73b-c845e320a319",
                            Name: codataccounting.String("Herbert Prohaska"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "947c9a86-7bc4-4242-a665-816ddca8ef51",
                            Name: codataccounting.String("Spencer Rath"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "593ec12c-daad-40ec-bafe-dbd80df448a4",
                            Name: codataccounting.String("Lela Moore"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0c588809-83da-4bf9-af3f-fdd9f7f079af",
                            Name: codataccounting.String("Muriel Durgan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "24cdb0f4-d281-4187-9568-44eded85a906",
                            Name: codataccounting.String("Laverne Hudson"),
                        },
                    },
                    UnitAmount: 6977.29,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("dfc2032b-6c87-4992-bb7e-13584f7ae12c"),
                        Name: codataccounting.String("Dr. Maxine Morissette"),
                    },
                    Description: codataccounting.String("aspernatur"),
                    DiscountAmount: codataccounting.Float64(7552.4),
                    DiscountPercentage: codataccounting.Float64(9178.77),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "11571723-0537-47dc-ba89-df975e356686",
                        Name: codataccounting.String("Katrina Considine"),
                    },
                    Quantity: 7677.78,
                    SubTotal: codataccounting.Float64(2250.01),
                    TaxAmount: codataccounting.Float64(8339.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8373.27),
                        ID: codataccounting.String("c5f111de-a102-46d5-81a4-d190feb21780"),
                        Name: codataccounting.String("Clark Schmitt PhD"),
                    },
                    TotalAmount: codataccounting.Float64(7188.79),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ddb48470-8fb4-4e39-9e6b-c158c4c4e545",
                                Name: codataccounting.String("Marion Ward"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "42260e9b-200c-4e78-a1bd-8fb7a0a116ce",
                                Name: codataccounting.String("Tammy Dickens"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "097fa30e-9af7-425b-a912-2030d83f5aeb",
                                Name: codataccounting.String("Stella Medhurst"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("sunt"),
                            ID: "2e8c1f84-9382-45fd-842c-876c2c2dfb4c",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1c76230f-841f-4b1b-923f-db14db6be5a6",
                            Name: codataccounting.String("Lewis Mante"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "22ae20da-16fc-42b2-b1a2-89c57e854e90",
                            Name: codataccounting.String("Tiffany Mayert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "22465694-6240-4708-8f7a-b37cef022251",
                            Name: codataccounting.String("Marcus Stehr"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5410adc6-69af-490a-a6c7-cdc981f06898",
                            Name: codataccounting.String("Ernestine Jast"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "33cfaa34-8c31-4bf4-87ee-4fcf0c42b78f",
                            Name: codataccounting.String("Samantha Huels"),
                        },
                    },
                    UnitAmount: 1913.12,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("corrupti"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7702.62),
                        TotalAmount: codataccounting.Float64(4972.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("66324ccb-06c8-4ca1-ad02-529270b8d572"),
                            Name: codataccounting.String("Lynette Stark"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(7399.38),
                        ID: codataccounting.String("8bcf24db-9596-4933-92f7-4533994d78de"),
                        Note: codataccounting.String("amet"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("vel"),
                        TotalAmount: codataccounting.Float64(9004.38),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5456.29),
                        TotalAmount: codataccounting.Float64(5873.37),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f5abb7f6-6255-40a2-8382-ac483afd2315"),
                            Name: codataccounting.String("Malcolm O'Connell"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(767.86),
                        ID: codataccounting.String("64e06f5b-f6ae-4591-bc8b-def3612b63c2"),
                        Note: codataccounting.String("alias"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("a"),
                        TotalAmount: codataccounting.Float64(8168.25),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(3106.48),
                        TotalAmount: codataccounting.Float64(458.5),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("774a68a9-a35d-4086-b6f6-6fef020e9f44"),
                            Name: codataccounting.String("Celia Gottlieb"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(7289.48),
                        ID: codataccounting.String("992c8dbd-a6a6-41ef-a219-8258fd0a9eba"),
                        Note: codataccounting.String("labore"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("voluptatibus"),
                        TotalAmount: codataccounting.Float64(4668.62),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("3ef04964-0d6a-4183-9c87-adf596fdf1ad"),
                    PurchaseOrderNumber: codataccounting.String("praesentium"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("37ae80c1-c19c-495b-a998-678fa3f69699"),
                    PurchaseOrderNumber: codataccounting.String("ab"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("af388ce0-3614-4448-8797-7a0ef2f53602"),
                    PurchaseOrderNumber: codataccounting.String("voluptatum"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("efeef934-152e-4d7e-a53f-4c157deaa717"),
                    PurchaseOrderNumber: codataccounting.String("consequatur"),
                },
            },
            Reference: codataccounting.String("delectus"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusOpen,
            SubTotal: 3626.93,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "quo": map[string]interface{}{
                        "delectus": "laboriosam",
                        "laboriosam": "quam",
                        "fuga": "officia",
                        "repellat": "cupiditate",
                    },
                    "soluta": map[string]interface{}{
                        "culpa": "fugiat",
                        "inventore": "atque",
                        "ad": "sapiente",
                    },
                    "voluptates": map[string]interface{}{
                        "nesciunt": "ab",
                        "quibusdam": "suscipit",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "bf5c838f-bb8c-420c-b67f-c4b425e99e62",
                SupplierName: codataccounting.String("amet"),
            },
            TaxAmount: 2796.79,
            TotalAmount: 7835.39,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 9707.03,
                    Name: "Whitney King",
                },
                shared.BillWithholdingTax{
                    Amount: 9503.37,
                    Name: "Rudolph Kshlerin",
                },
                shared.BillWithholdingTax{
                    Amount: 3216.54,
                    Name: "Philip Leannon",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(722889),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
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

[Bills](https://docs.codat.io/accounting-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support getting a bill attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        Query: codataccounting.String("saepe"),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.ListAttachments(ctx, operations.ListBillAttachmentsRequest{
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
            AmountDue: codataccounting.Float64(524.07),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(9101.32),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("890a54b4-75f1-46f5-ad38-5a3c4ac631b9"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("26ced8f9-fdb9-4410-b63b-bf817837b01a"),
                        Name: codataccounting.String("Josh Sporer"),
                    },
                    Description: codataccounting.String("rem"),
                    DiscountAmount: codataccounting.Float64(4245.53),
                    DiscountPercentage: codataccounting.Float64(1777.9),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4189eb44-873f-4503-bf19-dbf125ce4152",
                        Name: codataccounting.String("Julius Ratke"),
                    },
                    Quantity: 8382.27,
                    SubTotal: codataccounting.Float64(4547.61),
                    TaxAmount: codataccounting.Float64(9354.93),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3272.63),
                        ID: codataccounting.String("224a6a0e-123b-4784-bec5-9e1f67f3c4cc"),
                        Name: codataccounting.String("Ricky Rempel"),
                    },
                    TotalAmount: codataccounting.Float64(4785.76),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "96ff3c57-4750-4135-be44-f51f8b084c31",
                                Name: codataccounting.String("Ms. Jamie Torphy"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a245467f-9487-44c2-95cc-4972233e66bd",
                                Name: codataccounting.String("Jan Torphy"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("aut"),
                            ID: "0b979ef2-0387-4320-990c-cc1096400313",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e5044f65-fe72-4dc4-877d-0cc3f408efc1",
                            Name: codataccounting.String("Pat Upton"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6e1eae0f-75ae-4df2-acab-58b991c926dd",
                            Name: codataccounting.String("Jon Lockman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "61e7421c-be6d-4950-af0e-a930b69f7ac2",
                            Name: codataccounting.String("Erik Champlin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "85009049-1160-4820-b888-ec66183bfe96",
                            Name: codataccounting.String("Lindsey Trantow"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0ec16faf-75b0-4b53-aa4d-a37cbaaf4452",
                            Name: codataccounting.String("Jon Lemke DDS"),
                        },
                    },
                    UnitAmount: 5867.23,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b2ad32da-fe81-4a88-b444-4573fecd4735"),
                        Name: codataccounting.String("Lana Keeling"),
                    },
                    Description: codataccounting.String("quos"),
                    DiscountAmount: codataccounting.Float64(1523.02),
                    DiscountPercentage: codataccounting.Float64(544.98),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9379aa69-cd5f-4bcf-b9da-18a7822bf958",
                        Name: codataccounting.String("Tom Thiel"),
                    },
                    Quantity: 4331.94,
                    SubTotal: codataccounting.Float64(1151.37),
                    TaxAmount: codataccounting.Float64(6574.85),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8187.39),
                        ID: codataccounting.String("b55f9e5d-751c-49fe-8f75-02bfdc345084"),
                        Name: codataccounting.String("Josefina Borer"),
                    },
                    TotalAmount: codataccounting.Float64(2943.14),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "56379f3f-b27e-421f-8626-57b36fc6b9f5",
                                Name: codataccounting.String("Mathew Schmitt"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "25c67641-a831-42e5-847b-4c21ccb423ab",
                                Name: codataccounting.String("Josh Rutherford DVM"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("fuga"),
                            ID: "abdd88e7-1f6c-4482-92d7-771e7fd07400",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f8d29de1-dd70-497b-9da0-8c57fa6c78a2",
                            Name: codataccounting.String("Ms. Eileen Thompson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "feca6191-4981-440b-a4ff-8ae170ef03b5",
                            Name: codataccounting.String("Mike Kovacek"),
                        },
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
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5392c11a-25a8-4bf9-af97-428ad9a9f8bf"),
                        Name: codataccounting.String("Mr. Roy Conn"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(2449.9),
                    DiscountPercentage: codataccounting.Float64(3556.85),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9d98387f-7a79-4cd7-acd2-484da21729f2",
                        Name: codataccounting.String("Dr. Guillermo Funk"),
                    },
                    Quantity: 3495.58,
                    SubTotal: codataccounting.Float64(4682.21),
                    TaxAmount: codataccounting.Float64(1547.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3361.23),
                        ID: codataccounting.String("f1169ac1-e41d-48a2-bc23-e34f2dfa4a19"),
                        Name: codataccounting.String("Elsa Kerluke"),
                    },
                    TotalAmount: codataccounting.Float64(6107.22),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2151fe17-1209-4985-be9f-543d854439ee",
                                Name: codataccounting.String("Nicole Goyette"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("accusantium"),
                            ID: "443bc154-188c-42f5-ae85-da7832eabd61",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3b0d51a4-4bf0-41ba-9870-6d46082bfbdc",
                            Name: codataccounting.String("Jean Windler"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4e2ae4fb-5cb3-45d1-b638-f1edb78359ec",
                            Name: codataccounting.String("Warren Runolfsson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "60f8cd58-0ba7-4381-8e4f-e4447297cd3b",
                            Name: codataccounting.String("Angel Stokes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "bce247b7-684e-4ff5-8126-d71cffbd0eb7",
                            Name: codataccounting.String("Beulah Kuvalis"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1953b44b-d3c4-4315-9d33-e5953c001139",
                            Name: codataccounting.String("Gilbert Dickinson"),
                        },
                    },
                    UnitAmount: 2664.61,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1e6c31cc-2f1f-4cb5-9c9a-41ffbe9cbd79"),
                        Name: codataccounting.String("Casey Weber"),
                    },
                    Description: codataccounting.String("debitis"),
                    DiscountAmount: codataccounting.Float64(414.36),
                    DiscountPercentage: codataccounting.Float64(4582.74),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "6cc7abf6-16ea-45c7-9641-934b90f2e09d",
                        Name: codataccounting.String("Faye Senger"),
                    },
                    Quantity: 7687.72,
                    SubTotal: codataccounting.Float64(1277.84),
                    TaxAmount: codataccounting.Float64(9668.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6189.9),
                        ID: codataccounting.String("e2e10594-4b93-45d2-b7a7-2f90849d6aed"),
                        Name: codataccounting.String("Amelia Upton"),
                    },
                    TotalAmount: codataccounting.Float64(4929.22),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "37cd9222-c9ff-4574-91aa-bfa2e761f0ca",
                                Name: codataccounting.String("Meredith Green"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ef1031e6-899f-40c2-801e-22cd55cc0584",
                                Name: codataccounting.String("Carl Larson"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("nihil"),
                            ID: "6d971fc8-20c6-45b0-b7bb-8e0cc885187e",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e04af28c-5ddd-4b46-aa1c-fd6d828da013",
                            Name: codataccounting.String("Mr. Sabrina Blick"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "46645c1d-81f2-4904-af56-9b7aff0ea221",
                            Name: codataccounting.String("Nichole Renner IV"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1bc163e2-79a3-4b08-8da9-9257d04f4084",
                            Name: codataccounting.String("Kayla Kilback"),
                        },
                    },
                    UnitAmount: 8310.89,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("ut"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8124.31),
                        TotalAmount: codataccounting.Float64(7335.71),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("deecf6b9-9bc6-4356-aebf-df55c294c060"),
                            Name: codataccounting.String("Steven Jast Sr."),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4767.7),
                        ID: codataccounting.String("764eef6d-0c6d-46ed-9c73-dd634571509a"),
                        Note: codataccounting.String("blanditiis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("deleniti"),
                        TotalAmount: codataccounting.Float64(4930.17),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2381.67),
                        TotalAmount: codataccounting.Float64(7937.52),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5a1f9c24-2c7b-466a-9f30-c73df5b67198"),
                            Name: codataccounting.String("Gary Wisozk"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2723.96),
                        ID: codataccounting.String("bb438d85-b260-4591-9745-e3c2059c9c3f"),
                        Note: codataccounting.String("veniam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("esse"),
                        TotalAmount: codataccounting.Float64(8871.67),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("e252765b-1d62-4fcd-ace1-f01216ce2239"),
                    PurchaseOrderNumber: codataccounting.String("accusamus"),
                },
            },
            Reference: codataccounting.String("deleniti"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusUnknown,
            SubTotal: 3184.71,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "quibusdam": map[string]interface{}{
                        "nulla": "inventore",
                    },
                    "omnis": map[string]interface{}{
                        "excepturi": "nostrum",
                        "sint": "doloribus",
                        "magnam": "adipisci",
                        "natus": "necessitatibus",
                    },
                    "velit": map[string]interface{}{
                        "eos": "nisi",
                        "commodi": "impedit",
                        "facilis": "temporibus",
                    },
                    "error": map[string]interface{}{
                        "delectus": "molestiae",
                        "deserunt": "laborum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "2b241136-95d1-4e66-98fc-c4596217c297",
                SupplierName: codataccounting.String("molestiae"),
            },
            TaxAmount: 3968.2,
            TotalAmount: 4547.05,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 2325.57,
                    Name: "Rhonda Collins",
                },
                shared.BillWithholdingTax{
                    Amount: 17.82,
                    Name: "Penny Rice",
                },
            },
        },
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(584483),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
            Content: []byte("molestiae"),
            RequestBody: "et",
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

