# AccountBalances

## Overview

Balances for a bank account including end-of-day batch balance or running balances per transaction.

### Available Operations

* [List](#list) - List account balances

## List

Gets a list of balances for a bank account including end-of-day batch balance or running balances per transaction.

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
    req := operations.ListAccountBalancesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("provident"),
    }

    res, err := s.AccountBalances.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountBalances != nil {
        // handle response
    }
}
```