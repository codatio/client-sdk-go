# ListBankAccountTransactions
Available in: `BankAccountTransactions`

Gets bank transactions for a given bank account ID

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bankfeeds"
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListBankAccountTransactionsRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbankfeeds.String("-modifiedDate"),
        Page: 1,
        PageSize: codatbankfeeds.Int(100),
        Query: codatbankfeeds.String("aliquid"),
    }

    res, err := s.BankAccountTransactions.ListBankAccountTransactions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankTransactionsResponse != nil {
        // handle response
    }
}
```