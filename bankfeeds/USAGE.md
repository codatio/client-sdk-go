<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/codatio/client-sdk-go/bankfeeds"
    "github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
    "github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
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
    
    req := operations.GetBankAccountPushOptionsRequest{
        Security: operations.GetBankAccountPushOptionsSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
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
    
    res, err := s.BankAccountTransactions.GetBankAccountPushOptions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->