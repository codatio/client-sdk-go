# Transactions
(*Transactions*)

## Overview

An immutable source of up-to-date information on income and expenditure.

### Available Operations

* [Get](#get) - Get bank transaction
* [List](#list) - List transactions
* [~~ListBankTransactions~~](#listbanktransactions) - List banking transactions :warning: **Deprecated** Use `List` instead.

## Get

The *Get transaction* endpoint returns a single transaction for a given transactionId.

[Transactions](https://docs.codat.io/banking-api#/schemas/Transaction) provide an immutable source of up-to-date information on income and expenditure.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions) for integrations that support getting a specific transaction.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
)

func main() {
    s := banking.New(
        banking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.Get(ctx, operations.GetTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTransactionRequest](../../pkg/models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetTransactionResponse](../../pkg/models/operations/gettransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## List

The *List transactions* endpoint returns a list of [transactions](https://docs.codat.io/banking-api#/schemas/Transaction) for a given company's connection.

[Transactions](https://docs.codat.io/banking-api#/schemas/Transaction) provide an immutable source of up-to-date information on income and expenditure.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
)

func main() {
    s := banking.New(
        banking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: banking.String("-modifiedDate"),
        Page: banking.Int(1),
        PageSize: banking.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTransactionsRequest](../../pkg/models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.ListTransactionsResponse](../../pkg/models/operations/listtransactionsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |

## ~~ListBankTransactions~~

The *List transactions* endpoint returns a list of [transactions](https://docs.codat.io/banking-api#/schemas/Transaction) for a given company's connection.

[Transactions](https://docs.codat.io/banking-api#/schemas/Transaction) provide an immutable source of up-to-date information on income and expenditure.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible. Use `List` instead.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
)

func main() {
    s := banking.New(
        banking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListBankTransactions(ctx, operations.ListBankTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: banking.String("-modifiedDate"),
        Page: banking.Int(1),
        PageSize: banking.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListBankTransactionsRequest](../../pkg/models/operations/listbanktransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.ListBankTransactionsResponse](../../pkg/models/operations/listbanktransactionsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |
