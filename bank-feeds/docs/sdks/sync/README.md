# Sync
(*Sync*)

## Overview

Monitor the status of data syncs.

### Available Operations

* [GetLastSuccessfulSync](#getlastsuccessfulsync) - Get last successful sync

## GetLastSuccessfulSync

Use the _Get last successful sync_ endpoint to obtain the status information for the company's [most recent successful sync](https://docs.codat.io/bank-feeds-api#/schemas/CompanySyncStatus). 

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v6"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.GetLastSuccessfulSync(ctx, operations.GetLastSuccessfulRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetLastSuccessfulRequest](../../pkg/models/operations/getlastsuccessfulrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetLastSuccessfulResponse](../../pkg/models/operations/getlastsuccessfulresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |