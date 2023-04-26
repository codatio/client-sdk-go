# Settings

## Overview

Manage your Codat instance.

### Available Operations

* [GetProfile](#getprofile) - Get profile
* [GetProfileSyncSettings](#getprofilesyncsettings) - Get sync settings
* [UpdateProfile](#updateprofile) - Update profile
* [UpdateSyncSettings](#updatesyncsettings) - Update all sync settings

## GetProfile

Fetch your Codat profile.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
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

## GetProfileSyncSettings

Retrieve the sync settings for your client. This includes how often data types should be queued to be updated, and how much history should be fetched.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Settings.GetProfileSyncSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncSettings != nil {
        // handle response
    }
}
```

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.Profile{
        AlertAuthHeader: codatcommon.String("explicabo"),
        APIKey: codatcommon.String("nobis"),
        ConfirmCompanyName: codatcommon.Bool(false),
        IconURL: codatcommon.String("enim"),
        LogoURL: codatcommon.String("omnis"),
        Name: "Bob's Burgers",
        RedirectURL: "nemo",
        WhiteListUrls: []string{
            "https://bobs-burgers.com/redirect",
            "https://bobs-burgers.com/redirect",
        },
    }

    res, err := s.Settings.UpdateProfile(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Profile != nil {
        // handle response
    }
}
```

## UpdateSyncSettings

Update sync settings for all data types.

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
    req := operations.UpdateSyncSettingsRequestBody{
        ClientID: "367f7975-267b-439b-90c6-a6040ee680f3",
        OverridesDefaults: false,
        Settings: []shared.SyncSetting{
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("accusantium"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 438601,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("culpa"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 988374,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("sapiente"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 102044,
                SyncSchedule: 24,
            },
        },
    }

    res, err := s.Settings.UpdateSyncSettings(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
