<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/synccommerce"
    "github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
    "github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
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
    
    req := operations.AddDataConnectionRequest{
        PathParams: operations.AddDataConnectionPathParams{
            CompanyID: "unde",
        },
        Request: "deserunt",
    }

    ctx := context.Background()
    res, err := s.CompanyManagement.AddDataConnection(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddDataConnection200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->