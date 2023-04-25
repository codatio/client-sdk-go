# GetAgedDebtorsReport
Available in: `Reports`

Returns aged debtors report for company that shows the total outstanding balance due from customers to the business over time.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetAgedDebtorsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: codataccounting.Int(12),
        PeriodLengthDays: codataccounting.Int(30),
        ReportDate: types.MustDateFromString("2022-12-31"),
    }

    res, err := s.Reports.GetAgedDebtorsReport(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AgedDebtorReport != nil {
        // handle response
    }
}
```