# GetTransactionCategory
Available in: `TransactionCategories`

Gets a specified bank transaction category for a given company

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
    req := operations.GetTransactionCategoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransactionCategoryID: "quibusdam",
    }

    res, err := s.TransactionCategories.GetTransactionCategory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionCategory != nil {
        // handle response
    }
}
```