# PaymentMethods

## Overview

Payment methods

### Available Operations

* [GetPaymentMethod](#getpaymentmethod) - Get payment method
* [ListPaymentMethods](#listpaymentmethods) - List all payment methods

## GetPaymentMethod

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
        PaymentMethodID: "harum",
    }

    res, err := s.PaymentMethods.GetPaymentMethod(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

## ListPaymentMethods

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

    res, err := s.PaymentMethods.ListPaymentMethods(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethods != nil {
        // handle response
    }
}
```
