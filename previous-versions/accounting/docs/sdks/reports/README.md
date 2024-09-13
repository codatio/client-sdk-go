# Reports
(*Reports*)

## Overview

Access standardized Reports from linked accounting software.

### Available Operations

* [GetAgedCreditorsReport](#getagedcreditorsreport) - Aged creditors report
* [GetAgedDebtorsReport](#getageddebtorsreport) - Aged debtors report
* [GetBalanceSheet](#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](#getprofitandloss) - Get profit and loss
* [IsAgedCreditorsReportAvailable](#isagedcreditorsreportavailable) - Aged creditors report available
* [IsAgedDebtorReportAvailable](#isageddebtorreportavailable) - Aged debtors report available

## GetAgedCreditorsReport

Returns aged creditors report for company that shows the total balance owed by a business to its suppliers over time.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAgedCreditorsReport(ctx, operations.GetAgedCreditorsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: accounting.Int(12),
        PeriodLengthDays: accounting.Int(30),
        ReportDate: types.MustNewDateFromString("2022-12-31"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgedCreditorReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetAgedCreditorsReportRequest](../../pkg/models/operations/getagedcreditorsreportrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetAgedCreditorsReportResponse](../../pkg/models/operations/getagedcreditorsreportresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## GetAgedDebtorsReport

Returns aged debtors report for company that shows the total outstanding balance due from customers to the business over time.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAgedDebtorsReport(ctx, operations.GetAgedDebtorsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: accounting.Int(12),
        PeriodLengthDays: accounting.Int(30),
        ReportDate: types.MustNewDateFromString("2022-12-31"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgedDebtorReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAgedDebtorsReportRequest](../../pkg/models/operations/getageddebtorsreportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAgedDebtorsReportResponse](../../pkg/models/operations/getageddebtorsreportresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## GetBalanceSheet

Gets the latest balance sheet for a company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetBalanceSheet(ctx, operations.GetBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: accounting.String("2022-10-23T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BalanceSheet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetBalanceSheetRequest](../../pkg/models/operations/getbalancesheetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetBalanceSheetResponse](../../pkg/models/operations/getbalancesheetresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## GetCashFlowStatement

Gets the latest cash flow statement for a company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCashFlowStatement(ctx, operations.GetCashFlowStatementRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: accounting.String("2022-10-23T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashFlowStatement != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCashFlowStatementRequest](../../pkg/models/operations/getcashflowstatementrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetCashFlowStatementResponse](../../pkg/models/operations/getcashflowstatementresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## GetProfitAndLoss

Gets the latest profit and loss for a company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetProfitAndLoss(ctx, operations.GetProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: accounting.String("2022-10-23T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfitAndLossReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetProfitAndLossRequest](../../pkg/models/operations/getprofitandlossrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetProfitAndLossResponse](../../pkg/models/operations/getprofitandlossresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## IsAgedCreditorsReportAvailable

Indicates whether the aged creditor report is available for the company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.IsAgedCreditorsReportAvailable(ctx, operations.IsAgedCreditorsReportAvailableRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Boolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.IsAgedCreditorsReportAvailableRequest](../../pkg/models/operations/isagedcreditorsreportavailablerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.IsAgedCreditorsReportAvailableResponse](../../pkg/models/operations/isagedcreditorsreportavailableresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## IsAgedDebtorReportAvailable

Indicates whether the aged debtor report is available for the company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"log"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.IsAgedDebtorReportAvailable(ctx, operations.IsAgedDebtorReportAvailableRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Boolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.IsAgedDebtorReportAvailableRequest](../../pkg/models/operations/isageddebtorreportavailablerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.IsAgedDebtorReportAvailableResponse](../../pkg/models/operations/isageddebtorreportavailableresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
