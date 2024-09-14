# CommerceTransactions
(*CommerceTransactions*)

## Overview

Retrieve standardized Commerce transactions from linked commerce software.

### Available Operations

* [GetCommerceTransaction](#getcommercetransaction) - Get transaction
* [ListCommerceTransactions](#listcommercetransactions) - List transactions

## GetCommerceTransaction

The *Get transaction* endpoint returns a single transaction for a given transactionId.

[Transactions](https://docs.codat.io/commerce-api#/schemas/Transaction) detail all financial affairs recorded in the commerce or point of sale system.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-transactions) for integrations that support getting a specific transaction.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-commerce-v1-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CommerceTransactions.GetCommerceTransaction(ctx, operations.GetCommerceTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCommerceTransactionRequest](../../pkg/models/operations/getcommercetransactionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetCommerceTransactionResponse](../../pkg/models/operations/getcommercetransactionresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## ListCommerceTransactions

The *List transactions* endpoint returns a list of [transactions](https://docs.codat.io/commerce-api#/schemas/Transaction) for a given company's connection.

[Transactions](https://docs.codat.io/commerce-api#/schemas/Transaction) detail all financial affairs recorded in the commerce or point of sale system.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-commerce-v1-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CommerceTransactions.ListCommerceTransactions(ctx, operations.ListCommerceTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: syncforcommerceversion1.String("-modifiedDate"),
        Page: syncforcommerceversion1.Int(1),
        PageSize: syncforcommerceversion1.Int(100),
        Query: syncforcommerceversion1.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCommerceTransactionsRequest](../../pkg/models/operations/listcommercetransactionsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCommerceTransactionsResponse](../../pkg/models/operations/listcommercetransactionsresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
