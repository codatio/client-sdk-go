# CreateDirectIncome
Available in: `DirectIncomes`

Posts a new direct income to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating direct incomes.

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
    req := operations.CreateDirectIncomeRequest{
        DirectIncome: &shared.DirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("est"),
                ID: "aee79e3c-71ad-431b-acb8-3d2378ae3bfc",
            },
            Currency: "sed",
            CurrencyRate: codataccounting.Float64(2220.93),
            ID: codataccounting.String("d9450a98-6a49-45ba-8707-f06b28ecc864"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2386f62c-969c-44cc-ab78-890a3fd3c81d"),
                        Name: codataccounting.String("Samuel Aufderhar"),
                    },
                    Description: codataccounting.String("optio"),
                    DiscountAmount: codataccounting.Float64(1614.81),
                    DiscountPercentage: codataccounting.Float64(2485.3),
                    ItemRef: &shared.ItemRef{
                        ID: "df931da3-edb5-41fa-994a-cc9435137726",
                        Name: codataccounting.String("Peter Hand"),
                    },
                    Quantity: 1199.27,
                    SubTotal: codataccounting.Float64(7214.48),
                    TaxAmount: codataccounting.Float64(5545.08),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2429.72),
                        ID: codataccounting.String("2a56d691-80ff-460e-b9a6-658e69a4b843"),
                        Name: codataccounting.String("Chad Leannon"),
                    },
                    TotalAmount: codataccounting.Float64(6893.09),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c75c68c6-0659-4468-8e30-4d8849bf8214",
                            Name: codataccounting.String("Vincent Erdman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "96bb0c69-e372-4db1-b44b-a9f78a5c0ed7",
                            Name: codataccounting.String("Gerard Roberts DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "97261fb0-c58d-427b-9199-6b5b4b50eef7",
                            Name: codataccounting.String("Christina Rice"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7ab0344b-1710-4688-9eeb-ef897f3dd0cc",
                            Name: codataccounting.String("Curtis Fahey Sr."),
                        },
                    },
                    UnitAmount: 6984.45,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3e4e080a-a104-4186-ac75-9e02f3702c5c"),
                        Name: codataccounting.String("Clay Conn"),
                    },
                    Description: codataccounting.String("consequatur"),
                    DiscountAmount: codataccounting.Float64(9158.97),
                    DiscountPercentage: codataccounting.Float64(6344.86),
                    ItemRef: &shared.ItemRef{
                        ID: "d3104fa4-4707-4bf3-b5b4-4282821fdb2f",
                        Name: codataccounting.String("Velma Vandervort"),
                    },
                    Quantity: 1639.27,
                    SubTotal: codataccounting.Float64(3850.8),
                    TaxAmount: codataccounting.Float64(4833.56),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7908.48),
                        ID: codataccounting.String("71cc8d3c-d425-48d0-b58a-82c808fe2751"),
                        Name: codataccounting.String("Billy Aufderhar"),
                    },
                    TotalAmount: codataccounting.Float64(7989.54),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "449e143f-9619-4bb7-940d-5a11fa436e62",
                            Name: codataccounting.String("Faye Champlin"),
                        },
                    },
                    UnitAmount: 9421.67,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("95c9d237-397c-4785-b5db-4f500183febd"),
                        Name: codataccounting.String("Brett Kling"),
                    },
                    Description: codataccounting.String("reprehenderit"),
                    DiscountAmount: codataccounting.Float64(1612.08),
                    DiscountPercentage: codataccounting.Float64(378.34),
                    ItemRef: &shared.ItemRef{
                        ID: "6dab7500-52a5-4647-adc4-39ed8c4320f4",
                        Name: codataccounting.String("Dr. Julia Gibson"),
                    },
                    Quantity: 2790.04,
                    SubTotal: codataccounting.Float64(5208.11),
                    TaxAmount: codataccounting.Float64(4808.49),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6503.19),
                        ID: codataccounting.String("c693b94c-3b9d-4248-8d79-5aa42fc40566"),
                        Name: codataccounting.String("Roman Kassulke"),
                    },
                    TotalAmount: codataccounting.Float64(604.91),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6d212494-5081-49d7-83b1-b41844060e00",
                            Name: codataccounting.String("Carolyn Bednar I"),
                        },
                    },
                    UnitAmount: 2449.14,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("repellendus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("error"),
                        CurrencyRate: codataccounting.Float64(563.11),
                        TotalAmount: codataccounting.Float64(1094.16),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f5afd2a6-c448-446a-a9d8-9253c8962f48"),
                            Name: codataccounting.String("Franklin Prohaska"),
                        },
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(9102.24),
                        ID: codataccounting.String("4652d3c3-43d6-4177-8af4-91247725e621"),
                        Note: codataccounting.String("iste"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("aut"),
                        TotalAmount: codataccounting.Float64(5779.63),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("vero"),
                        CurrencyRate: codataccounting.Float64(6231.5),
                        TotalAmount: codataccounting.Float64(781.11),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("044a5de5-9ac7-4706-a70c-f1cf59326052"),
                            Name: codataccounting.String("Teresa VonRueden"),
                        },
                        Currency: codataccounting.String("harum"),
                        CurrencyRate: codataccounting.Float64(7414),
                        ID: codataccounting.String("426897d9-9a2d-4335-a70e-93ee6cf59f35"),
                        Note: codataccounting.String("voluptatum"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("dolorum"),
                        TotalAmount: codataccounting.Float64(6400.46),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("earum"),
                        CurrencyRate: codataccounting.Float64(6662.73),
                        TotalAmount: codataccounting.Float64(7892.14),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ae323a31-bf7b-4a1c-8977-16c802cc9e0c"),
                            Name: codataccounting.String("Marcella Monahan"),
                        },
                        Currency: codataccounting.String("dolores"),
                        CurrencyRate: codataccounting.Float64(2361.24),
                        ID: codataccounting.String("f1aa63ed-9cf1-4c85-abcb-a51ef2454a47"),
                        Note: codataccounting.String("voluptatibus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("dolorum"),
                        TotalAmount: codataccounting.Float64(8037.82),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("repellat"),
                        CurrencyRate: codataccounting.Float64(860.51),
                        TotalAmount: codataccounting.Float64(779.96),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6cdd5444-a756-4287-bc7d-d9efaf43dc62"),
                            Name: codataccounting.String("Dr. Terri Collier"),
                        },
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(2236.18),
                        ID: codataccounting.String("8f30df3d-b022-4faa-965f-b8f652ebb9d3"),
                        Note: codataccounting.String("praesentium"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("non"),
                        TotalAmount: codataccounting.Float64(5332.9),
                    },
                },
            },
            Reference: codataccounting.String("dolor"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SubTotal: 5464.65,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "molestias": map[string]interface{}{
                        "fugit": "labore",
                    },
                    "neque": map[string]interface{}{
                        "sed": "error",
                        "ratione": "facere",
                        "est": "soluta",
                    },
                },
            },
            TaxAmount: 2065,
            TotalAmount: 212.77,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(878395),
    }

    res, err := s.DirectIncomes.CreateDirectIncome(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectIncomeResponse != nil {
        // handle response
    }
}
```