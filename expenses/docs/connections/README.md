# Connections

## Overview

Create and manage partner expense connection.

### Available Operations

* [CreatePartnerExpenseConnection](#createpartnerexpenseconnection) - Create Partner Expense connection

## CreatePartnerExpenseConnection

Creates a Partner Expense data connection

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.CreatePartnerExpenseConnection(ctx, operations.CreatePartnerExpenseConnectionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataConnection != nil {
        // handle response
    }
}
```
