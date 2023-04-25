# DeleteBillPayment
Available in: `BillPayments`

Deletes a bill payment from the accounting package for a given company.

> **Supported Integrations**
> 
> This functionality is currently only supported for our Oracle NetSuite integration. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    req := operations.DeleteBillPaymentRequest{
        BillPaymentID: "autem",
        CompanyID: "quo",
        ConnectionID: "nesciunt",
    }

    res, err := s.BillPayments.DeleteBillPayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```