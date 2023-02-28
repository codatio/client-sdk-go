<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/common"
    "github.com/codatio/client-sdk-go/common/pkg/models/shared"
    "github.com/codatio/client-sdk-go/common/pkg/models/operations"
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
    
    req := operations.CreateCompanyRequest{
        Request: &operations.CreateCompanyRequestBody{
            Description: "unde",
            Name: "deserunt",
        },
    }

    ctx := context.Background()
    res, err := s.Companies.CreateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCompany200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->