# PushJournal
Available in: `Journals`

Posts a new journal to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create journal model](https://docs.codat.io/accounting-api#/operations/get-create-journals-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support creating journals.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.PushJournalRequest{
        Journal: &shared.Journal{
            CreatedOn: codataccounting.String("2022-10-23T00:00:00Z"),
            HasChildren: codataccounting.Bool(false),
            ID: codataccounting.String("e62f6aa5-58a6-45e2-8830-16ca34bb87d4"),
            JournalCode: codataccounting.String("repellat"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Name: codataccounting.String("Lori Bergstrom"),
            ParentID: codataccounting.String("culpa"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.JournalStatusEnumActive.ToPointer(),
            Type: codataccounting.String("accusantium"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(461123),
    }

    res, err := s.Journals.PushJournal(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushJournal200ApplicationJSONObject != nil {
        // handle response
    }
}
```