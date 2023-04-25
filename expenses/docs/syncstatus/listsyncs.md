# ListSyncs
Available in: `SyncStatus`

Gets a list of sync statuses

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
    req := operations.ListSyncsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.SyncStatus.ListSyncs(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatuses != nil {
        // handle response
    }
}
```