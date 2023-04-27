<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CreateCompany{
        Name: "Bob's Burgers",
    }

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