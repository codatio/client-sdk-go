# PurchaseOrders

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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Create(ctx, operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(5759.77),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("5075bc25-3825-4334-bfb0-a4e66ea47578"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("71e29418-18fc-4679-b6b2-f25359b855d0"),
                        Name: codataccounting.String("Cathy Powlowski"),
                    },
                    Description: codataccounting.String("eligendi"),
                    DiscountAmount: codataccounting.Float64(5505.76),
                    DiscountPercentage: codataccounting.Float64(7374.84),
                    ItemRef: &shared.ItemRef{
                        ID: "83a38a8a-88c1-4442-80c2-caeb1ae1ecf8",
                        Name: codataccounting.String("Chad Gislason"),
                    },
                    Quantity: codataccounting.Float64(4159.24),
                    SubTotal: codataccounting.Float64(6882.17),
                    TaxAmount: codataccounting.Float64(6973.55),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6797.33),
                        ID: codataccounting.String("7a05a8b4-a9ec-45b3-a88c-ca363272760e"),
                        Name: codataccounting.String("Chester Johnson"),
                    },
                    TotalAmount: codataccounting.Float64(4684.82),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "05410334-7d78-4ff2-8911-45fab9e59a4a",
                            Name: codataccounting.String("Nathan Fay"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "64eaa6bf-2ff1-44e8-81b3-52accedacc52",
                            Name: codataccounting.String("Mrs. Charlene Little"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ca016bc4-1ea1-4342-9410-4a25ef71de57",
                            Name: codataccounting.String("Keith Bradtke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "14a43176-92ea-4486-b3d5-22b828a90306",
                            Name: codataccounting.String("Mr. Michelle Zemlak"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(7532.31),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("sint"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("ut"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Runteview"),
                    Country: codataccounting.String("French Guiana"),
                    Line1: codataccounting.String("eligendi"),
                    Line2: codataccounting.String("magni"),
                    PostalCode: codataccounting.String("16117-2556"),
                    Region: codataccounting.String("facere"),
                    Type: shared.AddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Cathrine_Wisoky@yahoo.com"),
                    Name: codataccounting.String("Jon Heller"),
                    Phone: codataccounting.String("1-491-252-0038 x612"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusClosed.ToPointer(),
            SubTotal: codataccounting.Float64(5615.97),
            SupplierRef: &shared.SupplierRef{
                ID: "7d4f6212-7a60-47d1-a062-94514c3db9ca",
                SupplierName: codataccounting.String("omnis"),
            },
            TotalAmount: codataccounting.Float64(9600.37),
            TotalDiscount: codataccounting.Float64(2360.34),
            TotalTaxAmount: codataccounting.Float64(5514.18),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(694759),
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Get(ctx, operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "at",
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.List(ctx, operations.ListPurchaseOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("quia"),
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
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Update(ctx, operations.UpdatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(8782.83),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("8703493f-49aa-4846-9a32-83279b719d1c"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("673d86e3-b35e-449a-b135-778ce54cacb0"),
                        Name: codataccounting.String("Chris Terry"),
                    },
                    Description: codataccounting.String("ducimus"),
                    DiscountAmount: codataccounting.Float64(3264.82),
                    DiscountPercentage: codataccounting.Float64(276.36),
                    ItemRef: &shared.ItemRef{
                        ID: "45bacf63-b215-4186-ab5e-3a022614315d",
                        Name: codataccounting.String("Bernice Jakubowski"),
                    },
                    Quantity: codataccounting.Float64(5766.94),
                    SubTotal: codataccounting.Float64(5813.68),
                    TaxAmount: codataccounting.Float64(9207.5),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3778.77),
                        ID: codataccounting.String("1afc7186-ff20-4b7a-b3df-40ca0d7657c1"),
                        Name: codataccounting.String("Esther Bernier"),
                    },
                    TotalAmount: codataccounting.Float64(9823),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "55271b25-11dd-4606-9d1b-28272bc9c322",
                            Name: codataccounting.String("Tara Marquardt"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(822.04),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("880fcbb2-b93c-415f-a70b-d1784831653e"),
                        Name: codataccounting.String("Kerry Farrell"),
                    },
                    Description: codataccounting.String("vero"),
                    DiscountAmount: codataccounting.Float64(1407.83),
                    DiscountPercentage: codataccounting.Float64(3102.85),
                    ItemRef: &shared.ItemRef{
                        ID: "1c310998-3663-4c66-9cbb-7df6cb09c8b4",
                        Name: codataccounting.String("Ms. Irma Walter I"),
                    },
                    Quantity: codataccounting.Float64(4775.18),
                    SubTotal: codataccounting.Float64(4482.26),
                    TaxAmount: codataccounting.Float64(2986.6),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8337.81),
                        ID: codataccounting.String("e4fee101-d978-40a1-8c47-b95040d6c8b2"),
                        Name: codataccounting.String("Mr. Roberto Wilkinson"),
                    },
                    TotalAmount: codataccounting.Float64(1526.43),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "7e4048f9-0009-4ed2-9027-8eb4ae9d6416",
                            Name: codataccounting.String("Mrs. Delia Moore Jr."),
                        },
                    },
                    UnitAmount: codataccounting.Float64(1887.19),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("23b2c09b-9247-471f-9669-e5b7ec762664"),
                        Name: codataccounting.String("Woodrow Littel"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(5923.78),
                    DiscountPercentage: codataccounting.Float64(9275.83),
                    ItemRef: &shared.ItemRef{
                        ID: "4cfd2276-e0b8-48fb-87d6-fa5b6e8dbf81",
                        Name: codataccounting.String("Nettie Kuvalis"),
                    },
                    Quantity: codataccounting.Float64(809.04),
                    SubTotal: codataccounting.Float64(7530.97),
                    TaxAmount: codataccounting.Float64(6568.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4189.32),
                        ID: codataccounting.String("a9ffc561-929c-4ca9-960a-1395918da1d4"),
                        Name: codataccounting.String("Rolando Kovacek"),
                    },
                    TotalAmount: codataccounting.Float64(2178.8),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f8e1143d-a930-48b2-ba08-af22184439b3",
                            Name: codataccounting.String("Rolando Lesch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6ccce470-cd21-447b-ae61-52cf01d0d8c3",
                            Name: codataccounting.String("Clifford Quigley"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5bf935df-e974-4fa4-b1e9-c097eda62344",
                            Name: codataccounting.String("Alexis Braun"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "237e9984-c80b-4479-a891-923c18ca8d69",
                            Name: codataccounting.String("Jon Jacobs"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(1303.91),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("tempora"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("animi"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("North Caylatown"),
                    Country: codataccounting.String("Japan"),
                    Line1: codataccounting.String("officiis"),
                    Line2: codataccounting.String("eius"),
                    PostalCode: codataccounting.String("69025-7849"),
                    Region: codataccounting.String("ab"),
                    Type: shared.AddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Cicero_Schowalter67@hotmail.com"),
                    Name: codataccounting.String("Cody Zemlak"),
                    Phone: codataccounting.String("876-479-9089 x35968"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusClosed.ToPointer(),
            SubTotal: codataccounting.Float64(1625.48),
            SupplierRef: &shared.SupplierRef{
                ID: "ee70be06-9fb3-46ad-9704-080e0a3fc73a",
                SupplierName: codataccounting.String("ullam"),
            },
            TotalAmount: codataccounting.Float64(6283.94),
            TotalDiscount: codataccounting.Float64(465.74),
            TotalTaxAmount: codataccounting.Float64(2391.23),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "aliquam",
        TimeoutInMinutes: codataccounting.Int(744576),
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

