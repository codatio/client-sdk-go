# PurchaseOrders

## Overview

Purchase orders

### Available Operations

* [CreatePurchaseOrder](#createpurchaseorder) - Create purchase order
* [GetCreateUpdatePurchaseOrdersModel](#getcreateupdatepurchaseordersmodel) - Get create/update purchase order model
* [GetPurchaseOrder](#getpurchaseorder) - Get purchase order
* [ListPurchaseOrders](#listpurchaseorders) - List purchase orders
* [UpdatePurchaseOrder](#updatepurchaseorder) - Update purchase order

## CreatePurchaseOrder

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
    req := operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("et"),
            CurrencyRate: codataccounting.Float64(8193.41),
            DeliveryDate: codataccounting.String("unde"),
            ExpectedDeliveryDate: codataccounting.String("esse"),
            ID: codataccounting.String("80a10c47-b950-440d-ac8b-2a5f002207e4"),
            IssueDate: codataccounting.String("quae"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8f90009e-d290-4278-ab4a-e9d64161e915"),
                        Name: codataccounting.String("Patricia Dickens"),
                    },
                    Description: codataccounting.String("rerum"),
                    DiscountAmount: codataccounting.Float64(1793.89),
                    DiscountPercentage: codataccounting.Float64(7581.19),
                    ItemRef: &shared.ItemRef{
                        ID: "09b92477-1f56-469e-9b7e-c7626649d84e",
                        Name: codataccounting.String("Terrence Walter"),
                    },
                    Quantity: codataccounting.Float64(9525.87),
                    SubTotal: codataccounting.Float64(8696.62),
                    TaxAmount: codataccounting.Float64(1739.26),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1417.37),
                        ID: codataccounting.String("76e0b88f-b87d-46fa-9b6e-8dbf812f83b1"),
                        Name: codataccounting.String("Matt Johnston"),
                    },
                    TotalAmount: codataccounting.Float64(9738.23),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c561929c-ca95-460a-9395-918da1d48e78",
                            Name: codataccounting.String("Dale Schamberger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e1143da9-308b-427a-88af-22184439b3de",
                            Name: codataccounting.String("Brad Hayes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cce470cd-2147-4b6e-a152-cf01d0d8c3a4",
                            Name: codataccounting.String("Luther Osinski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f935dfe9-74fa-44b1-a9c0-97eda623442e",
                            Name: codataccounting.String("Alberta Mills"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(4631.93),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e9984c80-b479-4e89-9923-c18ca8d69c56"),
                        Name: codataccounting.String("Mrs. Terrance Christiansen"),
                    },
                    Description: codataccounting.String("animi"),
                    DiscountAmount: codataccounting.Float64(1643.63),
                    DiscountPercentage: codataccounting.Float64(411.11),
                    ItemRef: &shared.ItemRef{
                        ID: "207e4fae-038c-4d7f-9bc2-cabaf7fc2ccb",
                        Name: codataccounting.String("Jesus Reinger"),
                    },
                    Quantity: codataccounting.Float64(461.14),
                    SubTotal: codataccounting.Float64(8366.2),
                    TaxAmount: codataccounting.Float64(9710.36),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3915.17),
                        ID: codataccounting.String("8eaedb2e-e70b-4e06-9fb3-6add704080e0"),
                        Name: codataccounting.String("Leonard Wiegand"),
                    },
                    TotalAmount: codataccounting.Float64(2204.14),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5a034b11-4992-443a-ba69-87a472b709a1",
                            Name: codataccounting.String("Tracy Thiel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "30106853-9ce0-4932-910a-cd15d8cc306b",
                            Name: codataccounting.String("Irma Huels"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d37bd204-a1f3-440b-b36f-677a48519c33",
                            Name: codataccounting.String("Mr. Eva Marvin"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(3069.7),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("deleniti"),
            Note: codataccounting.String("quos"),
            PaymentDueDate: codataccounting.String("qui"),
            PurchaseOrderNumber: codataccounting.String("ex"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Rauton"),
                    Country: codataccounting.String("Denmark"),
                    Line1: codataccounting.String("porro"),
                    Line2: codataccounting.String("nihil"),
                    PostalCode: codataccounting.String("81138-2454"),
                    Region: codataccounting.String("architecto"),
                    Type: shared.AddressTypeEnumBilling,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Karelle.Treutel18@yahoo.com"),
                    Name: codataccounting.String("Roy Simonis"),
                    Phone: codataccounting.String("403.667.1809 x09357"),
                },
            },
            SourceModifiedDate: codataccounting.String("adipisci"),
            Status: shared.PurchaseOrderStatusEnumClosed.ToPointer(),
            SubTotal: codataccounting.Float64(5198.96),
            SupplierRef: &shared.SupplierRef{
                ID: "7b47040d-0d98-4e9d-82c5-e306f5576f5c",
                SupplierName: codataccounting.String("quibusdam"),
            },
            TotalAmount: codataccounting.Float64(9167.97),
            TotalDiscount: codataccounting.Float64(7367.93),
            TotalTaxAmount: codataccounting.Float64(57.18),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(178911),
    }

    res, err := s.PurchaseOrders.CreatePurchaseOrder(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

## GetCreateUpdatePurchaseOrdersModel

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
    req := operations.GetCreateUpdatePurchaseOrdersModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.PurchaseOrders.GetCreateUpdatePurchaseOrdersModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetPurchaseOrder

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
    req := operations.GetPurchaseOrderRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PurchaseOrderID: "totam",
    }

    res, err := s.PurchaseOrders.GetPurchaseOrder(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrder != nil {
        // handle response
    }
}
```

## ListPurchaseOrders

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
    req := operations.ListPurchaseOrdersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ea"),
    }

    res, err := s.PurchaseOrders.ListPurchaseOrders(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PurchaseOrders != nil {
        // handle response
    }
}
```

## UpdatePurchaseOrder

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
    req := operations.UpdatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("pariatur"),
            CurrencyRate: codataccounting.Float64(27.26),
            DeliveryDate: codataccounting.String("distinctio"),
            ExpectedDeliveryDate: codataccounting.String("maxime"),
            ID: codataccounting.String("43b18ab3-78f2-4fcf-b81d-df7e088f74ef"),
            IssueDate: codataccounting.String("quis"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c9216e89-2631-43bb-afc2-c8d2701096b6"),
                        Name: codataccounting.String("Genevieve Steuber"),
                    },
                    Description: codataccounting.String("amet"),
                    DiscountAmount: codataccounting.Float64(9260.74),
                    DiscountPercentage: codataccounting.Float64(963.03),
                    ItemRef: &shared.ItemRef{
                        ID: "d9d3b660-334a-411a-a1d5-d2247de9b3d4",
                        Name: codataccounting.String("Dr. Frances Kutch"),
                    },
                    Quantity: codataccounting.Float64(4290.14),
                    SubTotal: codataccounting.Float64(5505.38),
                    TaxAmount: codataccounting.Float64(6701.68),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5968.46),
                        ID: codataccounting.String("6bb39878-8398-4eba-9bbf-7143356f6349"),
                        Name: codataccounting.String("Keith Hoeger"),
                    },
                    TotalAmount: codataccounting.Float64(2761.09),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b211ce46-b951-4652-b158-ca9142f05263",
                            Name: codataccounting.String("Miss Opal Dickinson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d692ffc8-7450-405e-9d3d-934e036f5c38",
                            Name: codataccounting.String("Brett Hudson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6985530a-2e2a-4ed6-aaf8-63c28d040c69",
                            Name: codataccounting.String("Marvin Stracke III"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(9418.45),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6ebd5ad7-ec73-494f-a5f6-34b3730714e6"),
                        Name: codataccounting.String("Clay Lockman"),
                    },
                    Description: codataccounting.String("debitis"),
                    DiscountAmount: codataccounting.Float64(63.56),
                    DiscountPercentage: codataccounting.Float64(5959.44),
                    ItemRef: &shared.ItemRef{
                        ID: "c64d342a-c299-4a6e-9e7a-ef13402e945f",
                        Name: codataccounting.String("Cindy Koepp"),
                    },
                    Quantity: codataccounting.Float64(9094.5),
                    SubTotal: codataccounting.Float64(9428.69),
                    TaxAmount: codataccounting.Float64(8126.55),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9230.35),
                        ID: codataccounting.String("1198221f-9b1f-47d9-affe-69682aceefb0"),
                        Name: codataccounting.String("Lela Lemke"),
                    },
                    TotalAmount: codataccounting.Float64(750.04),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "caabea70-8ed5-4798-9385-d460599d5c33",
                            Name: codataccounting.String("Shelly Hickle"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(8165.54),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("minima"),
            Note: codataccounting.String("ullam"),
            PaymentDueDate: codataccounting.String("dolores"),
            PurchaseOrderNumber: codataccounting.String("accusantium"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Torphystead"),
                    Country: codataccounting.String("Oman"),
                    Line1: codataccounting.String("eos"),
                    Line2: codataccounting.String("qui"),
                    PostalCode: codataccounting.String("27384"),
                    Region: codataccounting.String("nisi"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Jordyn_Hyatt62@yahoo.com"),
                    Name: codataccounting.String("Vernon Will"),
                    Phone: codataccounting.String("359.660.2503 x4591"),
                },
            },
            SourceModifiedDate: codataccounting.String("dolor"),
            Status: shared.PurchaseOrderStatusEnumClosed.ToPointer(),
            SubTotal: codataccounting.Float64(4172.21),
            SupplierRef: &shared.SupplierRef{
                ID: "86d839fc-9e17-45ff-a906-ae559b72eb67",
                SupplierName: codataccounting.String("magnam"),
            },
            TotalAmount: codataccounting.Float64(4246.29),
            TotalDiscount: codataccounting.Float64(361.86),
            TotalTaxAmount: codataccounting.Float64(2316.09),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "sit",
        TimeoutInMinutes: codataccounting.Int(987456),
    }

    res, err := s.PurchaseOrders.UpdatePurchaseOrder(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
