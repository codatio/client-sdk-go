# Transactions
(*Transactions*)

## Overview

Create new bank account transactions for a company's connections, and see previous operations.

### Available Operations

* [Create](#create) - Create bank transactions
* [GetCreateOperation](#getcreateoperation) - Get create operation
* [ListCreateOperations](#listcreateoperations) - List create operations

## Create

﻿The *Create bank transactions* endpoint creates new [bank transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) for a given company's connection.

[Bank transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) are records of monetary amounts that have moved in and out of an SMB's bank account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create bank transaction model](https://docs.codat.io/bank-feeds-api#/operations/get-create-bankTransactions-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support creating a bank account transactions.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v5"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/types"
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.Create(ctx, operations.CreateBankTransactionsRequest{
        CreateBankTransactions: &shared.CreateBankTransactions{
            AccountID: "49cd5a42-b311-4750-9361-52e2ed1d4653",
            Transactions: []shared.BankTransactions{
                shared.BankTransactions{
                    Amount: types.MustNewDecimalFromString("100"),
                    Balance: types.MustNewDecimalFromString("100"),
                    Counterparty: bankfeeds.String("Bank of Example"),
                    Date: "2023-08-22T10:21:00Z",
                    Description: bankfeeds.String("Repayment of Credit Card"),
                    ID: "716422529",
                    Reconciled: bankfeeds.Bool(true),
                    Reference: bankfeeds.String("Ref-12345"),
                    TransactionType: shared.BankTransactionTypeCredit.ToPointer(),
                },
                shared.BankTransactions{
                    Amount: types.MustNewDecimalFromString("-100"),
                    Balance: types.MustNewDecimalFromString("0"),
                    Counterparty: bankfeeds.String("Amazon"),
                    Date: "2023-08-22T10:22:00Z",
                    Description: bankfeeds.String("Amazon Purchase"),
                    ID: "716422530",
                    Reconciled: bankfeeds.Bool(false),
                    Reference: bankfeeds.String("Ref-12346"),
                    TransactionType: shared.BankTransactionTypeDebit.ToPointer(),
                },
                shared.BankTransactions{
                    Amount: types.MustNewDecimalFromString("-60"),
                    Balance: types.MustNewDecimalFromString("-60"),
                    Counterparty: bankfeeds.String("Office Mart"),
                    Date: "2023-08-22T10:23:00Z",
                    Description: bankfeeds.String("Office Supplies"),
                    ID: "716422531",
                    Reconciled: bankfeeds.Bool(false),
                    Reference: bankfeeds.String("Ref-12347"),
                    TransactionType: shared.BankTransactionTypeDebit.ToPointer(),
                },
            },
        },
        AccountID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateBankTransactionsRequest](../../pkg/models/operations/createbanktransactionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateBankTransactionsResponse](../../pkg/models/operations/createbanktransactionsresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## GetCreateOperation

Retrieve push operation.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v5"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.GetCreateOperation(ctx, operations.GetCreateOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "1fb73c31-a851-46c2-ab8a-5ce6e25b57b8",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetCreateOperationRequest](../../pkg/models/operations/getcreateoperationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetCreateOperationResponse](../../pkg/models/operations/getcreateoperationresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ListCreateOperations

List create operations.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v5"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v5/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListCreateOperations(ctx, operations.ListCreateOperationsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: bankfeeds.String("-modifiedDate"),
        Page: bankfeeds.Int(1),
        PageSize: bankfeeds.Int(100),
        Query: bankfeeds.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListCreateOperationsRequest](../../pkg/models/operations/listcreateoperationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListCreateOperationsResponse](../../pkg/models/operations/listcreateoperationsresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
