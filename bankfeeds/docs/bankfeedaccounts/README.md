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
                AccountName: codatbankfeeds.String("provident"),
                AccountNumber: codatbankfeeds.String("necessitatibus"),
                AccountType: codatbankfeeds.String("sint"),
                Balance: codatbankfeeds.Float64(6389.21),
                Currency: codatbankfeeds.String("dolor"),
                FeedStartDate: codatbankfeeds.String("debitis"),
                ID: "fa77dfb1-4cd6-46ae-b95e-fb9ba88f3a66",
                ModifiedDate: codatbankfeeds.String("natus"),
                SortCode: codatbankfeeds.String("omnis"),
                Status: codatbankfeeds.String("molestiae"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("perferendis"),
                AccountNumber: codatbankfeeds.String("nihil"),
                AccountType: codatbankfeeds.String("magnam"),
                Balance: codatbankfeeds.Float64(7160.75),
                Currency: codatbankfeeds.String("id"),
                FeedStartDate: codatbankfeeds.String("labore"),
                ID: "469b6e21-4195-4989-8afa-563e2516fe4c",
                ModifiedDate: codatbankfeeds.String("deleniti"),
                SortCode: codatbankfeeds.String("facilis"),
                Status: codatbankfeeds.String("in"),
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
            AccountName: codatbankfeeds.String("architecto"),
            AccountNumber: codatbankfeeds.String("architecto"),
            AccountType: codatbankfeeds.String("repudiandae"),
            Balance: codatbankfeeds.Float64(3523.12),
            Currency: codatbankfeeds.String("expedita"),
            FeedStartDate: codatbankfeeds.String("nihil"),
            ID: "fd2ed028-921c-4ddc-a926-01fb576b0d5f",
            ModifiedDate: codatbankfeeds.String("perferendis"),
            SortCode: codatbankfeeds.String("fugiat"),
            Status: codatbankfeeds.String("amet"),
        },
        AccountID: "aut",
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
