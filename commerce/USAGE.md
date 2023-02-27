<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/commerce"
    "github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
    "github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
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
    
    req := operations.GetCommerceInfoRequest{
        PathParams: operations.GetCommerceInfoPathParams{
            CompanyID: "unde",
            ConnectionID: "deserunt",
        },
    }

    ctx := context.Background()
    res, err := s.CompanyInfo.GetCommerceInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SourceModifiedDate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->