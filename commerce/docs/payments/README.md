# Payments

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [List](#list) - List payments
* [ListMethods](#listmethods) - List payment methods

## List

List commerce payments for the given company & data connection.

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
    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("quibusdam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```

## ListMethods

Retrieve a list of payment methods, such as card, cash or other online payment methods, as held in the linked commerce platform.

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
    res, err := s.Payments.ListMethods(ctx, operations.ListPaymentMethodsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("unde"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethods != nil {
        // handle response
    }
}
```
