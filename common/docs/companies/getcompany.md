# GetCompany
Available in: `Companies`

Get metadata for a single company

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCompanyRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Companies.GetCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```