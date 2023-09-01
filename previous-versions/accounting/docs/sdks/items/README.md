# Items

## Overview

Items

### Available Operations

* [Create](#create) - Create item
* [Get](#get) - Get item
* [GetCreateModel](#getcreatemodel) - Get create item model
* [List](#list) - List items

## Create

The *Create item* endpoint creates a new [item](https://docs.codat.io/accounting-api#/schemas/Item) for a given company's connection.

[Items](https://docs.codat.io/accounting-api#/schemas/Item) allow your customers to save and track details of the products and services that they buy and sell.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create item model](https://docs.codat.io/accounting-api#/operations/get-create-items-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
                    ID: codataccounting.String("b7a73df4-0ca0-4d76-97c1-641bbf055271"),
                    Name: codataccounting.String("Mr. Alan Haley"),
                },
                Description: codataccounting.String("assumenda"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(3834.21),
                    ID: codataccounting.String("06dd1b28-272b-4c9c-b221-697b1880fcbb"),
                    Name: codataccounting.String("Verna Marks"),
                },
                UnitPrice: codataccounting.Float64(702.82),
            },
            Code: codataccounting.String("nostrum"),
            ID: codataccounting.String("f670bd17-8483-4165-beeb-3b6e241c3109"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("983663c6-6dcb-4b7d-b6cb-09c8b408e071"),
                    Name: codataccounting.String("Minnie Kirlin"),
                },
                Description: codataccounting.String("saepe"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(2505.06),
                    ID: codataccounting.String("fee101d9-780a-410c-87b9-5040d6c8b2a5"),
                    Name: codataccounting.String("John Aufderhar"),
                },
                UnitPrice: codataccounting.Float64(249.94),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusActive,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Name: codataccounting.String("Brenda Greenholt"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Type: shared.ItemTypeUnknown,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(4409),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateItemRequest](../../models/operations/createitemrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.CreateItemResponse](../../models/operations/createitemresponse.md), error**


## Get

The *Get item* endpoint returns a single item for a given itemId.

[Items](https://docs.codat.io/accounting-api#/schemas/Item) allow your customers to save and track details of the products and services that they buy and sell.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support getting a specific item.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        ItemID: "sit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Item != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetItemRequest](../../models/operations/getitemrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |


### Response

**[*operations.GetItemResponse](../../models/operations/getitemresponse.md), error**


## GetCreateModel

The *Get create item model* endpoint returns the expected data for the request payload when creating an [item](https://docs.codat.io/accounting-api#/schemas/Item) for a given company and integration.

[Items](https://docs.codat.io/accounting-api#/schemas/Item) allow your customers to save and track details of the products and services that they buy and sell.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating an item.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetCreateItemsModelRequest](../../models/operations/getcreateitemsmodelrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.GetCreateItemsModelResponse](../../models/operations/getcreateitemsmodelresponse.md), error**


## List

The *List items* endpoint returns a list of [items](https://docs.codat.io/accounting-api#/schemas/Item) for a given company's connection.

[Items](https://docs.codat.io/accounting-api#/schemas/Item) allow your customers to save and track details of the products and services that they buy and sell.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        Query: codataccounting.String("iste"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Items != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListItemsRequest](../../models/operations/listitemsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |


### Response

**[*operations.ListItemsResponse](../../models/operations/listitemsresponse.md), error**

