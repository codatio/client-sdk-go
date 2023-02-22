<!-- Start SDK Example Usage -->
```go
package main

import (
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
            }
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
            FeedStartDate: "2022-09-20T22:58:49.537Z",
            ID: "fuga",
            ModifiedDate: "2022-07-01T20:09:50.617Z",
            SortCode: "eum",
            Status: "iusto",
        },
    }
    
    res, err := s.BankFeeds.GetBankFeed(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedBankAccount != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->