# TrackingCategories

## Overview

Tracking categories

### Available Operations

* [GetTrackingCategory](#gettrackingcategory) - Get tracking categories
* [ListTrackingCategories](#listtrackingcategories) - List tracking categories

## GetTrackingCategory

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
    req := operations.GetTrackingCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TrackingCategoryID: "perferendis",
    }

    res, err := s.TrackingCategories.GetTrackingCategory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategoryTree != nil {
        // handle response
    }
}
```

## ListTrackingCategories

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
    req := operations.ListTrackingCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("nostrum"),
    }

    res, err := s.TrackingCategories.ListTrackingCategories(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TrackingCategories != nil {
        // handle response
    }
}
```