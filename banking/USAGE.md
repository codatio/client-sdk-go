<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/banking"
    "github.com/codatio/client-sdk-go/banking/pkg/models/shared"
    "github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatio.New(codatio.WithSecurity(
        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
    ))
    
    req := operations.ListBankingAccountBalancesRequest{
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

    ctx := context.Background()
    res, err := s.AccountBalances.ListBankingAccountBalances(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Links != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->