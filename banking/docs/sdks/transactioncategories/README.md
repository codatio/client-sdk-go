# TransactionCategories

## Overview

Hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.

### Available Operations

* [Get](#get) - Get transaction category
* [List](#list) - List all transaction categories

## Get

Gets a specified bank transaction category for a given company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionCategories.Get(ctx, operations.GetTransactionCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionCategoryID: "unde",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetTransactionCategoryRequest](../../models/operations/gettransactioncategoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GetTransactionCategoryResponse](../../models/operations/gettransactioncategoryresponse.md), error**


## List

Gets a list of hierarchical categories associated with a transaction for greater contextual meaning to transactionactivity.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionCategories.List(ctx, operations.ListTransactionCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: codatbanking.Int(1),
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("nulla"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListTransactionCategoriesRequest](../../models/operations/listtransactioncategoriesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.ListTransactionCategoriesResponse](../../models/operations/listtransactioncategoriesresponse.md), error**
