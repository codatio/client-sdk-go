# Journals

## Overview

Journals

### Available Operations

* [Create](#create) - Create journal
* [Get](#get) - Get journal
* [GetCreateModel](#getcreatemodel) - Get create journal model
* [List](#list) - List journals

## Create

The *Create journal* endpoint creates a new [journal](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) for a given company's connection.

[Journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) are used to record all the financial transactions of a company.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create journal model](https://docs.codat.io/sync-for-payroll-api#/operations/get-create-journals-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support creating a journal.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := codatsyncpayroll.New(
        codatsyncpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Journals.Create(ctx, operations.CreateJournalRequest{
        Journal: &shared.Journal{
            CreatedOn: codatsyncpayroll.String("2022-10-23T00:00:00.000Z"),
            HasChildren: codatsyncpayroll.Bool(false),
            ID: codatsyncpayroll.String("711e5b7f-d2ed-4028-921c-ddc692601fb5"),
            JournalCode: codatsyncpayroll.String("voluptate"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayroll.Bool(false),
            },
            ModifiedDate: codatsyncpayroll.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncpayroll.String("Gary Streich"),
            ParentID: codatsyncpayroll.String("perferendis"),
            SourceModifiedDate: codatsyncpayroll.String("2022-10-23T00:00:00.000Z"),
            Status: shared.JournalStatusUnknown.ToPointer(),
            Type: codatsyncpayroll.String("aut"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayroll.Int(764912),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateJournalRequest](../../models/operations/createjournalrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.CreateJournalResponse](../../models/operations/createjournalresponse.md), error**


## Get

The *Get journal* endpoint returns a single journal for a given `journalId`.

[Journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) are used to record all the financial transactions of a company.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support getting a specific journal.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := codatsyncpayroll.New(
        codatsyncpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Journals.Get(ctx, operations.GetJournalRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Journal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetJournalRequest](../../models/operations/getjournalrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.GetJournalResponse](../../models/operations/getjournalresponse.md), error**


## GetCreateModel

The *Get create journal model* endpoint returns the expected data for the request payload when creating a [journal](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) for a given company and integration.

[Journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) are used to record all the financial transactions of a company.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support creating a journal.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := codatsyncpayroll.New(
        codatsyncpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Journals.GetCreateModel(ctx, operations.GetCreateJournalModelRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCreateJournalModelRequest](../../models/operations/getcreatejournalmodelrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetCreateJournalModelResponse](../../models/operations/getcreatejournalmodelresponse.md), error**


## List

The *List journals* endpoint returns a list of [journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) for a given company's connection.

[Journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) are used to record all the financial transactions of a company.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := codatsyncpayroll.New(
        codatsyncpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Journals.List(ctx, operations.ListJournalsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncpayroll.String("-modifiedDate"),
        Page: codatsyncpayroll.Int(1),
        PageSize: codatsyncpayroll.Int(100),
        Query: codatsyncpayroll.String("hic"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Journals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListJournalsRequest](../../models/operations/listjournalsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.ListJournalsResponse](../../models/operations/listjournalsresponse.md), error**
