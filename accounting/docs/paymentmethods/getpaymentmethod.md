# GetPaymentMethod
Available in: `PaymentMethods`

Gets the specified payment method for a given company.

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
    req := operations.GetPaymentMethodRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentMethodID: "facere",
    }

    res, err := s.PaymentMethods.GetPaymentMethod(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```