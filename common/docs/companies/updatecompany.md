# UpdateCompany
Available in: `Companies`

Updates the given company with a new name and description

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateCompanyRequest{
        CompanyRequestBody: &shared.CompanyRequestBody{
            Description: codatcommon.String("magnam"),
            Name: "Larry Windler",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Companies.UpdateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```