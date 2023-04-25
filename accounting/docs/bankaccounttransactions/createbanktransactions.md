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
    req := operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codataccounting.String("cum"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 2165.5,
                    Balance: 5684.34,
                    ClearedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Counterparty: codataccounting.String("aspernatur"),
                    Description: codataccounting.String("perferendis"),
                    ID: codataccounting.String("5929396f-ea75-496e-b10f-aaa2352c5955"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codataccounting.String("excepturi"),
                    SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumUnknown,
                },
                shared.BankTransactionLine{
                    Amount: 4386.01,
                    Balance: 6342.74,
                    ClearedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Counterparty: codataccounting.String("doloribus"),
                    Description: codataccounting.String("sapiente"),
                    ID: codataccounting.String("1a3a2fa9-4677-4392-91aa-52c3f5ad019d"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codataccounting.String("laborum"),
                    SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumCredit,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(971945),
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