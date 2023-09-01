<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AdvancedControls.CreateCompany(ctx, shared.CreateCompany{
        Name: "Bob's Burgers",
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