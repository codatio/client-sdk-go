# CodatLendingReports
(*Sales.Reports*)

### Available Operations

* [GetOrders](#getorders) - Get orders report
* [GetRefunds](#getrefunds) - Get refunds report

## GetOrders

The *Get orders report* endpoint returns the number of orders, total value, and average order value for a specific company's commerce connection over one or more periods of time.

This detail helps you assess a merchant's health and advise them on performance improvement strategies. It also provides you with key insights you need to assess the credit risk of a company. 

Learn more about the formulas used to calculate the order metrics [here](https://docs.codat.io/lending/commerce-metrics/overview#what-metrics-are-available).

Refer to the [commerce reporting structure](https://docs.codat.io/lending/commerce-metrics/reporting-structure) page for more details on commerce reports in Lending.

#### Response structure

The Orders report's dimensions and measures are:

| Index         | Dimensions     |
|---------------|----------------|
|   `index` = 0 | Period         |
|   `index` = 1 | Order metrics  |

| Index         | Measures   |
|---------------|------------|
| `index` = 0   | Count      |
| `index` = 1   | Value      |

The report data then combines multiple reporting dimensions and measures and outputs the value of each combination.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sales.Reports.GetOrders(ctx, operations.GetCommerceOrdersReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 491586,
        PeriodLength: 776309,
        PeriodUnit: shared.PeriodUnitWeek,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCommerceOrdersReportRequest](../../pkg/models/operations/getcommerceordersreportrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCommerceOrdersReportResponse](../../pkg/models/operations/getcommerceordersreportresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## GetRefunds

The *Get refunds report* endpoint returns the number and total value of refunds and the refund rate for a specific company's commerce connection over one or more periods of time.

This detail helps you assess a merchant's health and advise them on performance improvement strategies. It also provides you with key insights you need to assess the credit risk of a company. 

Learn more about the formulas used to calculate the refunds metrics [here](https://docs.codat.io/lending/commerce-metrics/overview#what-metrics-are-available).

Refer to the [commerce reporting structure](https://docs.codat.io/lending/commerce-metrics/reporting-structure) page for more details on commerce reports in Lending.

#### Response structure

The Refunds report's dimensions and measures are:

| Index          | Dimensions     |
|----------------|----------------|
| `index` = 0    | Period         |
| `index` = 1    | Refund metrics |

| Index          | Measures    |
|----------------|------------|
| `index` = 0    | Count      |
| `index` = 1    | Value      |
| `index` = 2    | Percentage |

The report data then combines multiple reporting dimensions and measures and outputs the value of each combination.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sales.Reports.GetRefunds(ctx, operations.GetCommerceRefundsReportRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 277786,
        PeriodLength: 876670,
        PeriodUnit: shared.PeriodUnitYear,
        ReportDate: "29-09-2020",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCommerceRefundsReportRequest](../../pkg/models/operations/getcommercerefundsreportrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.GetCommerceRefundsReportResponse](../../pkg/models/operations/getcommercerefundsreportresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
