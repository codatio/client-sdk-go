# ListSyncTransactions
Available in: `TransactionStatus`

Get's the transactions and status for a sync

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
    req := operations.ListSyncTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: 1,
        PageSize: codatsyncexpenses.Int(100),
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    }

    res, err := s.TransactionStatus.ListSyncTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionMetadataList != nil {
        // handle response
    }
}
```