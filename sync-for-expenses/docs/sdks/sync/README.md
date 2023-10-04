# Sync
(*Sync*)

## Overview

Trigger and monitor expense syncs to accounting software.

### Available Operations

* [Get](#get) - Get sync status
* [GetLastSuccessfulSync](#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](#getlatestsync) - Latest sync status
* [InitiateSync](#initiatesync) - Initiate sync
* [List](#list) - List sync statuses

## Get

Get the sync status for a specified sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
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
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
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
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
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


## InitiateSync

Initiate sync of pending transactions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.InitiateSync(ctx, operations.InitiateSyncRequest{
        InitiateSync: &shared.InitiateSync{
            DatasetIds: []string{
                "acce2362-83d6-4e3e-a27f-f4a08e7217d5",
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncInitiated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.InitiateSyncRequest](../../models/operations/initiatesyncrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.InitiateSyncResponse](../../models/operations/initiatesyncresponse.md), error**


## List

Gets a list of sync statuses

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
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

