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
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CompanyRequestBody{
        Description: "corrupti",
        Name: "Kelvin Sporer",
    }

    res, err := s.Companies.CreateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->