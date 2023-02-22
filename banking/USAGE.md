<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/codatio/client-sdk-go/banking"
    "github.com/codatio/client-sdk-go/banking/pkg/models/shared"
    "github.com/codatio/client-sdk-go/banking/pkg/models/operations"
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
    
    req := operations.ListBankingAccountBalancesRequest{
        Security: operations.ListBankingAccountBalancesSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.ListBankingAccountBalancesPathParams{
            CompanyID: "unde",
            ConnectionID: "deserunt",
        },
        QueryParams: operations.ListBankingAccountBalancesQueryParams{
            OrderBy: "porro",
            Page: 8442.66,
            PageSize: 6027.63,
            Query: "vero",
        },
    }
    
    res, err := s.AccountBalances.ListBankingAccountBalances(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Links != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->