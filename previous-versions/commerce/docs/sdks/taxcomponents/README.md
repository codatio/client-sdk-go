# TaxComponents
(*TaxComponents*)

## Overview

Retrieve standardized data from linked commerce software.

### Available Operations

* [Get](#get) - Get tax component
* [List](#list) - List tax components

## Get

The *Get tax* endpoint returns a single tax for a given taxId.

[Tax components](https://docs.codat.io/commerce-api#/schemas/TaxComponent) are tax rates from the commerce software, including tax rate's name and value.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxComponents.Get(ctx, operations.GetTaxComponentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TaxID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaxComponent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTaxComponentRequest](../../pkg/models/operations/gettaxcomponentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetTaxComponentResponse](../../pkg/models/operations/gettaxcomponentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 401, 402, 403, 404, 409, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## List

The *List tax components* endpoint returns a list of [tax components](https://docs.codat.io/commerce-api#/schemas/TaxComponent) for a given company's connection.

[Tax components](https://docs.codat.io/commerce-api#/schemas/TaxComponent) are tax rates from the commerce software, including tax rate's name and value.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxComponents.List(ctx, operations.ListTaxComponentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: commerce.String("-modifiedDate"),
        Page: commerce.Int(1),
        PageSize: commerce.Int(100),
        Query: commerce.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaxComponents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTaxComponentsRequest](../../pkg/models/operations/listtaxcomponentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListTaxComponentsResponse](../../pkg/models/operations/listtaxcomponentsresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |