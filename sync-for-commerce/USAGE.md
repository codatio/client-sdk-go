<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AdvancedControls.CreateCompany(ctx, shared.CreateCompany{
        Description: codatsynccommerce.String("Requested early access to the new financing scheme."),
        Name: "Bank of Dave",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->