# GetWebhook
Available in: `Webhooks`

Get a single webhook

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
    req := operations.GetWebhookRequest{
        RuleID: "7318949f-c008-4936-a8ff-10d7ab563fa6",
    }

    res, err := s.Webhooks.GetWebhook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rule != nil {
        // handle response
    }
}
```