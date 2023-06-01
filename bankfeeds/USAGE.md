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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccountTransactions.Create(ctx, operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: codatbankfeeds.String("corrupti"),
            Amount: codatbankfeeds.Float64(5928.45),
            Balance: codatbankfeeds.Float64(7151.9),
            ClearedOnDate: codatbankfeeds.String("quibusdam"),
            Description: codatbankfeeds.String("unde"),
            ID: codatbankfeeds.String("d8d69a67-4e0f-4467-8c87-96ed151a05df"),
            ModifiedDate: codatbankfeeds.String("quo"),
            Reconciled: codatbankfeeds.Bool(false),
            SourceModifiedDate: codatbankfeeds.String("odit"),
            TransactionType: shared.BankTransactionTypeDirectDebit.ToPointer(),
        },
        AccountID: "at",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(978619),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankTransactionsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->