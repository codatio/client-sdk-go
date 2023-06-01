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
            AllocatedOnDate: codataccounting.String("molestias"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("excepturi"),
            CurrencyRate: codataccounting.Float64(8651.03),
            DiscountPercentage: 2653.89,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("praesentium"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e1e91e45-0ad2-4abd-8426-9802d502a94b"),
                        Name: codataccounting.String("Francisco Windler"),
                    },
                    Description: codataccounting.String("eligendi"),
                    DiscountAmount: codataccounting.Float64(5761.57),
                    DiscountPercentage: codataccounting.Float64(3960.98),
                    ItemRef: &shared.ItemRef{
                        ID: "9e9a3efa-77df-4b14-8d66-ae395efb9ba8",
                        Name: codataccounting.String("Timmy Feeney"),
                    },
                    Quantity: 4238.55,
                    SubTotal: codataccounting.Float64(6188.09),
                    TaxAmount: codataccounting.Float64(6063.93),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4748.67),
                        ID: codataccounting.String("074ba446-9b6e-4214-9959-890afa563e25"),
                        Name: codataccounting.String("Vera Wyman"),
                    },
                    TotalAmount: codataccounting.Float64(8061.94),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b711e5b7-fd2e-4d02-8921-cddc692601fb",
                                Name: codataccounting.String("Colleen Johnston PhD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5f0d30c5-fbb2-4587-8532-02c73d5fe9b9",
                                Name: codataccounting.String("Robyn Cruickshank"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "09b3fe49-a8d9-4cbf-8863-3323f9b77f3a",
                                Name: codataccounting.String("Ms. Christine Beer"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quaerat"),
                            ID: "ebf69280-d1ba-477a-89eb-f737ae4203ce",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6a95d8a0-d446-4ce2-af7a-73cf3be453f8",
                            Name: codataccounting.String("Karen Rath"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b5a73429-cdb1-4a84-a2bb-679d2322715b",
                            Name: codataccounting.String("George Runolfsdottir"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1e31b8b9-0f34-443a-9108-e0adcf4b9218",
                            Name: codataccounting.String("Toni Wolff"),
                        },
                    },
                    UnitAmount: 6064.76,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("53f73ef7-fbc7-4abd-b4dd-39c0f5d2cff7"),
                        Name: codataccounting.String("Kurt Abernathy"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(4104.92),
                    DiscountPercentage: codataccounting.Float64(1369),
                    ItemRef: &shared.ItemRef{
                        ID: "6d436813-f16d-49f5-bce6-c556146c3e25",
                        Name: codataccounting.String("Mr. Elsa Reinger"),
                    },
                    Quantity: 7705.81,
                    SubTotal: codataccounting.Float64(3045.82),
                    TaxAmount: codataccounting.Float64(1469.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8828.6),
                        ID: codataccounting.String("141aac36-6c8d-4d6b-9442-907474778a7b"),
                        Name: codataccounting.String("Bernard Kerluke"),
                    },
                    TotalAmount: codataccounting.Float64(1811.51),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c10ab3cd-ca42-4519-84e5-23c7e0bc7178",
                                Name: codataccounting.String("Tom Kuhn"),
                            },
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
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ex"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ad"),
                        Currency: codataccounting.String("expedita"),
                        CurrencyRate: codataccounting.Float64(299.5),
                        TotalAmount: codataccounting.Float64(5615.77),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b61891ba-a0fe-41ad-a008-e6f8c5f350d8"),
                            Name: codataccounting.String("Taylor Reichel"),
                        },
                        Currency: codataccounting.String("dolor"),
                        CurrencyRate: codataccounting.Float64(3073.76),
                        ID: codataccounting.String("18143010-4218-413d-9208-ece7e253b668"),
                        Note: codataccounting.String("magnam"),
                        PaidOnDate: codataccounting.String("exercitationem"),
                        Reference: codataccounting.String("ab"),
                        TotalAmount: codataccounting.Float64(7814.8),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("autem"),
                        Currency: codataccounting.String("nobis"),
                        CurrencyRate: codataccounting.Float64(3883.19),
                        TotalAmount: codataccounting.Float64(9272.12),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("205e16de-ab3f-4ec9-978a-64584273a841"),
                            Name: codataccounting.String("Clint Carroll"),
                        },
                        Currency: codataccounting.String("consectetur"),
                        CurrencyRate: codataccounting.Float64(468.06),
                        ID: codataccounting.String("9fb09299-21ae-4fb9-b58c-4d86e68e4be0"),
                        Note: codataccounting.String("ipsam"),
                        PaidOnDate: codataccounting.String("vel"),
                        Reference: codataccounting.String("alias"),
                        TotalAmount: codataccounting.Float64(938.94),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(2476.85),
            SourceModifiedDate: codataccounting.String("maiores"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sint": map[string]interface{}{
                        "deserunt": "esse",
                        "nemo": "reprehenderit",
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
            TotalAmount: 805.78,
            TotalDiscount: 7820.9,
            TotalTaxAmount: 3041.98,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 753.59,
                    Name: "Dr. Gina Jaskolski",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(431994),
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
        BillCreditNoteID: "velit",
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
        Query: codataccounting.String("ut"),
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
            AllocatedOnDate: codataccounting.String("perspiciatis"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("earum"),
            CurrencyRate: codataccounting.Float64(1175.25),
            DiscountPercentage: 7722.66,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("voluptatibus"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e06e3a43-7000-4ae6-b6bc-9b8f759eac55"),
                        Name: codataccounting.String("Jeremiah Koch PhD"),
                    },
                    Description: codataccounting.String("consectetur"),
                    DiscountAmount: codataccounting.Float64(1124.27),
                    DiscountPercentage: codataccounting.Float64(813.71),
                    ItemRef: &shared.ItemRef{
                        ID: "352965bb-8a72-4026-9143-5e139dbc2259",
                        Name: codataccounting.String("Gerald Ondricka"),
                    },
                    Quantity: 6374.62,
                    SubTotal: codataccounting.Float64(5546.03),
                    TaxAmount: codataccounting.Float64(8119.39),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(257.56),
                        ID: codataccounting.String("70e1084c-b067-42d1-ad87-9eeb9665b85e"),
                        Name: codataccounting.String("Mr. Jonathon Swaniawski"),
                    },
                    TotalAmount: codataccounting.Float64(6841.26),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "0be2d782-259e-43ea-8b51-97f92443da7c",
                                Name: codataccounting.String("Lewis Denesik"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "95c537c6-454e-4fb0-b348-96c3ca5acfbe",
                                Name: codataccounting.String("Winifred Streich"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "07577929-177d-4eac-a46e-cb573409e3eb",
                                Name: codataccounting.String("Lila Harvey"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b12eb07f-116d-4b99-945f-c95fa88970e1",
                                Name: codataccounting.String("Nick Shields"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("velit"),
                            ID: "0fcb33ea-055b-4197-8d44-e2f52d82d351",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b6f48b65-6bcd-4b35-bf2e-4b27537a8cd9",
                            Name: codataccounting.String("Miss Kelly Ernser"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "77d525f7-7b11-44ee-b52f-f785fc37814d",
                            Name: codataccounting.String("Alexandra McLaughlin"),
                        },
                    },
                    UnitAmount: 96.83,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c2bb89eb-75da-4d63-ac60-0503d8bb3118"),
                        Name: codataccounting.String("Lana Kris"),
                    },
                    Description: codataccounting.String("dolorum"),
                    DiscountAmount: codataccounting.Float64(8979.56),
                    DiscountPercentage: codataccounting.Float64(5928.8),
                    ItemRef: &shared.ItemRef{
                        ID: "e057eb80-9e28-4103-b1f3-981d4c700b60",
                        Name: codataccounting.String("Kristie Frami"),
                    },
                    Quantity: 2313.82,
                    SubTotal: codataccounting.Float64(7532.4),
                    TaxAmount: codataccounting.Float64(4901.1),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2358.34),
                        ID: codataccounting.String("b9da3f2c-eda7-4e23-b225-7411faf4b754"),
                        Name: codataccounting.String("Casey Gleason PhD"),
                    },
                    TotalAmount: codataccounting.Float64(5258.09),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2857a5b4-0463-4a7d-975f-1400e764ad73",
                                Name: codataccounting.String("Bertha Ward MD"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("iusto"),
                            ID: "81b36a08-088d-4100-afad-a200ef0422eb",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "64cf9ab8-366c-4723-bfda-9e06bee4825c",
                            Name: codataccounting.String("Dr. Shari Roob Sr."),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c80bff91-8544-4ec4-adef-cce8f1977773",
                            Name: codataccounting.String("Ben Durgan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2a7b408f-05e3-4d48-bdaf-313a1f5fd942",
                            Name: codataccounting.String("Miss Misty Ruecker"),
                        },
                    },
                    UnitAmount: 4124.33,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f25ea944-f3b7-456c-91f6-c37a51262438"),
                        Name: codataccounting.String("Annette Quitzon"),
                    },
                    Description: codataccounting.String("sit"),
                    DiscountAmount: codataccounting.Float64(3634.82),
                    DiscountPercentage: codataccounting.Float64(6339.87),
                    ItemRef: &shared.ItemRef{
                        ID: "23a45cef-c5fd-4e10-a0ce-2169e510019c",
                        Name: codataccounting.String("Doreen Schmeler"),
                    },
                    Quantity: 2467.72,
                    SubTotal: codataccounting.Float64(2991.53),
                    TaxAmount: codataccounting.Float64(4935.91),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3884.04),
                        ID: codataccounting.String("2799bfbb-e694-49fb-abb4-ecae6c3d5db3"),
                        Name: codataccounting.String("Kristopher Walter"),
                    },
                    TotalAmount: codataccounting.Float64(3233.65),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "aea4c506-a8aa-494c-8264-4cf5e9d9a457",
                                Name: codataccounting.String("Grant Schultz MD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c600dec0-01ac-4802-a2ec-09ff8f0f816f",
                                Name: codataccounting.String("Lee Gibson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c13e902c-1412-45b0-960a-668151a472af",
                                Name: codataccounting.String("Jeremy Douglas"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "949f83f3-50cf-4876-bfb9-01c6ecbb4e24",
                                Name: codataccounting.String("Marianne Zemlak"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("sint"),
                            ID: "ffafeda5-3e5a-4e6e-8ac1-84c2b9c247c8",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "73a40e19-42f3-42e5-9055-756f5d56d0bd",
                            Name: codataccounting.String("Amelia Williamson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e13db4f6-2cba-43f8-941a-ebc0b80a6924",
                            Name: codataccounting.String("Rodney Prohaska"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cfcc8f89-5010-4f5d-93d6-fa1804e54c82",
                            Name: codataccounting.String("Walter Jacobs"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "363c8873-e484-4380-b1f6-b8ca275a60a0",
                            Name: codataccounting.String("Jody Gutmann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cc699171-b51c-41bd-b1cf-4b888ebdfc4c",
                            Name: codataccounting.String("Lowell Olson"),
                        },
                    },
                    UnitAmount: 7268.51,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("quo"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("hic"),
                        Currency: codataccounting.String("maxime"),
                        CurrencyRate: codataccounting.Float64(366.91),
                        TotalAmount: codataccounting.Float64(7435.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("2dce1087-3e42-4b00-ad67-8878ba8581a5"),
                            Name: codataccounting.String("Martin Berge"),
                        },
                        Currency: codataccounting.String("enim"),
                        CurrencyRate: codataccounting.Float64(2864.53),
                        ID: codataccounting.String("fefa9c95-f2ea-4c55-a5d3-07cfee81206e"),
                        Note: codataccounting.String("sed"),
                        PaidOnDate: codataccounting.String("deleniti"),
                        Reference: codataccounting.String("sunt"),
                        TotalAmount: codataccounting.Float64(2001.9),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("delectus"),
                        Currency: codataccounting.String("laborum"),
                        CurrencyRate: codataccounting.Float64(3034.21),
                        TotalAmount: codataccounting.Float64(6442.23),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("41c480d3-f213-42af-8310-2d514f4cc6f1"),
                            Name: codataccounting.String("Rudolph Welch"),
                        },
                        Currency: codataccounting.String("sed"),
                        CurrencyRate: codataccounting.Float64(1066.82),
                        ID: codataccounting.String("a6a4f77a-87ee-43e4-be75-2c65b34418e3"),
                        Note: codataccounting.String("expedita"),
                        PaidOnDate: codataccounting.String("libero"),
                        Reference: codataccounting.String("iste"),
                        TotalAmount: codataccounting.Float64(749.21),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(7924.99),
            SourceModifiedDate: codataccounting.String("quos"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sint": map[string]interface{}{
                        "enim": "accusamus",
                        "aperiam": "voluptates",
                    },
                    "laudantium": map[string]interface{}{
                        "quae": "omnis",
                        "illum": "rem",
                    },
                    "tenetur": map[string]interface{}{
                        "modi": "earum",
                        "architecto": "aliquam",
                        "labore": "maiores",
                    },
                    "sequi": map[string]interface{}{
                        "consequatur": "esse",
                        "debitis": "facere",
                        "quisquam": "cumque",
                        "aliquam": "dolorum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "a5f3cabd-905a-4972-a056-728227b2d309",
                SupplierName: codataccounting.String("dolore"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 4570.63,
            TotalTaxAmount: 380.44,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9565.45,
                    Name: "Monique Hackett",
                },
                shared.WithholdingTaxitems{
                    Amount: 5314.94,
                    Name: "Jody Wolff",
                },
                shared.WithholdingTaxitems{
                    Amount: 3538.19,
                    Name: "Shane Yundt",
                },
            },
        },
        BillCreditNoteID: "corporis",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(252567),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```
