# PurchaseOrders
(*PurchaseOrders*)

## Overview

Purchase orders

### Available Operations

* [Create](#create) - Create purchase order
* [Get](#get) - Get purchase order
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update purchase order model
* [List](#list) - List purchase orders
* [Update](#update) - Update purchase order

## Create

The *Create purchase order* endpoint creates a new [purchase order](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) for a given company's connection.

[Purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) represent a business's intent to purchase goods or services from a supplier.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update purchase order model](https://docs.codat.io/accounting-api#/operations/get-create-update-purchaseOrders-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Create(ctx, operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: accounting.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("4893.82"),
            DeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ID: accounting.String("<ID>"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("South"),
                    },
                    Description: accounting.String("Vision-oriented responsive function"),
                    DiscountAmount: types.MustNewDecimalFromString("9510.62"),
                    DiscountPercentage: types.MustNewDecimalFromString("8915.1"),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("deposit"),
                    },
                    Quantity: types.MustNewDecimalFromString("3015.1"),
                    SubTotal: types.MustNewDecimalFromString("899.64"),
                    TaxAmount: types.MustNewDecimalFromString("7150.4"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("7926.2"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("Gasoline Screen mobile"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6562.56"),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("Durham after"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("5190.28"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("Fish"),
            PaymentDueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: accounting.String("functionalities Grocery Borders"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: accounting.String("West Mackville"),
                    Country: accounting.String("Greenland"),
                    Line1: accounting.String("Southfield Interactions Senior"),
                    Line2: accounting.String("red"),
                    PostalCode: accounting.String("53026-6579"),
                    Region: accounting.String("payment 1080p"),
                    Type: shared.AccountingAddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: accounting.String("Otho27@yahoo.com"),
                    Name: accounting.String("Toyota Neptunium round"),
                    Phone: accounting.String("1-394-849-5203"),
                },
            },
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusDraft.ToPointer(),
            SubTotal: types.MustNewDecimalFromString("4904.2"),
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: accounting.String("markets"),
            },
            TotalAmount: types.MustNewDecimalFromString("9244.84"),
            TotalDiscount: types.MustNewDecimalFromString("9824.5"),
            TotalTaxAmount: types.MustNewDecimalFromString("454.16"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: accounting.Int(589256),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreatePurchaseOrderRequest](../../models/operations/createpurchaseorderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.CreatePurchaseOrderResponse](../../models/operations/createpurchaseorderresponse.md), error**


## Get

The *Get purchase order* endpoint returns a single purchase order for a given purchaseOrderId.

[Purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) represent a business's intent to purchase goods or services from a supplier.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support getting a specific purchase order.

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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Get(ctx, operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "Northeast Hatchback Kia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPurchaseOrderRequest](../../models/operations/getpurchaseorderrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetPurchaseOrderResponse](../../models/operations/getpurchaseorderresponse.md), error**


## GetCreateUpdateModel

The *Get create/update purchase order model* endpoint returns the expected data for the request payload when creating and updating a [purchase order](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) for a given company and integration.

[Purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) represent a business's intent to purchase goods or services from a supplier.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating and updating a purchase order.


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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.GetCreateUpdateModel(ctx, operations.GetCreateUpdatePurchaseOrdersModelRequest{
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetCreateUpdatePurchaseOrdersModelRequest](../../models/operations/getcreateupdatepurchaseordersmodelrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.GetCreateUpdatePurchaseOrdersModelResponse](../../models/operations/getcreateupdatepurchaseordersmodelresponse.md), error**


## List

The *List purchase orders* endpoint returns a list of [purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) for a given company's connection.

[Purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) represent a business's intent to purchase goods or services from a supplier.

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
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.List(ctx, operations.ListPurchaseOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: accounting.String("-modifiedDate"),
        Page: accounting.Int(1),
        PageSize: accounting.Int(100),
        Query: accounting.String("Northeast Metal Canada"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPurchaseOrdersRequest](../../models/operations/listpurchaseordersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.ListPurchaseOrdersResponse](../../models/operations/listpurchaseordersresponse.md), error**


## Update

The *Update purchase order* endpoint updates an existing [purchase order](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) for a given company's connection.

[Purchase orders](https://docs.codat.io/accounting-api#/schemas/PurchaseOrder) represent a business's intent to purchase goods or services from a supplier.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update purchase order model](https://docs.codat.io/accounting-api#/operations/get-create-update-purchaseOrders-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Update(ctx, operations.UpdatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: accounting.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("245.55"),
            DeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ID: accounting.String("<ID>"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("dock Quality redundant"),
                    },
                    Description: accounting.String("Visionary bi-directional analyzer"),
                    DiscountAmount: types.MustNewDecimalFromString("2782.81"),
                    DiscountPercentage: types.MustNewDecimalFromString("8965.01"),
                    ItemRef: &shared.ItemRef{
                        ID: "<ID>",
                        Name: accounting.String("withdrawal extend"),
                    },
                    Quantity: types.MustNewDecimalFromString("2494.4"),
                    SubTotal: types.MustNewDecimalFromString("3668.07"),
                    TaxAmount: types.MustNewDecimalFromString("1395.79"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("6447.13"),
                        ID: accounting.String("<ID>"),
                        Name: accounting.String("syndicate East Baht"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("6298.17"),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "<ID>",
                            Name: accounting.String("guestbook driver users"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("3567.52"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("Lev Wooden"),
            PaymentDueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: accounting.String("Internal invoice"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: accounting.String("Fort Ari"),
                    Country: accounting.String("Poland"),
                    Line1: accounting.String("compressing convergence"),
                    Line2: accounting.String("modulo Kia"),
                    PostalCode: accounting.String("81011"),
                    Region: accounting.String("Reactive Global Northeast"),
                    Type: shared.AccountingAddressTypeBilling,
                },
                Contact: &shared.ShipToContact{
                    Email: accounting.String("Abe.Bogan@hotmail.com"),
                    Name: accounting.String("quisquam"),
                    Phone: accounting.String("(473) 342-1764 x99443"),
                },
            },
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusDraft.ToPointer(),
            SubTotal: types.MustNewDecimalFromString("3635.65"),
            SupplierRef: &shared.SupplierRef{
                ID: "<ID>",
                SupplierName: accounting.String("squat Omnigender"),
            },
            TotalAmount: types.MustNewDecimalFromString("8831.22"),
            TotalDiscount: types.MustNewDecimalFromString("8283.54"),
            TotalTaxAmount: types.MustNewDecimalFromString("9620.25"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: accounting.Bool(false),
        PurchaseOrderID: "aggregate Benz Granite",
        TimeoutInMinutes: accounting.Int(265006),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePurchaseOrderRequest](../../models/operations/updatepurchaseorderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.UpdatePurchaseOrderResponse](../../models/operations/updatepurchaseorderresponse.md), error**

