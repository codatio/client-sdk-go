# SyncStatus

## Overview

Check the status of ongoing or previous expense syncs.

### Available Operations

* [GetLastSuccessfulSync](#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](#getlatestsync) - Latest sync status
* [GetSyncByID](#getsyncbyid) - Get Sync status
* [ListSyncs](#listsyncs) - List sync statuses

## GetLastSuccessfulSync

Gets the status of the last successfull sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetLastSuccessfulSyncRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.SyncStatus.GetLastSuccessfulSync(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

## GetLatestSync

Gets the latest sync status

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetLatestSyncRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.SyncStatus.GetLatestSync(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

## GetSyncByID

Get the sync status for a specified sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetSyncByIDRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    }

    res, err := s.SyncStatus.GetSyncByID(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatus != nil {
        // handle response
    }
}
```

## ListSyncs

Gets a list of sync statuses

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListSyncsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.SyncStatus.ListSyncs(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanySyncStatuses != nil {
        // handle response
    }
}
```
