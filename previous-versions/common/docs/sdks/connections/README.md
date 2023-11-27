# Connections
(*Connections*)

## Overview

Manage your companies' data connections.

### Available Operations

* [Create](#create) - Create connection
* [Delete](#delete) - Delete connection
* [Get](#get) - Get connection
* [List](#list) - List connections
* [Unlink](#unlink) - Unlink connection
* [UpdateAuthorization](#updateauthorization) - Update authorization

## Create

﻿Creates a connection for the company by providing a valid `platformKey`. 

Use the [List Integrations](https://docs.codat.io/platform-api#/operations/list-integrations) endpoint to access valid platform keys. 

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.Create(ctx, operations.CreateConnectionRequest{
        RequestBody: &operations.CreateConnectionRequestBody{
            PlatformKey: common.String("gbol"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateConnectionRequest](../../pkg/models/operations/createconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.CreateConnectionResponse](../../pkg/models/operations/createconnectionresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## Delete

﻿Revoke and remove a connection from a company.
This operation is not reversible. The end user would need to reauthorize a new data connection if you wish to view new data for this company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.Delete(ctx, operations.DeleteConnectionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteConnectionRequest](../../pkg/models/operations/deleteconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.DeleteConnectionResponse](../../pkg/models/operations/deleteconnectionresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## Get

﻿Returns a specific connection for a company when valid identifiers are provided. If the identifiers are for a deleted company and/or connection, a not found response is returned.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.Get(ctx, operations.GetConnectionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetConnectionRequest](../../pkg/models/operations/getconnectionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetConnectionResponse](../../pkg/models/operations/getconnectionresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## List

﻿List the connections for a company.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.List(ctx, operations.ListConnectionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: common.String("-modifiedDate"),
        Page: common.Int(1),
        PageSize: common.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListConnectionsRequest](../../pkg/models/operations/listconnectionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ListConnectionsResponse](../../pkg/models/operations/listconnectionsresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## Unlink

﻿This allows you to deauthorize a connection, without deleting it from Codat. This means you can still view any data that has previously been pulled into Codat, and also lets you re-authorize in future if your customer wishes to resume sharing their data.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.Unlink(ctx, operations.UnlinkConnectionRequest{
        UpdateConnectionStatus: &shared.UpdateConnectionStatus{},
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UnlinkConnectionRequest](../../pkg/models/operations/unlinkconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.UnlinkConnectionResponse](../../pkg/models/operations/unlinkconnectionresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## UpdateAuthorization

Update data connection's authorization.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connections.UpdateAuthorization(ctx, operations.UpdateConnectionAuthorizationRequest{
        RequestBody: map[string]string{
            "key": "string",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateConnectionAuthorizationRequest](../../pkg/models/operations/updateconnectionauthorizationrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.UpdateConnectionAuthorizationResponse](../../pkg/models/operations/updateconnectionauthorizationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |
