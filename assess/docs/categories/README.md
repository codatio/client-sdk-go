# Categories

## Overview

Categorisation

### Available Operations

* [GetAccountCategory](#getaccountcategory) - Get suggested and/or confirmed category for a specific account
* [ListAccountsCategories](#listaccountscategories) - List suggested and confirmed account categories
* [ListAvailableAccountCategories](#listavailableaccountcategories) - List account categories
* [UpdateAccountCategory](#updateaccountcategory) - Patch account categories
* [UpdateAccountsCategories](#updateaccountscategories) - Confirm categories for accounts

## GetAccountCategory

Get category for specific nominal account.

### Example Usage

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
    req := operations.GetAccountCategoryRequest{
        AccountID: "provident",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Categories.GetAccountCategory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```

## ListAccountsCategories

Lists suggested and confirmed chart of account categories for the given company and data connection.

### Example Usage

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

## ListAvailableAccountCategories

Lists available account categories Codat's categorisation engine can provide. 

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.ListAvailableAccountCategories(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Categories != nil {
        // handle response
    }
}
```

## UpdateAccountCategory

Update category for a specific nominal account

### Example Usage

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
    req := operations.UpdateAccountCategoryRequest{
        ConfirmCategory: &shared.ConfirmCategory{
            Confirmed: shared.AccountCategory{
                DetailType: codatassess.String("quibusdam"),
                Subtype: codatassess.String("unde"),
                Type: codatassess.String("nulla"),
            },
        },
        AccountID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Categories.UpdateAccountCategory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```

## UpdateAccountsCategories

Comfirms the categories for all or a batch of accounts for a specific connection.

### Example Usage

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
