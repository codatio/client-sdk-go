# Products
(*Products*)

## Overview

Retrieve standardized data from linked commerce platforms.

### Available Operations

* [Get](#get) - Get product
* [GetCategory](#getcategory) - Get product category
* [List](#list) - List products
* [ListCategories](#listcategories) - List product categories

## Get

The *Get product* endpoint returns a single product for a given productId.

[Products](https://docs.codat.io/commerce-api#/schemas/Product) are items in the company's inventory that are available for sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-products) for integrations that support getting a specific product.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.Get(ctx, operations.GetProductRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ProductID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Product != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetProductRequest](../../models/operations/getproductrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.GetProductResponse](../../models/operations/getproductresponse.md), error**


## GetCategory

The *Get product* endpoint returns a single product for a given productId.

[Product categories](https://docs.codat.io/commerce-api#/schemas/ProductCategory) are used to classify a group of products together, either by type (e.g. "Furniture"), or sometimes by tax profile.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-productCategories) for integrations that support getting a specific product.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.GetCategory(ctx, operations.GetProductCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ProductID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProductCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetProductCategoryRequest](../../models/operations/getproductcategoryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetProductCategoryResponse](../../models/operations/getproductcategoryresponse.md), error**


## List

The *List products* endpoint returns a list of [products](https://docs.codat.io/commerce-api#/schemas/Product) for a given company's connection.

[Products](https://docs.codat.io/commerce-api#/schemas/Product) are items in the company's inventory that are available for sale.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.List(ctx, operations.ListProductsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: commerce.String("-modifiedDate"),
        Page: commerce.Int(1),
        PageSize: commerce.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Products != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListProductsRequest](../../models/operations/listproductsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.ListProductsResponse](../../models/operations/listproductsresponse.md), error**


## ListCategories

The *List product categories* endpoint returns a list of [product categories](https://docs.codat.io/commerce-api#/schemas/ProductCategory) for a given company's connection.

[Product categories](https://docs.codat.io/commerce-api#/schemas/ProductCategory) are used to classify a group of products together, either by type (e.g. "Furniture"), or sometimes by tax profile.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Products.ListCategories(ctx, operations.ListProductCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: commerce.String("-modifiedDate"),
        Page: commerce.Int(1),
        PageSize: commerce.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProductCategories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListProductCategoriesRequest](../../models/operations/listproductcategoriesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.ListProductCategoriesResponse](../../models/operations/listproductcategoriesresponse.md), error**

