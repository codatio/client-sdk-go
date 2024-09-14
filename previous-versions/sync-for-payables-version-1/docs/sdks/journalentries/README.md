# JournalEntries
(*JournalEntries*)

## Overview

Get, create, and update Journal entries.

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
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.JournalEntries.Create(ctx, operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: syncforpayables.String("2023-02-22T19:49:16.052Z"),
            Description: syncforpayables.String("record level description"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("80000019-1671793811"),
                        Name: syncforpayables.String("Office Supplies"),
                    },
                    Currency: syncforpayables.String("USD"),
                    Description: syncforpayables.String("journalLines.description debit"),
                    NetAmount: types.MustNewDecimalFromString("23.02"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.TrackingRecordRef{
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeCustomers.ToPointer(),
                                ID: syncforpayables.String("80000001-1674553252"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforpayables.String("8000001E-1671793811"),
                        Name: syncforpayables.String("Utilities"),
                    },
                    Currency: syncforpayables.String("USD"),
                    Description: syncforpayables.String("journalLines.description credit"),
                    NetAmount: types.MustNewDecimalFromString("-23.02"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.TrackingRecordRef{
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeTrackingCategories.ToPointer(),
                                ID: syncforpayables.String("80000002-1674553271"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "12",
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforpayables.Bool(true),
            },
            ModifiedDate: syncforpayables.String("2022-10-23T00:00:00Z"),
            PostedOn: syncforpayables.String("2023-02-23T19:49:16.052Z"),
            RecordRef: &shared.JournalEntryRecordRef{
                DataType: shared.JournalEntryRecordRefDataTypeBills.ToPointer(),
                ID: syncforpayables.String("80000002-6722155312"),
            },
            SourceModifiedDate: syncforpayables.String("2022-10-23T00:00:00Z"),
            UpdatedOn: syncforpayables.String("2023-02-21T19:49:16.052Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateJournalEntryRequest](../../pkg/models/operations/createjournalentryrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateJournalEntryResponse](../../pkg/models/operations/createjournalentryresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


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
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v3"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/operations"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCreateJournalEntryModelRequest](../../pkg/models/operations/getcreatejournalentrymodelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetCreateJournalEntryModelResponse](../../pkg/models/operations/getcreatejournalentrymodelresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
