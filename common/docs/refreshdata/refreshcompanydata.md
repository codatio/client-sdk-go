# RefreshCompanyData
Available in: `RefreshData`

Refreshes all data types marked Fetch on first link.

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
    req := operations.RefreshCompanyDataRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.RefreshData.RefreshCompanyData(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```