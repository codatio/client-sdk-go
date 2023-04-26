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

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating items.

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
    req := operations.CreateItemRequest{
        Item: &shared.Item{
            BillItem: &shared.BillItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("5472cdd1-4fc4-43b7-8bca-88fa70c43351"),
                    Name: codataccounting.String("Luis Swaniawski PhD"),
                },
                Description: codataccounting.String("harum"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(5331.06),
                    ID: codataccounting.String("f7f75f4f-23f1-4c0a-986c-3ae7d7b67fee"),
                    Name: codataccounting.String("Mrs. Floyd Torphy"),
                },
                UnitPrice: codataccounting.Float64(8642.28),
            },
            Code: codataccounting.String("perspiciatis"),
            ID: codataccounting.String("5b1dbece-ff7c-44b1-96e9-278275eea768"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("17468063-f799-4b79-96c0-b0fa0bb20a40"),
                    Name: codataccounting.String("Roland Ryan"),
                },
                Description: codataccounting.String("accusamus"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(3996.96),
                    ID: codataccounting.String("40642726-57b0-41a0-bc08-fd3921c25793"),
                    Name: codataccounting.String("Mercedes Kemmer V"),
                },
                UnitPrice: codataccounting.Float64(2095.62),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusEnumActive,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("sequi"),
            Name: codataccounting.String("Dominick Pagac"),
            SourceModifiedDate: codataccounting.String("temporibus"),
            Type: "adipisci",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(420757),
    }

    res, err := s.Items.Create(ctx, req)
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetItemRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ItemID: "ea",
    }

    res, err := s.Items.Get(ctx, req)
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating items.

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
    req := operations.GetCreateItemsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Items.GetCreateModel(ctx, req)
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListItemsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("nulla"),
    }

    res, err := s.Items.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Items != nil {
        // handle response
    }
}
```
