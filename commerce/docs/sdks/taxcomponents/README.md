# TaxComponents

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [Get](#get) - Get tax component
* [List](#list) - List tax components

## Get

This endpoint returns a specific tax rate from the commerce platform, including tax rate names and values. This supports the mapping of tax rates from the commerce platform to the accounting platform.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/commerce"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
)

func main() {
    s := codatcommerce.New(
        codatcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxComponents.Get(ctx, operations.GetTaxComponentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TaxID: "ipsa",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetTaxComponentRequest](../../models/operations/gettaxcomponentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetTaxComponentResponse](../../models/operations/gettaxcomponentresponse.md), error**


## List

This endpoint returns a lists of tax rates from the commerce platform, including tax rate names and values. This supports the mapping of tax rates from the commerce platform to the accounting platform.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/commerce"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
)

func main() {
    s := codatcommerce.New(
        codatcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxComponents.List(ctx, operations.ListTaxComponentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: codatcommerce.Int(1),
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("delectus"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListTaxComponentsRequest](../../models/operations/listtaxcomponentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ListTaxComponentsResponse](../../models/operations/listtaxcomponentsresponse.md), error**
