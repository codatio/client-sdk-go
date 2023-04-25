# CreateBankAccount
Available in: `BankAccounts`

Posts a new bank account to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating bank accounts.

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
    req := operations.CreateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("nihil"),
            AccountNumber: codataccounting.String("praesentium"),
            AccountType: shared.BankAccountBankAccountTypeEnumDebit.ToPointer(),
            AvailableBalance: codataccounting.Float64(557.14),
            Balance: codataccounting.Float64(6048.46),
            Currency: codataccounting.String("voluptate"),
            IBan: codataccounting.String("cum"),
            ID: codataccounting.String("0074f154-71b5-4e6e-93b9-9d488e1e91e4"),
            Institution: codataccounting.String("enim"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            NominalCode: codataccounting.String("consequatur"),
            OverdraftLimit: codataccounting.Float64(6674.11),
            SortCode: codataccounting.String("quibusdam"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(131797),
    }

    res, err := s.BankAccounts.CreateBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```