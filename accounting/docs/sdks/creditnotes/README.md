# CreditNotes

## Overview

Credit notes

### Available Operations

* [Create](#create) - Create credit note
* [Get](#get) - Get credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update credit note model
* [List](#list) - List credit notes
* [Update](#update) - Update creditNote

## Create

The *Create credit note* endpoint creates a new [credit note](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating an account.


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
    res, err := s.CreditNotes.Create(ctx, operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(5321.1),
            AdditionalTaxPercentage: codataccounting.Float64(8172.49),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("natus"),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(9143.99),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("illo"),
                ID: "dd7097b5-da08-4c57-ba6c-78a216e19baf",
            },
            DiscountPercentage: 9226.4,
            ID: codataccounting.String("ca619149-8140-4b64-bf8a-e170ef03b5f3"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4aa86855-5966-4732-aa5d-cb6682cb70f8"),
                        Name: codataccounting.String("Roman Shanahan"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(3888.51),
                    DiscountPercentage: codataccounting.Float64(9347.07),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "91b9a9f7-4846-4e2c-b309-db0536d9e75c",
                        Name: codataccounting.String("Robert Balistreri"),
                    },
                    Quantity: 3375.81,
                    SubTotal: codataccounting.Float64(2117.56),
                    TaxAmount: codataccounting.Float64(6091.64),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1500.91),
                        ID: codataccounting.String("c11a25a8-bf92-4f97-828a-d9a9f8bf8221"),
                        Name: codataccounting.String("Lori Hammes"),
                    },
                    TotalAmount: codataccounting.Float64(5929.46),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "98387f7a-79cd-472c-9248-4da21729f2ac",
                                Name: codataccounting.String("Amanda Tromp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "725f1169-ac1e-441d-8a23-c23e34f2dfa4",
                                Name: codataccounting.String("Lawrence Metz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6de92215-1fe1-4712-8998-53e9f543d854",
                                Name: codataccounting.String("Rosa Metz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "22446044-3bc1-4541-88c2-f56e85da7832",
                                Name: codataccounting.String("Hubert Rempel"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("et"),
                            ID: "7c3b0d51-a44b-4f01-bad8-706d46082bfb",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "41ff5d4e-2ae4-4fb5-8b35-d17638f1edb7",
                            Name: codataccounting.String("Jeffery Hilll"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c5cb860f-8cd5-480b-a738-10e4fe444729",
                            Name: codataccounting.String("Erma Streich"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1dd3bbce-247b-4768-8eff-50126d71cffb",
                            Name: codataccounting.String("David Waters"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4b842195-3b44-4bd3-8431-59d33e5953c0",
                            Name: codataccounting.String("Cheryl Bins"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "863aa41e-6c31-4cc2-b1fc-b51c9a41ffbe",
                            Name: codataccounting.String("Forrest Powlowski"),
                        },
                    },
                    UnitAmount: 6205.2,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5ee65e07-6cc7-4abf-a16e-a5c71641934b"),
                        Name: codataccounting.String("Daniel Zemlak"),
                    },
                    Description: codataccounting.String("sit"),
                    DiscountAmount: codataccounting.Float64(5776.72),
                    DiscountPercentage: codataccounting.Float64(8573.88),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "19d2fc2f-9e2e-4105-944b-935d237a72f9",
                        Name: codataccounting.String("Mattie Gutkowski"),
                    },
                    Quantity: 3849.39,
                    SubTotal: codataccounting.Float64(6596.96),
                    TaxAmount: codataccounting.Float64(9303.06),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8266.83),
                        ID: codataccounting.String("4aecb753-7cd9-4222-89ff-57491aabfa2e"),
                        Name: codataccounting.String("Agnes Boyle DDS"),
                    },
                    TotalAmount: codataccounting.Float64(6645.01),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d456ef10-31e6-4899-b0c2-001e22cd55cc",
                                Name: codataccounting.String("Ida MacGyver"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "184d76d9-71fc-4820-865b-037bb8e0cc88",
                                Name: codataccounting.String("Carolyn Macejkovic"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("labore"),
                            ID: "de04af28-c5dd-4db4-aaa1-cfd6d828da01",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "91129646-645c-41d8-9f29-042f569b7aff",
                            Name: codataccounting.String("Tasha Pagac"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6cbe071b-c163-4e27-9a3b-084da99257d0",
                            Name: codataccounting.String("Ms. Ollie Gibson"),
                        },
                    },
                    UnitAmount: 4782.16,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a742d844-96cb-4dee-8f6b-99bc63562ebf"),
                        Name: codataccounting.String("Van Halvorson"),
                    },
                    Description: codataccounting.String("dolores"),
                    DiscountAmount: codataccounting.Float64(6102.13),
                    DiscountPercentage: codataccounting.Float64(2859.22),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c060b06a-1287-4764-aef6-d0c6d6ed9c73",
                        Name: codataccounting.String("Lionel Kerluke"),
                    },
                    Quantity: 3419.34,
                    SubTotal: codataccounting.Float64(4922.27),
                    TaxAmount: codataccounting.Float64(768.18),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3613.31),
                        ID: codataccounting.String("09a8e870-d3c5-4a1f-9c24-2c7b66a1f30c"),
                        Name: codataccounting.String("Ethel Schultz"),
                    },
                    TotalAmount: codataccounting.Float64(6939.88),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "719890f4-2a4b-4b43-8d85-b260591d745e",
                                Name: codataccounting.String("Mrs. Sheri Cruickshank"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c9c3f567-e0e2-4527-a5b1-d62fcdace1f0",
                                Name: codataccounting.String("Christina Bode"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("repudiandae"),
                            ID: "2239e8f2-5cd0-4d19-9959-f439e39266cb",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "5f7aa2b2-4113-4695-91e6-698fcc459621",
                            Name: codataccounting.String("Kendra D'Amore"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "67633425-4038-4bfb-9971-e98190557389",
                            Name: codataccounting.String("Percy Swaniawski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c7fda395-94d6-46bc-aae4-80632b9954b6",
                            Name: codataccounting.String("Mr. Alfonso Dibbert"),
                        },
                    },
                    UnitAmount: 2026.42,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("69828553-cb10-4006-bef4-921ec2053b74"),
                        Name: codataccounting.String("Marvin Jacobson"),
                    },
                    Description: codataccounting.String("quisquam"),
                    DiscountAmount: codataccounting.Float64(5223.27),
                    DiscountPercentage: codataccounting.Float64(9109.84),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "e0f2bf19-588d-440d-83f3-deba297be3e9",
                        Name: codataccounting.String("Maryann Rolfson PhD"),
                    },
                    Quantity: 9843.02,
                    SubTotal: codataccounting.Float64(5544.17),
                    TaxAmount: codataccounting.Float64(4169.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5023.8),
                        ID: codataccounting.String("fd52405c-b331-4d49-af4f-127fb0e0bf1f"),
                        Name: codataccounting.String("Steve Block"),
                    },
                    TotalAmount: codataccounting.Float64(4940.93),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d0acca77-aeb7-4b70-a1a5-2046b64e99fb",
                                Name: codataccounting.String("Gretchen Huels"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "094fdfed-5540-4ef5-ba34-a1b8fe99731a",
                                Name: codataccounting.String("Rodolfo Baumbach"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "85ae2dfb-70fb-4387-8290-d336561eca16",
                                Name: codataccounting.String("Santos Langosh"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ad"),
                            ID: "1bd76eee-b518-4c4d-a1fa-d35512f06d4e",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "72f0f548-568a-4042-8e00-a1d6eb943464",
                            Name: codataccounting.String("Cristina Beer V"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fbba5cce-ff5c-4b01-be51-e528a45ac82b",
                            Name: codataccounting.String("Derrick Wunsch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c2caba8d-a412-47dd-997f-f4711aa1bc74",
                            Name: codataccounting.String("Alberto Hyatt"),
                        },
                    },
                    UnitAmount: 7880.36,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("ducimus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4976.33),
                        TotalAmount: codataccounting.Float64(7222),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4848bd6a-6f04-441d-ac3b-808094373e06"),
                            Name: codataccounting.String("Lucille Hermiston"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7001.12),
                        ID: codataccounting.String("bad02f25-86bc-4f15-a558-daa95be6cd02"),
                        Note: codataccounting.String("in"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("vel"),
                        TotalAmount: codataccounting.Float64(7736.78),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(2817.53),
                        TotalAmount: codataccounting.Float64(6681.55),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a432b47e-1763-4c52-88c2-3e9802d82f0d"),
                            Name: codataccounting.String("Dolores Waelchi"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(5034.79),
                        ID: codataccounting.String("b674ee5c-fc18-4edc-bf78-7e32e04b3d3e"),
                        Note: codataccounting.String("facere"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("cumque"),
                        TotalAmount: codataccounting.Float64(3532.32),
                    },
                },
            },
            RemainingCredit: 4271.07,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: 9262.25,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolore": map[string]interface{}{
                        "harum": "illum",
                    },
                    "dolor": map[string]interface{}{
                        "iste": "earum",
                        "vitae": "impedit",
                        "eligendi": "veniam",
                        "aperiam": "consectetur",
                    },
                    "repellat": map[string]interface{}{
                        "quod": "nesciunt",
                        "iste": "distinctio",
                    },
                    "cumque": map[string]interface{}{
                        "alias": "deserunt",
                        "vel": "qui",
                        "perspiciatis": "accusantium",
                        "voluptatibus": "occaecati",
                    },
                },
            },
            TotalAmount: 3658.63,
            TotalDiscount: 4626.73,
            TotalTaxAmount: 9704.91,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5155.49,
                    Name: "Virginia Littel",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(822084),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateCreditNoteRequest](../../models/operations/createcreditnoterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.CreateCreditNoteResponse](../../models/operations/createcreditnoteresponse.md), error**


## Get

The *Get credit note* endpoint returns a single credit note for a given creditNoteId.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support getting a specific credit note.

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
    res, err := s.CreditNotes.Get(ctx, operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCreditNoteRequest](../../models/operations/getcreditnoterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.GetCreditNoteResponse](../../models/operations/getcreditnoteresponse.md), error**


## GetCreateUpdateModel

﻿The *Get create/update credit note model* endpoint returns the expected data for the request payload when creating and updating a [credit note](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company and integration.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating and updating a credit note.


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
    res, err := s.CreditNotes.GetCreateUpdateModel(ctx, operations.GetCreateUpdateCreditNotesModelRequest{
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetCreateUpdateCreditNotesModelRequest](../../models/operations/getcreateupdatecreditnotesmodelrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.GetCreateUpdateCreditNotesModelResponse](../../models/operations/getcreateupdatecreditnotesmodelresponse.md), error**


## List

The *List credit notes* endpoint returns a list of [credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

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
    res, err := s.CreditNotes.List(ctx, operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("officiis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCreditNotesRequest](../../models/operations/listcreditnotesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.ListCreditNotesResponse](../../models/operations/listcreditnotesresponse.md), error**


## Update

The *Update credit note* endpoint updates an existing [credit note](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating an account.


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
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(9713.93),
            AdditionalTaxPercentage: codataccounting.Float64(5288.35),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("reprehenderit"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(6380.92),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("eveniet"),
                ID: "03f33ca7-9fb9-4de4-832b-a26fd368ba92",
            },
            DiscountPercentage: 1199.03,
            ID: codataccounting.String("6bcb4158-35c7-4364-9723-133edc046bc5"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3bbca492-27c4-42c2-ac55-350495c5dbb3"),
                        Name: codataccounting.String("Derek Kihn DVM"),
                    },
                    Description: codataccounting.String("tempora"),
                    DiscountAmount: codataccounting.Float64(6128.36),
                    DiscountPercentage: codataccounting.Float64(5607.36),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1e8aa257-ddc1-4912-abde-64bfcc5469d4",
                        Name: codataccounting.String("Diane Hayes"),
                    },
                    Quantity: 6658.07,
                    SubTotal: codataccounting.Float64(4582.2),
                    TaxAmount: codataccounting.Float64(6140.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4054.19),
                        ID: codataccounting.String("206bef2b-0a3e-442c-9aa0-10e9aac2e913"),
                        Name: codataccounting.String("Erica Luettgen"),
                    },
                    TotalAmount: codataccounting.Float64(1161.94),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f9f97a4b-fad2-4bf7-967c-a84ad99b41d6",
                                Name: codataccounting.String("Sara Funk"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "31870cf6-8b03-4ad4-a1bd-43d1f0cb0a00",
                                Name: codataccounting.String("Carmen Weber"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2d9b3a70-d94f-4aa7-81c5-7d1fedc2050d",
                                Name: codataccounting.String("Olga Stracke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quo"),
                            ID: "e185472f-9ee6-4916-aa8b-e3444eac8b3a",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "75c6c1fe-606d-407d-aa9c-87ae50c16661",
                            Name: codataccounting.String("Harold Smith I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a7e8d532-13f3-4f65-8752-db764c59f0a5",
                            Name: codataccounting.String("Nichole Treutel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ada29ca7-9181-4c95-a716-63c530b56651",
                            Name: codataccounting.String("Carmen Nicolas"),
                        },
                    },
                    UnitAmount: 2274.7,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8512ab25-21b9-4f2e-8724-67b8a40bc05f"),
                        Name: codataccounting.String("Wm Bartoletti"),
                    },
                    Description: codataccounting.String("quis"),
                    DiscountAmount: codataccounting.Float64(196.1),
                    DiscountPercentage: codataccounting.Float64(9319.91),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "df22a94d-20ec-490e-a41d-1f465e85156f",
                        Name: codataccounting.String("Randal Kris"),
                    },
                    Quantity: 8372.02,
                    SubTotal: codataccounting.Float64(9459.21),
                    TaxAmount: codataccounting.Float64(3424.58),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2726.64),
                        ID: codataccounting.String("fdd5ea95-4339-48da-bb42-a8d63388e4d8"),
                        Name: codataccounting.String("Peggy Muller"),
                    },
                    TotalAmount: codataccounting.Float64(3381.03),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9b18a244-fd61-4903-9dac-d38ed0dc671d",
                                Name: codataccounting.String("Dr. Jamie Wintheiser"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "af15920c-90d1-4b49-81f2-bd89c8a32639",
                                Name: codataccounting.String("Donnie Hauck"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b6902b88-1a94-4f64-b664-a8f0af8c691d",
                                Name: codataccounting.String("Carmen Crist"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fbaf9476-a2ae-48dc-850c-8a3512c73784",
                                Name: codataccounting.String("Ms. Eduardo Effertz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eaque"),
                            ID: "a00e966e-c736-4d43-9943-98c783c92398",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3d3ab7ca-3c5c-4a86-89a7-0cfd5d6989b7",
                            Name: codataccounting.String("Betty Hirthe"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "077d19ea-83d4-492e-914b-8a2c1954545e",
                            Name: codataccounting.String("Derrick Halvorson"),
                        },
                    },
                    UnitAmount: 7659,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("praesentium"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(3007.59),
                        TotalAmount: codataccounting.Float64(6029.32),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("01c7c43a-d2da-4a78-8aba-3d230edf7381"),
                            Name: codataccounting.String("Mrs. Maggie Breitenberg"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1641.54),
                        ID: codataccounting.String("bd7ed565-0762-41c5-8f4d-7396564c20a0"),
                        Note: codataccounting.String("reprehenderit"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("veritatis"),
                        TotalAmount: codataccounting.Float64(6283.25),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1166.35),
                        TotalAmount: codataccounting.Float64(8489.26),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("24a7dbb8-f532-4d89-acf7-812cb512c878"),
                            Name: codataccounting.String("Monica Bashirian"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(2743.68),
                        ID: codataccounting.String("8f88f8f1-bf0b-4c8e-9f20-6d5d831d0081"),
                        Note: codataccounting.String("voluptatem"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("ipsa"),
                        TotalAmount: codataccounting.Float64(9652.07),
                    },
                },
            },
            RemainingCredit: 3763.41,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: 3777.59,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ducimus": map[string]interface{}{
                        "doloribus": "ratione",
                    },
                    "id": map[string]interface{}{
                        "quos": "dicta",
                        "minus": "exercitationem",
                    },
                },
            },
            TotalAmount: 4741.06,
            TotalDiscount: 4348.74,
            TotalTaxAmount: 5007,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 7742.94,
                    Name: "Clinton Hackett",
                },
                shared.WithholdingTaxitems{
                    Amount: 334.24,
                    Name: "Mrs. Gerard Collins",
                },
                shared.WithholdingTaxitems{
                    Amount: 532.16,
                    Name: "Mrs. Shane Armstrong",
                },
                shared.WithholdingTaxitems{
                    Amount: 5831.69,
                    Name: "Darrell Yost",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "debitis",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(207202),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateCreditNoteRequest](../../models/operations/updatecreditnoterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.UpdateCreditNoteResponse](../../models/operations/updatecreditnoteresponse.md), error**

