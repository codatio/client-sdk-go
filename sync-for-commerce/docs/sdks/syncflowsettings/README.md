# SyncFlowSettings
(*SyncFlowSettings*)

## Overview

Configure preferences for any given Sync for Commerce company using sync flow.

### Available Operations

* [GetConfigTextSyncFlow](#getconfigtextsyncflow) - Get preferences for text fields
* [GetVisibleAccounts](#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](#updateconfigtextsyncflow) - Update preferences for text fields
* [UpdateVisibleAccountsSyncFlow](#updatevisibleaccountssyncflow) - Update visible accounts

## GetConfigTextSyncFlow

Return preferences set for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.GetConfigTextSyncFlowResponse](../../pkg/models/operations/getconfigtextsyncflowresponse.md), error**
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401,402,403,429,500,503 | application/json        |
| sdkerrors.SDKError      | 400-600                 | */*                     |

## GetVisibleAccounts

Return accounts which are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.GetVisibleAccounts(ctx, operations.GetVisibleAccountsRequest{
        ClientID: "86fe9741-738d-4f2c-8e96-9c3f84156e91",
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetVisibleAccountsRequest](../../pkg/models/operations/getvisibleaccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.GetVisibleAccountsResponse](../../pkg/models/operations/getvisibleaccountsresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## UpdateConfigTextSyncFlow

Set preferences for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateConfigTextSyncFlow(ctx, &map[string]shared.Localization{
        "key": shared.Localization{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [map[string]shared.Localization](../../.md)                  | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.UpdateConfigTextSyncFlowResponse](../../pkg/models/operations/updateconfigtextsyncflowresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## UpdateVisibleAccountsSyncFlow

Update which accounts are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateVisibleAccountsSyncFlow(ctx, operations.UpdateVisibleAccountsSyncFlowRequest{
        VisibleAccounts: &shared.VisibleAccounts{
            VisibleAccounts: []string{
                "string",
            },
        },
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateVisibleAccountsSyncFlowRequest](../../pkg/models/operations/updatevisibleaccountssyncflowrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.UpdateVisibleAccountsSyncFlowResponse](../../pkg/models/operations/updatevisibleaccountssyncflowresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
