# MappingOptions

## Overview

Mapping options for a companies expenses.

### Available Operations

* [GetMappingOptions](#getmappingoptions) - Mapping options

## GetMappingOptions

Gets the expense mapping options for a companies accounting software

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetMappingOptionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.MappingOptions.GetMappingOptions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingOptions != nil {
        // handle response
    }
}
```
