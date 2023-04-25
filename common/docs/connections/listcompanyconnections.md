# ListCompanyConnections
Available in: `Connections`

List the connections for a company

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
    req := operations.ListCompanyConnectionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("minus"),
    }

    res, err := s.Connections.ListCompanyConnections(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```