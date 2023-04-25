# UpdatePurchaseOrder
Available in: `PurchaseOrders`

Posts an updated purchase order to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call []().

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support updating purchase orders.

## Example Usage
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
            Currency: codataccounting.String("dolor"),
            CurrencyRate: codataccounting.Float64(2424.79),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("74902848-826b-4b03-87fd-225e47871a88"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d72a2d4a-f415-48ac-ad0f-0f58c3b87b47"),
                        Name: codataccounting.String("Rhonda Ankunding PhD"),
                    },
                    Description: codataccounting.String("excepturi"),
                    DiscountAmount: codataccounting.Float64(5506.48),
                    DiscountPercentage: codataccounting.Float64(9007.56),
                    ItemRef: &shared.ItemRef{
                        ID: "9d82c5e3-06f5-4576-b5cd-eb0286d0bc43",
                        Name: codataccounting.String("Patrick Lynch"),
                    },
                    Quantity: codataccounting.Float64(2190.95),
                    SubTotal: codataccounting.Float64(4751.31),
                    TaxAmount: codataccounting.Float64(5490.08),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9580.48),
                        ID: codataccounting.String("2fcff81d-df7e-4088-b74e-f54c9216e892"),
                        Name: codataccounting.String("Emily Bergnaum"),
                    },
                    TotalAmount: codataccounting.Float64(7427.97),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fc2c8d27-0109-46b6-aad6-e3e1d9d3b660",
                            Name: codataccounting.String("Ethel Hagenes Sr."),
                        },
                        shared.TrackingCategoryRef{
                            ID: "aa1d5d22-47de-49b3-9461-70e768a96bb3",
                            Name: codataccounting.String("Everett Kiehn"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(1975.93),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("98eba1bb-f714-4335-af63-49a164249b21"),
                        Name: codataccounting.String("Sheri Walsh"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(6131.78),
                    DiscountPercentage: codataccounting.Float64(3665.61),
                    ItemRef: &shared.ItemRef{
                        ID: "1652b158-ca91-442f-8526-32b31cad692f",
                        Name: codataccounting.String("Sherman Labadie"),
                    },
                    Quantity: codataccounting.Float64(3433.66),
                    SubTotal: codataccounting.Float64(555.05),
                    TaxAmount: codataccounting.Float64(27.62),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3143.21),
                        ID: codataccounting.String("e9d3d934-e036-4f5c-b886-64f6985530a2"),
                        Name: codataccounting.String("Jeremy Murazik"),
                    },
                    TotalAmount: codataccounting.Float64(4284.76),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "af863c28-d040-4c69-a3d9-06f6ebd5ad7e",
                            Name: codataccounting.String("Kelly DuBuque"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f25f634b-3730-4714-a6be-8c3e09c64d34",
                            Name: codataccounting.String("Jan Schowalter"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9a6e5e7a-ef13-4402-a945-f53743efde11",
                            Name: codataccounting.String("Hugh Crona DVM"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(6123.82),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b1f7d9af-fe69-4682-acee-fb04f8c512ca"),
                        Name: codataccounting.String("Marco Ullrich"),
                    },
                    Description: codataccounting.String("voluptatem"),
                    DiscountAmount: codataccounting.Float64(5033.38),
                    DiscountPercentage: codataccounting.Float64(9159.33),
                    ItemRef: &shared.ItemRef{
                        ID: "d5798d38-5d46-4059-9d5c-3349576d5520",
                        Name: codataccounting.String("Frankie Miller"),
                    },
                    Quantity: codataccounting.Float64(1818.58),
                    SubTotal: codataccounting.Float64(3579.41),
                    TaxAmount: codataccounting.Float64(2057.04),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7058.26),
                        ID: codataccounting.String("6d765886-eeae-45fd-8b39-f8a1490678f1"),
                        Name: codataccounting.String("Adrienne Johnson"),
                    },
                    TotalAmount: codataccounting.Float64(8309.31),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "39fc9e17-5ffa-4906-ae55-9b72eb674603",
                            Name: codataccounting.String("Ms. Ora Thompson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "76c2bede-e767-490e-90c1-6a7ba4784044",
                            Name: codataccounting.String("Austin Ziemann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "70ef0480-91a2-4ba2-9ee6-c75af8a60a7a",
                            Name: codataccounting.String("Curtis Grimes"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(193),
                },
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("979e5afe-60ac-4aca-a45d-e9867551a9cc"),
                        Name: codataccounting.String("Ben Berge"),
                    },
                    Description: codataccounting.String("sunt"),
                    DiscountAmount: codataccounting.Float64(7629.97),
                    DiscountPercentage: codataccounting.Float64(4878.89),
                    ItemRef: &shared.ItemRef{
                        ID: "9a39ae5a-4d5a-465b-8d55-62d9b7d9e2d6",
                        Name: codataccounting.String("Jermaine Wiegand"),
                    },
                    Quantity: codataccounting.Float64(4577.32),
                    SubTotal: codataccounting.Float64(3923.07),
                    TaxAmount: codataccounting.Float64(1751.03),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5752.06),
                        ID: codataccounting.String("db875c3a-8902-482a-91f4-1cf6796ed3d7"),
                        Name: codataccounting.String("Ms. Elaine Schroeder"),
                    },
                    TotalAmount: codataccounting.Float64(3231.27),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1e98cce3-f716-4600-9a0e-3aa61c6fe09d",
                            Name: codataccounting.String("Warren Champlin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3b32c8c7-c3c7-410e-9673-d905cb4bedef",
                            Name: codataccounting.String("Bernadette Block"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c3909955-2825-40dc-bbcd-3b121b88c1ee",
                            Name: codataccounting.String("Kerry Klocko III"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(959.1),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("adipisci"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            PurchaseOrderNumber: codataccounting.String("excepturi"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("Port Rachelmouth"),
                    Country: codataccounting.String("Venezuela"),
                    Line1: codataccounting.String("deserunt"),
                    Line2: codataccounting.String("aut"),
                    PostalCode: codataccounting.String("48144-2514"),
                    Region: codataccounting.String("facilis"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Xzavier.Hane77@hotmail.com"),
                    Name: codataccounting.String("Marc Treutel"),
                    Phone: codataccounting.String("938.276.5927 x0235"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.PurchaseOrderStatusEnumVoid.ToPointer(),
            SubTotal: codataccounting.Float64(3902.15),
            SupplierRef: &shared.SupplierRef{
                ID: "a87bb7ae-cbe5-469d-b0cb-069907f98944",
                SupplierName: codataccounting.String("vitae"),
            },
            TotalAmount: codataccounting.Float64(2776.66),
            TotalDiscount: codataccounting.Float64(3416.95),
            TotalTaxAmount: codataccounting.Float64(1275.33),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        PurchaseOrderID: "culpa",
        TimeoutInMinutes: codataccounting.Int(573150),
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