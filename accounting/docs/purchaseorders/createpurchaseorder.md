# CreatePurchaseOrder
Available in: `PurchaseOrders`

Posts a new purchase order to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update purchase order model](https://docs.codat.io/accounting-api#/operations/get-create-update-purchaseOrders-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders) for integrations that support creating purchase orders.

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
    req := operations.CreatePurchaseOrderRequest{
        PurchaseOrder: &shared.PurchaseOrder{
            Currency: codataccounting.String("et"),
            CurrencyRate: codataccounting.Float64(5307.21),
            DeliveryDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ExpectedDeliveryDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("4439b3de-8756-4ccc-a470-cd2147b6e615"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.PurchaseOrderLineItem{
                shared.PurchaseOrderLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("cf01d0d8-c3a4-4b9a-9bf9-35dfe974fa4b"),
                        Name: codataccounting.String("Tasha Mante V"),
                    },
                    Description: codataccounting.String("molestiae"),
                    DiscountAmount: codataccounting.Float64(8810.35),
                    DiscountPercentage: codataccounting.Float64(8368.04),
                    ItemRef: &shared.ItemRef{
                        ID: "a623442e-1a92-437e-9984-c80b479e8919",
                        Name: codataccounting.String("Ms. Connie Romaguera"),
                    },
                    Quantity: codataccounting.Float64(6399.68),
                    SubTotal: codataccounting.Float64(5263.14),
                    TaxAmount: codataccounting.Float64(8167.53),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4265.02),
                        ID: codataccounting.String("9c568921-4fa2-4020-be4f-ae038cd7f1bc"),
                        Name: codataccounting.String("Nichole Ortiz"),
                    },
                    TotalAmount: codataccounting.Float64(9965.51),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fc2ccba4-bef0-4df6-8eae-db2ee70be069",
                            Name: codataccounting.String("Archie Feeney"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "dd704080-e0a3-4fc7-ba5a-034b11499243",
                            Name: codataccounting.String("Emilio O'Keefe"),
                        },
                    },
                    UnitAmount: codataccounting.Float64(5196.11),
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("in"),
            PaymentDueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            PurchaseOrderNumber: codataccounting.String("est"),
            ShipTo: &shared.ShipTo{
                Address: &shared.Addressesitems{
                    City: codataccounting.String("South Caesar"),
                    Country: codataccounting.String("Reunion"),
                    Line1: codataccounting.String("esse"),
                    Line2: codataccounting.String("sit"),
                    PostalCode: codataccounting.String("61328-1120"),
                    Region: codataccounting.String("dicta"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                Contact: &shared.ShipToContact{
                    Email: codataccounting.String("Joyce60@yahoo.com"),
                    Name: codataccounting.String("Elias Bednar"),
                    Phone: codataccounting.String("1-810-778-1385"),
                },
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.PurchaseOrderStatusEnumClosed.ToPointer(),
            SubTotal: codataccounting.Float64(7865.82),
            SupplierRef: &shared.SupplierRef{
                ID: "306b786b-3d37-4bd2-84a1-f340bb36f677",
                SupplierName: codataccounting.String("dolorum"),
            },
            TotalAmount: codataccounting.Float64(2612.94),
            TotalDiscount: codataccounting.Float64(5072.49),
            TotalTaxAmount: codataccounting.Float64(3603.33),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(80773),
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