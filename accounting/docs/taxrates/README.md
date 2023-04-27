# TaxRates

## Overview

Tax rates

### Available Operations

* [GetTaxRate](#gettaxrate) - Get tax rate
* [ListTaxRates](#listtaxrates) - List all tax rates

## GetTaxRate

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetTaxRateRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TaxRateID: "at",
    }

    res, err := s.TaxRates.GetTaxRate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRate != nil {
        // handle response
    }
}
```

## ListTaxRates

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListTaxRatesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("excepturi"),
    }

    res, err := s.TaxRates.ListTaxRates(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRates != nil {
        // handle response
    }
}
```