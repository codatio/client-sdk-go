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
            CurrencyRate: types.MustNewDecimalFromString("3927.85"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: syncforpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("d151a05d-fc2d-4df7-8c78-ca1ba928fc81"),
                        Name: syncforpayables.String("Tanya Gleason"),
                    },
                    Description: syncforpayables.String("cum"),
                    DiscountAmount: types.MustNewDecimalFromString("4561.5"),
                    DiscountPercentage: types.MustNewDecimalFromString("2165.5"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "92059293-96fe-4a75-96eb-10faaa2352c5",
                        Name: syncforpayables.String("Corey Hane III"),
                    },
                    Quantity: types.MustNewDecimalFromString("6342.74"),
                    SubTotal: types.MustNewDecimalFromString("9883.74"),
                    TaxAmount: types.MustNewDecimalFromString("9589.5"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("1020.44"),
                        ID: syncforpayables.String("a3a2fa94-6773-4925-9aa5-2c3f5ad019da"),
                        Name: syncforpayables.String("Johanna Wolf"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("5096.24"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f097b007-4f15-4471-b5e6-e13b99d488e1",
                                Name: syncforpayables.String("Kirk Boehm"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: syncforpayables.String("enim"),
                            ID: "0ad2abd4-4269-4802-9502-a94bb4f63c96",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "9a3efa77-dfb1-44cd-a6ae-395efb9ba88f",
                            Name: syncforpayables.String("Sandy Huels"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "97074ba4-469b-46e2-9419-59890afa563e",
                            Name: syncforpayables.String("Vivian Boyle"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("8919.24"),
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
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("7038.89"),
                        TotalAmount: types.MustNewDecimalFromString("4479.26"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("11e5b7fd-2ed0-4289-a1cd-dc692601fb57"),
                            Name: syncforpayables.String("Candice Beatty"),
                        },
                        Currency: syncforpayables.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("166.27"),
                        ID: syncforpayables.String("d30c5fbb-2587-4053-a02c-73d5fe9b90c2"),
                        Note: syncforpayables.String("blanditiis"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("eaque"),
                        TotalAmount: types.MustNewDecimalFromString("5772.29"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "adipisci": map[string]interface{}{
                        "asperiores": "earum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "49a8d9cb-f486-4333-a3f9-b77f3a410067",
                SupplierName: syncforpayables.String("quaerat"),
            },
            TotalAmount: types.MustNewDecimalFromString("805.78"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("8810.05"),
                    Name: "Jan Hodkiewicz",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforpayables.Int(542499),
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
        BillCreditNoteID: "sit",
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
        Query: syncforpayables.String("fugiat"),
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
            Currency: syncforpayables.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("6793.93"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: syncforpayables.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("7a89ebf7-37ae-4420-bce5-e6a95d8a0d44"),
                        Name: syncforpayables.String("Bernadette Torp"),
                    },
                    Description: syncforpayables.String("a"),
                    DiscountAmount: types.MustNewDecimalFromString("4561.3"),
                    DiscountPercentage: types.MustNewDecimalFromString("6874.88"),
                    ItemRef: &shared.BillCreditNoteLineItemItemReference{
                        ID: "73cf3be4-53f8-470b-b26b-5a73429cdb1a",
                        Name: syncforpayables.String("Randall Cole"),
                    },
                    Quantity: types.MustNewDecimalFromString("7044.74"),
                    SubTotal: types.MustNewDecimalFromString("3960.6"),
                    TaxAmount: types.MustNewDecimalFromString("4631.5"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("5654.21"),
                        ID: syncforpayables.String("d2322715-bf0c-4bb1-a31b-8b90f3443a11"),
                        Name: syncforpayables.String("Miss Billie Ward"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("7851.53"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f4b92187-9fce-4953-b73e-f7fbc7abd74d",
                                Name: syncforpayables.String("Earl Mosciski DVM"),
                            },
                        },
                        CustomerRef: &shared.BillCreditNoteLineItemTrackingCustomerRef{
                            CompanyName: syncforpayables.String("exercitationem"),
                            ID: "d2cff7c7-0a45-4626-9436-813f16d9f5fc",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.BillCreditNoteLineItemTrackingProjectReference{
                            ID: "c556146c-3e25-40fb-808c-42e141aac366",
                            Name: syncforpayables.String("Clifton Simonis"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b1442907-4747-478a-bbd4-66d28c10ab3c",
                            Name: syncforpayables.String("Salvatore Parker"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("3738.13"),
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
                        CurrencyRate: types.MustNewDecimalFromString("2728.22"),
                        TotalAmount: types.MustNewDecimalFromString("8920.5"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: syncforpayables.String("523c7e0b-c717-48e4-b96f-2a70c688282a"),
                            Name: syncforpayables.String("Randall Lindgren"),
                        },
                        Currency: syncforpayables.String("USD"),
                        CurrencyRate: types.MustNewDecimalFromString("1470.14"),
                        ID: syncforpayables.String("f222e981-7ee1-47cb-a61e-6b7b95bc0ab3"),
                        Note: syncforpayables.String("cumque"),
                        PaidOnDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
                        Reference: syncforpayables.String("consequatur"),
                        TotalAmount: types.MustNewDecimalFromString("7963.92"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: types.MustNewDecimalFromString("805.78"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sapiente": map[string]interface{}{
                        "consectetur": "esse",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "89fd871f-99dd-42ef-9121-aa6f1e674bdb",
                SupplierName: syncforpayables.String("accusantium"),
            },
            TotalAmount: types.MustNewDecimalFromString("805.78"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: types.MustNewDecimalFromString("3069.86"),
                    Name: "Samuel Hermiston",
                },
            },
        },
        BillCreditNoteID: "nisi",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: syncforpayables.Bool(false),
        TimeoutInMinutes: syncforpayables.Int(16328),
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

