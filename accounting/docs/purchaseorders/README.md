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
            Currency: codataccounting.String("occaecati"),
            CurrencyRate: codataccounting.Float64(5196.11),
            DeliveryDate: codataccounting.String("in"),
            ExpectedDeliveryDate: codataccounting.String("est"),
            ID: codataccounting.String("472b709a-153e-4223-8106-8539ce0932d1"),
            IssueDate: codataccounting.String("aperiam"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("cd15d8cc-306b-4786-b3d3-7bd204a1f340"),
                        Name: codataccounting.String("Jean Fahey"),
                    },
                    Description: codataccounting.String("voluptas"),
                    DiscountAmount: codataccounting.Float64(4441.02),
                    DiscountPercentage: codataccounting.Float64(4856.38),
                    ItemRef: &shared.ItemRef{
                        ID: "a48519c3-3749-4028-8882-6bb03c7fd225",
                        Name: codataccounting.String("Micheal King"),
                    },
                    Quantity: codataccounting.Float64(1000.75),
                    SubTotal: codataccounting.Float64(6359.09),
                    TaxAmount: codataccounting.Float64(5010.45),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5418.34),
                        ID: codataccounting.String("ed72a2d4-af41-458a-82d0-f0f58c3b87b4"),
                        Name: codataccounting.String("Dr. Jessica Goldner PhD"),
                    },
                    TotalAmount: codataccounting.Float64(5697.11),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e9d82c5e-306f-4557-af5c-deb0286d0bc4",
                            Name: codataccounting.String("Maryann Boehm"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b378f2fc-ff81-4ddf-be08-8f74ef54c921",
                            Name: codataccounting.String("Meghan Lindgren"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6313bb6f-c2c8-4d27-8109-6b66ad6e3e1d",
                            Name: codataccounting.String("Nicolas Emmerich"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(3972.04),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0334a11a-a1d5-4d22-87de-9b3d46170e76"),
                        Name: codataccounting.String("Cameron Metz"),
                    },
                    Description: codataccounting.String("facilis"),
                    DiscountAmount: codataccounting.Float64(1958.51),
                    DiscountPercentage: codataccounting.Float64(6132.25),
                    ItemRef: &shared.ItemRef{
                        ID: "8788398e-ba1b-4bf7-9433-56f6349a1642",
                        Name: codataccounting.String("Shelly Quitzon Sr."),
                    },
                    Quantity: codataccounting.Float64(7569.85),
                    SubTotal: codataccounting.Float64(9251.22),
                    TaxAmount: codataccounting.Float64(3035.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4049.25),
                        ID: codataccounting.String("b951652b-158c-4a91-82f0-52632b31cad6"),
                        Name: codataccounting.String("Wayne Wilkinson"),
                    },
                    TotalAmount: codataccounting.Float64(5042.69),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "45005e9d-3d93-44e0-b6f5-c388664f6985",
                            Name: codataccounting.String("Tracy Bahringer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e2aed6aa-f863-4c28-9040-c69a3d906f6e",
                            Name: codataccounting.String("Josh Heathcote"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(4666.58),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ec7394f2-5f63-44b3-b307-14e6be8c3e09"),
                        Name: codataccounting.String("Cecil Hahn"),
                    },
                    Description: codataccounting.String("aliquam"),
                    DiscountAmount: codataccounting.Float64(1723.11),
                    DiscountPercentage: codataccounting.Float64(6784.76),
                    ItemRef: &shared.ItemRef{
                        ID: "c299a6e5-e7ae-4f13-802e-945f53743efd",
                        Name: codataccounting.String("Ryan Berge"),
                    },
                    Quantity: codataccounting.Float64(1577.51),
                    SubTotal: codataccounting.Float64(1756.68),
                    TaxAmount: codataccounting.Float64(716.71),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9955.38),
                        ID: codataccounting.String("9b1f7d9a-ffe6-4968-aace-efb04f8c512c"),
                        Name: codataccounting.String("Orlando Reilly"),
                    },
                    TotalAmount: codataccounting.Float64(4577.97),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8ed5798d-385d-4460-999d-5c3349576d55",
                            Name: codataccounting.String("Ruth Muller"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6702.81),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("eos"),
            Note: codataccounting.String("qui"),
            PaymentDueDate: codataccounting.String("corporis"),
            PurchaseOrderNumber: codataccounting.String("neque"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Jacobifield"),
                    Country: codataccounting.String("Kyrgyz Republic"),
                    Line1: codataccounting.String("nisi"),
                    Line2: codataccounting.String("veniam"),
                    PostalCode: codataccounting.String("53896-8398"),
                    Region: codataccounting.String("dolore"),
                    Type: shared.AddressTypeDelivery,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Lauren_Wolf@gmail.com"),
                    Name: codataccounting.String("Mrs. Esther Medhurst"),
                    Phone: codataccounting.String("912-845-3852 x6975"),
                },
            },
            SourceModifiedDate: codataccounting.String("saepe"),
            Status: shared.PurchaseOrderStatusUnknown.ToPointer(),
            SubTotal: codataccounting.Float64(4594.79),
            SupplierRef: &shared.SupplierRef{
                ID: "5ffa906a-e559-4b72-ab67-46030fe18376",
                SupplierName: codataccounting.String("eligendi"),
            },
            TotalAmount: codataccounting.Float64(1311.2),
            TotalDiscount: codataccounting.Float64(7109.21),
            TotalTaxAmount: codataccounting.Float64(9353.25),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(870347),
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
        PurchaseOrderID: "eveniet",
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
        Query: codataccounting.String("vero"),
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
            Currency: codataccounting.String("iure"),
            CurrencyRate: codataccounting.Float64(4003.93),
            DeliveryDate: codataccounting.String("dignissimos"),
            ExpectedDeliveryDate: codataccounting.String("perspiciatis"),
            ID: codataccounting.String("0ed0c16a-7ba4-4784-8448-9f6770ef0480"),
            IssueDate: codataccounting.String("occaecati"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a2ba25ee-6c75-4af8-a60a-7ae346e0979e"),
                        Name: codataccounting.String("Blanche Will"),
                    },
                    Description: codataccounting.String("consequatur"),
                    DiscountAmount: codataccounting.Float64(6343.86),
                    DiscountPercentage: codataccounting.Float64(7740.9),
                    ItemRef: &shared.ItemRef{
                        ID: "aca645de-9867-4551-a9cc-e61ec2c79a39",
                        Name: codataccounting.String("Stewart Hintz"),
                    },
                    Quantity: codataccounting.Float64(8250.33),
                    SubTotal: codataccounting.Float64(3666.61),
                    TaxAmount: codataccounting.Float64(6621.26),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3870.84),
                        ID: codataccounting.String("5b4d5562-d9b7-4d9e-ad6f-cf557629db87"),
                        Name: codataccounting.String("Adrienne Donnelly"),
                    },
                    TotalAmount: codataccounting.Float64(5958.7),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "282a51f4-1cf6-4796-ad3d-724c18f581e9",
                            Name: codataccounting.String("Lowell Schimmel"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(9716.2),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("esse"),
            Note: codataccounting.String("sunt"),
            PaymentDueDate: codataccounting.String("autem"),
            PurchaseOrderNumber: codataccounting.String("aliquid"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("North Retafurt"),
                    Country: codataccounting.String("Argentina"),
                    Line1: codataccounting.String("officiis"),
                    Line2: codataccounting.String("dolor"),
                    PostalCode: codataccounting.String("63074-9905"),
                    Region: codataccounting.String("assumenda"),
                    Type: shared.AddressTypeBilling,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Brandt23@gmail.com"),
                    Name: codataccounting.String("Danny Daugherty"),
                    Phone: codataccounting.String("581-840-0803 x42850"),
                },
            },
            SourceModifiedDate: codataccounting.String("nostrum"),
            Status: shared.PurchaseOrderStatusClosed.ToPointer(),
            SubTotal: codataccounting.Float64(7098.7),
            SupplierRef: &shared.SupplierRef{
                ID: "4bedef3c-127c-4390-9955-28250dcbbcd3",
                SupplierName: codataccounting.String("facilis"),
            },
            TotalAmount: codataccounting.Float64(1020.04),
            TotalDiscount: codataccounting.Float64(1677.86),
            TotalTaxAmount: codataccounting.Float64(1207.36),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "tempore",
        TimeoutInMinutes: codataccounting.Int(539092),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
