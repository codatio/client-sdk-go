# Transactions
(*Transactions*)

## Overview

Create new bank account transactions for a company's connections, and see previous operations.

### Available Operations

* [Create](#create) - Create bank transactions
* [GetCreateModel](#getcreatemodel) - Get create bank transactions model
* [GetCreateOperation](#getcreateoperation) - Get create operation
* [ListCreateOperations](#listcreateoperations) - List create operations

## Create

﻿The *Create bank transactions* endpoint creates new [bank transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) for a given company's connection.

[Bank transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) are records of monetary amounts that have moved in and out of an SMB's bank account.

**Integration-specific behaviour**

The required properties may vary based on the integration. For detailed requirements specific to each accounting software, refer to the API reference examples.
Alternatively, you can view the [Get create bank transaction model](https://docs.codat.io/bank-feeds-api#/operations/get-create-bank-transactions-model) for more information.

### Example Usage

```go
package main

import(
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v7"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/types"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/operations"
	"context"
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
                    Date: bankfeeds.String("2023-08-22T10:21:00Z"),
                    Description: bankfeeds.String("Repayment of Credit Card"),
                    ID: bankfeeds.String("716422529"),
                    Reconciled: bankfeeds.Bool(true),
                    Reference: bankfeeds.String("Ref-12345"),
                    TransactionType: shared.BankTransactionTypeCredit.ToPointer(),
                },
                shared.BankTransactions{
                    Amount: types.MustNewDecimalFromString("-100"),
                    Balance: types.MustNewDecimalFromString("0"),
                    Counterparty: bankfeeds.String("Amazon"),
                    Date: bankfeeds.String("2023-08-22T10:22:00Z"),
                    Description: bankfeeds.String("Amazon Purchase"),
                    ID: bankfeeds.String("716422530"),
                    Reconciled: bankfeeds.Bool(false),
                    Reference: bankfeeds.String("Ref-12346"),
                    TransactionType: shared.BankTransactionTypeDebit.ToPointer(),
                },
                shared.BankTransactions{
                    Amount: types.MustNewDecimalFromString("-60"),
                    Balance: types.MustNewDecimalFromString("-60"),
                    Counterparty: bankfeeds.String("Office Mart"),
                    Date: bankfeeds.String("2023-08-22T10:23:00Z"),
                    Description: bankfeeds.String("Office Supplies"),
                    ID: bankfeeds.String("716422531"),
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

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetCreateModel

The *Get create bank account transactions model* endpoint returns the expected data for the request payload when creating [bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) for a given company and integration.

[Bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) are records of money that has moved in and out of an SMB's bank account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

### Example Usage

```go
package main

import(
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v7"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.GetCreateModel(ctx, operations.GetCreateBankTransactionsModelRequest{
        AccountID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetCreateBankTransactionsModelRequest](../../pkg/models/operations/getcreatebanktransactionsmodelrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetCreateBankTransactionsModelResponse](../../pkg/models/operations/getcreatebanktransactionsmodelresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetCreateOperation

The **Get create operation** endpoint returns a specific [write operation](/using-the-api/push) identified by the `pushOperationKey` that was performed on the company.

Write operations are actions that send requests to Codat, enabling the creation, updating, deletion of records, or uploading attachments in the connected accounting software.

For bank feeds, your push operations will only relate to the `bankTransactions` data type.

### Example Usage

```go
package main

import(
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v7"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/operations"
	"context"
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

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## ListCreateOperations

The **List create operations** endpoint returns a list of [write operations](/using-the-api/push) performed on the company.

Write operations are actions that send requests to Codat, enabling the creation, updating, deletion of records, or uploading attachments in the connected accounting software. 

For bank feeds, use this endpoint to view write operations related to the `bankTransactions` data type.

### Example Usage

```go
package main

import(
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v7"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v7/pkg/models/operations"
	"context"
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

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |