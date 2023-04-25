# ListTransactionCategories
Available in: `TransactionCategories`

Gets a list of hierarchical categories associated with a transaction for greater contextual meaning to transactionactivity.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/banking"
	"github.com/codatio/client-sdk-go/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListTransactionCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("unde"),
    }

    res, err := s.TransactionCategories.ListTransactionCategories(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategories != nil {
        // handle response
    }
}
```