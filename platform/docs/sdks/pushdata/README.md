# PushData
(*PushData*)

## Overview

Initiate and monitor Create, Update, and Delete operations.

### Available Operations

* [GetModelOptions](#getmodeloptions) - Get push options
* [GetOperation](#getoperation) - Get push operation
* [ListOperations](#listoperations) - List push operations

## GetModelOptions

This is the generic documentation for creation and updating of data. See the equivalent endpoint for a given data type for more specific information. 

Before pushing data into accounting software, it is often necessary to collect some details from the user as to how they would like the data to be inserted. This includes names and amounts on transactional entities, but also factors such as categorisation of entities, which is often handled differently between different accounting software. A good example of this is specifying where on the balance sheet/profit and loss reports the user would like a newly-created nominal account to appear.

Codat tries not to limit users to pushing to a very limited number of standard categories, so we have implemented "options" endpoints, which allow us to expose to our clients the fields which are required to be pushed for a specific linked company, and the options which may be selected for each field.


> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/) for integrations that support push (POST/PUT methods).

### Example Usage

```go
package main

import(
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PushData.GetModelOptions(ctx, operations.GetCreateUpdateModelOptionsByDataTypeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DataType: shared.SchemaDataTypeInvoices,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetCreateUpdateModelOptionsByDataTypeRequest](../../pkg/models/operations/getcreateupdatemodeloptionsbydatatyperequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetCreateUpdateModelOptionsByDataTypeResponse](../../pkg/models/operations/getcreateupdatemodeloptionsbydatatyperesponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetOperation

The **Get push operation** endpoint returns a specific [push operation](/using-the-api/push) identified by the `pushOperationKey` that was performed on the company.

Write operations are actions that send requests to Codat, enabling the creation, updating, deletion of records, or uploading attachments in the connected accounting software.

### Example Usage

```go
package main

import(
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PushData.GetOperation(ctx, operations.GetPushOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "59acd79e-29d3-4138-91d3-91d4641bf7ed",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPushOperationRequest](../../pkg/models/operations/getpushoperationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPushOperationResponse](../../pkg/models/operations/getpushoperationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## ListOperations

The **List push operations** endpoint returns a list of [push operations](/using-the-api/push) performed on the company.

Write operations are actions that send requests to Codat, enabling the creation, updating, deletion of records, or uploading attachments in the connected accounting software.

### Example Usage

```go
package main

import(
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PushData.ListOperations(ctx, operations.GetCompanyPushHistoryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: platform.String("-modifiedDate"),
        Page: platform.Int(1),
        PageSize: platform.Int(100),
        Query: platform.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCompanyPushHistoryRequest](../../pkg/models/operations/getcompanypushhistoryrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetCompanyPushHistoryResponse](../../pkg/models/operations/getcompanypushhistoryresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |