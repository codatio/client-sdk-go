# Integrations
(*Integrations*)

## Overview

View useful information about codat's integrations.

### Available Operations

* [GetBranding](#getbranding) - Get branding for an integration
* [List](#list) - List integrations

## GetBranding

Retrieve Integration branding assets.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.GetBranding(ctx, operations.GetIntegrationBrandingRequest{
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


## List

Retrieve a list of available integrations support by data type and state of release.

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
)

func main() {
    s := syncforcommerce.New(
        syncforcommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.List(ctx, operations.ListIntegrationsRequest{
        OrderBy: syncforcommerce.String("-modifiedDate"),
        Page: syncforcommerce.Int(1),
        PageSize: syncforcommerce.Int(100),
        Query: syncforcommerce.String("Northeast Metal Canada"),
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

