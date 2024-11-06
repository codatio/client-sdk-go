# CodatLendingLoanWritebackAccounts
(*LoanWriteback.Accounts*)

## Overview

### Available Operations

* [Create](#create) - Create account
* [GetCreateModel](#getcreatemodel) - Get create account model

## Create

The *Create account* endpoint creates a new [account](https://docs.codat.io/lending-api#/schemas/Account) for a given company's connection.

[Accounts](https://docs.codat.io/lending-api#/schemas/Account) are the categories a business uses to record accounting transactions.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create account model](https://docs.codat.io/lending-api#/operations/get-create-chartOfAccounts-model).

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v6"
	"context"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/types"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Accounts.Create(ctx, operations.CreateAccountRequest{
        AccountPrototype: &shared.AccountPrototype{
            Currency: lending.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: lending.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: lending.String("Asset.Current"),
            FullyQualifiedName: lending.String("Cash On Hand"),
            Name: lending.String("Accounts Receivable"),
            NominalCode: lending.String("610"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreateAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateAccountRequest](../../pkg/models/operations/createaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetCreateModel

The *Get create account model* endpoint returns the expected data for the request payload when creating an [account](https://docs.codat.io/lending-api#/schemas/Account) for a given company and integration.

[Accounts](https://docs.codat.io/lending-api#/schemas/Account) are the categories a business uses to record accounting transactions.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v6"
	"context"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Accounts.GetCreateModel(ctx, operations.GetCreateChartOfAccountsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetCreateChartOfAccountsModelRequest](../../pkg/models/operations/getcreatechartofaccountsmodelrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.GetCreateChartOfAccountsModelResponse](../../pkg/models/operations/getcreatechartofaccountsmodelresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |