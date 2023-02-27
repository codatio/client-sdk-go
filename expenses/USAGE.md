<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/expenses"
    "github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
    "github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    opts := []codatio.SDKOption{
        codatio.WithSecurity(
            shared.Security{
                Authorization: shared.SchemeAuthorization{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            },
        ),
    }

    s := codatio.New(opts...)
    
    req := operations.GetCompanyConfigurationRequest{
        PathParams: operations.GetCompanyConfigurationPathParams{
            CompanyID: "unde",
        },
    }

    ctx := context.Background()
    res, err := s.Configuration.GetCompanyConfiguration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCompanyConfiguration200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->