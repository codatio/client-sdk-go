# ListDisputes
Available in: `Disputes`

List commerce disputes

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/commerce"
	"github.com/codatio/client-sdk-go/commerce/pkg/models/operations"
)

func main() {
    s := codatcommerce.New(
        codatcommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListDisputesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatcommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommerce.Int(100),
        Query: codatcommerce.String("provident"),
    }

    res, err := s.Disputes.ListDisputes(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Disputes != nil {
        // handle response
    }
}
```