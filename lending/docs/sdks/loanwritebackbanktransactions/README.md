# LoanWritebackBankTransactions
(*LoanWriteback.BankTransactions*)

### Available Operations

* [Create](#create) - Create bank account transactions
* [GetCreateModel](#getcreatemodel) - Get create bank account transactions model

## Create

The *Create bank account transactions* endpoint creates new [bank account transactions](https://docs.codat.io/lending-api#/schemas/BankTransactions) for a given company's connection.

[Bank account transactions](https://docs.codat.io/lending-api#/schemas/BankTransactions) are records of money that has moved in and out of an SMB's bank account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create bank transaction model](https://docs.codat.io/lending-api#/operations/get-create-bankTransactions-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support creating a bank account transactions.


### Example Usage

```go
package main

import(
	"context"
	"log"
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.BankTransactions.Create(ctx, operations.CreateBankTransactionsRequest{
        AccountingCreateBankTransactions: &shared.AccountingCreateBankTransactions{
            AccountID: lending.String("7110701885"),
            Transactions: []shared.CreateBankAccountTransaction{
                shared.CreateBankAccountTransaction{
                    Date: lending.String("2022-10-23T00:00:00.000Z"),
                },
            },
        },
        AccountID: "Extended South",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateBankTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateBankTransactionsRequest](../../models/operations/createbanktransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateBankTransactionsResponse](../../models/operations/createbanktransactionsresponse.md), error**


## GetCreateModel

The *Get create bank account transactions model* endpoint returns the expected data for the request payload when creating [bank account transactions](https://docs.codat.io/lending-api#/schemas/BankTransactions) for a given company and integration.

[Bank account transactions](https://docs.codat.io/lending-api#/schemas/BankTransactions) are records of money that has moved in and out of an SMB's bank account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions) for integrations that support creating an bank transaction.


### Example Usage

```go
package main

import(
	"context"
	"log"
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.BankTransactions.GetCreateModel(ctx, operations.GetCreateBankTransactionsModelRequest{
        AccountID: "Northwest",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetCreateBankTransactionsModelRequest](../../models/operations/getcreatebanktransactionsmodelrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.GetCreateBankTransactionsModelResponse](../../models/operations/getcreatebanktransactionsmodelresponse.md), error**
