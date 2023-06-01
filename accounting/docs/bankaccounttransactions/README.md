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
            AccountID: codataccounting.String("veritatis"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 202.18,
                    Balance: 3682.41,
                    ClearedOnDate: codataccounting.String("repellendus"),
                    Counterparty: codataccounting.String("sapiente"),
                    Description: codataccounting.String("quo"),
                    ID: codataccounting.String("2ddf7cc7-8ca1-4ba9-a8fc-816742cb7392"),
                    ModifiedDate: codataccounting.String("perferendis"),
                    Reconciled: false,
                    Reference: codataccounting.String("ad"),
                    SourceModifiedDate: codataccounting.String("natus"),
                    TransactionType: shared.BankTransactionTypeDebit,
                },
                shared.BankTransactionLine{
                    Amount: 6120.96,
                    Balance: 2223.21,
                    ClearedOnDate: codataccounting.String("natus"),
                    Counterparty: codataccounting.String("laboriosam"),
                    Description: codataccounting.String("hic"),
                    ID: codataccounting.String("ea7596eb-10fa-4aa2-b52c-5955907aff1a"),
                    ModifiedDate: codataccounting.String("dolorem"),
                    Reconciled: false,
                    Reference: codataccounting.String("culpa"),
                    SourceModifiedDate: codataccounting.String("consequuntur"),
                    TransactionType: shared.BankTransactionTypeOther,
                },
                shared.BankTransactionLine{
                    Amount: 6531.08,
                    Balance: 5818.5,
                    ClearedOnDate: codataccounting.String("numquam"),
                    Counterparty: codataccounting.String("commodi"),
                    Description: codataccounting.String("quam"),
                    ID: codataccounting.String("739251aa-52c3-4f5a-9019-da1ffe78f097"),
                    ModifiedDate: codataccounting.String("cum"),
                    Reconciled: false,
                    Reference: codataccounting.String("perferendis"),
                    SourceModifiedDate: codataccounting.String("doloremque"),
                    TransactionType: shared.BankTransactionTypeDep,
                },
            },
        },
        AccountID: "ut",
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(979587),
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
        AccountID: "dicta",
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
        AccountID: "corporis",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dolore"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```
