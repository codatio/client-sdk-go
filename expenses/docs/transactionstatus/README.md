# TransactionStatus

## Overview

Retrieve the status of transactions within a sync.

### Available Operations

* [GetSyncTransaction](#getsynctransaction) - Get Sync Transaction
* [ListSyncTransactions](#listsynctransactions) - Get Sync transactions

## GetSyncTransaction

Gets the status of a transaction for a sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
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

    if res.TransactionMetadata != nil {
        // handle response
    }
}
```

## ListSyncTransactions

Get's the transactions and status for a sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/expenses"
	"github.com/codatio/client-sdk-go/expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.TransactionStatus.ListSyncTransactions(ctx, operations.ListSyncTransactionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Page: codatsyncexpenses.Int(1),
        PageSize: codatsyncexpenses.Int(100),
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
