# Integrations

## Overview

View useful information about codat's integrations.

### Available Operations

* [GetIntegrationBranding](#getintegrationbranding) - Get branding for an integration
* [ListIntegrations](#listintegrations) - List information on Codat's supported integrations

## GetIntegrationBranding

Retrieve Integration branding assets.

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
    req := operations.GetIntegrationBrandingRequest{
        PlatformKey: "quibusdam",
    }

    res, err := s.Integrations.GetIntegrationBranding(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Branding != nil {
        // handle response
    }
}
```

## ListIntegrations

Retrieve a list of available integrations support by datatype and state of release.

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
    req := operations.ListIntegrationsRequest{
        OrderBy: codatsynccommerce.String("-modifiedDate"),
        Page: 1,
        PageSize: codatsynccommerce.Int(100),
        Query: codatsynccommerce.String("unde"),
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