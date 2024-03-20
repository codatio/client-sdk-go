# MappingOptions
(*MappingOptions*)

## Overview

Mapping options for a companies expenses.

### Available Operations

* [GetMappingOptions](#getmappingoptions) - Mapping options

## GetMappingOptions

Gets the expense mapping options for a companies accounting software

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.MappingOptions.GetMappingOptions(ctx, operations.GetMappingOptionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MappingOptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetMappingOptionsRequest](../../pkg/models/operations/getmappingoptionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.GetMappingOptionsResponse](../../pkg/models/operations/getmappingoptionsresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
