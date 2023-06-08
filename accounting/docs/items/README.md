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
                    ID: codataccounting.String("d64161e9-1500-4323-b2c0-9b924771f566"),
                    Name: codataccounting.String("Pat Hickle"),
                },
                Description: codataccounting.String("saepe"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(7828.51),
                    ID: codataccounting.String("7626649d-84eb-49e4-8fd2-276e0b88fb87"),
                    Name: codataccounting.String("Gilbert Zboncak"),
                },
                UnitPrice: codataccounting.Float64(7370.08),
            },
            Code: codataccounting.String("vel"),
            ID: codataccounting.String("e8dbf812-f83b-41ca-aa9f-fc561929cca9"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("560a1395-918d-4a1d-88e7-8e3cf8e1143d"),
                    Name: codataccounting.String("Ms. Wendell Dooley"),
                },
                Description: codataccounting.String("magni"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(4682.52),
                    ID: codataccounting.String("a08af221-8443-49b3-9e87-56ccce470cd2"),
                    Name: codataccounting.String("Elaine Kihn"),
                },
                UnitPrice: codataccounting.Float64(9267.48),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusActive,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Name: codataccounting.String("Nicole Schowalter Sr."),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Type: shared.ItemTypeUnknown,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(873982),
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
        ItemID: "voluptatum",
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
        Query: codataccounting.String("eligendi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Items != nil {
        // handle response
    }
}
```
