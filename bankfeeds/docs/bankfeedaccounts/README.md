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
                AccountName: codatbankfeeds.String("dolor"),
                AccountNumber: codatbankfeeds.String("debitis"),
                AccountType: codatbankfeeds.String("a"),
                Balance: codatbankfeeds.Float64(6800.56),
                Currency: codatbankfeeds.String("in"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "dfb14cd6-6ae3-495e-bb9b-a88f3a669970",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("magnam"),
                Status: codatbankfeeds.String("distinctio"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("id"),
                AccountNumber: codatbankfeeds.String("labore"),
                AccountType: codatbankfeeds.String("labore"),
                Balance: codatbankfeeds.Float64(3834.62),
                Currency: codatbankfeeds.String("natus"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "6e214195-9890-4afa-963e-2516fe4c8b71",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("repudiandae"),
                Status: codatbankfeeds.String("ullam"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("expedita"),
                AccountNumber: codatbankfeeds.String("nihil"),
                AccountType: codatbankfeeds.String("repellat"),
                Balance: codatbankfeeds.Float64(8411.4),
                Currency: codatbankfeeds.String("sed"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                ID: "d028921c-ddc6-4926-81fb-576b0d5f0d30",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                SortCode: codatbankfeeds.String("corporis"),
                Status: codatbankfeeds.String("hic"),
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
            AccountName: codatbankfeeds.String("libero"),
            AccountNumber: codatbankfeeds.String("nobis"),
            AccountType: codatbankfeeds.String("dolores"),
            Balance: codatbankfeeds.Float64(3394.04),
            Currency: codatbankfeeds.String("totam"),
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            ID: "053202c7-3d5f-4e9b-90c2-8909b3fe49a8",
            ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            SortCode: codatbankfeeds.String("provident"),
            Status: codatbankfeeds.String("nobis"),
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
