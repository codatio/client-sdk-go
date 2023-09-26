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

The *Create journal entry* endpoint creates a new [journal entry](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) for a given company's connection.

[Journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/sync-for-payroll-api#/operations/get-create-journalEntries-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating a journal entry.


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/types"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Create(ctx, operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            Description: syncforpayroll.String("temporibus"),
            ID: syncforpayroll.String("151a05df-c2dd-4f7c-878c-a1ba928fc816"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayroll.String("742cb739-2059-4293-96fe-a7596eb10faa"),
                        Name: syncforpayroll.String("Ernest Ebert"),
                    },
                    Currency: syncforpayroll.String("nobis"),
                    Description: syncforpayroll.String("enim"),
                    NetAmount: types.MustNewDecimalFromString("6078.31"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: syncforpayroll.String("invoice"),
                                ID: syncforpayroll.String("5907aff1-a3a2-4fa9-8677-39251aa52c3f"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "5ad019da-1ffe-478f-897b-0074f15471b5",
                Name: syncforpayroll.String("Mrs. Leslie VonRueden"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayroll.Bool(false),
            },
            ModifiedDate: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            PostedOn: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.JournalEntryRecordReference{
                DataType: syncforpayroll.String("transfer"),
                ID: syncforpayroll.String("488e1e91-e450-4ad2-abd4-4269802d502a"),
            },
            SourceModifiedDate: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.JournalEntrySupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempora": map[string]interface{}{
                        "facilis": "tempore",
                    },
                },
            },
            UpdatedOn: syncforpayroll.String("2022-10-23T00:00:00.000Z"),
        },
        AllowSyncOnPushComplete: syncforpayroll.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforpayroll.Int(962189),
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

[Journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) are made in a company's general ledger, or accounts, when transactions are approved.

### Process
1. Pass the `{journalEntryId}` to the *Delete journal entry* endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete by checking the status of push operation either via
   1. [Push operation webhook](https://docs.codat.io/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
   2. [Push operation status endpoint](https://docs.codat.io/sync-for-payroll-api#/operations/get-push-operation). 
   
   A `Success` status indicates that the journal entry object was deleted from the accounting platform.
3. (Optional) Check that the journal entry was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting a journal entry from an accounting platform might cause related objects to be modified. For example, if you delete the journal entry for a paid invoice in QuickBooks Online, the invoice is deleted but the payment against that invoice is not. The payment is converted to a payment on account.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Deleted | 
|-------------|--------------|
| QuickBooks Online | Yes    |       


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Delete(ctx, operations.DeleteJournalEntryRequest{
        AllowSyncOnPushComplete: syncforpayroll.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        JournalEntryID: "eum",
        TimeoutInMinutes: syncforpayroll.Int(248753),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperation != nil {
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

The *Get journal entry* endpoint returns a single journal entry for a given `journalEntryId`.

[Journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support getting a specific journal entry.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Get(ctx, operations.GetJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalEntryID: "eligendi",
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

﻿The *Get create journal entry model* endpoint returns the expected data for the request payload when creating a [journal entry](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) for a given company and integration.

[Journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating a journal entry.


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.GetCreateModel(ctx, operations.GetCreateJournalEntryModelRequest{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetCreateJournalEntryModelRequest](../../models/operations/getcreatejournalentrymodelrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.GetCreateJournalEntryModelResponse](../../models/operations/getcreatejournalentrymodelresponse.md), error**


## List

The *List journal entries* endpoint returns a list of [journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) for a given company's connection.

[Journal entries](https://docs.codat.io/sync-for-payroll-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payroll-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/models/operations"
)

func main() {
    s := syncforpayroll.New(
        syncforpayroll.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.List(ctx, operations.ListJournalEntriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforpayroll.String("-modifiedDate"),
        Page: syncforpayroll.Int(1),
        PageSize: syncforpayroll.Int(100),
        Query: syncforpayroll.String("sint"),
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

