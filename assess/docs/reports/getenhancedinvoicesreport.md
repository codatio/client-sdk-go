# GetEnhancedInvoicesReport
Available in: `Reports`

Gets a list of invoices linked to the corresponding banking transaction

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetEnhancedInvoicesReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: 1,
        PageSize: codatassess.Int(100),
        Query: codatassess.String("repudiandae"),
    }

    res, err := s.Reports.GetEnhancedInvoicesReport(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedInvoicesReport != nil {
        // handle response
    }
}
```