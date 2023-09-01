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
            AmountDue: codataccounting.Float64(3759.94),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(2420.99),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("a5acfbe2-fd57-4075-b792-9177deac646e"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("573409e3-eb1e-45a2-b12e-b07f116db995"),
                        Name: codataccounting.String("Bernice Yundt"),
                    },
                    Description: codataccounting.String("enim"),
                    DiscountAmount: codataccounting.Float64(9449.5),
                    DiscountPercentage: codataccounting.Float64(6573.19),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "88970e18-9dbb-430f-8b33-ea055b197cd4",
                        Name: codataccounting.String("Kellie Corwin"),
                    },
                    Quantity: 1645.32,
                    SubTotal: codataccounting.Float64(8138.8),
                    TaxAmount: codataccounting.Float64(5129.05),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1403.84),
                        ID: codataccounting.String("d3513bb6-f48b-4656-bcdb-35ff2e4b2753"),
                        Name: codataccounting.String("Genevieve Lebsack"),
                    },
                    TotalAmount: codataccounting.Float64(6040.78),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7319c177-d525-4f77-b114-eeb52ff785fc",
                                Name: codataccounting.String("Mrs. Claudia Leuschke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4c98e0c2-bb89-4eb7-9dad-636c600503d8",
                                Name: codataccounting.String("Mr. Jonathon Fay"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0f739ae9-e057-4eb8-89e2-810331f3981d",
                                Name: codataccounting.String("Mr. Bethany Koch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "607f3c93-c73b-49da-bf2c-eda7e23f2257",
                                Name: codataccounting.String("Virginia Bins"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("delectus"),
                            ID: "4b7544e4-72e8-4028-97a5-b40463a7d575",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "400e764a-d733-44ec-9b78-1b36a08088d1",
                            Name: codataccounting.String("Jessica Turner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a200ef04-22eb-4216-8cf9-ab8366c723ff",
                            Name: codataccounting.String("Cameron Mosciski III"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "bee4825c-1fc0-4e11-9c80-bff918544ec4",
                            Name: codataccounting.String("Nadine Terry"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ce8f1977-773e-4635-a2a7-b408f05e3d48",
                            Name: codataccounting.String("Clint Ortiz"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "13a1f5fd-9425-49c0-b36f-25ea944f3b75",
                            Name: codataccounting.String("Dr. Alexandra Bernhard"),
                        },
                    },
                    UnitAmount: 7869.54,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("37a51262-4383-45bb-805a-23a45cefc5fd"),
                        Name: codataccounting.String("Juan Abshire DDS"),
                    },
                    Description: codataccounting.String("necessitatibus"),
                    DiscountAmount: codataccounting.Float64(1559.78),
                    DiscountPercentage: codataccounting.Float64(1189.32),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "69e51001-9c6d-4c5e-b476-2799bfbbe694",
                        Name: codataccounting.String("Irvin Rippin"),
                    },
                    Quantity: 7202.66,
                    SubTotal: codataccounting.Float64(2791.72),
                    TaxAmount: codataccounting.Float64(9253.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7972.54),
                        ID: codataccounting.String("ae6c3d5d-b3ad-4ebd-9dae-a4c506a8aa94"),
                        Name: codataccounting.String("Thomas Conroy"),
                    },
                    TotalAmount: codataccounting.Float64(3085.28),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f5e9d9a4-578a-4dc1-ac60-0dec001ac802",
                                Name: codataccounting.String("Louis Treutel V"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ff8f0f81-6ff3-4477-813e-902c14125b09",
                                Name: codataccounting.String("Carol O'Reilly"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8151a472-af92-43c5-949f-83f350cf876f",
                                Name: codataccounting.String("Mr. Robin Miller"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6ecbb4e2-43cf-4789-bfaf-eda53e5ae6e0",
                                Name: codataccounting.String("Myron Boyle"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quisquam"),
                            ID: "2b9c247c-8837-43a4-8e19-42f32e550557",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f5d56d0b-d0af-42df-a13d-b4f62cba3f89",
                            Name: codataccounting.String("Joyce O'Connell"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0b80a692-4d3b-42ec-bcc8-f895010f5dd3",
                            Name: codataccounting.String("Chester Willms V"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "04e54c82-f168-4a36-bc88-73e484380b1f",
                            Name: codataccounting.String("Yvette Larson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "275a60a0-4c49-45cc-a991-71b51c1bdb1c",
                            Name: codataccounting.String("Leroy Ratke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8ebdfc4c-cca9-49bc-bfc0-b2dce10873e4",
                            Name: codataccounting.String("Ms. Susie Batz"),
                        },
                    },
                    UnitAmount: 4312.53,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("78878ba8-581a-4582-88c5-4fefa9c95f2e"),
                        Name: codataccounting.String("Noel Hauck"),
                    },
                    Description: codataccounting.String("nemo"),
                    DiscountAmount: codataccounting.Float64(8493.37),
                    DiscountPercentage: codataccounting.Float64(2012.66),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "07cfee81-206e-4281-bfa4-a41c480d3f21",
                        Name: codataccounting.String("Theresa Pfannerstill I"),
                    },
                    Quantity: 1018.54,
                    SubTotal: codataccounting.Float64(449.29),
                    TaxAmount: codataccounting.Float64(1341.73),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8603.62),
                        ID: codataccounting.String("514f4cc6-f18b-4f96-a1a6-a4f77a87ee3e"),
                        Name: codataccounting.String("Susie Ward"),
                    },
                    TotalAmount: codataccounting.Float64(1316.87),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "65b34418-e3bb-491c-8d97-5e0e8419d8f8",
                                Name: codataccounting.String("Lila Bradtke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f3e07edc-c4aa-45f3-8abd-905a972e0567",
                                Name: codataccounting.String("Myrtle Cremin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b2d30947-0bf7-4a4f-a87c-f535a6fae54e",
                                Name: codataccounting.String("Miss Cary Howe"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "21f023b7-5d23-467f-a1a0-cc8df79f0a39",
                                Name: codataccounting.String("Miss Estelle Mills"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("commodi"),
                            ID: "4b7c15df-bace-4188-b1c4-ee2c8c6ce611",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "eb1c7cbd-b6ee-4c74-b78b-a25317747dc9",
                            Name: codataccounting.String("Annette Osinski"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "af5dd672-3dc0-4f5a-a2f3-a6b700878756",
                            Name: codataccounting.String("Gail Fay"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a6c98b55-5540-480d-80bc-acc6cbd6b5f3",
                            Name: codataccounting.String("Ms. Wilbert McGlynn"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "04f926ba-d255-4381-9b47-4b0ed20e5624",
                            Name: codataccounting.String("Moses Wuckert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "39a910ab-dcab-4626-b669-6e1ec00221b3",
                            Name: codataccounting.String("Yvonne Stamm"),
                        },
                    },
                    UnitAmount: 6294.61,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("tempore"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(9890.89),
                        TotalAmount: codataccounting.Float64(8360.53),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a8d0c549-ef03-4004-978a-61fa1cf20688"),
                            Name: codataccounting.String("Jared Koepp DVM"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7989.53),
                        ID: codataccounting.String("71dca163-f2a3-4c80-a97f-f334cddf857a"),
                        Note: codataccounting.String("perspiciatis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("eum"),
                        TotalAmount: codataccounting.Float64(951.23),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("76c6ab21-d29d-4fc9-8d6f-ecd799390066"),
                    PurchaseOrderNumber: codataccounting.String("laborum"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("6d2d0003-5533-48ce-8086-fa21e9152cb3"),
                    PurchaseOrderNumber: codataccounting.String("beatae"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("19167b8e-3c8d-4b03-808d-6d364ffd4559"),
                    PurchaseOrderNumber: codataccounting.String("alias"),
                },
            },
            Reference: codataccounting.String("ex"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusUnknown,
            SubTotal: 1548.4,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "neque": map[string]interface{}{
                        "numquam": "rem",
                        "officiis": "omnis",
                        "neque": "corporis",
                        "quod": "dolores",
                    },
                    "placeat": map[string]interface{}{
                        "recusandae": "quos",
                        "dicta": "sapiente",
                        "ipsum": "consequatur",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "be3e4320-2d72-4165-b650-6641870d9d21",
                SupplierName: codataccounting.String("voluptatibus"),
            },
            TaxAmount: 6012.28,
            TotalAmount: 6456.09,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 79.19,
                    Name: "Sharon Ruecker",
                },
                shared.BillWithholdingTax{
                    Amount: 7639.37,
                    Name: "Carl Breitenberg V",
                },
                shared.BillWithholdingTax{
                    Amount: 1917.24,
                    Name: "Anita Dare III",
                },
                shared.BillWithholdingTax{
                    Amount: 5123.7,
                    Name: "Mr. Armando Hermann",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(312690),
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
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
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
        Query: codataccounting.String("esse"),
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
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
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
            AmountDue: codataccounting.Float64(6975.91),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(5062.45),
            DueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("5e320a31-9f4b-4adf-947c-9a867bc42426"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5816ddca-8ef5-41fc-b4c5-93ec12cdaad0"),
                        Name: codataccounting.String("Clark Kohler"),
                    },
                    Description: codataccounting.String("saepe"),
                    DiscountAmount: codataccounting.Float64(8139.75),
                    DiscountPercentage: codataccounting.Float64(7487.23),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "d80df448-a47f-4939-8c58-880983dabf9e",
                        Name: codataccounting.String("Jeffery Williamson"),
                    },
                    Quantity: 8301.49,
                    SubTotal: codataccounting.Float64(6077.42),
                    TaxAmount: codataccounting.Float64(9666.52),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4877.65),
                        ID: codataccounting.String("f079af4d-3572-44cd-b0f4-d281187d5684"),
                        Name: codataccounting.String("Eloise Stoltenberg"),
                    },
                    TotalAmount: codataccounting.Float64(5057.99),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "a9065e62-8bdf-4c20-b2b6-c879923b7e13",
                                Name: codataccounting.String("Leah Graham"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ae12c689-1f82-4ce1-9571-72305377dcfa",
                                Name: codataccounting.String("Terrance Strosin"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quam"),
                            ID: "5e356686-092e-49c3-9dc5-f111dea1026d",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1a4d190f-eb21-4780-bccc-0dbbddb48470",
                            Name: codataccounting.String("Dominick Purdy"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "91e6bc15-8c4c-44e5-8599-ea342260e9b2",
                            Name: codataccounting.String("Elizabeth Rutherford"),
                        },
                    },
                    UnitAmount: 5371.4,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a1bd8fb7-a0a1-416c-a723-d4097fa30e9a"),
                        Name: codataccounting.String("Kurt Cronin"),
                    },
                    Description: codataccounting.String("quia"),
                    DiscountAmount: codataccounting.Float64(6090.94),
                    DiscountPercentage: codataccounting.Float64(1206.46),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "22030d83-f5ae-4b77-99d2-2e8c1f849382",
                        Name: codataccounting.String("Marta Stanton"),
                    },
                    Quantity: 1663.24,
                    SubTotal: codataccounting.Float64(7639.28),
                    TaxAmount: codataccounting.Float64(5526.87),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4484.82),
                        ID: codataccounting.String("6c2c2dfb-4cfc-41c7-a230-f841fb1bd23f"),
                        Name: codataccounting.String("Alton Bernhard"),
                    },
                    TotalAmount: codataccounting.Float64(7092.34),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "be5a6859-98e2-42ae-a0da-16fc2b271a28",
                                Name: codataccounting.String("Clark Hermiston"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "854e9043-9d22-4246-9694-62407084f7ab",
                                Name: codataccounting.String("Nellie Ruecker"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("accusantium"),
                            ID: "2225194d-b554-410a-9c66-9af90a26c7cd",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "81f06898-1d6b-4b33-8faa-348c31bf407e",
                            Name: codataccounting.String("Francis Yundt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c42b78f1-5626-4398-a0dc-766324ccb06c",
                            Name: codataccounting.String("Mr. Benny O'Reilly"),
                        },
                    },
                    UnitAmount: 271.97,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("enim"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4866.06),
                        TotalAmount: codataccounting.Float64(27.58),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b8d5722d-d895-4b8b-8f24-db959693352f"),
                            Name: codataccounting.String("Joanne Hermiston"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(5812.69),
                        ID: codataccounting.String("4d78de3b-6e93-489f-9abb-7f662550a283"),
                        Note: codataccounting.String("totam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("deserunt"),
                        TotalAmount: codataccounting.Float64(7547.84),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("83afd231-5bba-4650-964e-06f5bf6ae591"),
                    PurchaseOrderNumber: codataccounting.String("expedita"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("c8bdef36-12b6-43c2-85fd-a840774a68a9"),
                    PurchaseOrderNumber: codataccounting.String("laborum"),
                },
            },
            Reference: codataccounting.String("dolor"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillStatusDraft,
            SubTotal: 229.66,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptas": map[string]interface{}{
                        "voluptas": "maiores",
                        "ea": "vel",
                        "delectus": "accusamus",
                    },
                    "reiciendis": map[string]interface{}{
                        "sed": "accusantium",
                    },
                    "voluptates": map[string]interface{}{
                        "maiores": "quaerat",
                        "numquam": "non",
                        "cum": "incidunt",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "257b992c-8dbd-4a6a-a1ef-a2198258fd0a",
                SupplierName: codataccounting.String("iste"),
            },
            TaxAmount: 9085.87,
            TotalAmount: 7236.23,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 2889.07,
                    Name: "Shari Konopelski",
                },
                shared.BillWithholdingTax{
                    Amount: 9156.53,
                    Name: "Thomas Hahn",
                },
                shared.BillWithholdingTax{
                    Amount: 2946.5,
                    Name: "Desiree Howell IV",
                },
            },
        },
        BillID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(114588),
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
            Content: []byte("quisquam"),
            RequestBody: "atque",
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

