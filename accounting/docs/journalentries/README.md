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
            CreatedOn: codataccounting.String("delectus"),
            Description: codataccounting.String("id"),
            ID: codataccounting.String("1011a091-b3ec-48b5-b862-de1a9d14fe72"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("521f9030-3dfc-4338-b97f-ffa6d1d32090"),
                        Name: codataccounting.String("Salvatore Boyer"),
                    },
                    Currency: codataccounting.String("mollitia"),
                    Description: codataccounting.String("cumque"),
                    NetAmount: 5632.6,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("accusamus"),
                                ID: codataccounting.String("1961ce9b-e41c-4869-9d7d-9719d07b200a"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("corporis"),
                                ID: codataccounting.String("8ffd2967-df8f-4d88-aa8e-60be620cd9c5"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("officia"),
                                ID: codataccounting.String("fdd04c37-5251-42be-ae1d-87ecc5fdcea8"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("eveniet"),
                                ID: codataccounting.String("7a883116-62cd-4a6d-b7c1-d86066237d42"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("27866db8-a749-4e39-8451-1cc75e4f0c00"),
                        Name: codataccounting.String("Patty Harber"),
                    },
                    Currency: codataccounting.String("molestiae"),
                    Description: codataccounting.String("ipsam"),
                    NetAmount: 5541.62,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("nobis"),
                                ID: codataccounting.String("94562f00-6968-45fc-91a1-73d84bbe24f2"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("error"),
                                ID: codataccounting.String("834afb07-35cb-4628-9d4a-29aaa1e16915"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("nisi"),
                                ID: codataccounting.String("f7d2ee20-9505-4bf0-ba93-e94480ca37fb"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("ab"),
                                ID: codataccounting.String("0789032a-c333-4172-a2dd-79ec74ba7e88"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ddb36fd1-ccc3-441c-8657-3474f0a740fb"),
                        Name: codataccounting.String("Sandy Reichel"),
                    },
                    Currency: codataccounting.String("illo"),
                    Description: codataccounting.String("impedit"),
                    NetAmount: 2164.48,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("doloremque"),
                                ID: codataccounting.String("9e763995-d808-4bbe-b944-55ebc550a1c4"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("qui"),
                                ID: codataccounting.String("6b59c836-6fdc-4c13-9582-c1b855e889d9"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("officiis"),
                                ID: codataccounting.String("f932e900-0a13-4ad8-9242-08efd2341189"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8e73879e-fbe8-4bae-babb-794536e90351"),
                        Name: codataccounting.String("Rickey Miller"),
                    },
                    Currency: codataccounting.String("adipisci"),
                    Description: codataccounting.String("architecto"),
                    NetAmount: 4393.34,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("voluptatem"),
                                ID: codataccounting.String("b77a5a53-65a7-49f1-9271-f01c0d361fed"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "8dc5effb-453e-4908-9e87-1fdb4d697bdd",
                Name: codataccounting.String("Sylvester Maggio"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("repudiandae"),
            PostedOn: codataccounting.String("incidunt"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("neque"),
                ID: codataccounting.String("734a5d72-d9ed-4d78-9be5-e7afe55297ba"),
            },
            SourceModifiedDate: codataccounting.String("laboriosam"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "laudantium": map[string]interface{}{
                        "repellat": "aliquam",
                    },
                },
            },
            UpdatedOn: codataccounting.String("modi"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(907650),
    }

    res, err := s.JournalEntries.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```

## Delete

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

    res, err := s.JournalEntries.Delete(ctx, req)
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetJournalEntryRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalEntryID: "dolorem",
    }

    res, err := s.JournalEntries.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntry != nil {
        // handle response
    }
}
```

## GetCreateModel

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

    res, err := s.JournalEntries.GetCreateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

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
        Query: codataccounting.String("laborum"),
    }

    res, err := s.JournalEntries.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntries != nil {
        // handle response
    }
}
```
