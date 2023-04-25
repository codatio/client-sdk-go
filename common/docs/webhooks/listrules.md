# ListRules
Available in: `Webhooks`

List webhooks that you are subscribed to.

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
    req := operations.ListRulesRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("architecto"),
    }

    res, err := s.Webhooks.ListRules(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rules != nil {
        // handle response
    }
}
```