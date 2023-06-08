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
                    ID: codataccounting.String("8e071377-4de4-4fee-901d-9780a10c47b9"),
                    Name: codataccounting.String("Dr. Lisa Haag"),
                },
                Description: codataccounting.String("minus"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(5048.47),
                    ID: codataccounting.String("b2a5f002-207e-4404-8f90-009ed290278e"),
                    Name: codataccounting.String("Tom Paucek"),
                },
                UnitPrice: codataccounting.Float64(8652.77),
            },
            Code: codataccounting.String("suscipit"),
            ID: codataccounting.String("4161e915-0032-43b2-809b-924771f5669e"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("5b7ec762-6649-4d84-ab9e-4cfd2276e0b8"),
                    Name: codataccounting.String("Edmond Pfeffer"),
                },
                Description: codataccounting.String("nulla"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(3968.06),
                    ID: codataccounting.String("fa5b6e8d-bf81-42f8-bb1c-a6a9ffc56192"),
                    Name: codataccounting.String("Simon Runolfsson"),
                },
                UnitPrice: codataccounting.Float64(3666.09),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusActive,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Name: codataccounting.String("Ryan Fisher"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Type: shared.ItemTypeUnknown,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(558963),
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
        ItemID: "facere",
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
        Query: codataccounting.String("fuga"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Items != nil {
        // handle response
    }
}
```
