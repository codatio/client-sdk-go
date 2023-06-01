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
            CreatedOn: codataccounting.String("corporis"),
            Description: codataccounting.String("omnis"),
            ID: codataccounting.String("b855d015-b62c-48b8-ba38-a8a88c144200"),
            JournalLines: []shared.JournalLine{
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2caeb1ae-1ecf-48c3-8946-bba7a05a8b4a"),
                        Name: codataccounting.String("Noah Rutherford"),
                    },
                    Currency: codataccounting.String("amet"),
                    Description: codataccounting.String("eum"),
                    NetAmount: 5102.81,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("maxime"),
                                ID: codataccounting.String("ca363272-760e-4966-a97e-054103347d78"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("asperiores"),
                                ID: codataccounting.String("f2491145-fab9-4e59-a4af-336664eaa6bf"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dolores"),
                                ID: codataccounting.String("ff14e8c1-b352-4acc-adac-c5227814eca0"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("16bc41ea-1342-4d41-84a2-5ef71de57a11"),
                        Name: codataccounting.String("Franklin Brown"),
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
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5e208301-6ca3-44bb-87d4-f62127a607d1"),
                        Name: codataccounting.String("Betty Jacobi"),
                    },
                    Currency: codataccounting.String("quaerat"),
                    Description: codataccounting.String("nostrum"),
                    NetAmount: 1080.29,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("cumque"),
                                ID: codataccounting.String("3db9ca9f-38bd-42be-8787-03493f49aa84"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("ex"),
                                ID: codataccounting.String("5a328327-9b71-49d1-8ea6-73d86e3b35e4"),
                            },
                        },
                    },
                },
                shared.JournalLine{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9a313577-8ce5-44ca-8b0e-3ea975045bac"),
                        Name: codataccounting.String("Nathaniel DuBuque"),
                    },
                    Currency: codataccounting.String("quasi"),
                    Description: codataccounting.String("nemo"),
                    NetAmount: 1217.04,
                    Tracking: &shared.Propertiestracking2{
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("commodi"),
                                ID: codataccounting.String("ab5e3a02-2614-4315-9156-8299e61afc71"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("totam"),
                                ID: codataccounting.String("6ff20b7a-73df-440c-a0d7-657c1641bbf0"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("nostrum"),
                                ID: codataccounting.String("5271b251-1dd6-406d-91b2-8272bc9c3221"),
                            },
                        },
                    },
                },
            },
            JournalRef: &shared.JournalRef{
                ID: "697b1880-fcbb-42b9-bc15-f670bd178483",
                Name: codataccounting.String("Kristin Herman"),
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("necessitatibus"),
            PostedOn: codataccounting.String("harum"),
            RecordRef: &shared.InvoiceTo{
                DataType: codataccounting.String("amet"),
                ID: codataccounting.String("b6e241c3-1099-4836-a3c6-6dcbb7df6cb0"),
            },
            SourceModifiedDate: codataccounting.String("molestias"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "totam": map[string]interface{}{
                        "modi": "aperiam",
                        "praesentium": "recusandae",
                        "eaque": "nihil",
                    },
                    "dicta": map[string]interface{}{
                        "molestiae": "in",
                    },
                    "magnam": map[string]interface{}{
                        "saepe": "non",
                        "a": "voluptates",
                        "vero": "quae",
                        "doloremque": "et",
                    },
                    "possimus": map[string]interface{}{
                        "esse": "praesentium",
                        "aperiam": "laborum",
                        "dicta": "doloremque",
                    },
                },
            },
            UpdatedOn: codataccounting.String("minus"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(260242),
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
        JournalEntryID: "odio",
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
        JournalEntryID: "rerum",
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
        Query: codataccounting.String("provident"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JournalEntries != nil {
        // handle response
    }
}
```
