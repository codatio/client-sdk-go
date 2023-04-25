# GenerateExcelReport
Available in: `ExcelReports`

Generate an Excel report which can subsequently be downloaded.

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
    req := operations.GenerateExcelReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnumAssess,
    }

    res, err := s.ExcelReports.GenerateExcelReport(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ExcelStatus != nil {
        // handle response
    }
}
```