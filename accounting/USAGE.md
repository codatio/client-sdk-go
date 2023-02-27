<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/accounting"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
    
    req := operations.GetAccountTransactionRequest{
        Security: operations.GetAccountTransactionSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.GetAccountTransactionPathParams{
            AccountTransactionID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
    }

    ctx := context.Background()
    res, err := s.AccountTransactions.GetAccountTransaction(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SourceModifiedDate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->