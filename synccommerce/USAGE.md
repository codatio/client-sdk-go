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
        codatio.WithSecurity(shared.Security{
            AuthHeader: "Basic YOUR_ENCODED_API_KEY",
        }),
    )

    req := operations.AddDataConnectionRequest{
        RequestBody: "unde",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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