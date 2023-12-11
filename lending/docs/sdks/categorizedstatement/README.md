# CategorizedStatement
(*Banking.CategorizedStatement*)

### Available Operations

* [Get](#get) - Get categorized bank statement

## Get

﻿> **Categorization engine**
>
> The categorization engine uses machine learning and has been fully trained against Plaid and TrueLayer banking data sources. It is not fully trained against the Basiq banking data source.

The _Get categorized bank statement_ endpoint provides a fully categorized list of banking transactions for a company. Accounts and transaction data are obtained from the company's banking data sources.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Banking.CategorizedStatement.Get(ctx, operations.GetCategorizedBankStatementRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: lending.Int(1),
        PageSize: lending.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedCashFlowTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetCategorizedBankStatementRequest](../../pkg/models/operations/getcategorizedbankstatementrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.GetCategorizedBankStatementResponse](../../pkg/models/operations/getcategorizedbankstatementresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |