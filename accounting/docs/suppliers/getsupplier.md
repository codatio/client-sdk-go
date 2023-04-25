# GetSupplier
Available in: `Suppliers`

Gets a single supplier corresponding to the given ID.

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
    req := operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Suppliers.GetSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Supplier != nil {
        // handle response
    }
}
```