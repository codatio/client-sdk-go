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
    res, err := s.RefreshData.All(ctx, operations.RefreshCompanyDataRequest{
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
    res, err := s.RefreshData.ByDataType(ctx, operations.CreatePullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: codatcommon.String("8ca1ba92-8fc8-4167-82cb-739205929396"),
        DataType: shared.DataTypeEnumInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```
