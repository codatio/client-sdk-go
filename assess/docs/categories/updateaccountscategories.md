# UpdateAccountsCategories
Available in: `Categories`

Comfirms the categories for all or a batch of accounts for a specific connection.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateAccountsCategoriesRequest{
        ConfirmCategories: &shared.ConfirmCategories{
            Categories: []shared.ConfirmCategoriesCategories{
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "69a674e0-f467-4cc8-b96e-d151a05dfc2d",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("at"),
                        Subtype: codatassess.String("maiores"),
                        Type: codatassess.String("molestiae"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "cc78ca1b-a928-4fc8-9674-2cb739205929",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("dolor"),
                        Subtype: codatassess.String("natus"),
                        Type: codatassess.String("laboriosam"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "fea7596e-b10f-4aaa-a352-c5955907aff1",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("mollitia"),
                        Subtype: codatassess.String("dolorem"),
                        Type: codatassess.String("culpa"),
                    },
                },
                shared.ConfirmCategoriesCategories{
                    AccountRef: &shared.ConfirmCategoriesCategoriesAccountRef{
                        ID: "2fa94677-3925-41aa-92c3-f5ad019da1ff",
                    },
                    Confirmed: &shared.AccountCategory{
                        DetailType: codatassess.String("vero"),
                        Subtype: codatassess.String("nihil"),
                        Type: codatassess.String("praesentium"),
                    },
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Categories.UpdateAccountsCategories(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccounts != nil {
        // handle response
    }
}
```