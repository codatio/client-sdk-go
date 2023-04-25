# SaveCompanyConfiguration
Available in: `Configuration`

Sets a companies expense sync configuration

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.SaveCompanyConfigurationRequest{
        CompanyConfiguration: &shared.CompanyConfiguration{
            BankAccount: shared.BankAccount{
                ID: codatsyncexpenses.String("32"),
            },
            Customer: shared.Customer{
                ID: codatsyncexpenses.String("142"),
            },
            Supplier: shared.Supplier{
                ID: codatsyncexpenses.String("124"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Configuration.SaveCompanyConfiguration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyConfiguration != nil {
        // handle response
    }
}
```