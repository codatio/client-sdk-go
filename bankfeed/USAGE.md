<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/bankfeed"
    "github.com/codatio/client-sdk-go/bankfeed/pkg/models/shared"
    "github.com/codatio/client-sdk-go/bankfeed/pkg/models/operations"
)

func main() {
    opts := []codatio.SDKOption{
        codatio.WithSecurity(
            shared.Security{
                APIKey: shared.SchemeAPIKey{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            },
        ),
    }

    s := codatio.New(opts...)
    
    req := operations.GetBankFeedRequest{
        Security: operations.GetBankFeedSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.GetBankFeedPathParams{
            BankAccountID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
        Request: &shared.BankFeedBankAccount{
            AccountName: "nulla",
            AccountNumber: "id",
            AccountType: "vero",
            Balance: 5448.83,
            Currency: "nulla",
            FeedStartDate: "2022-09-25T21:58:56.179Z",
            ID: "fuga",
            ModifiedDate: "2022-07-06T19:09:57.259Z",
            SortCode: "eum",
            Status: "iusto",
        },
    }

    ctx := context.Background()
    res, err := s.BankFeeds.GetBankFeed(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedBankAccount != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->