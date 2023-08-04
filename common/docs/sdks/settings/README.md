# Settings

## Overview

Manage your Codat instance.

### Available Operations

* [~~GetProfile~~](#getprofile) - Get profile :warning: **Deprecated**
* [GetSyncSettings](#getsyncsettings) - Get sync settings
* [UpdateProfile](#updateprofile) - Update profile
* [UpdateSyncSettings](#updatesyncsettings) - Update all sync settings

## ~~GetProfile~~

Fetch your Codat profile.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.GetProfile(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Profile != nil {
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

**[*operations.GetProfileResponse](../../models/operations/getprofileresponse.md), error**


## GetSyncSettings

Retrieve the sync settings for your client. This includes how often data types should be queued to be updated, and how much history should be fetched.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.GetSyncSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSettings != nil {
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

**[*operations.GetProfileSyncSettingsResponse](../../models/operations/getprofilesyncsettingsresponse.md), error**


## UpdateProfile

Update your Codat profile

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.UpdateProfile(ctx, shared.Profile{
        AlertAuthHeader: codatcommon.String("Bearer tXEiHiRK7XCtI8TNHbpGs1LI1pumdb4Cl1QIo7B2"),
        APIKey: codatcommon.String("sartANTjHAkLdbyDfaynoTQb7pkmj6hXHmnQKMrB"),
        ConfirmCompanyName: codatcommon.Bool(false),
        IconURL: codatcommon.String("https://client-images.codat.io/icon/042399f5-d104-4f38-9ce8-cac3524f4e88_3f5623af-d992-4c22-bc08-e58c520a8526.ico"),
        LogoURL: codatcommon.String("https://client-images.codat.io/logo/042399f5-d104-4f38-9ce8-cac3524f4e88_5806cb1f-7342-4c0e-a0a8-99bfbc47b0ff.png"),
        Name: "Bob's Burgers",
        RedirectURL: "https://bobs-burgers.{countrySuffix}/{companyId}",
        WhiteListUrls: []string{
            "https://bobs-burgers.com",
            "https://bobs-burgers.com",
            "https://bobs-burgers.com",
            "https://bobs-burgers.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Profile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.Profile](../../models/shared/profile.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.UpdateProfileResponse](../../models/operations/updateprofileresponse.md), error**


## UpdateSyncSettings

Update sync settings for all data types.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.UpdateSyncSettings(ctx, operations.UpdateProfileSyncSettingsRequestBody{
        ClientID: "367f7975-267b-439b-90c6-a6040ee680f3",
        OverridesDefaults: false,
        Settings: []shared.SyncSetting{
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00.000Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 449950,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00.000Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 613064,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00.000Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 902349,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00.000Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 99280,
                SyncSchedule: 24,
            },
        },
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateProfileSyncSettingsRequestBody](../../models/operations/updateprofilesyncsettingsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.UpdateProfileSyncSettingsResponse](../../models/operations/updateprofilesyncsettingsresponse.md), error**

