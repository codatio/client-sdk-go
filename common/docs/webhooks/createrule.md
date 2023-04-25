# CreateRule
Available in: `Webhooks`

Create a new webhook configuration

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
    req := shared.Rule{
        CompanyID: codatcommon.String("39b73b17-cc2e-429e-915d-71654e9dcd1e"),
        ID: "ff89c50e-a719-4ef5-a182-9917e53927b6",
        Notifiers: shared.RuleNotifiers{
            Emails: []string{
                "info@client.com",
                "info@client.com",
                "info@client.com",
                "info@client.com",
            },
            Webhook: codatcommon.String("https://webhook.client.com"),
        },
        Type: "sapiente",
    }

    res, err := s.Webhooks.CreateRule(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rule != nil {
        // handle response
    }
}
```