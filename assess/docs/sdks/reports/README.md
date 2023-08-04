# Reports

## Overview

Data integrity is important

### Available Operations

* [GetAccountsForEnhancedBalanceSheet](#getaccountsforenhancedbalancesheet) - Get enhanced balance sheet accounts
* [GetAccountsForEnhancedProfitAndLoss](#getaccountsforenhancedprofitandloss) - Get enhanced profit and loss accounts
* [GetCommerceCustomerRetentionMetrics](#getcommercecustomerretentionmetrics) - Get customer retention metrics
* [GetCommerceLifetimeValueMetrics](#getcommercelifetimevaluemetrics) - Get lifetime value metric
* [GetCommerceOrdersMetrics](#getcommerceordersmetrics) - Get orders report
* [GetCommerceRefundsMetrics](#getcommercerefundsmetrics) - Get refunds report
* [GetCommerceRevenueMetrics](#getcommercerevenuemetrics) - Get commerce revenue metrics
* [GetEnhancedCashFlowTransactions](#getenhancedcashflowtransactions) - Get enhanced cash flow report
* [GetEnhancedInvoicesReport](#getenhancedinvoicesreport) - Get enhanced invoices report
* [GetLoanSummary](#getloansummary) - Get enhanced loan summaries
* [GetRecurringRevenueMetrics](#getrecurringrevenuemetrics) - Get key subscription revenue metrics
* [ListLoanTransactions](#listloantransactions) - List enhanced loan transactions
* [RequestRecurringRevenueMetrics](#requestrecurringrevenuemetrics) - Generate key subscription revenue metrics

## GetAccountsForEnhancedBalanceSheet

The Enhanced Balance Sheet Accounts endpoint returns a list of categorized accounts that appear on a company’s Balance Sheet along with a balance per financial statement date.

Codat suggests a category for each account automatically, but you can [change it](/docs/assess-categorizing-accounts-ecommerce-lending) to a more suitable one.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAccountsForEnhancedBalanceSheet(ctx, operations.GetAccountsForEnhancedBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: codatassess.Int64(623564),
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetAccountsForEnhancedBalanceSheetRequest](../../models/operations/getaccountsforenhancedbalancesheetrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.GetAccountsForEnhancedBalanceSheetResponse](../../models/operations/getaccountsforenhancedbalancesheetresponse.md), error**


## GetAccountsForEnhancedProfitAndLoss

The Enhanced Profit and Loss Accounts endpoint returns a list of categorized accounts that appear on a company’s Profit and Loss. It also includes a balance per the financial statement date.

Codat suggests a category for each account automatically, but you can [change it](/docs/assess-categorizing-accounts-ecommerce-lending) to a more suitable one.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAccountsForEnhancedProfitAndLoss(ctx, operations.GetAccountsForEnhancedProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: codatassess.Int64(645894),
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetAccountsForEnhancedProfitAndLossRequest](../../models/operations/getaccountsforenhancedprofitandlossrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |


### Response

**[*operations.GetAccountsForEnhancedProfitAndLossResponse](../../models/operations/getaccountsforenhancedprofitandlossresponse.md), error**


## GetCommerceCustomerRetentionMetrics

Gets the customer retention metrics for a specific company connection, over one or more periods of time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceCustomerRetentionMetrics(ctx, operations.GetCommerceCustomerRetentionMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 384382,
        PeriodLength: 437587,
        PeriodUnit: shared.PeriodUnitWeek,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetCommerceCustomerRetentionMetricsRequest](../../models/operations/getcommercecustomerretentionmetricsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |


### Response

**[*operations.GetCommerceCustomerRetentionMetricsResponse](../../models/operations/getcommercecustomerretentionmetricsresponse.md), error**


## GetCommerceLifetimeValueMetrics

Gets the lifetime value metric for a specific company connection, over one or more periods of time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceLifetimeValueMetrics(ctx, operations.GetCommerceLifetimeValueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 891773,
        PeriodLength: 56713,
        PeriodUnit: shared.PeriodUnitYear,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetCommerceLifetimeValueMetricsRequest](../../models/operations/getcommercelifetimevaluemetricsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.GetCommerceLifetimeValueMetricsResponse](../../models/operations/getcommercelifetimevaluemetricsresponse.md), error**


## GetCommerceOrdersMetrics

Gets the order information for a specific company connection, over one or more periods of time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceOrdersMetrics(ctx, operations.GetCommerceOrdersMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 272656,
        PeriodLength: 383441,
        PeriodUnit: shared.PeriodUnitWeek,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCommerceOrdersMetricsRequest](../../models/operations/getcommerceordersmetricsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GetCommerceOrdersMetricsResponse](../../models/operations/getcommerceordersmetricsresponse.md), error**


## GetCommerceRefundsMetrics

Gets the refunds information for a specific company connection, over one or more periods of time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceRefundsMetrics(ctx, operations.GetCommerceRefundsMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 791725,
        PeriodLength: 812169,
        PeriodUnit: shared.PeriodUnitMonth,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCommerceRefundsMetricsRequest](../../models/operations/getcommercerefundsmetricsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCommerceRefundsMetricsResponse](../../models/operations/getcommercerefundsmetricsresponse.md), error**


## GetCommerceRevenueMetrics

Get the revenue and revenue growth for a specific company connection, over one or more periods of time.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceRevenueMetrics(ctx, operations.GetCommerceRevenueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        IncludeDisplayNames: codatassess.Bool(false),
        NumberOfPeriods: 479977,
        PeriodLength: 568045,
        PeriodUnit: shared.PeriodUnitWeek,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCommerceRevenueMetricsRequest](../../models/operations/getcommercerevenuemetricsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCommerceRevenueMetricsResponse](../../models/operations/getcommercerevenuemetricsresponse.md), error**


## GetEnhancedCashFlowTransactions

> **Categorization engine**
> 
> The categorization engine uses machine learning and has been fully trained against Plaid and TrueLayer banking data sources. It is not fully trained against the Basiq banking data source.

The Enhanced Cash Flow Transactions endpoint provides a fully categorized list of banking transactions for a company. Accounts and transaction data are obtained from the company's banking data sources.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetEnhancedCashFlowTransactions(ctx, operations.GetEnhancedCashFlowTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: codatassess.Int(1),
        PageSize: codatassess.Int(100),
        Query: codatassess.String("recusandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedCashFlowTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetEnhancedCashFlowTransactionsRequest](../../models/operations/getenhancedcashflowtransactionsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.GetEnhancedCashFlowTransactionsResponse](../../models/operations/getenhancedcashflowtransactionsresponse.md), error**


## GetEnhancedInvoicesReport

Gets a list of invoices linked to the corresponding banking transaction

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetEnhancedInvoicesReport(ctx, operations.GetEnhancedInvoicesReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: codatassess.Int(1),
        PageSize: codatassess.Int(100),
        Query: codatassess.String("temporibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedInvoicesReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetEnhancedInvoicesReportRequest](../../models/operations/getenhancedinvoicesreportrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetEnhancedInvoicesReportResponse](../../models/operations/getenhancedinvoicesreportresponse.md), error**


## GetLoanSummary

Get enhanced loan summaries

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetLoanSummary(ctx, operations.GetLoanSummaryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoanSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLoanSummaryRequest](../../models/operations/getloansummaryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.GetLoanSummaryResponse](../../models/operations/getloansummaryresponse.md), error**


## GetRecurringRevenueMetrics

Gets key metrics for subscription revenue.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetRecurringRevenueMetrics(ctx, operations.GetRecurringRevenueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetRecurringRevenueMetricsRequest](../../models/operations/getrecurringrevenuemetricsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.GetRecurringRevenueMetricsResponse](../../models/operations/getrecurringrevenuemetricsresponse.md), error**


## ListLoanTransactions

List enhanced loan transactions

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.ListLoanTransactions(ctx, operations.ListLoanTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SourceType: operations.ListLoanTransactionsSourceTypeBanking,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoanTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLoanTransactionsRequest](../../models/operations/listloantransactionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.ListLoanTransactionsResponse](../../models/operations/listloantransactionsresponse.md), error**


## RequestRecurringRevenueMetrics

Requests production of key subscription revenue metrics.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.RequestRecurringRevenueMetrics(ctx, operations.RequestRecurringRevenueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.RequestRecurringRevenueMetricsRequest](../../models/operations/requestrecurringrevenuemetricsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.RequestRecurringRevenueMetricsResponse](../../models/operations/requestrecurringrevenuemetricsresponse.md), error**

