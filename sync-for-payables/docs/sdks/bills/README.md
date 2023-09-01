# Bills

## Overview

Bills

### Available Operations

* [Create](#create) - Create bill
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
            AmountDue: codatsyncpayables.Float64(2986.13),
            Currency: codatsyncpayables.String("USD"),
            CurrencyRate: codatsyncpayables.Float64(4609.09),
            DueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            ID: codatsyncpayables.String("adc1ac60-0dec-4001-ac80-2e2ec09ff8f0"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("16ff3477-c13e-4902-8141-25b0960a6681"),
                        Name: codatsyncpayables.String("Amy Murray"),
                    },
                    Description: codatsyncpayables.String("magni"),
                    DiscountAmount: codatsyncpayables.Float64(6463.29),
                    DiscountPercentage: codatsyncpayables.Float64(9650.95),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "923c5949-f83f-4350-8f87-6ffb901c6ecb",
                        Name: codatsyncpayables.String("Joel Von"),
                    },
                    Quantity: 1940.58,
                    SubTotal: codatsyncpayables.Float64(7581.94),
                    TaxAmount: codatsyncpayables.Float64(9928.87),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(4598.75),
                        ID: codatsyncpayables.String("89ffafed-a53e-45ae-ae0a-c184c2b9c247"),
                        Name: codatsyncpayables.String("Isaac Lowe"),
                    },
                    TotalAmount: codatsyncpayables.Float64(2271.56),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "40e1942f-32e5-4505-9756-f5d56d0bd0af",
                                Name: codatsyncpayables.String("Elena Zieme I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "db4f62cb-a3f8-4941-aebc-0b80a6924d3b",
                                Name: codatsyncpayables.String("Eloise Rowe"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c8f89501-0f5d-4d3d-afa1-804e54c82f16",
                                Name: codatsyncpayables.String("Rex Ernser"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("cumque"),
                            ID: "8873e484-380b-41f6-b8ca-275a60a04c49",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "c699171b-51c1-4bdb-9cf4-b888ebdfc4cc",
                            Name: codatsyncpayables.String("Lynn Marks"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "7fc0b2dc-e108-473e-82b0-06d678878ba8",
                            Name: codatsyncpayables.String("Kay Bradtke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8208c54f-efa9-4c95-b2ea-c5565d307cfe",
                            Name: codatsyncpayables.String("Hugh Carroll III"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e2813fa4-a41c-4480-93f2-132af03102d5",
                            Name: codatsyncpayables.String("Danielle Willms"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c6f18bf9-621a-46a4-b77a-87ee3e4be752",
                            Name: codatsyncpayables.String("Gilbert Hayes"),
                        },
                    },
                    UnitAmount: 2848.85,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("418e3bb9-1c8d-4975-a0e8-419d8f84f144"),
                        Name: codatsyncpayables.String("Ms. Jimmy Turcotte"),
                    },
                    Description: codatsyncpayables.String("facere"),
                    DiscountAmount: codatsyncpayables.Float64(7890.16),
                    DiscountPercentage: codatsyncpayables.Float64(7690.47),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4aa5f3ca-bd90-45a9-b2e0-56728227b2d3",
                        Name: codatsyncpayables.String("Bobbie Greenfelder MD"),
                    },
                    Quantity: 9565.45,
                    SubTotal: codatsyncpayables.Float64(4630.5),
                    TaxAmount: codatsyncpayables.Float64(6671.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(3073.06),
                        ID: codatsyncpayables.String("fa87cf53-5a6f-4ae5-8ebf-60c321f023b7"),
                        Name: codatsyncpayables.String("Paulette Dibbert"),
                    },
                    TotalAmount: codatsyncpayables.Float64(4939.45),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e1a0cc8d-f79f-40a3-96d9-0c364b7c15df",
                                Name: codatsyncpayables.String("Hubert Russel V"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8b1c4ee2-c8c6-4ce6-91fe-eb1c7cbdb6ee",
                                Name: codatsyncpayables.String("Jessie Hahn"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8ba25317-747d-4c91-9ad2-caf5dd6723dc",
                                Name: codatsyncpayables.String("Shawna Heller"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2f3a6b70-0878-4756-943f-5a6c98b55554",
                                Name: codatsyncpayables.String("Naomi Bauch"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("doloremque"),
                            ID: "bcacc6cb-d6b5-4f3e-8909-304f926bad25",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "819b474b-0ed2-40e5-a248-fff639a910ab",
                            Name: codatsyncpayables.String("Salvatore Paucek"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "676696e1-ec00-4221-b335-d89acb3ecfda",
                            Name: codatsyncpayables.String("Ellis Balistreri"),
                        },
                    },
                    UnitAmount: 3075.32,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("9ef03004-978a-461f-a1cf-20688f77c1ff"),
                        Name: codatsyncpayables.String("Jared Blick"),
                    },
                    Description: codatsyncpayables.String("fuga"),
                    DiscountAmount: codatsyncpayables.Float64(665.96),
                    DiscountPercentage: codatsyncpayables.Float64(4057.89),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3f2a3c80-a97f-4f33-8cdd-f857a9e61876",
                        Name: codatsyncpayables.String("Marc O'Connell Sr."),
                    },
                    Quantity: 8600.27,
                    SubTotal: codatsyncpayables.Float64(1665.42),
                    TaxAmount: codatsyncpayables.Float64(6180.73),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(8511.99),
                        ID: codatsyncpayables.String("fc94d6fe-cd79-4939-8066-a6d2d0003553"),
                        Name: codatsyncpayables.String("Billie Schmitt"),
                    },
                    TotalAmount: codatsyncpayables.Float64(390.47),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "6fa21e91-52cb-4311-9167-b8e3c8db0340",
                                Name: codatsyncpayables.String("Levi Johns"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "64ffd455-906d-4126-bd48-e935c2c9e81f",
                                Name: codatsyncpayables.String("Elizabeth Roberts"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e43202d7-2165-4765-8664-1870d9d21f9a",
                                Name: codatsyncpayables.String("Miss Michael Ferry"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("debitis"),
                            ID: "cc11a083-6429-4068-b850-2a55e7f73bc8",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "e320a319-f4ba-4df9-87c9-a867bc424266",
                            Name: codatsyncpayables.String("Mrs. Geraldine Lueilwitz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ca8ef51f-cb4c-4593-ac12-cdaad0ec7afe",
                            Name: codatsyncpayables.String("Robin Strosin PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f448a47f-9390-4c58-8809-83dabf9ef3ff",
                            Name: codatsyncpayables.String("Levi Mohr"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f079af4d-3572-44cd-b0f4-d281187d5684",
                            Name: codatsyncpayables.String("Eloise Stoltenberg"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "85a9065e-628b-4dfc-a032-b6c879923b7e",
                            Name: codatsyncpayables.String("Rosa Hand"),
                        },
                    },
                    UnitAmount: 9897.65,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("deserunt"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(7733.55),
                        TotalAmount: codatsyncpayables.Float64(4013.88),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("891f82ce-1157-4172-b053-77dcfa89df97"),
                            Name: codatsyncpayables.String("Tasha Dickinson"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(5228.24),
                        ID: codatsyncpayables.String("6092e9c3-ddc5-4f11-9dea-1026d541a4d1"),
                        Note: codatsyncpayables.String("omnis"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("sapiente"),
                        TotalAmount: codatsyncpayables.Float64(8876),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(1127.51),
                        TotalAmount: codatsyncpayables.Float64(4878.39),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("80bccc0d-bbdd-4b48-8708-fb4e391e6bc1"),
                            Name: codatsyncpayables.String("Vickie Sauer"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(9365.18),
                        ID: codatsyncpayables.String("54599ea3-4226-40e9-b200-ce78a1bd8fb7"),
                        Note: codatsyncpayables.String("culpa"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("fuga"),
                        TotalAmount: codatsyncpayables.Float64(1175.46),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(7574.38),
                        TotalAmount: codatsyncpayables.Float64(8855.23),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("723d4097-fa30-4e9a-b725-b29122030d83"),
                            Name: codatsyncpayables.String("Dan Nolan"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(4938.65),
                        ID: codatsyncpayables.String("99d22e8c-1f84-4938-a5fd-c42c876c2c2d"),
                        Note: codatsyncpayables.String("delectus"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("aliquam"),
                        TotalAmount: codatsyncpayables.Float64(7579.62),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(756.1),
                        TotalAmount: codatsyncpayables.Float64(7513.47),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("76230f84-1fb1-4bd2-bfdb-14db6be5a685"),
                            Name: codatsyncpayables.String("Luther Leuschke"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(6715.68),
                        ID: codatsyncpayables.String("e20da16f-c2b2-471a-a89c-57e854e90439"),
                        Note: codatsyncpayables.String("quibusdam"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("consequuntur"),
                        TotalAmount: codatsyncpayables.Float64(1594.69),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("65694624-0708-44f7-ab37-cef02225194d"),
                    PurchaseOrderNumber: codatsyncpayables.String("quidem"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("55410adc-669a-4f90-a26c-7cdc981f0689"),
                    PurchaseOrderNumber: codatsyncpayables.String("corrupti"),
                },
            },
            Reference: codatsyncpayables.String("quae"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusPartiallyPaid,
            SubTotal: 7304.37,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "amet": map[string]interface{}{
                        "minus": "hic",
                    },
                    "similique": map[string]interface{}{
                        "consectetur": "labore",
                        "laudantium": "cumque",
                        "adipisci": "veritatis",
                    },
                    "nam": map[string]interface{}{
                        "magnam": "aperiam",
                        "ducimus": "itaque",
                        "necessitatibus": "numquam",
                        "doloribus": "eligendi",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "f0c42b78-f156-4263-98a0-dc766324ccb0",
                SupplierName: codatsyncpayables.String("ea"),
            },
            TaxAmount: 7713.21,
            TotalAmount: 5184.32,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 6651.83,
                    Name: "Mr. Louise Sipes",
                },
                shared.BillWithholdingTax{
                    Amount: 1364.32,
                    Name: "Miss Victor Kuhlman",
                },
                shared.BillWithholdingTax{
                    Amount: 8196.9,
                    Name: "Courtney Conroy",
                },
                shared.BillWithholdingTax{
                    Amount: 8506.28,
                    Name: "Kirk Heidenreich",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(704665),
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
        Query: codatsyncpayables.String("dolore"),
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
            AmountDue: codatsyncpayables.Float64(7053.07),
            Currency: codatsyncpayables.String("USD"),
            CurrencyRate: codatsyncpayables.Float64(3656.76),
            DueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            ID: codatsyncpayables.String("693352f7-4533-4994-978d-e3b6e9389f5a"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("7f662550-a283-482a-8483-afd2315bba65"),
                        Name: codatsyncpayables.String("Debra Kerluke"),
                    },
                    Description: codatsyncpayables.String("quae"),
                    DiscountAmount: codatsyncpayables.Float64(4090.21),
                    DiscountPercentage: codatsyncpayables.Float64(9891.22),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5bf6ae59-1bc8-4bde-b361-2b63c205fda8",
                        Name: codatsyncpayables.String("Sarah Kuhn"),
                    },
                    Quantity: 6839.8,
                    SubTotal: codatsyncpayables.Float64(4359.31),
                    TaxAmount: codatsyncpayables.Float64(5383.68),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(6382.19),
                        ID: codatsyncpayables.String("9a35d086-b6f6-46fe-b020-e9f443b4257b"),
                        Name: codatsyncpayables.String("Marshall Daugherty"),
                    },
                    TotalAmount: codatsyncpayables.Float64(8301.97),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "da6a61ef-a219-4825-8fd0-a9eba47f7d3e",
                                Name: codatsyncpayables.String("Thomas Hahn"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "40d6a183-1c87-4adf-996f-df1ad837ae80",
                                Name: codatsyncpayables.String("Ms. Terry Runolfsson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95ba9986-78fa-43f6-9699-1af388ce0361",
                                Name: codatsyncpayables.String("Eva Gleichner"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("voluptate"),
                            ID: "977a0ef2-f536-4028-afee-f934152ed7e2",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "f4c157de-aa71-470f-845a-ccf667aaf9bb",
                            Name: codatsyncpayables.String("Laurence Blick"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e431d6bf-5c83-48fb-b8c2-0cb67fc4b425",
                            Name: codatsyncpayables.String("Perry Mayert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "234c9f7b-79df-4eb7-ba5c-38d4baf91e50",
                            Name: codatsyncpayables.String("Olive Windler"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0a54b475-f16f-456d-b85a-3c4ac631b99e",
                            Name: codatsyncpayables.String("Ella Runolfsdottir"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8f9fdb94-10f6-43bb-b817-837b01afdd78",
                            Name: codatsyncpayables.String("Chester Daugherty IV"),
                        },
                    },
                    UnitAmount: 6138.48,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("eb44873f-5033-4f19-9bf1-25ce4152eab9"),
                        Name: codatsyncpayables.String("Lionel Klocko"),
                    },
                    Description: codatsyncpayables.String("odit"),
                    DiscountAmount: codatsyncpayables.Float64(1383.06),
                    DiscountPercentage: codatsyncpayables.Float64(2593.74),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a6a0e123-b784-47ec-99e1-f67f3c4cce4b",
                        Name: codatsyncpayables.String("Angel Kris"),
                    },
                    Quantity: 4001.45,
                    SubTotal: codatsyncpayables.Float64(9961.01),
                    TaxAmount: codatsyncpayables.Float64(9570.32),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(2322.09),
                        ID: codatsyncpayables.String("c5747501-357e-444f-91f8-b084c3197e19"),
                        Name: codatsyncpayables.String("Rosie Conroy"),
                    },
                    TotalAmount: codatsyncpayables.Float64(3088.66),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7f94874c-2d5c-4c49-b223-3e66bd8fe5d0",
                                Name: codatsyncpayables.String("Jeannette Mante"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ef203873-2059-40cc-8109-6400313b3e50",
                                Name: codatsyncpayables.String("Debbie Windler"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("maiores"),
                            ID: "e72dc407-7d0c-4c3f-808e-fc15ceb4d6e1",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "e0f75aed-f2ac-4ab5-8b99-1c926ddb5894",
                            Name: codatsyncpayables.String("Joyce Terry"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1cbe6d95-02f0-4ea9-b0b6-9f7ac2f72f88",
                            Name: codatsyncpayables.String("Karen Barrows I"),
                        },
                    },
                    UnitAmount: 5776.22,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("11608207-888e-4c66-983b-fe9659eb40ec"),
                        Name: codatsyncpayables.String("Loretta Wehner"),
                    },
                    Description: codatsyncpayables.String("nihil"),
                    DiscountAmount: codatsyncpayables.Float64(3452.7),
                    DiscountPercentage: codatsyncpayables.Float64(7023.39),
                    IsDirectCost: codatsyncpayables.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "0b532a4d-a37c-4baa-b445-2c4842c9b2ad",
                        Name: codatsyncpayables.String("Rose Stoltenberg"),
                    },
                    Quantity: 8968.11,
                    SubTotal: codatsyncpayables.Float64(5316.06),
                    TaxAmount: codatsyncpayables.Float64(981.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(6493.73),
                        ID: codatsyncpayables.String("88f44445-73fe-4cd4-b353-f63c8209379a"),
                        Name: codatsyncpayables.String("Reginald McClure"),
                    },
                    TotalAmount: codatsyncpayables.Float64(3210.07),
                    Tracking: &shared.Tracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "bcf79da1-8a78-422b-b958-94e6861adb55",
                                Name: codatsyncpayables.String("Nick Torp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "751c9fe8-f750-42bf-9c34-50841f176445",
                                Name: codatsyncpayables.String("Wendy Kling"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3fb27e21-f862-4657-b36f-c6b9f587ce52",
                                Name: codatsyncpayables.String("Brooke Jacobs"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "41a8312e-5047-4b4c-a1cc-b423abcdc91f",
                                Name: codatsyncpayables.String("Rex Rau"),
                            },
                        },
                        CustomerRef: &shared.TrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("totam"),
                            ID: "8e71f6c4-8252-4d77-b1e7-fd074009ef8d",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.TrackingProjectRef{
                            ID: "de1dd709-7b5d-4a08-857f-a6c78a216e19",
                            Name: codatsyncpayables.String("Pablo Wilkinson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "61914981-40b6-44ff-8ae1-70ef03b5f37e",
                            Name: codatsyncpayables.String("Angie O'Hara"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "85559667-32aa-45dc-b668-2cb70f8cfd5f",
                            Name: codatsyncpayables.String("Brent Weber MD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9a9f7484-6e2c-4330-9db0-536d9e75ca00",
                            Name: codatsyncpayables.String("Lana Hauck"),
                        },
                    },
                    UnitAmount: 1500.91,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("et"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(3322.37),
                        TotalAmount: codatsyncpayables.Float64(6650.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("8bf92f97-428a-4d9a-9f8b-f8221125359d"),
                            Name: codatsyncpayables.String("Guy Feest"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(4820.62),
                        ID: codatsyncpayables.String("a79cd72c-d248-44da-a172-9f2ac41ef572"),
                        Note: codatsyncpayables.String("quis"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("dicta"),
                        TotalAmount: codatsyncpayables.Float64(1158.49),
                    },
                },
            },
            PurchaseOrderRefs: []shared.BillPurchaseOrderReference{
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("9ac1e41d-8a23-4c23-a34f-2dfa4a197f6d"),
                    PurchaseOrderNumber: codatsyncpayables.String("recusandae"),
                },
                shared.BillPurchaseOrderReference{
                    ID: codatsyncpayables.String("922151fe-1712-4099-853e-9f543d854439"),
                    PurchaseOrderNumber: codatsyncpayables.String("accusamus"),
                },
            },
            Reference: codatsyncpayables.String("voluptates"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusUnknown,
            SubTotal: 2805.9,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "laboriosam": map[string]interface{}{
                        "tempora": "aliquam",
                    },
                    "dolorem": map[string]interface{}{
                        "impedit": "architecto",
                        "minima": "magnam",
                        "vitae": "quos",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "8c2f56e8-5da7-4832-aabd-617c3b0d51a4",
                SupplierName: codatsyncpayables.String("ut"),
            },
            TaxAmount: 6942.92,
            TotalAmount: 9852.65,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 1013.18,
                    Name: "Doug Stiedemann",
                },
            },
        },
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        TimeoutInMinutes: codatsyncpayables.Int(434330),
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
            Content: []byte("temporibus"),
            RequestBody: "incidunt",
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

