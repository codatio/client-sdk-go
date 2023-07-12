# BankFeedAccounts

## Overview

Bank feed bank accounts

### Available Operations

* [Create](#create) - Create a bank feed bank account
* [List](#list) - List bank feed bank accounts
* [~~PutBankFeed~~](#putbankfeed) - Create bank feed bank accounts :warning: **Deprecated**
* [Update](#update) - Update bank feed bank account

## Create

Post a BankFeed BankAccount for a single data source connected to a single company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankFeedAccounts.Create(ctx, operations.CreateBankFeedRequest{
        BankFeedAccount: &shared.BankFeedAccount{
            AccountName: codatbankfeeds.String("maiores"),
            AccountNumber: codatbankfeeds.String("dicta"),
            AccountType: codatbankfeeds.String("corporis"),
            Balance: codatbankfeeds.Float64(2961.4),
            Currency: codatbankfeeds.String("USD"),
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            ID: "b5e6e13b-99d4-488e-9e91-e450ad2abd44",
            ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            SortCode: codatbankfeeds.String("aliquid"),
            Status: codatbankfeeds.String("cupiditate"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateBankFeedRequest](../../models/operations/createbankfeedrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateBankFeedResponse](../../models/operations/createbankfeedresponse.md), error**


## List

﻿The *List bank feed bank accounts* endpoint returns a list of [bank feed accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankFeedAccount) for a given company's connection.

[Bank feed accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankFeedAccount) are the bank's bank account from which transactions are synced into the accounting platform.



### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankFeedAccounts.List(ctx, operations.ListBankFeedsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListBankFeedsRequest](../../models/operations/listbankfeedsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListBankFeedsResponse](../../models/operations/listbankfeedsresponse.md), error**


## ~~PutBankFeed~~

Put BankFeed BankAccounts for a single data source connected to a single company.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankFeedAccounts.PutBankFeed(ctx, operations.PutBankFeedRequest{
        RequestBody: []shared.BankFeedAccount{
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("perferendis"),
                AccountNumber: codatbankfeeds.String("magni"),
                AccountType: codatbankfeeds.String("assumenda"),
                Balance: codatbankfeeds.Float64(3698.08),
                Currency: codatbankfeeds.String("GBP"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "a94bb4f6-3c96-49e9-a3ef-a77dfb14cd66",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("accusamus"),
                Status: codatbankfeeds.String("non"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("occaecati"),
                AccountNumber: codatbankfeeds.String("enim"),
                AccountType: codatbankfeeds.String("accusamus"),
                Balance: codatbankfeeds.Float64(9654.17),
                Currency: codatbankfeeds.String("EUR"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "ba88f3a6-6997-4074-ba44-69b6e2141959",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("sint"),
                Status: codatbankfeeds.String("accusantium"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("mollitia"),
                AccountNumber: codatbankfeeds.String("reiciendis"),
                AccountType: codatbankfeeds.String("mollitia"),
                Balance: codatbankfeeds.Float64(3209.97),
                Currency: codatbankfeeds.String("USD"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "e2516fe4-c8b7-411e-9b7f-d2ed028921cd",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("maxime"),
                Status: codatbankfeeds.String("ea"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PutBankFeedRequest](../../models/operations/putbankfeedrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.PutBankFeedResponse](../../models/operations/putbankfeedresponse.md), error**


## Update

﻿The *Update bank feed bank account* endpoint updates a single bank feed bank account for a single data source connected to a single company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankFeedAccounts.Update(ctx, operations.UpdateBankFeedRequest{
        BankFeedAccount: &shared.BankFeedAccount{
            AccountName: codatbankfeeds.String("excepturi"),
            AccountNumber: codatbankfeeds.String("odit"),
            AccountType: codatbankfeeds.String("ea"),
            Balance: codatbankfeeds.Float64(332.22),
            Currency: codatbankfeeds.String("GBP"),
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            ID: "b576b0d5-f0d3-40c5-bbb2-587053202c73",
            ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            SortCode: codatbankfeeds.String("nostrum"),
            Status: codatbankfeeds.String("hic"),
        },
        AccountID: "EILBDVJVNUAGVKRQ",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateBankFeedRequest](../../models/operations/updatebankfeedrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.UpdateBankFeedResponse](../../models/operations/updatebankfeedresponse.md), error**

