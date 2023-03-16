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
        WithSecurity(        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.GetBankAccountPushOptionsRequest{
        PathParams: operations.GetBankAccountPushOptionsPathParams{
            AccountID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
        QueryParams: operations.GetBankAccountPushOptionsQueryParams{
            OrderBy: "nulla",
            Page: 6027.63,
            PageSize: 8579.46,
            Query: "perspiciatis",
        },
    }

    ctx := context.Background()
    res, err := s.BankAccountTransactions.GetBankAccountPushOptions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->