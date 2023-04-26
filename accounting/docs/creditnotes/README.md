# CreditNotes

## Overview

Credit notes

### Available Operations

* [CreateCreditNote](#createcreditnote) - Create credit note
* [GetCreateUpdateCreditNotesModel](#getcreateupdatecreditnotesmodel) - Get create/update credit note model
* [GetCreditNote](#getcreditnote) - Get credit note
* [ListCreditNotes](#listcreditnotes) - List credit notes
* [UpdateCreditNote](#updatecreditnote) - Update creditNote

## CreateCreditNote

Push credit note


Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating a credit note.

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
    req := operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(8282.67),
            AdditionalTaxPercentage: codataccounting.Float64(4716.93),
            AllocatedOnDate: codataccounting.String("sed"),
            CreditNoteNumber: codataccounting.String("optio"),
            Currency: codataccounting.String("nulla"),
            CurrencyRate: codataccounting.Float64(1664.81),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("modi"),
                ID: "84da2172-9f2a-4c41-af57-25f1169ac1e4",
            },
            DiscountPercentage: 1100.31,
            ID: codataccounting.String("d8a23c23-e34f-42df-a4a1-97f6de922151"),
            IssueDate: codataccounting.String("delectus"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("17120998-53e9-4f54-bd85-4439ee224460"),
                        Name: codataccounting.String("Alicia Ebert"),
                    },
                    Description: codataccounting.String("architecto"),
                    DiscountAmount: codataccounting.Float64(3290.16),
                    DiscountPercentage: codataccounting.Float64(2991.66),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "188c2f56-e85d-4a78-b2ea-bd617c3b0d51",
                        Name: codataccounting.String("Jim Grady"),
                    },
                    Quantity: 220.18,
                    SubTotal: codataccounting.Float64(1013.18),
                    TaxAmount: codataccounting.Float64(7340.76),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6450.47),
                        ID: codataccounting.String("d8706d46-082b-4fbd-841f-f5d4e2ae4fb5"),
                        Name: codataccounting.String("Andres Fay"),
                    },
                    TotalAmount: codataccounting.Float64(1049.9),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "638f1edb-7835-49ec-85cb-860f8cd580ba",
                                Name: codataccounting.String("Mr. Emily Macejkovic"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4fe44472-97cd-43b1-9d3b-bce247b7684e",
                                Name: codataccounting.String("Mr. Emmett Heidenreich"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("laboriosam"),
                            ID: "d71cffbd-0eb7-44b8-8219-53b44bd3c431",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d33e5953-c001-4139-863a-a41e6c31cc2f",
                            Name: codataccounting.String("May Sauer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c9a41ffb-e9cb-4d79-9ee6-5e076cc7abf6",
                            Name: codataccounting.String("Jeanette Veum"),
                        },
                    },
                    UnitAmount: 7817.77,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("71641934-b90f-42e0-9d19-d2fc2f9e2e10"),
                        Name: codataccounting.String("Daisy Fritsch"),
                    },
                    Description: codataccounting.String("molestias"),
                    DiscountAmount: codataccounting.Float64(2261.98),
                    DiscountPercentage: codataccounting.Float64(3180.3),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "d237a72f-9084-49d6-aed4-aecb7537cd92",
                        Name: codataccounting.String("Lori Schneider"),
                    },
                    Quantity: 9816.77,
                    SubTotal: codataccounting.Float64(3461.64),
                    TaxAmount: codataccounting.Float64(4998.74),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2935.12),
                        ID: codataccounting.String("91aabfa2-e761-4f0c-a4d4-56ef1031e689"),
                        Name: codataccounting.String("Terrell Bashirian Jr."),
                    },
                    TotalAmount: codataccounting.Float64(557.9),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e22cd55c-c058-44a1-84d7-6d971fc820c6",
                                Name: codataccounting.String("Maryann Ankunding"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("rerum"),
                            ID: "b8e0cc88-5187-4e4d-a04a-f28c5dddb46a",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "cfd6d828-da01-4319-9129-646645c1d81f",
                            Name: codataccounting.String("Margarita Auer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "569b7aff-0ea2-4216-8be0-71bc163e279a",
                            Name: codataccounting.String("Kelli Beier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "da99257d-04f4-4084-ba74-2d84496cbdee",
                            Name: codataccounting.String("Dominick Jakubowski"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9bc63562-ebfd-4f55-8294-c060b06a1287",
                            Name: codataccounting.String("Stacey Fritsch"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f6d0c6d6-ed9c-473d-9634-571509a8e870",
                            Name: codataccounting.String("Jeff Schimmel"),
                        },
                    },
                    UnitAmount: 1021.82,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f9c242c7-b66a-41f3-8c73-df5b6719890f"),
                        Name: codataccounting.String("Andrea Orn"),
                    },
                    Description: codataccounting.String("expedita"),
                    DiscountAmount: codataccounting.Float64(2678.39),
                    DiscountPercentage: codataccounting.Float64(2398.58),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "8d85b260-591d-4745-a3c2-059c9c3f567e",
                        Name: codataccounting.String("Kellie Cormier"),
                    },
                    Quantity: 4884.47,
                    SubTotal: codataccounting.Float64(3792.36),
                    TaxAmount: codataccounting.Float64(3395.03),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7473.58),
                        ID: codataccounting.String("1d62fcda-ce1f-4012-96ce-2239e8f25cd0"),
                        Name: codataccounting.String("Dennis Moen"),
                    },
                    TotalAmount: codataccounting.Float64(3460.81),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f439e392-66cb-4d95-b7aa-2b24113695d1",
                                Name: codataccounting.String("Charlie Jacobs"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "fcc45962-17c2-4977-a763-34254038bfb5",
                                Name: codataccounting.String("Clinton Bosco"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "81905573-89ce-4dba-87fd-a39594d66bc2",
                                Name: codataccounting.String("Erick Haag III"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolorem"),
                            ID: "2b9954b6-fa22-4063-a982-8553cb10006b",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "4921ec20-53b7-4493-a6ac-8ee0f2bf1958",
                            Name: codataccounting.String("Dr. Irving Gislason I"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3deba297-be3e-490b-840d-f868fd52405c",
                            Name: codataccounting.String("Dr. Rodney Dooley"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92f4f127-fb0e-40bf-9f82-17978d0acca7",
                            Name: codataccounting.String("Iris VonRueden"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b7021a52-046b-464e-99fb-0e67e094fdfe",
                            Name: codataccounting.String("Dustin Heidenreich DVM"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f53a34a1-b8fe-4997-b1ad-c05d85ae2dfb",
                            Name: codataccounting.String("Lisa Wuckert"),
                        },
                    },
                    UnitAmount: 5618.25,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("74290d33-6561-4eca-96ef-89451bd76eee"),
                        Name: codataccounting.String("Maurice Boehm"),
                    },
                    Description: codataccounting.String("modi"),
                    DiscountAmount: codataccounting.Float64(8226.31),
                    DiscountPercentage: codataccounting.Float64(6266.37),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1fad3551-2f06-4d4e-9b72-f0f548568a04",
                        Name: codataccounting.String("Mr. Lucille Weber"),
                    },
                    Quantity: 1009.76,
                    SubTotal: codataccounting.Float64(8459.84),
                    TaxAmount: codataccounting.Float64(4206.47),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9162.43),
                        ID: codataccounting.String("b9434645-d030-484f-bba5-cceff5cb01fe"),
                        Name: codataccounting.String("Joyce Torp"),
                    },
                    TotalAmount: codataccounting.Float64(5425.06),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "45ac82b8-5f8b-4c2c-aba8-da4127dd597f",
                                Name: codataccounting.String("Mr. Ray Koch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a1bc74b8-6cec-4c74-b77b-4848bd6a6f04",
                                Name: codataccounting.String("Diane Stokes"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3b808094-373e-4060-859b-ebbad02f2586",
                                Name: codataccounting.String("Mrs. Forrest Wilkinson"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("exercitationem"),
                            ID: "58daa95b-e6cd-4027-96c3-54aa432b47e1",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3c5208c2-3e98-402d-82f0-d45eb4a8b674",
                            Name: codataccounting.String("Clay Hintz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "18edc7f7-87e3-42e0-8b3d-3ed0c5670ef4",
                            Name: codataccounting.String("Maryann Stark"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9f1cc503-f6c3-49bc-90a6-290f957f3851",
                            Name: codataccounting.String("Luther Nader"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ef807aae-03f3-43ca-b9fb-9de4032ba26f",
                            Name: codataccounting.String("Jimmy Keebler"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a9216bcb-4158-435c-b364-1723133edc04",
                            Name: codataccounting.String("Luz Schmidt III"),
                        },
                    },
                    UnitAmount: 2277.13,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("facilis"),
            Note: codataccounting.String("libero"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("deserunt"),
                        Currency: codataccounting.String("eius"),
                        CurrencyRate: codataccounting.Float64(5630.24),
                        TotalAmount: codataccounting.Float64(1737.52),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("27c42c22-c553-4504-95c5-dbb3c57c1e49"),
                            Name: codataccounting.String("Jack Ward"),
                        },
                        Currency: codataccounting.String("similique"),
                        CurrencyRate: codataccounting.Float64(1847.97),
                        ID: codataccounting.String("57ddc191-2ebd-4e64-bfcc-5469d4015dfa"),
                        Note: codataccounting.String("esse"),
                        PaidOnDate: codataccounting.String("iste"),
                        Reference: codataccounting.String("ex"),
                        TotalAmount: codataccounting.Float64(1421.73),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("voluptatem"),
                        Currency: codataccounting.String("voluptas"),
                        CurrencyRate: codataccounting.Float64(7288.49),
                        TotalAmount: codataccounting.Float64(8766.36),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f2b0a3e4-2c1a-4a01-8e9a-ac2e9135586d"),
                            Name: codataccounting.String("Billie Windler"),
                        },
                        Currency: codataccounting.String("sint"),
                        CurrencyRate: codataccounting.Float64(4450.02),
                        ID: codataccounting.String("a4bfad2b-f7d6-47ca-84ad-99b41d612435"),
                        Note: codataccounting.String("nesciunt"),
                        PaidOnDate: codataccounting.String("sunt"),
                        Reference: codataccounting.String("blanditiis"),
                        TotalAmount: codataccounting.Float64(4643.01),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("perferendis"),
                        Currency: codataccounting.String("cumque"),
                        CurrencyRate: codataccounting.Float64(9639.68),
                        TotalAmount: codataccounting.Float64(4071.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8b03ad42-1bd4-43d1-b0cb-0a0003eb22d9"),
                            Name: codataccounting.String("Leonard Padberg PhD"),
                        },
                        Currency: codataccounting.String("excepturi"),
                        CurrencyRate: codataccounting.Float64(3025.96),
                        ID: codataccounting.String("faa741c5-7d1f-4edc-a050-d38dc3ce1854"),
                        Note: codataccounting.String("iusto"),
                        PaidOnDate: codataccounting.String("sunt"),
                        Reference: codataccounting.String("tenetur"),
                        TotalAmount: codataccounting.Float64(5799.52),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("necessitatibus"),
                        Currency: codataccounting.String("necessitatibus"),
                        CurrencyRate: codataccounting.Float64(4230.32),
                        TotalAmount: codataccounting.Float64(6203.63),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("166a8be3-444e-4ac8-b3a2-875c6c1fe606"),
                            Name: codataccounting.String("James Klocko"),
                        },
                        Currency: codataccounting.String("deserunt"),
                        CurrencyRate: codataccounting.Float64(6229.89),
                        ID: codataccounting.String("c87ae50c-1666-41a1-9913-6a7e8d53213f"),
                        Note: codataccounting.String("velit"),
                        PaidOnDate: codataccounting.String("asperiores"),
                        Reference: codataccounting.String("commodi"),
                        TotalAmount: codataccounting.Float64(3744.95),
                    },
                },
            },
            RemainingCredit: 5538.05,
            SourceModifiedDate: codataccounting.String("esse"),
            Status: shared.CreditNoteStatusEnumDraft,
            SubTotal: 1453.33,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "expedita": map[string]interface{}{
                        "autem": "aliquam",
                        "maxime": "nostrum",
                    },
                    "occaecati": map[string]interface{}{
                        "doloremque": "id",
                        "veniam": "ea",
                        "placeat": "necessitatibus",
                        "harum": "cumque",
                    },
                    "culpa": map[string]interface{}{
                        "laborum": "consequuntur",
                        "omnis": "maxime",
                        "officia": "iusto",
                        "natus": "ab",
                    },
                    "deleniti": map[string]interface{}{
                        "eligendi": "sint",
                    },
                },
            },
            TotalAmount: 3694.92,
            TotalDiscount: 3887.15,
            TotalTaxAmount: 4752.38,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4014.49,
                    Name: "Dawn Schiller",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(55015),
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

## GetCreateUpdateCreditNotesModel

Get create/update credit note model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating and updating a credit note.

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
    req := operations.GetCreateUpdateCreditNotesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.CreditNotes.GetCreateUpdateCreditNotesModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetCreditNote

Gets a single creditNote corresponding to the given ID.

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
    req := operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "nam",
    }

    res, err := s.CreditNotes.GetCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## ListCreditNotes

Gets a list of all credit notes for a company, with pagination

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
    req := operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("minima"),
    }

    res, err := s.CreditNotes.ListCreditNotes(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNotes != nil {
        // handle response
    }
}
```

## UpdateCreditNote

Posts an updated credit note to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support updating a credit note.

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
    req := operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(4263.08),
            AdditionalTaxPercentage: codataccounting.Float64(3905.83),
            AllocatedOnDate: codataccounting.String("minima"),
            CreditNoteNumber: codataccounting.String("et"),
            Currency: codataccounting.String("autem"),
            CurrencyRate: codataccounting.Float64(2204.55),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("culpa"),
                ID: "3638512a-b252-41b9-b2e0-72467b8a40bc",
            },
            DiscountPercentage: 240.78,
            ID: codataccounting.String("5fab0d65-0edf-422a-94d2-0ec90ea41d1f"),
            IssueDate: codataccounting.String("labore"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5e85156f-ff73-4fdf-94fd-d5ea9543398d"),
                        Name: codataccounting.String("Timmy Robel"),
                    },
                    Description: codataccounting.String("dolorum"),
                    DiscountAmount: codataccounting.Float64(5358.88),
                    DiscountPercentage: codataccounting.Float64(8253.69),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "63388e4d-8039-4ea5-b9b1-8a244fd61903",
                        Name: codataccounting.String("Drew Padberg"),
                    },
                    Quantity: 2456.35,
                    SubTotal: codataccounting.Float64(5302.25),
                    TaxAmount: codataccounting.Float64(9110.49),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8423.7),
                        ID: codataccounting.String("0dc671dc-7f1e-43af-9592-0c90d1b4901f"),
                        Name: codataccounting.String("Opal Sporer"),
                    },
                    TotalAmount: codataccounting.Float64(8051.67),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "a32639da-5b7b-4690-ab88-1a94f643664a",
                                Name: codataccounting.String("Irvin Baumbach"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8c691d73-2d9f-4baf-9476-a2ae8dcc50c8",
                                Name: codataccounting.String("Mr. Jeffery Hegmann"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "73784893-0750-4a00-a966-ec736d431943",
                                Name: codataccounting.String("Sidney Ruecker"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolor"),
                            ID: "c92398ed-3d3a-4b7c-a3c5-ca8649a70cfd",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6989b720-6451-4077-919e-a83d492ed14b",
                            Name: codataccounting.String("Blake Connelly V"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4545e955-dcc1-485e-a490-1c7c43ad2daa",
                            Name: codataccounting.String("Christy Gorczany"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a3d230ed-f738-411a-9153-82bd7ed56507",
                            Name: codataccounting.String("Judy Bogan"),
                        },
                    },
                    UnitAmount: 5330.96,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f4d73965-64c2-40a0-b11a-961d24a7dbb8"),
                        Name: codataccounting.String("Bill Frami"),
                    },
                    Description: codataccounting.String("atque"),
                    DiscountAmount: codataccounting.Float64(5968.02),
                    DiscountPercentage: codataccounting.Float64(1421.56),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "cf7812cb-512c-4878-a40b-f548f88f8f1b",
                        Name: codataccounting.String("James Reynolds"),
                    },
                    Quantity: 8881.27,
                    SubTotal: codataccounting.Float64(1202.57),
                    TaxAmount: codataccounting.Float64(9822.77),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1756.76),
                        ID: codataccounting.String("06d5d831-d008-4109-8f67-06673f3a681c"),
                        Name: codataccounting.String("Nellie Kerluke"),
                    },
                    TotalAmount: codataccounting.Float64(7742.94),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "742409a2-15e0-4860-9489-a5f63e3af3dd",
                                Name: codataccounting.String("Marty Spencer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3dcd6348-3e4a-47a9-8e4d-f37e45b8955d",
                                Name: codataccounting.String("Gloria Emard I"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a4823109-07bd-4354-8092-bd734f02449d",
                                Name: codataccounting.String("Cecil Wintheiser"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b20fe5d9-11cb-4fe7-89ca-f45a27f69e2c",
                                Name: codataccounting.String("Merle Keebler Jr."),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("debitis"),
                            ID: "9db3ad4c-6b03-4108-99c3-37473082b94f",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b1fd5671-e9c3-4263-90a4-67143789ce0e",
                            Name: codataccounting.String("Evan Buckridge"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d93a74c0-252f-4e3b-8b4d-b8b778ebb6e1",
                            Name: codataccounting.String("Brandon Rutherford"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "02bafb2c-bc46-435d-9e65-da028c3e951a",
                            Name: codataccounting.String("Dr. Leigh Dickens"),
                        },
                    },
                    UnitAmount: 6458.29,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("sint"),
            Note: codataccounting.String("eum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("magnam"),
                        Currency: codataccounting.String("rem"),
                        CurrencyRate: codataccounting.Float64(5808.35),
                        TotalAmount: codataccounting.Float64(8240.62),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7b78673e-13a1-42a6-b992-494594487f5c"),
                            Name: codataccounting.String("Mario Fay"),
                        },
                        Currency: codataccounting.String("ex"),
                        CurrencyRate: codataccounting.Float64(7339.99),
                        ID: codataccounting.String("86b3cdf6-415b-4044-9f9d-f13f4eedbe78"),
                        Note: codataccounting.String("tempore"),
                        PaidOnDate: codataccounting.String("reiciendis"),
                        Reference: codataccounting.String("commodi"),
                        TotalAmount: codataccounting.Float64(261.97),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ea"),
                        Currency: codataccounting.String("molestias"),
                        CurrencyRate: codataccounting.Float64(1564.16),
                        TotalAmount: codataccounting.Float64(3730.95),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("894ea763-d5c7-4279-9b78-5148d6d549e5"),
                            Name: codataccounting.String("Kim Hegmann"),
                        },
                        Currency: codataccounting.String("dolorem"),
                        CurrencyRate: codataccounting.Float64(7110.04),
                        ID: codataccounting.String("c0f970c4-2fc9-4f48-8422-5e75b796065c"),
                        Note: codataccounting.String("eaque"),
                        PaidOnDate: codataccounting.String("earum"),
                        Reference: codataccounting.String("earum"),
                        TotalAmount: codataccounting.Float64(6367.75),
                    },
                },
            },
            RemainingCredit: 4254.84,
            SourceModifiedDate: codataccounting.String("sapiente"),
            Status: shared.CreditNoteStatusEnumPaid,
            SubTotal: 1904.44,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "molestias": map[string]interface{}{
                        "fuga": "beatae",
                    },
                    "distinctio": map[string]interface{}{
                        "eligendi": "unde",
                        "veniam": "nam",
                        "accusamus": "vitae",
                    },
                    "explicabo": map[string]interface{}{
                        "incidunt": "soluta",
                        "nihil": "adipisci",
                    },
                },
            },
            TotalAmount: 5919.98,
            TotalDiscount: 9674.76,
            TotalTaxAmount: 2952.84,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9078.26,
                    Name: "Mr. Maureen Christiansen",
                },
                shared.WithholdingTaxitems{
                    Amount: 1182.21,
                    Name: "Hector Hegmann",
                },
                shared.WithholdingTaxitems{
                    Amount: 7728.04,
                    Name: "Marshall Schmitt",
                },
                shared.WithholdingTaxitems{
                    Amount: 1411.42,
                    Name: "Billy Reichert DVM",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "excepturi",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(256890),
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
