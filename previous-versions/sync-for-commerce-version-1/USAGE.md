<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingAccounts.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
        AccountingAccount: &shared.AccountingAccount{
            Currency: syncforcommerceversion1.String("GBP"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: syncforcommerceversion1.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: syncforcommerceversion1.String("Asset.Current"),
            FullyQualifiedName: syncforcommerceversion1.String("Cash On Hand"),
            ID: syncforcommerceversion1.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            Metadata: &shared.AccountingAccountMetadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Name: syncforcommerceversion1.String("Accounts Receivable"),
            NominalCode: syncforcommerceversion1.String("610"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountingAccountValidDataTypeLinks{
                shared.AccountingAccountValidDataTypeLinks{
                    Links: []string{
                        "Gasoline",
                    },
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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