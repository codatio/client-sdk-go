# GetDataIntegritySummaries
Available in: `DataIntegrity`

Gets match summary for a given company and datatype, optionally restricted by a Codat query string.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetDataIntegritySummariesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeEnumBankingAccounts,
        Query: codatassess.String("ipsa"),
    }

    res, err := s.DataIntegrity.GetDataIntegritySummaries(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Summaries != nil {
        // handle response
    }
}
```