# Reports

## Overview

Reports

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
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAgedCreditorsReport(ctx, operations.GetAgedCreditorsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: codataccounting.Int(12),
        PeriodLengthDays: codataccounting.Int(30),
        ReportDate: types.MustDateFromString("2022-12-31"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAgedCreditorsReportRequest](../../models/operations/getagedcreditorsreportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GetAgedCreditorsReportResponse](../../models/operations/getagedcreditorsreportresponse.md), error**


## GetAgedDebtorsReport

Returns aged debtors report for company that shows the total outstanding balance due from customers to the business over time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAgedDebtorsReport(ctx, operations.GetAgedDebtorsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: codataccounting.Int(12),
        PeriodLengthDays: codataccounting.Int(30),
        ReportDate: types.MustDateFromString("2022-12-31"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAgedDebtorsReportRequest](../../models/operations/getageddebtorsreportrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.GetAgedDebtorsReportResponse](../../models/operations/getageddebtorsreportresponse.md), error**


## GetBalanceSheet

Gets the latest balance sheet for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetBalanceSheet(ctx, operations.GetBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: codataccounting.String("2022-10-23T00:00:00.000Z"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetBalanceSheetRequest](../../models/operations/getbalancesheetrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetBalanceSheetResponse](../../models/operations/getbalancesheetresponse.md), error**


## GetCashFlowStatement

Gets the latest cash flow statement for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCashFlowStatement(ctx, operations.GetCashFlowStatementRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: codataccounting.String("2022-10-23T00:00:00.000Z"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetCashFlowStatementRequest](../../models/operations/getcashflowstatementrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.GetCashFlowStatementResponse](../../models/operations/getcashflowstatementresponse.md), error**


## GetProfitAndLoss

Gets the latest profit and loss for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetProfitAndLoss(ctx, operations.GetProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 4,
        PeriodsToCompare: 20,
        StartMonth: codataccounting.String("2022-10-23T00:00:00.000Z"),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetProfitAndLossRequest](../../models/operations/getprofitandlossrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetProfitAndLossResponse](../../models/operations/getprofitandlossresponse.md), error**


## IsAgedCreditorsReportAvailable

Indicates whether the aged creditor report is available for the company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
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

    if res.IsAgedCreditorsReportAvailable200ApplicationJSONBoolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.IsAgedCreditorsReportAvailableRequest](../../models/operations/isagedcreditorsreportavailablerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.IsAgedCreditorsReportAvailableResponse](../../models/operations/isagedcreditorsreportavailableresponse.md), error**


## IsAgedDebtorReportAvailable

Indicates whether the aged debtor report is available for the company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
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

    if res.IsAgedDebtorReportAvailable200ApplicationJSONBoolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.IsAgedDebtorReportAvailableRequest](../../models/operations/isageddebtorreportavailablerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.IsAgedDebtorReportAvailableResponse](../../models/operations/isageddebtorreportavailableresponse.md), error**

