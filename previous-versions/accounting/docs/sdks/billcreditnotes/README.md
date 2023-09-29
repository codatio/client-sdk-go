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
* [UploadAttachment](#uploadattachment) - Upload bill credit note attachment

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
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: accounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: accounting.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("6384.24"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: accounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("innovative blue"),
                    },
                    Description: accounting.String("Vision-oriented responsive function"),
                    DiscountAmount: types.MustNewDecimalFromString("9510.62"),
                    DiscountPercentage: types.MustNewDecimalFromString("8915.1"),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("deposit"),
                    },
                    Quantity: types.MustNewDecimalFromString("3015.1"),
                    SubTotal: types.MustNewDecimalFromString("899.64"),
                    TaxAmount: types.MustNewDecimalFromString("7150.4"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("7926.2"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("Gasoline Screen mobile"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6562.56"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: accounting.String("Durham after"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("Fay - Durgan"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "<ID>",
                            Name: accounting.String("Fiat"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("Grocery Borders Northwest"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("6519.85"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: accounting.String("EUR"),
                        CurrencyRate: types.MustNewDecimalFromString("365.21"),
                        TotalAmount: types.MustNewDecimalFromString("8424.64"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: accounting.String("<ID>"),
                            Name: accounting.String("though East"),
                        },
                        Currency: accounting.String("GBP"),
                        CurrencyRate: types.MustNewDecimalFromString("155.52"),
                        ID: accounting.String("<ID>"),
                        Note: accounting.String("array Edinburg Investor"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: accounting.String("likewise payment 1080p"),
                        TotalAmount: types.MustNewDecimalFromString("2597.72"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
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
                SupplierName: accounting.String("Toyota Neptunium round"),
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
        TimeoutInMinutes: accounting.Int(863813),
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
    res, err := s.BillCreditNotes.List(ctx, operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: accounting.String("-modifiedDate"),
        Page: accounting.Int(1),
        PageSize: accounting.Int(100),
        Query: accounting.String("Northeast Metal Canada"),
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
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: accounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: accounting.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("5971.29"),
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            ID: accounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("male Metal"),
                    },
                    Description: accounting.String("Visionary bi-directional analyzer"),
                    DiscountAmount: types.MustNewDecimalFromString("2782.81"),
                    DiscountPercentage: types.MustNewDecimalFromString("8965.01"),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("withdrawal extend"),
                    },
                    Quantity: types.MustNewDecimalFromString("2494.4"),
                    SubTotal: types.MustNewDecimalFromString("3668.07"),
                    TaxAmount: types.MustNewDecimalFromString("1395.79"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("6447.13"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("syndicate East Baht"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6298.17"),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "<ID>",
                                Name: accounting.String("guestbook driver users"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: accounting.String("Schroeder - Nienow"),
                            ID: "<ID>",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "<ID>",
                            Name: accounting.String("Wooden Internal"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("Dodge brightly"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7115.64"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
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
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
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
                SupplierName: accounting.String("quisquam"),
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
        ForceUpdate: accounting.Bool(false),
        TimeoutInMinutes: accounting.Int(639383),
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
    res, err := s.BillCreditNotes.UploadAttachment(ctx, operations.UploadBillCreditNoteAttachmentRequest{
        RequestBody: &operations.UploadBillCreditNoteAttachmentRequestBody{
            Content: []byte("v/ghW&IC$x"),
            RequestBody: "Elegant Producer Electric",
        },
        BillCreditNoteID: "Iowa Bentley",
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

