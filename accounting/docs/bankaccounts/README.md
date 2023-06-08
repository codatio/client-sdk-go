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
            AccountName: codataccounting.String("perferendis"),
            AccountNumber: codataccounting.String("ad"),
            AccountType: shared.BankAccountBankAccountTypeCredit.ToPointer(),
            AvailableBalance: codataccounting.Float64(1496.75),
            Balance: codataccounting.Float64(6120.96),
            Currency: codataccounting.String("GBP"),
            IBan: codataccounting.String("natus"),
            ID: codataccounting.String("6fea7596-eb10-4faa-a235-2c5955907aff"),
            Institution: codataccounting.String("architecto"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            NominalCode: codataccounting.String("dolorem"),
            OverdraftLimit: codataccounting.Float64(6350.59),
            SortCode: codataccounting.String("consequuntur"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(653108),
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
        AccountID: "occaecati",
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
        Query: codataccounting.String("numquam"),
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
            AccountName: codataccounting.String("commodi"),
            AccountNumber: codataccounting.String("quam"),
            AccountType: shared.BankAccountBankAccountTypeCredit.ToPointer(),
            AvailableBalance: codataccounting.Float64(2444.25),
            Balance: codataccounting.Float64(6235.1),
            Currency: codataccounting.String("GBP"),
            IBan: codataccounting.String("quis"),
            ID: codataccounting.String("1aa52c3f-5ad0-419d-a1ff-e78f097b0074"),
            Institution: codataccounting.String("maiores"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            NominalCode: codataccounting.String("corporis"),
            OverdraftLimit: codataccounting.Float64(2961.4),
            SortCode: codataccounting.String("iusto"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
        },
        BankAccountID: "harum",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(317983),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```
