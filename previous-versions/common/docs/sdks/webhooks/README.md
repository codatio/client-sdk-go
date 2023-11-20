# Webhooks
(*Webhooks*)

## Overview

Manage webhooks, rules, and events.

### Available Operations

* [Create](#create) - Create webhook
* [Get](#get) - Get webhook
* [List](#list) - List webhooks

## Create

Create a new webhook configuration

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Create(ctx, &shared.CreateRule{
        CompanyID: common.String("8a210b68-6988-11ed-a1eb-0242ac120002"),
        Notifiers: shared.WebhookNotifier{
            Emails: []string{
                "i",
                "n",
                "f",
                "o",
                "@",
                "c",
                "l",
                "i",
                "e",
                "n",
                "t",
                ".",
                "c",
                "o",
                "m",
            },
            Webhook: common.String("https://webhook.client.com"),
        },
        Type: "string",
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
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401,402,403,429,500,503 | application/json        |
| sdkerrors.SDKError      | 400-600                 | */*                     |

## Get

Get a single webhook

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
            AuthHeader: "",
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
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## List

List webhooks that you are subscribed to.

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
            AuthHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.List(ctx, operations.ListRulesRequest{
        OrderBy: common.String("-modifiedDate"),
        Page: common.Int(1),
        PageSize: common.Int(100),
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
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
