<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := codatsyncpayroll.New(
        codatsyncpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codatsyncpayroll.String("USD"),
            CurrentBalance: codatsyncpayroll.Float64(0),
            Description: codatsyncpayroll.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsyncpayroll.String("Asset.Current"),
            FullyQualifiedName: codatsyncpayroll.String("Fixed Asset"),
            ID: codatsyncpayroll.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsyncpayroll.Bool(false),
            Metadata: &shared.AccountMetadata{
                IsDeleted: codatsyncpayroll.Bool(false),
            },
            ModifiedDate: codatsyncpayroll.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncpayroll.String("Accounts Receivable"),
            NominalCode: codatsyncpayroll.String("610"),
            SourceModifiedDate: codatsyncpayroll.String("2022-10-23T00:00:00.000Z"),
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
                    Property: codatsyncpayroll.String("deserunt"),
                },
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "iure",
                        "magnam",
                    },
                    Property: codatsyncpayroll.String("debitis"),
                },
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "delectus",
                    },
                    Property: codatsyncpayroll.String("tempora"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayroll.Int(383441),
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