# IsAgedCreditorsReportAvailable
Available in: `Reports`

Indicates whether the aged creditor report is available for the company.

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
    req := operations.IsAgedCreditorsReportAvailableRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Reports.IsAgedCreditorsReportAvailable(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.IsAgedCreditorsReportAvailable200ApplicationJSONBoolean != nil {
        // handle response
    }
}
```