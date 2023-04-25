# DownloadExcelReport
Available in: `ExcelReports`

Download the previously generated Excel report to a local drive.

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
    req := operations.DownloadExcelReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnumAudit,
    }

    res, err := s.ExcelReports.DownloadExcelReport(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```