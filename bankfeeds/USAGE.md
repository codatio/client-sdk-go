<!-- Start SDK Example Usage -->
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
            AccountID: codatbankfeeds.String("corrupti"),
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 7151.9,
                    Balance: 8442.66,
                    ClearedOnDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Counterparty: codatbankfeeds.String("unde"),
                    Description: codatbankfeeds.String("nulla"),
                    ID: codatbankfeeds.String("8d69a674-e0f4-467c-8879-6ed151a05dfc"),
                    ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("odit"),
                    SourceModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumDirectDebit,
                },
                shared.BankTransactionLine{
                    Amount: 8700.88,
                    Balance: 9786.19,
                    ClearedOnDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Counterparty: codatbankfeeds.String("molestiae"),
                    Description: codatbankfeeds.String("quod"),
                    ID: codatbankfeeds.String("c78ca1ba-928f-4c81-a742-cb7392059293"),
                    ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("natus"),
                    SourceModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumSerChg,
                },
                shared.BankTransactionLine{
                    Amount: 9437.49,
                    Balance: 9025.99,
                    ClearedOnDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Counterparty: codatbankfeeds.String("fuga"),
                    Description: codatbankfeeds.String("in"),
                    ID: codatbankfeeds.String("596eb10f-aaa2-4352-8595-5907aff1a3a2"),
                    ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    Reconciled: false,
                    Reference: codatbankfeeds.String("repellat"),
                    SourceModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                    TransactionType: shared.BankTransactionTypeEnumCheck,
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(581850),
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
<!-- End SDK Example Usage -->