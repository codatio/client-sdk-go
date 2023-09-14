# BankingTransactionCategories

### Available Operations

* [Get](#get) - Get transaction category
* [List](#list) - List transaction categories

## Get

The *Get transaction category* endpoint returns a single transaction category for a given transactionCategoryId.

[Transaction categories](https://docs.codat.io/banking-api#/schemas/TransactionCategory) are associated with a transaction to provide greater contextual meaning to transaction activity.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactionCategories) for integrations that support getting a specific transaction category.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending/v3"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankingTransactionCategories.Get(ctx, operations.GetBankingTransactionCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionCategoryID: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankingTransactionCategory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetBankingTransactionCategoryRequest](../../models/operations/getbankingtransactioncategoryrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.GetBankingTransactionCategoryResponse](../../models/operations/getbankingtransactioncategoryresponse.md), error**


## List

The *List transaction categories* endpoint returns a list of [transaction categories](https://docs.codat.io/banking-api#/schemas/TransactionCategory) for a given company's connection.

[Transaction categories](https://docs.codat.io/banking-api#/schemas/TransactionCategory) are associated with a transaction to provide greater contextual meaning to transaction activity.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/lending-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending/v3"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankingTransactionCategories.List(ctx, operations.ListBankingTransactionCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatlending.String("-modifiedDate"),
        Page: codatlending.Int(1),
        PageSize: codatlending.Int(100),
        Query: codatlending.String("modi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankingTransactionCategories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListBankingTransactionCategoriesRequest](../../models/operations/listbankingtransactioncategoriesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |


### Response

**[*operations.ListBankingTransactionCategoriesResponse](../../models/operations/listbankingtransactioncategoriesresponse.md), error**

