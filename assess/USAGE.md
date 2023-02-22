<!-- Start SDK Example Usage -->
```go
package main

import (
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
            }
        ),
    }

    s := codatio.New(opts...)
    
    res, err := s.Categories.GetDataAssessAccountsCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Categories != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->