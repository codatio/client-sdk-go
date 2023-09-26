# Transactions

## Overview

Transactions represent debits and credits from a source account.

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
	"context"
	"log"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v3"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/types"
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
            AccountID: bankfeeds.String("nobis"),
            Transactions: []shared.CreateBankTransaction{
                shared.CreateBankTransaction{
                    Amount: types.MustNewDecimalFromString("3154.28"),
                    Balance: types.MustNewDecimalFromString("6078.31"),
                    Date: bankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: bankfeeds.String("minima"),
                    ID: bankfeeds.String("907aff1a-3a2f-4a94-a773-9251aa52c3f5"),
                },
            },
        },
        AccountID: "7110701885",
        AllowSyncOnPushComplete: bankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: bankfeeds.Int(820994),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateBankTransactionsRequest](../../models/operations/createbanktransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateBankTransactionsResponse](../../models/operations/createbanktransactionsresponse.md), error**


## GetCreateOperation

Retrieve push operation.

### Example Usage

```go
package main

import(
	"context"
	"log"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v3"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/operations"
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
        PushOperationKey: "019da1ff-e78f-4097-b007-4f15471b5e6e",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCreateOperationRequest](../../models/operations/getcreateoperationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetCreateOperationResponse](../../models/operations/getcreateoperationresponse.md), error**


## ListCreateOperations

List create operations.

### Example Usage

```go
package main

import(
	"context"
	"log"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v3"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v3/pkg/models/operations"
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
        Query: bankfeeds.String("quae"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCreateOperationsRequest](../../models/operations/listcreateoperationsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.ListCreateOperationsResponse](../../models/operations/listcreateoperationsresponse.md), error**

