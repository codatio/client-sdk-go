# CreateCompany
Available in: `Companies`

Create a new company

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CompanyRequestBody{
        Description: codatcommon.String("corrupti"),
        Name: "Ben Mueller",
    }

    res, err := s.Companies.CreateCompany(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```