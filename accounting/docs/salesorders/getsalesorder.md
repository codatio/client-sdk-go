# GetSalesOrder
Available in: `SalesOrders`

Get sales order

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
    req := operations.GetSalesOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SalesOrderID: "voluptatibus",
    }

    res, err := s.SalesOrders.GetSalesOrder(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SalesOrder != nil {
        // handle response
    }
}
```