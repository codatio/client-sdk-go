# Accounts

## Overview

Accounts

### Available Operations

* [CreateAccount](#createaccount) - Create account
* [GetAccount](#getaccount) - Get account
* [GetCreateChartOfAccountsModel](#getcreatechartofaccountsmodel) - Get create account model
* [ListAccounts](#listaccounts) - List accounts

## CreateAccount

Creates a new account for a given company.

Required data may vary by integration. To see what data to post, first call [Get create account model](https://docs.codat.io/accounting-api#/operations/get-create-chartOfAccounts-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts) for integrations that support creating an account.

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
    req := operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codataccounting.String("quibusdam"),
            CurrentBalance: codataccounting.Float64(6027.63),
            Description: codataccounting.String("nulla"),
            FullyQualifiedCategory: codataccounting.String("corrupti"),
            FullyQualifiedName: codataccounting.String("illum"),
            ID: codataccounting.String("69a674e0-f467-4cc8-b96e-d151a05dfc2d"),
            IsBankAccount: false,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("at"),
            Name: codataccounting.String("Javier Schmidt"),
            NominalCode: codataccounting.String("totam"),
            SourceModifiedDate: codataccounting.String("porro"),
            Status: shared.AccountStatusEnumArchived,
            Type: shared.AccountTypeEnumUnknown,
            ValidDatatypeLinks: []shared.ValidDataTypeLinks{
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "occaecati",
                        "fugit",
                        "deleniti",
                    },
                    Property: codataccounting.String("hic"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "totam",
                        "beatae",
                        "commodi",
                        "molestiae",
                    },
                    Property: codataccounting.String("modi"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "impedit",
                    },
                    Property: codataccounting.String("cum"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(456150),
    }

    res, err := s.Accounts.CreateAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```

## GetAccount

Gets a single account corresponding to the given ID.

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
    req := operations.GetAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Accounts.GetAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

## GetCreateChartOfAccountsModel

Get create account model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts) for integrations that support creating an account.

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
    req := operations.GetCreateChartOfAccountsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Accounts.GetCreateChartOfAccountsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## ListAccounts

Gets the latest accounts for a company

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
    req := operations.ListAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ipsum"),
    }

    res, err := s.Accounts.ListAccounts(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Accounts != nil {
        // handle response
    }
}
```
