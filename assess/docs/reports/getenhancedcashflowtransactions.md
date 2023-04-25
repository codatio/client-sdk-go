# GetEnhancedCashFlowTransactions
Available in: `Reports`

The Enhanced Cash Flow Transactions endpoint provides a fully categorized list of banking transactions for a company. Accounts and transaction data are obtained from the company's banking data sources.

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
    req := operations.GetEnhancedCashFlowTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: 1,
        PageSize: codatassess.Int(100),
        Query: codatassess.String("rem"),
    }

    res, err := s.Reports.GetEnhancedCashFlowTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnhancedCashFlowTransactions != nil {
        // handle response
    }
}
```