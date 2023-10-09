# Integrations
(*Integrations*)

## Overview

View useful information about codat's integrations.

### Available Operations

* [GetIntegrationBranding](#getintegrationbranding) - Get branding for an integration
* [ListIntegrations](#listintegrations) - List integrations

## GetIntegrationBranding

Retrieve Integration branding assets.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.GetIntegrationBranding(ctx, operations.GetIntegrationBrandingRequest{
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Branding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetIntegrationBrandingRequest](../../models/operations/getintegrationbrandingrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GetIntegrationBrandingResponse](../../models/operations/getintegrationbrandingresponse.md), error**


## ListIntegrations

Retrieve a list of available integrations support by datatype and state of release.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.ListIntegrations(ctx, operations.ListIntegrationsRequest{
        OrderBy: syncforcommerceversion1.String("-modifiedDate"),
        Page: syncforcommerceversion1.Int(1),
        PageSize: syncforcommerceversion1.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListIntegrationsRequest](../../models/operations/listintegrationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.ListIntegrationsResponse](../../models/operations/listintegrationsresponse.md), error**

