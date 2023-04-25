# GetAllBankAccount
Available in: `BankAccounts`

Gets the bank account for given account ID.

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
    req := operations.GetAllBankAccountRequest{
        AccountID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Query: codataccounting.String("deserunt"),
    }

    res, err := s.BankAccounts.GetAllBankAccount(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankStatementAccount != nil {
        // handle response
    }
}
```