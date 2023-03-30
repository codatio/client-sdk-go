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
    s := codatio.New(
        codatio.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.GetCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    ctx := context.Background()
    res, err := s.CompanyInfo.GetCompanyInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyInfo != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->