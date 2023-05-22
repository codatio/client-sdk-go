# TrackingCategories

## Overview

Tracking categories

### Available Operations

* [Get](#get) - Get tracking categories
* [List](#list) - List tracking categories

## Get

Gets the specified tracking categories for a given company.

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
    res, err := s.TrackingCategories.Get(ctx, operations.GetTrackingCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TrackingCategoryID: "quas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategoryTree != nil {
        // handle response
    }
}
```

## List

Gets the latest tracking categories for a given company.

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
    res, err := s.TrackingCategories.List(ctx, operations.ListTrackingCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("expedita"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategories != nil {
        // handle response
    }
}
```
