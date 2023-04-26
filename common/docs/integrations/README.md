# Integrations

## Overview

View and manage your available integrations in Codat.

### Available Operations

* [GetIntegration](#getintegration) - Get integration
* [GetIntegrationsBranding](#getintegrationsbranding) - Get branding
* [ListIntegrations](#listintegrations) - List integrations

## GetIntegration

Get single integration, by platformKey

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetIntegrationRequest{
        PlatformKey: "gbol",
    }

    res, err := s.Integrations.GetIntegration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

## GetIntegrationsBranding

Get branding for platform.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetIntegrationsBrandingRequest{
        PlatformKey: "gbol",
    }

    res, err := s.Integrations.GetIntegrationsBranding(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Branding != nil {
        // handle response
    }
}
```

## ListIntegrations

List your available integrations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListIntegrationsRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("veritatis"),
    }

    res, err := s.Integrations.ListIntegrations(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```
