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
            Currency: codatsyncpayables.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("3843.82"),
            DiscountPercentage: *types.MustNewDecimalFromString("0"),
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("4e0f467c-c879-46ed-951a-05dfc2ddf7cc"),
                        Name: codatsyncpayables.String("Deanna Sauer MD"),
                    },
                    Description: codatsyncpayables.String("officia"),
                    DiscountAmount: types.MustNewDecimalFromString("5820.2"),
                    DiscountPercentage: types.MustNewDecimalFromString("1433.53"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "8fc81674-2cb7-4392-8592-9396fea7596e",
                        Name: codatsyncpayables.String("Roger Beier"),
                    },
                    Quantity: *types.MustNewDecimalFromString("6531.4"),
                    SubTotal: types.MustNewDecimalFromString("6706.38"),
                    TaxAmount: types.MustNewDecimalFromString("1709.09"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("2103.82"),
                        ID: codatsyncpayables.String("52c59559-07af-4f1a-ba2f-a9467739251a"),
                        Name: codatsyncpayables.String("Bill Conn"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("9495.72"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5ad019da-1ffe-478f-897b-0074f15471b5",
                                Name: codatsyncpayables.String("Mrs. Leslie VonRueden"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("molestias"),
                            ID: "9d488e1e-91e4-450a-92ab-d44269802d50",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "94bb4f63-c969-4e9a-befa-77dfb14cd66a",
                            Name: codatsyncpayables.String("Alfred McClure"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fb9ba88f-3a66-4997-874b-a4469b6e2141",
                            Name: codatsyncpayables.String("Derrick McLaughlin"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("336.25"),
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
                        CurrencyRate: types.MustNewDecimalFromString("3209.97"),
                        TotalAmount: types.MustNewDecimalFromString("4314.18"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("3e2516fe-4c8b-4711-a5b7-fd2ed028921c"),
                            Name: codatsyncpayables.String("Ervin Schoen"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("4071.83"),
                        ID: codatsyncpayables.String("01fb576b-0d5f-40d3-8c5f-bb2587053202"),
                        Note: codatsyncpayables.String("minus"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("dolor"),
                        TotalAmount: types.MustNewDecimalFromString("8745.73"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: *types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "hic": map[string]interface{}{
                        "recusandae": "omnis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "b90c2890-9b3f-4e49-a8d9-cbf48633323f",
                SupplierName: codatsyncpayables.String("excepturi"),
            },
            TotalAmount: *types.MustNewDecimalFromString("805.78"),
            TotalDiscount: *types.MustNewDecimalFromString("0"),
            TotalTaxAmount: *types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: *types.MustNewDecimalFromString("7395.51"),
                    Name: "Marian Wisozk",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(254356),
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
        BillCreditNoteID: "veritatis",
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
        Query: codatsyncpayables.String("ipsa"),
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
            CurrencyRate: types.MustNewDecimalFromString("4878.38"),
            DiscountPercentage: *types.MustNewDecimalFromString("0"),
            ID: codatsyncpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("ebf69280-d1ba-477a-89eb-f737ae4203ce"),
                        Name: codatsyncpayables.String("Jenna Hoppe"),
                    },
                    Description: codatsyncpayables.String("minima"),
                    DiscountAmount: types.MustNewDecimalFromString("8310.49"),
                    DiscountPercentage: types.MustNewDecimalFromString("5197.11"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "a0d446ce-2af7-4a73-8f3b-e453f870b326",
                        Name: codatsyncpayables.String("Glen Oberbrunner"),
                    },
                    Quantity: *types.MustNewDecimalFromString("2776.28"),
                    SubTotal: types.MustNewDecimalFromString("1864.58"),
                    TaxAmount: types.MustNewDecimalFromString("5867.84"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("8075.81"),
                        ID: codatsyncpayables.String("db1a8422-bb67-49d2-b227-15bf0cbb1e31"),
                        Name: codatsyncpayables.String("Isaac Reynolds DVM"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("2091.57"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "443a1108-e0ad-4cf4-b921-879fce953f73",
                                Name: codatsyncpayables.String("Roman Kulas"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: codatsyncpayables.String("quod"),
                            ID: "7abd74dd-39c0-4f5d-acff-7c70a45626d4",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "813f16d9-f5fc-4e6c-9561-46c3e250fb00",
                            Name: codatsyncpayables.String("Myron Haag"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "141aac36-6c8d-4d6b-9442-907474778a7b",
                            Name: codatsyncpayables.String("Bernard Kerluke"),
                        },
                    },
                    UnitAmount: *types.MustNewDecimalFromString("1811.51"),
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
                        CurrencyRate: types.MustNewDecimalFromString("568.48"),
                        TotalAmount: types.MustNewDecimalFromString("6600.4"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsyncpayables.String("b3cdca42-5190-44e5-a3c7-e0bc7178e479"),
                            Name: codatsyncpayables.String("Miranda Daniel"),
                        },
                        Currency: codatsyncpayables.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("7836.48"),
                        ID: codatsyncpayables.String("688282aa-4825-462f-a22e-9817ee17cbe6"),
                        Note: codatsyncpayables.String("quasi"),
                        PaidOnDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsyncpayables.String("vel"),
                        TotalAmount: types.MustNewDecimalFromString("6900.25"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: *types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "rerum": map[string]interface{}{
                        "occaecati": "minima",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "bc0ab3c2-0c4f-4378-9fd8-71f99dd2efd1",
                SupplierName: codatsyncpayables.String("consequuntur"),
            },
            TotalAmount: *types.MustNewDecimalFromString("805.78"),
            TotalDiscount: *types.MustNewDecimalFromString("0"),
            TotalTaxAmount: *types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: *types.MustNewDecimalFromString("944.58"),
                    Name: "Shannon Jacobi DVM",
                },
            },
        },
        BillCreditNoteID: "vel",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        TimeoutInMinutes: codatsyncpayables.Int(447378),
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

