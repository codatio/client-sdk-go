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
            Currency: codataccounting.String("amet"),
            CurrencyRate: codataccounting.Float64(256.85),
            DeliveryDate: codataccounting.String("quos"),
            ExpectedDeliveryDate: codataccounting.String("voluptas"),
            ID: codataccounting.String("c10a856a-19d4-4665-ba97-259875dc0cec"),
            IssueDate: codataccounting.String("facilis"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("78bd248e-c6e8-4b24-8b1c-06c9c0649d2b"),
                        Name: codataccounting.String("Laurence Mraz"),
                    },
                    Description: codataccounting.String("deleniti"),
                    DiscountAmount: codataccounting.Float64(8562.89),
                    DiscountPercentage: codataccounting.Float64(8359.95),
                    ItemRef: &shared.ItemRef{
                        ID: "b1665c31-2c7f-4550-9472-1c176292dd78",
                        Name: codataccounting.String("Mattie Treutel DDS"),
                    },
                    Quantity: codataccounting.Float64(9633.21),
                    SubTotal: codataccounting.Float64(5503.18),
                    TaxAmount: codataccounting.Float64(7788.66),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1193.26),
                        ID: codataccounting.String("41841fe1-f87e-4a10-ba98-06ea1606399e"),
                        Name: codataccounting.String("Jack Kris"),
                    },
                    TotalAmount: codataccounting.Float64(1221.4),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "58d4ab5b-c80d-4ea7-bfd9-931ec9d106cf",
                            Name: codataccounting.String("Bobby Schiller"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ab840a28-ea06-472d-ab73-a34ca434cdb3",
                            Name: codataccounting.String("Rhonda Medhurst V"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(9691.25),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("252078a1-8a4b-40da-ad4b-5cf0616ee922"),
                        Name: codataccounting.String("Samantha Kuhic"),
                    },
                    Description: codataccounting.String("nulla"),
                    DiscountAmount: codataccounting.Float64(4354.13),
                    DiscountPercentage: codataccounting.Float64(416.21),
                    ItemRef: &shared.ItemRef{
                        ID: "daa0e149-cd1c-4cdd-b62b-bf92390015f2",
                        Name: codataccounting.String("Cassandra Mann"),
                    },
                    Quantity: codataccounting.Float64(9894.06),
                    SubTotal: codataccounting.Float64(2550.98),
                    TaxAmount: codataccounting.Float64(9668.03),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9639.86),
                        ID: codataccounting.String("eb9bec50-318a-481e-b01d-297f7b456a85"),
                        Name: codataccounting.String("Ms. Casey Heathcote"),
                    },
                    TotalAmount: codataccounting.Float64(8073.54),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a326341a-cccc-4a66-bd4a-8595c1b32bb2",
                            Name: codataccounting.String("Edith VonRueden"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "31cd6a5b-e749-406b-96c6-36e74d28a481",
                            Name: codataccounting.String("Miss Velma Murphy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "41198640-5876-4b30-8711-3de4061082d0",
                            Name: codataccounting.String("Zachary Effertz"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(5757.79),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("cd927a5b-a551-41bb-8370-d9784fb14647"),
                        Name: codataccounting.String("Emily Brakus"),
                    },
                    Description: codataccounting.String("minima"),
                    DiscountAmount: codataccounting.Float64(7006.85),
                    DiscountPercentage: codataccounting.Float64(9129.06),
                    ItemRef: &shared.ItemRef{
                        ID: "61b3f371-7287-44c3-b7c8-d439ec6bd2ca",
                        Name: codataccounting.String("John Beatty"),
                    },
                    Quantity: codataccounting.Float64(4610.28),
                    SubTotal: codataccounting.Float64(4580.46),
                    TaxAmount: codataccounting.Float64(7955.57),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4356.49),
                        ID: codataccounting.String("ebbbc9e9-744c-45b6-85a4-af2fcabccbca"),
                        Name: codataccounting.String("Denise Marks"),
                    },
                    TotalAmount: codataccounting.Float64(8858.4),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "06a6cabe-22a1-41f7-ba75-d8ff4452bed7",
                            Name: codataccounting.String("Dora Ankunding"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "48c282b8-716c-427f-ab17-5780304c40ac",
                            Name: codataccounting.String("Lula Gulgowski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8254fde3-7724-4350-ad85-a7f8cc2911a6",
                            Name: codataccounting.String("Cory Toy"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(7840.43),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6009f01d-d385-423c-9a4e-bb9fd83f83df"),
                        Name: codataccounting.String("Abraham Anderson"),
                    },
                    Description: codataccounting.String("vel"),
                    DiscountAmount: codataccounting.Float64(6843.48),
                    DiscountPercentage: codataccounting.Float64(534.75),
                    ItemRef: &shared.ItemRef{
                        ID: "94e2e9c2-205d-4fe7-a5bf-fbcb86015f21",
                        Name: codataccounting.String("Mrs. Genevieve Pfannerstill IV"),
                    },
                    Quantity: codataccounting.Float64(99.38),
                    SubTotal: codataccounting.Float64(4855.86),
                    TaxAmount: codataccounting.Float64(9424.35),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6764.72),
                        ID: codataccounting.String("7398247a-8721-47fe-9962-df3eee7c385c"),
                        Name: codataccounting.String("Dr. Ada King MD"),
                    },
                    TotalAmount: codataccounting.Float64(1095.45),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8dbcf6e1-9b90-412c-844e-231ba147727d",
                            Name: codataccounting.String("Nick Hermann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2adabf68-00b0-41bc-bc03-2f2c19dbf711",
                            Name: codataccounting.String("Kristi Mann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f21523e3-7136-4521-ba6e-596aa41b9e20",
                            Name: codataccounting.String("Edna O'Reilly"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(8049.73),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nesciunt"),
            Note: codataccounting.String("ab"),
            PaymentDueDate: codataccounting.String("ullam"),
            PurchaseOrderNumber: codataccounting.String("consectetur"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("West Melvinaborough"),
                    Country: codataccounting.String("Turkey"),
                    Line1: codataccounting.String("eaque"),
                    Line2: codataccounting.String("dolores"),
                    PostalCode: codataccounting.String("62413"),
                    Region: codataccounting.String("assumenda"),
                    Type: shared.AddressTypeBilling,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Bo_Ernser@yahoo.com"),
                    Name: codataccounting.String("Lucas Ward"),
                    Phone: codataccounting.String("(360) 661-7480 x44002"),
                },
            },
            SourceModifiedDate: codataccounting.String("sequi"),
            Status: shared.PurchaseOrderStatusClosed.ToPointer(),
            SubTotal: codataccounting.Float64(4563.81),
            SupplierRef: &shared.SupplierRef{
                ID: "99a2a18d-b129-4dc5-a4ab-b7b10caf244d",
                SupplierName: codataccounting.String("itaque"),
            },
            TotalAmount: codataccounting.Float64(1039.82),
            TotalDiscount: codataccounting.Float64(522.31),
            TotalTaxAmount: codataccounting.Float64(2042.81),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(861729),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

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
        PurchaseOrderID: "consequatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrder != nil {
        // handle response
    }
}
```

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

## List

Get purchase orders

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
        Query: codataccounting.String("quos"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrders != nil {
        // handle response
    }
}
```

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
            Currency: codataccounting.String("ratione"),
            CurrencyRate: codataccounting.Float64(313.23),
            DeliveryDate: codataccounting.String("id"),
            ExpectedDeliveryDate: codataccounting.String("eligendi"),
            ID: codataccounting.String("4d070c4e-396e-4562-85cc-b16373314dad"),
            IssueDate: codataccounting.String("minima"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b890542e-5a55-4a10-bd8a-c0f482f9e9a5"),
                        Name: codataccounting.String("Mrs. Marianne Dibbert"),
                    },
                    Description: codataccounting.String("dolorum"),
                    DiscountAmount: codataccounting.Float64(8885.29),
                    DiscountPercentage: codataccounting.Float64(1224.04),
                    ItemRef: &shared.ItemRef{
                        ID: "22f0bfec-c419-432d-b04b-3ae70934d9eb",
                        Name: codataccounting.String("Rick O'Reilly"),
                    },
                    Quantity: codataccounting.Float64(9449.55),
                    SubTotal: codataccounting.Float64(4261.48),
                    TaxAmount: codataccounting.Float64(9895.04),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4380.7),
                        ID: codataccounting.String("1b0652fe-6536-4fb3-8a41-4aa294d64c08"),
                        Name: codataccounting.String("Alan Torp MD"),
                    },
                    TotalAmount: codataccounting.Float64(5153.47),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "7151a354-ba1a-46d7-9dc3-9917b844c850",
                            Name: codataccounting.String("Brandi Moen"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(3811.97),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2f4946ca-2d72-466c-9763-812723aa03f8"),
                        Name: codataccounting.String("Kristi Bins PhD"),
                    },
                    Description: codataccounting.String("quam"),
                    DiscountAmount: codataccounting.Float64(6684.99),
                    DiscountPercentage: codataccounting.Float64(6945.05),
                    ItemRef: &shared.ItemRef{
                        ID: "3e07c05e-13db-488f-991f-98329f91922f",
                        Name: codataccounting.String("Angelica Nienow IV"),
                    },
                    Quantity: codataccounting.Float64(7196.87),
                    SubTotal: codataccounting.Float64(3786.8),
                    TaxAmount: codataccounting.Float64(3422.36),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2715.17),
                        ID: codataccounting.String("5000a5b3-6a22-42e3-bf77-0a2b42e5edf6"),
                        Name: codataccounting.String("Isabel Dare"),
                    },
                    TotalAmount: codataccounting.Float64(4624.39),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c6a710e5-4b47-4ec6-aaf9-bd8327e8f7d3",
                            Name: codataccounting.String("Alma Brown"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ebdd822a-f2c1-4679-98a0-a46646ec658e",
                            Name: codataccounting.String("Gwendolyn Bernhard"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(8534.49),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e0aee8c7-2213-4f42-9a03-38b71b3d2fd3"),
                        Name: codataccounting.String("Brad Farrell"),
                    },
                    Description: codataccounting.String("accusantium"),
                    DiscountAmount: codataccounting.Float64(4920.29),
                    DiscountPercentage: codataccounting.Float64(2368.27),
                    ItemRef: &shared.ItemRef{
                        ID: "088e75ab-7ff2-4a12-bb07-4cd44c23c0b1",
                        Name: codataccounting.String("Dr. Norma Kuvalis"),
                    },
                    Quantity: codataccounting.Float64(4240.66),
                    SubTotal: codataccounting.Float64(5411.69),
                    TaxAmount: codataccounting.Float64(2221.75),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3622.4),
                        ID: codataccounting.String("b93ced68-7bb4-453f-84af-461c7cb91c79"),
                        Name: codataccounting.String("Stuart Howe"),
                    },
                    TotalAmount: codataccounting.Float64(8904.64),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "23875a4a-2e87-4d87-b51e-22e77c0e6e11",
                            Name: codataccounting.String("Dave Lockman"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(1771.68),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("laboriosam"),
            Note: codataccounting.String("pariatur"),
            PaymentDueDate: codataccounting.String("minus"),
            PurchaseOrderNumber: codataccounting.String("ipsam"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Breitenbergton"),
                    Country: codataccounting.String("Cuba"),
                    Line1: codataccounting.String("voluptatem"),
                    Line2: codataccounting.String("suscipit"),
                    PostalCode: codataccounting.String("65659"),
                    Region: codataccounting.String("aspernatur"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Grady78@hotmail.com"),
                    Name: codataccounting.String("Ellen Willms V"),
                    Phone: codataccounting.String("454.712.0503 x206"),
                },
            },
            SourceModifiedDate: codataccounting.String("dicta"),
            Status: shared.PurchaseOrderStatusDraft.ToPointer(),
            SubTotal: codataccounting.Float64(4036.82),
            SupplierRef: &shared.SupplierRef{
                ID: "0070c0bc-de7e-450e-a441-101c138b4629",
                SupplierName: codataccounting.String("fugit"),
            },
            TotalAmount: codataccounting.Float64(2485.73),
            TotalDiscount: codataccounting.Float64(8394.11),
            TotalTaxAmount: codataccounting.Float64(3242.71),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "sit",
        TimeoutInMinutes: codataccounting.Int(985872),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
