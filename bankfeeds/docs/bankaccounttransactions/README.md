# BankAccountTransactions

## Overview

Bank feed bank accounts

### Available Operations

* [Create](#create) - Create bank transactions
* [Get](#get) - List push options for bank account bank transactions
* [List](#list) - List bank transactions for bank account

## Create

Posts bank transactions to the accounting package for a given company.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) to see which integrations support this endpoint.

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
    res, err := s.BankAccountTransactions.Create(ctx, operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codatbankfeeds.String("animi"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 1381.83,
                    Balance: 7783.46,
                    ClearedOnDate: codatbankfeeds.String("sequi"),
                    Counterparty: codatbankfeeds.String("tenetur"),
                    Description: codatbankfeeds.String("ipsam"),
                    ID: codatbankfeeds.String("ad019da1-ffe7-48f0-97b0-074f15471b5e"),
                    ModifiedDate: codatbankfeeds.String("commodi"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("repudiandae"),
                    SourceModifiedDate: codatbankfeeds.String("quae"),
                    TransactionType: shared.BankTransactionTypeInt,
                },
                shared.BankTransactionLine{
                    Amount: 6924.72,
                    Balance: 5651.89,
                    ClearedOnDate: codatbankfeeds.String("excepturi"),
                    Counterparty: codatbankfeeds.String("pariatur"),
                    Description: codatbankfeeds.String("modi"),
                    ID: codatbankfeeds.String("88e1e91e-450a-4d2a-bd44-269802d502a9"),
                    ModifiedDate: codatbankfeeds.String("tempora"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("facilis"),
                    SourceModifiedDate: codatbankfeeds.String("tempore"),
                    TransactionType: shared.BankTransactionTypeFee,
                },
            },
        },
        AccountID: "delectus",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(433288),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

## Get

Gets the options of pushing bank account transactions.

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
    res, err := s.BankAccountTransactions.Get(ctx, operations.GetCreateBankAccountModelRequest{
        AccountID: "non",
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

## List

Gets bank transactions for a given bank account ID

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
    res, err := s.BankAccountTransactions.List(ctx, operations.ListBankAccountTransactionsRequest{
        AccountID: "eligendi",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbankfeeds.String("-modifiedDate"),
        Page: codatbankfeeds.Int(1),
        PageSize: codatbankfeeds.Int(100),
        Query: codatbankfeeds.String("sint"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```
