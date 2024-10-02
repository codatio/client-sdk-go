# BankAccounts
(*BankAccounts*)

## Overview

Create a bank account for a given company's connection.

### Available Operations

* [Create](#create) - Create bank account

## Create

The *Create bank account* endpoint creates a new [bank account](https://docs.codat.io/sync-for-payables-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/sync-for-payables-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.Create(ctx, operations.CreateBankAccountRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BankAccountPrototype: &shared.BankAccountPrototype{
            NominalCode: syncforpayables.String("22"),
            Name: syncforpayables.String("Plutus - Payables - Bank Account 12"),
            AccountType: shared.AccountTypeDebit,
            AccountNumber: syncforpayables.String("0120 0440"),
            SortCode: syncforpayables.String("50-50-50"),
            Currency: "GBP",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateBankAccountRequest](../../pkg/models/operations/createbankaccountrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateBankAccountResponse](../../pkg/models/operations/createbankaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |