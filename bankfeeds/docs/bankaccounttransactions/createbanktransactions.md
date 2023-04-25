# CreateBankTransactions
Available in: `BankAccountTransactions`

Posts bank transactions to the accounting package for a given company.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support POST methods.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codatbankfeeds.String("numquam"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 4663.11,
                    Balance: 4746.97,
                    ClearedOnDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Counterparty: codatbankfeeds.String("velit"),
                    Description: codatbankfeeds.String("error"),
                    ID: codatbankfeeds.String("251aa52c-3f5a-4d01-9da1-ffe78f097b00"),
                    ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("reprehenderit"),
                    SourceModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumFee,
                },
                shared.BankTransactionLine{
                    Amount: 9795.87,
                    Balance: 1201.96,
                    ClearedOnDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Counterparty: codatbankfeeds.String("corporis"),
                    Description: codatbankfeeds.String("dolore"),
                    ID: codatbankfeeds.String("71b5e6e1-3b99-4d48-8e1e-91e450ad2abd"),
                    ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("labore"),
                    SourceModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumDiv,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(183191),
    }

    res, err := s.BankAccountTransactions.CreateBankTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```