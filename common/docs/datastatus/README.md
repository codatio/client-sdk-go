# DataStatus

## Overview

Understand the state of data within Codat.

### Available Operations

* [GetCompanyDataHistory](#getcompanydatahistory) - Get pull operations
* [GetCompanyDataStatus](#getcompanydatastatus) - Get data status
* [GetPullOperation](#getpulloperation) - Get pull operation

## GetCompanyDataHistory

Gets the pull operation history (datasets) for a given company.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCompanyDataHistoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("quis"),
    }

    res, err := s.DataStatus.GetCompanyDataHistory(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataConnectionHistory != nil {
        // handle response
    }
}
```

## GetCompanyDataStatus

Get the state of each data type for a company

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCompanyDataStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DataStatus.GetCompanyDataStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataStatusResponse != nil {
        // handle response
    }
}
```

## GetPullOperation

Retrieve information about a single dataset or pull operation.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetPullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DatasetID: "eaed9f0f-e77b-4bc9-a58f-ab8b4b99ab18",
    }

    res, err := s.DataStatus.GetPullOperation(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```
