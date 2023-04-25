# UpdateSyncSettings
Available in: `Settings`

Update sync settings for all data types.

## Example Usage
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
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 38425,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 438601,
                SyncSchedule: 24,
            },
            shared.SyncSetting{
                DataType: shared.SyncSettingDataTypeEnumInvoices,
                FetchOnFirstLink: false,
                IsLocked: codatcommon.Bool(false),
                MonthsToSync: codatcommon.Int64(24),
                SyncFromUtc: codatcommon.String("2022-10-23T00:00:00Z"),
                SyncFromWindow: codatcommon.Int64(24),
                SyncOrder: 634274,
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