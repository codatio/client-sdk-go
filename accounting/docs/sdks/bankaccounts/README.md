# BankAccounts

## Overview

Bank accounts

### Available Operations

* [Create](#create) - Create bank account
* [Get](#get) - Get bank account
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bank account model
* [List](#list) - List bank accounts
* [Update](#update) - Update bank account

## Create

The *Create bank account* endpoint creates a new [bank account](https://docs.codat.io/accounting-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bank account model](https://docs.codat.io/accounting-api#/operations/get-create-update-bankAccounts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating an account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.Create(ctx, operations.CreateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("accusantium"),
            AccountNumber: codataccounting.String("iure"),
            AccountType: shared.BankAccountBankAccountTypeCredit.ToPointer(),
            AvailableBalance: codataccounting.Float64(9883.74),
            Balance: codataccounting.Float64(9589.5),
            Currency: codataccounting.String("GBP"),
            IBan: codataccounting.String("mollitia"),
            ID: codataccounting.String("3a2fa946-7739-4251-aa52-c3f5ad019da1"),
            Institution: codataccounting.String("reiciendis"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            NominalCode: codataccounting.String("vero"),
            OverdraftLimit: codataccounting.Float64(4686.51),
            SortCode: codataccounting.String("praesentium"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
        },
        AllowSyncOnPushComplete: codataccounting.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(55714),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateBankAccountRequest](../../models/operations/createbankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.CreateBankAccountResponse](../../models/operations/createbankaccountresponse.md), error**


## Get

The *Get bank account* endpoint returns a single account for a given accountId.

[Bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support getting a specific bank account.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.Get(ctx, operations.GetBankAccountRequest{
        AccountID: "omnis",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetBankAccountRequest](../../models/operations/getbankaccountrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.GetBankAccountResponse](../../models/operations/getbankaccountresponse.md), error**


## GetCreateUpdateModel

The *Get create/update bank account model* endpoint returns the expected data for the request payload when creating and updating a [bank account](https://docs.codat.io/accounting-api#/schemas/BankAccount) for a given company and integration.

[Bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating and updating a bank account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBankAccountsModelRequest{
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetCreateUpdateBankAccountsModelRequest](../../models/operations/getcreateupdatebankaccountsmodelrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |


### Response

**[*operations.GetCreateUpdateBankAccountsModelResponse](../../models/operations/getcreateupdatebankaccountsmodelresponse.md), error**


## List

The *List bank accounts* endpoint returns a list of [bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.List(ctx, operations.ListBankAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("voluptate"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListBankAccountsRequest](../../models/operations/listbankaccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.ListBankAccountsResponse](../../models/operations/listbankaccountsresponse.md), error**


## Update

The *Update bank account* endpoint updates an existing [bank account](https://docs.codat.io/accounting-api#/schemas/BankAccount) for a given company's connection.

[Bank accounts](https://docs.codat.io/accounting-api#/schemas/BankAccount) are financial accounts maintained by a bank or other financial institution.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update bank account model](https://docs.codat.io/accounting-api#/operations/get-create-update-bankAccounts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BankAccounts.Update(ctx, operations.UpdateBankAccountRequest{
        BankAccount: &shared.BankAccount{
            AccountName: codataccounting.String("cum"),
            AccountNumber: codataccounting.String("perferendis"),
            AccountType: shared.BankAccountBankAccountTypeUnknown.ToPointer(),
            AvailableBalance: codataccounting.Float64(4417.11),
            Balance: codataccounting.Float64(2828.07),
            Currency: codataccounting.String("EUR"),
            IBan: codataccounting.String("dicta"),
            ID: codataccounting.String("5471b5e6-e13b-499d-888e-1e91e450ad2a"),
            Institution: codataccounting.String("distinctio"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            NominalCode: codataccounting.String("labore"),
            OverdraftLimit: codataccounting.Float64(2647.3),
            SortCode: codataccounting.String("qui"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
        },
        BankAccountID: "cupiditate",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(552822),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateBankAccountRequest](../../models/operations/updatebankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.UpdateBankAccountResponse](../../models/operations/updatebankaccountresponse.md), error**

