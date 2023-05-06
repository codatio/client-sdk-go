# Settings

## Overview

Manage your Codat instance.

### Available Operations

* [GetProfile](#getprofile) - Get profile
* [GetSyncSettings](#getsyncsettings) - Update all sync settings
* [UpdateProfile](#updateprofile) - Update profile

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

## GetSyncSettings

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
    res, err := s.Settings.GetSyncSettings(ctx, operations.UpdateSyncSettingsRequestBody{
        ClientID: "367f7975-267b-439b-90c6-a6040ee680f3",
        OverridesDefaults: false,
        Settings: []shared.SyncSetting{
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("saepe"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 681820,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("in"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 359508,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("iste"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 437032,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("saepe"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 697631,
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
