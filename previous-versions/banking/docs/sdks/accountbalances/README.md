# AccountBalances
(*AccountBalances*)

## Overview

Balances for a bank account including end-of-day batch balance or running balances per transaction.

### Available Operations

* [List](#list) - List account balances

## List

The *List account balances* endpoint returns a list of [account balances](https://docs.codat.io/banking-api#/schemas/AccountBalance) for a given company's connection.

[Account balances](https://docs.codat.io/banking-api#/schemas/AccountBalance) are balances for a bank account, including end-of-day batch balance or running balances per transaction.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
	"log"
)

func main() {
    s := banking.New(
        banking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountBalances.List(ctx, operations.ListAccountBalancesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: banking.String("-modifiedDate"),
        Page: banking.Int(1),
        PageSize: banking.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountBalances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAccountBalancesRequest](../../pkg/models/operations/listaccountbalancesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.ListAccountBalancesResponse](../../pkg/models/operations/listaccountbalancesresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |
