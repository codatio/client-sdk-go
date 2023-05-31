# CreditNotes

## Overview

Credit notes

### Available Operations

* [Create](#create) - Create credit note
* [Get](#get) - Get credit note
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update credit note model
* [List](#list) - List credit notes
* [Update](#update) - Update creditNote

## Create

Push credit note


Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) to see which integrations support this endpoint.


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
    res, err := s.CreditNotes.Create(ctx, operations.CreateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(4169.63),
            AdditionalTaxPercentage: codataccounting.Float64(5118.07),
            AllocatedOnDate: codataccounting.String("est"),
            CreditNoteNumber: codataccounting.String("consequatur"),
            Currency: codataccounting.String("incidunt"),
            CurrencyRate: codataccounting.Float64(1748.36),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("labore"),
                ID: "e00a1d6e-b943-4464-9d03-084fbba5ccef",
            },
            DiscountPercentage: 9488.1,
            ID: codataccounting.String("5cb01fe5-1e52-48a4-9ac8-2b85f8bc2cab"),
            IssueDate: codataccounting.String("laborum"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("da4127dd-597f-4f47-91aa-1bc74b86cecc"),
                        Name: codataccounting.String("Eva Wolf"),
                    },
                    Description: codataccounting.String("nam"),
                    DiscountAmount: codataccounting.Float64(2787.91),
                    DiscountPercentage: codataccounting.Float64(5341.75),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "48bd6a6f-0441-4d2c-bb80-8094373e0604",
                        Name: codataccounting.String("Velma Rogahn"),
                    },
                    Quantity: 7386.39,
                    SubTotal: codataccounting.Float64(6653.11),
                    TaxAmount: codataccounting.Float64(8566.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(214.7),
                        ID: codataccounting.String("2f2586bc-f152-4558-9aa9-5be6cd02756c"),
                        Name: codataccounting.String("Regina Grady"),
                    },
                    TotalAmount: codataccounting.Float64(2877.97),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2b47e176-3c52-408c-a3e9-802d82f0d45e",
                                Name: codataccounting.String("Ray Pfannerstill"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ea"),
                            ID: "74ee5cfc-18ed-4c7f-b87e-32e04b3d3ed0",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "670ef42b-d3c9-4f1c-8503-f6c39bcd0a62",
                            Name: codataccounting.String("Edward Wolf"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f385189a-d7ef-4807-aae0-3f33ca79fb9d",
                            Name: codataccounting.String("Francis Anderson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ba26fd36-8ba9-4216-bcb4-15835c736417",
                            Name: codataccounting.String("Diana Bogisich"),
                        },
                    },
                    UnitAmount: 9262.29,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("dc046bc5-163b-4bca-8922-7c42c22c5535"),
                        Name: codataccounting.String("Eva McDermott"),
                    },
                    Description: codataccounting.String("quis"),
                    DiscountAmount: codataccounting.Float64(8177.29),
                    DiscountPercentage: codataccounting.Float64(6963.68),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b3c57c1e-4981-4e8a-a257-ddc1912ebde6",
                        Name: codataccounting.String("Juana Williamson"),
                    },
                    Quantity: 3303.58,
                    SubTotal: codataccounting.Float64(3003.21),
                    TaxAmount: codataccounting.Float64(4362.35),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6168.21),
                        ID: codataccounting.String("d4015dfa-7962-406b-af2b-0a3e42c1aa01"),
                        Name: codataccounting.String("Laverne Mante"),
                    },
                    TotalAmount: codataccounting.Float64(7755.85),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e9135586-d18f-49f9-ba4b-fad2bf7d67ca",
                                Name: codataccounting.String("Edwin Olson"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("unde"),
                            ID: "b41d6124-3531-4870-8f68-b03ad421bd43",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f0cb0a00-03eb-422d-9b3a-70d94faa741c",
                            Name: codataccounting.String("Dr. Lydia Spencer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c2050d38-dc3c-4e18-9472-f9ee69166a8b",
                            Name: codataccounting.String("Jeff Grimes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "eac8b3a2-875c-46c1-be60-6d07d2a9c87a",
                            Name: codataccounting.String("Herman Barrows III"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "661a1d91-36a7-4e8d-9321-3f3f658752db",
                            Name: codataccounting.String("Stacey Haag"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9f0a56ce-bcad-4a29-8a79-181c95671663",
                            Name: codataccounting.String("Miss Tommy Emard"),
                        },
                    },
                    UnitAmount: 4263.08,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("65163a36-3851-42ab-a521-b9f2e072467b"),
                        Name: codataccounting.String("Miss Homer Gislason"),
                    },
                    Description: codataccounting.String("sit"),
                    DiscountAmount: codataccounting.Float64(3356.88),
                    DiscountPercentage: codataccounting.Float64(9426.04),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "ab0d650e-df22-4a94-920e-c90ea41d1f46",
                        Name: codataccounting.String("Jenna Lakin II"),
                    },
                    Quantity: 3916.67,
                    SubTotal: codataccounting.Float64(9957.33),
                    TaxAmount: codataccounting.Float64(9573.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9524.11),
                        ID: codataccounting.String("73fdf54f-dd5e-4a95-8339-8dafb42a8d63"),
                        Name: codataccounting.String("Marsha Lueilwitz"),
                    },
                    TotalAmount: codataccounting.Float64(8648.87),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "039ea5f9-b18a-4244-bd61-9039dacd38ed",
                                Name: codataccounting.String("Krystal Rolfson"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "1dc7f1e3-af15-4920-890d-1b4901f2bd89",
                                Name: codataccounting.String("Everett Ondricka"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "639da5b7-b690-42b8-81a9-4f643664a8f0",
                                Name: codataccounting.String("Marlon Leffler"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("natus"),
                            ID: "1d732d9f-baf9-4476-a2ae-8dcc50c8a351",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "73784893-0750-4a00-a966-ec736d431943",
                            Name: codataccounting.String("Sidney Ruecker"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c92398ed-3d3a-4b7c-a3c5-ca8649a70cfd",
                            Name: codataccounting.String("Freda Johnson"),
                        },
                    },
                    UnitAmount: 6125.84,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("distinctio"),
            Note: codataccounting.String("in"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("voluptatem"),
                        Currency: codataccounting.String("voluptas"),
                        CurrencyRate: codataccounting.Float64(2991.8),
                        TotalAmount: codataccounting.Float64(3189),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1077d19e-a83d-4492-ad14-b8a2c1954545"),
                            Name: codataccounting.String("Austin Hermann"),
                        },
                        Currency: codataccounting.String("cumque"),
                        CurrencyRate: codataccounting.Float64(7659),
                        ID: codataccounting.String("185ea490-1c7c-443a-92da-a784aba3d230"),
                        Note: codataccounting.String("voluptates"),
                        PaidOnDate: codataccounting.String("nulla"),
                        Reference: codataccounting.String("tenetur"),
                        TotalAmount: codataccounting.Float64(4923.66),
                    },
                },
            },
            RemainingCredit: 2200.49,
            SourceModifiedDate: codataccounting.String("totam"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: 1147.62,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "beatae": map[string]interface{}{
                        "veniam": "non",
                    },
                    "laudantium": map[string]interface{}{
                        "rerum": "nulla",
                    },
                    "ducimus": map[string]interface{}{
                        "repellendus": "enim",
                        "voluptas": "veniam",
                        "voluptatem": "quam",
                        "vel": "aspernatur",
                    },
                },
            },
            TotalAmount: 850.02,
            TotalDiscount: 7810.03,
            TotalTaxAmount: 3220.54,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 9851.55,
                    Name: "Janis Kshlerin",
                },
                shared.WithholdingTaxitems{
                    Amount: 3759.62,
                    Name: "Natalie Gleichner",
                },
                shared.WithholdingTaxitems{
                    Amount: 334.57,
                    Name: "Mr. Brian Kihn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(608208),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCreditNoteResponse != nil {
        // handle response
    }
}
```

## Get

﻿Gets a single creditNote corresponding to the given ID.

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
    res, err := s.CreditNotes.Get(ctx, operations.GetCreditNoteRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CreditNoteID: "commodi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

﻿Get create/update credit note model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating and updating a credit note.

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
    res, err := s.CreditNotes.GetCreateUpdateModel(ctx, operations.GetCreateUpdateCreditNotesModelRequest{
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

﻿Gets a list of all credit notes for a company, with pagination.

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
    res, err := s.CreditNotes.List(ctx, operations.ListCreditNotesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dicta"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNotes != nil {
        // handle response
    }
}
```

## Update

﻿Posts an updated credit note to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support updating a credit note.

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
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNote: &shared.CreditNote{
            AdditionalTaxAmount: codataccounting.Float64(8489.26),
            AdditionalTaxPercentage: codataccounting.Float64(1334.52),
            AllocatedOnDate: codataccounting.String("ut"),
            CreditNoteNumber: codataccounting.String("deserunt"),
            Currency: codataccounting.String("dignissimos"),
            CurrencyRate: codataccounting.Float64(8604.11),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("facilis"),
                ID: "b8f532d8-92cf-4781-acb5-12c878240bf5",
            },
            DiscountPercentage: 2743.68,
            ID: codataccounting.String("8f88f8f1-bf0b-4c8e-9f20-6d5d831d0081"),
            IssueDate: codataccounting.String("voluptatem"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("0f670667-3f3a-4681-8576-8dce742409a2"),
                        Name: codataccounting.String("Ms. Annette Towne"),
                    },
                    Description: codataccounting.String("aut"),
                    DiscountAmount: codataccounting.Float64(1141.51),
                    DiscountPercentage: codataccounting.Float64(2710.98),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "89a5f63e-3af3-4dd9-9da3-3dcd63483e4a",
                        Name: codataccounting.String("Amelia Morissette"),
                    },
                    Quantity: 2688.31,
                    SubTotal: codataccounting.Float64(8278.13),
                    TaxAmount: codataccounting.Float64(9784.16),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2045.17),
                        ID: codataccounting.String("7e45b895-5d41-43e1-ba48-2310907bd354"),
                        Name: codataccounting.String("Charles McGlynn"),
                    },
                    TotalAmount: codataccounting.Float64(8744.46),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "34f02449-d86f-44bb-a0fe-5d911cbfe749",
                                Name: codataccounting.String("Omar Wolf"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a27f69e2-c9e6-4d10-a9db-3ad4c6b03108",
                                Name: codataccounting.String("Seth Schaefer"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("in"),
                            ID: "473082b9-4f2a-4b1f-9567-1e9c326350a4",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "143789ce-0e99-4159-8d93-a74c0252fe3b",
                            Name: codataccounting.String("Bridget Greenholt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b778ebb6-e1d2-4cf5-82ba-fb2cbc4635d5",
                            Name: codataccounting.String("Andre Heidenreich"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "028c3e95-1a1e-430f-9a96-6489d7b78673",
                            Name: codataccounting.String("Peter Feeney Sr."),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a6b99249-4594-4487-b5c8-43836b86b3cd",
                            Name: codataccounting.String("Mrs. Marc Grimes"),
                        },
                    },
                    UnitAmount: 223.76,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("449f9df1-3f4e-4edb-a78b-f606825894ea"),
                        Name: codataccounting.String("Jeanne Farrell"),
                    },
                    Description: codataccounting.String("impedit"),
                    DiscountAmount: codataccounting.Float64(4457.32),
                    DiscountPercentage: codataccounting.Float64(1479.33),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "795b7851-48d6-4d54-9e56-35b33bc0f970",
                        Name: codataccounting.String("Micheal Cummings"),
                    },
                    Quantity: 5918.35,
                    SubTotal: codataccounting.Float64(9624.68),
                    TaxAmount: codataccounting.Float64(2930.13),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5170.23),
                        ID: codataccounting.String("44225e75-b796-4065-80ef-a6f93b90a1b8"),
                        Name: codataccounting.String("Austin Hartmann"),
                    },
                    TotalAmount: codataccounting.Float64(1147.52),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "54b739f4-fe77-4210-91f6-558c99c722d2",
                                Name: codataccounting.String("Lucas Barton"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("numquam"),
                            ID: "087d9caa-e042-4dd7-8aac-9b4caa1cfe9e",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "df903907-f378-4319-83d4-2e54a8546659",
                            Name: codataccounting.String("Mr. Kara Hintz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c1471d51-aaa6-4ddf-9abd-6487c5fc2b86",
                            Name: codataccounting.String("Miss Monique Bartell"),
                        },
                    },
                    UnitAmount: 9832.72,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("69e10015-7630-4bda-bafd-ed84a35a4123"),
                        Name: codataccounting.String("Stewart Bode"),
                    },
                    Description: codataccounting.String("sequi"),
                    DiscountAmount: codataccounting.Float64(3217.24),
                    DiscountPercentage: codataccounting.Float64(6847.17),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c26ae33b-ef97-41a8-b46b-ca1106fe965b",
                        Name: codataccounting.String("Evelyn Braun IV"),
                    },
                    Quantity: 7565.71,
                    SubTotal: codataccounting.Float64(9475.73),
                    TaxAmount: codataccounting.Float64(5358.76),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5373.21),
                        ID: codataccounting.String("ec9f7b99-a550-4a65-aed3-33bb0ce8aa65"),
                        Name: codataccounting.String("Victoria Denesik"),
                    },
                    TotalAmount: codataccounting.Float64(5250.89),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "eb7e14ca-5640-4891-9009-7019a48f88ec",
                                Name: codataccounting.String("Ron Ratke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "04e01105-d389-4081-a2c6-beb68a0f657b",
                                Name: codataccounting.String("Teri Abshire"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("beatae"),
                            ID: "480f8de3-0f06-49d8-9061-8d97e1522975",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "da803122-92cc-461c-aa70-2bb97ee102da",
                            Name: codataccounting.String("Hope Walsh"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8e01bf33-eaab-4454-82ac-1704bf1cc9fc",
                            Name: codataccounting.String("Katherine Oberbrunner"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5eb5f0c4-92b5-4744-908a-2267aaee79e3",
                            Name: codataccounting.String("Fernando Blanda"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "31becb83-d237-48ae-bbfc-23d9450a986a",
                            Name: codataccounting.String("Faye Hammes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c707f06b-28ec-4c86-8923-86f62c969c4c",
                            Name: codataccounting.String("Rick Predovic"),
                        },
                    },
                    UnitAmount: 5207.16,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("error"),
            Note: codataccounting.String("alias"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sequi"),
                        Currency: codataccounting.String("sapiente"),
                        CurrencyRate: codataccounting.Float64(8471.37),
                        TotalAmount: codataccounting.Float64(2026.93),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("c81da10f-8c23-4df9-b1da-3edb51fad94a"),
                            Name: codataccounting.String("Roosevelt Marvin"),
                        },
                        Currency: codataccounting.String("ullam"),
                        CurrencyRate: codataccounting.Float64(1158.79),
                        ID: codataccounting.String("37726d15-321b-4832-a56d-69180ff60eb9"),
                        Note: codataccounting.String("fuga"),
                        PaidOnDate: codataccounting.String("ex"),
                        Reference: codataccounting.String("autem"),
                        TotalAmount: codataccounting.Float64(3460.72),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("atque"),
                        Currency: codataccounting.String("saepe"),
                        CurrencyRate: codataccounting.Float64(4320.55),
                        TotalAmount: codataccounting.Float64(5652.45),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a4b843d3-82db-4ec7-9c68-c60659468ce3"),
                            Name: codataccounting.String("Anita Spinka"),
                        },
                        Currency: codataccounting.String("magnam"),
                        CurrencyRate: codataccounting.Float64(5876.67),
                        ID: codataccounting.String("bf8214c3-37f9-46bb-8c69-e372db1344ba"),
                        Note: codataccounting.String("error"),
                        PaidOnDate: codataccounting.String("doloribus"),
                        Reference: codataccounting.String("reprehenderit"),
                        TotalAmount: codataccounting.Float64(5256.79),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("est"),
                        Currency: codataccounting.String("quis"),
                        CurrencyRate: codataccounting.Float64(7709.69),
                        TotalAmount: codataccounting.Float64(343.42),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ed7aab62-e972-461f-b0c5-8d27b51996b5"),
                            Name: codataccounting.String("Joel Rath DVM"),
                        },
                        Currency: codataccounting.String("saepe"),
                        CurrencyRate: codataccounting.Float64(9596.73),
                        ID: codataccounting.String("712b7a7a-b034-44b1-b106-88deebef897f"),
                        Note: codataccounting.String("velit"),
                        PaidOnDate: codataccounting.String("fugiat"),
                        Reference: codataccounting.String("pariatur"),
                        TotalAmount: codataccounting.Float64(276.76),
                    },
                },
            },
            RemainingCredit: 7998.57,
            SourceModifiedDate: codataccounting.String("minus"),
            Status: shared.CreditNoteStatusPartiallyPaid,
            SubTotal: 2232.35,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tenetur": map[string]interface{}{
                        "dicta": "rerum",
                    },
                },
            },
            TotalAmount: 2054.73,
            TotalDiscount: 9131.52,
            TotalTaxAmount: 3033.65,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 30.91,
                    Name: "Richard Nitzsche Jr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 2752.88,
                    Name: "Cassandra Kerluke",
                },
                shared.WithholdingTaxitems{
                    Amount: 4841.4,
                    Name: "Mr. Sherri Torphy",
                },
                shared.WithholdingTaxitems{
                    Amount: 2244.51,
                    Name: "Laura Cronin",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "cumque",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(513307),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```
