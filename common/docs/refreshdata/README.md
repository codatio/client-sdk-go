# RefreshData

## Overview

Queue pull operations to refresh data in Codat.

### Available Operations

* [All](#all) - Queue pull operations
* [ByDataType](#bydatatype) - Queue pull operation

## All

Refreshes all data types marked Fetch on first link.

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
    req := operations.RefreshCompanyDataRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.RefreshData.All(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ByDataType

Queue a single pull operation for the given company and data type.

This will bring updated data into Codat from the linked integration for you to view.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreatePullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: codatcommon.String("b7392059-2939-46fe-a759-6eb10faaa235"),
        DataType: shared.DataTypeEnumInvoices,
    }

    res, err := s.RefreshData.ByDataType(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```
