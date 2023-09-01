<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codatsyncpayables.String("USD"),
            CurrentBalance: codatsyncpayables.Float64(0),
            Description: codatsyncpayables.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsyncpayables.String("Asset.Current"),
            FullyQualifiedName: codatsyncpayables.String("Fixed Asset"),
            ID: codatsyncpayables.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsyncpayables.Bool(false),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncpayables.String("Accounts Receivable"),
            NominalCode: codatsyncpayables.String("610"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "corrupti",
                        "illum",
                        "vel",
                        "error",
                    },
                    Property: codatsyncpayables.String("deserunt"),
                },
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "iure",
                        "magnam",
                    },
                    Property: codatsyncpayables.String("debitis"),
                },
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "delectus",
                    },
                    Property: codatsyncpayables.String("tempora"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(383441),
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