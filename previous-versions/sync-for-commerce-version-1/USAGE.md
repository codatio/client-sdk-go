<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingAccounts.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
        AccountingAccount: &shared.AccountingAccount{
            Currency: codatsynccommerce.String("USD"),
            CurrentBalance: codatsynccommerce.Float64(0),
            Description: codatsynccommerce.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsynccommerce.String("Asset.Current"),
            FullyQualifiedName: codatsynccommerce.String("Fixed Asset"),
            ID: codatsynccommerce.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsynccommerce.Bool(false),
            Metadata: &shared.AccountingAccountMetadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Name: codatsynccommerce.String("Accounts Receivable"),
            NominalCode: codatsynccommerce.String("610"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountingAccountValidDataTypeLinks{
                shared.AccountingAccountValidDataTypeLinks{
                    Links: []string{
                        "corrupti",
                        "illum",
                        "vel",
                        "error",
                    },
                    Property: codatsynccommerce.String("deserunt"),
                },
                shared.AccountingAccountValidDataTypeLinks{
                    Links: []string{
                        "iure",
                        "magnam",
                    },
                    Property: codatsynccommerce.String("debitis"),
                },
                shared.AccountingAccountValidDataTypeLinks{
                    Links: []string{
                        "delectus",
                    },
                    Property: codatsynccommerce.String("tempora"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(383441),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateAccountResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->