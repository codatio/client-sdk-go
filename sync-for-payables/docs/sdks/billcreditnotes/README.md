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
	"github.com/ericlagergren/decimal"
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
            CurrencyRate: types.MustNewDecimalFromString("3834.41"),
            DiscountPercentage: *types.MustNewDecimalFromString("0"),
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("cc8796ed-151a-405d-bc2d-df7cc78ca1ba"),
                        Name: codatsyncpayables.String("Wayne Lind"),
                    },
                    Description: codatsyncpayables.String("totam"),
                    DiscountAmount: types.MustNewDecimalFromString("1059.07"),
                    DiscountPercentage: types.MustNewDecimalFromString("4146.62"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "742cb739-2059-4293-96fe-a7596eb10faa",
                        Name: codatsyncpayables.String("Ernest Ebert"),
                    },
                    Quantity: *types.MustNewDecimalFromString("7506.86"),
                    SubTotal: types.MustNewDecimalFromString("3154.28"),
                    TaxAmount: types.MustNewDecimalFromString("6078.31"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("3637.11"),
                        ID: codatsyncpayables.String("5907aff1-a3a2-4fa9-8677-39251aa52c3f"),
                        Name: codatsyncpayables.String("Mr. Alberta Schuster"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("8379.45"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "a1ffe78f-097b-4007-8f15-471b5e6e13b9",
                                Name: codatsyncpayables.String("Ervin Gleason"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("voluptates"),
                            ID: "1e91e450-ad2a-4bd4-8269-802d502a94bb",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "63c969e9-a3ef-4a77-9fb1-4cd66ae395ef",
                            Name: codatsyncpayables.String("Rene Reinger"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8f3a6699-7074-4ba4-869b-6e2141959890",
                            Name: codatsyncpayables.String("Abel O'Hara"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("2212.62"),
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
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("972.6"),
                        TotalAmount: types.MustNewDecimalFromString("4358.65"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("fe4c8b71-1e5b-47fd-aed0-28921cddc692"),
                            Name: codatsyncpayables.String("Donna Bernhard"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("4535.43"),
                        ID: codatsyncpayables.String("6b0d5f0d-30c5-4fbb-a587-053202c73d5f"),
                        Note: codatsyncpayables.String("recusandae"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("facilis"),
                        TotalAmount: types.MustNewDecimalFromString("5966.56"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: *types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "porro": map[string]interface{}{
                        "consequuntur": "blanditiis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "909b3fe4-9a8d-49cb-b486-33323f9b77f3",
                SupplierName: codatsyncpayables.String("dolorum"),
            },
            TotalAmount: *types.MustNewDecimalFromString("805.78"),
            TotalDiscount: *types.MustNewDecimalFromString("0"),
            TotalTaxAmount: *types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: *types.MustNewDecimalFromString("2543.56"),
                    Name: "Melissa Bednar",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(311796),
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
        BillCreditNoteID: "accusamus",
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
        Query: codatsyncpayables.String("quidem"),
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
	"github.com/ericlagergren/decimal"
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
            CurrencyRate: types.MustNewDecimalFromString("6176.58"),
            DiscountPercentage: *types.MustNewDecimalFromString("0"),
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("80d1ba77-a89e-4bf7-b7ae-4203ce5e6a95"),
                        Name: codatsyncpayables.String("Dr. Jimmie Murphy"),
                    },
                    Description: codatsyncpayables.String("tempora"),
                    DiscountAmount: types.MustNewDecimalFromString("4254.51"),
                    DiscountPercentage: types.MustNewDecimalFromString("7980.47"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "e2af7a73-cf3b-4e45-bf87-0b326b5a7342",
                        Name: codatsyncpayables.String("Simon Stracke MD"),
                    },
                    Quantity: *types.MustNewDecimalFromString("5173.79"),
                    SubTotal: types.MustNewDecimalFromString("2768.94"),
                    TaxAmount: types.MustNewDecimalFromString("1320.68"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("1749.09"),
                        ID: codatsyncpayables.String("bb679d23-2271-45bf-8cbb-1e31b8b90f34"),
                        Name: codatsyncpayables.String("Mr. Josephine Pagac V"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("9295.3"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0adcf4b9-2187-49fc-a953-f73ef7fbc7ab",
                                Name: codatsyncpayables.String("Allan Greenholt"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("sequi"),
                            ID: "9c0f5d2c-ff7c-470a-8562-6d436813f16d",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "5fce6c55-6146-4c3e-a50f-b008c42e141a",
                            Name: codatsyncpayables.String("Clark Franecki"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c8dd6b14-4290-4747-8778-a7bd466d28c1",
                            Name: codatsyncpayables.String("Amelia Predovic"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("8472.76"),
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
                        CurrencyRate: types.MustNewDecimalFromString("1783.67"),
                        TotalAmount: types.MustNewDecimalFromString("3738.13"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("1904e523-c7e0-4bc7-978e-4796f2a70c68"),
                            Name: codatsyncpayables.String("Eugene Leuschke"),
                        },
                        Currency: codatsyncpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("2775.96"),
                        ID: codatsyncpayables.String("82562f22-2e98-417e-a17c-be61e6b7b95b"),
                        Note: codatsyncpayables.String("eligendi"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("culpa"),
                        TotalAmount: types.MustNewDecimalFromString("7313.98"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: *types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "cumque": map[string]interface{}{
                        "consequuntur": "consequatur",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "c4f3789f-d871-4f99-9d2e-fd121aa6f1e6",
                SupplierName: codatsyncpayables.String("in"),
            },
            TotalAmount: *types.MustNewDecimalFromString("805.78"),
            TotalDiscount: *types.MustNewDecimalFromString("0"),
            TotalTaxAmount: *types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: *types.MustNewDecimalFromString("2586.84"),
                    Name: "Mrs. Gilberto Roberts",
                },
            },
        },
        BillCreditNoteID: "dicta",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        TimeoutInMinutes: codatsyncpayables.Int(355369),
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

