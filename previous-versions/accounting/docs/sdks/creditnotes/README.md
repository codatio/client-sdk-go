# CreditNotes
(*CreditNotes*)

## Overview

Credit notes

### Available Operations

* [Create](#create) - Create credit note
* [Get](#get) - Get credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update credit note model
* [List](#list) - List credit notes
* [Update](#update) - Update credit note

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
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CreditNotes.Create(ctx, operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: types.MustNewDecimalFromString("4865.89"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("4893.82"),
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: accounting.String("Money blue shred"),
            Currency: accounting.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("9510.62"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: accounting.String("Abbott, Klocko and Dach"),
                ID: "<ID>",
            },
            DiscountPercentage: types.MustNewDecimalFromString("3015.1"),
            ID: accounting.String("<ID>"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("fuchsia Gasoline Screen"),
                    },
                    Description: accounting.String("Networked web-enabled monitoring"),
                    DiscountAmount: types.MustNewDecimalFromString("3570.21"),
                    DiscountPercentage: types.MustNewDecimalFromString("285.48"),
                    IsDirectIncome: accounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("after"),
                    },
                    Quantity: types.MustNewDecimalFromString("5190.28"),
                    SubTotal: types.MustNewDecimalFromString("2303.13"),
                    TaxAmount: types.MustNewDecimalFromString("2075.65"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("2113.37"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("Buckinghamshire functionalities Grocery"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("738.99"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: accounting.String("Northwest Direct"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("Stracke - Bashirian"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "<ID>",
                            Name: accounting.String("Senior Mouse West"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: accounting.String("accountTransaction"),
                            ID: accounting.String("<ID>"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("Edinburg Investor"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("5504.83"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("Dollar 1080p Rubber"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: accounting.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("1373.24"),
                        TotalAmount: types.MustNewDecimalFromString("2734.46"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: accounting.String("<ID>"),
                            Name: accounting.String("Toyota Neptunium round"),
                        },
                        Currency: accounting.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("2199.74"),
                        ID: accounting.String("<ID>"),
                        Note: accounting.String("invoice Dollar Electronic"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: accounting.String("markets"),
                        TotalAmount: types.MustNewDecimalFromString("9244.84"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("9824.5"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("1296.63"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nemo": map[string]interface{}{
                        "labore": "pixel",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("3583.09"),
            TotalDiscount: types.MustNewDecimalFromString("5193.59"),
            TotalTaxAmount: types.MustNewDecimalFromString("1932.39"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("3259.9"),
                    Name: "Credit female facere",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: accounting.Int(445608),
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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CreditNotes.Get(ctx, operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "Northeast Hatchback Kia",
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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CreditNotes.List(ctx, operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: accounting.String("-modifiedDate"),
        Page: accounting.Int(1),
        PageSize: accounting.Int(100),
        Query: accounting.String("Northeast Metal Canada"),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: types.MustNewDecimalFromString("8574.78"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("245.55"),
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: accounting.String("Reactive"),
            Currency: accounting.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("2703.24"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: accounting.String("Paucek - Kuhn"),
                ID: "<ID>",
            },
            DiscountPercentage: types.MustNewDecimalFromString("4430.76"),
            ID: accounting.String("<ID>"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("Islands"),
                    },
                    Description: accounting.String("Networked heuristic framework"),
                    DiscountAmount: types.MustNewDecimalFromString("3115.07"),
                    DiscountPercentage: types.MustNewDecimalFromString("7884.4"),
                    IsDirectIncome: accounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("bifurcated"),
                    },
                    Quantity: types.MustNewDecimalFromString("6447.13"),
                    SubTotal: types.MustNewDecimalFromString("7892.75"),
                    TaxAmount: types.MustNewDecimalFromString("9936.8"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("8898.38"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("implement JBOD"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("7712.03"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: accounting.String("Representative Home"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("Roob - Hermiston"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "<ID>",
                            Name: accounting.String("Northeast Wooden"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: accounting.String("invoice"),
                            ID: accounting.String("<ID>"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("Internal invoice"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("2803.6"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("female"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: accounting.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("3297.12"),
                        TotalAmount: types.MustNewDecimalFromString("4939.56"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: accounting.String("<ID>"),
                            Name: accounting.String("modulo Kia"),
                        },
                        Currency: accounting.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("8000.94"),
                        ID: accounting.String("<ID>"),
                        Note: accounting.String("Avon"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: accounting.String("Reactive Global Northeast"),
                        TotalAmount: types.MustNewDecimalFromString("6090.5"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("9920.1"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: types.MustNewDecimalFromString("4798.89"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "perferendis": map[string]interface{}{
                        "architecto": "quisquam",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("8523.4"),
            TotalDiscount: types.MustNewDecimalFromString("3379.02"),
            TotalTaxAmount: types.MustNewDecimalFromString("7416.05"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("3486.27"),
                    Name: "Money",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "Group wank Latvian",
        ForceUpdate: accounting.Bool(false),
        TimeoutInMinutes: accounting.Int(934395),
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

