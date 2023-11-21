# AccountMapping
(*AccountMapping*)

## Overview

Bank feed bank account mapping.

### Available Operations

* [Create](#create) - Create bank feed account mapping
* [Get](#get) - List bank feed account mappings

## Create

﻿The *Create bank account mapping* endpoint creates a new mapping between a source bank account and a potential account in the accounting platform (target account).

A bank feed account mapping is a specified link between the source account (provided by the Codat user) and the target account (the end users account in the underlying platform).

To find valid target account options, first call list bank feed account mappings.

This endpoint is only needed if building an account management UI.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountMapping.Create(ctx, operations.CreateBankAccountMappingRequest{
        Zero: &shared.Zero{
            FeedStartDate: bankfeeds.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccountMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateBankAccountMappingRequest](../../pkg/models/operations/createbankaccountmappingrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.CreateBankAccountMappingResponse](../../pkg/models/operations/createbankaccountmappingresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## Get

﻿The *List bank account mappings* endpoint returns information about a source bank account and any current or potential target mapping accounts.

A bank feed account mapping is a specified link between the source account (provided by the Codat user) and the target account (the end users account in the underlying platform).

This endpoint is only needed if building an account management UI.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"context"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/operations"
	"log"
)

func main() {
    s := bankfeeds.New(
        bankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountMapping.Get(ctx, operations.GetBankAccountMappingRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedMapping != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetBankAccountMappingRequest](../../pkg/models/operations/getbankaccountmappingrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetBankAccountMappingResponse](../../pkg/models/operations/getbankaccountmappingresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |
