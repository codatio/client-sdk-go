# PushData

## Overview

View push options and get push statuses.

### Available Operations

* [GetModelOptions](#getmodeloptions) - Get push options
* [GetOperation](#getoperation) - Get push operation
* [ListOperations](#listoperations) - List push operations

## GetModelOptions

This is the generic documentation for creation and updating of data. See the equivalent endpoint for a given data type for more specific information. 

Before pushing data into accounting software, it is often necessary to collect some details from the user as to how they would like the data to be inserted. This includes names and amounts on transactional entities, but also factors such as categorisation of entities, which is often handled differently between different accounting packages. A good example of this is specifying where on the balance sheet/profit and loss reports the user would like a newly-created nominal account to appear.

Codat tries not to limit users to pushing to a very limited number of standard categories, so we have implemented "options" endpoints, which allow us to expose to our clients the fields which are required to be pushed for a specific linked company, and the options which may be selected for each field.


> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/) for integrations that support push (POST/PUT methods).

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
    res, err := s.PushData.GetModelOptions(ctx, operations.GetCreateUpdateModelOptionsByDataTypeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DataType: shared.DataTypeInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetOperation

Retrieve push operation.

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
    res, err := s.PushData.GetOperation(ctx, operations.GetPushOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "74e0f467-cc87-496e-9151-a05dfc2ddf7c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperation != nil {
        // handle response
    }
}
```

## ListOperations

List push operation records.

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
    res, err := s.PushData.ListOperations(ctx, operations.GetCompanyPushHistoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: codatcommon.Int(1),
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("quod"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushHistoryResponse != nil {
        // handle response
    }
}
```
