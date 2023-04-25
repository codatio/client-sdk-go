# UpdateCreditNote
Available in: `CreditNotes`

Posts an updated credit note to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support updating a credit note.

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
    req := operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(613),
            AdditionalTaxPercentage: codataccounting.Float64(3298.36),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            CreditNoteNumber: codataccounting.String("sint"),
            Currency: codataccounting.String("minus"),
            CurrencyRate: codataccounting.Float64(5674.41),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("porro"),
                ID: "3f567e0e-2527-465b-9d62-fcdace1f0121",
            },
            DiscountPercentage: 3888.89,
            ID: codataccounting.String("ce2239e8-f25c-4d0d-99d9-59f439e39266"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("bd95f7aa-2b24-4113-a95d-1e6698fcc459"),
                        Name: codataccounting.String("Bonnie Carroll"),
                    },
                    Description: codataccounting.String("dolores"),
                    DiscountAmount: codataccounting.Float64(5658.09),
                    DiscountPercentage: codataccounting.Float64(4624.26),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "76763342-5403-48bf-b597-1e9819055738",
                        Name: codataccounting.String("Benny Ward"),
                    },
                    Quantity: 6636.42,
                    SubTotal: codataccounting.Float64(7673.61),
                    TaxAmount: codataccounting.Float64(4498.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9508.94),
                        ID: codataccounting.String("da39594d-66bc-42ae-8806-32b9954b6fa2"),
                        Name: codataccounting.String("Ruth Kemmer"),
                    },
                    TotalAmount: codataccounting.Float64(5646.65),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "28553cb1-0006-4bef-8921-ec2053b74936",
                                Name: codataccounting.String("Brandi Schaefer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e0f2bf19-588d-440d-83f3-deba297be3e9",
                                Name: codataccounting.String("Maryann Rolfson PhD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f868fd52-405c-4b33-9d49-2f4f127fb0e0",
                                Name: codataccounting.String("Moses Brekke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("fugit"),
                            ID: "17978d0a-cca7-47ae-b7b7-021a52046b64",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "9fb0e67e-094f-4dfe-9554-0ef53a34a1b8",
                            Name: codataccounting.String("Lamar Mertz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1adc05d8-5ae2-4dfb-b0fb-3874290d3365",
                            Name: codataccounting.String("Janet Torphy"),
                        },
                    },
                    UnitAmount: 993.43,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6ef89451-bd76-4eee-b518-c4da1fad3551"),
                        Name: codataccounting.String("Faith Baumbach"),
                    },
                    Description: codataccounting.String("dolore"),
                    DiscountAmount: codataccounting.Float64(9328.83),
                    DiscountPercentage: codataccounting.Float64(3163.35),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b72f0f54-8568-4a04-a4e0-0a1d6eb94346",
                        Name: codataccounting.String("Mrs. Dana Swaniawski V"),
                    },
                    Quantity: 3105.42,
                    SubTotal: codataccounting.Float64(9999.32),
                    TaxAmount: codataccounting.Float64(7093),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7284.66),
                        ID: codataccounting.String("a5cceff5-cb01-4fe5-9e52-8a45ac82b85f"),
                        Name: codataccounting.String("Dr. Marco Schulist"),
                    },
                    TotalAmount: codataccounting.Float64(7142.4),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "8da4127d-d597-4ff4-b11a-a1bc74b86cec",
                                Name: codataccounting.String("Kelly Goyette"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7b4848bd-6a6f-4044-9d2c-3b808094373e",
                                Name: codataccounting.String("Katie Bailey"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9bebbad0-2f25-486b-8f15-2558daa95be6",
                                Name: codataccounting.String("Jody Altenwerth"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("enim"),
                            ID: "6c354aa4-32b4-47e1-b63c-5208c23e9802",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "2f0d45eb-4a8b-4674-ae5c-fc18edc7f787",
                            Name: codataccounting.String("Travis Dach I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3d3ed0c5-670e-4f42-bd3c-9f1cc503f6c3",
                            Name: codataccounting.String("Preston Runte MD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6290f957-f385-4189-ad7e-f807aae03f33",
                            Name: codataccounting.String("Dana Kuhlman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b9de4032-ba26-4fd3-a8ba-9216bcb41583",
                            Name: codataccounting.String("Marianne Koelpin"),
                        },
                    },
                    UnitAmount: 2708.4,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1723133e-dc04-46bc-9163-bbca49227c42"),
                        Name: codataccounting.String("Brandon Cummings"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(1931.99),
                    DiscountPercentage: codataccounting.Float64(3126.17),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "0495c5db-b3c5-47c1-a498-1e8aa257ddc1",
                        Name: codataccounting.String("Dennis Considine"),
                    },
                    Quantity: 8734.29,
                    SubTotal: codataccounting.Float64(8953.49),
                    TaxAmount: codataccounting.Float64(4286.45),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2501.01),
                        ID: codataccounting.String("bfcc5469-d401-45df-a796-206bef2b0a3e"),
                        Name: codataccounting.String("Miss Lori Sawayn"),
                    },
                    TotalAmount: codataccounting.Float64(173.42),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0e9aac2e-9135-4586-918f-9f97a4bfad2b",
                                Name: codataccounting.String("Erik Stehr"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quo"),
                            ID: "a84ad99b-41d6-4124-b531-870cf68b03ad",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1bd43d1f-0cb0-4a00-83eb-22d9b3a70d94",
                            Name: codataccounting.String("Grant Oberbrunner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c57d1fed-c205-40d3-8dc3-ce185472f9ee",
                            Name: codataccounting.String("Natasha Boyer"),
                        },
                    },
                    UnitAmount: 6250.09,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8be3444e-ac8b-43a2-875c-6c1fe606d07d"),
                        Name: codataccounting.String("Joanna Mueller"),
                    },
                    Description: codataccounting.String("nihil"),
                    DiscountAmount: codataccounting.Float64(6406.81),
                    DiscountPercentage: codataccounting.Float64(9201.94),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "50c16661-a1d9-4136-a7e8-d53213f3f658",
                        Name: codataccounting.String("Lynn Considine"),
                    },
                    Quantity: 4758.59,
                    SubTotal: codataccounting.Float64(4218.34),
                    TaxAmount: codataccounting.Float64(3058.33),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8030.15),
                        ID: codataccounting.String("59f0a56c-ebca-4da2-9ca7-9181c9567166"),
                        Name: codataccounting.String("Brooke Hand MD"),
                    },
                    TotalAmount: codataccounting.Float64(3255.27),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "65163a36-3851-42ab-a521-b9f2e072467b",
                                Name: codataccounting.String("Miss Homer Gislason"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "05fab0d6-50ed-4f22-a94d-20ec90ea41d1",
                                Name: codataccounting.String("Leroy Huels"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("praesentium"),
                            ID: "5156fff7-3fdf-454f-9d5e-a9543398dafb",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "a8d63388-e4d8-4039-aa5f-9b18a244fd61",
                            Name: codataccounting.String("Charles Emmerich"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cd38ed0d-c671-4dc7-b1e3-af15920c90d1",
                            Name: codataccounting.String("Mr. Clifford Morissette"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2bd89c8a-3263-49da-9b7b-6902b881a94f",
                            Name: codataccounting.String("Emma Dickinson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4a8f0af8-c691-4d73-ad9f-baf9476a2ae8",
                            Name: codataccounting.String("Benny Ryan DDS"),
                        },
                    },
                    UnitAmount: 5567.19,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("culpa"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("nostrum"),
                        CurrencyRate: codataccounting.Float64(1196.85),
                        TotalAmount: codataccounting.Float64(1793.01),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("c7378489-3075-40a0-8e96-6ec736d43194"),
                            Name: codataccounting.String("Belinda Lockman"),
                        },
                        Currency: codataccounting.String("atque"),
                        CurrencyRate: codataccounting.Float64(2239.23),
                        ID: codataccounting.String("c92398ed-3d3a-4b7c-a3c5-ca8649a70cfd"),
                        Note: codataccounting.String("veniam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("pariatur"),
                        TotalAmount: codataccounting.Float64(4175.19),
                    },
                },
            },
            RemainingCredit: 6115.25,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CreditNoteStatusEnumPaid,
            SubTotal: 6125.84,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "in": map[string]interface{}{
                        "voluptatem": "voluptas",
                    },
                    "magnam": map[string]interface{}{
                        "quae": "ipsa",
                        "iure": "voluptate",
                    },
                    "illum": map[string]interface{}{
                        "perspiciatis": "accusamus",
                    },
                },
            },
            TotalAmount: 6755.33,
            TotalDiscount: 5296.77,
            TotalTaxAmount: 1880.08,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 2619.53,
                    Name: "Russell Weber II",
                },
                shared.WithholdingTaxitems{
                    Amount: 7320.16,
                    Name: "Blake Connelly V",
                },
                shared.WithholdingTaxitems{
                    Amount: 3475.83,
                    Name: "Lynn Hahn",
                },
                shared.WithholdingTaxitems{
                    Amount: 6009.33,
                    Name: "Veronica Schultz",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "et",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(510079),
    }

    res, err := s.CreditNotes.UpdateCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```