# Integrations
(*Integrations*)

## Overview

Get a list of integrations supported by Codat and their logos.

### Available Operations

* [Get](#get) - Get integration
* [GetBranding](#getbranding) - Get branding
* [List](#list) - List integrations

## Get

Get single integration, by platformKey

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.Get(ctx, operations.GetIntegrationRequest{
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetIntegrationRequest](../../pkg/models/operations/getintegrationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetIntegrationResponse](../../pkg/models/operations/getintegrationresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## GetBranding

Get branding for platform.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.GetBranding(ctx, operations.GetIntegrationsBrandingRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetIntegrationsBrandingRequest](../../pkg/models/operations/getintegrationsbrandingrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetIntegrationsBrandingResponse](../../pkg/models/operations/getintegrationsbrandingresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## List

List your available integrations

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Integrations.List(ctx, operations.ListIntegrationsRequest{
        OrderBy: common.String("-modifiedDate"),
        Page: common.Int(1),
        PageSize: common.Int(100),
        Query: common.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListIntegrationsRequest](../../pkg/models/operations/listintegrationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListIntegrationsResponse](../../pkg/models/operations/listintegrationsresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
