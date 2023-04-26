# ExcelReports

## Overview

Downloadable reports

### Available Operations

* [DownloadExcelReport](#downloadexcelreport) - Download generated excel report
* [GenerateExcelReport](#generateexcelreport) - Generate an Excel report
* [GetAccountingMarketingMetrics](#getaccountingmarketingmetrics) - Get the marketing metrics from an accounting source for a given company.
* [GetExcelReport](#getexcelreport) - Download generated excel report
* [GetExcelReportGenerationStatus](#getexcelreportgenerationstatus) - Get status of Excel report

## DownloadExcelReport

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

## GetAccountingMarketingMetrics

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
    req := operations.GetExcelReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnumAssess,
    }

    res, err := s.ExcelReports.GetExcelReport(ctx, req)
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
    req := operations.GetExcelReportGenerationStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ReportType: shared.ExcelReportTypeEnumAssess,
    }

    res, err := s.ExcelReports.GetExcelReportGenerationStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ExcelStatus != nil {
        // handle response
    }
}
```