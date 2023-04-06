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
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.ListAccountBalancesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: "-modifiedDate",
        Page: 1,
        PageSize: 100,
        Query: "corrupti",
    }

    ctx := context.Background()
    res, err := s.AccountBalances.ListAccountBalances(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountBalances != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->