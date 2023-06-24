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

Posts a new purchase order to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update purchase order model](https://docs.codat.io/accounting-api#/operations/get-create-update-purchaseOrders-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) to see which integrations support this endpoint.

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
    res, err := s.PurchaseOrders.Create(ctx, operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(8513.15),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("0c69a3d9-06f6-4ebd-9ad7-ec7394f25f63"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3730714e-6be8-4c3e-89c6-4d342ac299a6"),
                        Name: codataccounting.String("Glen Toy"),
                    },
                    Description: codataccounting.String("earum"),
                    DiscountAmount: codataccounting.Float64(9610.27),
                    DiscountPercentage: codataccounting.Float64(791.87),
                    ItemRef: &shared.ItemRef{
                        ID: "3402e945-f537-443e-bde1-198221f9b1f7",
                        Name: codataccounting.String("Nick Nolan"),
                    },
                    Quantity: codataccounting.Float64(9212.73),
                    SubTotal: codataccounting.Float64(4322.15),
                    TaxAmount: codataccounting.Float64(6074.58),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3980.98),
                        ID: codataccounting.String("82aceefb-04f8-4c51-acaa-bea708ed5798"),
                        Name: codataccounting.String("Lee Luettgen"),
                    },
                    TotalAmount: codataccounting.Float64(2876.48),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0599d5c3-3495-476d-9520-9e9a2253b6d7",
                            Name: codataccounting.String("Lynn Lang"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "eeae5fd4-b39f-48a1-8906-78f13c686d83",
                            Name: codataccounting.String("Amos Roob"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(756.85),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("75ffa906-ae55-49b7-aeb6-746030fe1837"),
                        Name: codataccounting.String("Sheri Cole"),
                    },
                    Description: codataccounting.String("at"),
                    DiscountAmount: codataccounting.Float64(9103.82),
                    DiscountPercentage: codataccounting.Float64(8766.6),
                    ItemRef: &shared.ItemRef{
                        ID: "76790ed0-c16a-47ba-8784-04489f6770ef",
                        Name: codataccounting.String("Ms. Suzanne Lang MD"),
                    },
                    Quantity: codataccounting.Float64(1711.34),
                    SubTotal: codataccounting.Float64(7457.67),
                    TaxAmount: codataccounting.Float64(6341.57),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1475.97),
                        ID: codataccounting.String("5ee6c75a-f8a6-40a7-ae34-6e0979e5afe6"),
                        Name: codataccounting.String("Cecilia Ryan"),
                    },
                    TotalAmount: codataccounting.Float64(6433.94),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "45de9867-551a-49cc-a61e-c2c79a39ae5a",
                            Name: codataccounting.String("Krystal Hilll"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5b4d5562-d9b7-4d9e-ad6f-cf557629db87",
                            Name: codataccounting.String("Adrienne Donnelly"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(5958.7),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0282a51f-41cf-4679-aed3-d724c18f581e"),
                        Name: codataccounting.String("Dwight Sauer"),
                    },
                    Description: codataccounting.String("consectetur"),
                    DiscountAmount: codataccounting.Float64(9716.2),
                    DiscountPercentage: codataccounting.Float64(4600.54),
                    ItemRef: &shared.ItemRef{
                        ID: "16600da0-e3aa-461c-afe0-9d852b53b32c",
                        Name: codataccounting.String("Spencer Kuhn"),
                    },
                    Quantity: codataccounting.Float64(7573.4),
                    SubTotal: codataccounting.Float64(4608.74),
                    TaxAmount: codataccounting.Float64(958.62),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6.22),
                        ID: codataccounting.String("e1673d90-5cb4-4bed-af3c-127c39099552"),
                        Name: codataccounting.String("Dr. Phillip Hilpert"),
                    },
                    TotalAmount: codataccounting.Float64(7498.85),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cd3b121b-88c1-4ee5-a7a0-61391cc8fa0b",
                            Name: codataccounting.String("Kari Breitenberg"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4926b0cf-5e6c-4b6e-babe-5e0b99f3b135",
                            Name: codataccounting.String("Orville Hudson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7bb7aecb-e569-4d70-8b06-9907f9894414",
                            Name: codataccounting.String("Judith Nader"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(592.69),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("sapiente"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("quaerat"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("East Nya"),
                    Country: codataccounting.String("Isle of Man"),
                    Line1: codataccounting.String("inventore"),
                    Line2: codataccounting.String("quidem"),
                    PostalCode: codataccounting.String("02176-7883"),
                    Region: codataccounting.String("neque"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Halie38@hotmail.com"),
                    Name: codataccounting.String("Max Homenick"),
                    Phone: codataccounting.String("(399) 353-5882 x111"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusUnknown.ToPointer(),
            SubTotal: codataccounting.Float64(9856.77),
            SupplierRef: &shared.SupplierRef{
                ID: "e6664c41-d2fb-4a5c-ba06-9b8d291beb81",
                SupplierName: codataccounting.String("eaque"),
            },
            TotalAmount: codataccounting.Float64(6653.38),
            TotalDiscount: codataccounting.Float64(1485.57),
            TotalTaxAmount: codataccounting.Float64(6798.42),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(672553),
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

Get purchase order

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
    res, err := s.PurchaseOrders.Get(ctx, operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "atque",
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

Get create/update purchase order model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating and updating purchase orders.

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
        Query: codataccounting.String("molestiae"),
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

Posts an updated purchase order to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support updating purchase orders.

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
    res, err := s.PurchaseOrders.Update(ctx, operations.UpdatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(5976.63),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("9edd4fcf-7b50-4cf8-bf08-f39271076a24"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0c8f08bf-f108-41e8-8f86-996c8e22be0a"),
                        Name: codataccounting.String("Kendra White"),
                    },
                    Description: codataccounting.String("blanditiis"),
                    DiscountAmount: codataccounting.Float64(5720.7),
                    DiscountPercentage: codataccounting.Float64(2089.59),
                    ItemRef: &shared.ItemRef{
                        ID: "bd23f866-00c6-41c7-8342-73caa9118b38",
                        Name: codataccounting.String("Harold Ritchie MD"),
                    },
                    Quantity: codataccounting.Float64(2482.29),
                    SubTotal: codataccounting.Float64(1910.69),
                    TaxAmount: codataccounting.Float64(781.25),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6791.41),
                        ID: codataccounting.String("54dc1029-4f92-4fed-939b-a8f71e2992c2"),
                        Name: codataccounting.String("Mr. Casey Towne"),
                    },
                    TotalAmount: codataccounting.Float64(5593.62),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c3adc647-d240-4bc1-9ea4-82824ccc6a2f",
                            Name: codataccounting.String("Johanna Heathcote"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d3cb11a7-687d-4310-8e8e-2b9b0d746d2a",
                            Name: codataccounting.String("Adrienne Kreiger PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a0e79fa9-bbe5-4f17-9f65-0b1e707e7e43",
                            Name: codataccounting.String("Mrs. Angel Kuhn"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6307.18),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("cce072ab-d619-418d-a79c-10c18516fd78"),
                        Name: codataccounting.String("Grady Considine"),
                    },
                    Description: codataccounting.String("ab"),
                    DiscountAmount: codataccounting.Float64(1562.87),
                    DiscountPercentage: codataccounting.Float64(4579.62),
                    ItemRef: &shared.ItemRef{
                        ID: "2628fa50-3962-4867-a72b-3a65024b157f",
                        Name: codataccounting.String("Pete Rohan"),
                    },
                    Quantity: codataccounting.Float64(9998.39),
                    SubTotal: codataccounting.Float64(4376.21),
                    TaxAmount: codataccounting.Float64(1341.14),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6278.1),
                        ID: codataccounting.String("50871d99-b661-4a7d-af16-8b6ccb2822b4"),
                        Name: codataccounting.String("Enrique Larson DVM"),
                    },
                    TotalAmount: codataccounting.Float64(8556.65),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f4a1e9c4-ae55-4140-a757-26e003c2f029",
                            Name: codataccounting.String("Diane Marks"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(991.13),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("cumque"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("accusamus"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("North Orin"),
                    Country: codataccounting.String("Nauru"),
                    Line1: codataccounting.String("excepturi"),
                    Line2: codataccounting.String("natus"),
                    PostalCode: codataccounting.String("23594-9179"),
                    Region: codataccounting.String("quasi"),
                    Type: shared.AddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Velma19@gmail.com"),
                    Name: codataccounting.String("Ricardo Johnston"),
                    Phone: codataccounting.String("1-741-500-1876 x034"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusOpen.ToPointer(),
            SubTotal: codataccounting.Float64(5264.74),
            SupplierRef: &shared.SupplierRef{
                ID: "c6720c31-03f1-4a40-80f3-ec8688fd8ec6",
                SupplierName: codataccounting.String("tenetur"),
            },
            TotalAmount: codataccounting.Float64(7601.6),
            TotalDiscount: codataccounting.Float64(64.47),
            TotalTaxAmount: codataccounting.Float64(2341.76),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "illo",
        TimeoutInMinutes: codataccounting.Int(175186),
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

