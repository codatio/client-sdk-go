# Sync

## Overview

Triggering a new sync of expenses to accounting software.

### Available Operations

* [IntiateSync](#intiatesync) - Initiate sync

## IntiateSync

Initiate sync of pending transactions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Sync.IntiateSync(ctx, operations.IntiateSyncRequest{
        PostSync: &shared.PostSync{
            DatasetIds: []string{
                "96ed151a-05df-4c2d-9f7c-c78ca1ba928f",
                "c816742c-b739-4205-9293-96fea7596eb1",
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncInitiated != nil {
        // handle response
    }
}
```
