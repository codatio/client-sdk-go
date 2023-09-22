# AccountingJournalEntries

## Overview

Journal entries

### Available Operations

* [CreateAccountingJournalEntry](#createaccountingjournalentry) - Create journal entry

## CreateAccountingJournalEntry

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
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingJournalEntries.CreateAccountingJournalEntry(ctx, operations.CreateAccountingJournalEntryRequest{
        AccountingJournalEntry: &shared.AccountingJournalEntry{
            CreatedOn: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Description: syncforcommerceversion1.String("nam"),
            ID: syncforcommerceversion1.String("5a341814-3010-4421-813d-5208ece7e253"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("b668451c-6c6e-4205-a16d-eab3fec9578a"),
                        Name: syncforcommerceversion1.String("Marjorie Hickle"),
                    },
                    Currency: syncforcommerceversion1.String("aspernatur"),
                    Description: syncforcommerceversion1.String("ducimus"),
                    NetAmount: types.MustNewDecimalFromString("2005.16"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: syncforcommerceversion1.String("accountTransaction"),
                                ID: syncforcommerceversion1.String("8418d162-309f-4b09-a992-1aefb9f58c4d"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.AccountingJournalEntryJournalReference{
                ID: "86e68e4b-e056-4013-b59d-a757a59ecfef",
                Name: syncforcommerceversion1.String("Loretta Tremblay DDS"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PostedOn: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.AccountingJournalEntryRecordReference{
                DataType: syncforcommerceversion1.String("journalEntry"),
                ID: syncforcommerceversion1.String("383c2beb-4773-473c-8d72-f64d1db1f2c4"),
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "illo": map[string]interface{}{
                        "accusantium": "vel",
                    },
                },
            },
            UpdatedOn: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(107617),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateAccountingJournalEntryRequest](../../models/operations/createaccountingjournalentryrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.CreateAccountingJournalEntryResponse](../../models/operations/createaccountingjournalentryresponse.md), error**

