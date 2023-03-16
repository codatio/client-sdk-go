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
    s := codatio.New(
        WithSecurity(        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

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