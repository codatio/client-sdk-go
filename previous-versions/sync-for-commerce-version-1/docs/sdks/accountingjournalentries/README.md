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
	"github.com/ericlagergren/decimal"
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
            Description: codatsynccommerce.String("pariatur"),
            ID: codatsynccommerce.String("e008e6f8-c5f3-450d-8cdb-5a3418143010"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("421813d5-208e-4ce7-a253-b668451c6c6e"),
                        Name: codatsynccommerce.String("Helen Heller III"),
                    },
                    Currency: codatsynccommerce.String("at"),
                    Description: codatsynccommerce.String("vero"),
                    NetAmount: *types.MustNewDecimalFromString("6675.93"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.RecordRef{
                            shared.RecordRef{
                                DataType: codatsynccommerce.String("accountTransaction"),
                                ID: codatsynccommerce.String("3fec9578-a645-4842-b3a8-418d162309fb"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.AccountingJournalEntryJournalReference{
                ID: "0929921a-efb9-4f58-84d8-6e68e4be0560",
                Name: codatsynccommerce.String("Sheila Wolff"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            PostedOn: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            RecordRef: &shared.AccountingJournalEntryRecordReference{
                DataType: codatsynccommerce.String("invoice"),
                ID: codatsynccommerce.String("57a59ecf-ef66-4ef1-8aa3-383c2beb4773"),
            },
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "neque": map[string]interface{}{
                        "quo": "deleniti",
                    },
                },
            },
            UpdatedOn: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(437814),
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

