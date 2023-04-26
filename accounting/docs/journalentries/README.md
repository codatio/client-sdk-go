# JournalEntries

## Overview

Journal entries

### Available Operations

* [CreateJournalEntry](#createjournalentry) - Create journal entry
* [DeleteJournalEntry](#deletejournalentry) - Delete journal entry
* [GetCreateJournalEntriesModel](#getcreatejournalentriesmodel) - Get create journal entry model
* [GetJournalEntry](#getjournalentry) - Get journal entry
* [ListJournalEntries](#listjournalentries) - List journal entries

## CreateJournalEntry

Posts a new journalEntry to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create journal entry model](https://docs.codat.io/accounting-api#/operations/get-create-journalEntries-model).

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
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateJournalEntryRequest{
        JournalEntry: &shared.JournalEntry{
            CreatedOn: codataccounting.String("tempora"),
            Description: codataccounting.String("nesciunt"),
            ID: codataccounting.String("fb0a4e66-ea47-4578-9171-e2941818fc67"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b6b2f253-59b8-455d-815b-62c8b83a38a8"),
                        Name: codataccounting.String("Dwayne MacGyver I"),
                    },
                    Currency: codataccounting.String("labore"),
                    Description: codataccounting.String("consequuntur"),
                    NetAmount: 317.03,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("optio"),
                                ID: codataccounting.String("2caeb1ae-1ecf-48c3-8946-bba7a05a8b4a"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9ec5b368-8cca-4363-a727-60e966e97e05"),
                        Name: codataccounting.String("Teresa Anderson"),
                    },
                    Currency: codataccounting.String("aliquam"),
                    Description: codataccounting.String("esse"),
                    NetAmount: 8634.7,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("corrupti"),
                                ID: codataccounting.String("ff249114-5fab-49e5-9a4a-f336664eaa6b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("sapiente"),
                                ID: codataccounting.String("2ff14e8c-1b35-42ac-8eda-cc5227814eca"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("016bc41e-a134-42d4-904a-25ef71de57a1"),
                        Name: codataccounting.String("Mrs. Janis Keeling"),
                    },
                    Currency: codataccounting.String("tempora"),
                    Description: codataccounting.String("velit"),
                    NetAmount: 1191.73,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("laboriosam"),
                                ID: codataccounting.String("92ea4867-3d52-42b8-a8a9-030660f024c7"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("sint"),
                                ID: codataccounting.String("b4cc64c2-b3a3-42c4-88ad-e62f6aa558a6"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "5e208301-6ca3-44bb-87d4-f62127a607d1",
                Name: codataccounting.String("Betty Jacobi"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("quaerat"),
            PostedOn: codataccounting.String("nostrum"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("beatae"),
                ID: codataccounting.String("4c3db9ca-9f38-4bd2-be87-8703493f49aa"),
            },
            SourceModifiedDate: codataccounting.String("laudantium"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ex": map[string]interface{}{
                        "mollitia": "sequi",
                        "eos": "laudantium",
                    },
                    "adipisci": map[string]interface{}{
                        "iusto": "natus",
                    },
                },
            },
            UpdatedOn: codataccounting.String("facilis"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(465310),
    }

    res, err := s.JournalEntries.CreateJournalEntry(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```

## DeleteJournalEntry

Deletes a journal entry from the accounting package for a given company.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DeleteJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        JournalEntryID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.JournalEntries.DeleteJournalEntry(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## GetCreateJournalEntriesModel

Get create journal entry model. Returns the expected data for the request payload.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateJournalEntriesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.JournalEntries.GetCreateJournalEntriesModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetJournalEntry

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalEntryID: "beatae",
    }

    res, err := s.JournalEntries.GetJournalEntry(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntry != nil {
        // handle response
    }
}
```

## ListJournalEntries

Gets the latest journal entries for a company, with pagination

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListJournalEntriesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("error"),
    }

    res, err := s.JournalEntries.ListJournalEntries(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntries != nil {
        // handle response
    }
}
```
