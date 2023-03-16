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
    s := codatio.New(
        WithSecurity(        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

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