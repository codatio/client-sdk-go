# Sync

## Overview

Initiate a sync of Sync for Commerce company data into their respective accounting software.

### Available Operations

* [RequestSync](#requestsync) - Run a Commerce sync from the last successful sync
* [RequestSyncForDateRange](#requestsyncfordaterange) - Run a Commerce sync from a given date range

## RequestSync

Run a Commerce sync from the last successful sync up to the date provided (optional), otherwise UtcNow is used.
If there was no previously successful sync, the start date in the config is used.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.RequestSync(ctx, operations.RequestSyncRequest{
        SyncToLatestArgs: &shared.SyncToLatestArgs{
            SyncTo: codatsynccommerce.String("nulla"),
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

## RequestSyncForDateRange

Run a Commerce sync from the specified start date to the specified finish date in the request payload.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.RequestSyncForDateRange(ctx, operations.RequestSyncForDateRangeRequest{
        DateRange: &shared.DateRange{
            Finish: codatsynccommerce.String("corrupti"),
            Start: codatsynccommerce.String("illum"),
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
