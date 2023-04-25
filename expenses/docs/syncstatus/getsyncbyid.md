# GetSyncByID
Available in: `SyncStatus`

Get the sync status for a specified sync

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetSyncByIDRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    }

    res, err := s.SyncStatus.GetSyncByID(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```