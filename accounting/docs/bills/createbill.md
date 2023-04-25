# CreateBill
Available in: `Bills`

Posts a new bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating a bill.

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
    req := operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(8487.22),
            Currency: codataccounting.String("facilis"),
            CurrencyRate: codataccounting.Float64(2476.18),
            DueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("adebd5da-ea4c-4506-a8aa-94c02644cf5e"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d9a4578a-dc1a-4c60-8dec-001ac802e2ec"),
                        Name: codataccounting.String("Sonia Wiegand"),
                    },
                    Description: codataccounting.String("maiores"),
                    DiscountAmount: codataccounting.Float64(6.91),
                    DiscountPercentage: codataccounting.Float64(9926.67),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "816ff347-7c13-4e90-ac14-125b0960a668",
                        Name: codataccounting.String("Dana Berge"),
                    },
                    Quantity: 4625.83,
                    SubTotal: codataccounting.Float64(1693.12),
                    TaxAmount: codataccounting.Float64(6463.29),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9650.95),
                        ID: codataccounting.String("923c5949-f83f-4350-8f87-6ffb901c6ecb"),
                        Name: codataccounting.String("Joel Von"),
                    },
                    TotalAmount: codataccounting.Float64(1940.58),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f789ffaf-eda5-43e5-ae6e-0ac184c2b9c2",
                                Name: codataccounting.String("Bessie Schmidt"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "373a40e1-942f-432e-9505-5756f5d56d0b",
                                Name: codataccounting.String("Joseph Olson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "dfe13db4-f62c-4ba3-b894-1aebc0b80a69",
                                Name: codataccounting.String("Rhonda Schuster"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2ecfcc8f-8950-410f-9dd3-d6fa1804e54c",
                                Name: codataccounting.String("Ms. Russell Wunsch"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("est"),
                            ID: "363c8873-e484-4380-b1f6-b8ca275a60a0",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "495cc699-171b-451c-9bdb-1cf4b888ebdf",
                            Name: codataccounting.String("Randall Schmeler"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "99bc7fc0-b2dc-4e10-873e-42b006d67887",
                            Name: codataccounting.String("Rickey Oberbrunner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "81a58208-c54f-4efa-9c95-f2eac5565d30",
                            Name: codataccounting.String("Bethany Zboncak"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "81206e28-13fa-44a4-9c48-0d3f2132af03",
                            Name: codataccounting.String("Sarah Collier"),
                        },
                    },
                    UnitAmount: 735.74,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4f4cc6f1-8bf9-4621-a6a4-f77a87ee3e4b"),
                        Name: codataccounting.String("Lance Hintz"),
                    },
                    Description: codataccounting.String("aliquid"),
                    DiscountAmount: codataccounting.Float64(3398.43),
                    DiscountPercentage: codataccounting.Float64(7044.02),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "34418e3b-b91c-48d9-b5e0-e8419d8f84f1",
                        Name: codataccounting.String("Clara Wyman"),
                    },
                    Quantity: 89.04,
                    SubTotal: codataccounting.Float64(4587.23),
                    TaxAmount: codataccounting.Float64(8913.02),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8174.54),
                        ID: codataccounting.String("cc4aa5f3-cabd-4905-a972-e056728227b2"),
                        Name: codataccounting.String("Antonio Beer"),
                    },
                    TotalAmount: codataccounting.Float64(4570.63),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "bf7a4fa8-7cf5-435a-afae-54ebf60c321f",
                                Name: codataccounting.String("Tammy Farrell"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nostrum"),
                            ID: "d2367fe1-a0cc-48df-b9f0-a396d90c364b",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "15dfbace-188b-41c4-ae2c-8c6ce611feeb",
                            Name: codataccounting.String("Angelica Kulas"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b6eec743-78ba-4253-9774-7dc915ad2caf",
                            Name: codataccounting.String("Angel Sporer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "23dc0f5a-e2f3-4a6b-b008-78756143f5a6",
                            Name: codataccounting.String("Eduardo Larkin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5554080d-40bc-4acc-acbd-6b5f3ec90930",
                            Name: codataccounting.String("Kristie Mohr"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "bad25538-19b4-474b-8ed2-0e56248fff63",
                            Name: codataccounting.String("Mr. Matt McLaughlin"),
                        },
                    },
                    UnitAmount: 7211.38,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("dcab6267-6696-4e1e-8002-21b335d89acb"),
                        Name: codataccounting.String("Raquel Rolfson"),
                    },
                    Description: codataccounting.String("id"),
                    DiscountAmount: codataccounting.Float64(5420.17),
                    DiscountPercentage: codataccounting.Float64(8453.65),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "0c549ef0-3004-4978-a61f-a1cf20688f77",
                        Name: codataccounting.String("Justin Willms"),
                    },
                    Quantity: 4515.89,
                    SubTotal: codataccounting.Float64(779.92),
                    TaxAmount: codataccounting.Float64(8156.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8049.36),
                        ID: codataccounting.String("a163f2a3-c80a-497f-b334-cddf857a9e61"),
                        Name: codataccounting.String("Darren Johnson"),
                    },
                    TotalAmount: codataccounting.Float64(6484.97),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "21d29dfc-94d6-4fec-9799-390066a6d2d0",
                                Name: codataccounting.String("Linda Frami"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "338cec08-6fa2-41e9-952c-b3119167b8e3",
                                Name: codataccounting.String("Julian Stanton I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "408d6d36-4ffd-4455-906d-1263d48e935c",
                                Name: codataccounting.String("Nichole Marks"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dicta"),
                            ID: "f30be3e4-3202-4d72-9657-6506641870d9",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1f9ad030-c4ec-4c11-a083-6429068b8502",
                            Name: codataccounting.String("Jon Hessel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "73bc845e-320a-4319-b4ba-df947c9a867b",
                            Name: codataccounting.String("Barry Daugherty"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6665816d-dca8-4ef5-9fcb-4c593ec12cda",
                            Name: codataccounting.String("Marty Beer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7afedbd8-0df4-448a-87f9-390c58880983",
                            Name: codataccounting.String("Blake Purdy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ef3ffdd9-f7f0-479a-b4d3-5724cdb0f4d2",
                            Name: codataccounting.String("Roger Bradtke"),
                        },
                    },
                    UnitAmount: 8548,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("enim"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("laudantium"),
                        CurrencyRate: codataccounting.Float64(2653.03),
                        TotalAmount: codataccounting.Float64(3017.01),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("eded85a9-065e-4628-bdfc-2032b6c87992"),
                            Name: codataccounting.String("Melody Kreiger I"),
                        },
                        Currency: codataccounting.String("ad"),
                        CurrencyRate: codataccounting.Float64(5398.86),
                        ID: codataccounting.String("4f7ae12c-6891-4f82-8e11-57172305377d"),
                        Note: codataccounting.String("minus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("a"),
                        TotalAmount: codataccounting.Float64(6863.01),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("totam"),
                        CurrencyRate: codataccounting.Float64(5875.39),
                        TotalAmount: codataccounting.Float64(8701),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f975e356-6860-492e-9c3d-dc5f111dea10"),
                            Name: codataccounting.String("Jeanne Stracke"),
                        },
                        Currency: codataccounting.String("architecto"),
                        CurrencyRate: codataccounting.Float64(6659.52),
                        ID: codataccounting.String("4d190feb-2178-40bc-8c0d-bbddb484708f"),
                        Note: codataccounting.String("facilis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("aliquam"),
                        TotalAmount: codataccounting.Float64(9229.15),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("91e6bc15-8c4c-44e5-8599-ea342260e9b2"),
                    PurchaseOrderNumber: codataccounting.String("accusantium"),
                },
            },
            Reference: codataccounting.String("consequatur"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.BillStatusEnumVoid,
            SubTotal: 9277.54,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "deleniti": map[string]interface{}{
                        "et": "expedita",
                        "quibusdam": "quos",
                        "maiores": "quidem",
                    },
                    "in": map[string]interface{}{
                        "doloremque": "fuga",
                        "dicta": "architecto",
                        "suscipit": "eligendi",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "e723d409-7fa3-40e9-af72-5b29122030d8",
                SupplierName: codataccounting.String("velit"),
            },
            TaxAmount: 9427.8,
            TotalAmount: 3564.85,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 9319.53,
                    Name: "Mathew Kunde",
                },
                shared.BillWithholdingTax{
                    Amount: 8287.35,
                    Name: "Theresa Terry",
                },
                shared.BillWithholdingTax{
                    Amount: 988.05,
                    Name: "Ivan Gulgowski",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(546089),
    }

    res, err := s.Bills.CreateBill(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillResponse != nil {
        // handle response
    }
}
```