# UpdateBankAccount
Available in: `BankAccounts`

Posts an updated bank account to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support updating bank accounts.

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
    req := operations.UpdateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("quibusdam"),
            AccountNumber: codataccounting.String("labore"),
            AccountType: shared.BankAccountBankAccountTypeEnumUnknown.ToPointer(),
            AvailableBalance: codataccounting.Float64(1831.91),
            Balance: codataccounting.Float64(3978.21),
            Currency: codataccounting.String("cupiditate"),
            IBan: codataccounting.String("quos"),
            ID: codataccounting.String("02d502a9-4bb4-4f63-8969-e9a3efa77dfb"),
            Institution: codataccounting.String("dicta"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            NominalCode: codataccounting.String("magnam"),
            OverdraftLimit: codataccounting.Float64(7670.24),
            SortCode: codataccounting.String("facere"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
        },
        BankAccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(411820),
    }

    res, err := s.BankAccounts.UpdateBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```