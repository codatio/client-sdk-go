# RefreshData
(*RefreshData*)

## Overview

Initiate data refreshes, view pull status and history.

### Available Operations

* [All](#all) - Refresh all data
* [ByDataType](#bydatatype) - Refresh data type
* [Get](#get) - Get data status
* [GetPullOperation](#getpulloperation) - Get pull operation
* [ListPullOperations](#listpulloperations) - List pull operations

## All

Refreshes all data types with `fetch on first link` set to `true` for a given company.

This is an asynchronous operation, and will bring updated data into Codat from the linked integration for you to view.

[Read more](https://docs.codat.io/core-concepts/data-type-settings) about data type settings and `fetch on first link`.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.RefreshData.All(ctx, operations.RefreshCompanyDataRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RefreshCompanyDataRequest](../../pkg/models/operations/refreshcompanydatarequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RefreshCompanyDataResponse](../../pkg/models/operations/refreshcompanydataresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ByDataType

Refreshes a given data type for a given company.

This is an asynchronous operation, and will bring updated data into Codat from the linked integration for you to view.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.RefreshData.ByDataType(ctx, operations.RefreshDataTypeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.SchemaDataTypeInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RefreshDataTypeRequest](../../pkg/models/operations/refreshdatatyperequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RefreshDataTypeResponse](../../pkg/models/operations/refreshdatatyperesponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## Get

Get the state of each data type for a company

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.RefreshData.Get(ctx, operations.GetCompanyDataStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCompanyDataStatusRequest](../../pkg/models/operations/getcompanydatastatusrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetCompanyDataStatusResponse](../../pkg/models/operations/getcompanydatastatusresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## GetPullOperation

Retrieve information about a single dataset or pull operation.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.RefreshData.GetPullOperation(ctx, operations.GetPullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DatasetID: "7911a54a-c808-4f4b-b87e-b195f52b4da5",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPullOperationRequest](../../pkg/models/operations/getpulloperationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPullOperationResponse](../../pkg/models/operations/getpulloperationresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ListPullOperations

Gets the pull operation history (datasets) for a given company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.RefreshData.ListPullOperations(ctx, operations.ListPullOperationsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: platform.String("-modifiedDate"),
        Page: platform.Int(1),
        PageSize: platform.Int(100),
        Query: platform.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PullOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPullOperationsRequest](../../pkg/models/operations/listpulloperationsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPullOperationsResponse](../../pkg/models/operations/listpulloperationsresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
