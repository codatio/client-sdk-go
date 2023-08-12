# Transactions

## Overview

Transactions represent debits and credits from a source account.

### Available Operations

* [Create](#create) - Create bank account transactions
* [GetOperation](#getoperation) - Get push operation
* [ListOperations](#listoperations) - List push operations

## Create

﻿The *Create bank account transactions* endpoint creates new [bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) for a given company's connection.

[Bank account transactions](https://docs.codat.io/bank-feeds-api#/schemas/BankTransactions) are records of monetary amounts that have moved in and out of an SMB's bank account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create bank transaction model](https://docs.codat.io/bank-feeds-api#/operations/get-create-bankTransactions-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support creating a bank account transactions.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.Create(ctx, operations.CreateBankTransactionsRequest{
        CreateBankTransactions: &shared.CreateBankTransactions{
            AccountID: codatbankfeeds.String("dolorem"),
            Transactions: []shared.CreateBankAccountTransaction{
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(1289.26),
                    Balance: codatbankfeeds.Float64(7506.86),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("omnis"),
                    ID: codatbankfeeds.String("55907aff-1a3a-42fa-9467-739251aa52c3"),
                },
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(9495.72),
                    Balance: codatbankfeeds.Float64(3687.25),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("possimus"),
                    ID: codatbankfeeds.String("019da1ff-e78f-4097-b007-4f15471b5e6e"),
                },
            },
        },
        AccountID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(216822),
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


## GetOperation

Retrieve push operation.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.GetOperation(ctx, operations.GetPushOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "b99d488e-1e91-4e45-8ad2-abd44269802d",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPushOperationRequest](../../models/operations/getpushoperationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetPushOperationResponse](../../models/operations/getpushoperationresponse.md), error**


## ListOperations

List push operation records.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transactions.ListOperations(ctx, operations.GetCompanyPushHistoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatbankfeeds.String("-modifiedDate"),
        Page: codatbankfeeds.Int(1),
        PageSize: codatbankfeeds.Int(100),
        Query: codatbankfeeds.String("ipsam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCompanyPushHistoryRequest](../../models/operations/getcompanypushhistoryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetCompanyPushHistoryResponse](../../models/operations/getcompanypushhistoryresponse.md), error**

