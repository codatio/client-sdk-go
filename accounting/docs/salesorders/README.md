# SalesOrders

## Overview

Sales orders

### Available Operations

* [Get](#get) - Get sales order
* [List](#list) - List sales orders

## Get

Get sales order

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetSalesOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SalesOrderID: "vel",
    }

    res, err := s.SalesOrders.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SalesOrder != nil {
        // handle response
    }
}
```

## List

Get sales orders

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListSalesOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("exercitationem"),
    }

    res, err := s.SalesOrders.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SalesOrders != nil {
        // handle response
    }
}
```
