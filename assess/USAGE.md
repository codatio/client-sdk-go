<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/assess"
    "github.com/codatio/client-sdk-go/assess/pkg/models/shared"
    "github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatio.New(
        WithSecurity(        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.GetAccountCategoryRequest{
        PathParams: operations.GetAccountCategoryPathParams{
            AccountID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
    }

    ctx := context.Background()
    res, err := s.Categories.GetAccountCategory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->