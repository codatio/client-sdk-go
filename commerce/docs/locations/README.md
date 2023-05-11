# Locations

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [List](#list) - List locations

## List

Retrieve a list of locations as seen in the commerce platform.

A `location` is a geographic place at which stocks of products may be held, or from where orders were placed.

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
    res, err := s.Locations.List(ctx, operations.ListLocationsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LocationsResponse != nil {
        // handle response
    }
}
```
