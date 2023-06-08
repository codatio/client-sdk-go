# TaxRates

## Overview

Tax rates

### Available Operations

* [Get](#get) - Get tax rate
* [List](#list) - List all tax rates

## Get

Gets the specified tax rate for a given company.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxRates.Get(ctx, operations.GetTaxRateRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TaxRateID: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRate != nil {
        // handle response
    }
}
```

## List

Gets the latest tax rates for a given company.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TaxRates.List(ctx, operations.ListTaxRatesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("sed"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRates != nil {
        // handle response
    }
}
```
