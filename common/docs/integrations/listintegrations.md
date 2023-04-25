# ListIntegrations
Available in: `Integrations`

List your available integrations

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
    req := operations.ListIntegrationsRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("veritatis"),
    }

    res, err := s.Integrations.ListIntegrations(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```