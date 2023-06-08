# BankAccountTransactions

## Overview

Bank transactions for bank accounts

### Available Operations

* [Create](#create) - Create bank transactions
* [GetCreateModel](#getcreatemodel) - List push options for bank account bank transactions
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.Create(ctx, operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codataccounting.String("excepturi"),
            Amount: codataccounting.Float64(3927.85),
            Balance: codataccounting.Float64(9255.97),
            ClearedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Description: codataccounting.String("ab"),
            ID: codataccounting.String("51a05dfc-2ddf-47cc-b8ca-1ba928fc8167"),
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Reconciled: codataccounting.Bool(false),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            TransactionType: shared.BankTransactionTypeCash.ToPointer(),
        },
        AccountID: "cum",
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(456150),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

## GetCreateModel

Gets the options of pushing bank account transactions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.GetCreateModel(ctx, operations.GetCreateBankAccountModelRequest{
        AccountID: "ipsum",
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.List(ctx, operations.ListBankAccountTransactionsRequest{
        AccountID: "excepturi",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("aspernatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```
