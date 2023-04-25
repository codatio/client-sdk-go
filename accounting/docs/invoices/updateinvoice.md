# UpdateInvoice
Available in: `Invoices`

Posts an updated invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support updating invoices.

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
    req := operations.UpdateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(8977.85),
            AdditionalTaxPercentage: codataccounting.Float64(5673.2),
            AmountDue: 721.26,
            Currency: codataccounting.String("tempora"),
            CurrencyRate: codataccounting.Float64(2813.81),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("perspiciatis"),
                ID: "8a9ba460-addf-4de4-90c3-7daa9182a49d",
            },
            DiscountPercentage: codataccounting.Float64(5768.41),
            DueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("625d3caf-fc19-48ee-a445-2792bcd440ea"),
            InvoiceNumber: codataccounting.String("occaecati"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("becce048-6de0-4d56-973b-005503e8dc62"),
                        Name: codataccounting.String("Ora Wuckert"),
                    },
                    Description: codataccounting.String("impedit"),
                    DiscountAmount: codataccounting.Float64(4366.89),
                    DiscountPercentage: codataccounting.Float64(3524.16),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "675f5b70-e3e4-4cfc-86a9-1ec52624d000",
                        Name: codataccounting.String("Rhonda VonRueden"),
                    },
                    Quantity: 3242.38,
                    SubTotal: codataccounting.Float64(7743.39),
                    TaxAmount: codataccounting.Float64(8885.73),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6571.73),
                        ID: codataccounting.String("11ac53eb-b658-47f3-8041-4c5b9acee400"),
                        Name: codataccounting.String("Noah Medhurst"),
                    },
                    TotalAmount: codataccounting.Float64(1671.44),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "af1b025f-1d14-4718-86fa-2fad0c06c5d9",
                                Name: codataccounting.String("Gail King"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "dd14fc43-b70b-4ca8-8fa7-0c43351c3dd1",
                                Name: codataccounting.String("Kerry Lesch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f75f4f23-f1c0-4a58-ac3a-e7d7b67feef5",
                                Name: codataccounting.String("Eric Gibson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95b1dbec-eff7-4c4b-956e-9278275eea76",
                                Name: codataccounting.String("Carl Kozey"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("praesentium"),
                            ID: "063f799b-7956-4c0b-8fa0-bb20a40e7c4a",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "40642726-57b0-41a0-bc08-fd3921c25793",
                            Name: codataccounting.String("Mercedes Kemmer V"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a3efa46d-366d-4fa1-811a-091b3ec8b538",
                            Name: codataccounting.String("Bonnie Stokes MD"),
                        },
                    },
                    UnitAmount: 5648.16,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d14fe72e-521f-4903-83df-c338397fffa6"),
                        Name: codataccounting.String("Justin Stroman"),
                    },
                    Description: codataccounting.String("consequatur"),
                    DiscountAmount: codataccounting.Float64(5788.09),
                    DiscountPercentage: codataccounting.Float64(9.47),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "fc157ac9-fe19-461c-a9be-41c869dd7d97",
                        Name: codataccounting.String("Ms. Geneva Skiles"),
                    },
                    Quantity: 1382.61,
                    SubTotal: codataccounting.Float64(154.46),
                    TaxAmount: codataccounting.Float64(498.92),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6470.89),
                        ID: codataccounting.String("58ffd296-7df8-4fd8-82a8-e60be620cd9c"),
                        Name: codataccounting.String("Brandi Williamson"),
                    },
                    TotalAmount: codataccounting.Float64(419.17),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c3752512-beae-41d8-becc-5fdcea8e7a88",
                                Name: codataccounting.String("Frances Bosco"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2cda6d77-c1d8-4606-a237-d4227866db8a",
                                Name: codataccounting.String("Clara Mertz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("provident"),
                            ID: "84511cc7-5e4f-40c0-84b5-bb758cc94562",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "069685fc-d1a1-473d-84bb-e24f29834afb",
                            Name: codataccounting.String("Viola Dicki"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6285d4a2-9aaa-41e1-a915-6f7d2ee20950",
                            Name: codataccounting.String("Mrs. Whitney Wuckert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "93e94480-ca37-4fb1-8789-032ac333172e",
                            Name: codataccounting.String("Angelina Sporer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ec74ba7e-88dd-4b36-bd1c-cc341c865734",
                            Name: codataccounting.String("Miss Charlotte Yundt"),
                        },
                    },
                    UnitAmount: 2786.03,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0fb4ab44-1c3a-409e-b639-95d808bbe794"),
                        Name: codataccounting.String("Joann Herman"),
                    },
                    Description: codataccounting.String("quo"),
                    DiscountAmount: codataccounting.Float64(3668.22),
                    DiscountPercentage: codataccounting.Float64(3290.47),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "0a1c426b-59c8-4366-bdcc-135582c1b855",
                        Name: codataccounting.String("Dwight Lindgren"),
                    },
                    Quantity: 6061.92,
                    SubTotal: codataccounting.Float64(8896.83),
                    TaxAmount: codataccounting.Float64(9407.66),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6238.98),
                        ID: codataccounting.String("32e9000a-13ad-4812-8208-efd23411898e"),
                        Name: codataccounting.String("Sylvia Lang"),
                    },
                    TotalAmount: codataccounting.Float64(9076.15),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "be8baeba-bb79-4453-ae90-351bb9763172",
                                Name: codataccounting.String("Melody Kreiger"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5a5365a7-9f15-4271-b01c-0d361fed8dc5",
                                Name: codataccounting.String("Tommie Wilkinson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "53e9089e-871f-4db4-9697-bdd9c985e437",
                                Name: codataccounting.String("Marjorie Nitzsche"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "72d9edd7-85be-45e7-afe5-5297ba6281f4",
                                Name: codataccounting.String("Tricia Ebert"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolor"),
                            ID: "394a68cc-80d3-40ff-b216-4d0a91fe9d96",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3b89e000-9c66-492d-a7b3-562201a6aab4",
                            Name: codataccounting.String("Noah Kirlin MD"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b908d4e3-0491-4aa3-9d4a-839f03bab77b",
                            Name: codataccounting.String("Ryan Labadie I"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "13984507-e0e3-49c7-a23e-cb0604652e23",
                            Name: codataccounting.String("Norman Stiedemann"),
                        },
                    },
                    UnitAmount: 3772.06,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("quis"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("vero"),
                        CurrencyRate: codataccounting.Float64(6032.05),
                        TotalAmount: codataccounting.Float64(8413.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e8f7f002-d198-46aa-99d3-a1d32329e458"),
                            Name: codataccounting.String("Maureen Veum"),
                        },
                        Currency: codataccounting.String("qui"),
                        CurrencyRate: codataccounting.Float64(6377.7),
                        ID: codataccounting.String("d6bb10e2-55fd-4c48-8d6e-3308675cbf18"),
                        Note: codataccounting.String("voluptas"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("molestias"),
                        TotalAmount: codataccounting.Float64(3424.33),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("vel"),
                        CurrencyRate: codataccounting.Float64(6468.22),
                        TotalAmount: codataccounting.Float64(4981.83),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e82cdf9d-0fc2-482c-a66a-f3c3f5589bea"),
                            Name: codataccounting.String("Marcella Daniel"),
                        },
                        Currency: codataccounting.String("voluptates"),
                        CurrencyRate: codataccounting.Float64(2652.54),
                        ID: codataccounting.String("1e2ca848-22e5-413f-ad9d-2ad37c309907"),
                        Note: codataccounting.String("esse"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("nobis"),
                        TotalAmount: codataccounting.Float64(640.01),
                    },
                },
            },
            SalesOrderRefs: []string{
                "rerum",
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.InvoiceStatusEnumSubmitted,
            SubTotal: codataccounting.Float64(5362.23),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "iste": map[string]interface{}{
                        "veritatis": "commodi",
                    },
                    "amet": map[string]interface{}{
                        "autem": "nihil",
                        "repellendus": "non",
                        "rem": "voluptatum",
                        "vel": "quae",
                    },
                },
            },
            TotalAmount: 3358.37,
            TotalDiscount: codataccounting.Float64(2942.68),
            TotalTaxAmount: 2316.11,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 470,
                    Name: "Jeff Bailey",
                },
                shared.WithholdingTaxitems{
                    Amount: 7879.41,
                    Name: "Traci Wolf",
                },
                shared.WithholdingTaxitems{
                    Amount: 251.21,
                    Name: "Marcia Howe",
                },
                shared.WithholdingTaxitems{
                    Amount: 7295.54,
                    Name: "Debra Purdy",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        InvoiceID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(590737),
    }

    res, err := s.Invoices.UpdateInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateInvoiceResponse != nil {
        // handle response
    }
}
```