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

Posts a new journalEntry to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/accounting-api#/operations/get-create-journalEntries-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) to see which integrations support this endpoint.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
            CreatedOn: codataccounting.String("quas"),
            Description: codataccounting.String("soluta"),
            ID: codataccounting.String("d380c29a-a8dd-471b-9daa-30b7b91449ae"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9c088d41-8bb7-4180-8f42-3d543935f377"),
                        Name: codataccounting.String("Devin Hintz"),
                    },
                    Currency: codataccounting.String("nam"),
                    Description: codataccounting.String("ducimus"),
                    NetAmount: 8873.03,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("velit"),
                                ID: codataccounting.String("b6a3c523-105e-47c3-8cab-0ecb812a6614"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("atque"),
                                ID: codataccounting.String("944a8e90-8507-45bc-a538-253343fb0a4e"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("autem"),
                                ID: codataccounting.String("6ea47578-d171-4e29-8181-8fc679b6b2f2"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5359b855-d015-4b62-88b8-3a38a8a88c14"),
                        Name: codataccounting.String("Miss Andrea Bartell"),
                    },
                    Currency: codataccounting.String("maxime"),
                    Description: codataccounting.String("deserunt"),
                    NetAmount: 9204.81,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("et"),
                                ID: codataccounting.String("ae1ecf8c-3494-46bb-a7a0-5a8b4a9ec5b3"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("eum"),
                                ID: codataccounting.String("88cca363-2727-460e-966e-97e054103347"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("pariatur"),
                                ID: codataccounting.String("78ff2491-145f-4ab9-a59a-4af336664eaa"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "6bf2ff14-e8c1-4b35-aacc-edacc5227814",
                Name: codataccounting.String("Mr. Sherman Pfannerstill"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("cum"),
            PostedOn: codataccounting.String("quo"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("aliquam"),
                ID: codataccounting.String("1ea1342d-4104-4a25-af71-de57a11d614a"),
            },
            SourceModifiedDate: codataccounting.String("tempora"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dicta": map[string]interface{}{
                        "laboriosam": "sint",
                        "dolores": "repudiandae",
                    },
                },
            },
            UpdatedOn: codataccounting.String("fuga"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(284514),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```

## Delete

﻿> **Use with caution**
>
>Because Journal Entries underpin every transaction in an accounting platform, deleting a Journal Entry can affect every transaction for a given company.
> 
> **Before you proceed, make sure you understand the implications of deleting Journal Entries from an accounting perspective.**

The _Delete Journal entries_ endpoint allows you to delete a specified Journal entry from an accounting platform.

### Process
1. Pass the `{journalEntryId}` to the _Delete Journal Entries_ endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete by checking the status of push operation either via
   1. [Push operation webhook](/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
   2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation). 
   
   A `Success` status indicates that the Journal Entry object was deleted from the accounting platform.
3. (Optional) Check that the Journal Entry was deleted from the accounting platform.

### Effect on related objects

Be aware that deleting a Journal Entry from an accounting platform might cause related objects to be modified. For example, if you delete the Journal Entry for a paid invoice in QuickBooks Online, the invoice is deleted but the payment against that invoice is not. The payment is converted to a payment on account.

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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
        JournalEntryID: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## Get

Gets a single JournalEntry corresponding to the given ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
        JournalEntryID: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntry != nil {
        // handle response
    }
}
```

## GetCreateModel

﻿Get create journal entry model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries) for integrations that support creating journal entries.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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

## List

﻿Gets the latest journal entries for a company, with pagination.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
        Query: codataccounting.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntries != nil {
        // handle response
    }
}
```
