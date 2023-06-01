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
            Currency: codataccounting.String("accusamus"),
            CurrencyRate: codataccounting.Float64(2806.9),
            DeliveryDate: codataccounting.String("quae"),
            ExpectedDeliveryDate: codataccounting.String("dolore"),
            ID: codataccounting.String("8f90009e-d290-4278-ab4a-e9d64161e915"),
            IssueDate: codataccounting.String("quae"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("323b2c09-b924-4771-b566-9e5b7ec76266"),
                        Name: codataccounting.String("Sherri Schuster"),
                    },
                    Description: codataccounting.String("necessitatibus"),
                    DiscountAmount: codataccounting.Float64(7309.44),
                    DiscountPercentage: codataccounting.Float64(5923.78),
                    ItemRef: &shared.ItemRef{
                        ID: "e4cfd227-6e0b-488f-b87d-6fa5b6e8dbf8",
                        Name: codataccounting.String("Phyllis Zboncak"),
                    },
                    Quantity: codataccounting.Float64(7212.12),
                    SubTotal: codataccounting.Float64(809.04),
                    TaxAmount: codataccounting.Float64(7530.97),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6568.11),
                        ID: codataccounting.String("6a9ffc56-1929-4cca-9560-a1395918da1d"),
                        Name: codataccounting.String("Penny Walsh"),
                    },
                    TotalAmount: codataccounting.Float64(8885.45),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cf8e1143-da93-408b-a7a0-8af22184439b",
                            Name: codataccounting.String("Desiree Walsh"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(3395.66),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("eum"),
            Note: codataccounting.String("eligendi"),
            PaymentDueDate: codataccounting.String("quisquam"),
            PurchaseOrderNumber: codataccounting.String("optio"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("El Paso"),
                    Country: codataccounting.String("Lesotho"),
                    Line1: codataccounting.String("accusantium"),
                    Line2: codataccounting.String("impedit"),
                    PostalCode: codataccounting.String("10246-4931"),
                    Region: codataccounting.String("ipsam"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Wendell84@hotmail.com"),
                    Name: codataccounting.String("Cristina Lemke"),
                    Phone: codataccounting.String("1-476-736-9523 x8985"),
                },
            },
            SourceModifiedDate: codataccounting.String("iure"),
            Status: shared.PurchaseOrderStatusDraft.ToPointer(),
            SubTotal: codataccounting.Float64(9973.37),
            SupplierRef: &shared.SupplierRef{
                ID: "a4b1e9c0-97ed-4a62-b442-e1a9237e9984",
                SupplierName: codataccounting.String("porro"),
            },
            TotalAmount: codataccounting.Float64(5572.49),
            TotalDiscount: codataccounting.Float64(62.92),
            TotalTaxAmount: codataccounting.Float64(7081.21),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(281064),
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
        PurchaseOrderID: "nihil",
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
        Query: codataccounting.String("sint"),
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
            Currency: codataccounting.String("saepe"),
            CurrencyRate: codataccounting.Float64(5389),
            DeliveryDate: codataccounting.String("excepturi"),
            ExpectedDeliveryDate: codataccounting.String("architecto"),
            ID: codataccounting.String("923c18ca-8d69-4c56-8921-4fa20207e4fa"),
            IssueDate: codataccounting.String("saepe"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("38cd7f1b-c2ca-4baf-bfc2-ccba4bef0df6"),
                        Name: codataccounting.String("Darin O'Hara"),
                    },
                    Description: codataccounting.String("facilis"),
                    DiscountAmount: codataccounting.Float64(1625.48),
                    DiscountPercentage: codataccounting.Float64(9369.6),
                    ItemRef: &shared.ItemRef{
                        ID: "e70be069-fb36-4add-b040-80e0a3fc73a5",
                        Name: codataccounting.String("Jason Fisher"),
                    },
                    Quantity: codataccounting.Float64(798.12),
                    SubTotal: codataccounting.Float64(679.91),
                    TaxAmount: codataccounting.Float64(2846.73),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5751.24),
                        ID: codataccounting.String("9243afa6-987a-4472-b709-a153e2230106"),
                        Name: codataccounting.String("Tommy Emard"),
                    },
                    TotalAmount: codataccounting.Float64(9037.93),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "932d10ac-d15d-48cc-b06b-786b3d37bd20",
                            Name: codataccounting.String("Harriet Berge"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(2936.48),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("accusantium"),
            Note: codataccounting.String("nam"),
            PaymentDueDate: codataccounting.String("rerum"),
            PurchaseOrderNumber: codataccounting.String("dolor"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Wavabury"),
                    Country: codataccounting.String("Jersey"),
                    Line1: codataccounting.String("odio"),
                    Line2: codataccounting.String("dolorum"),
                    PostalCode: codataccounting.String("53057"),
                    Region: codataccounting.String("dolor"),
                    Type: shared.AddressTypeUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Duncan13@yahoo.com"),
                    Name: codataccounting.String("Tom Lind"),
                    Phone: codataccounting.String("1-770-374-9811 x382"),
                },
            },
            SourceModifiedDate: codataccounting.String("in"),
            Status: shared.PurchaseOrderStatusOpen.ToPointer(),
            SubTotal: codataccounting.Float64(4798.4),
            SupplierRef: &shared.SupplierRef{
                ID: "1a88ed72-a2d4-4af4-958a-c2d0f0f58c3b",
                SupplierName: codataccounting.String("totam"),
            },
            TotalAmount: codataccounting.Float64(4462.83),
            TotalDiscount: codataccounting.Float64(7360.32),
            TotalTaxAmount: codataccounting.Float64(2850.04),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "molestiae",
        TimeoutInMinutes: codataccounting.Int(50921),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```
