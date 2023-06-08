# Companies

## Overview

Create and manage your Codat companies.

### Available Operations

* [Create](#create) - Create company
* [Delete](#delete) - Delete a company
* [Get](#get) - Get company
* [List](#list) - List companies
* [Update](#update) - Update company

## Create

Create a new company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Create(ctx, shared.CompanyRequestBody{
        Description: codatcommon.String("Requested early access to the new financing scheme."),
        Name: "Bank of Dave",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```

## Delete

Delete the given company from Codat.
This operation is not reversible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Delete(ctx, operations.DeleteCompanyRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Get

Get metadata for a single company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Get(ctx, operations.GetCompanyRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```

## List

List all companies that you have created in Codat.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: codatcommon.Int(1),
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Companies != nil {
        // handle response
    }
}
```

## Update

Updates the given company with a new name and description

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Update(ctx, operations.UpdateCompanyRequest{
        CompanyRequestBody: &shared.CompanyRequestBody{
            Description: codatcommon.String("Requested early access to the new financing scheme."),
            Name: "Bank of Dave",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Company != nil {
        // handle response
    }
}
```
