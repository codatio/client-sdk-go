# GetTaxRate
Available in: `TaxRates`

Gets the specified tax rate for a given company.

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
    req := operations.GetTaxRateRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TaxRateID: "est",
    }

    res, err := s.TaxRates.GetTaxRate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaxRate != nil {
        // handle response
    }
}
```