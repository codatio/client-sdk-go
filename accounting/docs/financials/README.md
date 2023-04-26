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
    req := operations.GetBalanceSheetRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 690865,
        PeriodsToCompare: 125404,
        StartMonth: codataccounting.String("facere"),
    }

    res, err := s.Financials.GetBalanceSheet(ctx, req)
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
    req := operations.GetCashFlowStatementRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 170992,
        PeriodsToCompare: 446583,
        StartMonth: codataccounting.String("repudiandae"),
    }

    res, err := s.Financials.GetCashFlowStatement(ctx, req)
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
    req := operations.GetProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 698306,
        PeriodsToCompare: 457632,
        StartMonth: codataccounting.String("accusantium"),
    }

    res, err := s.Financials.GetProfitAndLoss(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfitAndLossReport != nil {
        // handle response
    }
}
```
