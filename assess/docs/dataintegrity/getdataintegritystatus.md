# GetDataIntegrityStatus
Available in: `DataIntegrity`

Gets match status for a given company and datatype.

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
    req := operations.GetDataIntegrityStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeEnumBankingAccounts,
    }

    res, err := s.DataIntegrity.GetDataIntegrityStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Status != nil {
        // handle response
    }
}
```