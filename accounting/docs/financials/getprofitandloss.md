# GetProfitAndLoss
Available in: `Financials`

Gets the latest profit and loss for a company.

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
    req := operations.GetProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PeriodLength: 251070,
        PeriodsToCompare: 785469,
        StartMonth: codataccounting.String("2022-10-23T00:00:00Z"),
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