# Connections
(*Connections*)

## Overview

Create new and manage existing data connections for a company.

### Available Operations

* [Create](#create) - Create connection
* [GetSyncFlowURL](#getsyncflowurl) - Start new sync flow
* [List](#list) - List connections
* [UpdateAuthorization](#updateauthorization) - Update authorization
* [UpdateConnection](#updateconnection) - Update connection

## Create

﻿Creates a connection for the company by providing a valid `platformKey`. 

Use the [List Integrations](https://docs.codat.io/platform-api#/operations/list-integrations) endpoint to access valid platform keys. 

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Connections.Create(ctx, operations.CreateConnectionRequest{
        RequestBody: &operations.CreateConnectionRequestBody{
            PlatformKey: syncforcommerce.String("gbol"),
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

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401, 402, 403, 404, 429 | application/json        |
| sdkerrors.ErrorMessage  | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetSyncFlowURL

Create a new company and connections. Get a URL for Sync Flow, including a one time passcode.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Connections.GetSyncFlowURL(ctx, operations.GetSyncFlowURLRequest{
        AccountingKey: "<value>",
        CommerceKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SyncFlowURL != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetSyncFlowURLRequest](../../pkg/models/operations/getsyncflowurlrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetSyncFlowURLResponse](../../pkg/models/operations/getsyncflowurlresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorMessage       | 400, 401, 402, 403, 404, 429 | application/json             |
| sdkerrors.ErrorMessage       | 500, 503                     | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## List

﻿List the connections for a company.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Connections.List(ctx, operations.ListConnectionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforcommerce.String("-modifiedDate"),
        Query: syncforcommerce.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
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

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.ErrorMessage       | 400, 401, 402, 403, 404, 429 | application/json             |
| sdkerrors.ErrorMessage       | 500, 503                     | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## UpdateAuthorization

Update data connection's authorization.

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Connections.UpdateAuthorization(ctx, operations.UpdateConnectionAuthorizationRequest{
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

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401, 402, 403, 404, 429 | application/json        |
| sdkerrors.ErrorMessage  | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateConnection

Update a data connection

### Example Usage

```go
package main

import(
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.Connections.UpdateConnection(ctx, operations.UpdateConnectionRequest{
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
| `request`                                                                                    | [operations.UpdateConnectionRequest](../../pkg/models/operations/updateconnectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateConnectionResponse](../../pkg/models/operations/updateconnectionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401, 402, 403, 404, 429 | application/json        |
| sdkerrors.ErrorMessage  | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |