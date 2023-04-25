# UpdateProfile
Available in: `Settings`

Update your Codat profile

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
    req := shared.Profile{
        AlertAuthHeader: codatcommon.String("explicabo"),
        APIKey: codatcommon.String("nobis"),
        ConfirmCompanyName: codatcommon.Bool(false),
        IconURL: codatcommon.String("enim"),
        LogoURL: codatcommon.String("omnis"),
        Name: "Bob's Burgers",
        RedirectURL: "nemo",
        WhiteListUrls: []string{
            "https://bobs-burgers.com/redirect",
            "https://bobs-burgers.com/redirect",
        },
    }

    res, err := s.Settings.UpdateProfile(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Profile != nil {
        // handle response
    }
}
```