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

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) to see which integrations support this endpoint.


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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Create(ctx, operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("quae"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("ipsum"),
            CurrencyRate: codataccounting.Float64(6924.72),
            DiscountPercentage: 5651.89,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("excepturi"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("488e1e91-e450-4ad2-abd4-4269802d502a"),
                        Name: codataccounting.String("Eddie Prosacco"),
                    },
                    Description: codataccounting.String("delectus"),
                    DiscountAmount: codataccounting.Float64(4332.88),
                    DiscountPercentage: codataccounting.Float64(2487.53),
                    ItemRef: &shared.ItemRef{
                        ID: "c969e9a3-efa7-47df-b14c-d66ae395efb9",
                        Name: codataccounting.String("Lynn Kuvalis"),
                    },
                    Quantity: 2305.33,
                    SubTotal: codataccounting.Float64(6439.9),
                    TaxAmount: codataccounting.Float64(3948.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4238.55),
                        ID: codataccounting.String("997074ba-4469-4b6e-a141-959890afa563"),
                        Name: codataccounting.String("Ms. Fred Hilll"),
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
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
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
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
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
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("698f447f-603e-48b4-85e8-0ca55efd20e4"),
                        Name: codataccounting.String("Ms. Pearl Towne"),
                    },
                    Description: codataccounting.String("praesentium"),
                    DiscountAmount: codataccounting.Float64(7400.98),
                    DiscountPercentage: codataccounting.Float64(3868.27),
                    ItemRef: &shared.ItemRef{
                        ID: "a89fbe3a-5aa8-4e48-a4d0-ab4075088e51",
                        Name: codataccounting.String("Ms. Ruben Cremin"),
                    },
                    Quantity: 9061.72,
                    SubTotal: codataccounting.Float64(6222.31),
                    TaxAmount: codataccounting.Float64(85.11),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2790.68),
                        ID: codataccounting.String("f3b1194b-8abf-4603-a79f-9dfe0ab7da8a"),
                        Name: codataccounting.String("Helen Schiller IV"),
                    },
                    TotalAmount: codataccounting.Float64(4420.36),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "86bc173d-689e-4ee9-926f-8d986e881ead",
                                Name: codataccounting.String("Lela Baumbach Jr."),
                            },
                            shared.TrackingCategoryRef{
                                ID: "12563f94-e29e-4973-a922-a57a15be3e06",
                                Name: codataccounting.String("Lena Beier"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2b6e3ab8-845f-4059-ba60-ff2a54a31e94",
                                Name: codataccounting.String("Carla Graham"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e865e795-6f92-451a-9a9d-a660ff57bfaa",
                                Name: codataccounting.String("Edwin Wolf"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("sapiente"),
                            ID: "c1b4512c-1032-4648-9c2f-615199ebfd0e",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e6c632ca-3aed-4011-b996-312fde047717",
                            Name: codataccounting.String("Irma Wuckert"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d0174763-60a1-45db-aa66-0659a1adeaab",
                            Name: codataccounting.String("Dr. Cassandra Halvorson"),
                        },
                    },
                    UnitAmount: 7758.03,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("645b08b6-1891-4baa-8fe1-ade008e6f8c5"),
                        Name: codataccounting.String("Dr. Chris Hermiston"),
                    },
                    Description: codataccounting.String("impedit"),
                    DiscountAmount: codataccounting.Float64(8427.77),
                    DiscountPercentage: codataccounting.Float64(7205.28),
                    ItemRef: &shared.ItemRef{
                        ID: "5a341814-3010-4421-813d-5208ece7e253",
                        Name: codataccounting.String("Andre Kautzer"),
                    },
                    Quantity: 3494.4,
                    SubTotal: codataccounting.Float64(704.1),
                    TaxAmount: codataccounting.Float64(7814.8),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4218.44),
                        ID: codataccounting.String("c6e205e1-6dea-4b3f-ac95-78a64584273a"),
                        Name: codataccounting.String("Randall Boyle"),
                    },
                    TotalAmount: codataccounting.Float64(1173.8),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2309fb09-2992-41ae-bb9f-58c4d86e68e4",
                                Name: codataccounting.String("Ignacio Bartoletti"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "013f59da-757a-459e-8fef-66ef1caa3383",
                                Name: codataccounting.String("Victor Rogahn"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolore"),
                            ID: "77373c8d-72f6-44d1-9b1f-2c4310661e96",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "9e1cf9e0-6e3a-4437-800a-e6b6bc9b8f75",
                            Name: codataccounting.String("Terence O'Keefe"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a9741d31-1352-4965-bb8a-7202611435e1",
                            Name: codataccounting.String("Lindsay Stiedemann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2259b1ab-da8c-4070-a108-4cb0672d1ad8",
                            Name: codataccounting.String("Daisy Tillman"),
                        },
                    },
                    UnitAmount: 5750.78,
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
                        AllocatedOnDate: codataccounting.String("ipsam"),
                        Currency: codataccounting.String("rerum"),
                        CurrencyRate: codataccounting.Float64(5156.38),
                        TotalAmount: codataccounting.Float64(3572.07),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("efbd02ba-e0be-42d7-8225-9e3ea4b5197f"),
                            Name: codataccounting.String("Steve Fritsch"),
                        },
                        Currency: codataccounting.String("at"),
                        CurrencyRate: codataccounting.Float64(6378.56),
                        ID: codataccounting.String("7ce52b89-5c53-47c6-854e-fb0b34896c3c"),
                        Note: codataccounting.String("fuga"),
                        PaidOnDate: codataccounting.String("nostrum"),
                        Reference: codataccounting.String("est"),
                        TotalAmount: codataccounting.Float64(7708.73),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("delectus"),
                        Currency: codataccounting.String("tempore"),
                        CurrencyRate: codataccounting.Float64(8786.01),
                        TotalAmount: codataccounting.Float64(1415.06),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fd570757-7929-4177-9eac-646ecb573409"),
                            Name: codataccounting.String("Earl VonRueden DVM"),
                        },
                        Currency: codataccounting.String("veniam"),
                        CurrencyRate: codataccounting.Float64(6592.68),
                        ID: codataccounting.String("2b12eb07-f116-4db9-9545-fc95fa88970e"),
                        Note: codataccounting.String("architecto"),
                        PaidOnDate: codataccounting.String("quos"),
                        Reference: codataccounting.String("iste"),
                        TotalAmount: codataccounting.Float64(8268.62),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(7316.34),
            SourceModifiedDate: codataccounting.String("libero"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "doloremque": map[string]interface{}{
                        "impedit": "cum",
                        "ipsum": "adipisci",
                        "saepe": "deserunt",
                        "doloremque": "quis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "5b197cd4-4e2f-452d-82d3-513bb6f48b65",
                SupplierName: codataccounting.String("nisi"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 7277.71,
            TotalTaxAmount: 7945.07,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 7060.61,
                    Name: "Erin Wiza",
                },
                shared.WithholdingTaxitems{
                    Amount: 8915.81,
                    Name: "Susie Davis",
                },
                shared.WithholdingTaxitems{
                    Amount: 2072.96,
                    Name: "Genevieve Lebsack",
                },
                shared.WithholdingTaxitems{
                    Amount: 6040.78,
                    Name: "Miss Kelly Ernser",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(111496),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

## Get

﻿Gets a single billCreditNote corresponding to the given ID.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Get(ctx, operations.GetBillCreditNoteRequest{
        BillCreditNoteID: "dignissimos",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNote != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

﻿Get create/update bill credit note model.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating and updating bill credit notes.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.GetCreateUpdateModel(ctx, operations.GetCreateUpdateBillCreditNotesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

﻿Gets a list of all bill credit notes for a company, with pagination.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.List(ctx, operations.ListBillCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNotes != nil {
        // handle response
    }
}
```

## Update

﻿Posts an updated billCreditNote to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support updating bill credit notes.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillCreditNotes.Update(ctx, operations.UpdateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("fugiat"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("ad"),
            CurrencyRate: codataccounting.Float64(1348.18),
            DiscountPercentage: 3165.01,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("delectus"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7b114eeb-52ff-4785-bc37-814d4c98e0c2"),
                        Name: codataccounting.String("Rudolph Macejkovic"),
                    },
                    Description: codataccounting.String("rerum"),
                    DiscountAmount: codataccounting.Float64(4923.61),
                    DiscountPercentage: codataccounting.Float64(3609.34),
                    ItemRef: &shared.ItemRef{
                        ID: "dad636c6-0050-43d8-bb31-180f739ae9e0",
                        Name: codataccounting.String("Nellie Waters"),
                    },
                    Quantity: 439.75,
                    SubTotal: codataccounting.Float64(5740.92),
                    TaxAmount: codataccounting.Float64(8795.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1786.35),
                        ID: codataccounting.String("810331f3-981d-44c7-80b6-07f3c93c73b9"),
                        Name: codataccounting.String("Luke Fay"),
                    },
                    TotalAmount: codataccounting.Float64(7782.76),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "da7e23f2-2574-411f-af4b-7544e472e802",
                                Name: codataccounting.String("Bill Kling"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b40463a7-d575-4f14-80e7-64ad7334ec1b",
                                Name: codataccounting.String("Tracey Bosco"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6a08088d-100e-4fad-a200-ef0422eb2164",
                                Name: codataccounting.String("Courtney Maggio"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8366c723-ffda-49e0-abee-4825c1fc0e11",
                                Name: codataccounting.String("Miss Marianne Leffler"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("a"),
                            ID: "918544ec-42de-4fcc-a8f1-977773e63562",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b408f05e-3d48-4fda-b313-a1f5fd94259c",
                            Name: codataccounting.String("Yvette Dooley"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5ea944f3-b756-4c11-b6c3-7a5126243835",
                            Name: codataccounting.String("Mrs. Johnathan Russel"),
                        },
                    },
                    UnitAmount: 1593.93,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3a45cefc-5fde-410a-8ce2-169e510019c6"),
                        Name: codataccounting.String("Jermaine Hettinger"),
                    },
                    Description: codataccounting.String("magnam"),
                    DiscountAmount: codataccounting.Float64(4935.91),
                    DiscountPercentage: codataccounting.Float64(3884.04),
                    ItemRef: &shared.ItemRef{
                        ID: "2799bfbb-e694-49fb-abb4-ecae6c3d5db3",
                        Name: codataccounting.String("Kristopher Walter"),
                    },
                    Quantity: 3233.65,
                    SubTotal: codataccounting.Float64(8161.51),
                    TaxAmount: codataccounting.Float64(6746.83),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9114.51),
                        ID: codataccounting.String("a4c506a8-aa94-4c02-a44c-f5e9d9a4578a"),
                        Name: codataccounting.String("Edmund Boyle"),
                    },
                    TotalAmount: codataccounting.Float64(3857.39),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0dec001a-c802-4e2e-809f-f8f0f816ff34",
                                Name: codataccounting.String("Mrs. Pearl Rosenbaum"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("excepturi"),
                            ID: "02c14125-b096-40a6-a815-1a472af923c5",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "9f83f350-cf87-46ff-b901-c6ecbb4e243c",
                            Name: codataccounting.String("Claude Kutch"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "afeda53e-5ae6-4e0a-8184-c2b9c247c883",
                            Name: codataccounting.String("Grace Padberg PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1942f32e-5505-4575-af5d-56d0bd0af2df",
                            Name: codataccounting.String("Joe Fisher"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4f62cba3-f894-41ae-bc0b-80a6924d3b2e",
                            Name: codataccounting.String("Van Schiller"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f895010f-5dd3-4d6f-a180-4e54c82f168a",
                            Name: codataccounting.String("Vicki Feeney"),
                        },
                    },
                    UnitAmount: 5277.15,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ducimus"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("recusandae"),
                        Currency: codataccounting.String("tempora"),
                        CurrencyRate: codataccounting.Float64(5034.49),
                        TotalAmount: codataccounting.Float64(2580.59),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("380b1f6b-8ca2-475a-a0a0-4c495cc69917"),
                            Name: codataccounting.String("Miss Juana Hilpert MD"),
                        },
                        Currency: codataccounting.String("facere"),
                        CurrencyRate: codataccounting.Float64(7079.83),
                        ID: codataccounting.String("1cf4b888-ebdf-4c4c-8ca9-9bc7fc0b2dce"),
                        Note: codataccounting.String("veritatis"),
                        PaidOnDate: codataccounting.String("aut"),
                        Reference: codataccounting.String("laudantium"),
                        TotalAmount: codataccounting.Float64(4804.21),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(2198.6),
            SourceModifiedDate: codataccounting.String("voluptates"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "magni": map[string]interface{}{
                        "doloremque": "voluptatem",
                        "eum": "at",
                        "eum": "reprehenderit",
                    },
                    "voluptatum": map[string]interface{}{
                        "nihil": "atque",
                        "rerum": "deserunt",
                        "atque": "nostrum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "81a58208-c54f-4efa-9c95-f2eac5565d30",
                SupplierName: codataccounting.String("odio"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 7943.06,
            TotalTaxAmount: 9903.69,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9121.51,
                    Name: "Mrs. Samuel Considine",
                },
                shared.WithholdingTaxitems{
                    Amount: 1489.75,
                    Name: "Ralph Dooley",
                },
                shared.WithholdingTaxitems{
                    Amount: 3034.21,
                    Name: "Edwin Cartwright",
                },
                shared.WithholdingTaxitems{
                    Amount: 5283.2,
                    Name: "Cristina Ebert",
                },
            },
        },
        BillCreditNoteID: "inventore",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(193236),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```
