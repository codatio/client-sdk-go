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
            AdditionalTaxAmount: codataccounting.Float64(5015.91),
            AdditionalTaxPercentage: codataccounting.Float64(6138.48),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("distinctio"),
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(2875.03),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("blanditiis"),
                ID: "73f5033f-19db-4f12-9ce4-152eab9cd7e5",
            },
            DiscountPercentage: 1419.86,
            ID: codataccounting.String("24a6a0e1-23b7-4847-ac59-e1f67f3c4cce"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6d7696ff-3c57-4475-8135-7e44f51f8b08"),
                        Name: codataccounting.String("Miss Roxanne Fahey"),
                    },
                    Description: codataccounting.String("officiis"),
                    DiscountAmount: codataccounting.Float64(1182.36),
                    DiscountPercentage: codataccounting.Float64(5664.44),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3a245467-f948-474c-ad5c-c4972233e66b",
                        Name: codataccounting.String("Wallace Wolff"),
                    },
                    Quantity: 8309.09,
                    SubTotal: codataccounting.Float64(129.17),
                    TaxAmount: codataccounting.Float64(316.05),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7277.89),
                        ID: codataccounting.String("979ef203-8732-4059-8ccc-1096400313b3"),
                        Name: codataccounting.String("Tim Becker"),
                    },
                    TotalAmount: codataccounting.Float64(9622.8),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "5fe72dc4-077d-40cc-bf40-8efc15ceb4d6",
                                Name: codataccounting.String("Ryan Tillman"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0f75aedf-2aca-4b58-b991-c926ddb58946",
                                Name: codataccounting.String("Leigh Kuhic"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("inventore"),
                            ID: "cbe6d950-2f0e-4a93-8b69-f7ac2f72f885",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "90491160-8207-4888-ac66-183bfe9659eb",
                            Name: codataccounting.String("Laura Thompson III"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "af75b0b5-32a4-4da3-bcba-af4452c4842c",
                            Name: codataccounting.String("Garry Conn"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "32dafe81-a88f-4444-8573-fecd47353f63",
                            Name: codataccounting.String("Miss Isaac Cremin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "79aa69cd-5fbc-4f79-9a18-a7822bf95894",
                            Name: codataccounting.String("Tyler Lakin MD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "db55f9e5-d751-4c9f-a8f7-502bfdc34508",
                            Name: codataccounting.String("Ms. Pamela Wilkinson"),
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
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
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
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
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
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("laborum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8037.63),
                        TotalAmount: codataccounting.Float64(3236.14),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("dddb46aa-1cfd-46d8-a8da-013191129646"),
                            Name: codataccounting.String("Bertha Halvorson PhD"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(932.99),
                        ID: codataccounting.String("f29042f5-69b7-4aff-8ea2-216cbe071bc1"),
                        Note: codataccounting.String("ex"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("officiis"),
                        TotalAmount: codataccounting.Float64(1493.76),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(6872.4),
                        TotalAmount: codataccounting.Float64(1975.4),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b084da99-257d-404f-8084-7a742d84496c"),
                            Name: codataccounting.String("Carroll Wehner"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(4056.34),
                        ID: codataccounting.String("b99bc635-62eb-4fdf-95c2-94c060b06a12"),
                        Note: codataccounting.String("quos"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("voluptate"),
                        TotalAmount: codataccounting.Float64(4210.18),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8921.14),
                        TotalAmount: codataccounting.Float64(9377.43),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6d0c6d6e-d9c7-43dd-a345-71509a8e870d"),
                            Name: codataccounting.String("Brooke Harris DVM"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8021.88),
                        ID: codataccounting.String("242c7b66-a1f3-40c7-bdf5-b6719890f42a"),
                        Note: codataccounting.String("tempora"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("expedita"),
                        TotalAmount: codataccounting.Float64(2678.39),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8712.98),
                        TotalAmount: codataccounting.Float64(5236.07),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("5b260591-d745-4e3c-a059-c9c3f567e0e2"),
                            Name: codataccounting.String("Louise Kuhn"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1074.24),
                        ID: codataccounting.String("d62fcdac-e1f0-4121-ace2-239e8f25cd0d"),
                        Note: codataccounting.String("inventore"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quibusdam"),
                        TotalAmount: codataccounting.Float64(5660.89),
                    },
                },
            },
            RemainingCredit: 3460.81,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusPartiallyPaid,
            SubTotal: 3012.89,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "natus": map[string]interface{}{
                        "velit": "sint",
                        "eos": "nisi",
                        "commodi": "impedit",
                        "facilis": "temporibus",
                    },
                },
            },
            TotalAmount: 6214.28,
            TotalDiscount: 3126.08,
            TotalTaxAmount: 9648.83,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 6464.56,
                    Name: "Fred Price",
                },
                shared.WithholdingTaxitems{
                    Amount: 930.67,
                    Name: "Diana Hudson",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(838287),
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
        CreditNoteID: "et",
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
        Query: codataccounting.String("debitis"),
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
            AdditionalTaxAmount: codataccounting.Float64(3931.89),
            AdditionalTaxPercentage: codataccounting.Float64(4000.44),
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codataccounting.String("quas"),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(8092),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("eligendi"),
                ID: "4596217c-2977-4676-b342-54038bfb5971",
            },
            DiscountPercentage: 8798.24,
            ID: codataccounting.String("98190557-389c-4edb-ac7f-da39594d66bc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e480632b-9954-4b6f-a220-6369828553cb"),
                        Name: codataccounting.String("Ms. Patricia Barton"),
                    },
                    Description: codataccounting.String("necessitatibus"),
                    DiscountAmount: codataccounting.Float64(9685.91),
                    DiscountPercentage: codataccounting.Float64(2775.69),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "921ec205-3b74-4936-aac8-ee0f2bf19588",
                        Name: codataccounting.String("Ray Batz I"),
                    },
                    Quantity: 9674.95,
                    SubTotal: codataccounting.Float64(2044.66),
                    TaxAmount: codataccounting.Float64(8285.71),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9049.68),
                        ID: codataccounting.String("ba297be3-e90b-4c40-9f86-8fd52405cb33"),
                        Name: codataccounting.String("Mercedes Fritsch"),
                    },
                    TotalAmount: codataccounting.Float64(9699.11),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f127fb0e-0bf1-4f82-9797-8d0acca77aeb",
                                Name: codataccounting.String("Mr. Bridget Koepp MD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "52046b64-e99f-4b0e-a7e0-94fdfed5540e",
                                Name: codataccounting.String("Dustin Doyle"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ut"),
                            ID: "a1b8fe99-731a-4dc0-9d85-ae2dfb70fb38",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "290d3365-61ec-4a16-af89-451bd76eeeb5",
                            Name: codataccounting.String("Jenny Rutherford"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1fad3551-2f06-4d4e-9b72-f0f548568a04",
                            Name: codataccounting.String("Mr. Lucille Weber"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1d6eb943-4645-4d03-884f-bba5cceff5cb",
                            Name: codataccounting.String("Martha Wisozk"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1e528a45-ac82-4b85-b8bc-2caba8da4127",
                            Name: codataccounting.String("Drew Hickle"),
                        },
                    },
                    UnitAmount: 9416.84,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f4711aa1-bc74-4b86-8ecc-74f77b4848bd"),
                        Name: codataccounting.String("Angie Johnston I"),
                    },
                    Description: codataccounting.String("aliquam"),
                    DiscountAmount: codataccounting.Float64(997.32),
                    DiscountPercentage: codataccounting.Float64(8603.83),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "2c3b8080-9437-43e0-a045-9bebbad02f25",
                        Name: codataccounting.String("Elmer Predovic"),
                    },
                    Quantity: 915.29,
                    SubTotal: codataccounting.Float64(3715.34),
                    TaxAmount: codataccounting.Float64(1768.52),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3465.57),
                        ID: codataccounting.String("58daa95b-e6cd-4027-96c3-54aa432b47e1"),
                        Name: codataccounting.String("Rosemary Fahey"),
                    },
                    TotalAmount: codataccounting.Float64(1292.57),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "8c23e980-2d82-4f0d-85eb-4a8b674ee5cf",
                                Name: codataccounting.String("Ralph Leffler"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("optio"),
                            ID: "7f787e32-e04b-43d3-ad0c-5670ef42bd3c",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "1cc503f6-c39b-4cd0-a629-0f957f385189",
                            Name: codataccounting.String("Woodrow Kreiger"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "07aae03f-33ca-479f-b9de-4032ba26fd36",
                            Name: codataccounting.String("Damon O'Kon"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "16bcb415-835c-4736-8172-3133edc046bc",
                            Name: codataccounting.String("Alice Kautzer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "bca49227-c42c-422c-9535-0495c5dbb3c5",
                            Name: codataccounting.String("Roxanne Buckridge"),
                        },
                    },
                    UnitAmount: 6128.36,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("81e8aa25-7ddc-4191-aebd-e64bfcc5469d"),
                        Name: codataccounting.String("Maria Bradtke"),
                    },
                    Description: codataccounting.String("reiciendis"),
                    DiscountAmount: codataccounting.Float64(6658.07),
                    DiscountPercentage: codataccounting.Float64(4582.2),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "96206bef-2b0a-43e4-ac1a-a010e9aac2e9",
                        Name: codataccounting.String("Josephine Hauck"),
                    },
                    Quantity: 4088.31,
                    SubTotal: codataccounting.Float64(8538.36),
                    TaxAmount: codataccounting.Float64(1161.94),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5559.88),
                        ID: codataccounting.String("f9f97a4b-fad2-4bf7-967c-a84ad99b41d6"),
                        Name: codataccounting.String("Sara Funk"),
                    },
                    TotalAmount: codataccounting.Float64(2008.24),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "870cf68b-03ad-4421-bd43-d1f0cb0a0003",
                                Name: codataccounting.String("Cesar Daugherty"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("sint"),
                            ID: "b3a70d94-faa7-441c-97d1-fedc2050d38d",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "ce185472-f9ee-4691-a6a8-be3444eac8b3",
                            Name: codataccounting.String("Billy Leannon"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6c1fe606-d07d-42a9-887a-e50c16661a1d",
                            Name: codataccounting.String("Harold Emmerich"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7e8d5321-3f3f-4658-b52d-b764c59f0a56",
                            Name: codataccounting.String("Terence Pfeffer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "da29ca79-181c-4956-b166-3c530b566516",
                            Name: codataccounting.String("Cecilia Friesen"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8512ab25-21b9-4f2e-8724-67b8a40bc05f",
                            Name: codataccounting.String("Wm Bartoletti"),
                        },
                    },
                    UnitAmount: 3385.42,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("itaque"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1794.45),
                        TotalAmount: codataccounting.Float64(6745.48),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("94d20ec9-0ea4-41d1-b465-e85156fff73f"),
                            Name: codataccounting.String("Irvin Heaney"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8498.18),
                        ID: codataccounting.String("5ea95433-98da-4fb4-aa8d-63388e4d8039"),
                        Note: codataccounting.String("necessitatibus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quis"),
                        TotalAmount: codataccounting.Float64(9910.9),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1219.33),
                        TotalAmount: codataccounting.Float64(5397.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a244fd61-9039-4dac-938e-d0dc671dc7f1"),
                            Name: codataccounting.String("Allen Nader II"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1603.05),
                        ID: codataccounting.String("0c90d1b4-901f-42bd-89c8-a32639da5b7b"),
                        Note: codataccounting.String("nisi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("eaque"),
                        TotalAmount: codataccounting.Float64(1555.84),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(5400.74),
                        TotalAmount: codataccounting.Float64(748.14),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a94f6436-64a8-4f0a-b8c6-91d732d9fbaf"),
                            Name: codataccounting.String("Clifford Klein"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(6562.84),
                        ID: codataccounting.String("e8dcc50c-8a35-412c-b378-48930750a00e"),
                        Note: codataccounting.String("sint"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("voluptas"),
                        TotalAmount: codataccounting.Float64(9074.2),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(2124.05),
                        TotalAmount: codataccounting.Float64(4311.08),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d4319439-8c78-43c9-a398-ed3d3ab7ca3c"),
                            Name: codataccounting.String("Alexandra Murray"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(6240.11),
                        ID: codataccounting.String("a70cfd5d-6989-4b72-8645-1077d19ea83d"),
                        Note: codataccounting.String("eius"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("consequuntur"),
                        TotalAmount: codataccounting.Float64(9357.71),
                    },
                },
            },
            RemainingCredit: 8426.89,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: 7320.16,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "fuga": map[string]interface{}{
                        "minus": "sunt",
                    },
                    "sint": map[string]interface{}{
                        "magnam": "veniam",
                        "quaerat": "minima",
                    },
                    "officiis": map[string]interface{}{
                        "ullam": "enim",
                        "facere": "cumque",
                        "cumque": "et",
                    },
                },
            },
            TotalAmount: 5100.79,
            TotalDiscount: 3247.86,
            TotalTaxAmount: 9120.76,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 3007.59,
                    Name: "Michael Carter",
                },
                shared.WithholdingTaxitems{
                    Amount: 7647.82,
                    Name: "Lillian Murray",
                },
                shared.WithholdingTaxitems{
                    Amount: 8589.87,
                    Name: "Matt Kohler",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "similique",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(747444),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```
