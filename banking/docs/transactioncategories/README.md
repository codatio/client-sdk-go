# TransactionCategories

## Overview

Hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.

### Available Operations

* [Get](#get) - Get transaction category
* [List](#list) - List all transaction categories

## Get

Gets a specified bank transaction category for a given company

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
    res, err := s.TransactionCategories.Get(ctx, operations.GetTransactionCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionCategoryID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategory != nil {
        // handle response
    }
}
```

## List

Gets a list of hierarchical categories associated with a transaction for greater contextual meaning to transactionactivity.

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
    res, err := s.TransactionCategories.List(ctx, operations.ListTransactionCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("unde"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategories != nil {
        // handle response
    }
}
```
