# GetEnhancedProfitAndLoss
Available in: `Reports`

Gets a fully categorized profit and loss statement for a given company, over one or more period(s).

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
    req := operations.GetEnhancedProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 575947,
        PeriodLength: 83112,
        ReportDate: "29-09-2020",
    }

    res, err := s.Reports.GetEnhancedProfitAndLoss(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```