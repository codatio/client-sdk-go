# BankAccounts
(*BankAccounts*)

## Overview

Access bank accounts in an SMBs accounting software.

### Available Operations

* [Create](#create) - Create bank account
* [GetCreateModel](#getcreatemodel) - Get create/update bank account model
* [List](#list) - List bank accounts

## Create

The *Create bank account* endpoint creates a new [bank account](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bank account model](https://docs.codat.io/bank-feeds-api#/operations/get-create-update-bankAccounts-model).

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v6"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.Create(ctx, operations.CreateBankAccountRequest{
        BankAccountPrototype: &shared.BankAccountPrototype{
            Currency: bankfeeds.String("USD"),
            Status: shared.BankAccountStatusActive.ToPointer(),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateBankAccountRequest](../../pkg/models/operations/createbankaccountrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateBankAccountResponse](../../pkg/models/operations/createbankaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetCreateModel

The *Get create/update bank account model* endpoint returns the expected data for the request payload when creating and updating a [bank account](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) for a given company and integration.

[Bank accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v6"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.GetCreateModel(ctx, operations.GetCreateBankAccountsModelRequest{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCreateBankAccountsModelRequest](../../pkg/models/operations/getcreatebankaccountsmodelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetCreateBankAccountsModelResponse](../../pkg/models/operations/getcreatebankaccountsmodelresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## List

The *List bank accounts* endpoint returns a list of [bank accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/bank-feeds-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v6"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.List(ctx, operations.ListBankAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: bankfeeds.String("-modifiedDate"),
        Page: bankfeeds.Int(1),
        PageSize: bankfeeds.Int(100),
        Query: bankfeeds.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListBankAccountsRequest](../../pkg/models/operations/listbankaccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListBankAccountsResponse](../../pkg/models/operations/listbankaccountsresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |