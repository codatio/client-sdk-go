# CompanyInfo

## Overview

Company info

### Available Operations

* [GetCompanyInfo](#getcompanyinfo) - Get company info
* [PostSyncInfo](#postsyncinfo) - Refresh company info

## GetCompanyInfo

Gets the latest basic info for a company.

### Example Usage

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
    req := operations.GetCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.CompanyInfo.GetCompanyInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyDataset != nil {
        // handle response
    }
}
```

## PostSyncInfo

Initiates the process of synchronising basic info for a company

### Example Usage

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
    req := operations.PostSyncInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.CompanyInfo.PostSyncInfo(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```
