# UpdateBankFeed
Available in: `BankFeedAccounts`

Update a single BankFeed BankAccount for a single data source connected to a single company.

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
    req := operations.UpdateBankFeedRequest{
        BankFeedAccount: &shared.BankFeedAccount{
            AccountName: codatbankfeeds.String("magni"),
            AccountNumber: codatbankfeeds.String("sunt"),
            AccountType: codatbankfeeds.String("quo"),
            Balance: codatbankfeeds.Float64(8480.09),
            Currency: codatbankfeeds.String("pariatur"),
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
            ID: "c692601f-b576-4b0d-9f0d-30c5fbb25870",
            ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
            SortCode: codatbankfeeds.String("quis"),
            Status: codatbankfeeds.String("nesciunt"),
        },
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankFeedAccounts.UpdateBankFeed(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccount != nil {
        // handle response
    }
}
```