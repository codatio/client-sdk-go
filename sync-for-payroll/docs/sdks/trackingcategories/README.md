# TrackingCategories
(*TrackingCategories*)

## Overview

Tracking categories

### Available Operations

* [Get](#get) - Get tracking categories
* [List](#list) - List tracking categories

## Get

The *Get tracking category* endpoint returns a single tracking category for a given `trackingCategoryId`.

[Tracking categories](https://docs.codat.io/sync-for-payroll-api#/schemas/TrackingCategory) are used to monitor cost centres and control budgets that sit outside the standard set of accounts.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=trackingCategories) for integrations that support getting a specific tracking category.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TrackingCategories.Get(ctx, operations.GetTrackingCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TrackingCategoryID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategoryTree != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetTrackingCategoryRequest](../../pkg/models/operations/gettrackingcategoryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetTrackingCategoryResponse](../../pkg/models/operations/gettrackingcategoryresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## List

The *List tracking categories* endpoint returns a list of [tracking categories](https://docs.codat.io/sync-for-payroll-api#/schemas/TrackingCategory) for a given company's connection.

[Tracking categories](https://docs.codat.io/sync-for-payroll-api#/schemas/TrackingCategory) are used to monitor cost centres and control budgets that sit outside the standard set of accounts.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TrackingCategories.List(ctx, operations.ListTrackingCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforpayroll.String("-modifiedDate"),
        Page: syncforpayroll.Int(1),
        PageSize: syncforpayroll.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListTrackingCategoriesRequest](../../pkg/models/operations/listtrackingcategoriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.ListTrackingCategoriesResponse](../../pkg/models/operations/listtrackingcategoriesresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |
