# CreateCreditNote
Available in: `CreditNotes`

Push credit note


Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating a credit note.

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
    req := operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(4884.33),
            AdditionalTaxPercentage: codataccounting.Float64(3427.72),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            CreditNoteNumber: codataccounting.String("maiores"),
            Currency: codataccounting.String("veritatis"),
            CurrencyRate: codataccounting.Float64(4226.1),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("earum"),
                ID: "56d385a3-c4ac-4631-b99e-26ced8f9fdb9",
            },
            DiscountPercentage: 2913.44,
            ID: codataccounting.String("10f63bbf-8178-437b-81af-dd788624189e"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("44873f50-33f1-49db-b125-ce4152eab9cd"),
                        Name: codataccounting.String("Sonja Hansen"),
                    },
                    Description: codataccounting.String("eius"),
                    DiscountAmount: codataccounting.Float64(6253.78),
                    DiscountPercentage: codataccounting.Float64(4274.61),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a0e123b7-847e-4c59-a1f6-7f3c4cce4b6d",
                        Name: codataccounting.String("Vicki Mayer"),
                    },
                    Quantity: 9570.32,
                    SubTotal: codataccounting.Float64(2322.09),
                    TaxAmount: codataccounting.Float64(7574.94),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3534.24),
                        ID: codataccounting.String("74750135-7e44-4f51-b8b0-84c3197e193a"),
                        Name: codataccounting.String("Eva Hettinger"),
                    },
                    TotalAmount: codataccounting.Float64(4990.05),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "94874c2d-5cc4-4972-a33e-66bd8fe5d00b",
                                Name: codataccounting.String("Darren Monahan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "20387320-590c-4cc1-8964-00313b3e5044",
                                Name: codataccounting.String("Gene Herman"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "72dc4077-d0cc-43f4-88ef-c15ceb4d6e1e",
                                Name: codataccounting.String("Alonzo Bartell"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5aedf2ac-ab58-4b99-9c92-6ddb589461e7",
                                Name: codataccounting.String("Beverly Block"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("recusandae"),
                            ID: "6d9502f0-ea93-40b6-9f7a-c2f72f885009",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "91160820-7888-4ec6-a183-bfe9659eb40e",
                            Name: codataccounting.String("Willie Hodkiewicz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "75b0b532-a4da-437c-baaf-4452c4842c9b",
                            Name: codataccounting.String("Jodi Stroman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "dafe81a8-8f44-4445-b3fe-cd47353f63c8",
                            Name: codataccounting.String("Cynthia Morissette"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9aa69cd5-fbcf-479d-a18a-7822bf95894e",
                            Name: codataccounting.String("Miss Maxine Kemmer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b55f9e5d-751c-49fe-8f75-02bfdc345084",
                            Name: codataccounting.String("Josefina Borer"),
                        },
                    },
                    UnitAmount: 2943.14,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("456379f3-fb27-4e21-b862-657b36fc6b9f"),
                        Name: codataccounting.String("Nina Kshlerin"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(1676.13),
                    DiscountPercentage: codataccounting.Float64(3457.46),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c67641a8-312e-4504-bb4c-21ccb423abcd",
                        Name: codataccounting.String("Arturo Bosco"),
                    },
                    Quantity: 6654.27,
                    SubTotal: codataccounting.Float64(7160.24),
                    TaxAmount: codataccounting.Float64(8546.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8621.67),
                        ID: codataccounting.String("88e71f6c-4825-42d7-b71e-7fd074009ef8"),
                        Name: codataccounting.String("Carlos Morissette"),
                    },
                    TotalAmount: codataccounting.Float64(717.41),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d7097b5d-a08c-457f-a6c7-8a216e19bafe",
                                Name: codataccounting.String("Ms. Gerard Hudson II"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "98140b64-ff8a-4e17-8ef0-3b5f37e4aa86",
                                Name: codataccounting.String("Tim Hamill"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "66732aa5-dcb6-4682-8b70-f8cfd5fb6e91",
                                Name: codataccounting.String("Alejandro Pacocha"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "74846e2c-3309-4db0-936d-9e75ca006f53",
                                Name: codataccounting.String("Mr. Jeremy Sanford"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quia"),
                            ID: "5a8bf92f-9742-48ad-9a9f-8bf822112535",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "98387f7a-79cd-472c-9248-4da21729f2ac",
                            Name: codataccounting.String("Amanda Tromp"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "25f1169a-c1e4-41d8-a23c-23e34f2dfa4a",
                            Name: codataccounting.String("Claire Koepp"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "de922151-fe17-4120-9985-3e9f543d8544",
                            Name: codataccounting.String("Katrina Thompson"),
                        },
                    },
                    UnitAmount: 1351.1,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4460443b-c154-4188-82f5-6e85da7832ea"),
                        Name: codataccounting.String("Ms. Ismael Hodkiewicz"),
                    },
                    Description: codataccounting.String("nesciunt"),
                    DiscountAmount: codataccounting.Float64(6902.11),
                    DiscountPercentage: codataccounting.Float64(454.45),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "d51a44bf-01ba-4d87-86d4-6082bfbdc41f",
                        Name: codataccounting.String("Maurice Strosin"),
                    },
                    Quantity: 1528.07,
                    SubTotal: codataccounting.Float64(6511.36),
                    TaxAmount: codataccounting.Float64(9016.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2893.73),
                        ID: codataccounting.String("fb5cb35d-1763-48f1-adb7-8359ecc5cb86"),
                        Name: codataccounting.String("Marta Macejkovic"),
                    },
                    TotalAmount: codataccounting.Float64(3563.51),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0ba73810-e4fe-4444-b297-cd3b1dd3bbce",
                                Name: codataccounting.String("Debbie Kozey"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "684eff50-126d-471c-bfbd-0eb74b842195",
                                Name: codataccounting.String("Melody Grady"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "d3c43159-d33e-4595-bc00-1139863aa41e",
                                Name: codataccounting.String("Dr. Kara Feil"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("explicabo"),
                            ID: "f1fcb51c-9a41-4ffb-a9cb-d795ee65e076",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7abf616e-a5c7-4164-9934-b90f2e09d19d",
                            Name: codataccounting.String("Dr. Lorene Runte"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2e105944-b935-4d23-ba72-f90849d6aed4",
                            Name: codataccounting.String("Ramiro Rowe"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "537cd922-2c9f-4f57-891a-abfa2e761f0c",
                            Name: codataccounting.String("Troy Stroman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6ef1031e-6899-4f0c-a001-e22cd55cc058",
                            Name: codataccounting.String("Hattie Botsford"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d76d971f-c820-4c65-b037-bb8e0cc88518",
                            Name: codataccounting.String("Rochelle Grant"),
                        },
                    },
                    UnitAmount: 96.87,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("ut"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("hic"),
                        CurrencyRate: codataccounting.Float64(1531.31),
                        TotalAmount: codataccounting.Float64(5461.33),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("c5dddb46-aa1c-4fd6-9828-da0131911296"),
                            Name: codataccounting.String("Beth Keeling"),
                        },
                        Currency: codataccounting.String("cumque"),
                        CurrencyRate: codataccounting.Float64(701.73),
                        ID: codataccounting.String("d81f2904-2f56-49b7-aff0-ea2216cbe071"),
                        Note: codataccounting.String("harum"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("cumque"),
                        TotalAmount: codataccounting.Float64(697.88),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("ex"),
                        CurrencyRate: codataccounting.Float64(2198.04),
                        TotalAmount: codataccounting.Float64(8847.04),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("279a3b08-4da9-4925-bd04-f40847a742d8"),
                            Name: codataccounting.String("Charlotte McGlynn"),
                        },
                        Currency: codataccounting.String("tempore"),
                        CurrencyRate: codataccounting.Float64(8142.92),
                        ID: codataccounting.String("eecf6b99-bc63-4562-abfd-f55c294c060b"),
                        Note: codataccounting.String("accusantium"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("ea"),
                        TotalAmount: codataccounting.Float64(6715.04),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(1448.56),
                        TotalAmount: codataccounting.Float64(5505.79),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7764eef6-d0c6-4d6e-99c7-3dd634571509"),
                            Name: codataccounting.String("Nelson Walker"),
                        },
                        Currency: codataccounting.String("doloremque"),
                        CurrencyRate: codataccounting.Float64(8440.94),
                        ID: codataccounting.String("3c5a1f9c-242c-47b6-aa1f-30c73df5b671"),
                        Note: codataccounting.String("excepturi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("voluptatum"),
                        TotalAmount: codataccounting.Float64(6108.73),
                    },
                },
            },
            RemainingCredit: 509.03,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CreditNoteStatusEnumPartiallyPaid,
            SubTotal: 3021.47,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "est": map[string]interface{}{
                        "nobis": "expedita",
                        "modi": "adipisci",
                    },
                },
            },
            TotalAmount: 5401.92,
            TotalDiscount: 8712.98,
            TotalTaxAmount: 5236.07,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 7311.86,
                    Name: "Vanessa Beahan",
                },
                shared.WithholdingTaxitems{
                    Amount: 1195.43,
                    Name: "Ron Goodwin",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(195946),
    }

    res, err := s.CreditNotes.CreateCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCreditNoteResponse != nil {
        // handle response
    }
}
```