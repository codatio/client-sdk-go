# DirectIncomes
(*AccountsReceivable.DirectIncomes*)

## Overview

### Available Operations

* [DownloadAttachment](#downloadattachment) - Download direct income attachment
* [Get](#get) - Get direct income
* [GetAttachment](#getattachment) - Get direct income attachment
* [List](#list) - List direct incomes
* [ListAttachments](#listattachments) - List direct income attachments

## DownloadAttachment

The *Download direct income attachment* endpoint downloads a specific attachment for a given `directIncomeId` and `attachmentId`.

[Direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support downloading a direct income attachment.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountsReceivable.DirectIncomes.DownloadAttachment(ctx, operations.DownloadAccountingDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "<value>",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.DownloadAccountingDirectIncomeAttachmentRequest](../../pkg/models/operations/downloadaccountingdirectincomeattachmentrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.DownloadAccountingDirectIncomeAttachmentResponse](../../pkg/models/operations/downloadaccountingdirectincomeattachmentresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## Get

The *Get direct income* endpoint returns a single direct income for a given directIncomeId.

[Direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support getting a specific direct income.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountsReceivable.DirectIncomes.Get(ctx, operations.GetAccountingDirectIncomeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingDirectIncome != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAccountingDirectIncomeRequest](../../pkg/models/operations/getaccountingdirectincomerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetAccountingDirectIncomeResponse](../../pkg/models/operations/getaccountingdirectincomeresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## GetAttachment

The *Get direct income attachment* endpoint returns a specific attachment for a given `directIncomeId` and `attachmentId`.

[Direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support getting a direct income attachment.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountsReceivable.DirectIncomes.GetAttachment(ctx, operations.GetAccountingDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingAttachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetAccountingDirectIncomeAttachmentRequest](../../pkg/models/operations/getaccountingdirectincomeattachmentrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.GetAccountingDirectIncomeAttachmentResponse](../../pkg/models/operations/getaccountingdirectincomeattachmentresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## List

The *List direct incomes* endpoint returns a list of [direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) for a given company's connection.

[Direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountsReceivable.DirectIncomes.List(ctx, operations.ListAccountingDirectIncomesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: lending.String("-modifiedDate"),
        Page: lending.Int(1),
        PageSize: lending.Int(100),
        Query: lending.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingDirectIncomes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListAccountingDirectIncomesRequest](../../pkg/models/operations/listaccountingdirectincomesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListAccountingDirectIncomesResponse](../../pkg/models/operations/listaccountingdirectincomesresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |


## ListAttachments

The *List direct income attachments* endpoint returns a list of attachments available to download for given `directIncomeId`.

[Direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support listing direct income attachments.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountsReceivable.DirectIncomes.ListAttachments(ctx, operations.ListAccountingDirectIncomeAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "<value>",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.ListAccountingDirectIncomeAttachmentsRequest](../../pkg/models/operations/listaccountingdirectincomeattachmentsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.ListAccountingDirectIncomeAttachmentsResponse](../../pkg/models/operations/listaccountingdirectincomeattachmentsresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
