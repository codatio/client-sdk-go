# Categories

## Overview

Categorisation

### Available Operations

* [~~GetAccountCategory~~](#getaccountcategory) - Get suggested and/or confirmed category for a specific account :warning: **Deprecated**
* [~~ListAccountsCategories~~](#listaccountscategories) - List suggested and confirmed account categories :warning: **Deprecated**
* [~~ListAvailableAccountCategories~~](#listavailableaccountcategories) - List account categories :warning: **Deprecated**
* [~~UpdateAccountCategory~~](#updateaccountcategory) - Update account categories :warning: **Deprecated**
* [~~UpdateAccountsCategories~~](#updateaccountscategories) - Confirm categories for accounts :warning: **Deprecated**

## ~~GetAccountCategory~~

Get category for specific nominal account.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.GetAccountCategory(ctx, operations.GetAccountCategoryRequest{
        AccountID: "provident",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAccountCategoryRequest](../../models/operations/getaccountcategoryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetAccountCategoryResponse](../../models/operations/getaccountcategoryresponse.md), error**


## ~~ListAccountsCategories~~

Lists suggested and confirmed chart of account categories for the given company and data connection.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.ListAccountsCategories(ctx, operations.ListAccountsCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatassess.String("-modifiedDate"),
        Page: codatassess.Int(1),
        PageSize: codatassess.Int(100),
        Query: codatassess.String("distinctio"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAccountsCategoriesRequest](../../models/operations/listaccountscategoriesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.ListAccountsCategoriesResponse](../../models/operations/listaccountscategoriesresponse.md), error**


## ~~ListAvailableAccountCategories~~

Lists available account categories Codat's categorisation engine can provide. 

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.ListAvailableAccountCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Categories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.ListAvailableAccountCategoriesResponse](../../models/operations/listavailableaccountcategoriesresponse.md), error**


## ~~UpdateAccountCategory~~

Update category for a specific nominal account

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.UpdateAccountCategory(ctx, operations.UpdateAccountCategoryRequest{
        ConfirmCategory: &shared.ConfirmCategory{
            Confirmed: shared.AccountCategory{
                DetailType: codatassess.String("quibusdam"),
                Subtype: codatassess.String("unde"),
                Type: codatassess.String("nulla"),
            },
        },
        AccountID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateAccountCategoryRequest](../../models/operations/updateaccountcategoryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.UpdateAccountCategoryResponse](../../models/operations/updateaccountcategoryresponse.md), error**


## ~~UpdateAccountsCategories~~

Comfirms the categories for all or a batch of accounts for a specific connection.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.UpdateAccountsCategories(ctx, operations.UpdateAccountsCategoriesRequest{
        ConfirmCategories: &shared.ConfirmCategories{
            Categories: []shared.ConfirmCategoriesCategories{
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("error"),
                        Subtype: codatassess.String("deserunt"),
                        Type: codatassess.String("suscipit"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("magnam"),
                        Subtype: codatassess.String("debitis"),
                        Type: codatassess.String("ipsa"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "EILBDVJVNUAGVKRQ",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("tempora"),
                        Subtype: codatassess.String("suscipit"),
                        Type: codatassess.String("molestiae"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "EILBDVJVNUAGVKRQ",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("placeat"),
                        Subtype: codatassess.String("voluptatum"),
                        Type: codatassess.String("iusto"),
                    },
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAccountsCategoriesRequest](../../models/operations/updateaccountscategoriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.UpdateAccountsCategoriesResponse](../../models/operations/updateaccountscategoriesresponse.md), error**

