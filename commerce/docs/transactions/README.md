# Transactions

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [ListTransactions](#listtransactions) - List transactions

## ListTransactions

Details of all financial transactions recorded in the commerce or point of sale system are added to the Transactions data type. For example, payments, service charges, and fees.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("illum"),
    }

    res, err := s.Transactions.ListTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```
