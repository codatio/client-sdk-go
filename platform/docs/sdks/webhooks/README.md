# Webhooks
(*Webhooks*)

## Overview

Create and manage webhooks that listen to Codat's events.

### Available Operations

* [~~Create~~](#create) - Create webhook :warning: **Deprecated**
* [CreateConsumer](#createconsumer) - Create webhook consumer
* [DeleteConsumer](#deleteconsumer) - Delete webhook consumer
* [~~Get~~](#get) - Get webhook :warning: **Deprecated**
* [~~List~~](#list) - List webhooks :warning: **Deprecated**
* [ListConsumers](#listconsumers) - List webhook consumers

## ~~Create~~

Create a new webhook configuration

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Create(ctx, &shared.CreateRule{
        CompanyID: platform.String("39b73b17-cc2e-429e-915d-71654e9dcd1e"),
        Notifiers: shared.WebhookNotifier{
            Emails: []string{
                "info@client.com",
            },
            Webhook: platform.String("https://webhook.client.com"),
        },
        Type: "DataConnectionStatusChanged",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.CreateRule](../../pkg/models/shared/createrule.md)   | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreateRuleResponse](../../pkg/models/operations/createruleresponse.md), error**

### Errors

| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401,402,403,429,500,503 | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |


## CreateConsumer

﻿Use the *Create webhook consumer* endpoint to create a new webhook consumer that will listen to messages we send you.

[Webhook consumer](https://docs.codat.io/platform-api#/schemas/WebhookConsumer) is an HTTP endpoint that you configure to subscribe to specific events. See our documentation for more details on [Codat's webhook service](https://docs.codat.io/using-the-api/webhooks/overview).

### Tips and traps
- The number of webhook consumers you can create is limited to 50. If you have reached the maximum number of consumers, use the [*Delete webhook consumer*](https://docs.codat.io/platform-api#/operations/delete-webhook-consumer) endpoint to delete an unused consumer first.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.CreateConsumer(ctx, &shared.WebhookConsumerPrototype{
        CompanyID: platform.String("8a210b68-6988-11ed-a1eb-0242ac120002"),
        EventTypes: []string{
            "DataSyncCompleted",
            "Dataset data changed",
        },
        URL: platform.String("https://example.com/webhoook-consumer"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookConsumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.WebhookConsumerPrototype](../../pkg/models/shared/webhookconsumerprototype.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateWebhookConsumerResponse](../../pkg/models/operations/createwebhookconsumerresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## DeleteConsumer

﻿Use the *Delete webhook consumer* endpoint to delete an existing webhoook consumer, providing its valid `id` as a parameter.

[Webhook consumer](https://docs.codat.io/platform-api#/schemas/WebhookConsumer) is an HTTP endpoint that you configure to subscribe to specific events. See our documentation for more details on [Codat's webhook service](https://docs.codat.io/using-the-api/webhooks/overview).

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.DeleteConsumer(ctx, operations.DeleteWebhookConsumerRequest{
        WebhookID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteWebhookConsumerRequest](../../pkg/models/operations/deletewebhookconsumerrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DeleteWebhookConsumerResponse](../../pkg/models/operations/deletewebhookconsumerresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ~~Get~~

Get a single webhook

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Get(ctx, operations.GetWebhookRequest{
        RuleID: "7318949f-c008-4936-a8ff-10d7ab563fa6",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetWebhookRequest](../../pkg/models/operations/getwebhookrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetWebhookResponse](../../pkg/models/operations/getwebhookresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ~~List~~

List webhooks that you are subscribed to.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.List(ctx, operations.ListRulesRequest{
        OrderBy: platform.String("-modifiedDate"),
        Page: platform.Int(1),
        PageSize: platform.Int(100),
        Query: platform.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListRulesRequest](../../pkg/models/operations/listrulesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListRulesResponse](../../pkg/models/operations/listrulesresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## ListConsumers

﻿Use the *List webhook consumers* endpoint to return a list of all webhook consumers that currently exist for your client.

[Webhook consumer](https://docs.codat.io/platform-api#/schemas/WebhookConsumer) is an HTTP endpoint that you configure to subscribe to specific events. See our documentation for more details on [Codat's webhook service](https://docs.codat.io/using-the-api/webhooks/overview).

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.ListConsumers(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookConsumers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.ListWebhookConsumersResponse](../../pkg/models/operations/listwebhookconsumersresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
