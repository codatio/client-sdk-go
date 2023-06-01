# BankAccounts

## Overview

Bank accounts

### Available Operations

* [Create](#create) - Create bank account
* [~~Get~~](#get) - Get bank account :warning: **Deprecated**
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bank account model
* [List](#list) - List bank accounts
* [Update](#update) - Update bank account

## Create

Posts a new bank account to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bank account model](https://docs.codat.io/accounting-api#/operations/get-create-update-bankAccounts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) to see which integrations support this endpoint.


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
    res, err := s.BankAccounts.Create(ctx, operations.CreateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("iusto"),
            AccountNumber: codataccounting.String("dicta"),
            AccountType: shared.BankAccountBankAccountTypeDebit.ToPointer(),
            AvailableBalance: codataccounting.Float64(3179.83),
            Balance: codataccounting.Float64(8804.76),
            Currency: codataccounting.String("commodi"),
            IBan: codataccounting.String("repudiandae"),
            ID: codataccounting.String("13b99d48-8e1e-491e-850a-d2abd4426980"),
            Institution: codataccounting.String("magni"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("assumenda"),
            NominalCode: codataccounting.String("ipsam"),
            OverdraftLimit: codataccounting.Float64(46.95),
            SortCode: codataccounting.String("fugit"),
            SourceModifiedDate: codataccounting.String("dolorum"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(569618),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```

## ~~Get~~

Gets the bank account with a given ID

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.BankAccounts.Get(ctx, operations.GetBankAccountRequest{
        AccountID: "tempora",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.List(ctx, operations.ListBankAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("facilis"),
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
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support updating bank accounts.

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
    res, err := s.BankAccounts.Update(ctx, operations.UpdateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("tempore"),
            AccountNumber: codataccounting.String("labore"),
            AccountType: shared.BankAccountBankAccountTypeDebit.ToPointer(),
            AvailableBalance: codataccounting.Float64(4332.88),
            Balance: codataccounting.Float64(2487.53),
            Currency: codataccounting.String("eligendi"),
            IBan: codataccounting.String("sint"),
            ID: codataccounting.String("69e9a3ef-a77d-4fb1-8cd6-6ae395efb9ba"),
            Institution: codataccounting.String("blanditiis"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("deleniti"),
            NominalCode: codataccounting.String("sapiente"),
            OverdraftLimit: codataccounting.Float64(2305.33),
            SortCode: codataccounting.String("deserunt"),
            SourceModifiedDate: codataccounting.String("nisi"),
        },
        BankAccountID: "vel",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(618809),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```
