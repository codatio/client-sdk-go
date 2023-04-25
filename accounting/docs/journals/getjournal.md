# GetJournal
Available in: `Journals`

Gets a single journal corresponding to the given ID.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetJournalRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalID: "mollitia",
    }

    res, err := s.Journals.GetJournal(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Journal != nil {
        // handle response
    }
}
```