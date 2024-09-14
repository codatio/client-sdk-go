# TransactionStatus
(*TransactionStatus*)

## Overview

Retrieve the status of transactions within a sync.

### Available Operations

* [GetSyncTransaction](#getsynctransaction) - Get sync transaction
* [ListSyncTransactions](#listsynctransactions) - Get sync transactions

## GetSyncTransaction

Gets the status of a transaction for a sync

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionStatus.GetSyncTransaction(ctx, operations.GetSyncTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetSyncTransactionRequest](../../pkg/models/operations/getsynctransactionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetSyncTransactionResponse](../../pkg/models/operations/getsynctransactionresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |


## ListSyncTransactions

Get's the transactions and status for a sync

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionStatus.ListSyncTransactions(ctx, operations.ListSyncTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: syncforexpensesversion1.Int(1),
        PageSize: syncforexpensesversion1.Int(100),
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionMetadataList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListSyncTransactionsRequest](../../pkg/models/operations/listsynctransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListSyncTransactionsResponse](../../pkg/models/operations/listsynctransactionsresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
