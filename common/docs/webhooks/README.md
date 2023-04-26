# Webhooks

## Overview

Manage webhooks, rules and alerts.

### Available Operations

* [CreateRule](#createrule) - Create webhook
* [GetWebhook](#getwebhook) - Get webhook
* [ListRules](#listrules) - List webhooks

## CreateRule

Create a new webhook configuration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.Rule{
        CompanyID: codatcommon.String("39b73b17-cc2e-429e-915d-71654e9dcd1e"),
        ID: "ff89c50e-a719-4ef5-a182-9917e53927b6",
        Notifiers: shared.RuleNotifiers{
            Emails: []string{
                "info@client.com",
                "info@client.com",
                "info@client.com",
            },
            Webhook: codatcommon.String("https://webhook.client.com"),
        },
        Type: "dolorem",
    }

    res, err := s.Webhooks.CreateRule(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rule != nil {
        // handle response
    }
}
```

## GetWebhook

Get a single webhook

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
    req := operations.GetWebhookRequest{
        RuleID: "7318949f-c008-4936-a8ff-10d7ab563fa6",
    }

    res, err := s.Webhooks.GetWebhook(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rule != nil {
        // handle response
    }
}
```

## ListRules

List webhooks that you are subscribed to.

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
    req := operations.ListRulesRequest{
        OrderBy: codatcommon.String("-modifiedDate"),
        Page: 1,
        PageSize: codatcommon.Int(100),
        Query: codatcommon.String("culpa"),
    }

    res, err := s.Webhooks.ListRules(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Rules != nil {
        // handle response
    }
}
```
