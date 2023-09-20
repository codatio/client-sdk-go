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
            Currency: accounting.String("EUR"),
            CurrencyRate: types.MustNewDecimalFromString("1356.92"),
            DeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ID: accounting.String("309db053-6d9e-475c-a006-f5392c11a25a"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("bf92f974-28ad-49a9-b8bf-8221125359d9"),
                        Name: accounting.String("Glenn Lebsack"),
                    },
                    Description: accounting.String("iusto"),
                    DiscountAmount: types.MustNewDecimalFromString("6339.56"),
                    DiscountPercentage: types.MustNewDecimalFromString("4531.11"),
                    ItemRef: &shared.ItemRef{
                        ID: "9cd72cd2-484d-4a21-b29f-2ac41ef5725f",
                        Name: accounting.String("Jean Kemmer"),
                    },
                    Quantity: types.MustNewDecimalFromString("7500.75"),
                    SubTotal: types.MustNewDecimalFromString("944.87"),
                    TaxAmount: types.MustNewDecimalFromString("9313.93"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("3062.69"),
                        ID: accounting.String("1d8a23c2-3e34-4f2d-ba4a-197f6de92215"),
                        Name: accounting.String("Ms. Dixie Turner Sr."),
                    },
                    TotalAmount: types.MustNewDecimalFromString("309.62"),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "99853e9f-543d-4854-839e-e224460443bc",
                            Name: accounting.String("Ms. Lorraine Gusikowski"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("7911.29"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("asperiores"),
            PaymentDueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: accounting.String("vel"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: accounting.String("Madison"),
                    Country: accounting.String("Guyana"),
                    Line1: accounting.String("at"),
                    Line2: accounting.String("culpa"),
                    PostalCode: accounting.String("52186"),
                    Region: accounting.String("libero"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: accounting.String("Augustine.Kuhlman69@yahoo.com"),
                    Name: accounting.String("Miss Meredith Hand"),
                    Phone: accounting.String("790-276-8540 x482"),
                },
            },
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusUnknown.ToPointer(),
            SubTotal: types.MustNewDecimalFromString("5142.34"),
            SupplierRef: &shared.SupplierRef{
                ID: "2bfbdc41-ff5d-44e2-ae4f-b5cb35d17638",
                SupplierName: accounting.String("delectus"),
            },
            TotalAmount: types.MustNewDecimalFromString("1230.16"),
            TotalDiscount: types.MustNewDecimalFromString("9225.93"),
            TotalTaxAmount: types.MustNewDecimalFromString("8182.73"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: accounting.Int(716538),
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
        PurchaseOrderID: "ducimus",
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
        Query: accounting.String("atque"),
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
            Currency: accounting.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("3670.96"),
            DeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: accounting.String("2022-10-23T00:00:00.000Z"),
            ID: accounting.String("cc5cb860-f8cd-4580-ba73-810e4fe44472"),
            IssueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: accounting.String("7cd3b1dd-3bbc-4e24-bb76-84eff50126d7"),
                        Name: accounting.String("Vicky Wolf"),
                    },
                    Description: accounting.String("facere"),
                    DiscountAmount: types.MustNewDecimalFromString("113.92"),
                    DiscountPercentage: types.MustNewDecimalFromString("9309.42"),
                    ItemRef: &shared.ItemRef{
                        ID: "b74b8421-953b-444b-93c4-3159d33e5953",
                        Name: accounting.String("Mr. Michael Bashirian"),
                    },
                    Quantity: types.MustNewDecimalFromString("5784.52"),
                    SubTotal: types.MustNewDecimalFromString("5525.12"),
                    TaxAmount: types.MustNewDecimalFromString("3971.12"),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: types.MustNewDecimalFromString("1930.48"),
                        ID: accounting.String("aa41e6c3-1cc2-4f1f-8b51-c9a41ffbe9cb"),
                        Name: accounting.String("Brad Mraz"),
                    },
                    TotalAmount: types.MustNewDecimalFromString("9358"),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "65e076cc-7abf-4616-aa5c-71641934b90f",
                            Name: accounting.String("Rochelle Bailey"),
                        },
                    },
                    UnitAmount: types.MustNewDecimalFromString("683"),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Note: accounting.String("possimus"),
            PaymentDueDate: accounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: accounting.String("a"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: accounting.String("Champlinworth"),
                    Country: accounting.String("Nepal"),
                    Line1: accounting.String("earum"),
                    Line2: accounting.String("fugit"),
                    PostalCode: accounting.String("10352-2752"),
                    Region: accounting.String("enim"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: accounting.String("Dayana.Kilback@gmail.com"),
                    Name: accounting.String("Ms. Nettie McCullough"),
                    Phone: accounting.String("(836) 983-6977 x4324"),
                },
            },
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusVoid.ToPointer(),
            SubTotal: types.MustNewDecimalFromString("5678.21"),
            SupplierRef: &shared.SupplierRef{
                ID: "222c9ff5-7491-4aab-ba2e-761f0ca4d456",
                SupplierName: accounting.String("vero"),
            },
            TotalAmount: types.MustNewDecimalFromString("9708.48"),
            TotalDiscount: types.MustNewDecimalFromString("1170.53"),
            TotalTaxAmount: types.MustNewDecimalFromString("234.1"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: accounting.Bool(false),
        PurchaseOrderID: "adipisci",
        TimeoutInMinutes: accounting.Int(80294),
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

