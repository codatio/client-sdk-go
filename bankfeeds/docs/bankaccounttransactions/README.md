# BankAccountTransactions

## Overview

Bank feed bank accounts

### Available Operations

* [CreateBankTransactions](#createbanktransactions) - Create bank transactions
* [GetCreateBankAccountModel](#getcreatebankaccountmodel) - List push options for bank account bank transactions
* [ListBankAccountTransactions](#listbankaccounttransactions) - List bank transactions for bank account

## CreateBankTransactions

Posts bank transactions to the accounting package for a given company.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support POST methods.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.CreateBankTransactions(ctx, operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codatbankfeeds.String("laborum"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 3172.02,
                    Balance: 1381.83,
                    ClearedOnDate: codatbankfeeds.String("quo"),
                    Counterparty: codatbankfeeds.String("sequi"),
                    Description: codatbankfeeds.String("tenetur"),
                    ID: codatbankfeeds.String("5ad019da-1ffe-478f-897b-0074f15471b5"),
                    ModifiedDate: codatbankfeeds.String("accusamus"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("commodi"),
                    SourceModifiedDate: codatbankfeeds.String("repudiandae"),
                    TransactionType: shared.BankTransactionTypeEnumCredit,
                },
                shared.BankTransactionLine{
                    Amount: 2168.22,
                    Balance: 6924.72,
                    ClearedOnDate: codatbankfeeds.String("molestias"),
                    Counterparty: codatbankfeeds.String("excepturi"),
                    Description: codatbankfeeds.String("pariatur"),
                    ID: codatbankfeeds.String("488e1e91-e450-4ad2-abd4-4269802d502a"),
                    ModifiedDate: codatbankfeeds.String("excepturi"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("tempora"),
                    SourceModifiedDate: codatbankfeeds.String("facilis"),
                    TransactionType: shared.BankTransactionTypeEnumCash,
                },
                shared.BankTransactionLine{
                    Amount: 2884.76,
                    Balance: 9621.89,
                    ClearedOnDate: codatbankfeeds.String("eum"),
                    Counterparty: codatbankfeeds.String("non"),
                    Description: codatbankfeeds.String("eligendi"),
                    ID: codatbankfeeds.String("969e9a3e-fa77-4dfb-94cd-66ae395efb9b"),
                    ModifiedDate: codatbankfeeds.String("id"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("blanditiis"),
                    SourceModifiedDate: codatbankfeeds.String("deleniti"),
                    TransactionType: shared.BankTransactionTypeEnumOther,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(230533),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

## GetCreateBankAccountModel

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.GetCreateBankAccountModel(ctx, operations.GetCreateBankAccountModelRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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

## ListBankAccountTransactions

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.ListBankAccountTransactions(ctx, operations.ListBankAccountTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbankfeeds.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbankfeeds.Int(100),
        Query: codatbankfeeds.String("deserunt"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```
