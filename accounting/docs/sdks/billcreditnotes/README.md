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
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(8289.4),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2a94bb4f-63c9-469e-9a3e-fa77dfb14cd6"),
                        Name: codataccounting.String("Kayla Thompson"),
                    },
                    Description: codataccounting.String("enim"),
                    DiscountAmount: codataccounting.Float64(8817.36),
                    DiscountPercentage: codataccounting.Float64(9654.17),
                    ItemRef: &shared.ItemRef{
                        ID: "b9ba88f3-a669-4970-b4ba-4469b6e21419",
                        Name: codataccounting.String("Ramona Lueilwitz MD"),
                    },
                    Quantity: 9689.62,
                    SubTotal: codataccounting.Float64(6521.03),
                    TaxAmount: codataccounting.Float64(3209.97),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4314.18),
                        ID: codataccounting.String("3e2516fe-4c8b-4711-a5b7-fd2ed028921c"),
                        Name: codataccounting.String("Ervin Schoen"),
                    },
                    TotalAmount: codataccounting.Float64(1399.72),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "01fb576b-0d5f-40d3-8c5f-bb2587053202",
                                Name: codataccounting.String("Darryl Fadel"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fe9b90c2-8909-4b3f-a49a-8d9cbf486333",
                                Name: codataccounting.String("Tiffany Welch"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptate"),
                            ID: "7f3a4100-674e-4bf6-9280-d1ba77a89ebf",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7ae4203c-e5e6-4a95-98a0-d446ce2af7a7",
                            Name: codataccounting.String("Rosalie White"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "453f870b-326b-45a7-b429-cdb1a8422bb6",
                            Name: codataccounting.String("Felicia Spencer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "22715bf0-cbb1-4e31-b8b9-0f3443a1108e",
                            Name: codataccounting.String("Jodi Skiles"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4b921879-fce9-453f-b3ef-7fbc7abd74dd",
                            Name: codataccounting.String("Dr. Faye Rutherford"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d2cff7c7-0a45-4626-9436-813f16d9f5fc",
                            Name: codataccounting.String("Nathaniel Ryan"),
                        },
                    },
                    UnitAmount: 3994.99,
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
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2322.34),
                        TotalAmount: codataccounting.Float64(9262.13),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("250fb008-c42e-4141-aac3-66c8dd6b1442"),
                            Name: codataccounting.String("Jose Kreiger"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4585.15),
                        ID: codataccounting.String("78a7bd46-6d28-4c10-ab3c-dca4251904e5"),
                        Note: codataccounting.String("aspernatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quo"),
                        TotalAmount: codataccounting.Float64(4598.56),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(7151.79),
                        TotalAmount: codataccounting.Float64(7997.96),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7178e479-6f2a-470c-a882-82aa482562f2"),
                            Name: codataccounting.String("Rose Turner"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4569.11),
                        ID: codataccounting.String("ee17cbe6-1e6b-47b9-9bc0-ab3c20c4f378"),
                        Note: codataccounting.String("provident"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("nulla"),
                        TotalAmount: codataccounting.Float64(5578.11),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "a": map[string]interface{}{
                        "sint": "pariatur",
                        "possimus": "quia",
                        "eveniet": "asperiores",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "d121aa6f-1e67-44bd-b04f-15756082d68e",
                SupplierName: codataccounting.String("dolorum"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 6091.78,
                    Name: "Ms. Roger Strosin II",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(86532),
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
        BillCreditNoteID: "consectetur",
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
        Query: codataccounting.String("adipisci"),
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
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(330.74),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("86a18403-94c2-4607-9f93-f5f0642dac7a"),
                        Name: codataccounting.String("Vernon Bergnaum"),
                    },
                    Description: codataccounting.String("quod"),
                    DiscountAmount: codataccounting.Float64(2883.98),
                    DiscountPercentage: codataccounting.Float64(704.47),
                    ItemRef: &shared.ItemRef{
                        ID: "3aa63aae-8d67-4864-9bb6-75fd5e60b375",
                        Name: codataccounting.String("Carroll Gerhold"),
                    },
                    Quantity: 9689.72,
                    SubTotal: codataccounting.Float64(6971.42),
                    TaxAmount: codataccounting.Float64(9049.49),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8970.71),
                        ID: codataccounting.String("41f33317-fe35-4b60-ab1e-a426555ba3c2"),
                        Name: codataccounting.String("Harvey Gulgowski"),
                    },
                    TotalAmount: codataccounting.Float64(8391.89),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "3b88f3a8-d8f5-4c0b-af2f-b7b194a276b2",
                                Name: codataccounting.String("Geneva Bradtke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e1f08f42-94e3-4698-b447-f603e8b445e8",
                                Name: codataccounting.String("Della Muller"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("recusandae"),
                            ID: "fd20e457-e185-48b6-a89f-be3a5aa8e482",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0ab40750-88e5-4186-a065-e904f3b1194b",
                            Name: codataccounting.String("Cameron Reilly"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3a79f9df-e0ab-47da-8a50-ce187f86bc17",
                            Name: codataccounting.String("Angelina Jenkins"),
                        },
                    },
                    UnitAmount: 8872.65,
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
                        CurrencyRate: codataccounting.Float64(1334.61),
                        TotalAmount: codataccounting.Float64(4044.25),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f8d986e8-81ea-4d4f-8e10-12563f94e29e"),
                            Name: codataccounting.String("Arnold Ferry"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1458.7),
                        ID: codataccounting.String("a57a15be-3e06-4080-be2b-6e3ab8845f05"),
                        Note: codataccounting.String("perspiciatis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("mollitia"),
                        TotalAmount: codataccounting.Float64(3782.45),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(9702.22),
                        TotalAmount: codataccounting.Float64(1746.58),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a54a31e9-4764-4a3e-865e-7956f9251a5a"),
                            Name: codataccounting.String("Rufus Okuneva"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(9992.78),
                        ID: codataccounting.String("f57bfaad-4f9e-4fc1-b451-2c1032648dc2"),
                        Note: codataccounting.String("sapiente"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("dicta"),
                        TotalAmount: codataccounting.Float64(3251.18),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(5896.95),
                        TotalAmount: codataccounting.Float64(9364.69),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("bfd0e9fe-6c63-42ca-baed-0117996312fd"),
                            Name: codataccounting.String("Jeffrey Goldner"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4797.54),
                        ID: codataccounting.String("78ff61d0-1747-4636-8a15-db6a660659a1"),
                        Note: codataccounting.String("error"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("voluptates"),
                        TotalAmount: codataccounting.Float64(6534.21),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(3240.83),
                        TotalAmount: codataccounting.Float64(5369.23),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("51d6c645-b08b-4618-91ba-a0fe1ade008e"),
                            Name: codataccounting.String("Miranda Ledner"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1905.67),
                        ID: codataccounting.String("50d8cdb5-a341-4814-b010-421813d5208e"),
                        Note: codataccounting.String("impedit"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("esse"),
                        TotalAmount: codataccounting.Float64(8972.77),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nesciunt": map[string]interface{}{
                        "eum": "vel",
                        "voluptatum": "magnam",
                        "exercitationem": "ab",
                    },
                    "porro": map[string]interface{}{
                        "nobis": "laboriosam",
                        "recusandae": "consequuntur",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "05e16dea-b3fe-4c95-b8a6-4584273a8418",
                SupplierName: codataccounting.String("fugiat"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 3955.44,
                    Name: "Edith Beahan",
                },
            },
        },
        BillCreditNoteID: "soluta",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(3860),
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

