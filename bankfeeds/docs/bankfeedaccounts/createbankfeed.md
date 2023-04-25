# CreateBankFeed
Available in: `BankFeedAccounts`

Put BankFeed BankAccounts for a single data source connected to a single company.

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
    req := operations.CreateBankFeedRequest{
        RequestBody: []shared.BankFeedAccount{
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("quos"),
                AccountNumber: codatbankfeeds.String("perferendis"),
                AccountType: codatbankfeeds.String("magni"),
                Balance: codatbankfeeds.Float64(8289.4),
                Currency: codatbankfeeds.String("ipsam"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                ID: "02a94bb4-f63c-4969-a9a3-efa77dfb14cd",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                SortCode: codatbankfeeds.String("ea"),
                Status: codatbankfeeds.String("aliquid"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("laborum"),
                AccountNumber: codatbankfeeds.String("accusamus"),
                AccountType: codatbankfeeds.String("non"),
                Balance: codatbankfeeds.Float64(5812.73),
                Currency: codatbankfeeds.String("enim"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                ID: "efb9ba88-f3a6-4699-b074-ba4469b6e214",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                SortCode: codatbankfeeds.String("et"),
                Status: codatbankfeeds.String("excepturi"),
            },
            shared.BankFeedAccount{
                AccountName: codatbankfeeds.String("ullam"),
                AccountNumber: codatbankfeeds.String("provident"),
                AccountType: codatbankfeeds.String("quos"),
                Balance: codatbankfeeds.Float64(5743.25),
                Currency: codatbankfeeds.String("accusantium"),
                FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                ID: "afa563e2-516f-4e4c-8b71-1e5b7fd2ed02",
                ModifiedDate: codatbankfeeds.String("2022-10-23T00:00:00Z"),
                SortCode: codatbankfeeds.String("praesentium"),
                Status: codatbankfeeds.String("natus"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankFeedAccounts.CreateBankFeed(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccounts != nil {
        // handle response
    }
}
```