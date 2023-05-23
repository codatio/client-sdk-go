# ExcelReports

## Overview

Downloadable reports

### Available Operations

* [GenerateExcelReport](#generateexcelreport) - Generate Excel report
* [GetAccountingMarketingMetrics](#getaccountingmarketingmetrics) - Get marketing metrics report
* [GetExcelReport](#getexcelreport) - Download Excel report
* [GetExcelReportGenerationStatus](#getexcelreportgenerationstatus) - Get Excel report status

## GenerateExcelReport

Generate an Excel report which can subsequently be downloaded.

### Example Usage

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
    res, err := s.ExcelReports.GenerateExcelReport(ctx, operations.GenerateExcelReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnhancedInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExcelStatus != nil {
        // handle response
    }
}
```

## GetAccountingMarketingMetrics

Get the marketing metrics from an accounting source for a given company.

Request an Excel report for download.

### Example Usage

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
    res, err := s.ExcelReports.GetAccountingMarketingMetrics(ctx, operations.GetAccountingMarketingMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 451159,
        PeriodLength: 739264,
        PeriodUnit: shared.PeriodUnitDay,
        ReportDate: "29-09-2020",
        ShowInputValues: codatassess.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

## GetExcelReport

Download the previously generated Excel report to a local drive.

### Example Usage

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
    res, err := s.ExcelReports.GetExcelReport(ctx, operations.GetExcelReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeAssess,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

## GetExcelReportGenerationStatus

Returns the status of the latest report requested.

### Example Usage

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
    res, err := s.ExcelReports.GetExcelReportGenerationStatus(ctx, operations.GetExcelReportGenerationStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnhancedFinancials,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExcelStatus != nil {
        // handle response
    }
}
```
