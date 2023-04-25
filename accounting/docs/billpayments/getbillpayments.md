# GetBillPayments
Available in: `BillPayments`

Get a bill payment

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
    req := operations.GetBillPaymentsRequest{
        BillPaymentID: "illum",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.BillPayments.GetBillPayments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayment != nil {
        // handle response
    }
}
```