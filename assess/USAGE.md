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

    ctx := context.Background()
    res, err := s.Categories.GetDataAssessAccountsCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDataAssessAccountsCategoriesChartOfAccountCategoryAllOfs != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->