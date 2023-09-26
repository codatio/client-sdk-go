<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/types"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: syncforpayroll.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: syncforpayroll.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: syncforpayroll.String("Asset.Current"),
            FullyQualifiedName: syncforpayroll.String("Fixed Asset"),
            ID: syncforpayroll.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: syncforpayroll.Bool(false),
            Metadata: &shared.AccountMetadata{
                IsDeleted: syncforpayroll.Bool(false),
            },
            ModifiedDate: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            Name: syncforpayroll.String("Accounts Receivable"),
            NominalCode: syncforpayroll.String("610"),
            SourceModifiedDate: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "unde",
                    },
                    Property: syncforpayroll.String("nulla"),
                },
            },
        },
        AllowSyncOnPushComplete: syncforpayroll.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforpayroll.Int(544883),
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