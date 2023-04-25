# GetBalanceSheet
Available in: `Financials`

Gets the latest balance sheet for a company.

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
    req := operations.GetBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 992382,
        PeriodsToCompare: 843875,
        StartMonth: codataccounting.String("2022-10-23T00:00:00Z"),
    }

    res, err := s.Financials.GetBalanceSheet(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BalanceSheet != nil {
        // handle response
    }
}
```