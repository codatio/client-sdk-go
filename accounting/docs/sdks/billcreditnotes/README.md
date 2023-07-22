# BillCreditNotes

## Overview

Bill credit notes

### Available Operations

* [Create](#create) - Create bill credit note
* [Get](#get) - Get bill credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bill credit note model
* [List](#list) - List bill credit notes
* [Update](#update) - Update bill credit note
* [UploadAttachment](#uploadattachment) - Push invoice attachment

## Create

The *Create bill credit note* endpoint creates a new [bill credit note](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating an account.


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
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(9654.17),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ba88f3a6-6997-4074-ba44-69b6e2141959"),
                        Name: codataccounting.String("Kirk Bartoletti"),
                    },
                    Description: codataccounting.String("mollitia"),
                    DiscountAmount: codataccounting.Float64(3209.97),
                    DiscountPercentage: codataccounting.Float64(4314.18),
                    ItemRef: &shared.ItemRef{
                        ID: "3e2516fe-4c8b-4711-a5b7-fd2ed028921c",
                        Name: codataccounting.String("Ervin Schoen"),
                    },
                    Quantity: 1399.72,
                    SubTotal: codataccounting.Float64(4071.83),
                    TaxAmount: codataccounting.Float64(332.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(691.67),
                        ID: codataccounting.String("fb576b0d-5f0d-430c-9fbb-2587053202c7"),
                        Name: codataccounting.String("Eula Hegmann"),
                    },
                    TotalAmount: codataccounting.Float64(6082.53),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "90c28909-b3fe-449a-8d9c-bf48633323f9",
                                Name: codataccounting.String("Adrian Kuhn"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a4100674-ebf6-4928-8d1b-a77a89ebf737",
                                Name: codataccounting.String("Elbert Gislason I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ce5e6a95-d8a0-4d44-ace2-af7a73cf3be4",
                                Name: codataccounting.String("Florence Will"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("sit"),
                            ID: "b326b5a7-3429-4cdb-9a84-22bb679d2322",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "5bf0cbb1-e31b-48b9-8f34-43a1108e0adc",
                            Name: codataccounting.String("Alexander Prosacco"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "879fce95-3f73-4ef7-bbc7-abd74dd39c0f",
                            Name: codataccounting.String("Freda Cormier"),
                        },
                    },
                    UnitAmount: 9850.33,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7c70a456-26d4-4368-93f1-6d9f5fce6c55"),
                        Name: codataccounting.String("Stephanie Gutkowski"),
                    },
                    Description: codataccounting.String("consectetur"),
                    DiscountAmount: codataccounting.Float64(9262.13),
                    DiscountPercentage: codataccounting.Float64(1324.87),
                    ItemRef: &shared.ItemRef{
                        ID: "50fb008c-42e1-441a-ac36-6c8dd6b14429",
                        Name: codataccounting.String("Minnie Gutkowski"),
                    },
                    Quantity: 4585.15,
                    SubTotal: codataccounting.Float64(4561.41),
                    TaxAmount: codataccounting.Float64(5245.93),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6832.82),
                        ID: codataccounting.String("7bd466d2-8c10-4ab3-8dca-4251904e523c"),
                        Name: codataccounting.String("Sophie Bayer"),
                    },
                    TotalAmount: codataccounting.Float64(4908.19),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "78e4796f-2a70-4c68-8282-aa482562f222",
                                Name: codataccounting.String("Ms. Marion Little"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("accusamus"),
                            ID: "17cbe61e-6b7b-495b-80ab-3c20c4f3789f",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "71f99dd2-efd1-421a-a6f1-e674bdb04f15",
                            Name: codataccounting.String("Ms. Dana Huel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "68ea19f1-d170-4513-b9d0-8086a1840394",
                            Name: codataccounting.String("Ms. Benjamin Hirthe DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "93f5f064-2dac-47af-915c-c413aa63aae8",
                            Name: codataccounting.String("Chester Kuphal"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4dbb675f-d5e6-40b3-b5ed-4f6fbee41f33",
                            Name: codataccounting.String("Heather Kuhn"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "35b60eb1-ea42-4655-9ba3-c28744ed53b8",
                            Name: codataccounting.String("Moses Douglas"),
                        },
                    },
                    UnitAmount: 8672.9,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8f5c0b2f-2fb7-4b19-8a27-6b26916fe1f0"),
                        Name: codataccounting.String("Emilio Goodwin"),
                    },
                    Description: codataccounting.String("eius"),
                    DiscountAmount: codataccounting.Float64(8967.62),
                    DiscountPercentage: codataccounting.Float64(2155.29),
                    ItemRef: &shared.ItemRef{
                        ID: "698f447f-603e-48b4-85e8-0ca55efd20e4",
                        Name: codataccounting.String("Ms. Pearl Towne"),
                    },
                    Quantity: 5106.29,
                    SubTotal: codataccounting.Float64(7400.98),
                    TaxAmount: codataccounting.Float64(3868.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6805.15),
                        ID: codataccounting.String("89fbe3a5-aa8e-4482-8d0a-b4075088e518"),
                        Name: codataccounting.String("Jane Bailey"),
                    },
                    TotalAmount: codataccounting.Float64(9061.72),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "04f3b119-4b8a-4bf6-83a7-9f9dfe0ab7da",
                                Name: codataccounting.String("Miss Hubert Hartmann"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "187f86bc-173d-4689-aee9-526f8d986e88",
                                Name: codataccounting.String("Delia Parisian"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f0e10125-63f9-44e2-9e97-3e922a57a15b",
                                Name: codataccounting.String("Ms. Melvin Thiel IV"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quae"),
                            ID: "7e2b6e3a-b884-45f0-997a-60ff2a54a31e",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "764a3e86-5e79-456f-9251-a5a9da660ff5",
                            Name: codataccounting.String("Antoinette Wehner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4f9efc1b-4512-4c10-b264-8dc2f615199e",
                            Name: codataccounting.String("Dr. Terrell Stanton"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fe6c632c-a3ae-4d01-9799-6312fde04771",
                            Name: codataccounting.String("Tamara Lang"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "61d01747-6360-4a15-9b6a-660659a1adea",
                            Name: codataccounting.String("Wm Hane"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1d6c645b-08b6-4189-9baa-0fe1ade008e6",
                            Name: codataccounting.String("Ian Schinner"),
                        },
                    },
                    UnitAmount: 1905.67,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(7706.75),
                        TotalAmount: codataccounting.Float64(8427.77),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b5a34181-4301-4042-9813-d5208ece7e25"),
                            Name: codataccounting.String("Lula Kemmer"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(3494.4),
                        ID: codataccounting.String("1c6c6e20-5e16-4dea-b3fe-c9578a645842"),
                        Note: codataccounting.String("ducimus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("fuga"),
                        TotalAmount: codataccounting.Float64(5140.54),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "rem": map[string]interface{}{
                        "dicta": "nisi",
                        "consequuntur": "consectetur",
                        "aperiam": "cupiditate",
                        "reiciendis": "soluta",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "0929921a-efb9-4f58-84d8-6e68e4be0560",
                SupplierName: codataccounting.String("quasi"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9785.48,
                    Name: "Bobbie Stokes",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(364463),
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

The *Get bill credit note* endpoint returns a single bill credit note for a given billCreditNoteId.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support getting a specific bill credit note.

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
    res, err := s.BillCreditNotes.Get(ctx, operations.GetBillCreditNoteRequest{
        BillCreditNoteID: "reprehenderit",
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

The *Get create/update bill credit note model* endpoint returns the expected data for the request payload when creating and updating a [bill credit note](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) for a given company and integration.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating and updating a bill credit note.


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
    res, err := s.BillCreditNotes.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBillCreditNotesModelRequest{
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetCreateUpdateBillCreditNotesModelRequest](../../models/operations/getcreateupdatebillcreditnotesmodelrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |


### Response

**[*operations.GetCreateUpdateBillCreditNotesModelResponse](../../models/operations/getcreateupdatebillcreditnotesmodelresponse.md), error**


## List

The *List bill credit notes* endpoint returns a list of [bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

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
    res, err := s.BillCreditNotes.List(ctx, operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("est"),
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

The *Update bill credit note* endpoint updates an existing [bill credit note](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating an account.


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
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(8806.79),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ef66ef1c-aa33-483c-abeb-477373c8d72f"),
                        Name: codataccounting.String("Dr. Megan Spinka"),
                    },
                    Description: codataccounting.String("architecto"),
                    DiscountAmount: codataccounting.Float64(9754.25),
                    DiscountPercentage: codataccounting.Float64(1563.83),
                    ItemRef: &shared.ItemRef{
                        ID: "c4310661-e963-449e-9cf9-e06e3a437000",
                        Name: codataccounting.String("Clay Jast"),
                    },
                    Quantity: 7051.48,
                    SubTotal: codataccounting.Float64(8093.65),
                    TaxAmount: codataccounting.Float64(5960.65),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7090.36),
                        ID: codataccounting.String("8f759eac-55a9-4741-9311-352965bb8a72"),
                        Name: codataccounting.String("Mr. Anne Kautzer"),
                    },
                    TotalAmount: codataccounting.Float64(2082.53),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e139dbc2-259b-41ab-9a8c-070e1084cb06",
                                Name: codataccounting.String("Miss Wanda Shanahan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "879eeb96-65b8-45ef-bd02-bae0be2d7822",
                                Name: codataccounting.String("Natasha Wehner"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("similique"),
                            ID: "4b5197f9-2443-4da7-8e52-b895c537c645",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "fb0b3489-6c3c-4a5a-8fbe-2fd570757792",
                            Name: codataccounting.String("Peter Kuphal"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ac646ecb-5734-409e-beb1-e5a2b12eb07f",
                            Name: codataccounting.String("Joyce Howe"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "99545fc9-5fa8-4897-8e18-9dbb30fcb33e",
                            Name: codataccounting.String("Anthony Hayes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "197cd44e-2f52-4d82-9351-3bb6f48b656b",
                            Name: codataccounting.String("Carroll Purdy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ff2e4b27-537a-48cd-9e73-19c177d525f7",
                            Name: codataccounting.String("Mrs. Pam Bins"),
                        },
                    },
                    UnitAmount: 9025.81,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b52ff785-fc37-4814-94c9-8e0c2bb89eb7"),
                        Name: codataccounting.String("Cristina Murphy"),
                    },
                    Description: codataccounting.String("dolorem"),
                    DiscountAmount: codataccounting.Float64(4138.01),
                    DiscountPercentage: codataccounting.Float64(7712.26),
                    ItemRef: &shared.ItemRef{
                        ID: "600503d8-bb31-4180-b739-ae9e057eb809",
                        Name: codataccounting.String("Mr. Craig Leannon"),
                    },
                    Quantity: 2244.13,
                    SubTotal: codataccounting.Float64(1242.89),
                    TaxAmount: codataccounting.Float64(9536.76),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2232.91),
                        ID: codataccounting.String("981d4c70-0b60-47f3-893c-73b9da3f2ced"),
                        Name: codataccounting.String("Cody Terry"),
                    },
                    TotalAmount: codataccounting.Float64(9958.16),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "257411fa-f4b7-4544-a472-e802857a5b40",
                                Name: codataccounting.String("Natalie Dooley"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("fugiat"),
                            ID: "575f1400-e764-4ad7-b34e-c1b781b36a08",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8d100efa-da20-40ef-8422-eb2164cf9ab8",
                            Name: codataccounting.String("Beth Jenkins"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3ffda9e0-6bee-4482-9c1f-c0e115c80bff",
                            Name: codataccounting.String("Keith Lueilwitz"),
                        },
                    },
                    UnitAmount: 2662.84,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ec42defc-ce8f-4197-b773-e63562a7b408"),
                        Name: codataccounting.String("Steven Harris"),
                    },
                    Description: codataccounting.String("facere"),
                    DiscountAmount: codataccounting.Float64(3071.73),
                    DiscountPercentage: codataccounting.Float64(5525.81),
                    ItemRef: &shared.ItemRef{
                        ID: "fdaf313a-1f5f-4d94-a59c-0b36f25ea944",
                        Name: codataccounting.String("Travis Rempel"),
                    },
                    Quantity: 4031.47,
                    SubTotal: codataccounting.Float64(7917.62),
                    TaxAmount: codataccounting.Float64(688.8),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1081.65),
                        ID: codataccounting.String("f6c37a51-2624-4383-9bbc-05a23a45cefc"),
                        Name: codataccounting.String("Ora Shields Jr."),
                    },
                    TotalAmount: codataccounting.Float64(6339.82),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ce2169e5-1001-49c6-9c5e-34762799bfbb",
                                Name: codataccounting.String("Brent Mills"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("hic"),
                            ID: "b2bb4eca-e6c3-4d5d-b3ad-ebd5daea4c50",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8aa94c02-644c-4f5e-9d9a-4578adc1ac60",
                            Name: codataccounting.String("Ernestine Turcotte Jr."),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ac802e2e-c09f-4f8f-8f81-6ff3477c13e9",
                            Name: codataccounting.String("Mrs. Phyllis Russel Sr."),
                        },
                    },
                    UnitAmount: 3576.39,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b0960a66-8151-4a47-aaf9-23c5949f83f3"),
                        Name: codataccounting.String("Angela Schaefer"),
                    },
                    Description: codataccounting.String("molestiae"),
                    DiscountAmount: codataccounting.Float64(3956.34),
                    DiscountPercentage: codataccounting.Float64(9890.33),
                    ItemRef: &shared.ItemRef{
                        ID: "fb901c6e-cbb4-4e24-bcf7-89ffafeda53e",
                        Name: codataccounting.String("Brandi Turner"),
                    },
                    Quantity: 357.42,
                    SubTotal: codataccounting.Float64(6378.4),
                    TaxAmount: codataccounting.Float64(7701.28),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(978.96),
                        ID: codataccounting.String("84c2b9c2-47c8-4837-ba40-e1942f32e550"),
                        Name: codataccounting.String("Lynn Kovacek"),
                    },
                    TotalAmount: codataccounting.Float64(9468.08),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d56d0bd0-af2d-4fe1-bdb4-f62cba3f8941",
                                Name: codataccounting.String("Erick Pfeffer MD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "80a6924d-3b2e-4cfc-88f8-95010f5dd3d6",
                                Name: codataccounting.String("Shaun Bergnaum I"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptates"),
                            ID: "54c82f16-8a36-43c8-873e-484380b1f6b8",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "275a60a0-4c49-45cc-a991-71b51c1bdb1c",
                            Name: codataccounting.String("Leroy Ratke"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ebdfc4cc-ca99-4bc7-bc0b-2dce10873e42",
                            Name: codataccounting.String("Brian Bartell"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "678878ba-8581-4a58-a08c-54fefa9c95f2",
                            Name: codataccounting.String("Angelo Runolfsdottir"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "65d307cf-ee81-4206-a281-3fa4a41c480d",
                            Name: codataccounting.String("Mr. Kristie Cole"),
                        },
                    },
                    UnitAmount: 6854.67,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(449.29),
                        TotalAmount: codataccounting.Float64(1341.73),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d514f4cc-6f18-4bf9-a21a-6a4f77a87ee3"),
                            Name: codataccounting.String("Ray Prosacco"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1316.87),
                        ID: codataccounting.String("c65b3441-8e3b-4b91-88d9-75e0e8419d8f"),
                        Note: codataccounting.String("deleniti"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("earum"),
                        TotalAmount: codataccounting.Float64(1013.74),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "maiores": map[string]interface{}{
                        "saepe": "consequatur",
                    },
                    "esse": map[string]interface{}{
                        "facere": "quisquam",
                        "cumque": "aliquam",
                        "dolorum": "deserunt",
                        "ad": "reiciendis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "3cabd905-a972-4e05-a728-227b2d309470",
                SupplierName: codataccounting.String("distinctio"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4630.5,
                    Name: "Tom Wintheiser",
                },
                shared.WithholdingTaxitems{
                    Amount: 4826.28,
                    Name: "Jan Hirthe",
                },
                shared.WithholdingTaxitems{
                    Amount: 6750.58,
                    Name: "Ora Okuneva",
                },
                shared.WithholdingTaxitems{
                    Amount: 2525.67,
                    Name: "Preston Wyman DDS",
                },
            },
        },
        BillCreditNoteID: "sequi",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(125707),
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


## UploadAttachment

---
stoplight-id: c26f5b1b19168
---

The *Upload bill credit note attachment* endpoint uploads an attachment and assigns it against a specific `billCreditNoteId`.

[Bill Credit Notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/accounting-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support uploading a bill credit note attachment.


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
    res, err := s.BillCreditNotes.UploadAttachment(ctx, operations.UploadBillCreditNoteAttachmentRequest{
        RequestBody: &operations.UploadBillCreditNoteAttachmentRequestBody{
            Content: []byte("vitae"),
            RequestBody: "voluptatibus",
        },
        BillCreditNoteID: "doloremque",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.UploadBillCreditNoteAttachmentRequest](../../models/operations/uploadbillcreditnoteattachmentrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.UploadBillCreditNoteAttachmentResponse](../../models/operations/uploadbillcreditnoteattachmentresponse.md), error**

