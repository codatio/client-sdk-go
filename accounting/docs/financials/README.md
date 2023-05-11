# Financials

## Overview

Financials

### Available Operations

* [GetBalanceSheet](#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](#getprofitandloss) - Get profit and loss

## GetBalanceSheet

Gets the latest balance sheet for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Financials.GetBalanceSheet(ctx, operations.GetBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 181836,
        PeriodsToCompare: 754041,
        StartMonth: codataccounting.String("perspiciatis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BalanceSheet != nil {
        // handle response
    }
}
```

## GetCashFlowStatement

Gets the latest cash flow statement for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Financials.GetCashFlowStatement(ctx, operations.GetCashFlowStatementRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 431723,
        PeriodsToCompare: 576232,
        StartMonth: codataccounting.String("eligendi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CashFlowStatement != nil {
        // handle response
    }
}
```

## GetProfitAndLoss

Gets the latest profit and loss for a company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Financials.GetProfitAndLoss(ctx, operations.GetProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 276507,
        PeriodsToCompare: 790080,
        StartMonth: codataccounting.String("quod"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfitAndLossReport != nil {
        // handle response
    }
}
```
