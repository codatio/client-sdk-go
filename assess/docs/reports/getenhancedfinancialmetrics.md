# GetEnhancedFinancialMetrics
Available in: `Reports`

Gets all the available financial metrics for a given company, over one or more periods.

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
    req := operations.GetEnhancedFinancialMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 916723,
        PeriodLength: 93940,
        ReportDate: "29-09-2020",
        ShowMetricInputs: codatassess.Bool(false),
    }

    res, err := s.Reports.GetEnhancedFinancialMetrics(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.FinancialMetrics != nil {
        // handle response
    }
}
```