# CreateDirectCost
Available in: `DirectCosts`

Posts a new direct cost to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/accounting-api#/operations/get-create-directCosts-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating direct costs.

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
    req := operations.CreateDirectCostRequest{
        DirectCost: &shared.DirectCost{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("modi"),
                ID: "c0252fe3-b4b4-4db8-b778-ebb6e1d2cf50",
            },
            Currency: "aspernatur",
            CurrencyRate: codataccounting.Float64(7116.2),
            ID: codataccounting.String("afb2cbc4-635d-45e6-9da0-28c3e951a1e3"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("fda96648-9d7b-4786-b3e1-3a12a6b99249"),
                        Name: codataccounting.String("Renee Morissette"),
                    },
                    Description: codataccounting.String("corrupti"),
                    DiscountAmount: codataccounting.Float64(4996.21),
                    DiscountPercentage: codataccounting.Float64(9890.6),
                    ItemRef: &shared.ItemRef{
                        ID: "5c843836-b86b-43cd-b641-5b0449f9df13",
                        Name: codataccounting.String("Mario Torphy"),
                    },
                    Quantity: 7182.93,
                    SubTotal: codataccounting.Float64(9239.82),
                    TaxAmount: codataccounting.Float64(4493.31),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5339.88),
                        ID: codataccounting.String("bf606825-894e-4a76-bd5c-72795b785148"),
                        Name: codataccounting.String("Andre Sporer"),
                    },
                    TotalAmount: codataccounting.Float64(5825.36),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("debitis"),
                            ID: codataccounting.String("5635b33b-c0f9-470c-82fc-9f4844225e75"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("ducimus"),
                                ID: codataccounting.String("96065c0e-fa6f-493b-90a1-b8c95be1254b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("nihil"),
                                ID: codataccounting.String("39f4fe77-210d-41f6-958c-99c722d2bc0f"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("excepturi"),
                                ID: codataccounting.String("4087d9ca-ae04-42dd-bcaa-c9b4caa1cfe9"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "15df9039-07f3-4783-9983-d42e54a85466",
                            Name: codataccounting.String("Ramona Kub"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0233c147-1d51-4aaa-addf-5abd6487c5fc",
                            Name: codataccounting.String("Gayle Lehner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a00bef69-e100-4157-a30b-da7afded84a3",
                            Name: codataccounting.String("Mr. Sonya Gutmann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8e1a735a-c26a-4e33-bef9-71a8f46bca11",
                            Name: codataccounting.String("Tonya Wintheiser"),
                        },
                    },
                    UnitAmount: 4078.27,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("ipsam"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("esse"),
                        CurrencyRate: codataccounting.Float64(1133.82),
                        TotalAmount: codataccounting.Float64(1054.97),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d08cf88e-c9f7-4b99-a550-a656ed333bb0"),
                            Name: codataccounting.String("Frankie Larkin"),
                        },
                        Currency: codataccounting.String("autem"),
                        CurrencyRate: codataccounting.Float64(3351.76),
                        ID: codataccounting.String("432a986e-b7e1-44ca-9640-89150097019a"),
                        Note: codataccounting.String("eius"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("rem"),
                        TotalAmount: codataccounting.Float64(9654.94),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("blanditiis"),
                        CurrencyRate: codataccounting.Float64(5149.76),
                        TotalAmount: codataccounting.Float64(9368.8),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ce7bf904-e011-405d-b890-8162c6beb68a"),
                            Name: codataccounting.String("Melba Keebler"),
                        },
                        Currency: codataccounting.String("quidem"),
                        CurrencyRate: codataccounting.Float64(4478.48),
                        ID: codataccounting.String("d03a1480-f8de-430f-869d-810618d97e15"),
                        Note: codataccounting.String("quia"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("aspernatur"),
                        TotalAmount: codataccounting.Float64(6114.23),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("molestiae"),
                        CurrencyRate: codataccounting.Float64(3286.66),
                        TotalAmount: codataccounting.Float64(900.67),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0da80312-292c-4c61-82a7-02bb97ee102d"),
                            Name: codataccounting.String("Aaron Stehr"),
                        },
                        Currency: codataccounting.String("minima"),
                        CurrencyRate: codataccounting.Float64(9483.74),
                        ID: codataccounting.String("8e01bf33-eaab-4454-82ac-1704bf1cc9fc"),
                        Note: codataccounting.String("ex"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("dicta"),
                        TotalAmount: codataccounting.Float64(6460.23),
                    },
                },
            },
            Reference: codataccounting.String("laborum"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SubTotal: 9299.41,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptates": map[string]interface{}{
                        "quaerat": "delectus",
                        "sit": "porro",
                        "labore": "molestias",
                    },
                    "qui": map[string]interface{}{
                        "ullam": "nihil",
                        "ut": "incidunt",
                        "quibusdam": "doloremque",
                    },
                },
            },
            TaxAmount: 5244.63,
            TotalAmount: 6772.03,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(129394),
    }

    res, err := s.DirectCosts.CreateDirectCost(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectCostResponse != nil {
        // handle response
    }
}
```