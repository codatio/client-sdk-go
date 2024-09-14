# Transfers
(*Transfers*)

## Overview

Create and update transactions that represent the movement of your customers' money.

### Available Operations

* [Create](#create) - Create transfer transaction

## Create

Use the *Create transfer* endpoint to create or update a [transfer transaction](https://docs.codat.io/sync-for-expenses-api#/schemas/TransferTransactionRequest) in the accounting software for a given company's connection. 

Transfers record the movement of money between two bank accounts, or between a bank account and a nominal account. Use them to represent actions such as topping up a debit card account or a balance transfer to another credit card.

The `from.amount` and `to.amount` fields are in the native currency of the account.

### Supported Integrations
| Integration           | Supported |
|-----------------------|-----------|
| FreeAgent             | Yes       |
| QuickBooks Desktop    | Yes       |
| QuickBooks Online     | Yes       |
| Xero                  | Yes       |

### Example Usage

```go
package main

import(
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Transfers.Create(ctx, operations.CreateTransferTransactionRequest{
        TransferTransactionRequest: &shared.TransferTransactionRequest{
            Date: "2021-05-21T00:00:00+00:00",
            Description: syncforexpenses.String("Sample transfer description"),
            From: shared.From{
                AccountRef: shared.AccountReference{
                    ID: "787dfb37-5707-4dc0-8a86-8d74e4cc78ea",
                },
                Amount: types.MustNewDecimalFromString("100"),
            },
            To: shared.To{
                AccountRef: shared.TransferTransactionRequestAccountReference{
                    ID: "777dfb37-5506-3dc0-6g86-8d34z4cc78ea",
                },
                Amount: types.MustNewDecimalFromString("100"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateTransferTransactionRequest](../../pkg/models/operations/createtransfertransactionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateTransferTransactionResponse](../../pkg/models/operations/createtransfertransactionresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
