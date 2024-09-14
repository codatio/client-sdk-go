# TaxRates
(*TaxRates*)

## Overview

Get, create, and update Tax rates.

### Available Operations

* [Get](#get) - Get tax rate
* [List](#list) - List all tax rates

## Get

The *Get tax rate* endpoint returns a single tax rate for a given `taxRateId`.

[Tax rates](https://docs.codat.io/sync-for-payables-api#/schemas/TaxRate) are a set of taxes and associated rates within the SMB's accounting software.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=taxRates) for integrations that support getting a specific tax rate.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxRates.Get(ctx, operations.GetTaxRateRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TaxRateID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaxRate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetTaxRateRequest](../../pkg/models/operations/gettaxraterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetTaxRateResponse](../../pkg/models/operations/gettaxrateresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## List

The *List tax rates* endpoint returns a list of [tax rates](https://docs.codat.io/sync-for-payables-api#/schemas/TaxRate) for a given company's connection.

[Tax rates](https://docs.codat.io/sync-for-payables-api#/schemas/TaxRate) are a set of taxes and associated rates within the SMB's accounting software.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxRates.List(ctx, operations.ListTaxRatesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforpayables.String("-modifiedDate"),
        Page: syncforpayables.Int(1),
        PageSize: syncforpayables.Int(100),
        Query: syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaxRates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListTaxRatesRequest](../../pkg/models/operations/listtaxratesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListTaxRatesResponse](../../pkg/models/operations/listtaxratesresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
