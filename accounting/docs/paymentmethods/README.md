# PaymentMethods

## Overview

Payment methods

### Available Operations

* [Get](#get) - Get payment method
* [List](#list) - List all payment methods

## Get

Gets the specified payment method for a given company.

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
    req := operations.GetPaymentMethodRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentMethodID: "ea",
    }

    res, err := s.PaymentMethods.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

## List

Gets the payment methods for a given company.

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
    req := operations.ListPaymentMethodsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("error"),
    }

    res, err := s.PaymentMethods.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethods != nil {
        // handle response
    }
}
```
