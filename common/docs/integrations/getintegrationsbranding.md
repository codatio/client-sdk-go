# GetIntegrationsBranding
Available in: `Integrations`

Get branding for platform.

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
    req := operations.GetIntegrationsBrandingRequest{
        PlatformKey: "gbol",
    }

    res, err := s.Integrations.GetIntegrationsBranding(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Branding != nil {
        // handle response
    }
}
```