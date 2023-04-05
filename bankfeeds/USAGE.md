<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/bankfeeds"
    "github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
    "github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatio.New(
        codatio.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.CreateBankTransactionsRequest{
        BankTransactions: &shared.BankTransactions{
            AccountID: "corrupti",
            Transactions: []shared.BankTransactionLine{
                shared.BankTransactionLine{
                    Amount: 7151.9,
                    Balance: 8442.66,
                    ClearedOnDate: "2022-10-23T00:00:00Z",
                    Counterparty: "unde",
                    Description: "nulla",
                    ID: "corrupti",
                    ModifiedDate: "2022-10-23T00:00:00Z",
                    Reconciled: false,
                    Reference: "illum",
                    SourceModifiedDate: "2022-10-23T00:00:00Z",
                    TransactionType: "Dep",
                },
                shared.BankTransactionLine{
                    Amount: 6235.64,
                    Balance: 6458.94,
                    ClearedOnDate: "2022-10-23T00:00:00Z",
                    Counterparty: "suscipit",
                    Description: "iure",
                    ID: "magnam",
                    ModifiedDate: "2022-10-23T00:00:00Z",
                    Reconciled: false,
                    Reference: "debitis",
                    SourceModifiedDate: "2022-10-23T00:00:00Z",
                    TransactionType: "Credit",
                },
                shared.BankTransactionLine{
                    Amount: 9636.63,
                    Balance: 2726.56,
                    ClearedOnDate: "2022-10-23T00:00:00Z",
                    Counterparty: "suscipit",
                    Description: "molestiae",
                    ID: "minus",
                    ModifiedDate: "2022-10-23T00:00:00Z",
                    Reconciled: false,
                    Reference: "placeat",
                    SourceModifiedDate: "2022-10-23T00:00:00Z",
                    TransactionType: "Pos",
                },
            },
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        AllowSyncOnPushComplete: false,
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: 479977,
    }

    ctx := context.Background()
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