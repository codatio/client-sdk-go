# BankAccounts

## Overview

Bank accounts

### Available Operations

* [Create](#create) - Create bank account
* [Get](#get) - Get bank account
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bank account model
* [List](#list) - List bank accounts
* [Update](#update) - Update bank account

## Create

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
    res, err := s.BankAccounts.Create(ctx, operations.CreateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("natus"),
            AccountNumber: codataccounting.String("laboriosam"),
            AccountType: shared.BankAccountBankAccountTypeEnumDebit.ToPointer(),
            AvailableBalance: codataccounting.Float64(9025.99),
            Balance: codataccounting.Float64(6818.2),
            Currency: codataccounting.String("in"),
            IBan: codataccounting.String("corporis"),
            ID: codataccounting.String("96eb10fa-aa23-452c-9955-907aff1a3a2f"),
            Institution: codataccounting.String("mollitia"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("occaecati"),
            NominalCode: codataccounting.String("numquam"),
            OverdraftLimit: codataccounting.Float64(4143.69),
            SortCode: codataccounting.String("quam"),
            SourceModifiedDate: codataccounting.String("molestiae"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(244425),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```

## Get

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
    res, err := s.BankAccounts.Get(ctx, operations.GetBankAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccount != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

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
    res, err := s.BankAccounts.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBankAccountsModelRequest{
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
    res, err := s.BankAccounts.List(ctx, operations.ListBankAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("error"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccounts != nil {
        // handle response
    }
}
```

## Update

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
    res, err := s.BankAccounts.Update(ctx, operations.UpdateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("quia"),
            AccountNumber: codataccounting.String("quis"),
            AccountType: shared.BankAccountBankAccountTypeEnumUnknown.ToPointer(),
            AvailableBalance: codataccounting.Float64(6747.52),
            Balance: codataccounting.Float64(6563.3),
            Currency: codataccounting.String("enim"),
            IBan: codataccounting.String("odit"),
            ID: codataccounting.String("c3f5ad01-9da1-4ffe-b8f0-97b0074f1547"),
            Institution: codataccounting.String("dicta"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("harum"),
            NominalCode: codataccounting.String("enim"),
            OverdraftLimit: codataccounting.Float64(8804.76),
            SortCode: codataccounting.String("commodi"),
            SourceModifiedDate: codataccounting.String("repudiandae"),
        },
        BankAccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(64147),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```
