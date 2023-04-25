# IntiateSync
Available in: `Sync`

Initiate sync of pending transactions.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.IntiateSyncRequest{
        PostSync: &shared.PostSync{
            DatasetIds: []string{
                "7cc8796e-d151-4a05-9fc2-ddf7cc78ca1b",
                "a928fc81-6742-4cb7-b920-5929396fea75",
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Sync.IntiateSync(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncInitiated != nil {
        // handle response
    }
}
```