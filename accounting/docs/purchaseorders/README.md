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

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating purchase orders.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Create(ctx, operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("quo"),
            CurrencyRate: codataccounting.Float64(9804.63),
            DeliveryDate: codataccounting.String("maxime"),
            ExpectedDeliveryDate: codataccounting.String("suscipit"),
            ID: codataccounting.String("c0e503f5-6831-4f1d-8ed8-7b28e8afabc9"),
            IssueDate: codataccounting.String("laudantium"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e241e43b-2342-4417-913e-3f62aa9ae4ae"),
                        Name: codataccounting.String("Lyle Rath"),
                    },
                    Description: codataccounting.String("provident"),
                    DiscountAmount: codataccounting.Float64(7674.79),
                    DiscountPercentage: codataccounting.Float64(2960.36),
                    ItemRef: &shared.ItemRef{
                        ID: "92c5e8ba-5d4a-4a4a-908b-d380c29aa8dd",
                        Name: codataccounting.String("Teresa Predovic"),
                    },
                    Quantity: codataccounting.Float64(6509.18),
                    SubTotal: codataccounting.Float64(6484.89),
                    TaxAmount: codataccounting.Float64(2149.29),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(623.49),
                        ID: codataccounting.String("b7b91449-ae69-4c08-8d41-8bb71804f423"),
                        Name: codataccounting.String("Clyde Goldner"),
                    },
                    TotalAmount: codataccounting.Float64(1897.28),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f377ac5c-9b7e-493b-aa3c-523105e7c34c",
                            Name: codataccounting.String("Preston Abernathy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b812a661-4894-44a8-a908-5075bc253825",
                            Name: codataccounting.String("Robin Goodwin"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(7488.6),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0a4e66ea-4757-48d1-b1e2-941818fc679b"),
                        Name: codataccounting.String("Shelley Cronin II"),
                    },
                    Description: codataccounting.String("dolorem"),
                    DiscountAmount: codataccounting.Float64(3596.63),
                    DiscountPercentage: codataccounting.Float64(6053.38),
                    ItemRef: &shared.ItemRef{
                        ID: "b855d015-b62c-48b8-ba38-a8a88c144200",
                        Name: codataccounting.String("Benjamin Schoen"),
                    },
                    Quantity: codataccounting.Float64(7210.53),
                    SubTotal: codataccounting.Float64(879.15),
                    TaxAmount: codataccounting.Float64(6295.87),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9062.13),
                        ID: codataccounting.String("1ecf8c34-946b-4ba7-a05a-8b4a9ec5b368"),
                        Name: codataccounting.String("Simon Rutherford"),
                    },
                    TotalAmount: codataccounting.Float64(4247.55),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "272760e9-66e9-47e0-9410-3347d78ff249",
                            Name: codataccounting.String("Evelyn Gutkowski"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6574.81),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("soluta"),
            Note: codataccounting.String("excepturi"),
            PaymentDueDate: codataccounting.String("voluptates"),
            PurchaseOrderNumber: codataccounting.String("veniam"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Marcusboro"),
                    Country: codataccounting.String("Palestinian Territory"),
                    Line1: codataccounting.String("a"),
                    Line2: codataccounting.String("ipsum"),
                    PostalCode: codataccounting.String("33428"),
                    Region: codataccounting.String("deserunt"),
                    Type: shared.AddressTypeEnumBilling,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Maxine96@yahoo.com"),
                    Name: codataccounting.String("Ryan Hagenes"),
                    Phone: codataccounting.String("272-416-7898 x68731"),
                },
            },
            SourceModifiedDate: codataccounting.String("qui"),
            Status: shared.PurchaseOrderStatusEnumOpen.ToPointer(),
            SubTotal: codataccounting.Float64(5437.75),
            SupplierRef: &shared.SupplierRef{
                ID: "14eca016-bc41-4ea1-b42d-4104a25ef71d",
                SupplierName: codataccounting.String("debitis"),
            },
            TotalAmount: codataccounting.Float64(3264.43),
            TotalDiscount: codataccounting.Float64(4979.58),
            TotalTaxAmount: codataccounting.Float64(6690.5),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(117700),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Get(ctx, operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "architecto",
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating and updating purchase orders.

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
            AuthHeader: "YOUR_API_KEY_HERE",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.List(ctx, operations.ListPurchaseOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("fugiat"),
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support updating purchase orders.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.PurchaseOrders.Update(ctx, operations.UpdatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("eum"),
            CurrencyRate: codataccounting.Float64(1110.45),
            DeliveryDate: codataccounting.String("numquam"),
            ExpectedDeliveryDate: codataccounting.String("deserunt"),
            ID: codataccounting.String("4317692e-a486-473d-922b-828a9030660f"),
            IssueDate: codataccounting.String("doloremque"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4c79b4cc-64c2-4b3a-b2c4-88ade62f6aa5"),
                        Name: codataccounting.String("Leona Olson"),
                    },
                    Description: codataccounting.String("eveniet"),
                    DiscountAmount: codataccounting.Float64(1787.12),
                    DiscountPercentage: codataccounting.Float64(483.47),
                    ItemRef: &shared.ItemRef{
                        ID: "83016ca3-4bb8-47d4-b621-27a607d16062",
                        Name: codataccounting.String("Mrs. Alex Heaney"),
                    },
                    Quantity: codataccounting.Float64(2049.69),
                    SubTotal: codataccounting.Float64(8634.15),
                    TaxAmount: codataccounting.Float64(7188.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5826.59),
                        ID: codataccounting.String("ca9f38bd-2be8-4787-8349-3f49aa8465a3"),
                        Name: codataccounting.String("Myrtle Feil"),
                    },
                    TotalAmount: codataccounting.Float64(6184.81),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "719d1cea-673d-486e-bb35-e49a3135778c",
                            Name: codataccounting.String("Derek Funk"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cb0e3ea9-7504-45ba-8f63-b215186ab5e3",
                            Name: codataccounting.String("Robert Daugherty"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "14315d15-6829-49e6-9afc-7186ff20b7a7",
                            Name: codataccounting.String("Elena Wisoky DDS"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6606.02),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("alias"),
            Note: codataccounting.String("at"),
            PaymentDueDate: codataccounting.String("dignissimos"),
            PurchaseOrderNumber: codataccounting.String("aliquid"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("South Raheem"),
                    Country: codataccounting.String("Bahrain"),
                    Line1: codataccounting.String("ex"),
                    Line2: codataccounting.String("eius"),
                    PostalCode: codataccounting.String("77903"),
                    Region: codataccounting.String("veniam"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Augusta31@yahoo.com"),
                    Name: codataccounting.String("Heather Simonis"),
                    Phone: codataccounting.String("588-271-5141"),
                },
            },
            SourceModifiedDate: codataccounting.String("quidem"),
            Status: shared.PurchaseOrderStatusEnumClosed.ToPointer(),
            SubTotal: codataccounting.Float64(5873.6),
            SupplierRef: &shared.SupplierRef{
                ID: "c3221697-b188-40fc-bb2b-93c15f670bd1",
                SupplierName: codataccounting.String("nihil"),
            },
            TotalAmount: codataccounting.Float64(5049.32),
            TotalDiscount: codataccounting.Float64(2535.46),
            TotalTaxAmount: codataccounting.Float64(5040.97),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "sequi",
        TimeoutInMinutes: codataccounting.Int(122858),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
