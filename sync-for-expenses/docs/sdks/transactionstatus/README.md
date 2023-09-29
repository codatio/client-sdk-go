# TransactionStatus
(*TransactionStatus*)

## Overview

Retrieve the status of transactions within a sync.

### Available Operations

* [Get](#get) - Get sync transaction
* [List](#list) - List sync transactions

## Get

Gets the status of a transaction for a sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionStatus.Get(ctx, operations.GetSyncTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetSyncTransactionRequest](../../models/operations/getsynctransactionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GetSyncTransactionResponse](../../models/operations/getsynctransactionresponse.md), error**


## List

Gets the transactions and status for a sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionStatus.List(ctx, operations.ListSyncTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: syncforexpenses.Int(1),
        PageSize: syncforexpenses.Int(100),
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListSyncTransactionsRequest](../../models/operations/listsynctransactionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.ListSyncTransactionsResponse](../../models/operations/listsynctransactionsresponse.md), error**

