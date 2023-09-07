# JournalEntries

## Overview

Journal entries

### Available Operations

* [Create](#create) - Create journal entry
* [GetCreateModel](#getcreatemodel) - Get create journal entry model

## Create

The *Create journal entry* endpoint creates a new [journal entry](https://docs.codat.io/sync-for-payables-api#/schemas/JournalEntry) for a given company's connection.

[Journal entries](https://docs.codat.io/sync-for-payables-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-journalEntries-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating a journal entry.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Create(ctx, operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Description: codatsyncpayables.String("animi"),
            ID: codatsyncpayables.String("88970e18-9dbb-430f-8b33-ea055b197cd4"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("4e2f52d8-2d35-413b-b6f4-8b656bcdb35f"),
                        Name: codatsyncpayables.String("Russell Toy"),
                    },
                    Currency: codatsyncpayables.String("eos"),
                    Description: codatsyncpayables.String("reprehenderit"),
                    NetAmount: 3455.06,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("7a8cd9e7-319c-4177-9525-f77b114eeb52"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "ff785fc3-7814-4d4c-98e0-c2bb89eb75da",
                Name: codatsyncpayables.String("Elmer Emard"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.JournalEntryRecordReference{
                DataType: codatsyncpayables.String("journalEntry"),
                ID: codatsyncpayables.String("503d8bb3-1180-4f73-9ae9-e057eb809e28"),
            },
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptatem": map[string]interface{}{
                        "velit": "dolor",
                    },
                },
            },
            UpdatedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(953676),
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


## GetCreateModel

﻿The *Get create journal entry model* endpoint returns the expected data for the request payload when creating a [journal entry](https://docs.codat.io/sync-for-payables-api#/schemas/JournalEntry) for a given company and integration.

[Journal entries](https://docs.codat.io/sync-for-payables-api#/schemas/JournalEntry) are  made in a company's general ledger, or accounts, when transactions are approved.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating a journal entry.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

