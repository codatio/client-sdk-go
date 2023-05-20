# Orders

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [List](#list) - List orders

## List

Get a list of orders placed or held on the linked commerce platform

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
    res, err := s.Orders.List(ctx, operations.ListOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: codatcommerce.Int(1),
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("distinctio"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Orders != nil {
        // handle response
    }
}
```
