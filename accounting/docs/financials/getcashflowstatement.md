# GetCashFlowStatement
Available in: `Financials`

Gets the latest cash flow statement for a company.

## Example Usage
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
        PeriodLength: 656964,
        PeriodsToCompare: 1990,
        StartMonth: codataccounting.String("2022-10-23T00:00:00Z"),
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