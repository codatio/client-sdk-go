# BankFeedAccounts

## Overview

Bank feed bank accounts

### Available Operations

* [Create](#create) - Create bank feed bank accounts
* [Get](#get) - List bank feed bank accounts
* [Update](#update) - Update bank feed bank account

## Create

Put BankFeed BankAccounts for a single data source connected to a single company.

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
        RequestBody: []shared.BankFeedAccount{
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("ipsa"),
                AccountNumber: codatbankfeeds.String("reiciendis"),
                AccountType: codatbankfeeds.String("est"),
                Balance: codatbankfeeds.Float64(6531.4),
                Currency: codatbankfeeds.String("laborum"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "352c5955-907a-4ff1-a3a2-fa9467739251",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("animi"),
                Status: codatbankfeeds.String("enim"),
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

## Get

Get BankFeed BankAccounts for a single data source connected to a single company.

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
    res, err := s.BankFeedAccounts.Get(ctx, operations.GetBankFeedsRequest{
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

## Update

Update a single BankFeed BankAccount for a single data source connected to a single company.

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
            AccountName: codatbankfeeds.String("odit"),
            AccountNumber: codatbankfeeds.String("quo"),
            AccountType: codatbankfeeds.String("sequi"),
            Balance: codatbankfeeds.Float64(9495.72),
            Currency: codatbankfeeds.String("ipsam"),
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            ID: "d019da1f-fe78-4f09-bb00-74f15471b5e6",
            ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            SortCode: codatbankfeeds.String("quae"),
            Status: codatbankfeeds.String("ipsum"),
        },
        AccountID: "7110701885",
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
