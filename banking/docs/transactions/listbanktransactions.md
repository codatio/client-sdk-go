# ListBankTransactions
Available in: `Transactions`

Gets a list of transactions incurred by a company across all bank accounts.

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
    req := operations.ListBankTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("corrupti"),
    }

    res, err := s.Transactions.ListBankTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```