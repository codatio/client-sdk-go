# AccountingJournalEntries
(*AccountingJournalEntries*)

## Overview

Retrieve standardized Journal entries from linked accounting software.

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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"log"
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
            CreatedOn: syncforcommerceversion1.String("2023-02-22T19:49:16.052Z"),
            Description: syncforcommerceversion1.String("record level description"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("80000019-1671793811"),
                        Name: syncforcommerceversion1.String("Office Supplies"),
                    },
                    Currency: syncforcommerceversion1.String("USD"),
                    Description: syncforcommerceversion1.String("journalLines.description debit"),
                    NetAmount: types.MustNewDecimalFromString("23.02"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.TrackingRecordRef{
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeCustomers.ToPointer(),
                                ID: syncforcommerceversion1.String("80000001-1674553252"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: syncforcommerceversion1.String("8000001E-1671793811"),
                        Name: syncforcommerceversion1.String("Utilities"),
                    },
                    Currency: syncforcommerceversion1.String("USD"),
                    Description: syncforcommerceversion1.String("journalLines.description credit"),
                    NetAmount: types.MustNewDecimalFromString("-23.02"),
                    Tracking: &shared.JournalLineTracking{
                        RecordRefs: []shared.TrackingRecordRef{
                            shared.TrackingRecordRef{
                                DataType: shared.TrackingRecordRefDataTypeTrackingCategories.ToPointer(),
                                ID: syncforcommerceversion1.String("80000002-1674553271"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalReference{
                ID: "12",
            },
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(true),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            PostedOn: syncforcommerceversion1.String("2023-02-23T19:49:16.052Z"),
            RecordRef: &shared.JournalEntryRecordRef{
                DataType: shared.JournalEntryRecordRefDataTypeBills.ToPointer(),
                ID: syncforcommerceversion1.String("80000002-6722155312"),
            },
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00Z"),
            UpdatedOn: syncforcommerceversion1.String("2023-02-21T19:49:16.052Z"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateAccountingJournalEntryRequest](../../pkg/models/operations/createaccountingjournalentryrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CreateAccountingJournalEntryResponse](../../pkg/models/operations/createaccountingjournalentryresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
