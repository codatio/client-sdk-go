# SyncFlowSettings

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
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetConfigTextSyncFlowResponse](../../models/operations/getconfigtextsyncflowresponse.md), error**


## GetVisibleAccounts

Return accounts which are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.GetVisibleAccounts(ctx, operations.GetVisibleAccountsRequest{
        ClientID: "7cc8796e-d151-4a05-9fc2-ddf7cc78ca1b",
        PlatformKey: "a928fc81-6742-4cb7-b920-5929396fea75",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVisibleAccountsRequest](../../models/operations/getvisibleaccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetVisibleAccountsResponse](../../models/operations/getvisibleaccountsresponse.md), error**


## UpdateConfigTextSyncFlow

Set preferences for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateConfigTextSyncFlow(ctx, map[string]shared.Localization{
        "iure": shared.Localization{
            Required: codatsynccommerce.Bool(false),
            Text: codatsynccommerce.String("saepe"),
        },
        "quidem": shared.Localization{
            Required: codatsynccommerce.Bool(false),
            Text: codatsynccommerce.String("architecto"),
        },
        "ipsa": shared.Localization{
            Required: codatsynccommerce.Bool(false),
            Text: codatsynccommerce.String("reiciendis"),
        },
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [map[string]shared.Localization](../../models//.md)      | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.UpdateConfigTextSyncFlowResponse](../../models/operations/updateconfigtextsyncflowresponse.md), error**


## UpdateVisibleAccountsSyncFlow

Update which accounts are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowSettings.UpdateVisibleAccountsSyncFlow(ctx, operations.UpdateVisibleAccountsSyncFlowRequest{
        VisibleAccounts: &shared.VisibleAccounts{
            VisibleAccounts: []string{
                "mollitia",
                "laborum",
                "dolores",
            },
        },
        PlatformKey: "352c5955-907a-4ff1-a3a2-fa9467739251",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateVisibleAccountsSyncFlowRequest](../../models/operations/updatevisibleaccountssyncflowrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.UpdateVisibleAccountsSyncFlowResponse](../../models/operations/updatevisibleaccountssyncflowresponse.md), error**

