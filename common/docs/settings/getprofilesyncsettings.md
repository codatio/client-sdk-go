# GetProfileSyncSettings
Available in: `Settings`

Retrieve the sync settings for your client. This includes how often data types should be queued to be updated, and how much history should be fetched.

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
    res, err := s.Settings.GetProfileSyncSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSettings != nil {
        // handle response
    }
}
```