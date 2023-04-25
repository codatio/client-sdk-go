# ListBankTransactions
Available in: `BankAccountTransactions`

Gets the latest bank transactions for given account ID and company. Doesn't require connection ID.

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
    req := operations.ListBankTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("vero"),
    }

    res, err := s.BankAccountTransactions.ListBankTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountTransactions != nil {
        // handle response
    }
}
```