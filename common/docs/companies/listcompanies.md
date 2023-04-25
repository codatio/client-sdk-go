# ListCompanies
Available in: `Companies`

List all companies that you have created in Codat.

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
    req := operations.ListCompaniesRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("iure"),
    }

    res, err := s.Companies.ListCompanies(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Companies != nil {
        // handle response
    }
}
```