# GetCompanyDataStatus
Available in: `DataStatus`

Get the state of each data type for a company

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
    req := operations.GetCompanyDataStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DataStatus.GetCompanyDataStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataStatusResponse != nil {
        // handle response
    }
}
```