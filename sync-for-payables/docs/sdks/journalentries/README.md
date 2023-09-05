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
            Description: codatsyncpayables.String("culpa"),
            ID: codatsyncpayables.String("bcdc91fa-abdd-488e-b1f6-c48252d7771e"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("fd074009-ef8d-429d-a1dd-7097b5da08c5"),
                        Name: codatsyncpayables.String("Ora Olson"),
                    },
                    Currency: codatsyncpayables.String("odio"),
                    Description: codatsyncpayables.String("atque"),
                    NetAmount: 6288.11,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("6e19bafe-ca61-4914-9814-0b64ff8ae170"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("ef03b5f3-7e4a-4a86-8555-966732aa5dcb"),
                        Name: codatsyncpayables.String("Ella Lang"),
                    },
                    Currency: codatsyncpayables.String("expedita"),
                    Description: codatsyncpayables.String("in"),
                    NetAmount: 526.59,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("accountTransaction"),
                                ID: codatsyncpayables.String("cfd5fb6e-91b9-4a9f-b484-6e2c3309db05"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("6d9e75ca-006f-4539-ac11-a25a8bf92f97"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("invoice"),
                                ID: codatsyncpayables.String("28ad9a9f-8bf8-4221-9253-59d98387f7a7"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("accountTransaction"),
                                ID: codatsyncpayables.String("cd72cd24-84da-4217-a9f2-ac41ef5725f1"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "169ac1e4-1d8a-423c-a3e3-4f2dfa4a197f",
                Name: codatsyncpayables.String("Betsy Walter"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.JournalEntryRecordReference{
                DataType: codatsyncpayables.String("invoice"),
                ID: codatsyncpayables.String("1fe17120-9985-43e9-b543-d854439ee224"),
            },
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "accusantium": map[string]interface{}{
                        "aliquam": "dolorem",
                        "expedita": "impedit",
                    },
                    "architecto": map[string]interface{}{
                        "magnam": "vitae",
                        "quos": "atque",
                    },
                },
            },
            UpdatedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(125488),
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

