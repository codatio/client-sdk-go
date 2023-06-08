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
        CreateBankTransactions: &shared.CreateBankTransactions{
            AccountID: codatbankfeeds.String("corrupti"),
            Transactions: []shared.CreateBankAccountTransaction{
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(7151.9),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("unde"),
                    ID: codatbankfeeds.String("d8d69a67-4e0f-4467-8c87-96ed151a05df"),
                },
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(7781.57),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("at"),
                    ID: codatbankfeeds.String("df7cc78c-a1ba-4928-bc81-6742cb739205"),
                },
                shared.CreateBankAccountTransaction{
                    Amount: codatbankfeeds.Float64(6176.36),
                    Date: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
                    Description: codatbankfeeds.String("iste"),
                    ID: codatbankfeeds.String("396fea75-96eb-410f-aaa2-352c5955907a"),
                },
            },
        },
        AccountID: "EILBDVJVNUAGVKRQ",
        AllowSyncOnPushComplete: codatbankfeeds.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatbankfeeds.Int(958950),
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