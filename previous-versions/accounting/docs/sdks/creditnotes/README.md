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
            AdditionalTaxAmount: types.MustNewDecimalFromString("7239.42"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("7119.91"),
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: accounting.String("provident"),
            Currency: accounting.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("7000.45"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: accounting.String("dignissimos"),
                ID: "5dad636c-6005-403d-8bb3-1180f739ae9e",
            },
            DiscountPercentage: types.MustNewDecimalFromString("100.63"),
            ID: accounting.String("57eb809e-2810-4331-b398-1d4c700b607f"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("c93c73b9-da3f-42ce-9a7e-23f2257411fa"),
                        Name: accounting.String("Kyle Reichel"),
                    },
                    Description: accounting.String("labore"),
                    DiscountAmount: types.MustNewDecimalFromString("2543.82"),
                    DiscountPercentage: types.MustNewDecimalFromString("9211.93"),
                    IsDirectIncome: accounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "472e8028-57a5-4b40-863a-7d575f1400e7",
                        Name: accounting.String("Carrie Pagac"),
                    },
                    Quantity: types.MustNewDecimalFromString("2327.72"),
                    SubTotal: types.MustNewDecimalFromString("2006.37"),
                    TaxAmount: types.MustNewDecimalFromString("3106.29"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("9294.76"),
                        ID: accounting.String("c1b781b3-6a08-4088-9100-efada200ef04"),
                        Name: accounting.String("Phyllis Tremblay Sr."),
                    },
                    TotalAmount: types.MustNewDecimalFromString("3979.88"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "4cf9ab83-66c7-423f-bda9-e06bee4825c1",
                                Name: accounting.String("Colin Berge Sr."),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("enim"),
                            ID: "c80bff91-8544-4ec4-adef-cce8f1977773",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3562a7b4-08f0-45e3-948f-daf313a1f5fd",
                            Name: accounting.String("Troy Champlin"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: accounting.String("transfer"),
                            ID: accounting.String("0b36f25e-a944-4f3b-b56c-11f6c37a5126"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "243835bb-c05a-423a-85ce-fc5fde10a0ce",
                            Name: accounting.String("Mildred Kautzer"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("3548.21"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("accusantium"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: accounting.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("5905.85"),
                        TotalAmount: types.MustNewDecimalFromString("7658.33"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: accounting.String("6dc5e347-6279-49bf-bbe6-949fb2bb4eca"),
                            Name: accounting.String("Ben Satterfield"),
                        },
                        Currency: accounting.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("8487.22"),
                        ID: accounting.String("b3adebd5-daea-44c5-86a8-aa94c02644cf"),
                        Note: accounting.String("nostrum"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: accounting.String("unde"),
                        TotalAmount: types.MustNewDecimalFromString("8603.11"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("6213.93"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: types.MustNewDecimalFromString("3442.89"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "esse": map[string]interface{}{
                        "corrupti": "fuga",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("8152.25"),
            TotalDiscount: types.MustNewDecimalFromString("7736.59"),
            TotalTaxAmount: types.MustNewDecimalFromString("986.1"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("6472.18"),
                    Name: "Dr. Rick Bauch",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: accounting.Int(807564),
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
        CreditNoteID: "consequatur",
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
        Query: accounting.String("eaque"),
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
            AdditionalTaxAmount: types.MustNewDecimalFromString("1023.9"),
            AdditionalTaxPercentage: types.MustNewDecimalFromString("6271.61"),
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: accounting.String("blanditiis"),
            Currency: accounting.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("1698.19"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: accounting.String("officiis"),
                ID: "2ec09ff8-f0f8-416f-b347-7c13e902c141",
            },
            DiscountPercentage: types.MustNewDecimalFromString("1391.33"),
            ID: accounting.String("5b0960a6-6815-41a4-b2af-923c5949f83f"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("50cf876f-fb90-41c6-acbb-4e243cf789ff"),
                        Name: accounting.String("Emilio Waters"),
                    },
                    Description: accounting.String("corporis"),
                    DiscountAmount: types.MustNewDecimalFromString("2465.77"),
                    DiscountPercentage: types.MustNewDecimalFromString("8877.01"),
                    IsDirectIncome: accounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "5ae6e0ac-184c-42b9-8247-c88373a40e19",
                        Name: accounting.String("Ashley Wunsch"),
                    },
                    Quantity: types.MustNewDecimalFromString("9368.45"),
                    SubTotal: types.MustNewDecimalFromString("3305.96"),
                    TaxAmount: types.MustNewDecimalFromString("3731.06"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("510.53"),
                        ID: accounting.String("55756f5d-56d0-4bd0-af2d-fe13db4f62cb"),
                        Name: accounting.String("Jacob Wehner"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("2524.73"),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "1aebc0b8-0a69-424d-bb2e-cfcc8f895010",
                                Name: accounting.String("Gordon Strosin"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("pariatur"),
                            ID: "6fa1804e-54c8-42f1-a8a3-63c8873e4843",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b1f6b8ca-275a-460a-84c4-95cc699171b5",
                            Name: accounting.String("Blanca Carroll"),
                        },
                        RecordRef: &shared.InvoiceTo{
                            DataType: accounting.String("accountTransaction"),
                            ID: accounting.String("1cf4b888-ebdf-4c4c-8ca9-9bc7fc0b2dce"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "10873e42-b006-4d67-8878-ba8581a58208",
                            Name: accounting.String("Lloyd Grant"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("9657.35"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("natus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: accounting.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("3125.11"),
                        TotalAmount: types.MustNewDecimalFromString("9853.79"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: accounting.String("2eac5565-d307-4cfe-a812-06e2813fa4a4"),
                            Name: accounting.String("Leticia Gerlach PhD"),
                        },
                        Currency: accounting.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("9998.54"),
                        ID: accounting.String("2132af03-102d-4514-b4cc-6f18bf9621a6"),
                        Note: accounting.String("animi"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: accounting.String("tenetur"),
                        TotalAmount: types.MustNewDecimalFromString("4934.07"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("4578.35"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("4610.5"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "eveniet": map[string]interface{}{
                        "earum": "velit",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("8847.65"),
            TotalDiscount: types.MustNewDecimalFromString("2633.46"),
            TotalTaxAmount: types.MustNewDecimalFromString("7019.78"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("9301.11"),
                    Name: "Brittany Cole",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "quis",
        ForceUpdate: accounting.Bool(false),
        TimeoutInMinutes: accounting.Int(704402),
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

