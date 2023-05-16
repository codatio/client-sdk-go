# TaxComponents

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [Get](#get) - List tax components

## Get

This endpoint returns a lists of tax rates from the commerce platform, including tax rate names and values. This supports the mapping of tax rates from the commerce platform to the accounting platform.

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
    res, err := s.TaxComponents.Get(ctx, operations.GetTaxComponentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxComponents != nil {
        // handle response
    }
}
```
