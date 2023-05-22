# Transactions

## Overview

An immutable source of up-to-date information on income and expenditure.

### Available Operations

* [Get](#get) - Get bank transaction
* [List](#list) - List transactions
* [~~ListBankTransactions~~](#listbanktransactions) - List banking transactions :warning: **Deprecated** - Use `List` instead.

## Get

Gets a specified bank transaction for a given company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.Get(ctx, operations.GetTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionID: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

## List

Gets a list of transactions incurred by a bank account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: codatbanking.Int(1),
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```

## ~~ListBankTransactions~~

Gets a list of transactions incurred by a company across all bank accounts.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible. Use `List` instead.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListBankTransactions(ctx, operations.ListBankTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: codatbanking.Int(1),
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("illum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```
