# CreditNotes
(*CreditNotes*)

## Overview

Access standardized Credit notes from linked accounting software.

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
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
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
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00Z"),
            CreditNoteNumber: accounting.String("09/03 17.15"),
            Currency: accounting.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("1"),
            CustomerRef: &shared.AccountingCustomerRef{
                ID: "80000002-1674552702",
            },
            DiscountPercentage: types.MustNewDecimalFromString("0"),
            IssueDate: accounting.String("2023-03-09T02:21:26.726327+00:00"),
            LineItems: []shared.CreditNoteLineItem{

            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00Z"),
            Note: accounting.String("credit note 20230309 17.15"),
            PaymentAllocations: []shared.PaymentAllocationItems{
                shared.PaymentAllocationItems{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00Z"),
                        Currency: accounting.String("EUR"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        Currency: accounting.String("GBP"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00Z"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("0"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CreditNoteStatusSubmitted,
            SubTotal: types.MustNewDecimalFromString("10.2"),
            TotalAmount: types.MustNewDecimalFromString("1.25"),
            TotalDiscount: types.MustNewDecimalFromString("0"),
            TotalTaxAmount: types.MustNewDecimalFromString("0"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateCreditNoteRequest](../../pkg/models/operations/createcreditnoterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateCreditNoteResponse](../../pkg/models/operations/createcreditnoteresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## Get

The *Get credit note* endpoint returns a single credit note for a given creditNoteId.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support getting a specific credit note.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
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
        CreditNoteID: "<value>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetCreditNoteRequest](../../pkg/models/operations/getcreditnoterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetCreditNoteResponse](../../pkg/models/operations/getcreditnoteresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


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
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetCreateUpdateCreditNotesModelRequest](../../pkg/models/operations/getcreateupdatecreditnotesmodelrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.GetCreateUpdateCreditNotesModelResponse](../../pkg/models/operations/getcreateupdatecreditnotesmodelresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## List

The *List credit notes* endpoint returns a list of [credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
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
        Query: accounting.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCreditNotesRequest](../../pkg/models/operations/listcreditnotesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListCreditNotesResponse](../../pkg/models/operations/listcreditnotesresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |


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
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
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
            AllocatedOnDate: accounting.String("2022-10-23T00:00:00Z"),
            Currency: accounting.String("GBP"),
            DiscountPercentage: types.MustNewDecimalFromString("5971.29"),
            IssueDate: accounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.CreditNoteLineItem{

            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00Z"),
            PaymentAllocations: []shared.PaymentAllocationItems{
                shared.PaymentAllocationItems{
                    Allocation: shared.Allocation{
                        AllocatedOnDate: accounting.String("2022-10-23T00:00:00Z"),
                        Currency: accounting.String("GBP"),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        Currency: accounting.String("EUR"),
                        PaidOnDate: accounting.String("2022-10-23T00:00:00Z"),
                    },
                },
            },
            RemainingCredit: types.MustNewDecimalFromString("7084.55"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: types.MustNewDecimalFromString("6276.9"),
            TotalAmount: types.MustNewDecimalFromString("6841.99"),
            TotalDiscount: types.MustNewDecimalFromString("4888.52"),
            TotalTaxAmount: types.MustNewDecimalFromString("4430.76"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "<value>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateCreditNoteRequest](../../pkg/models/operations/updatecreditnoterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateCreditNoteResponse](../../pkg/models/operations/updatecreditnoteresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
