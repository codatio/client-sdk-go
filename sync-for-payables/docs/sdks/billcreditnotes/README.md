# BillCreditNotes

## Overview

Bill credit notes

### Available Operations

* [Create](#create) - Create bill credit note
* [Get](#get) - Get bill credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bill credit note model
* [List](#list) - List bill credit notes
* [Update](#update) - Update bill credit note

## Create

The *Create bill credit note* endpoint creates a new [bill credit note](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-billCreditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating a bill credit note.


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
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codatsyncpayables.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codatsyncpayables.String("GBP"),
            CurrencyRate: codatsyncpayables.Float64(8700.13),
            DiscountPercentage: 0,
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("7cc78ca1-ba92-48fc-8167-42cb73920592"),
                        Name: codatsyncpayables.String("Curtis Morissette"),
                    },
                    Description: codatsyncpayables.String("saepe"),
                    DiscountAmount: codatsyncpayables.Float64(6818.2),
                    DiscountPercentage: codatsyncpayables.Float64(4499.5),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "596eb10f-aaa2-4352-8595-5907aff1a3a2",
                        Name: codatsyncpayables.String("Shaun McCullough"),
                    },
                    Quantity: 4663.11,
                    SubTotal: codatsyncpayables.Float64(4746.97),
                    TaxAmount: codatsyncpayables.Float64(2444.25),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(6235.1),
                        ID: codatsyncpayables.String("251aa52c-3f5a-4d01-9da1-ffe78f097b00"),
                        Name: codatsyncpayables.String("Mrs. April Wuckert"),
                    },
                    TotalAmount: codatsyncpayables.Float64(4808.94),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b5e6e13b-99d4-488e-9e91-e450ad2abd44",
                                Name: codatsyncpayables.String("Beth McGlynn Sr."),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("assumenda"),
                            ID: "502a94bb-4f63-4c96-9e9a-3efa77dfb14c",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "6ae395ef-b9ba-488f-ba66-997074ba4469",
                            Name: codatsyncpayables.String("Duane Thiel II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "959890af-a563-4e25-96fe-4c8b711e5b7f",
                            Name: codatsyncpayables.String("Louis Turner Sr."),
                        },
                    },
                    UnitAmount: 5083.15,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("921cddc6-9260-41fb-976b-0d5f0d30c5fb"),
                        Name: codatsyncpayables.String("Ernest Hayes"),
                    },
                    Description: codatsyncpayables.String("eaque"),
                    DiscountAmount: codatsyncpayables.Float64(3389.85),
                    DiscountPercentage: codatsyncpayables.Float64(1999.96),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "202c73d5-fe9b-490c-a890-9b3fe49a8d9c",
                        Name: codatsyncpayables.String("Toby Hahn"),
                    },
                    Quantity: 2123.9,
                    SubTotal: codatsyncpayables.Float64(2098.43),
                    TaxAmount: codatsyncpayables.Float64(2224.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(1861.93),
                        ID: codatsyncpayables.String("3f9b77f3-a410-4067-8ebf-69280d1ba77a"),
                        Name: codatsyncpayables.String("Arturo Treutel"),
                    },
                    TotalAmount: codatsyncpayables.Float64(4694.97),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "7ae4203c-e5e6-4a95-98a0-d446ce2af7a7",
                                Name: codatsyncpayables.String("Rosalie White"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("accusamus"),
                            ID: "453f870b-326b-45a7-b429-cdb1a8422bb6",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "d2322715-bf0c-4bb1-a31b-8b90f3443a11",
                            Name: codatsyncpayables.String("Miss Billie Ward"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f4b92187-9fce-4953-b73e-f7fbc7abd74d",
                            Name: codatsyncpayables.String("Earl Mosciski DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5d2cff7c-70a4-4562-ad43-6813f16d9f5f",
                            Name: codatsyncpayables.String("Elbert Jenkins"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "56146c3e-250f-4b00-8c42-e141aac366c8",
                            Name: codatsyncpayables.String("Drew Hoeger I"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "42907474-778a-47bd-866d-28c10ab3cdca",
                            Name: codatsyncpayables.String("Ms. Ruby Hintz II"),
                        },
                    },
                    UnitAmount: 8920.5,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("523c7e0b-c717-48e4-b96f-2a70c688282a"),
                        Name: codatsyncpayables.String("Randall Lindgren"),
                    },
                    Description: codatsyncpayables.String("nisi"),
                    DiscountAmount: codatsyncpayables.Float64(1470.14),
                    DiscountPercentage: codatsyncpayables.Float64(9564.06),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "222e9817-ee17-4cbe-a1e6-b7b95bc0ab3c",
                        Name: codatsyncpayables.String("Elizabeth Schinner"),
                    },
                    Quantity: 2328.65,
                    SubTotal: codatsyncpayables.Float64(4581.39),
                    TaxAmount: codatsyncpayables.Float64(5034.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(5909.84),
                        ID: codatsyncpayables.String("fd871f99-dd2e-4fd1-a1aa-6f1e674bdb04"),
                        Name: codatsyncpayables.String("Samuel Hermiston"),
                    },
                    TotalAmount: codatsyncpayables.Float64(3917.74),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "82d68ea1-9f1d-4170-9133-9d08086a1840",
                                Name: codatsyncpayables.String("Toni Fritsch"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("voluptas"),
                            ID: "071f93f5-f064-42da-87af-515cc413aa63",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "e8d67864-dbb6-475f-95e6-0b375ed4f6fb",
                            Name: codatsyncpayables.String("Dr. Terence Gulgowski"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "317fe35b-60eb-41ea-8265-55ba3c28744e",
                            Name: codatsyncpayables.String("Dustin Ferry"),
                        },
                    },
                    UnitAmount: 5553.61,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("f3a8d8f5-c0b2-4f2f-b7b1-94a276b26916"),
                        Name: codatsyncpayables.String("Rogelio Bins V"),
                    },
                    Description: codatsyncpayables.String("maiores"),
                    DiscountAmount: codatsyncpayables.Float64(2748.23),
                    DiscountPercentage: codatsyncpayables.Float64(1484.78),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "94e3698f-447f-4603-a8b4-45e80ca55efd",
                        Name: codatsyncpayables.String("Deborah Turcotte"),
                    },
                    Quantity: 4461.35,
                    SubTotal: codatsyncpayables.Float64(8892.34),
                    TaxAmount: codatsyncpayables.Float64(1046.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(5124.52),
                        ID: codatsyncpayables.String("58b6a89f-be3a-45aa-8e48-24d0ab407508"),
                        Name: codatsyncpayables.String("Ms. Lamar Hessel"),
                    },
                    TotalAmount: codatsyncpayables.Float64(1536.27),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "65e904f3-b119-44b8-abf6-03a79f9dfe0a",
                                Name: codatsyncpayables.String("Ron Schulist"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("mollitia"),
                            ID: "50ce187f-86bc-4173-9689-eee9526f8d98",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "881ead4f-0e10-4125-a3f9-4e29e973e922",
                            Name: codatsyncpayables.String("Leo Kiehn II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e3e06080-7e2b-46e3-ab88-45f0597a60ff",
                            Name: codatsyncpayables.String("Alberta Harber"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "31e94764-a3e8-465e-b956-f9251a5a9da6",
                            Name: codatsyncpayables.String("Ruth Zulauf"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7bfaad4f-9efc-41b4-912c-1032648dc2f6",
                            Name: codatsyncpayables.String("Cathy Breitenberg"),
                        },
                    },
                    UnitAmount: 9364.69,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(9358.33),
                        TotalAmount: codatsyncpayables.Float64(5962.11),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("fe6c632c-a3ae-4d01-9799-6312fde04771"),
                            Name: codatsyncpayables.String("Tamara Lang"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(999.58),
                        ID: codatsyncpayables.String("d0174763-60a1-45db-aa66-0659a1adeaab"),
                        Note: codatsyncpayables.String("ad"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("enim"),
                        TotalAmount: codatsyncpayables.Float64(1104.77),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(7758.03),
                        TotalAmount: codatsyncpayables.Float64(4053.73),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("45b08b61-891b-4aa0-be1a-de008e6f8c5f"),
                            Name: codatsyncpayables.String("Marion Aufderhar"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(8427.77),
                        ID: codatsyncpayables.String("b5a34181-4301-4042-9813-d5208ece7e25"),
                        Note: codatsyncpayables.String("nesciunt"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("eum"),
                        TotalAmount: codatsyncpayables.Float64(4269.43),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(3494.4),
                        TotalAmount: codatsyncpayables.Float64(704.1),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("c6c6e205-e16d-4eab-bfec-9578a6458427"),
                            Name: codatsyncpayables.String("Lee Larkin IV"),
                        },
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(1173.8),
                        ID: codatsyncpayables.String("62309fb0-9299-421a-afb9-f58c4d86e68e"),
                        Note: codatsyncpayables.String("modi"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("vero"),
                        TotalAmount: codatsyncpayables.Float64(329.01),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(13.83),
                        TotalAmount: codatsyncpayables.Float64(938.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("3f59da75-7a59-4ecf-af66-ef1caa3383c2"),
                            Name: codatsyncpayables.String("Lamar Reichert"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: codatsyncpayables.Float64(1940.23),
                        ID: codatsyncpayables.String("73c8d72f-64d1-4db1-b2c4-310661e96349"),
                        Note: codatsyncpayables.String("earum"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("impedit"),
                        TotalAmount: codatsyncpayables.Float64(9758.84),
                    },
                },
            },
            RemainingCredit: codatsyncpayables.Float64(0),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "alias": map[string]interface{}{
                        "itaque": "velit",
                        "laborum": "non",
                    },
                    "dolor": map[string]interface{}{
                        "sit": "doloremque",
                        "consequatur": "officia",
                    },
                    "recusandae": map[string]interface{}{
                        "quidem": "voluptas",
                        "facilis": "placeat",
                    },
                    "perspiciatis": map[string]interface{}{
                        "deleniti": "a",
                        "voluptate": "ullam",
                        "unde": "necessitatibus",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "ac55a974-1d31-4135-a965-bb8a72026114",
                SupplierName: codatsyncpayables.String("neque"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9323.94,
                    Name: "Tracy Mills",
                },
                shared.WithholdingTaxitems{
                    Amount: 8028.94,
                    Name: "Marilyn Heaney",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(115661),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateBillCreditNoteRequest](../../models/operations/createbillcreditnoterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreateBillCreditNoteResponse](../../models/operations/createbillcreditnoteresponse.md), error**


## Get

The *Get bill credit note* endpoint returns a single bill credit note for a given `billCreditNoteId`.

[Bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support getting a specific bill credit note.

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
    res, err := s.BillCreditNotes.Get(ctx, operations.GetBillCreditNoteRequest{
        BillCreditNoteID: "id",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetBillCreditNoteRequest](../../models/operations/getbillcreditnoterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.GetBillCreditNoteResponse](../../models/operations/getbillcreditnoteresponse.md), error**


## GetCreateUpdateModel

The *Get create/update bill credit note model* endpoint returns the expected data for the request payload when creating and updating a [bill credit note](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) for a given company and integration.

[Bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating and updating a bill credit note.


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
    res, err := s.BillCreditNotes.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBillCreditNoteModelRequest{
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetCreateUpdateBillCreditNoteModelRequest](../../models/operations/getcreateupdatebillcreditnotemodelrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.GetCreateUpdateBillCreditNoteModelResponse](../../models/operations/getcreateupdatebillcreditnotemodelresponse.md), error**


## List

The *List bill credit notes* endpoint returns a list of [bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

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
    res, err := s.BillCreditNotes.List(ctx, operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncpayables.String("-modifiedDate"),
        Page: codatsyncpayables.Int(1),
        PageSize: codatsyncpayables.Int(100),
        Query: codatsyncpayables.String("libero"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListBillCreditNotesRequest](../../models/operations/listbillcreditnotesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.ListBillCreditNotesResponse](../../models/operations/listbillcreditnotesresponse.md), error**


## Update

The *Update bill credit note* endpoint updates an existing [bill credit note](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/sync-for-payables-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-billCreditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating a bill credit note.


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
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codatsyncpayables.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codatsyncpayables.String("USD"),
            CurrencyRate: codatsyncpayables.Float64(5546.03),
            DiscountPercentage: 0,
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("70e1084c-b067-42d1-ad87-9eeb9665b85e"),
                        Name: codatsyncpayables.String("Mr. Jonathon Swaniawski"),
                    },
                    Description: codatsyncpayables.String("fuga"),
                    DiscountAmount: codatsyncpayables.Float64(9195.08),
                    DiscountPercentage: codatsyncpayables.Float64(340.7),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "be2d7822-59e3-4ea4-b519-7f92443da7ce",
                        Name: codatsyncpayables.String("Phyllis Quitzon"),
                    },
                    Quantity: 3262.69,
                    SubTotal: codatsyncpayables.Float64(8095.94),
                    TaxAmount: codatsyncpayables.Float64(3165.42),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsyncpayables.Float64(2040.72),
                        ID: codatsyncpayables.String("7c6454ef-b0b3-4489-ac3c-a5acfbe2fd57"),
                        Name: codatsyncpayables.String("Viola Hane"),
                    },
                    TotalAmount: codatsyncpayables.Float64(5678.46),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9177deac-646e-4cb5-b340-9e3eb1e5a2b1",
                                Name: codatsyncpayables.String("Ms. Kelley Rogahn"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("veritatis"),
                            ID: "16db9954-5fc9-45fa-8897-0e189dbb30fc",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "3ea055b1-97cd-444e-af52-d82d3513bb6f",
                            Name: codatsyncpayables.String("Mattie Raynor"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "bcdb35ff-2e4b-4275-b7a8-cd9e7319c177",
                            Name: codatsyncpayables.String("Leon Collier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "77b114ee-b52f-4f78-9fc3-7814d4c98e0c",
                            Name: codatsyncpayables.String("Candice Rath"),
                        },
                    },
                    UnitAmount: 9222.99,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Note: codatsyncpayables.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(6293.77),
                        TotalAmount: codatsyncpayables.Float64(8339.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("636c6005-03d8-4bb3-9180-f739ae9e057e"),
                            Name: codatsyncpayables.String("Johnnie Baumbach"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(5200.81),
                        ID: codatsyncpayables.String("10331f39-81d4-4c70-8b60-7f3c93c73b9d"),
                        Note: codatsyncpayables.String("officia"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("tenetur"),
                        TotalAmount: codatsyncpayables.Float64(1339.49),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsyncpayables.String("EUR"),
                        CurrencyRate: codatsyncpayables.Float64(8483.46),
                        TotalAmount: codatsyncpayables.Float64(6707.62),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("7e23f225-7411-4faf-8b75-44e472e80285"),
                            Name: codatsyncpayables.String("Marguerite Hansen"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: codatsyncpayables.Float64(2667.88),
                        ID: codatsyncpayables.String("63a7d575-f140-40e7-a4ad-7334ec1b781b"),
                        Note: codatsyncpayables.String("amet"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("fuga"),
                        TotalAmount: codatsyncpayables.Float64(53.1),
                    },
                },
            },
            RemainingCredit: codatsyncpayables.Float64(0),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quos": map[string]interface{}{
                        "repellendus": "veritatis",
                        "quae": "eaque",
                        "saepe": "delectus",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "ada200ef-0422-4eb2-964c-f9ab8366c723",
                SupplierName: codatsyncpayables.String("reiciendis"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 8611.23,
                    Name: "Mrs. Luther Torp",
                },
                shared.WithholdingTaxitems{
                    Amount: 9248.4,
                    Name: "Kyle Leffler",
                },
                shared.WithholdingTaxitems{
                    Amount: 7868.6,
                    Name: "Dr. Shari Roob Sr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 3178.98,
                    Name: "Ian Auer",
                },
            },
        },
        BillCreditNoteID: "a",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        TimeoutInMinutes: codatsyncpayables.Int(615058),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateBillCreditNoteRequest](../../models/operations/updatebillcreditnoterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.UpdateBillCreditNoteResponse](../../models/operations/updatebillcreditnoteresponse.md), error**

