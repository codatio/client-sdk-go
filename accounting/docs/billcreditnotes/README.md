# BillCreditNotes

## Overview

Bill credit notes

### Available Operations

* [CreateBillCreditNote](#createbillcreditnote) - Create bill credit note
* [GetBillCreditNote](#getbillcreditnote) - Get bill credit note
* [GetCreateUpdateBillCreditNotesModel](#getcreateupdatebillcreditnotesmodel) - Get create/update bill credit note model
* [ListBillCreditNotes](#listbillcreditnotes) - List bill credit notes
* [UpdateBillCreditNote](#updatebillcreditnote) - Update bill credit note

## CreateBillCreditNote

Posts a new billCreditNote to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating bill credit notes.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("facilis"),
            BillCreditNoteNumber: codataccounting.String("tempore"),
            Currency: codataccounting.String("labore"),
            CurrencyRate: codataccounting.Float64(9621.89),
            DiscountPercentage: 4332.88,
            ID: codataccounting.String("3c969e9a-3efa-477d-bb14-cd66ae395efb"),
            IssueDate: codataccounting.String("provident"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a88f3a66-9970-474b-a446-9b6e21419598"),
                        Name: codataccounting.String("Kenneth O'Hara"),
                    },
                    Description: codataccounting.String("ad"),
                    DiscountAmount: codataccounting.Float64(4314.18),
                    DiscountPercentage: codataccounting.Float64(2212.62),
                    ItemRef: &shared.ItemRef{
                        ID: "e2516fe4-c8b7-411e-9b7f-d2ed028921cd",
                        Name: codataccounting.String("Simon Jenkins"),
                    },
                    Quantity: 4071.83,
                    SubTotal: codataccounting.Float64(332.22),
                    TaxAmount: codataccounting.Float64(691.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9825.75),
                        ID: codataccounting.String("b576b0d5-f0d3-40c5-bbb2-587053202c73"),
                        Name: codataccounting.String("Dean Welch"),
                    },
                    TotalAmount: codataccounting.Float64(7044.15),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0c28909b-3fe4-49a8-99cb-f48633323f9b",
                                Name: codataccounting.String("Marian Wisozk"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4100674e-bf69-4280-91ba-77a89ebf737a",
                                Name: codataccounting.String("Mrs. Ray Collins"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e5e6a95d-8a0d-4446-8e2a-f7a73cf3be45",
                                Name: codataccounting.String("Jeannie Leannon MD"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("neque"),
                            ID: "26b5a734-29cd-4b1a-8422-bb679d232271",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f0cbb1e3-1b8b-490f-b443-a1108e0adcf4",
                            Name: codataccounting.String("Ms. Terrance Davis"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fce953f7-3ef7-4fbc-babd-74dd39c0f5d2",
                            Name: codataccounting.String("Elijah Wyman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "70a45626-d436-4813-b16d-9f5fce6c5561",
                            Name: codataccounting.String("Rosemary Ryan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "250fb008-c42e-4141-aac3-66c8dd6b1442",
                            Name: codataccounting.String("Jose Kreiger"),
                        },
                    },
                    UnitAmount: 2621.18,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("778a7bd4-66d2-48c1-8ab3-cdca4251904e"),
                        Name: codataccounting.String("Kelly Donnelly"),
                    },
                    Description: codataccounting.String("recusandae"),
                    DiscountAmount: codataccounting.Float64(446.12),
                    DiscountPercentage: codataccounting.Float64(7151.79),
                    ItemRef: &shared.ItemRef{
                        ID: "c7178e47-96f2-4a70-8688-282aa482562f",
                        Name: codataccounting.String("Norma Christiansen"),
                    },
                    Quantity: 5438.06,
                    SubTotal: codataccounting.Float64(922.6),
                    TaxAmount: codataccounting.Float64(4569.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9105.45),
                        ID: codataccounting.String("e17cbe61-e6b7-4b95-bc0a-b3c20c4f3789"),
                        Name: codataccounting.String("Ismael Lynch DVM"),
                    },
                    TotalAmount: codataccounting.Float64(6216.79),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "dd2efd12-1aa6-4f1e-a74b-db04f1575608",
                                Name: codataccounting.String("Rosemarie Jacobs"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a19f1d17-0513-439d-8808-6a1840394c26",
                                Name: codataccounting.String("Marian Buckridge"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3f5f0642-dac7-4af5-95cc-413aa63aae8d",
                                Name: codataccounting.String("Dora Luettgen"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("possimus"),
                            ID: "bb675fd5-e60b-4375-ad4f-6fbee41f3331",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e35b60eb-1ea4-4265-95ba-3c28744ed53b",
                            Name: codataccounting.String("Morris Weissnat"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d8f5c0b2-f2fb-47b1-94a2-76b26916fe1f",
                            Name: codataccounting.String("Naomi Wuckert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "94e3698f-447f-4603-a8b4-45e80ca55efd",
                            Name: codataccounting.String("Deborah Turcotte"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7e1858b6-a89f-4be3-a5aa-8e4824d0ab40",
                            Name: codataccounting.String("Brittany Bailey"),
                        },
                    },
                    UnitAmount: 9221.12,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("51862065-e904-4f3b-9194-b8abf603a79f"),
                        Name: codataccounting.String("Marcos Windler MD"),
                    },
                    Description: codataccounting.String("quidem"),
                    DiscountAmount: codataccounting.Float64(4406.66),
                    DiscountPercentage: codataccounting.Float64(8136.79),
                    ItemRef: &shared.ItemRef{
                        ID: "a8a50ce1-87f8-46bc-973d-689eee9526f8",
                        Name: codataccounting.String("Jeremiah Kuvalis"),
                    },
                    Quantity: 5421.29,
                    SubTotal: codataccounting.Float64(5413.81),
                    TaxAmount: codataccounting.Float64(1209.19),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9233.06),
                        ID: codataccounting.String("ad4f0e10-1256-43f9-8e29-e973e922a57a"),
                        Name: codataccounting.String("Ana Predovic"),
                    },
                    TotalAmount: codataccounting.Float64(8784.93),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "60807e2b-6e3a-4b88-85f0-597a60ff2a54",
                                Name: codataccounting.String("Danny Berge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quaerat"),
                            ID: "764a3e86-5e79-456f-9251-a5a9da660ff5",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "faad4f9e-fc1b-4451-ac10-32648dc2f615",
                            Name: codataccounting.String("Misty McKenzie"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d0e9fe6c-632c-4a3a-ad01-17996312fde0",
                            Name: codataccounting.String("Ms. Marcia Kozey"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8ff61d01-7476-4360-a15d-b6a660659a1a",
                            Name: codataccounting.String("Pat O'Keefe"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5851d6c6-45b0-48b6-9891-baa0fe1ade00",
                            Name: codataccounting.String("Darin Jakubowski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c5f350d8-cdb5-4a34-9814-3010421813d5",
                            Name: codataccounting.String("Cynthia Macejkovic"),
                        },
                    },
                    UnitAmount: 8849.52,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("esse"),
            Note: codataccounting.String("necessitatibus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("veniam"),
                        Currency: codataccounting.String("nesciunt"),
                        CurrencyRate: codataccounting.Float64(7129.27),
                        TotalAmount: codataccounting.Float64(4329.84),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("68451c6c-6e20-45e1-adea-b3fec9578a64"),
                            Name: codataccounting.String("Brandy Gibson"),
                        },
                        Currency: codataccounting.String("nesciunt"),
                        CurrencyRate: codataccounting.Float64(6817.4),
                        ID: codataccounting.String("8418d162-309f-4b09-a992-1aefb9f58c4d"),
                        Note: codataccounting.String("quos"),
                        PaidOnDate: codataccounting.String("commodi"),
                        Reference: codataccounting.String("itaque"),
                        TotalAmount: codataccounting.Float64(4156.08),
                    },
                },
            },
            RemainingCredit: 5207.61,
            SourceModifiedDate: codataccounting.String("earum"),
            Status: shared.BillCreditNoteStatusEnumDraft,
            SubTotal: 7221.84,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptatem": map[string]interface{}{
                        "vel": "alias",
                        "quasi": "non",
                    },
                    "maiores": map[string]interface{}{
                        "sint": "nulla",
                        "deserunt": "esse",
                    },
                    "nemo": map[string]interface{}{
                        "est": "quis",
                        "sint": "accusamus",
                    },
                    "impedit": map[string]interface{}{
                        "necessitatibus": "asperiores",
                        "ex": "voluptas",
                        "debitis": "delectus",
                        "quae": "minus",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "aa3383c2-beb4-4773-b3c8-d72f64d1db1f",
                SupplierName: codataccounting.String("quia"),
            },
            TotalAmount: 7820.9,
            TotalDiscount: 3041.98,
            TotalTaxAmount: 2470.45,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 365.61,
                    Name: "Rosemary Breitenberg",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(431994),
    }

    res, err := s.BillCreditNotes.CreateBillCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

## GetBillCreditNote

Gets a single billCreditNote corresponding to the given ID.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetBillCreditNoteRequest{
        BillCreditNoteID: "velit",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.BillCreditNotes.GetBillCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNote != nil {
        // handle response
    }
}
```

## GetCreateUpdateBillCreditNotesModel

Get create/update bill credit note model.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating and updating bill credit notes.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateUpdateBillCreditNotesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BillCreditNotes.GetCreateUpdateBillCreditNotesModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## ListBillCreditNotes

Gets a list of all bill credit notes for a company, with pagination

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ut"),
    }

    res, err := s.BillCreditNotes.ListBillCreditNotes(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNotes != nil {
        // handle response
    }
}
```

## UpdateBillCreditNote

Posts an updated billCreditNote to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support updating bill credit notes.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("perspiciatis"),
            BillCreditNoteNumber: codataccounting.String("earum"),
            Currency: codataccounting.String("dicta"),
            CurrencyRate: codataccounting.Float64(7722.66),
            DiscountPercentage: 9758.84,
            ID: codataccounting.String("9e06e3a4-3700-40ae-ab6b-c9b8f759eac5"),
            IssueDate: codataccounting.String("corporis"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9741d311-3529-465b-b8a7-202611435e13"),
                        Name: codataccounting.String("Orville Ratke"),
                    },
                    Description: codataccounting.String("quia"),
                    DiscountAmount: codataccounting.Float64(3421.37),
                    DiscountPercentage: codataccounting.Float64(6057.12),
                    ItemRef: &shared.ItemRef{
                        ID: "b1abda8c-070e-4108-8cb0-672d1ad879ee",
                        Name: codataccounting.String("Kirk Jast"),
                    },
                    Quantity: 7029.52,
                    SubTotal: codataccounting.Float64(5156.38),
                    TaxAmount: codataccounting.Float64(3572.07),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8890.6),
                        ID: codataccounting.String("fbd02bae-0be2-4d78-a259-e3ea4b5197f9"),
                        Name: codataccounting.String("Elaine Gerhold"),
                    },
                    TotalAmount: codataccounting.Float64(6378.56),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ce52b895-c537-4c64-94ef-b0b34896c3ca",
                                Name: codataccounting.String("Jodi Russel"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e2fd5707-5779-4291-b7de-ac646ecb5734",
                                Name: codataccounting.String("Bobbie Terry"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("cum"),
                            ID: "1e5a2b12-eb07-4f11-adb9-9545fc95fa88",
                        },
                        IsBilledTo: shared.BilledToTypeEnumCustomer,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0e189dbb-30fc-4b33-aa05-5b197cd44e2f",
                            Name: codataccounting.String("Louise Schulist"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3513bb6f-48b6-456b-8db3-5ff2e4b27537",
                            Name: codataccounting.String("Jordan Romaguera"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e7319c17-7d52-45f7-bb11-4eeb52ff785f",
                            Name: codataccounting.String("Jimmy Konopelski II"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d4c98e0c-2bb8-49eb-b5da-d636c600503d",
                            Name: codataccounting.String("Willis Rippin Sr."),
                        },
                        shared.TrackingCategoryRef{
                            ID: "80f739ae-9e05-47eb-809e-2810331f3981",
                            Name: codataccounting.String("Leroy Schinner Jr."),
                        },
                    },
                    UnitAmount: 6985.58,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("607f3c93-c73b-49da-bf2c-eda7e23f2257"),
                        Name: codataccounting.String("Virginia Bins"),
                    },
                    Description: codataccounting.String("delectus"),
                    DiscountAmount: codataccounting.Float64(2512.12),
                    DiscountPercentage: codataccounting.Float64(7193.89),
                    ItemRef: &shared.ItemRef{
                        ID: "7544e472-e802-4857-a5b4-0463a7d575f1",
                        Name: codataccounting.String("Maria Ankunding"),
                    },
                    Quantity: 4137.58,
                    SubTotal: codataccounting.Float64(2561.14),
                    TaxAmount: codataccounting.Float64(6770.45),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8237.18),
                        ID: codataccounting.String("7334ec1b-781b-436a-8808-8d100efada20"),
                        Name: codataccounting.String("Mrs. Olive Weissnat"),
                    },
                    TotalAmount: codataccounting.Float64(1858.97),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b2164cf9-ab83-466c-b23f-fda9e06bee48",
                                Name: codataccounting.String("Dr. Marion Schaefer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0e115c80-bff9-4185-84ec-42defcce8f19",
                                Name: codataccounting.String("Joy Kessler"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e63562a7-b408-4f05-a3d4-8fdaf313a1f5",
                                Name: codataccounting.String("Woodrow Mitchell III"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9c0b36f2-5ea9-444f-bb75-6c11f6c37a51",
                                Name: codataccounting.String("Beth Cummerata"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("praesentium"),
                            ID: "35bbc05a-23a4-45ce-bc5f-de10a0ce2169",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "10019c6d-c5e3-4476-a799-bfbbe6949fb2",
                            Name: codataccounting.String("Alton Goyette"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e6c3d5db-3ade-4bd5-9aea-4c506a8aa94c",
                            Name: codataccounting.String("Theresa Jacobi"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cf5e9d9a-4578-4adc-9ac6-00dec001ac80",
                            Name: codataccounting.String("Rochelle Cormier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "09ff8f0f-816f-4f34-b7c1-3e902c14125b",
                            Name: codataccounting.String("Miss Isabel Kassulke"),
                        },
                    },
                    UnitAmount: 4016.88,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("8151a472-af92-43c5-949f-83f350cf876f"),
                        Name: codataccounting.String("Mr. Robin Miller"),
                    },
                    Description: codataccounting.String("commodi"),
                    DiscountAmount: codataccounting.Float64(9087.34),
                    DiscountPercentage: codataccounting.Float64(7815.82),
                    ItemRef: &shared.ItemRef{
                        ID: "bb4e243c-f789-4ffa-beda-53e5ae6e0ac1",
                        Name: codataccounting.String("Frederick Schaden"),
                    },
                    Quantity: 6200.54,
                    SubTotal: codataccounting.Float64(7935.68),
                    TaxAmount: codataccounting.Float64(1543.89),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3006.51),
                        ID: codataccounting.String("7c88373a-40e1-4942-b32e-55055756f5d5"),
                        Name: codataccounting.String("Meredith Bailey"),
                    },
                    TotalAmount: codataccounting.Float64(177.68),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f2dfe13d-b4f6-42cb-a3f8-941aebc0b80a",
                                Name: codataccounting.String("Velma Cummings"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3b2ecfcc-8f89-4501-8f5d-d3d6fa1804e5",
                                Name: codataccounting.String("Rosalie Lynch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "168a363c-8873-4e48-8380-b1f6b8ca275a",
                                Name: codataccounting.String("Mrs. Mary Pfannerstill"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("aliquam"),
                            ID: "95cc6991-71b5-41c1-bdb1-cf4b888ebdfc",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "cca99bc7-fc0b-42dc-a108-73e42b006d67",
                            Name: codataccounting.String("Guy Kovacek"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8581a582-08c5-44fe-ba9c-95f2eac5565d",
                            Name: codataccounting.String("Nancy Kuhlman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ee81206e-2813-4fa4-a41c-480d3f2132af",
                            Name: codataccounting.String("Mr. Connie Brakus"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "514f4cc6-f18b-4f96-a1a6-a4f77a87ee3e",
                            Name: codataccounting.String("Susie Ward"),
                        },
                    },
                    UnitAmount: 1316.87,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("impedit"),
            Note: codataccounting.String("aliquid"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("facilis"),
                        Currency: codataccounting.String("ipsum"),
                        CurrencyRate: codataccounting.Float64(2848.85),
                        TotalAmount: codataccounting.Float64(3088.19),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("18e3bb91-c8d9-475e-8e84-19d8f84f144f"),
                            Name: codataccounting.String("Kerry Altenwerth"),
                        },
                        Currency: codataccounting.String("facere"),
                        CurrencyRate: codataccounting.Float64(7890.16),
                        ID: codataccounting.String("c4aa5f3c-abd9-405a-972e-056728227b2d"),
                        Note: codataccounting.String("nesciunt"),
                        PaidOnDate: codataccounting.String("ipsa"),
                        Reference: codataccounting.String("sint"),
                        TotalAmount: codataccounting.Float64(2913.89),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("esse"),
                        Currency: codataccounting.String("accusantium"),
                        CurrencyRate: codataccounting.Float64(7181.19),
                        TotalAmount: codataccounting.Float64(9565.45),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7a4fa87c-f535-4a6f-ae54-ebf60c321f02"),
                            Name: codataccounting.String("Luz King"),
                        },
                        Currency: codataccounting.String("ratione"),
                        CurrencyRate: codataccounting.Float64(2218.24),
                        ID: codataccounting.String("67fe1a0c-c8df-479f-8a39-6d90c364b7c1"),
                        Note: codataccounting.String("enim"),
                        PaidOnDate: codataccounting.String("nulla"),
                        Reference: codataccounting.String("maiores"),
                        TotalAmount: codataccounting.Float64(7156.22),
                    },
                },
            },
            RemainingCredit: 6496.57,
            SourceModifiedDate: codataccounting.String("impedit"),
            Status: shared.BillCreditNoteStatusEnumPartiallyPaid,
            SubTotal: 911.2,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "blanditiis": map[string]interface{}{
                        "dicta": "impedit",
                        "tempora": "eveniet",
                        "repudiandae": "sed",
                    },
                    "impedit": map[string]interface{}{
                        "impedit": "vel",
                        "eligendi": "recusandae",
                        "ex": "beatae",
                    },
                    "veritatis": map[string]interface{}{
                        "itaque": "vero",
                        "quidem": "illo",
                        "quo": "dignissimos",
                        "minus": "distinctio",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "db6eec74-378b-4a25-b177-47dc915ad2ca",
                SupplierName: codataccounting.String("repellat"),
            },
            TotalAmount: 3419.83,
            TotalDiscount: 8473.08,
            TotalTaxAmount: 8450.86,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4563.71,
                    Name: "Edith Spencer DVM",
                },
                shared.WithholdingTaxitems{
                    Amount: 3502.71,
                    Name: "Phil Collier",
                },
            },
        },
        BillCreditNoteID: "officia",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(379661),
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
