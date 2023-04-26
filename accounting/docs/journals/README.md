# Journals

## Overview

Journals

### Available Operations

* [GetCreateJournalsModel](#getcreatejournalsmodel) - Get create journal model
* [GetJournal](#getjournal) - Get journal
* [ListJournals](#listjournals) - List journals
* [PushJournal](#pushjournal) - Create journal

## GetCreateJournalsModel

Get create journal model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support creating journals.

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
    req := operations.GetCreateJournalsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Journals.GetCreateJournalsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetJournal

Gets a single journal corresponding to the given ID.

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
    req := operations.GetJournalRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        JournalID: "quibusdam",
    }

    res, err := s.Journals.GetJournal(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Journal != nil {
        // handle response
    }
}
```

## ListJournals

Gets the latest journals for a company, with pagination

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
    req := operations.ListJournalsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ab"),
    }

    res, err := s.Journals.ListJournals(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Journals != nil {
        // handle response
    }
}
```

## PushJournal

Posts a new journal to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create journal model](https://docs.codat.io/accounting-api#/operations/get-create-journals-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals) for integrations that support creating journals.

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
    req := operations.PushJournalRequest{
        Journal: &shared.Journal{
            CreatedOn: codataccounting.String("eligendi"),
            HasChildren: codataccounting.Bool(false),
            ID: codataccounting.String("ea673d86-e3b3-45e4-9a31-35778ce54cac"),
            JournalCode: codataccounting.String("libero"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("perferendis"),
            Name: codataccounting.String("Chris Terry"),
            ParentID: codataccounting.String("ducimus"),
            SourceModifiedDate: codataccounting.String("minima"),
            Status: shared.JournalStatusEnumUnknown.ToPointer(),
            Type: codataccounting.String("labore"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(359649),
    }

    res, err := s.Journals.PushJournal(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushJournal200ApplicationJSONObject != nil {
        // handle response
    }
}
```
