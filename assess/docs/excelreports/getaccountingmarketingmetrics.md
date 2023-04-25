# GetAccountingMarketingMetrics
Available in: `ExcelReports`

Request an Excel report for download.

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
    req := operations.GetAccountingMarketingMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 739264,
        PeriodLength: 19987,
        PeriodUnit: shared.PeriodUnitEnumDay,
        ReportDate: "29-09-2020",
        ShowInputValues: codatassess.Bool(false),
    }

    res, err := s.ExcelReports.GetAccountingMarketingMetrics(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```