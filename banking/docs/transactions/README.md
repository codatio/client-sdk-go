# Transactions

## Overview

An immutable source of up-to-date information on income and expenditure.

### Available Operations

* [Get](#get) - Get bank transaction
* [List](#list) - List banking transactions
* [List](#list) - List transactions

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
    req := operations.GetTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionID: "nulla",
    }

    res, err := s.Transactions.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

## List

Gets a list of transactions incurred by a company across all bank accounts.

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
    req := operations.ListBankTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("corrupti"),
    }

    res, err := s.Transactions.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
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
    req := operations.ListTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("illum"),
    }

    res, err := s.Transactions.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```
