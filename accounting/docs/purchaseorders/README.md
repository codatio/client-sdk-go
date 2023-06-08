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
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(1678.91),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("7de9b3d4-6170-4e76-8a96-bb398788398e"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1bbf7143-356f-4634-9a16-4249b211ce46"),
                        Name: codataccounting.String("Ms. Joey Hilll"),
                    },
                    Description: codataccounting.String("sunt"),
                    DiscountAmount: codataccounting.Float64(7068.01),
                    DiscountPercentage: codataccounting.Float64(1068.42),
                    ItemRef: &shared.ItemRef{
                        ID: "58ca9142-f052-4632-b31c-ad692ffc8745",
                        Name: codataccounting.String("Patricia Haley"),
                    },
                    Quantity: codataccounting.Float64(8634.66),
                    SubTotal: codataccounting.Float64(2457.73),
                    TaxAmount: codataccounting.Float64(8374.91),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5692.42),
                        ID: codataccounting.String("34e036f5-c388-4664-b698-5530a2e2aed6"),
                        Name: codataccounting.String("Carlton White"),
                    },
                    TotalAmount: codataccounting.Float64(1991.8),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "28d040c6-9a3d-4906-b6eb-d5ad7ec7394f",
                            Name: codataccounting.String("Stacy Zulauf"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4b373071-4e6b-4e8c-be09-c64d342ac299",
                            Name: codataccounting.String("Tyler Walsh"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7aef1340-2e94-45f5-b743-efde1198221f",
                            Name: codataccounting.String("Randolph Blick"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d9affe69-682a-4cee-bb04-f8c512caabea",
                            Name: codataccounting.String("Sandra Labadie"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(3622.31),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("798d385d-4605-499d-9c33-49576d55209e"),
                        Name: codataccounting.String("Cameron Davis"),
                    },
                    Description: codataccounting.String("neque"),
                    DiscountAmount: codataccounting.Float64(7058.26),
                    DiscountPercentage: codataccounting.Float64(3982),
                    ItemRef: &shared.ItemRef{
                        ID: "d765886e-eae5-4fd4-b39f-8a1490678f13",
                        Name: codataccounting.String("Edgar Langworth"),
                    },
                    Quantity: codataccounting.Float64(5517.59),
                    SubTotal: codataccounting.Float64(2068.12),
                    TaxAmount: codataccounting.Float64(6087.38),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9380.85),
                        ID: codataccounting.String("c9e175ff-a906-4ae5-99b7-2eb6746030fe"),
                        Name: codataccounting.String("Leona Doyle"),
                    },
                    TotalAmount: codataccounting.Float64(7569.56),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "bedee767-90ed-40c1-aa7b-a478404489f6",
                            Name: codataccounting.String("Courtney Bednar"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(559.47),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("48091a2b-a25e-4e6c-b5af-8a60a7ae346e"),
                        Name: codataccounting.String("Cora Koch"),
                    },
                    Description: codataccounting.String("nemo"),
                    DiscountAmount: codataccounting.Float64(6415.96),
                    DiscountPercentage: codataccounting.Float64(9569.24),
                    ItemRef: &shared.ItemRef{
                        ID: "e60acaca-645d-4e98-a755-1a9cce61ec2c",
                        Name: codataccounting.String("Misty Nitzsche"),
                    },
                    Quantity: codataccounting.Float64(6860.16),
                    SubTotal: codataccounting.Float64(8894.95),
                    TaxAmount: codataccounting.Float64(3723.63),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6840.34),
                        ID: codataccounting.String("4d5a65b4-d556-42d9-b7d9-e2d6fcf55762"),
                        Name: codataccounting.String("Drew Pollich"),
                    },
                    TotalAmount: codataccounting.Float64(3233.18),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3a890282-a51f-441c-b679-6ed3d724c18f",
                            Name: codataccounting.String("Hilda Brekke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8cce3f71-6600-4da0-a3aa-61c6fe09d852",
                            Name: codataccounting.String("Leon Ferry"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2c8c7c3c-710e-4167-bd90-5cb4bedef3c1",
                            Name: codataccounting.String("Pearl Ruecker"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "09955282-50dc-4bbc-93b1-21b88c1ee5e7",
                            Name: codataccounting.String("Mrs. Anthony Johnston"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(976),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("maxime"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("delectus"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Ankundingfort"),
                    Country: codataccounting.String("Kyrgyz Republic"),
                    Line1: codataccounting.String("repellendus"),
                    Line2: codataccounting.String("beatae"),
                    PostalCode: codataccounting.String("42514"),
                    Region: codataccounting.String("facilis"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Xzavier.Hane77@hotmail.com"),
                    Name: codataccounting.String("Marc Treutel"),
                    Phone: codataccounting.String("938.276.5927 x0235"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusDraft.ToPointer(),
            SubTotal: codataccounting.Float64(6429.43),
            SupplierRef: &shared.SupplierRef{
                ID: "87bb7aec-be56-49d7-8cb0-69907f989441",
                SupplierName: codataccounting.String("incidunt"),
            },
            TotalAmount: codataccounting.Float64(3416.95),
            TotalDiscount: codataccounting.Float64(1275.33),
            TotalTaxAmount: codataccounting.Float64(6340.79),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(573150),
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
        PurchaseOrderID: "voluptatibus",
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
        Query: codataccounting.String("ipsa"),
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
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(9589.65),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            ID: codataccounting.String("42c61be1-33ba-4cde-932b-6526f862853f"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("859ce322-231f-4e66-a4c4-1d2fba5cba06"),
                        Name: codataccounting.String("Randolph Little"),
                    },
                    Description: codataccounting.String("cupiditate"),
                    DiscountAmount: codataccounting.Float64(1187.35),
                    DiscountPercentage: codataccounting.Float64(7294.91),
                    ItemRef: &shared.ItemRef{
                        ID: "eb810a2a-a874-4947-9edd-4fcf7b50cf87",
                        Name: codataccounting.String("Frank Lind"),
                    },
                    Quantity: codataccounting.Float64(5703.77),
                    SubTotal: codataccounting.Float64(1432.47),
                    TaxAmount: codataccounting.Float64(4590.04),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(762.27),
                        ID: codataccounting.String("076a24b4-0c8f-408b-bf10-81e88f86996c"),
                        Name: codataccounting.String("Clay Connelly"),
                    },
                    TotalAmount: codataccounting.Float64(9114.26),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a3cf4789-3bd2-43f8-a600-c61c7834273c",
                            Name: codataccounting.String("Mr. Ernesto Medhurst"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(7330.09),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("atque"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            PurchaseOrderNumber: codataccounting.String("veritatis"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Jenkinston"),
                    Country: codataccounting.String("Norway"),
                    Line1: codataccounting.String("non"),
                    Line2: codataccounting.String("ratione"),
                    PostalCode: codataccounting.String("63287"),
                    Region: codataccounting.String("ab"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Leonie_Gleichner17@gmail.com"),
                    Name: codataccounting.String("Doyle Senger"),
                    Phone: codataccounting.String("(765) 940-9156 x1810"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.PurchaseOrderStatusVoid.ToPointer(),
            SubTotal: codataccounting.Float64(802.05),
            SupplierRef: &shared.SupplierRef{
                ID: "228ac3ad-c647-4d24-8bc1-1ea482824ccc",
                SupplierName: codataccounting.String("aliquid"),
            },
            TotalAmount: codataccounting.Float64(6728.65),
            TotalDiscount: codataccounting.Float64(1512.54),
            TotalTaxAmount: codataccounting.Float64(9779.19),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "ipsa",
        TimeoutInMinutes: codataccounting.Int(971477),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
