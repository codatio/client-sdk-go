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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingJournalEntries.CreateAccountingJournalEntry(ctx, operations.CreateAccountingJournalEntryRequest{
        AccountingJournalEntry: &shared.AccountingJournalEntry{
            CreatedOn: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Description: codatsynccommerce.String("excepturi"),
            ID: codatsynccommerce.String("e81f30be-3e43-4202-9721-657650664187"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("d9d21f9a-d030-4c4e-8c11-a0836429068b"),
                        Name: codatsynccommerce.String("Pedro Armstrong"),
                    },
                    Currency: codatsynccommerce.String("quaerat"),
                    Description: codatsynccommerce.String("corporis"),
                    NetAmount: 8843.25,
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsynccommerce.String("transfer"),
                                ID: codatsynccommerce.String("73bc845e-320a-4319-b4ba-df947c9a867b"),
                            },
                            shared.RecordRef{
                                DataType: codatsynccommerce.String("transfer"),
                                ID: codatsynccommerce.String("42426665-816d-4dca-8ef5-1fcb4c593ec1"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.AccountingJournalEntryJournalReference{
                ID: "2cdaad0e-c7af-4edb-980d-f448a47f9390",
                Name: codatsynccommerce.String("Derek Lubowitz"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.AccountingJournalEntryRecordReference{
                DataType: codatsynccommerce.String("accountTransaction"),
                ID: codatsynccommerce.String("3dabf9ef-3ffd-4d9f-bf07-9af4d35724cd"),
            },
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "reiciendis": map[string]interface{}{
                        "vero": "eos",
                        "quas": "quasi",
                    },
                },
            },
            UpdatedOn: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(509799),
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

