# CreateAccount
Available in: `Accounts`

Creates a new account for a given company.

Required data may vary by integration. To see what data to post, first call [Get create account model](https://docs.codat.io/accounting-api#/operations/get-create-chartOfAccounts-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts) for integrations that support creating an account.

## Example Usage
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
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Name: codataccounting.String("Emilio Krajcik"),
            NominalCode: codataccounting.String("esse"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.AccountStatusEnumArchived,
            Type: shared.AccountTypeEnumLiability,
            ValidDatatypeLinks: []shared.ValidDataTypeLinks{
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "nam",
                    },
                    Property: codataccounting.String("officia"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "fugit",
                        "deleniti",
                        "hic",
                    },
                    Property: codataccounting.String("optio"),
                },
                shared.ValidDataTypeLinks{
                    Links: []string{
                        "beatae",
                        "commodi",
                        "molestiae",
                    },
                    Property: codataccounting.String("modi"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(186332),
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