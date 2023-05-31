# Products

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [List](#list) - List products
* [ListCategories](#listcategories) - List product categories

## List

The Products data type provides the company's product inventory, and includes the price and quantity of all products, and product variants, available for sale.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.List(ctx, operations.ListProductsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: codatcommerce.Int(1),
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("nulla"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Products != nil {
        // handle response
    }
}
```

## ListCategories

Product categories are used to classify a group of products together, either by type (eg "Furniture"), or sometimes by tax profile.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.ListCategories(ctx, operations.ListProductCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: codatcommerce.Int(1),
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProductCategories != nil {
        // handle response
    }
}
```
