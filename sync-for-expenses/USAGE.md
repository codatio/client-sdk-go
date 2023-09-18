<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codatsyncexpenses.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: codatsyncexpenses.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsyncexpenses.String("Asset.Current"),
            FullyQualifiedName: codatsyncexpenses.String("Fixed Asset"),
            ID: codatsyncexpenses.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsyncexpenses.Bool(false),
            Metadata: &shared.AccountMetadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncexpenses.String("Accounts Receivable"),
            NominalCode: codatsyncexpenses.String("610"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "unde",
                    },
                    Property: codatsyncexpenses.String("nulla"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncexpenses.Int(544883),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->