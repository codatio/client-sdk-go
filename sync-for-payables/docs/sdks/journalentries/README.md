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
            Description: codatsyncpayables.String("nam"),
            ID: codatsyncpayables.String("dc41ff5d-4e2a-4e4f-b5cb-35d17638f1ed"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("78359ecc-5cb8-460f-8cd5-80ba73810e4f"),
                        Name: codatsyncpayables.String("Don Hagenes"),
                    },
                    Currency: codatsyncpayables.String("magni"),
                    Description: codatsyncpayables.String("excepturi"),
                    NetAmount: 4576.85,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("transfer"),
                                ID: codatsyncpayables.String("3b1dd3bb-ce24-47b7-a84e-ff50126d71cf"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("transfer"),
                                ID: codatsyncpayables.String("bd0eb74b-8421-4953-b44b-d3c43159d33e"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("invoice"),
                                ID: codatsyncpayables.String("953c0011-3986-43aa-81e6-c31cc2f1fcb5"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("c9a41ffb-e9cb-4d79-9ee6-5e076cc7abf6"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("16ea5c71-6419-434b-90f2-e09d19d2fc2f"),
                        Name: codatsyncpayables.String("Merle Cormier Jr."),
                    },
                    Currency: codatsyncpayables.String("nemo"),
                    Description: codatsyncpayables.String("provident"),
                    NetAmount: 2529.57,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("accountTransaction"),
                                ID: codatsyncpayables.String("935d237a-72f9-4084-9d6a-ed4aecb7537c"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("transfer"),
                                ID: codatsyncpayables.String("9222c9ff-5749-41aa-bfa2-e761f0ca4d45"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsyncpayables.String("6ef1031e-6899-4f0c-a001-e22cd55cc058"),
                        Name: codatsyncpayables.String("Hattie Botsford"),
                    },
                    Currency: codatsyncpayables.String("possimus"),
                    Description: codatsyncpayables.String("nihil"),
                    NetAmount: 3758.77,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("accountTransaction"),
                                ID: codatsyncpayables.String("71fc820c-65b0-437b-b8e0-cc885187e4de"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("4af28c5d-ddb4-46aa-9cfd-6d828da01319"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("journalEntry"),
                                ID: codatsyncpayables.String("12964664-5c1d-481f-a904-2f569b7aff0e"),
                            },
                            shared.RecordRef{
                                DataType: codatsyncpayables.String("accountTransaction"),
                                ID: codatsyncpayables.String("2216cbe0-71bc-4163-a279-a3b084da9925"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "7d04f408-47a7-442d-8449-6cbdeecf6b99",
                Name: codatsyncpayables.String("Wilbert Jerde"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.JournalEntryRecordReference{
                DataType: codatsyncpayables.String("transfer"),
                ID: codatsyncpayables.String("bfdf55c2-94c0-460b-86a1-287764eef6d0"),
            },
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "temporibus": map[string]interface{}{
                        "itaque": "nulla",
                        "excepturi": "quod",
                    },
                    "in": map[string]interface{}{
                        "temporibus": "temporibus",
                    },
                },
            },
            UpdatedOn: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(247927),
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

