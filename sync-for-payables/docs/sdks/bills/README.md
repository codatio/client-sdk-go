# Bills
(*Bills*)

## Overview

Get, create, and update Bills.

### Available Operations

* [GetBillOptions](#getbilloptions) - Get bill mapping options
* [List](#list) - List bills
* [Create](#create) - Create bill
* [UploadAttachment](#uploadattachment) - Upload bill attachment
* [ListAttachments](#listattachments) - List bill attachments
* [DownloadAttachment](#downloadattachment) - Download bill attachment

## GetBillOptions

﻿Use the *Get mapping options - Bills* endpoint to return a list of available mapping options for a given company's connection ID.

By default, this endpoint returns a list of active accounts and tax rates. You can use [querying](https://docs.codat.io/using-the-api/querying) to change that.

Mapping options are a set of accounts and tax rates used to configure the SMB's payables integration.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.GetBillOptions(ctx, operations.GetMappingOptionsBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ContinuationToken: syncforpayables.String("continuationToken=eyJwYWdlIjoyLCJwYWdlU2l6ZSI6MTAwLCJwYWdlQ291bnQiOjExfQ=="),
        StatusQuery: syncforpayables.String("status=Archived"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillMappingOptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetMappingOptionsBillsRequest](../../pkg/models/operations/getmappingoptionsbillsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetMappingOptionsBillsResponse](../../pkg/models/operations/getmappingoptionsbillsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## List

The *List bills* endpoint returns a list of [bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

By default, the endpoint will return all bills with a status of 'Open' & 'PartiallyPaid' to show all oustanding bills.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.List(ctx, operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ContinuationToken: syncforpayables.String("continuationToken=eyJwYWdlIjoyLCJwYWdlU2l6ZSI6MTAwLCJwYWdlQ291bnQiOjExfQ=="),
        Query: syncforpayables.String("status=Open"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Bills != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListBillsRequest](../../pkg/models/operations/listbillsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListBillsResponse](../../pkg/models/operations/listbillsresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |

## Create

The *Create bill* endpoint creates a new [bill](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) for a given company's connection.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.Create(ctx, operations.CreateBillRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BillPrototype: &shared.BillPrototype{
            Reference: syncforpayables.String("bill_b8qmmj4ksf1suax"),
            SupplierRef: shared.SupplierRef{
                ID: "1262c350-fe0f-40ec-aeff-41c95b4a45af",
                SupplierName: syncforpayables.String("DIISR - Small Business Services"),
            },
            IssueDate: "2023-04-23T00:00:00",
            DueDate: "2023-04-23T00:00:00",
            Currency: "GBP",
            CurrencyRate: types.MustNewDecimalFromString("1"),
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    Description: syncforpayables.String("Half day training - Microsoft Office"),
                    UnitAmount: types.MustNewDecimalFromString("1800"),
                    Quantity: types.MustNewDecimalFromString("1"),
                    TaxAmount: types.MustNewDecimalFromString("360"),
                    AccountRef: shared.BillAccountRef{
                        ID: syncforpayables.String("46f9461e-788b-4906-8b74-d1ea17f6dc10"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("2160"),
                    TaxRateRef: &shared.BillTaxRateRef{
                        ID: syncforpayables.String("INPUT2"),
                    },
                },
                shared.BillLineItem{
                    Description: syncforpayables.String("Desktop/network support via email & phone.Per month fixed fee for minimum 20 hours/month."),
                    UnitAmount: types.MustNewDecimalFromString("4000"),
                    Quantity: types.MustNewDecimalFromString("1"),
                    TaxAmount: types.MustNewDecimalFromString("800"),
                    AccountRef: shared.BillAccountRef{
                        ID: syncforpayables.String("f96c9458-d724-47bf-8f74-a9d5726465ce"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("4800"),
                    TaxRateRef: &shared.BillTaxRateRef{
                        ID: syncforpayables.String("INPUT2"),
                    },
                },
                shared.BillLineItem{
                    Description: syncforpayables.String("Stationery charges"),
                    UnitAmount: types.MustNewDecimalFromString("32"),
                    Quantity: types.MustNewDecimalFromString("8"),
                    TaxAmount: types.MustNewDecimalFromString("51.2"),
                    AccountRef: shared.BillAccountRef{
                        ID: syncforpayables.String("cba6527d-f102-4538-b421-e483233e9d5a"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("307.2"),
                    TaxRateRef: &shared.BillTaxRateRef{
                        ID: syncforpayables.String("INPUT2"),
                    },
                },
            },
            Status: shared.BillStatusOpen,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Bill != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateBillRequest](../../pkg/models/operations/createbillrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateBillResponse](../../pkg/models/operations/createbillresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |

## UploadAttachment

The *Upload bill attachment* endpoint uploads an attachment and assigns it against a specific `billId`.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"os"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.Bills.UploadAttachment(ctx, operations.UploadBillAttachmentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BillID: "EILBDVJVNUAGVKRQ",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UploadBillAttachmentRequest](../../pkg/models/operations/uploadbillattachmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UploadBillAttachmentResponse](../../pkg/models/operations/uploadbillattachmentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## ListAttachments

The *List bill attachments* endpoint returns a list of attachments available to download for a given `billId`.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.ListAttachments(ctx, operations.ListBillAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BillID: "EILBDVJVNUAGVKRQ",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListBillAttachmentsRequest](../../pkg/models/operations/listbillattachmentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListBillAttachmentsResponse](../../pkg/models/operations/listbillattachmentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 401, 402, 403, 404, 409, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## DownloadAttachment

The *Download bill attachment* endpoint downloads a specific attachment for a given `billId` and `attachmentId`.

[Bills](https://docs.codat.io/sync-for-payables-api#/schemas/Bill) are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support downloading a bill attachment.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Bills.DownloadAttachment(ctx, operations.DownloadBillAttachmentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BillID: "EILBDVJVNUAGVKRQ",
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DownloadBillAttachmentRequest](../../pkg/models/operations/downloadbillattachmentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DownloadBillAttachmentResponse](../../pkg/models/operations/downloadbillattachmentresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |