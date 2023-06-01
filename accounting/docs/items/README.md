# Items

## Overview

Items

### Available Operations

* [Create](#create) - Create item
* [Get](#get) - Get item
* [GetCreateModel](#getcreatemodel) - Get create item model
* [List](#list) - List items

## Create

Posts a new item to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create item model](https://docs.codat.io/accounting-api#/operations/get-create-items-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) to see which integrations support this endpoint.

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
    res, err := s.Items.Create(ctx, operations.CreateItemRequest{
        Item: &shared.Item{
            BillItem: &shared.BillItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("e8ab4a9c-492c-45e8-ba5d-4aa4a508bd38"),
                    Name: codataccounting.String("Kara Cremin"),
                },
                Description: codataccounting.String("deserunt"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(5230.55),
                    ID: codataccounting.String("dd71bdda-a30b-47b9-9449-ae69c088d418"),
                    Name: codataccounting.String("Ms. Wm Kohler II"),
                },
                UnitPrice: codataccounting.Float64(9965.22),
            },
            Code: codataccounting.String("modi"),
            ID: codataccounting.String("23d54393-5f37-47ac-9c9b-7e93b6a3c523"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("105e7c34-cab0-4ecb-812a-66148944a8e9"),
                    Name: codataccounting.String("Ms. Jennie Hartmann"),
                },
                Description: codataccounting.String("nam"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(7904.63),
                    ID: codataccounting.String("25382533-43fb-40a4-a66e-a47578d171e2"),
                    Name: codataccounting.String("Frederick Bogan IV"),
                },
                UnitPrice: codataccounting.Float64(9823.15),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusArchived,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("aliquid"),
            Name: codataccounting.String("Shelly Purdy"),
            SourceModifiedDate: codataccounting.String("quia"),
            Type: shared.ItemTypeService,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(125551),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateItemResponse != nil {
        // handle response
    }
}
```

## Get

Gets the specified item for a given company.

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
    res, err := s.Items.Get(ctx, operations.GetItemRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ItemID: "veniam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Item != nil {
        // handle response
    }
}
```

## GetCreateModel

Get create item model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating items.

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
    res, err := s.Items.GetCreateModel(ctx, operations.GetCreateItemsModelRequest{
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

Gets the items for a given company.

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
    res, err := s.Items.List(ctx, operations.ListItemsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dolorem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Items != nil {
        // handle response
    }
}
```
