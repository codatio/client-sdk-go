# BankAccounts

## Overview

Bank accounts

### Available Operations

* [CreateBankAccount](#createbankaccount) - Create bank account
* [GetAllBankAccount](#getallbankaccount) - Get bank account
* [GetBankAccount](#getbankaccount) - Get bank account
* [GetCreateUpdateBankAccountsModel](#getcreateupdatebankaccountsmodel) - Get create/update bank account model
* [ListBankAccounts](#listbankaccounts) - List bank accounts
* [UpdateBankAccount](#updatebankaccount) - Update bank account

## CreateBankAccount

Posts a new bank account to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating bank accounts.

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
    req := operations.CreateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("repellat"),
            AccountNumber: codataccounting.String("mollitia"),
            AccountType: shared.BankAccountBankAccountTypeEnumCredit.ToPointer(),
            AvailableBalance: codataccounting.Float64(2532.91),
            Balance: codataccounting.Float64(4143.69),
            Currency: codataccounting.String("quam"),
            IBan: codataccounting.String("molestiae"),
            ID: codataccounting.String("39251aa5-2c3f-45ad-819d-a1ffe78f097b"),
            Institution: codataccounting.String("perferendis"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("doloremque"),
            NominalCode: codataccounting.String("reprehenderit"),
            OverdraftLimit: codataccounting.Float64(2828.07),
            SortCode: codataccounting.String("maiores"),
            SourceModifiedDate: codataccounting.String("dicta"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(359444),
    }

    res, err := s.BankAccounts.CreateBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```

## GetAllBankAccount

Gets the bank account for given account ID.

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
    req := operations.GetAllBankAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Query: codataccounting.String("dolore"),
    }

    res, err := s.BankAccounts.GetAllBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankStatementAccount != nil {
        // handle response
    }
}
```

## GetBankAccount

Gets the bank account with a given ID

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
    req := operations.GetBankAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankAccounts.GetBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccount != nil {
        // handle response
    }
}
```

## GetCreateUpdateBankAccountsModel

Get create/update bank account model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating and updating bank accounts.

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
    req := operations.GetCreateUpdateBankAccountsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankAccounts.GetCreateUpdateBankAccountsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## ListBankAccounts

Gets the list of bank accounts for a given connection

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
    req := operations.ListBankAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("iusto"),
    }

    res, err := s.BankAccounts.ListBankAccounts(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccounts != nil {
        // handle response
    }
}
```

## UpdateBankAccount

Posts an updated bank account to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support updating bank accounts.

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
    req := operations.UpdateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("dicta"),
            AccountNumber: codataccounting.String("harum"),
            AccountType: shared.BankAccountBankAccountTypeEnumUnknown.ToPointer(),
            AvailableBalance: codataccounting.Float64(8804.76),
            Balance: codataccounting.Float64(4142.63),
            Currency: codataccounting.String("repudiandae"),
            IBan: codataccounting.String("quae"),
            ID: codataccounting.String("3b99d488-e1e9-41e4-90ad-2abd44269802"),
            Institution: codataccounting.String("assumenda"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ipsam"),
            NominalCode: codataccounting.String("alias"),
            OverdraftLimit: codataccounting.Float64(1464.41),
            SortCode: codataccounting.String("dolorum"),
            SourceModifiedDate: codataccounting.String("excepturi"),
        },
        BankAccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(270008),
    }

    res, err := s.BankAccounts.UpdateBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```
