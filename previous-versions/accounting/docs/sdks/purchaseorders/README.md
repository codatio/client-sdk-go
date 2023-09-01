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
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(9054.35),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("5ffa906a-e559-4b72-ab67-46030fe18376"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("bedee767-90ed-40c1-aa7b-a478404489f6"),
                        Name: codataccounting.String("Courtney Bednar"),
                    },
                    Description: codataccounting.String("ipsa"),
                    DiscountAmount: codataccounting.Float64(3059.71),
                    DiscountPercentage: codataccounting.Float64(5075.68),
                    ItemRef: &shared.ItemRef{
                        ID: "091a2ba2-5ee6-4c75-af8a-60a7ae346e09",
                        Name: codataccounting.String("Misty Toy"),
                    },
                    Quantity: codataccounting.Float64(9569.24),
                    SubTotal: codataccounting.Float64(8931.29),
                    TaxAmount: codataccounting.Float64(4306.16),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(72.7),
                        ID: codataccounting.String("acaca645-de98-4675-91a9-cce61ec2c79a"),
                        Name: codataccounting.String("Miriam Pfannerstill"),
                    },
                    TotalAmount: codataccounting.Float64(6840.34),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d5a65b4d-5562-4d9b-bd9e-2d6fcf557629",
                            Name: codataccounting.String("Kerry Lockman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c3a89028-2a51-4f41-8f67-96ed3d724c18",
                            Name: codataccounting.String("Dr. Leon Ledner"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(5132.8),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("minus"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("consectetur"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("La Habra"),
                    Country: codataccounting.String("British Indian Ocean Territory (Chagos Archipelago)"),
                    Line1: codataccounting.String("autem"),
                    Line2: codataccounting.String("aliquid"),
                    PostalCode: codataccounting.String("08608"),
                    Region: codataccounting.String("dolor"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Giles.Bogan99@yahoo.com"),
                    Name: codataccounting.String("James Marvin"),
                    Phone: codataccounting.String("(373) 362-1857 x481"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusOpen.ToPointer(),
            SubTotal: codataccounting.Float64(958.62),
            SupplierRef: &shared.SupplierRef{
                ID: "0e1673d9-05cb-44be-9ef3-c127c3909955",
                SupplierName: codataccounting.String("consequuntur"),
            },
            TotalAmount: codataccounting.Float64(5021.78),
            TotalDiscount: codataccounting.Float64(1730.52),
            TotalTaxAmount: codataccounting.Float64(3712.26),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(57599),
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
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Get(ctx, operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "pariatur",
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
    res, err := s.PurchaseOrders.List(ctx, operations.ListPurchaseOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("maxime"),
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
            CurrencyRate: codataccounting.Float64(7439.49),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("3b121b88-c1ee-45e7-a061-391cc8fa0b7d"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("64926b0c-f5e6-4cb6-abab-e5e0b99f3b13"),
                        Name: codataccounting.String("Billie Stokes"),
                    },
                    Description: codataccounting.String("corrupti"),
                    DiscountAmount: codataccounting.Float64(4401.02),
                    DiscountPercentage: codataccounting.Float64(7240.6),
                    ItemRef: &shared.ItemRef{
                        ID: "b7aecbe5-69d7-40cb-8699-07f989441452",
                        Name: codataccounting.String("Mr. Perry Wiza"),
                    },
                    Quantity: codataccounting.Float64(2102.37),
                    SubTotal: codataccounting.Float64(3116.04),
                    TaxAmount: codataccounting.Float64(2769.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1795.24),
                        ID: codataccounting.String("c61be133-bacd-4e53-ab65-26f862853fe2"),
                        Name: codataccounting.String("Ronnie McKenzie"),
                    },
                    TotalAmount: codataccounting.Float64(2499.78),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2231fe66-64c4-41d2-bba5-cba069b8d291",
                            Name: codataccounting.String("Percy Renner Jr."),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6653.38),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2aa87494-79ed-4d4f-8f7b-50cf87f08f39"),
                        Name: codataccounting.String("Ms. Joy Blick"),
                    },
                    Description: codataccounting.String("culpa"),
                    DiscountAmount: codataccounting.Float64(1832.7),
                    DiscountPercentage: codataccounting.Float64(2964.52),
                    ItemRef: &shared.ItemRef{
                        ID: "b40c8f08-bff1-4081-a88f-86996c8e22be",
                        Name: codataccounting.String("Marguerite Emmerich"),
                    },
                    Quantity: codataccounting.Float64(2955.55),
                    SubTotal: codataccounting.Float64(4539.1),
                    TaxAmount: codataccounting.Float64(5001.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5720.7),
                        ID: codataccounting.String("3bd23f86-600c-461c-b834-273caa9118b3"),
                        Name: codataccounting.String("Emilio Bogisich"),
                    },
                    TotalAmount: codataccounting.Float64(810),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "331a54dc-1029-44f9-afed-939ba8f71e29",
                            Name: codataccounting.String("Adam Schroeder DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e1228ac3-adc6-447d-a40b-c11ea482824c",
                            Name: codataccounting.String("Gregg Jacobs"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f0f5b9d3-cb11-4a76-87d3-100e8e2b9b0d",
                            Name: codataccounting.String("Bertha Hudson"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6766.65),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("quisquam"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("fugiat"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Fort Mariannatown"),
                    Country: codataccounting.String("Turkey"),
                    Line1: codataccounting.String("quam"),
                    Line2: codataccounting.String("iste"),
                    PostalCode: codataccounting.String("65779-3904"),
                    Region: codataccounting.String("sint"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Flossie.Bailey93@yahoo.com"),
                    Name: codataccounting.String("Margaret Kutch"),
                    Phone: codataccounting.String("1-425-541-2767 x78041"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusClosed.ToPointer(),
            SubTotal: codataccounting.Float64(8436.48),
            SupplierRef: &shared.SupplierRef{
                ID: "61918d27-9c10-4c18-916f-d78be2621272",
                SupplierName: codataccounting.String("ea"),
            },
            TotalAmount: codataccounting.Float64(1382.52),
            TotalDiscount: codataccounting.Float64(5494.57),
            TotalTaxAmount: codataccounting.Float64(9854.77),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "officia",
        TimeoutInMinutes: codataccounting.Int(352126),
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

