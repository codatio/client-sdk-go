# Sync

## Overview

Initiate and monitor the sync of company data into accounting software.

### Available Operations

* [Get](#get) - Get sync status
* [GetLastSuccessfulSync](#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](#getlatestsync) - Latest sync status
* [GetStatus](#getstatus) - Get sync status
* [List](#list) - List sync statuses
* [Request](#request) - Initiate new sync
* [RequestForDateRange](#requestfordaterange) - Initiate sync for specific range

## Get

Get the sync status for a specified sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.Get(ctx, operations.GetSyncByIDRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSyncByIDRequest](../../models/operations/getsyncbyidrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.GetSyncByIDResponse](../../models/operations/getsyncbyidresponse.md), error**


## GetLastSuccessfulSync

Gets the status of the last successful sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.GetLastSuccessfulSync(ctx, operations.GetLastSuccessfulSyncRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetLastSuccessfulSyncRequest](../../models/operations/getlastsuccessfulsyncrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetLastSuccessfulSyncResponse](../../models/operations/getlastsuccessfulsyncresponse.md), error**


## GetLatestSync

Gets the latest sync status

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.GetLatestSync(ctx, operations.GetLatestSyncRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLatestSyncRequest](../../models/operations/getlatestsyncrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.GetLatestSyncResponse](../../models/operations/getlatestsyncresponse.md), error**


## GetStatus

Gets a list of sync statuses.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.GetStatus(ctx, operations.GetSyncStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetSyncStatusRequest](../../models/operations/getsyncstatusrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.GetSyncStatusResponse](../../models/operations/getsyncstatusresponse.md), error**


## List

Gets a list of sync statuses

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.List(ctx, operations.ListSyncsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListSyncsRequest](../../models/operations/listsyncsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |


### Response

**[*operations.ListSyncsResponse](../../models/operations/listsyncsresponse.md), error**


## Request

Run a Commerce sync from the last successful sync up to the date provided (optional), otherwise UtcNow is used.\r\nIf there was no previously successful sync, the start date in the config is used.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.Request(ctx, operations.RequestSyncRequest{
        SyncToLatestArgs: &shared.SyncToLatestArgs{
            SyncTo: syncforcommerce.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RequestSyncRequest](../../models/operations/requestsyncrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.RequestSyncResponse](../../models/operations/requestsyncresponse.md), error**


## RequestForDateRange

Initiate a sync for the specified start date to the specified finish date in the request payload.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.RequestForDateRange(ctx, operations.RequestSyncForDateRangeRequest{
        SyncRange: &shared.SyncRange{
            DateRange: shared.SyncRangeDateRange{
                Finish: "2022-10-23T00:00:00.000Z",
                Start: "2022-10-23T00:00:00.000Z",
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RequestSyncForDateRangeRequest](../../models/operations/requestsyncfordaterangerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.RequestSyncForDateRangeResponse](../../models/operations/requestsyncfordaterangeresponse.md), error**

