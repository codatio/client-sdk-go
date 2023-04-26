# BillCreditNotes

## Overview

Bill credit notes

### Available Operations

* [Create](#create) - Create bill credit note
* [Get](#get) - Get bill credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update bill credit note model
* [List](#list) - List bill credit notes
* [Update](#update) - Update bill credit note

## Create

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
            AllocatedOnDate: codataccounting.String("ipsum"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("quidem"),
            CurrencyRate: codataccounting.Float64(5651.89),
            DiscountPercentage: 5666.02,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("pariatur"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("88e1e91e-450a-4d2a-bd44-269802d502a9"),
                        Name: codataccounting.String("Olivia Rice"),
                    },
                    Description: codataccounting.String("eum"),
                    DiscountAmount: codataccounting.Float64(2487.53),
                    DiscountPercentage: codataccounting.Float64(7561.07),
                    ItemRef: &shared.ItemRef{
                        ID: "969e9a3e-fa77-4dfb-94cd-66ae395efb9b",
                        Name: codataccounting.String("Nelson Lesch"),
                    },
                    Quantity: 6439.9,
                    SubTotal: codataccounting.Float64(3948.69),
                    TaxAmount: codataccounting.Float64(4238.55),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6188.09),
                        ID: codataccounting.String("97074ba4-469b-46e2-9419-59890afa563e"),
                        Name: codataccounting.String("Vivian Boyle"),
                    },
                    TotalAmount: codataccounting.Float64(8919.24),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c8b711e5-b7fd-42ed-8289-21cddc692601",
                                Name: codataccounting.String("Rickey Hintz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b0d5f0d3-0c5f-4bb2-9870-53202c73d5fe",
                                Name: codataccounting.String("Miss Cesar Metz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("blanditiis"),
                            ID: "909b3fe4-9a8d-49cb-b486-33323f9b77f3",
                        },
                        IsBilledTo: shared.BilledToTypeEnumCustomer,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "100674eb-f692-480d-9ba7-7a89ebf737ae",
                            Name: codataccounting.String("Judy Aufderhar"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5e6a95d8-a0d4-446c-a2af-7a73cf3be453",
                            Name: codataccounting.String("Miss Jimmie Kozey"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "26b5a734-29cd-4b1a-8422-bb679d232271",
                            Name: codataccounting.String("Miss Candice Weimann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b1e31b8b-90f3-4443-a110-8e0adcf4b921",
                            Name: codataccounting.String("Darren McClure"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e953f73e-f7fb-4c7a-bd74-dd39c0f5d2cf",
                            Name: codataccounting.String("Ted Romaguera MD"),
                        },
                    },
                    UnitAmount: 2694.79,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5626d436-813f-416d-9f5f-ce6c556146c3"),
                        Name: codataccounting.String("Dr. Bruce Hane"),
                    },
                    Description: codataccounting.String("aut"),
                    DiscountAmount: codataccounting.Float64(114.27),
                    DiscountPercentage: codataccounting.Float64(5334.66),
                    ItemRef: &shared.ItemRef{
                        ID: "c42e141a-ac36-46c8-9d6b-144290747477",
                        Name: codataccounting.String("Blake Kihn"),
                    },
                    Quantity: 2835.19,
                    SubTotal: codataccounting.Float64(4334.39),
                    TaxAmount: codataccounting.Float64(3799.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8268.71),
                        ID: codataccounting.String("28c10ab3-cdca-4425-9904-e523c7e0bc71"),
                        Name: codataccounting.String("Christy Tillman"),
                    },
                    TotalAmount: codataccounting.Float64(5775.43),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f2a70c68-8282-4aa4-8256-2f222e9817ee",
                                Name: codataccounting.String("Joy Schmidt"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "61e6b7b9-5bc0-4ab3-820c-4f3789fd871f",
                                Name: codataccounting.String("Kirk Stracke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eveniet"),
                            ID: "fd121aa6-f1e6-474b-9b04-f15756082d68",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "19f1d170-5133-49d0-8086-a1840394c260",
                            Name: codataccounting.String("Jean Wunsch"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5f0642da-c7af-4515-8c41-3aa63aae8d67",
                            Name: codataccounting.String("Cecil Grant"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b675fd5e-60b3-475e-94f6-fbee41f33317",
                            Name: codataccounting.String("Doyle Feest"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "60eb1ea4-2655-45ba-bc28-744ed53b88f3",
                            Name: codataccounting.String("Byron Stroman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5c0b2f2f-b7b1-494a-a76b-26916fe1f08f",
                            Name: codataccounting.String("Tammy Medhurst"),
                        },
                    },
                    UnitAmount: 2155.29,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ea"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quos"),
                        Currency: codataccounting.String("voluptatibus"),
                        CurrencyRate: codataccounting.Float64(2716.53),
                        TotalAmount: codataccounting.Float64(2730.09),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7f603e8b-445e-480c-a55e-fd20e457e185"),
                            Name: codataccounting.String("Bennie Howe"),
                        },
                        Currency: codataccounting.String("error"),
                        CurrencyRate: codataccounting.Float64(9447.08),
                        ID: codataccounting.String("be3a5aa8-e482-44d0-ab40-75088e518620"),
                        Note: codataccounting.String("vel"),
                        PaidOnDate: codataccounting.String("nostrum"),
                        Reference: codataccounting.String("saepe"),
                        TotalAmount: codataccounting.Float64(6222.31),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("consequatur"),
                        Currency: codataccounting.String("incidunt"),
                        CurrencyRate: codataccounting.Float64(9688.65),
                        TotalAmount: codataccounting.Float64(2097.5),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b1194b8a-bf60-43a7-9f9d-fe0ab7da8a50"),
                            Name: codataccounting.String("Phil Boyer"),
                        },
                        Currency: codataccounting.String("asperiores"),
                        CurrencyRate: codataccounting.Float64(5199.52),
                        ID: codataccounting.String("6bc173d6-89ee-4e95-a6f8-d986e881ead4"),
                        Note: codataccounting.String("reiciendis"),
                        PaidOnDate: codataccounting.String("doloremque"),
                        Reference: codataccounting.String("repudiandae"),
                        TotalAmount: codataccounting.Float64(1160.98),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("accusantium"),
                        Currency: codataccounting.String("beatae"),
                        CurrencyRate: codataccounting.Float64(1747.72),
                        TotalAmount: codataccounting.Float64(3164.88),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("63f94e29-e973-4e92-aa57-a15be3e06080"),
                            Name: codataccounting.String("Tricia Denesik"),
                        },
                        Currency: codataccounting.String("necessitatibus"),
                        CurrencyRate: codataccounting.Float64(1875.52),
                        ID: codataccounting.String("ab8845f0-597a-460f-b2a5-4a31e94764a3"),
                        Note: codataccounting.String("debitis"),
                        PaidOnDate: codataccounting.String("laudantium"),
                        Reference: codataccounting.String("eum"),
                        TotalAmount: codataccounting.Float64(3679.27),
                    },
                },
            },
            RemainingCredit: 9282.19,
            SourceModifiedDate: codataccounting.String("esse"),
            Status: shared.BillCreditNoteStatusEnumPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quis": map[string]interface{}{
                        "reiciendis": "provident",
                        "aspernatur": "ullam",
                    },
                    "quasi": map[string]interface{}{
                        "nostrum": "mollitia",
                        "provident": "possimus",
                        "animi": "ex",
                    },
                    "aliquid": map[string]interface{}{
                        "repellat": "doloribus",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "57bfaad4-f9ef-4c1b-8512-c1032648dc2f",
                SupplierName: codataccounting.String("eum"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 1173.2,
            TotalTaxAmount: 3251.18,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5834.04,
                    Name: "Darin Rodriguez",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(52508),
    }

    res, err := s.BillCreditNotes.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

## Get

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
        BillCreditNoteID: "earum",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.BillCreditNotes.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNote != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

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

    res, err := s.BillCreditNotes.GetCreateUpdateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

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
        Query: codataccounting.String("perspiciatis"),
    }

    res, err := s.BillCreditNotes.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNotes != nil {
        // handle response
    }
}
```

## Update

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
            AllocatedOnDate: codataccounting.String("maiores"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("debitis"),
            CurrencyRate: codataccounting.Float64(3998.02),
            DiscountPercentage: 7809.31,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("suscipit"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("2ca3aed0-1179-4963-92fd-e04771778ff6"),
                        Name: codataccounting.String("Ms. Janis Batz"),
                    },
                    Description: codataccounting.String("esse"),
                    DiscountAmount: codataccounting.Float64(4037.93),
                    DiscountPercentage: codataccounting.Float64(2352.63),
                    ItemRef: &shared.ItemRef{
                        ID: "60a15db6-a660-4659-a1ad-eaab5851d6c6",
                        Name: codataccounting.String("Ms. Geraldine Ratke"),
                    },
                    Quantity: 3996.6,
                    SubTotal: codataccounting.Float64(1097.84),
                    TaxAmount: codataccounting.Float64(5308.6),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6063.08),
                        ID: codataccounting.String("1baa0fe1-ade0-408e-af8c-5f350d8cdb5a"),
                        Name: codataccounting.String("Michele Bode II"),
                    },
                    TotalAmount: codataccounting.Float64(2213.96),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "10421813-d520-48ec-a7e2-53b668451c6c",
                                Name: codataccounting.String("Mrs. Kate Cronin"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quasi"),
                            ID: "6deab3fe-c957-48a6-8584-273a8418d162",
                        },
                        IsBilledTo: shared.BilledToTypeEnumUnknown,
                        IsRebilledTo: shared.BilledToTypeEnumUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "9fb09299-21ae-4fb9-b58c-4d86e68e4be0",
                            Name: codataccounting.String("Mrs. Gina Abbott"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9da757a5-9ecf-4ef6-aef1-caa3383c2beb",
                            Name: codataccounting.String("Glenda Kling"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3c8d72f6-4d1d-4b1f-ac43-10661e96349e",
                            Name: codataccounting.String("Pat Wolf"),
                        },
                    },
                    UnitAmount: 26.77,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nisi"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("velit"),
                        Currency: codataccounting.String("laborum"),
                        CurrencyRate: codataccounting.Float64(2503.98),
                        TotalAmount: codataccounting.Float64(2244.67),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7000ae6b-6bc9-4b8f-b59e-ac55a9741d31"),
                            Name: codataccounting.String("Florence Hand"),
                        },
                        Currency: codataccounting.String("ex"),
                        CurrencyRate: codataccounting.Float64(3676.26),
                        ID: codataccounting.String("bb8a7202-6114-435e-939d-bc2259b1abda"),
                        Note: codataccounting.String("quos"),
                        PaidOnDate: codataccounting.String("placeat"),
                        Reference: codataccounting.String("sit"),
                        TotalAmount: codataccounting.Float64(4793.85),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ipsa"),
                        Currency: codataccounting.String("voluptates"),
                        CurrencyRate: codataccounting.Float64(800.61),
                        TotalAmount: codataccounting.Float64(493.48),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("84cb0672-d1ad-4879-aeb9-665b85efbd02"),
                            Name: codataccounting.String("Miss Grant VonRueden"),
                        },
                        Currency: codataccounting.String("eos"),
                        CurrencyRate: codataccounting.Float64(8448.54),
                        ID: codataccounting.String("782259e3-ea4b-4519-bf92-443da7ce52b8"),
                        Note: codataccounting.String("cupiditate"),
                        PaidOnDate: codataccounting.String("minima"),
                        Reference: codataccounting.String("placeat"),
                        TotalAmount: codataccounting.Float64(3165.42),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("neque"),
                        Currency: codataccounting.String("in"),
                        CurrencyRate: codataccounting.Float64(7963.97),
                        TotalAmount: codataccounting.Float64(4330.77),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("454efb0b-3489-46c3-8a5a-cfbe2fd57075"),
                            Name: codataccounting.String("Dora Mante"),
                        },
                        Currency: codataccounting.String("veritatis"),
                        CurrencyRate: codataccounting.Float64(4981.8),
                        ID: codataccounting.String("7deac646-ecb5-4734-89e3-eb1e5a2b12eb"),
                        Note: codataccounting.String("ipsa"),
                        PaidOnDate: codataccounting.String("ducimus"),
                        Reference: codataccounting.String("maiores"),
                        TotalAmount: codataccounting.Float64(873.82),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quasi"),
                        Currency: codataccounting.String("laboriosam"),
                        CurrencyRate: codataccounting.Float64(8634.71),
                        TotalAmount: codataccounting.Float64(7294.48),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("99545fc9-5fa8-4897-8e18-9dbb30fcb33e"),
                            Name: codataccounting.String("Anthony Hayes"),
                        },
                        Currency: codataccounting.String("architecto"),
                        CurrencyRate: codataccounting.Float64(5845.93),
                        ID: codataccounting.String("7cd44e2f-52d8-42d3-913b-b6f48b656bcd"),
                        Note: codataccounting.String("facilis"),
                        PaidOnDate: codataccounting.String("ipsum"),
                        Reference: codataccounting.String("ad"),
                        TotalAmount: codataccounting.Float64(9738.19),
                    },
                },
            },
            RemainingCredit: 9745.89,
            SourceModifiedDate: codataccounting.String("consequuntur"),
            Status: shared.BillCreditNoteStatusEnumPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "labore": map[string]interface{}{
                        "eos": "reprehenderit",
                        "nostrum": "neque",
                        "iusto": "est",
                    },
                    "rem": map[string]interface{}{
                        "fugiat": "unde",
                        "officiis": "ducimus",
                        "dolor": "dicta",
                        "error": "porro",
                    },
                    "vitae": map[string]interface{}{
                        "esse": "fugiat",
                        "ad": "aspernatur",
                    },
                    "enim": map[string]interface{}{
                        "iusto": "dignissimos",
                        "libero": "illo",
                        "ab": "incidunt",
                        "accusamus": "saepe",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "b52ff785-fc37-4814-94c9-8e0c2bb89eb7",
                SupplierName: codataccounting.String("corporis"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 8738.33,
            TotalTaxAmount: 6293.77,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 4348.27,
                    Name: "Gertrude Russel Jr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 3228.29,
                    Name: "Wendy Stanton",
                },
                shared.WithholdingTaxitems{
                    Amount: 7368.53,
                    Name: "Joyce Carroll DVM",
                },
                shared.WithholdingTaxitems{
                    Amount: 4797.07,
                    Name: "Shelly Pagac",
                },
            },
        },
        BillCreditNoteID: "repudiandae",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(10063),
    }

    res, err := s.BillCreditNotes.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```
