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
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(7740.9),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("a645de98-6755-41a9-8ce6-1ec2c79a39ae"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4d5a65b4-d556-42d9-b7d9-e2d6fcf55762"),
                        Name: codataccounting.String("Drew Pollich"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(7877.69),
                    DiscountPercentage: codataccounting.Float64(1972.67),
                    ItemRef: &shared.ItemRef{
                        ID: "a890282a-51f4-41cf-a796-ed3d724c18f5",
                        Name: codataccounting.String("Jonathan Tromp"),
                    },
                    Quantity: codataccounting.Float64(7803.32),
                    SubTotal: codataccounting.Float64(7948.94),
                    TaxAmount: codataccounting.Float64(8856.75),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2355.79),
                        ID: codataccounting.String("f716600d-a0e3-4aa6-9c6f-e09d852b53b3"),
                        Name: codataccounting.String("Nichole Kutch"),
                    },
                    TotalAmount: codataccounting.Float64(8112.21),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c710e167-3d90-45cb-8bed-ef3c127c3909",
                            Name: codataccounting.String("Alvin Hartmann"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(1730.52),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("50dcbbcd-3b12-41b8-8c1e-e5e7a061391c"),
                        Name: codataccounting.String("Dwight Wisoky DDS"),
                    },
                    Description: codataccounting.String("iusto"),
                    DiscountAmount: codataccounting.Float64(8302.89),
                    DiscountPercentage: codataccounting.Float64(1060.35),
                    ItemRef: &shared.ItemRef{
                        ID: "764926b0-cf5e-46cb-aeba-be5e0b99f3b1",
                        Name: codataccounting.String("Veronica Luettgen"),
                    },
                    Quantity: codataccounting.Float64(6429.43),
                    SubTotal: codataccounting.Float64(5465.01),
                    TaxAmount: codataccounting.Float64(4401.02),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7240.6),
                        ID: codataccounting.String("b7aecbe5-69d7-40cb-8699-07f989441452"),
                        Name: codataccounting.String("Mr. Perry Wiza"),
                    },
                    TotalAmount: codataccounting.Float64(2102.37),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "42c61be1-33ba-4cde-932b-6526f862853f",
                            Name: codataccounting.String("Alan Langworth"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ce322231-fe66-464c-81d2-fba5cba069b8",
                            Name: codataccounting.String("Miss Billy McCullough"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(7299.76),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("810a2aa8-7494-479e-9d4f-cf7b50cf87f0"),
                        Name: codataccounting.String("Emmett Ernser"),
                    },
                    Description: codataccounting.String("esse"),
                    DiscountAmount: codataccounting.Float64(762.27),
                    DiscountPercentage: codataccounting.Float64(197.89),
                    ItemRef: &shared.ItemRef{
                        ID: "76a24b40-c8f0-48bf-b108-1e88f86996c8",
                        Name: codataccounting.String("Fred Dare"),
                    },
                    Quantity: codataccounting.Float64(164.13),
                    SubTotal: codataccounting.Float64(6264.24),
                    TaxAmount: codataccounting.Float64(2164.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8056.07),
                        ID: codataccounting.String("f47893bd-23f8-4660-8c61-c7834273caa9"),
                        Name: codataccounting.String("Anna Lemke"),
                    },
                    TotalAmount: codataccounting.Float64(5420.06),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1b61a331-a54d-4c10-a94f-92fed939ba8f",
                            Name: codataccounting.String("Janet Upton"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92c20ee1-228a-4c3a-9c64-7d240bc11ea4",
                            Name: codataccounting.String("Mrs. Nicholas Leannon"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cc6a2f0f-5b9d-43cb-91a7-687d3100e8e2",
                            Name: codataccounting.String("Dr. Salvador Roberts"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "46d2a7c7-d1ea-40e7-9fa9-bbe5f179f650",
                            Name: codataccounting.String("Arthur Wehner IV"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(9205.44),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("officiis"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("neque"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Haliemouth"),
                    Country: codataccounting.String("Brazil"),
                    Line1: codataccounting.String("non"),
                    Line2: codataccounting.String("soluta"),
                    PostalCode: codataccounting.String("77804-1678"),
                    Region: codataccounting.String("ex"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Blaise.Leffler44@yahoo.com"),
                    Name: codataccounting.String("Dr. Wilbert Buckridge IV"),
                    Phone: codataccounting.String("(239) 845-6914 x101"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusUnknown.ToPointer(),
            SubTotal: codataccounting.Float64(4110.09),
            SupplierRef: &shared.SupplierRef{
                ID: "28fa5039-6286-47e7-ab3a-65024b157f9b",
                SupplierName: codataccounting.String("nobis"),
            },
            TotalAmount: codataccounting.Float64(3859.93),
            TotalDiscount: codataccounting.Float64(9232.4),
            TotalTaxAmount: codataccounting.Float64(9998.39),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(437621),
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
        PurchaseOrderID: "aspernatur",
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
        Query: codataccounting.String("similique"),
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
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(600.68),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("1d99b661-a7de-4f16-8b6c-cb2822b4a985"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d2f4a1e9-c4ae-4551-80e7-5726e003c2f0"),
                        Name: codataccounting.String("Ms. Ada Haag"),
                    },
                    Description: codataccounting.String("enim"),
                    DiscountAmount: codataccounting.Float64(991.13),
                    DiscountPercentage: codataccounting.Float64(5431.77),
                    ItemRef: &shared.ItemRef{
                        ID: "cee41c99-9f46-49f6-b1cf-1a3f023c669e",
                        Name: codataccounting.String("Fannie Jerde"),
                    },
                    Quantity: codataccounting.Float64(324.74),
                    SubTotal: codataccounting.Float64(680.19),
                    TaxAmount: codataccounting.Float64(1256.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8751.44),
                        ID: codataccounting.String("ba057988-c672-40c3-903f-1a40c0f3ec86"),
                        Name: codataccounting.String("Clayton Yost"),
                    },
                    TotalAmount: codataccounting.Float64(8912.43),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6fc03128-f0aa-4aee-a004-eba7bf8732be",
                            Name: codataccounting.String("Linda McKenzie"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "087131f0-6f0b-4ce5-9a86-87143c97905f",
                            Name: codataccounting.String("Kurt Marvin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5da664b7-e778-4a74-baaa-2832bb65862d",
                            Name: codataccounting.String("Dr. Lee Frami"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b14aa6bd-ec7f-4444-a32e-9a5dee1acd72",
                            Name: codataccounting.String("Jaime McKenzie"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(711.51),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b58fe682-e1c2-4dbe-a3d5-8e8247d122c9"),
                        Name: codataccounting.String("Gilbert Kshlerin"),
                    },
                    Description: codataccounting.String("praesentium"),
                    DiscountAmount: codataccounting.Float64(9796.97),
                    DiscountPercentage: codataccounting.Float64(6607.28),
                    ItemRef: &shared.ItemRef{
                        ID: "27958367-363d-4a07-9096-faeb86480730",
                        Name: codataccounting.String("Armando Wiza"),
                    },
                    Quantity: codataccounting.Float64(5237.4),
                    SubTotal: codataccounting.Float64(6038.54),
                    TaxAmount: codataccounting.Float64(8675.31),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6062.34),
                        ID: codataccounting.String("ca607565-6fc0-4ebe-a715-5e2d06a3070d"),
                        Name: codataccounting.String("Sonja Christiansen"),
                    },
                    TotalAmount: codataccounting.Float64(9410.3),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "81fabaaa-7d80-4108-8076-ff5f6ed29814",
                            Name: codataccounting.String("Mrs. Brandy Lehner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b6a70b0d-d82f-494f-bfbd-1e1e21ddc690",
                            Name: codataccounting.String("Dr. Margie Powlowski IV"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(4675.79),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b51eb5fd-30bf-4e03-890c-f20254a95904"),
                        Name: codataccounting.String("Kendra Rau"),
                    },
                    Description: codataccounting.String("aspernatur"),
                    DiscountAmount: codataccounting.Float64(8202.86),
                    DiscountPercentage: codataccounting.Float64(3958.97),
                    ItemRef: &shared.ItemRef{
                        ID: "bc9917f9-8e47-492b-979a-413d6a8c9168",
                        Name: codataccounting.String("Alison Stroman"),
                    },
                    Quantity: codataccounting.Float64(992.09),
                    SubTotal: codataccounting.Float64(8469.9),
                    TaxAmount: codataccounting.Float64(606.59),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8453.97),
                        ID: codataccounting.String("98ccf89d-3861-4186-ad76-c002facb13ac"),
                        Name: codataccounting.String("Debbie Schowalter II"),
                    },
                    TotalAmount: codataccounting.Float64(1929.94),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "866c575a-1e26-4687-b0be-37b0e8fbc48d",
                            Name: codataccounting.String("Cedric Koch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9b535105-0501-44dc-a105-882484c36e94",
                            Name: codataccounting.String("Byron Mayert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "82d34e0b-8fc0-4d59-b57b-9f9820be0780",
                            Name: codataccounting.String("Leland Douglas"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(5874.85),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e2f70344-e00f-4478-ab53-9483f748eefc"),
                        Name: codataccounting.String("Kelvin Keebler"),
                    },
                    Description: codataccounting.String("minima"),
                    DiscountAmount: codataccounting.Float64(2545.32),
                    DiscountPercentage: codataccounting.Float64(700.77),
                    ItemRef: &shared.ItemRef{
                        ID: "b4b393f3-5666-425b-aa32-201dec379c59",
                        Name: codataccounting.String("Dr. Elias Nitzsche"),
                    },
                    Quantity: codataccounting.Float64(5072.63),
                    SubTotal: codataccounting.Float64(7816.86),
                    TaxAmount: codataccounting.Float64(3272.41),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8119.2),
                        ID: codataccounting.String("2f9e21d9-0fd5-4377-abfc-7677f0f504a6"),
                        Name: codataccounting.String("Oscar Lesch"),
                    },
                    TotalAmount: codataccounting.Float64(9417.31),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6daee19c-26c0-4cb6-98c6-331cabdab767",
                            Name: codataccounting.String("Miguel Goldner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d0da0abe-58eb-43d5-8ba1-cb3ad49b8e5c",
                            Name: codataccounting.String("Lena Pouros"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e87f6482-3255-4be9-9c0c-bcb2ca87908d",
                            Name: codataccounting.String("Miss Tiffany Lowe"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(116.47),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("doloribus"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("eum"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("New Brunswick"),
                    Country: codataccounting.String("Cameroon"),
                    Line1: codataccounting.String("qui"),
                    Line2: codataccounting.String("totam"),
                    PostalCode: codataccounting.String("34033-6304"),
                    Region: codataccounting.String("numquam"),
                    Type: shared.AddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Zula_Ritchie@hotmail.com"),
                    Name: codataccounting.String("Duane Legros"),
                    Phone: codataccounting.String("(365) 620-2242 x60633"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusUnknown.ToPointer(),
            SubTotal: codataccounting.Float64(3727.48),
            SupplierRef: &shared.SupplierRef{
                ID: "3876e39d-ef9c-4765-9fd7-354e5cb94977",
                SupplierName: codataccounting.String("voluptatem"),
            },
            TotalAmount: codataccounting.Float64(963.17),
            TotalDiscount: codataccounting.Float64(4507.31),
            TotalTaxAmount: codataccounting.Float64(6874.97),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "quia",
        TimeoutInMinutes: codataccounting.Int(375533),
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

