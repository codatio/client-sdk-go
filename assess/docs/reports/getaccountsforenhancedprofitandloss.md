# GetAccountsForEnhancedProfitAndLoss
Available in: `Reports`

The Enhanced Profit and Loss Accounts endpoint returns a list of categorized accounts that appear on a company’s Profit and Loss. It also includes a balance per the financial statement date.

Codat suggests a category for each account automatically, but you can [change it](/docs/assess-categorizing-accounts-ecommerce-lending) to a more suitable one.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetAccountsForEnhancedProfitAndLossRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        NumberOfPeriods: 120196,
        ReportDate: "29-09-2020",
    }

    res, err := s.Reports.GetAccountsForEnhancedProfitAndLoss(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedReport != nil {
        // handle response
    }
}
```