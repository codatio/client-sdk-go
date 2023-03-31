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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := shared.CreateCompany{
        Name: "Bob's Burgers",
    }

    ctx := context.Background()
    res, err := s.CompanyManagement.CreateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->