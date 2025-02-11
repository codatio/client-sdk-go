# SourceAccounts
(*LoanWriteback.SourceAccounts*)

## Overview

### Available Operations

* [Create](#create) - Create source account
* [CreateMapping](#createmapping) - Create bank feed account mapping
* [ListMappings](#listmappings) - List bank feed account mappings

## Create

The _Create Source Account_ endpoint allows you to create a representation of a bank account within Codat's domain. The company can then map the source account to an existing or new target account in their accounting software.

> ### Versioning
> If you are integrating the Bank Feeds API with Codat after August 1, 2024, please use the v2 version of the API, as detailed in the schema below. For integrations completed before August 1, 2024, select the v1 version from the schema dropdown below.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.LoanWriteback.SourceAccounts.Create(ctx, operations.CreateSourceAccountRequest{
        RequestBody: lending.Pointer(operations.CreateCreateSourceAccountRequestBodySourceAccountV2Prototype(
            shared.SourceAccountV2Prototype{
                AccountInfo: &shared.AccountInfo{
                    AccountOpenDate: lending.String("2022-10-23"),
                },
                Currency: lending.String("USD"),
                ModifiedDate: lending.String("2022-10-23T00:00:00Z"),
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

## CreateMapping

﻿The *Create bank account mapping* endpoint creates a new mapping between a source bank account and a potential account in the accounting software (target account).

A bank feed account mapping is a specified link between the source account (provided by the Codat user) and the target account (the end user's account in the underlying software).

To find valid target account options, first call the [List bank feed account mappings](https://docs.codat.io//bank-feeds-api#/operations/get-bank-account-mapping) endpoint.

> **For custom builds only**
>
> Only use this endpoint if you are building your own account management UI.

#### Account mapping variability

The method of mapping the source account to the target account varies depending on the accounting software your company uses.

#### Mapping options:

1. **API Mapping**: Integrate the mapping journey directly into your application for a seamless user experience.
2. **Codat UI Mapping**: If you prefer a quicker setup, you can utilize Codat's provided user interface for mapping.
3. **Accounting Platform Mapping**: For some accounting software, the mapping process must be conducted within the software itself.

### Integration-specific behaviour

| Bank Feed Integration | API Mapping | Codat UI Mapping | Accounting Platform Mapping |
| --------------------- | ----------- | ---------------- | --------------------------- |
| Xero                  | ✅          | ✅               |                             |
| FreeAgent             | ✅          | ✅               |                             |
| Oracle NetSuite       | ✅          | ✅               |                             |
| Exact Online (NL)     | ✅          | ✅               |                             |
| QuickBooks Online     |             |                  | ✅                          |
| Sage                  |             |                  | ✅                          |

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.LoanWriteback.SourceAccounts.CreateMapping(ctx, operations.CreateBankAccountMappingRequest{
        BankFeedBankAccountMapping: &shared.BankFeedBankAccountMapping{
            SourceAccountID: "acc-002",
            TargetAccountID: lending.String("account-081"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankFeedBankAccountMappingResponse != nil {
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

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## ListMappings

﻿The *List bank accounts* endpoint returns information about a source bank account and any current or potential target mapping accounts.

A bank feed account mapping is a specified link between the source account (provided by the Codat user) and the target account (the end user's account in the underlying software).

> **For custom builds only**
> 
> Only use this endpoint if you are building your own account management UI.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.LoanWriteback.SourceAccounts.ListMappings(ctx, operations.GetBankAccountMappingRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankFeedMappings != nil {
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

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |