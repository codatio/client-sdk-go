# ListAvailableAccountCategories
Available in: `Categories`

Lists available account categories Codat's categorisation engine can provide. 

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.ListAvailableAccountCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Categories != nil {
        // handle response
    }
}
```