# GetCompanyConfiguration
Available in: `Configuration`

Gets a companies expense sync configuration

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCompanyConfigurationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Configuration.GetCompanyConfiguration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyConfiguration != nil {
        // handle response
    }
}
```