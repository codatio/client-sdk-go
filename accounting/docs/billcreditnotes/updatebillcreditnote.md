# UpdateBillCreditNote
Available in: `BillCreditNotes`

Posts an updated billCreditNote to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support updating bill credit notes.

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
    req := operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            BillCreditNoteNumber: codataccounting.String("voluptas"),
            Currency: codataccounting.String("asperiores"),
            CurrencyRate: codataccounting.Float64(456.59),
            DiscountPercentage: 4090.54,
            ID: codataccounting.String("42dac7af-515c-4c41-baa6-3aae8d67864d"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b675fd5e-60b3-475e-94f6-fbee41f33317"),
                        Name: codataccounting.String("Doyle Feest"),
                    },
                    Description: codataccounting.String("laboriosam"),
                    DiscountAmount: codataccounting.Float64(583.56),
                    DiscountPercentage: codataccounting.Float64(9167.27),
                    ItemRef: &shared.ItemRef{
                        ID: "b1ea4265-55ba-43c2-8744-ed53b88f3a8d",
                        Name: codataccounting.String("Terrell Heidenreich MD"),
                    },
                    Quantity: 1488.29,
                    SubTotal: codataccounting.Float64(9679.66),
                    TaxAmount: codataccounting.Float64(1318.52),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9944.01),
                        ID: codataccounting.String("b7b194a2-76b2-4691-afe1-f08f4294e369"),
                        Name: codataccounting.String("Courtney Goldner"),
                    },
                    TotalAmount: codataccounting.Float64(9700.76),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "03e8b445-e80c-4a55-afd2-0e457e1858b6",
                                Name: codataccounting.String("Bob Mueller"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e3a5aa8e-4824-4d0a-b407-5088e5186206",
                                Name: codataccounting.String("Mrs. Tricia Mueller"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolorem"),
                            ID: "b1194b8a-bf60-43a7-9f9d-fe0ab7da8a50",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "187f86bc-173d-4689-aee9-526f8d986e88",
                            Name: codataccounting.String("Delia Parisian"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0e101256-3f94-4e29-a973-e922a57a15be",
                            Name: codataccounting.String("Meghan Batz IV"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "07e2b6e3-ab88-445f-8597-a60ff2a54a31",
                            Name: codataccounting.String("Arturo Hagenes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4a3e865e-7956-4f92-91a5-a9da660ff57b",
                            Name: codataccounting.String("Oliver Osinski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f9efc1b4-512c-4103-a648-dc2f615199eb",
                            Name: codataccounting.String("Gilberto Bechtelar"),
                        },
                    },
                    UnitAmount: 9834.27,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e6c632ca-3aed-4011-b996-312fde047717"),
                        Name: codataccounting.String("Irma Wuckert"),
                    },
                    Description: codataccounting.String("architecto"),
                    DiscountAmount: codataccounting.Float64(8571.25),
                    DiscountPercentage: codataccounting.Float64(396.5),
                    ItemRef: &shared.ItemRef{
                        ID: "17476360-a15d-4b6a-a606-59a1adeaab58",
                        Name: codataccounting.String("Gloria Skiles"),
                    },
                    Quantity: 4053.73,
                    SubTotal: codataccounting.Float64(2811.53),
                    TaxAmount: codataccounting.Float64(3210.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7139.27),
                        ID: codataccounting.String("08b61891-baa0-4fe1-ade0-08e6f8c5f350"),
                        Name: codataccounting.String("Jimmie Russel"),
                    },
                    TotalAmount: codataccounting.Float64(3732.16),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "34181430-1042-4181-bd52-08ece7e253b6",
                                Name: codataccounting.String("Jennie Gutkowski DDS"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6c6e205e-16de-4ab3-bec9-578a64584273",
                                Name: codataccounting.String("Ms. Armando Gottlieb"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "162309fb-0929-4921-aefb-9f58c4d86e68",
                                Name: codataccounting.String("Edwin Reichert III"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("vel"),
                            ID: "013f59da-757a-459e-8fef-66ef1caa3383",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "beb47737-3c8d-472f-a4d1-db1f2c431066",
                            Name: codataccounting.String("Leigh Mante"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9e1cf9e0-6e3a-4437-800a-e6b6bc9b8f75",
                            Name: codataccounting.String("Terence O'Keefe"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5a9741d3-1135-4296-9bb8-a7202611435e",
                            Name: codataccounting.String("Tracy Mills"),
                        },
                    },
                    UnitAmount: 8028.94,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2259b1ab-da8c-4070-a108-4cb0672d1ad8"),
                        Name: codataccounting.String("Daisy Tillman"),
                    },
                    Description: codataccounting.String("sint"),
                    DiscountAmount: codataccounting.Float64(4097.26),
                    DiscountPercentage: codataccounting.Float64(4218.19),
                    ItemRef: &shared.ItemRef{
                        ID: "5b85efbd-02ba-4e0b-a2d7-82259e3ea4b5",
                        Name: codataccounting.String("Lindsey Kreiger"),
                    },
                    Quantity: 1478.01,
                    SubTotal: codataccounting.Float64(2536.25),
                    TaxAmount: codataccounting.Float64(2569.16),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2010.1),
                        ID: codataccounting.String("da7ce52b-895c-4537-8645-4efb0b34896c"),
                        Name: codataccounting.String("Bethany Paucek"),
                    },
                    TotalAmount: codataccounting.Float64(7708.73),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "be2fd570-7577-4929-977d-eac646ecb573",
                                Name: codataccounting.String("Melissa Mayer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "eb1e5a2b-12eb-407f-916d-b99545fc95fa",
                                Name: codataccounting.String("Felix Mann PhD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "189dbb30-fcb3-43ea-855b-197cd44e2f52",
                                Name: codataccounting.String("Dwight Connelly"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "513bb6f4-8b65-46bc-9b35-ff2e4b27537a",
                                Name: codataccounting.String("Delbert Stehr"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ducimus"),
                            ID: "319c177d-525f-477b-914e-eb52ff785fc3",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "14d4c98e-0c2b-4b89-ab75-dad636c60050",
                            Name: codataccounting.String("Desiree Lakin"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1180f739-ae9e-4057-ab80-9e2810331f39",
                            Name: codataccounting.String("Albert Stroman"),
                        },
                    },
                    UnitAmount: 4567.04,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("voluptatem"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("rerum"),
                        CurrencyRate: codataccounting.Float64(4116.15),
                        TotalAmount: codataccounting.Float64(467.91),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7f3c93c7-3b9d-4a3f-aced-a7e23f225741"),
                            Name: codataccounting.String("May Nolan"),
                        },
                        Currency: codataccounting.String("distinctio"),
                        CurrencyRate: codataccounting.Float64(4502.24),
                        ID: codataccounting.String("544e472e-8028-457a-9b40-463a7d575f14"),
                        Note: codataccounting.String("aut"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("aut"),
                        TotalAmount: codataccounting.Float64(9116.57),
                    },
                },
            },
            RemainingCredit: 4837.53,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.BillCreditNoteStatusEnumSubmitted,
            SubTotal: 2561.14,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "possimus": map[string]interface{}{
                        "consectetur": "nesciunt",
                        "quaerat": "itaque",
                    },
                    "minus": map[string]interface{}{
                        "distinctio": "iusto",
                    },
                    "quas": map[string]interface{}{
                        "facilis": "amet",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "6a08088d-100e-4fad-a200-ef0422eb2164",
                SupplierName: codataccounting.String("optio"),
            },
            TotalAmount: 9750.95,
            TotalDiscount: 5629.48,
            TotalTaxAmount: 6394.63,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5206.78,
                    Name: "Beth Jenkins",
                },
                shared.WithholdingTaxitems{
                    Amount: 1409.57,
                    Name: "Faith Zulauf",
                },
                shared.WithholdingTaxitems{
                    Amount: 6176.57,
                    Name: "Anthony Huel",
                },
            },
        },
        BillCreditNoteID: "voluptates",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(251627),
    }

    res, err := s.BillCreditNotes.UpdateBillCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```