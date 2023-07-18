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
            AdditionalTaxAmount: codataccounting.Float64(6319.04),
            AdditionalTaxPercentage: codataccounting.Float64(8377.39),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("ad"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(3831.96),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("reiciendis"),
                ID: "df1ad837-ae80-4c1c-99c9-5ba998678fa3",
            },
            DiscountPercentage: 9409.51,
            ID: codataccounting.String("696991af-388c-4e03-a144-48c7977a0ef2"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("36028efe-ef93-4415-aed7-e253f4c157de"),
                        Name: codataccounting.String("Ms. Ernesto King DVM"),
                    },
                    Description: codataccounting.String("incidunt"),
                    DiscountAmount: codataccounting.Float64(2930.23),
                    DiscountPercentage: codataccounting.Float64(3626.93),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "accf667a-af9b-4bad-985f-e431d6bf5c83",
                        Name: codataccounting.String("Emilio Ratke"),
                    },
                    Quantity: 7992.36,
                    SubTotal: codataccounting.Float64(1330.76),
                    TaxAmount: codataccounting.Float64(538.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7734.55),
                        ID: codataccounting.String("b67fc4b4-25e9-49e6-a34c-9f7b79dfeb77"),
                        Name: codataccounting.String("Tommy Schmidt"),
                    },
                    TotalAmount: codataccounting.Float64(8736.81),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "baf91e50-6ef8-490a-94b4-75f16f56d385",
                                Name: codataccounting.String("Earl Schoen"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c631b99e-26ce-4d8f-9fdb-9410f63bbf81",
                                Name: codataccounting.String("Kay Frami"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("aperiam"),
                            ID: "1afdd788-6241-489e-b448-73f5033f19db",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "25ce4152-eab9-4cd7-a522-4a6a0e123b78",
                            Name: codataccounting.String("Tamara Terry"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e1f67f3c-4cce-44b6-9769-6ff3c5747501",
                            Name: codataccounting.String("Stacy Kovacek"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4f51f8b0-84c3-4197-a193-a245467f9487",
                            Name: codataccounting.String("Rachael Corkery"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cc497223-3e66-4bd8-be5d-00b979ef2038",
                            Name: codataccounting.String("Mrs. Gladys Collins"),
                        },
                    },
                    UnitAmount: 155.63,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ccc10964-0031-43b3-a504-4f65fe72dc40"),
                        Name: codataccounting.String("Miss Constance Shanahan"),
                    },
                    Description: codataccounting.String("adipisci"),
                    DiscountAmount: codataccounting.Float64(9504.86),
                    DiscountPercentage: codataccounting.Float64(2527.17),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "08efc15c-eb4d-46e1-aae0-f75aedf2acab",
                        Name: codataccounting.String("Tracey Rodriguez"),
                    },
                    Quantity: 912.7,
                    SubTotal: codataccounting.Float64(7863.19),
                    TaxAmount: codataccounting.Float64(6004.71),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1535.13),
                        ID: codataccounting.String("6ddb5894-61e7-4421-8be6-d9502f0ea930"),
                        Name: codataccounting.String("Hector Mayer"),
                    },
                    TotalAmount: codataccounting.Float64(6813.31),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2f72f885-0090-4491-9608-207888ec6618",
                                Name: codataccounting.String("Latoya West"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "659eb40e-c16f-4af7-9b0b-532a4da37cba",
                                Name: codataccounting.String("Tommie Gutkowski"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2c4842c9-b2ad-432d-afe8-1a88f4444573",
                                Name: codataccounting.String("Alonzo Schmidt"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7353f63c-8209-4379-aa69-cd5fbcf79da1",
                                Name: codataccounting.String("Blake Kuhic"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eos"),
                            ID: "bf95894e-6861-4adb-95f9-e5d751c9fe8f",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "02bfdc34-5084-41f1-b644-56379f3fb27e",
                            Name: codataccounting.String("Stephanie Yundt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "657b36fc-6b9f-4587-8e52-5c67641a8312",
                            Name: codataccounting.String("Jerome Berge"),
                        },
                    },
                    UnitAmount: 7303.7,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("cumque"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7678.8),
                        TotalAmount: codataccounting.Float64(7149.33),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("423abcdc-91fa-4abd-988e-71f6c48252d7"),
                            Name: codataccounting.String("Delores Bosco"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8719.69),
                        ID: codataccounting.String("074009ef-8d29-4de1-9d70-97b5da08c57f"),
                        Note: codataccounting.String("id"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("cumque"),
                        TotalAmount: codataccounting.Float64(4846.96),
                    },
                },
            },
            RemainingCredit: 5395.02,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: 991.6,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "accusamus": map[string]interface{}{
                        "excepturi": "harum",
                    },
                    "laborum": map[string]interface{}{
                        "repudiandae": "minus",
                        "officia": "laboriosam",
                        "illo": "cupiditate",
                        "veritatis": "aliquam",
                    },
                },
            },
            TotalAmount: 5682.31,
            TotalDiscount: 5410.46,
            TotalTaxAmount: 1166.65,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 428.84,
                    Name: "Sam Gerlach",
                },
                shared.WithholdingTaxitems{
                    Amount: 5497.1,
                    Name: "Erick Bernhard DVM",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(997047),
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
        CreditNoteID: "voluptatem",
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
        Query: codataccounting.String("dolor"),
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
            AdditionalTaxAmount: codataccounting.Float64(7194.5),
            AdditionalTaxPercentage: codataccounting.Float64(3125.63),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("neque"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(9281.02),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("numquam"),
                ID: "aa868555-9667-432a-a5dc-b6682cb70f8c",
            },
            DiscountPercentage: 9497.96,
            ID: codataccounting.String("d5fb6e91-b9a9-4f74-846e-2c3309db0536"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e75ca006-f539-42c1-9a25-a8bf92f97428"),
                        Name: codataccounting.String("Al Mills"),
                    },
                    Description: codataccounting.String("hic"),
                    DiscountAmount: codataccounting.Float64(5458.54),
                    DiscountPercentage: codataccounting.Float64(7433.13),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "f8221125-359d-4983-87f7-a79cd72cd248",
                        Name: codataccounting.String("Krystal Pagac IV"),
                    },
                    Quantity: 1794.33,
                    SubTotal: codataccounting.Float64(6090.79),
                    TaxAmount: codataccounting.Float64(9701.08),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1847.74),
                        ID: codataccounting.String("ac41ef57-25f1-4169-ac1e-41d8a23c23e3"),
                        Name: codataccounting.String("Kristie Daugherty"),
                    },
                    TotalAmount: codataccounting.Float64(6505.38),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "a197f6de-9221-451f-a171-2099853e9f54",
                                Name: codataccounting.String("Meredith Langosh"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "439ee224-4604-443b-8154-188c2f56e85d",
                                Name: codataccounting.String("Jessie Larkin"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("officiis"),
                            ID: "abd617c3-b0d5-41a4-8bf0-1bad8706d460",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "bfbdc41f-f5d4-4e2a-a4fb-5cb35d17638f",
                            Name: codataccounting.String("Delia Schuppe"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "359ecc5c-b860-4f8c-9580-ba73810e4fe4",
                            Name: codataccounting.String("Megan Kling"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7cd3b1dd-3bbc-4e24-bb76-84eff50126d7",
                            Name: codataccounting.String("Vicky Wolf"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d0eb74b8-4219-453b-84bd-3c43159d33e5",
                            Name: codataccounting.String("Gordon Ernser Jr."),
                        },
                    },
                    UnitAmount: 1167.42,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("139863aa-41e6-4c31-8c2f-1fcb51c9a41f"),
                        Name: codataccounting.String("Rudy Waters"),
                    },
                    Description: codataccounting.String("quidem"),
                    DiscountAmount: codataccounting.Float64(8383.74),
                    DiscountPercentage: codataccounting.Float64(4379.79),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "95ee65e0-76cc-47ab-b616-ea5c71641934",
                        Name: codataccounting.String("Tracy Bahringer"),
                    },
                    Quantity: 8851.03,
                    SubTotal: codataccounting.Float64(252.1),
                    TaxAmount: codataccounting.Float64(5776.72),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8573.88),
                        ID: codataccounting.String("19d2fc2f-9e2e-4105-944b-935d237a72f9"),
                        Name: codataccounting.String("Mattie Gutkowski"),
                    },
                    TotalAmount: codataccounting.Float64(3849.39),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ed4aecb7-537c-4d92-a2c9-ff57491aabfa",
                                Name: codataccounting.String("Eula Kuhic DVM"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0ca4d456-ef10-431e-a899-f0c2001e22cd",
                                Name: codataccounting.String("June Schmeler III"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "84a184d7-6d97-41fc-820c-65b037bb8e0c",
                                Name: codataccounting.String("Byron MacGyver V"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("molestiae"),
                            ID: "e4de04af-28c5-4ddd-b46a-a1cfd6d828da",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "31911296-4664-45c1-981f-29042f569b7a",
                            Name: codataccounting.String("Randal Altenwerth"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "216cbe07-1bc1-463e-a79a-3b084da99257",
                            Name: codataccounting.String("Gary Gorczany"),
                        },
                    },
                    UnitAmount: 572.07,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("847a742d-8449-46cb-9eec-f6b99bc63562"),
                        Name: codataccounting.String("Jonathon Zboncak"),
                    },
                    Description: codataccounting.String("enim"),
                    DiscountAmount: codataccounting.Float64(3686.58),
                    DiscountPercentage: codataccounting.Float64(7835.08),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "294c060b-06a1-4287-b64e-ef6d0c6d6ed9",
                        Name: codataccounting.String("Arnold Dooley"),
                    },
                    Quantity: 4337.98,
                    SubTotal: codataccounting.Float64(2479.27),
                    TaxAmount: codataccounting.Float64(2855.44),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3419.34),
                        ID: codataccounting.String("71509a8e-870d-43c5-a1f9-c242c7b66a1f"),
                        Name: codataccounting.String("Michelle Schmeler"),
                    },
                    TotalAmount: codataccounting.Float64(8150.46),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5b671989-0f42-4a4b-b438-d85b260591d7",
                                Name: codataccounting.String("Geraldine Von"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2059c9c3-f567-4e0e-a527-65b1d62fcdac",
                                Name: codataccounting.String("Mr. Joe Wisozk"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "16ce2239-e8f2-45cd-8d19-d959f439e392",
                                Name: codataccounting.String("Gertrude Runte"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95f7aa2b-2411-4369-9d1e-6698fcc45962",
                                Name: codataccounting.String("Minnie Schneider"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quam"),
                            ID: "76763342-5403-48bf-b597-1e9819055738",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "edbac7fd-a395-494d-a6bc-2ae480632b99",
                            Name: codataccounting.String("Hazel Renner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "22063698-2855-43cb-9000-6bef4921ec20",
                            Name: codataccounting.String("Robin Rempel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9366ac8e-e0f2-4bf1-9588-d40d03f3deba",
                            Name: codataccounting.String("Faye Kreiger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3e90bc40-df86-48fd-9240-5cb331d492f4",
                            Name: codataccounting.String("Douglas Denesik"),
                        },
                    },
                    UnitAmount: 6941.48,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("saepe"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1087.18),
                        TotalAmount: codataccounting.Float64(9623.97),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8217978d-0acc-4a77-aeb7-b7021a52046b"),
                            Name: codataccounting.String("Michele Waelchi"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7063.51),
                        ID: codataccounting.String("0e67e094-fdfe-4d55-80ef-53a34a1b8fe9"),
                        Note: codataccounting.String("natus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("dolorem"),
                        TotalAmount: codataccounting.Float64(1243.81),
                    },
                },
            },
            RemainingCredit: 6824.3,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusVoid,
            SubTotal: 442.64,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "fugiat": map[string]interface{}{
                        "quis": "fuga",
                        "eveniet": "consequuntur",
                        "illum": "delectus",
                    },
                    "rerum": map[string]interface{}{
                        "perferendis": "maiores",
                        "harum": "ratione",
                    },
                },
            },
            TotalAmount: 5618.25,
            TotalDiscount: 4837.74,
            TotalTaxAmount: 2521.83,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5768.7,
                    Name: "Desiree Ferry",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "quis",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(431131),
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

