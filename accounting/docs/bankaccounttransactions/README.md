# BankAccountTransactions

## Overview

Bank transactions for bank accounts

### Available Operations

* [CreateBankTransactions](#createbanktransactions) - Create bank transactions
* [GetCreateBankAccountModel](#getcreatebankaccountmodel) - List push options for bank account bank transactions
* [ListBankAccountTransactions](#listbankaccounttransactions) - List bank transactions for bank account
* [ListBankTransactions](#listbanktransactions) - List all bank transactions

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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codataccounting.String("excepturi"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 187.89,
                    Balance: 3241.41,
                    ClearedOnDate: codataccounting.String("natus"),
                    Counterparty: codataccounting.String("sed"),
                    Description: codataccounting.String("iste"),
                    ID: codataccounting.String("396fea75-96eb-410f-aaa2-352c5955907a"),
                    ModifiedDate: codataccounting.String("doloribus"),
                    Reconciled: false,
                    Reference: codataccounting.String("sapiente"),
                    SourceModifiedDate: codataccounting.String("architecto"),
                    TransactionType: shared.BankTransactionTypeEnumCheck,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(208876),
    }

    res, err := s.BankAccountTransactions.CreateBankTransactions(ctx, req)
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateBankAccountModelRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankAccountTransactions.GetCreateBankAccountModel(ctx, req)
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListBankAccountTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("culpa"),
    }

    res, err := s.BankAccountTransactions.ListBankAccountTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```

## ListBankTransactions

Gets the latest bank transactions for given account ID and company. Doesn't require connection ID.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListBankTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("consequuntur"),
    }

    res, err := s.BankAccountTransactions.ListBankTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountTransactions != nil {
        // handle response
    }
}
```
