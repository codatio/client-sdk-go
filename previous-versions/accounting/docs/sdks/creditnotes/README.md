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
    res, err := s.CreditNotes.Create(ctx, operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(5662.62),
            AdditionalTaxPercentage: codataccounting.Float64(5345.33),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("perspiciatis"),
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(3581.57),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codataccounting.String("ullam"),
                ID: "7389cedb-ac7f-4da3-9594-d66bc2ae4806",
            },
            DiscountPercentage: 2099.2,
            ID: codataccounting.String("2b9954b6-fa22-4063-a982-8553cb10006b"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4921ec20-53b7-4493-a6ac-8ee0f2bf1958"),
                        Name: codataccounting.String("Dr. Irving Gislason I"),
                    },
                    Description: codataccounting.String("reiciendis"),
                    DiscountAmount: codataccounting.Float64(2044.66),
                    DiscountPercentage: codataccounting.Float64(8285.71),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "eba297be-3e90-4bc4-8df8-68fd52405cb3",
                        Name: codataccounting.String("Evelyn Stracke"),
                    },
                    Quantity: 1574.21,
                    SubTotal: codataccounting.Float64(9699.11),
                    TaxAmount: codataccounting.Float64(2650.6),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9870.85),
                        ID: codataccounting.String("127fb0e0-bf1f-4821-b978-d0acca77aeb7"),
                        Name: codataccounting.String("Erik Abbott MD"),
                    },
                    TotalAmount: codataccounting.Float64(3636.6),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "046b64e9-9fb0-4e67-a094-fdfed5540ef5",
                                Name: codataccounting.String("Angie Flatley"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("quasi"),
                            ID: "b8fe9973-1adc-405d-85ae-2dfb70fb3874",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0d336561-eca1-46ef-8945-1bd76eeeb518",
                            Name: codataccounting.String("Joel Senger DVM"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("accountTransaction"),
                            ID: codataccounting.String("d35512f0-6d4e-45b7-af0f-548568a0424e"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0a1d6eb9-4346-445d-8308-4fbba5cceff5",
                            Name: codataccounting.String("Dr. Malcolm Bauch"),
                        },
                    },
                    UnitAmount: 3147.28,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1e528a45-ac82-4b85-b8bc-2caba8da4127"),
                        Name: codataccounting.String("Drew Hickle"),
                    },
                    Description: codataccounting.String("hic"),
                    DiscountAmount: codataccounting.Float64(9598.28),
                    DiscountPercentage: codataccounting.Float64(2621.78),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "711aa1bc-74b8-46ce-8c74-f77b4848bd6a",
                        Name: codataccounting.String("Darla Aufderhar"),
                    },
                    Quantity: 997.32,
                    SubTotal: codataccounting.Float64(8603.83),
                    TaxAmount: codataccounting.Float64(1859.83),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8035.05),
                        ID: codataccounting.String("3b808094-373e-4060-859b-ebbad02f2586"),
                        Name: codataccounting.String("Mrs. Forrest Wilkinson"),
                    },
                    TotalAmount: codataccounting.Float64(3465.57),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "8daa95be-6cd0-4275-ac35-4aa432b47e17",
                                Name: codataccounting.String("Cindy Schiller"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "08c23e98-02d8-42f0-945e-b4a8b674ee5c",
                                Name: codataccounting.String("Lucas Cartwright"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("assumenda"),
                            ID: "c7f787e3-2e04-4b3d-bed0-c5670ef42bd3",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f1cc503f-6c39-4bcd-8a62-90f957f38518",
                            Name: codataccounting.String("Shannon Schuster"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("transfer"),
                            ID: codataccounting.String("807aae03-f33c-4a79-bb9d-e4032ba26fd3"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8ba9216b-cb41-4583-9c73-641723133edc",
                            Name: codataccounting.String("Juanita Kemmer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5163bbca-4922-47c4-ac22-c55350495c5d",
                            Name: codataccounting.String("Roderick Fisher"),
                        },
                    },
                    UnitAmount: 4435.65,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c1e4981e-8aa2-457d-9c19-12ebde64bfcc"),
                        Name: codataccounting.String("Megan Kertzmann"),
                    },
                    Description: codataccounting.String("quaerat"),
                    DiscountAmount: codataccounting.Float64(133.69),
                    DiscountPercentage: codataccounting.Float64(994.67),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5dfa7962-06be-4f2b-8a3e-42c1aa010e9a",
                        Name: codataccounting.String("Garrett Cassin"),
                    },
                    Quantity: 826.45,
                    SubTotal: codataccounting.Float64(2403.52),
                    TaxAmount: codataccounting.Float64(3369.7),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3367.21),
                        ID: codataccounting.String("86d18f9f-97a4-4bfa-92bf-7d67ca84ad99"),
                        Name: codataccounting.String("Ray Botsford"),
                    },
                    TotalAmount: codataccounting.Float64(984.52),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "43531870-cf68-4b03-ad42-1bd43d1f0cb0",
                                Name: codataccounting.String("Mrs. Brian Adams"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("facilis"),
                            ID: "22d9b3a7-0d94-4faa-b41c-57d1fedc2050",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8dc3ce18-5472-4f9e-a691-66a8be3444ea",
                            Name: codataccounting.String("Clayton Reinger"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("journalEntry"),
                            ID: codataccounting.String("875c6c1f-e606-4d07-92a9-c87ae50c1666"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a1d9136a-7e8d-4532-93f3-f658752db764",
                            Name: codataccounting.String("Dean Mayert MD"),
                        },
                    },
                    UnitAmount: 3349.54,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6cebcada-29ca-4791-81c9-5671663c530b"),
                        Name: codataccounting.String("Kristin Hudson III"),
                    },
                    Description: codataccounting.String("dolor"),
                    DiscountAmount: codataccounting.Float64(6359.03),
                    DiscountPercentage: codataccounting.Float64(2495.41),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "638512ab-2521-4b9f-ae07-2467b8a40bc0",
                        Name: codataccounting.String("May Parisian PhD"),
                    },
                    Quantity: 3937.88,
                    SubTotal: codataccounting.Float64(3385.42),
                    TaxAmount: codataccounting.Float64(196.1),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9319.91),
                        ID: codataccounting.String("df22a94d-20ec-490e-a41d-1f465e85156f"),
                        Name: codataccounting.String("Randal Kris"),
                    },
                    TotalAmount: codataccounting.Float64(8372.02),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "54fdd5ea-9543-4398-9afb-42a8d63388e4",
                                Name: codataccounting.String("Casey Anderson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ea5f9b18-a244-4fd6-9903-9dacd38ed0dc",
                                Name: codataccounting.String("Tamara Borer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7f1e3af1-5920-4c90-91b4-901f2bd89c8a",
                                Name: codataccounting.String("Wanda Kassulke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "da5b7b69-02b8-481a-94f6-43664a8f0af8",
                                Name: codataccounting.String("Dr. Edgar Mosciski"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("dolor"),
                            ID: "2d9fbaf9-476a-42ae-8dcc-50c8a3512c73",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "48930750-a00e-4966-ac73-6d43194398c7",
                            Name: codataccounting.String("Curtis Rosenbaum"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("journalEntry"),
                            ID: codataccounting.String("98ed3d3a-b7ca-43c5-8a86-49a70cfd5d69"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9b720645-1077-4d19-aa83-d492ed14b8a2",
                            Name: codataccounting.String("Ralph Marquardt"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "545e955d-cc18-45ea-8901-c7c43ad2daa7",
                            Name: codataccounting.String("Troy Murphy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3d230edf-7381-41a1-9538-2bd7ed565076",
                            Name: codataccounting.String("Christine Sauer"),
                        },
                    },
                    UnitAmount: 9851.55,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("nulla"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(3759.62),
                        TotalAmount: codataccounting.Float64(3441.48),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("64c20a07-11a9-461d-a4a7-dbb8f532d892"),
                            Name: codataccounting.String("Amos Kshlerin Sr."),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7206.43),
                        ID: codataccounting.String("512c8782-40bf-4548-b88f-8f1bf0bc8e1f"),
                        Note: codataccounting.String("dolores"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("eum"),
                        TotalAmount: codataccounting.Float64(8596.04),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(5581.38),
                        TotalAmount: codataccounting.Float64(1927.18),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1d008109-0f67-4066-b3f3-a681c5768dce"),
                            Name: codataccounting.String("Gail Christiansen V"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1374.14),
                        ID: codataccounting.String("15e08601-489a-45f6-be3a-f3dd9dda33dc"),
                        Note: codataccounting.String("pariatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("non"),
                        TotalAmount: codataccounting.Float64(2799.65),
                    },
                },
            },
            RemainingCredit: 5083.73,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusPartiallyPaid,
            SubTotal: 2523.23,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "esse": map[string]interface{}{
                        "natus": "quas",
                        "saepe": "modi",
                        "assumenda": "maiores",
                    },
                    "neque": map[string]interface{}{
                        "debitis": "quaerat",
                        "nostrum": "libero",
                    },
                    "totam": map[string]interface{}{
                        "veniam": "nostrum",
                        "facere": "aliquam",
                        "vitae": "ipsum",
                    },
                },
            },
            TotalAmount: 9233.27,
            TotalDiscount: 784.86,
            TotalTaxAmount: 2180.93,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 2578.35,
                    Name: "Mr. Todd Feil",
                },
                shared.WithholdingTaxitems{
                    Amount: 189.01,
                    Name: "Karla Schumm",
                },
                shared.WithholdingTaxitems{
                    Amount: 2666.43,
                    Name: "Charles McGlynn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(874446),
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
    res, err := s.CreditNotes.Get(ctx, operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "nihil",
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
    res, err := s.CreditNotes.List(ctx, operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("velit"),
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
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(2769.16),
            AdditionalTaxPercentage: codataccounting.Float64(9539.59),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("consequuntur"),
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(2532.61),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codataccounting.String("excepturi"),
                ID: "d86f4bb2-0fe5-4d91-9cbf-e749caf45a27",
            },
            DiscountPercentage: 9648.57,
            ID: codataccounting.String("69e2c9e6-d10e-49db-bad4-c6b03108d9c3"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("473082b9-4f2a-4b1f-9567-1e9c326350a4"),
                        Name: codataccounting.String("Tanya Braun"),
                    },
                    Description: codataccounting.String("nihil"),
                    DiscountAmount: codataccounting.Float64(5004.45),
                    DiscountPercentage: codataccounting.Float64(6160.39),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "ce0e9915-94d9-43a7-8c02-52fe3b4b4db8",
                        Name: codataccounting.String("Lance Kilback"),
                    },
                    Quantity: 7170.18,
                    SubTotal: codataccounting.Float64(7252.47),
                    TaxAmount: codataccounting.Float64(3909.37),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8879.8),
                        ID: codataccounting.String("1d2cf502-bafb-42cb-8463-5d5e65da028c"),
                        Name: codataccounting.String("Sophia Marvin MD"),
                    },
                    TotalAmount: codataccounting.Float64(847.08),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "30fda966-489d-47b7-8673-e13a12a6b992",
                                Name: codataccounting.String("Katrina Grimes"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4487f5c8-4383-46b8-ab3c-df6415b0449f",
                                Name: codataccounting.String("Mr. Marcos Windler"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4eedbe78-bf60-4682-9894-ea763d5c7279",
                                Name: codataccounting.String("Lula Koepp"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "148d6d54-9e56-435b-b3bc-0f970c42fc9f",
                                Name: codataccounting.String("Lena Green Sr."),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("veniam"),
                            ID: "e75b7960-65c0-4efa-af93-b90a1b8c95be",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "54b739f4-fe77-4210-91f6-558c99c722d2",
                            Name: codataccounting.String("Lucas Barton"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("invoice"),
                            ID: codataccounting.String("087d9caa-e042-4dd7-8aac-9b4caa1cfe9e"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5df90390-7f37-4831-983d-42e54a854665",
                            Name: codataccounting.String("Tyrone Ruecker Sr."),
                        },
                    },
                    UnitAmount: 2138.71,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3c1471d5-1aaa-46dd-b5ab-d6487c5fc2b8"),
                        Name: codataccounting.String("Mr. Kelly Orn"),
                    },
                    Description: codataccounting.String("saepe"),
                    DiscountAmount: codataccounting.Float64(9832.72),
                    DiscountPercentage: codataccounting.Float64(4228.23),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "9e100157-630b-4da7-afde-d84a35a41238",
                        Name: codataccounting.String("Walter Nitzsche"),
                    },
                    Quantity: 3217.24,
                    SubTotal: codataccounting.Float64(6847.17),
                    TaxAmount: codataccounting.Float64(7905.34),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1866.88),
                        ID: codataccounting.String("6ae33bef-971a-48f4-abca-1106fe965b71"),
                        Name: codataccounting.String("Meredith Barrows"),
                    },
                    TotalAmount: codataccounting.Float64(9475.73),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "8ec9f7b9-9a55-40a6-96ed-333bb0ce8aa6",
                                Name: codataccounting.String("Pauline Feeney"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "986eb7e1-4ca5-4640-8915-0097019a48f8",
                                Name: codataccounting.String("Darin Ryan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "bf904e01-105d-4389-8816-2c6beb68a0f6",
                                Name: codataccounting.String("Jo Pouros"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codataccounting.String("alias"),
                            ID: "3a1480f8-de30-4f06-9d81-0618d97e1522",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "510da803-1229-42cc-a1c2-a702bb97ee10",
                            Name: codataccounting.String("Ginger O'Reilly"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("transfer"),
                            ID: codataccounting.String("35f8e01b-f33e-4aab-8540-2ac1704bf1cc"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fc61aae5-eb5f-40c4-92b5-744d08a2267a",
                            Name: codataccounting.String("Sheldon Toy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e3c71ad3-1bec-4b83-9237-8ae3bfc23d94",
                            Name: codataccounting.String("Sarah Okuneva"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6a495bac-707f-406b-a8ec-c86492386f62",
                            Name: codataccounting.String("Wade Kemmer"),
                        },
                    },
                    UnitAmount: 2765.07,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("quod"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(5109.78),
                        TotalAmount: codataccounting.Float64(5207.16),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("90a3fd3c-81da-410f-8c23-df931da3edb5"),
                            Name: codataccounting.String("Winifred O'Reilly"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(6313.52),
                        ID: codataccounting.String("cc943513-7726-4d15-b21b-832a56d69180"),
                        Note: codataccounting.String("hic"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("iure"),
                        TotalAmount: codataccounting.Float64(84.69),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(5778.54),
                        TotalAmount: codataccounting.Float64(6825.06),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6658e69a-4b84-43d3-82db-ec75c68c6065"),
                            Name: codataccounting.String("Tom Homenick"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2073.26),
                        ID: codataccounting.String("04d8849b-f821-44c3-b7f9-6bb0c69e372d"),
                        Note: codataccounting.String("libero"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("non"),
                        TotalAmount: codataccounting.Float64(2896.81),
                    },
                },
            },
            RemainingCredit: 2774.88,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusPaid,
            SubTotal: 6213.54,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "reprehenderit": map[string]interface{}{
                        "est": "quis",
                        "impedit": "accusantium",
                        "necessitatibus": "facere",
                    },
                    "reprehenderit": map[string]interface{}{
                        "officia": "soluta",
                        "suscipit": "explicabo",
                        "recusandae": "unde",
                    },
                    "iusto": map[string]interface{}{
                        "ea": "architecto",
                    },
                    "earum": map[string]interface{}{
                        "alias": "quod",
                        "veniam": "corrupti",
                        "temporibus": "odit",
                    },
                },
            },
            TotalAmount: 4514.48,
            TotalDiscount: 7493.43,
            TotalTaxAmount: 3483.86,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5878.17,
                    Name: "Ruben Rice",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "modi",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(711732),
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

