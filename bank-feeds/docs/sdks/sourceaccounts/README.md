# SourceAccounts
(*SourceAccounts*)

## Overview

Provide and manage lists of source bank accounts.

### Available Operations

* [Create](#create) - Create single source account
* [CreateBatch](#createbatch) - Create source accounts
* [Delete](#delete) - Delete source account
* [DeleteCredentials](#deletecredentials) - Delete all source account credentials
* [GenerateCredentials](#generatecredentials) - Generate source account credentials
* [List](#list) - List source accounts
* [Update](#update) - Update source account

## Create

The _Create Source Account_ endpoint allows you to create a representation of a bank account within Codat's domain. The company can then map the source account to an existing or new target account in their accounting software.

> ### Versioning
> If you are integrating the Bank Feeds API with Codat after August 1, 2024, please use the v2 version of the API, as detailed in the schema below. For integrations completed before August 1, 2024, select the v1 version from the schema dropdown below.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.Create(ctx, operations.CreateSourceAccountRequest{
        RequestBody: bankfeeds.Pointer(operations.CreateCreateSourceAccountRequestBodySourceAccountV2Prototype(
            shared.SourceAccountV2Prototype{
                AccountInfo: &shared.AccountInfo{
                    AccountOpenDate: bankfeeds.String("2022-10-23"),
                },
                Currency: bankfeeds.String("USD"),
                ModifiedDate: bankfeeds.String("2022-10-23T00:00:00Z"),
            },
        )),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateSourceAccountRequest](../../pkg/models/operations/createsourceaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateSourceAccountResponse](../../pkg/models/operations/createsourceaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateBatch

The _Batch create source accounts_ endpoint allows you to create multiple representations of your SMB's bank accounts within Codat's domain. The company can then map the source account to an existing or new target account in their accounting software.

> ### Versioning
> If you are integrating the Bank Feeds API with Codat after August 1, 2024, please use the v2 version of the API, as detailed in the schema below. For integrations completed before August 1, 2024, select the v1 version from the schema dropdown below.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/types"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.CreateBatch(ctx, operations.CreateBatchSourceAccountRequest{
        RequestBody: bankfeeds.Pointer(operations.CreateCreateBatchSourceAccountRequestBodyArrayOfSourceAccount(
            []shared.SourceAccount{
                shared.SourceAccount{
                    AccountName: bankfeeds.String("account-081"),
                    AccountNumber: bankfeeds.String("12345670"),
                    AccountType: bankfeeds.String("Credit"),
                    Balance: types.MustNewDecimalFromString("99.99"),
                    Currency: bankfeeds.String("GBP"),
                    ID: "acc-002",
                    ModifiedDate: bankfeeds.String("2023-01-09T14:14:14.1057478Z"),
                    SortCode: bankfeeds.String("123456"),
                    Status: shared.StatusPending.ToPointer(),
                },
                shared.SourceAccount{
                    AccountName: bankfeeds.String("account-095"),
                    AccountNumber: bankfeeds.String("12345671"),
                    AccountType: bankfeeds.String("Credit"),
                    Balance: types.MustNewDecimalFromString("0"),
                    Currency: bankfeeds.String("USD"),
                    ID: "acc-003",
                    ModifiedDate: bankfeeds.String("2023-01-09T14:14:14.1057478Z"),
                    SortCode: bankfeeds.String("123456"),
                    Status: shared.StatusPending.ToPointer(),
                },
            },
        )),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateBatchSourceAccountRequest](../../pkg/models/operations/createbatchsourceaccountrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateBatchSourceAccountResponse](../../pkg/models/operations/createbatchsourceaccountresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |

## Delete

The _Delete source account_ endpoint enables you to remove a source account.

Removing a source account will also remove any mapping between the source bank feed bank accounts and the target bankfeed bank account.


### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.Delete(ctx, operations.DeleteSourceAccountRequest{
        AccountID: "7110701885",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteSourceAccountRequest](../../pkg/models/operations/deletesourceaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeleteSourceAccountResponse](../../pkg/models/operations/deletesourceaccountresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## DeleteCredentials

The _Delete Bank Account Credentials_ endpoint serves as a comprehensive mechanism for revoking all existing authorization credentials that a company employs to establish its Bank Feed connection.

In cases where multiple credential sets have been generated, a single API call to this endpoint revokes all of them.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.DeleteCredentials(ctx, operations.DeleteBankFeedCredentialsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteBankFeedCredentialsRequest](../../pkg/models/operations/deletebankfeedcredentialsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.DeleteBankFeedCredentialsResponse](../../pkg/models/operations/deletebankfeedcredentialsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GenerateCredentials

﻿The _Generate bank account credentials_ endpoint can be used to generate credentials for QuickBooks Online to authenticate the Bank Feed in the QBO portal. Each time this endpoint is called, a new set of credentials will be generated.

The old credentials will still be valid until the revoke credentials endpoint is used, which will revoke all credentials associated to the data connection.

> **For QuickBooks Online only**
>
> Only call this endpoint when onboarding SMBs that use  QuickBooks Online.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"os"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    requestBody, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }


    res, err := s.SourceAccounts.GenerateCredentials(ctx, operations.GenerateCredentialsRequest{
        RequestBody: requestBody,
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountCredentials != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GenerateCredentialsRequest](../../pkg/models/operations/generatecredentialsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GenerateCredentialsResponse](../../pkg/models/operations/generatecredentialsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## List

﻿The _List source accounts_ endpoint returns a list of [source accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankFeedAccount) for a given company's connection.

[Source accounts](https://docs.codat.io/bank-feeds-api#/schemas/BankFeedAccount) are the bank's bank account within Codat's domain from which transactions are synced into the accounting platform.

> ### Versioning
> If you are integrating the Bank Feeds API with Codat after August 1, 2024, please use the v2 version of the API, as detailed in the schema below. For integrations completed before August 1, 2024, select the v1 version from the schema dropdown below.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.List(ctx, operations.ListSourceAccountsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListSourceAccountsRequest](../../pkg/models/operations/listsourceaccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListSourceAccountsResponse](../../pkg/models/operations/listsourceaccountsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## Update

﻿The _Update source account_ endpoint updates a single source account for a single data connection connected to a single company.

### Tips and pitfalls

* This endpoint only updates the `accountName` field.
* Updates made here apply exclusively to source accounts and will not affect target accounts in the accounting software.

### Example Usage

```go
package main

import(
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/types"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.SourceAccounts.Update(ctx, operations.UpdateSourceAccountRequest{
        SourceAccount: &shared.SourceAccount{
            AccountName: bankfeeds.String("account-095"),
            AccountNumber: bankfeeds.String("12345671"),
            AccountType: bankfeeds.String("Credit"),
            Balance: types.MustNewDecimalFromString("0"),
            Currency: bankfeeds.String("USD"),
            FeedStartDate: bankfeeds.String("2022-10-23T00:00:00Z"),
            ID: "acc-003",
            ModifiedDate: bankfeeds.String("2023-01-09T14:14:14.1057478Z"),
            SortCode: bankfeeds.String("123456"),
            Status: shared.StatusPending.ToPointer(),
        },
        AccountID: "7110701885",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SourceAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateSourceAccountRequest](../../pkg/models/operations/updatesourceaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateSourceAccountResponse](../../pkg/models/operations/updatesourceaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |