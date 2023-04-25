# GetProfile
Available in: `Settings`

Fetch your Codat profile.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.GetProfile(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Profile != nil {
        // handle response
    }
}
```