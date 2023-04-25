# ListAccountsCategories
Available in: `Categories`

Lists suggested and confirmed chart of account categories for the given company and data connection.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListAccountsCategoriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatassess.String("-modifiedDate"),
        Page: 1,
        PageSize: codatassess.Int(100),
        Query: codatassess.String("distinctio"),
    }

    res, err := s.Categories.ListAccountsCategories(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccounts != nil {
        // handle response
    }
}
```