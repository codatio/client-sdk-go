# Reimbursements
(*Reimbursements*)

## Overview

Create reimbursable expense transactions.

### Available Operations

* [Create](#create) - Create reimbursable expense transaction
* [Update](#update) - Update reimbursable expense transaction

## Create

Use the *Create reimbursable expense* endpoint to create a [reimbursement request](https://docs.codat.io/sync-for-expenses-api#/schemas/Reimburseable-Expense-Transactions) in the accounting platform for a given company's connection. 

Employee reimbursement requests are reflected in the accounting system in the form of Bills against an employee, who is a supplier.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Reimbursements.Create(ctx, operations.CreateReimbursableExpenseTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateReimbursableExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.CreateReimbursableExpenseTransactionRequest](../../pkg/models/operations/createreimbursableexpensetransactionrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |


### Response

**[*operations.CreateReimbursableExpenseTransactionResponse](../../pkg/models/operations/createreimbursableexpensetransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## Update

The *Update reimbursable expense* endpoint updates an existing [reimbursable expense transaction](https://docs.codat.io/sync-for-expenses-api#/operations/create-reimbursable-expense-transaction) in the accounting platform for a given company's connection. 

Employee reimbursement requests are reflected in the accounting system in the form of Bills against an employee, who is a supplier.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Reimbursements.Update(ctx, operations.UpdateReimbursableExpenseTransactionRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateReimbursableExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.UpdateReimbursableExpenseTransactionRequest](../../pkg/models/operations/updatereimbursableexpensetransactionrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |


### Response

**[*operations.UpdateReimbursableExpenseTransactionResponse](../../pkg/models/operations/updatereimbursableexpensetransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
