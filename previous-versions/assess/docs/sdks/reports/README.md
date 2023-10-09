# Reports
(*Reports*)

## Overview

Enriched reports and analyses of financial data

### Available Operations

* [GenerateLoanSummary](#generateloansummary) - Generate loan summaries report
* [GenerateLoanTransactions](#generateloantransactions) - Generate loan transactions report
* [GetAccountsForEnhancedBalanceSheet](#getaccountsforenhancedbalancesheet) - Get enhanced balance sheet accounts
* [GetAccountsForEnhancedProfitAndLoss](#getaccountsforenhancedprofitandloss) - Get enhanced profit and loss accounts
* [GetCommerceCustomerRetentionMetrics](#getcommercecustomerretentionmetrics) - Get customer retention metrics
* [GetCommerceLifetimeValueMetrics](#getcommercelifetimevaluemetrics) - Get lifetime value metric
* [GetCommerceOrdersMetrics](#getcommerceordersmetrics) - Get orders report
* [GetCommerceRefundsMetrics](#getcommercerefundsmetrics) - Get refunds report
* [GetCommerceRevenueMetrics](#getcommercerevenuemetrics) - Get commerce revenue metrics
* [GetEnhancedCashFlowTransactions](#getenhancedcashflowtransactions) - Get enhanced cash flow report
* [GetEnhancedInvoicesReport](#getenhancedinvoicesreport) - Get enhanced invoices report
* [GetLoanSummary](#getloansummary) - Get loan summaries
* [GetRecurringRevenueMetrics](#getrecurringrevenuemetrics) - Get key subscription revenue metrics
* [ListLoanTransactions](#listloantransactions) - List loan transactions
* [RequestRecurringRevenueMetrics](#requestrecurringrevenuemetrics) - Generate key subscription revenue metrics

## GenerateLoanSummary

The _Generate loan summaries_ endpoint requests the generation of the Loan Summaries report.

Learn more about Codat's liabilities feature [here](https://docs.codat.io/lending/features/liabilities-overview).

Make sure you have [synced a company](https://docs.codat.io/codat-api#/operations/refresh-company-data) recently before calling the endpoint.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GenerateLoanSummary(ctx, operations.GenerateLoanSummaryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SourceType: operations.GenerateLoanSummarySourceTypeAccounting,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GenerateLoanSummaryRequest](../../models/operations/generateloansummaryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.GenerateLoanSummaryResponse](../../models/operations/generateloansummaryresponse.md), error**


## GenerateLoanTransactions

The _Generate loan transactions_ endpoint requests the generation of the Loan Transactions report.

Learn more about Codat's liabilities feature [here](https://docs.codat.io/lending/features/liabilities-overview).

Make sure you have [synced a company](https://docs.codat.io/codat-api#/operations/refresh-company-data) recently before calling the endpoint.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GenerateLoanTransactions(ctx, operations.GenerateLoanTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SourceType: operations.GenerateLoanTransactionsSourceTypeAccounting,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GenerateLoanTransactionsRequest](../../models/operations/generateloantransactionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GenerateLoanTransactionsResponse](../../models/operations/generateloantransactionsresponse.md), error**


## GetAccountsForEnhancedBalanceSheet

The Enhanced Balance Sheet Accounts endpoint returns a list of categorized accounts that appear on a company’s Balance Sheet along with a balance per financial statement date.

Codat suggests a category for each account automatically, but you can [change it](/docs/assess-categorizing-accounts-ecommerce-lending) to a more suitable one.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAccountsForEnhancedBalanceSheet(ctx, operations.GetAccountsForEnhancedBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetAccountsForEnhancedProfitAndLoss(ctx, operations.GetAccountsForEnhancedProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceCustomerRetentionMetrics(ctx, operations.GetCommerceCustomerRetentionMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 474636,
        PeriodLength: 781048,
        PeriodUnit: shared.PeriodUnitDay,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceLifetimeValueMetrics(ctx, operations.GetCommerceLifetimeValueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 463554,
        PeriodLength: 892968,
        PeriodUnit: shared.PeriodUnitDay,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceOrdersMetrics(ctx, operations.GetCommerceOrdersMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 661381,
        PeriodLength: 875123,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceRefundsMetrics(ctx, operations.GetCommerceRefundsMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 806705,
        PeriodLength: 498153,
        PeriodUnit: shared.PeriodUnitDay,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetCommerceRevenueMetrics(ctx, operations.GetCommerceRevenueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 58448,
        PeriodLength: 864392,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetEnhancedCashFlowTransactions(ctx, operations.GetEnhancedCashFlowTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: assess.Int(1),
        PageSize: assess.Int(100),
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetEnhancedInvoicesReport(ctx, operations.GetEnhancedInvoicesReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: assess.Int(1),
        PageSize: assess.Int(100),
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

The *Get loan summaries* endpoint returns a summary by integration type of all loans identified from a company's accounting, banking, and commerce integrations.

The endpoint returns a list of a company's [loan summaries](https://docs.codat.io/codat-api#/schemas/LoanSummary) for each valid data connection.

Make sure you have [synced a company](https://docs.codat.io/codat-api#/operations/refresh-company-data) recently before calling the endpoint.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.GetLoanSummary(ctx, operations.GetLoanSummaryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SourceType: operations.GetLoanSummarySourceTypeBanking,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
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

The *List loan transactions* endpoint returns all [loan transactions](https://docs.codat.io/codat-api#/schemas/LoanTransactions) identified from a company's accounting, banking, and commerce integrations.

This detail gives analysts a better idea of the loan obligations a company may have.

Make sure you have [synced a company](https://docs.codat.io/codat-api#/operations/refresh-company-data) recently before calling the endpoint.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Reports.ListLoanTransactions(ctx, operations.ListLoanTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SourceType: operations.ListLoanTransactionsSourceTypeCommerce,
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
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
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

