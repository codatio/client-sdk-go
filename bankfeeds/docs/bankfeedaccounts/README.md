# BankFeedAccounts

## Overview

Bank feed bank accounts

### Available Operations

* [CreateBankFeed](#createbankfeed) - Create bank feed bank accounts
* [GetBankFeeds](#getbankfeeds) - List bank feed bank accounts
* [UpdateBankFeed](#updatebankfeed) - Update bank feed bank account

## CreateBankFeed

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
    res, err := s.BankFeedAccounts.CreateBankFeed(ctx, operations.CreateBankFeedRequest{
        RequestBody: []shared.BankFeedAccount{
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("vel"),
                AccountNumber: codatbankfeeds.String("natus"),
                AccountType: codatbankfeeds.String("omnis"),
                Balance: codatbankfeeds.Float64(4748.67),
                Currency: codatbankfeeds.String("perferendis"),
                FeedStartDate: codatbankfeeds.String("nihil"),
                ID: "4ba4469b-6e21-4419-9989-0afa563e2516",
                ModifiedDate: codatbankfeeds.String("doloribus"),
                SortCode: codatbankfeeds.String("debitis"),
                Status: codatbankfeeds.String("eius"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("maxime"),
                AccountNumber: codatbankfeeds.String("deleniti"),
                AccountType: codatbankfeeds.String("facilis"),
                Balance: codatbankfeeds.Float64(4479.26),
                Currency: codatbankfeeds.String("architecto"),
                FeedStartDate: codatbankfeeds.String("architecto"),
                ID: "e5b7fd2e-d028-4921-8ddc-692601fb576b",
                ModifiedDate: codatbankfeeds.String("eaque"),
                SortCode: codatbankfeeds.String("pariatur"),
                Status: codatbankfeeds.String("nemo"),
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

## GetBankFeeds

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
    res, err := s.BankFeedAccounts.GetBankFeeds(ctx, operations.GetBankFeedsRequest{
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

## UpdateBankFeed

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
    res, err := s.BankFeedAccounts.UpdateBankFeed(ctx, operations.UpdateBankFeedRequest{
        BankFeedAccount: &shared.BankFeedAccount{
            AccountName: codatbankfeeds.String("voluptatibus"),
            AccountNumber: codatbankfeeds.String("perferendis"),
            AccountType: codatbankfeeds.String("fugiat"),
            Balance: codatbankfeeds.Float64(2307.42),
            Currency: codatbankfeeds.String("aut"),
            FeedStartDate: codatbankfeeds.String("cumque"),
            ID: "5fbb2587-0532-402c-b3d5-fe9b90c28909",
            ModifiedDate: codatbankfeeds.String("rerum"),
            SortCode: codatbankfeeds.String("adipisci"),
            Status: codatbankfeeds.String("asperiores"),
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
