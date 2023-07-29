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
    res, err := s.CreditNotes.Create(ctx, operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(6522.45),
            AdditionalTaxPercentage: codataccounting.Float64(2425.82),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("est"),
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(7021.83),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("blanditiis"),
                ID: "fe99731a-dc05-4d85-ae2d-fb70fb387429",
            },
            DiscountPercentage: 101.8,
            ID: codataccounting.String("d336561e-ca16-4ef8-9451-bd76eeeb518c"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a1fad355-12f0-46d4-a5b7-2f0f548568a0"),
                        Name: codataccounting.String("Bonnie Greenfelder Jr."),
                    },
                    Description: codataccounting.String("officia"),
                    DiscountAmount: codataccounting.Float64(1009.76),
                    DiscountPercentage: codataccounting.Float64(8459.84),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "6eb94346-45d0-4308-8fbb-a5cceff5cb01",
                        Name: codataccounting.String("Dr. Rogelio Haley"),
                    },
                    Quantity: 1344.12,
                    SubTotal: codataccounting.Float64(5425.06),
                    TaxAmount: codataccounting.Float64(6556.94),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2635.77),
                        ID: codataccounting.String("5ac82b85-f8bc-42ca-ba8d-a4127dd597ff"),
                        Name: codataccounting.String("Miss Joy Boyer"),
                    },
                    TotalAmount: codataccounting.Float64(909.26),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c74b86ce-cc74-4f77-b484-8bd6a6f0441d",
                                Name: codataccounting.String("Jody Dickens"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "08094373-e060-4459-bebb-ad02f2586bcf",
                                Name: codataccounting.String("Holly Dare"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8daa95be-6cd0-4275-ac35-4aa432b47e17",
                                Name: codataccounting.String("Cindy Schiller"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("alias"),
                            ID: "8c23e980-2d82-4f0d-85eb-4a8b674ee5cf",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8edc7f78-7e32-4e04-b3d3-ed0c5670ef42",
                            Name: codataccounting.String("Clint Ernser"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("transfer"),
                            ID: codataccounting.String("1cc503f6-c39b-4cd0-a629-0f957f385189"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d7ef807a-ae03-4f33-8a79-fb9de4032ba2",
                            Name: codataccounting.String("Jeannie Smith"),
                        },
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
                        CustomerRef: &shared.CustomerRef{
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
                        CustomerRef: &shared.CustomerRef{
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
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4d739656-4c20-4a07-91a9-61d24a7dbb8f"),
                        Name: codataccounting.String("Shannon Cremin"),
                    },
                    Description: codataccounting.String("perspiciatis"),
                    DiscountAmount: codataccounting.Float64(1421.56),
                    DiscountPercentage: codataccounting.Float64(7540.53),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f7812cb5-12c8-4782-80bf-548f88f8f1bf",
                        Name: codataccounting.String("Hannah Schmidt"),
                    },
                    Quantity: 1202.57,
                    SubTotal: codataccounting.Float64(9822.77),
                    TaxAmount: codataccounting.Float64(1756.76),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(195.51),
                        ID: codataccounting.String("6d5d831d-0081-4090-b670-6673f3a681c5"),
                        Name: codataccounting.String("Vera Kutch"),
                    },
                    TotalAmount: codataccounting.Float64(9328.07),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "42409a21-5e08-4601-889a-5f63e3af3dd9",
                                Name: codataccounting.String("Marty Nikolaus"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "dcd63483-e4a7-4a98-a4df-37e45b8955d4",
                                Name: codataccounting.String("Mrs. Tracy Walker"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("numquam"),
                            ID: "82310907-bd35-44c0-92bd-734f02449d86",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "bb20fe5d-911c-4bfe-b49c-af45a27f69e2",
                            Name: codataccounting.String("Tracy Weber"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("journalEntry"),
                            ID: codataccounting.String("0e9db3ad-4c6b-4031-88d9-c337473082b9"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f2ab1fd5-671e-49c3-a635-0a467143789c",
                            Name: codataccounting.String("William Walker"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1594d93a-74c0-4252-be3b-4b4db8b778eb",
                            Name: codataccounting.String("Dr. Ramon Towne"),
                        },
                    },
                    UnitAmount: 7725.93,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("ad"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6732.9),
                        TotalAmount: codataccounting.Float64(9436.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b2cbc463-5d5e-465d-a028-c3e951a1e30f"),
                            Name: codataccounting.String("Doug Marquardt"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5226.16),
                        ID: codataccounting.String("9d7b7867-3e13-4a12-a6b9-92494594487f"),
                        Note: codataccounting.String("veniam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("praesentium"),
                        TotalAmount: codataccounting.Float64(2852.56),
                    },
                },
            },
            RemainingCredit: 2293.64,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: 4060.7,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "rem": map[string]interface{}{
                        "harum": "consectetur",
                        "quisquam": "nulla",
                    },
                    "a": map[string]interface{}{
                        "dolore": "dicta",
                        "minima": "facilis",
                    },
                    "sit": map[string]interface{}{
                        "magnam": "molestias",
                        "hic": "error",
                    },
                },
            },
            TotalAmount: 8334,
            TotalDiscount: 9625.75,
            TotalTaxAmount: 1189.86,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9646.4,
                    Name: "Gretchen Waters",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(923982),
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
    res, err := s.CreditNotes.Get(ctx, operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "in",
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
    res, err := s.CreditNotes.List(ctx, operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("deleniti"),
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
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(7347.74),
            AdditionalTaxPercentage: codataccounting.Float64(9709.57),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("sit"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(5624.3),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("quia"),
                ID: "5894ea76-3d5c-4727-95b7-85148d6d549e",
            },
            DiscountPercentage: 3248.72,
            ID: codataccounting.String("635b33bc-0f97-40c4-afc9-f4844225e75b"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6065c0ef-a6f9-43b9-8a1b-8c95be1254b7"),
                        Name: codataccounting.String("Ramona Wisoky"),
                    },
                    Description: codataccounting.String("eveniet"),
                    DiscountAmount: codataccounting.Float64(4855.06),
                    DiscountPercentage: codataccounting.Float64(4518.47),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "210d1f65-58c9-49c7-a2d2-bc0f94087d9c",
                        Name: codataccounting.String("Mrs. Gerard Walter"),
                    },
                    Quantity: 8445.66,
                    SubTotal: codataccounting.Float64(8404.68),
                    TaxAmount: codataccounting.Float64(4522.21),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8119.81),
                        ID: codataccounting.String("aac9b4ca-a1cf-4e9e-95df-903907f37831"),
                        Name: codataccounting.String("Isaac Dickinson"),
                    },
                    TotalAmount: codataccounting.Float64(1409.72),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "54a85466-597c-4502-b3c1-471d51aaa6dd",
                                Name: codataccounting.String("Corey Pacocha"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6487c5fc-2b86-42a0-8bef-69e100157630",
                                Name: codataccounting.String("Taylor Paucek"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fded84a3-5a41-4238-a1a7-35ac26ae33be",
                                Name: codataccounting.String("Miss Terrence Kulas"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f46bca11-06fe-4965-b711-d08cf88ec9f7",
                                Name: codataccounting.String("Freddie Mayert"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quis"),
                            ID: "0a656ed3-33bb-40ce-8aa6-5432a986eb7e",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "ca564089-1500-4970-99a4-8f88ece7bf90",
                            Name: codataccounting.String("Mr. Kerry Adams II"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("transfer"),
                            ID: codataccounting.String("38908162-c6be-4b68-a0f6-57b7d03a1480"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8de30f06-9d81-4061-8d97-e152297510da",
                            Name: codataccounting.String("Mr. Kenneth Fay"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92cc61c2-a702-4bb9-bee1-02da2de35f8e",
                            Name: codataccounting.String("Diane Rempel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3eaab454-02ac-4170-8bf1-cc9fc61aae5e",
                            Name: codataccounting.String("Miss Jon Willms"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92b5744d-08a2-4267-aaee-79e3c71ad31b",
                            Name: codataccounting.String("Benny Raynor"),
                        },
                    },
                    UnitAmount: 8588.02,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2378ae3b-fc23-4d94-90a9-86a495bac707"),
                        Name: codataccounting.String("Timothy Jaskolski"),
                    },
                    Description: codataccounting.String("laudantium"),
                    DiscountAmount: codataccounting.Float64(9219.81),
                    DiscountPercentage: codataccounting.Float64(7954.89),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c8649238-6f62-4c96-9c4c-c6b78890a3fd",
                        Name: codataccounting.String("Dr. Kara Lowe"),
                    },
                    Quantity: 1194.73,
                    SubTotal: codataccounting.Float64(207.35),
                    TaxAmount: codataccounting.Float64(9901.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5342.04),
                        ID: codataccounting.String("c23df931-da3e-4db5-9fad-94acc9435137"),
                        Name: codataccounting.String("Sara Jast II"),
                    },
                    TotalAmount: codataccounting.Float64(2426.06),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "1b832a56-d691-480f-b60e-b9a6658e69a4",
                                Name: codataccounting.String("Jimmie Grady"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("adipisci"),
                            ID: "82dbec75-c68c-4606-9946-8ce304d8849b",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "214c337f-96bb-40c6-9e37-2db1344ba9f7",
                            Name: codataccounting.String("Ernesto Heaney PhD"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("transfer"),
                            ID: codataccounting.String("7aab62e9-7261-4fb0-858d-27b51996b5b4"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "50eef712-b7a7-4ab0-b44b-1710688deebe",
                            Name: codataccounting.String("Wallace Medhurst"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3dd0ccd3-3f11-4b3e-8e08-0aa104186ec7",
                            Name: codataccounting.String("Mr. Sherri Torphy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3702c5c8-e2d3-40ea-9310-4fa44707bf37",
                            Name: codataccounting.String("Jeannette Graham"),
                        },
                    },
                    UnitAmount: 5596.68,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2821fdb2-f69e-4592-a7c7-1cc8d3cd4258"),
                        Name: codataccounting.String("Anthony Fahey"),
                    },
                    Description: codataccounting.String("laborum"),
                    DiscountAmount: codataccounting.Float64(5588.69),
                    DiscountPercentage: codataccounting.Float64(1387.08),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c808fe27-51a2-4047-8044-9e143f9619bb",
                        Name: codataccounting.String("Dr. Estelle Goyette"),
                    },
                    Quantity: 6731.58,
                    SubTotal: codataccounting.Float64(640.7),
                    TaxAmount: codataccounting.Float64(641.84),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9726.54),
                        ID: codataccounting.String("a436e625-9233-4f95-89d2-37397c785b5d"),
                        Name: codataccounting.String("Miguel Wuckert Jr."),
                    },
                    TotalAmount: codataccounting.Float64(1104.36),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "3febdf67-6b72-406d-ab75-0052a5647edc",
                                Name: codataccounting.String("Annie Morissette"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8c4320f4-1240-4d44-87ac-693b94c3b9d2",
                                Name: codataccounting.String("Leah Kuvalis"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95aa42fc-4056-469f-a9a0-06d212494508",
                                Name: codataccounting.String("Natasha Stark"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ipsum"),
                            ID: "b1b41844-060e-4003-90d0-23dc901f5afd",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6c44846a-e9d8-4925-bc89-62f4896bf51e",
                            Name: codataccounting.String("Eileen Hegmann"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: codataccounting.String("journalEntry"),
                            ID: codataccounting.String("c343d617-78af-4491-a477-25e621909e91"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "44a5de59-ac77-4066-b0cf-1cf593260525",
                            Name: codataccounting.String("Elvira Jacobson"),
                        },
                    },
                    UnitAmount: 7414,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("quia"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4937.34),
                        TotalAmount: codataccounting.Float64(8134.63),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("99a2d335-670e-493e-a6cf-59f358aaeaca"),
                            Name: codataccounting.String("Philip Crooks"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(799.07),
                        ID: codataccounting.String("bf7ba1cc-9771-46c8-82cc-9e0c7d9d323f"),
                        Note: codataccounting.String("quae"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("est"),
                        TotalAmount: codataccounting.Float64(4209.27),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8610.9),
                        TotalAmount: codataccounting.Float64(5823.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("cf1c856b-cba5-41ef-a454-a47facf116cd"),
                            Name: codataccounting.String("Jorge Grady"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4413.58),
                        ID: codataccounting.String("562873c7-dd9e-4faf-83dc-623620f3138f"),
                        Note: codataccounting.String("nesciunt"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("at"),
                        TotalAmount: codataccounting.Float64(9458.52),
                    },
                },
            },
            RemainingCredit: 1945.26,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusVoid,
            SubTotal: 224.79,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aspernatur": map[string]interface{}{
                        "similique": "id",
                        "exercitationem": "commodi",
                        "nostrum": "delectus",
                        "quidem": "rem",
                    },
                },
            },
            TotalAmount: 9947.56,
            TotalDiscount: 3840.74,
            TotalTaxAmount: 3327.12,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9230.44,
                    Name: "Pete Metz",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "praesentium",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(249962),
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

