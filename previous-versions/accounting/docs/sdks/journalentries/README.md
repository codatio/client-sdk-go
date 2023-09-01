# JournalEntries

## Overview

Journal entries

### Available Operations

* [Create](#create) - Create journal entry
* [Delete](#delete) - Delete journal entry
* [Get](#get) - Get journal entry
* [GetCreateModel](#getcreatemodel) - Get create journal entry model
* [List](#list) - List journal entries

## Create

The *Create journal entry* endpoint creates a new [journal entry](https://docs.codat.io/accounting-api#/schemas/JournalEntry) for a given company's connection.

[Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/accounting-api#/operations/get-create-journalEntries-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Create(ctx, operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Description: codataccounting.String("illum"),
            ID: codataccounting.String("290278eb-4ae9-4d64-961e-91500323b2c0"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b924771f-5669-4e5b-bec7-626649d84eb9"),
                        Name: codataccounting.String("Jesus Runte"),
                    },
                    Currency: codataccounting.String("dolores"),
                    Description: codataccounting.String("odit"),
                    NetAmount: 4764.38,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("transfer"),
                                ID: codataccounting.String("0b88fb87-d6fa-45b6-a8db-f812f83b1ca6"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("accountTransaction"),
                                ID: codataccounting.String("9ffc5619-29cc-4a95-a0a1-395918da1d48"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e78e3cf8-e114-43da-9308-b27a08af2218"),
                        Name: codataccounting.String("Bertha Fay"),
                    },
                    Currency: codataccounting.String("nesciunt"),
                    Description: codataccounting.String("illum"),
                    NetAmount: 9253.93,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("invoice"),
                                ID: codataccounting.String("56ccce47-0cd2-4147-b6e6-152cf01d0d8c"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("journalEntry"),
                                ID: codataccounting.String("a4b9a5bf-935d-4fe9-b4fa-4b1e9c097eda"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("invoice"),
                                ID: codataccounting.String("23442e1a-9237-4e99-84c8-0b479e891923"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c18ca8d6-9c56-4892-94fa-20207e4fae03"),
                        Name: codataccounting.String("Colin Schultz"),
                    },
                    Currency: codataccounting.String("ab"),
                    Description: codataccounting.String("cum"),
                    NetAmount: 7949.86,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("transfer"),
                                ID: codataccounting.String("abaf7fc2-ccba-44be-b0df-68eaedb2ee70"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "be069fb3-6add-4704-880e-0a3fc73a5a03",
                Name: codataccounting.String("Mrs. Patty Block"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codataccounting.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("invoice"),
                ID: codataccounting.String("3afa6987-a472-4b70-9a15-3e2230106853"),
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "saepe": map[string]interface{}{
                        "perspiciatis": "consectetur",
                    },
                    "eos": map[string]interface{}{
                        "vitae": "aperiam",
                        "animi": "quisquam",
                        "possimus": "vitae",
                        "ad": "quibusdam",
                    },
                    "praesentium": map[string]interface{}{
                        "quisquam": "ipsum",
                        "perferendis": "ex",
                        "quidem": "esse",
                        "praesentium": "nisi",
                    },
                    "libero": map[string]interface{}{
                        "at": "amet",
                    },
                },
            },
            UpdatedOn: codataccounting.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(690696),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateJournalEntryRequest](../../models/operations/createjournalentryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.CreateJournalEntryResponse](../../models/operations/createjournalentryresponse.md), error**


## Delete

﻿> **Use with caution**
>
>Because journal entries underpin every transaction in an accounting platform, deleting a journal entry can affect every transaction for a given company.
> 
> **Before you proceed, make sure you understand the implications of deleting journal entries from an accounting perspective.**

The *Delete journal entry* endpoint allows you to delete a specified journal entry from an accounting platform.

[Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) are made in a company's general ledger, or accounts, when transactions are approved.

### Process
1. Pass the `{journalEntryId}` to the *Delete journal entry* endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete by checking the status of push operation either via
   1. [Push operation webhook](https://docs.codat.io/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
   2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation). 
   
   A `Success` status indicates that the journal entry object was deleted from the accounting platform.
3. (Optional) Check that the journal entry was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting a journal entry from an accounting platform might cause related objects to be modified. For example, if you delete the journal entry for a paid invoice in QuickBooks Online, the invoice is deleted but the payment against that invoice is not. The payment is converted to a payment on account.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Deleted | 
|-------------|--------------|
| QuickBooks Online | Yes    |       

> **Supported Integrations**
> 
> This functionality is currently only supported for our QuickBooks Online integration. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Delete(ctx, operations.DeleteJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        JournalEntryID: "fugiat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteJournalEntryRequest](../../models/operations/deletejournalentryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.DeleteJournalEntryResponse](../../models/operations/deletejournalentryresponse.md), error**


## Get

The *Get journal entry* endpoint returns a single journal entry for a given journalEntryId.

[Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support getting a specific journal entry.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Get(ctx, operations.GetJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalEntryID: "sed",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetJournalEntryRequest](../../models/operations/getjournalentryrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetJournalEntryResponse](../../models/operations/getjournalentryresponse.md), error**


## GetCreateModel

﻿The *Get create journal entry model* endpoint returns the expected data for the request payload when creating a [journal entry](https://docs.codat.io/accounting-api#/schemas/JournalEntry) for a given company and integration.

[Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating a journal entry.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.GetCreateModel(ctx, operations.GetCreateJournalEntriesModelRequest{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCreateJournalEntriesModelRequest](../../models/operations/getcreatejournalentriesmodelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetCreateJournalEntriesModelResponse](../../models/operations/getcreatejournalentriesmodelresponse.md), error**


## List

The *List journal entries* endpoint returns a list of [journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) for a given company's connection.

[Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.List(ctx, operations.ListJournalEntriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("aut"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListJournalEntriesRequest](../../models/operations/listjournalentriesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.ListJournalEntriesResponse](../../models/operations/listjournalentriesresponse.md), error**

