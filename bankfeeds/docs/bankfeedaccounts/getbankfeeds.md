# GetBankFeeds
Available in: `BankFeedAccounts`

Get BankFeed BankAccounts for a single data source connected to a single company.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetBankFeedsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BankFeedAccounts.GetBankFeeds(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccounts != nil {
        // handle response
    }
}
```