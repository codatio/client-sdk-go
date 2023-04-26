# SyncFlowPreferences

## Overview

Configure preferences for any given Sync for Commerce company using sync flow.

### Available Operations

* [GetConfigTextSyncFlow](#getconfigtextsyncflow) - Retrieve preferences for text fields on Sync Flow
* [GetSyncFlowURL](#getsyncflowurl) - Retrieve sync flow url
* [GetVisibleAccounts](#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](#updateconfigtextsyncflow) - Update preferences for text fields on sync flow
* [UpdateVisibleAccountsSyncFlow](#updatevisibleaccountssyncflow) - Update the visible accounts on Sync Flow

## GetConfigTextSyncFlow

To enable retrieval of preferences set for the text fields on Sync Flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

## GetSyncFlowURL

Get a URL for Sync Flow including a one time passcode.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetSyncFlowURLRequest{
        AccountingKey: "vel",
        CommerceKey: "error",
        MerchantIdentifier: codatsynccommerce.String("deserunt"),
    }

    res, err := s.SyncFlowPreferences.GetSyncFlowURL(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncFlowURL != nil {
        // handle response
    }
}
```

## GetVisibleAccounts

Enable retrieval for accounts which are visible on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetVisibleAccountsRequest{
        ClientID: "674e0f46-7cc8-4796-ad15-1a05dfc2ddf7",
        PlatformKey: "cc78ca1b-a928-4fc8-9674-2cb739205929",
    }

    res, err := s.SyncFlowPreferences.GetVisibleAccounts(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```

## UpdateConfigTextSyncFlow

To enable update of preferences set for the text fields on sync flow.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/synccommerce"
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := map[string]shared.Localization{
        "natus": shared.Localization{
            Required: codatsynccommerce.Bool(false),
            Text: codatsynccommerce.String("laboriosam"),
        },
    }

    res, err := s.SyncFlowPreferences.UpdateConfigTextSyncFlow(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.LocalizationInfo != nil {
        // handle response
    }
}
```

## UpdateVisibleAccountsSyncFlow

To enable update of accounts visible preferences set on Sync Flow.

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
    req := operations.UpdateVisibleAccountsSyncFlowRequest{
        VisibleAccounts: &shared.VisibleAccounts{
            VisibleAccounts: []string{
                "saepe",
                "fuga",
                "in",
                "corporis",
            },
        },
        CommerceKey: "96eb10fa-aa23-452c-9955-907aff1a3a2f",
    }

    res, err := s.SyncFlowPreferences.UpdateVisibleAccountsSyncFlow(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.VisibleAccounts != nil {
        // handle response
    }
}
```
