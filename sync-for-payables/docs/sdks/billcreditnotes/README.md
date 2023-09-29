# BillCreditNotes
(*BillCreditNotes*)

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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/types"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: syncforpayables.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: syncforpayables.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("6384.24"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: syncforpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("innovative blue"),
                    },
                    Description: syncforpayables.String("Vision-oriented responsive function"),
                    DiscountAmount: types.MustNewDecimalFromString("9510.62"),
                    DiscountPercentage: types.MustNewDecimalFromString("8915.1"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "<ID>",
                        Name: syncforpayables.String("deposit"),
                    },
                    Quantity: types.MustNewDecimalFromString("3015.1"),
                    SubTotal: types.MustNewDecimalFromString("899.64"),
                    TaxAmount: types.MustNewDecimalFromString("7150.4"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("7926.2"),
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("Gasoline Screen mobile"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6562.56"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: syncforpayables.String("Durham after"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: syncforpayables.String("Fay - Durgan"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforpayables.String("Fiat"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: syncforpayables.String("Grocery Borders Northwest"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("6519.85"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayables.Bool(false),
            },
            ModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Note: syncforpayables.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforpayables.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("365.21"),
                        TotalAmount: types.MustNewDecimalFromString("8424.64"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("<ID>"),
                            Name: syncforpayables.String("though East"),
                        },
                        Currency: syncforpayables.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("155.52"),
                        ID: syncforpayables.String("<ID>"),
                        Note: syncforpayables.String("array Edinburg Investor"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("likewise payment 1080p"),
                        TotalAmount: types.MustNewDecimalFromString("2597.72"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "porro": map[string]interface{}{
                        "asperiores": "Indiana",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: syncforpayables.String("Toyota Neptunium round"),
            },
            TotalAmount: types.MustNewDecimalFromString("805.78"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("1406.49"),
                    Name: "meanwhile",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforpayables.Int(863813),
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Get(ctx, operations.GetBillCreditNoteRequest{
        BillCreditNoteID: "Northeast Hatchback Kia",
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.List(ctx, operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforpayables.String("-modifiedDate"),
        Page: syncforpayables.Int(1),
        PageSize: syncforpayables.Int(100),
        Query: syncforpayables.String("Northeast Metal Canada"),
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/types"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: syncforpayables.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: syncforpayables.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("5971.29"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: syncforpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("male Metal"),
                    },
                    Description: syncforpayables.String("Visionary bi-directional analyzer"),
                    DiscountAmount: types.MustNewDecimalFromString("2782.81"),
                    DiscountPercentage: types.MustNewDecimalFromString("8965.01"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "<ID>",
                        Name: syncforpayables.String("withdrawal extend"),
                    },
                    Quantity: types.MustNewDecimalFromString("2494.4"),
                    SubTotal: types.MustNewDecimalFromString("3668.07"),
                    TaxAmount: types.MustNewDecimalFromString("1395.79"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("6447.13"),
                        ID: syncforpayables.String("<ID>"),
                        Name: syncforpayables.String("syndicate East Baht"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6298.17"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: syncforpayables.String("guestbook driver users"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: syncforpayables.String("Schroeder - Nienow"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "<ID>",
                            Name: syncforpayables.String("Wooden Internal"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: syncforpayables.String("Dodge brightly"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7115.64"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayables.Bool(false),
            },
            ModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Note: syncforpayables.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Currency: syncforpayables.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("3297.12"),
                        TotalAmount: types.MustNewDecimalFromString("4939.56"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("<ID>"),
                            Name: syncforpayables.String("modulo Kia"),
                        },
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("8000.94"),
                        ID: syncforpayables.String("<ID>"),
                        Note: syncforpayables.String("Avon"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("Reactive Global Northeast"),
                        TotalAmount: types.MustNewDecimalFromString("6090.5"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "alias": map[string]interface{}{
                        "veritatis": "ADP",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: syncforpayables.String("quisquam"),
            },
            TotalAmount: types.MustNewDecimalFromString("805.78"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("8523.4"),
                    Name: "pascal Plastic",
                },
            },
        },
        BillCreditNoteID: "magenta",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: syncforpayables.Bool(false),
        TimeoutInMinutes: syncforpayables.Int(639383),
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

