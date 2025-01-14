# Metrics
(*Sales.Metrics*)

## Overview

### Available Operations

* [GetCustomerRetention](#getcustomerretention) - Get customer retention metrics
* [GetLifetimeValue](#getlifetimevalue) - Get lifetime value metrics
* [GetRevenue](#getrevenue) - Get commerce revenue metrics

## GetCustomerRetention

The *Get customer retention metrics* endpoint returns customer retention insights for a specific company's commerce connection over one or more periods of time.

This detail helps you assess a merchant's health and advise them on performance improvement strategies. It also provides you with key insights you need to assess the credit risk of a company. 

#### Customer retention metrics

- __Existing customers__: the number of unique customers that have placed an order(s) in the specified period and any previous period. 
- __New customers__: the number of unique customers that have placed an order(s) in the specified period and none in any previous period.
- __Total customers__: the total number of existing and new customers within the specified period.
- __Retention rate__: the ratio of existing customers within the specified period compared to the total customers at the end of the previous period represented as a percentage.
- __Repeat rate__: the ratio of existing customers to total customers over the specified period represented as a percentage.

[Learn more](https://docs.codat.io/lending/features/sales-overview#metrics) about the formulas used to calculate customer retention metrics.

#### Response structure

The Customer retention report's dimensions and measures are:

| Index                       | Dimensions                 |
|-----------------------------|----------------------------|
| `index` = 0                 | Period                     |
| `index` = 1                 | Customer retention metrics |

| Index                | Measures    |
|----------------------|------------|
| `index` = 0          | Count      |
| `index` = 1          | Percentage |

The report data then combines multiple reporting dimensions and measures and outputs the value of each combination.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Sales.Metrics.GetCustomerRetention(ctx, operations.GetCommerceCustomerRetentionMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 431272,
        PeriodLength: 497588,
        PeriodUnit: shared.PeriodUnitDay,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetCommerceCustomerRetentionMetricsRequest](../../pkg/models/operations/getcommercecustomerretentionmetricsrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.GetCommerceCustomerRetentionMetricsResponse](../../pkg/models/operations/getcommercecustomerretentionmetricsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetLifetimeValue

The *Get lifetime value metrics* endpoint returns the average revenue that a specific company will generate throughout its lifespan over one or more periods of time.

This detail helps you assess a merchant's health and advise them on performance improvement strategies. It also provides you with key insights you need to assess the credit risk of a company.

[Learn more](https://docs.codat.io/lending/features/sales-overview#metrics) about the formulas used to calculate the lifetime value metrics.

#### Response structure

The Lifetime value report's dimensions and measures are:

| Index         | Dimensions             |
|---------------|------------------------|
| `index` = 0   | Period                 |
| `index` = 1   | Lifetime value metrics |

| Index             | Measures |
|-------------------|---------|
|   `index` = 1     | Value   |

The report data then combines multiple reporting dimensions and measures and outputs the value of each combination.


### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Sales.Metrics.GetLifetimeValue(ctx, operations.GetCommerceLifetimeValueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 500610,
        PeriodLength: 900865,
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetCommerceLifetimeValueMetricsRequest](../../pkg/models/operations/getcommercelifetimevaluemetricsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.GetCommerceLifetimeValueMetricsResponse](../../pkg/models/operations/getcommercelifetimevaluemetricsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## GetRevenue

The *Get revenue report* endpoint returns the revenue and revenue growth for a specific company connection over one or more periods of time.

This detail helps you assess a merchant's health and advise them on performance improvement strategies. It also provides you with key insights you need to assess the credit risk of a company. 

[Learn more](https://docs.codat.io/lending/features/sales-overview#metrics) about the formulas used to calculate the revenue metrics.

#### Response structure

The Revenue report's dimensions and measures are:

| Index         | Dimensions |
|---------------|------------|
|   `index` = 0 | Period     |
|   `index` = 1 | Revenue    |

| Index         | Measures                                                                                                                 |
|---------------|--------------------------------------------------------------------------------------------------------------------------|
| `index` = 0   | Value                                                                                                                    |
| `index` = 1   | Percentage change, defined as the change between the current and previous periods' values and expressed as a percentage. |

The report data then combines multiple reporting dimensions and measures and outputs the value of each combination.


### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Sales.Metrics.GetRevenue(ctx, operations.GetCommerceRevenueMetricsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        NumberOfPeriods: 120092,
        PeriodLength: 307462,
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCommerceRevenueMetricsRequest](../../pkg/models/operations/getcommercerevenuemetricsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCommerceRevenueMetricsResponse](../../pkg/models/operations/getcommercerevenuemetricsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |