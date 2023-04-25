# GetCommerceCustomerRetentionMetrics
Available in: `Reports`

Gets the customer retention metrics for a specific company connection, over one or more periods of time.

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
    req := operations.GetCommerceCustomerRetentionMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 359444,
        PeriodLength: 296140,
        PeriodUnit: shared.PeriodUnitEnumWeek,
        ReportDate: "29-09-2020",
    }

    res, err := s.Reports.GetCommerceCustomerRetentionMetrics(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```