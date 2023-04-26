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
            AdditionalTaxAmount: codataccounting.Float64(4618.53),
            AdditionalTaxPercentage: codataccounting.Float64(5345.09),
            AllocatedOnDate: codataccounting.String("rem"),
            CreditNoteNumber: codataccounting.String("vel"),
            Currency: codataccounting.String("eos"),
            CurrencyRate: codataccounting.Float64(2864.64),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("sunt"),
                ID: "89eb4487-3f50-433f-99db-f125ce4152ea",
            },
            DiscountPercentage: 7125.23,
            ID: codataccounting.String("9cd7e522-4a6a-40e1-a3b7-847ec59e1f67"),
            IssueDate: codataccounting.String("repellat"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c4cce4b6-d769-46ff-bc57-47501357e44f"),
                        Name: codataccounting.String("Jean Welch"),
                    },
                    Description: codataccounting.String("consequatur"),
                    DiscountAmount: codataccounting.Float64(5167.39),
                    DiscountPercentage: codataccounting.Float64(2725.18),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c3197e19-3a24-4546-bf94-874c2d5cc497",
                        Name: codataccounting.String("Norma Feest"),
                    },
                    Quantity: 4358.41,
                    SubTotal: codataccounting.Float64(3961.88),
                    TaxAmount: codataccounting.Float64(7385.92),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8537.01),
                        ID: codataccounting.String("8fe5d00b-979e-4f20-b873-20590ccc1096"),
                        Name: codataccounting.String("Sarah Bartell I"),
                    },
                    TotalAmount: codataccounting.Float64(6992.15),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e5044f65-fe72-4dc4-877d-0cc3f408efc1",
                                Name: codataccounting.String("Pat Upton"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("illum"),
                            ID: "6e1eae0f-75ae-4df2-acab-58b991c926dd",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "89461e74-21cb-4e6d-9502-f0ea930b69f7",
                            Name: codataccounting.String("Spencer Crooks"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f8850090-4911-4608-a078-88ec66183bfe",
                            Name: codataccounting.String("Ricardo Hermiston"),
                        },
                    },
                    UnitAmount: 6946.11,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("magnam"),
            Note: codataccounting.String("doloremque"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quod"),
                        Currency: codataccounting.String("sunt"),
                        CurrencyRate: codataccounting.Float64(3774.3),
                        TotalAmount: codataccounting.Float64(9382.3),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("af75b0b5-32a4-4da3-bcba-af4452c4842c"),
                            Name: codataccounting.String("Garry Conn"),
                        },
                        Currency: codataccounting.String("ipsum"),
                        CurrencyRate: codataccounting.Float64(1292.7),
                        ID: codataccounting.String("dafe81a8-8f44-4445-b3fe-cd47353f63c8"),
                        Note: codataccounting.String("sed"),
                        PaidOnDate: codataccounting.String("eaque"),
                        Reference: codataccounting.String("natus"),
                        TotalAmount: codataccounting.Float64(1912.02),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nihil"),
                        Currency: codataccounting.String("unde"),
                        CurrencyRate: codataccounting.Float64(6463.21),
                        TotalAmount: codataccounting.Float64(6621.84),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("69cd5fbc-f79d-4a18-a782-2bf95894e686"),
                            Name: codataccounting.String("Lynda Schuppe"),
                        },
                        Currency: codataccounting.String("quaerat"),
                        CurrencyRate: codataccounting.Float64(9830.67),
                        ID: codataccounting.String("9e5d751c-9fe8-4f75-82bf-dc3450841f17"),
                        Note: codataccounting.String("autem"),
                        PaidOnDate: codataccounting.String("dolore"),
                        Reference: codataccounting.String("eius"),
                        TotalAmount: codataccounting.Float64(3423.93),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ex"),
                        Currency: codataccounting.String("amet"),
                        CurrencyRate: codataccounting.Float64(4543.86),
                        TotalAmount: codataccounting.Float64(5653.04),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f3fb27e2-1f86-4265-bb36-fc6b9f587ce5"),
                            Name: codataccounting.String("Audrey Schimmel"),
                        },
                        Currency: codataccounting.String("ea"),
                        CurrencyRate: codataccounting.Float64(2569.41),
                        ID: codataccounting.String("1a8312e5-047b-44c2-9ccb-423abcdc91fa"),
                        Note: codataccounting.String("est"),
                        PaidOnDate: codataccounting.String("distinctio"),
                        Reference: codataccounting.String("fugiat"),
                        TotalAmount: codataccounting.Float64(8621.67),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("totam"),
                        Currency: codataccounting.String("praesentium"),
                        CurrencyRate: codataccounting.Float64(8857.21),
                        TotalAmount: codataccounting.Float64(4610.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1f6c4825-2d77-471e-bfd0-74009ef8d29d"),
                            Name: codataccounting.String("Raymond Schulist"),
                        },
                        Currency: codataccounting.String("aut"),
                        CurrencyRate: codataccounting.Float64(5912.2),
                        ID: codataccounting.String("7b5da08c-57fa-46c7-8a21-6e19bafeca61"),
                        Note: codataccounting.String("cupiditate"),
                        PaidOnDate: codataccounting.String("veritatis"),
                        Reference: codataccounting.String("aliquam"),
                        TotalAmount: codataccounting.Float64(5682.31),
                    },
                },
            },
            RemainingCredit: 5410.46,
            SourceModifiedDate: codataccounting.String("dicta"),
            Status: shared.CreditNoteStatusEnumDraft,
            SubTotal: 428.84,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "suscipit": map[string]interface{}{
                        "maiores": "delectus",
                        "quos": "id",
                    },
                    "officiis": map[string]interface{}{
                        "voluptate": "consequatur",
                    },
                    "itaque": map[string]interface{}{
                        "voluptatem": "dolor",
                        "distinctio": "quaerat",
                        "a": "neque",
                        "nihil": "recusandae",
                    },
                },
            },
            TotalAmount: 2538.55,
            TotalDiscount: 6520.13,
            TotalTaxAmount: 6515.04,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 3819.74,
                    Name: "Tim Hamill",
                },
                shared.WithholdingTaxitems{
                    Amount: 4116.69,
                    Name: "Maureen Dooley",
                },
                shared.WithholdingTaxitems{
                    Amount: 6762.74,
                    Name: "Mona Schaden",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(419585),
    }

    res, err := s.CreditNotes.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCreditNoteResponse != nil {
        // handle response
    }
}
```

## Get

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
        CreditNoteID: "praesentium",
    }

    res, err := s.CreditNotes.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

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

    res, err := s.CreditNotes.GetCreateUpdateModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## List

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
        Query: codataccounting.String("magni"),
    }

    res, err := s.CreditNotes.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNotes != nil {
        // handle response
    }
}
```

## Update

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
            AdditionalTaxAmount: codataccounting.Float64(7874.67),
            AdditionalTaxPercentage: codataccounting.Float64(7118.19),
            AllocatedOnDate: codataccounting.String("in"),
            CreditNoteNumber: codataccounting.String("eaque"),
            Currency: codataccounting.String("delectus"),
            CurrencyRate: codataccounting.Float64(5019.46),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("minus"),
                ID: "fd5fb6e9-1b9a-49f7-8846-e2c3309db053",
            },
            DiscountPercentage: 4343.82,
            ID: codataccounting.String("d9e75ca0-06f5-4392-811a-25a8bf92f974"),
            IssueDate: codataccounting.String("aspernatur"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ad9a9f8b-f822-4112-9359-d98387f7a79c"),
                        Name: codataccounting.String("Christian Corwin"),
                    },
                    Description: codataccounting.String("magni"),
                    DiscountAmount: codataccounting.Float64(2657.08),
                    DiscountPercentage: codataccounting.Float64(5319.67),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4da21729-f2ac-441e-b572-5f1169ac1e41",
                        Name: codataccounting.String("Wallace Oberbrunner"),
                    },
                    Quantity: 7720.48,
                    SubTotal: codataccounting.Float64(1609.09),
                    TaxAmount: codataccounting.Float64(2042.92),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9044.85),
                        ID: codataccounting.String("34f2dfa4-a197-4f6d-a922-151fe1712099"),
                        Name: codataccounting.String("Ronnie Donnelly"),
                    },
                    TotalAmount: codataccounting.Float64(9711.55),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "43d85443-9ee2-4244-a044-3bc154188c2f",
                                Name: codataccounting.String("Kristin Tillman"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "da7832ea-bd61-47c3-b0d5-1a44bf01bad8",
                                Name: codataccounting.String("Patricia Kerluke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ea"),
                            ID: "082bfbdc-41ff-45d4-a2ae-4fb5cb35d176",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f1edb783-59ec-4c5c-b860-f8cd580ba738",
                            Name: codataccounting.String("Kimberly Waters"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4447297c-d3b1-4dd3-bbce-247b7684eff5",
                            Name: codataccounting.String("Teresa Denesik"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "71cffbd0-eb74-4b84-a195-3b44bd3c4315",
                            Name: codataccounting.String("Sammy Fisher"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5953c001-1398-463a-a41e-6c31cc2f1fcb",
                            Name: codataccounting.String("Joan Schaefer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "41ffbe9c-bd79-45ee-a5e0-76cc7abf616e",
                            Name: codataccounting.String("Vernon Sauer III"),
                        },
                    },
                    UnitAmount: 2640.9,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1934b90f-2e09-4d19-92fc-2f9e2e105944"),
                        Name: codataccounting.String("Virgil Fahey"),
                    },
                    Description: codataccounting.String("sed"),
                    DiscountAmount: codataccounting.Float64(2405.55),
                    DiscountPercentage: codataccounting.Float64(4444.94),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a72f9084-9d6a-4ed4-aecb-7537cd9222c9",
                        Name: codataccounting.String("Elijah Hegmann"),
                    },
                    Quantity: 6152.06,
                    SubTotal: codataccounting.Float64(817.75),
                    TaxAmount: codataccounting.Float64(6258.15),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6644.91),
                        ID: codataccounting.String("bfa2e761-f0ca-44d4-96ef-1031e6899f0c"),
                        Name: codataccounting.String("Dr. Sharon Bednar"),
                    },
                    TotalAmount: codataccounting.Float64(1347.95),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "d55cc058-4a18-44d7-ad97-1fc820c65b03",
                                Name: codataccounting.String("Luz Rau"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0cc88518-7e4d-4e04-af28-c5dddb46aa1c",
                                Name: codataccounting.String("Clint Jast"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "28da0131-9112-4964-a645-c1d81f29042f",
                                Name: codataccounting.String("Vanessa Monahan"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "aff0ea22-16cb-4e07-9bc1-63e279a3b084",
                                Name: codataccounting.String("Lorenzo Monahan"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("veniam"),
                            ID: "7d04f408-47a7-442d-8449-6cbdeecf6b99",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "63562ebf-df55-4c29-8c06-0b06a1287764",
                            Name: codataccounting.String("Darrel Wehner"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c6d6ed9c-73dd-4634-9715-09a8e870d3c5",
                            Name: codataccounting.String("Juan Wolff"),
                        },
                    },
                    UnitAmount: 1690.72,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("42c7b66a-1f30-4c73-9f5b-6719890f42a4"),
                        Name: codataccounting.String("Malcolm Gleichner"),
                    },
                    Description: codataccounting.String("at"),
                    DiscountAmount: codataccounting.Float64(5236.07),
                    DiscountPercentage: codataccounting.Float64(3465.34),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b260591d-745e-43c2-859c-9c3f567e0e25",
                        Name: codataccounting.String("Courtney Hoeger"),
                    },
                    Quantity: 1074.24,
                    SubTotal: codataccounting.Float64(8313.04),
                    TaxAmount: codataccounting.Float64(4021.21),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1628.49),
                        ID: codataccounting.String("fcdace1f-0121-46ce-a239-e8f25cd0d19d"),
                        Name: codataccounting.String("Greg Mayer"),
                    },
                    TotalAmount: codataccounting.Float64(2363.72),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e39266cb-d95f-47aa-ab24-113695d1e669",
                                Name: codataccounting.String("Jerald Schowalter"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "596217c2-9776-4763-b425-4038bfb5971e",
                                Name: codataccounting.String("Casey Block II"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "57389ced-bac7-4fda-b959-4d66bc2ae480",
                                Name: codataccounting.String("Dawn Cole"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("occaecati"),
                            ID: "54b6fa22-0636-4982-8553-cb10006bef49",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Unknown,
                        IsRebilledTo: shared.BilledToTypeEnum1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "ec2053b7-4936-46ac-8ee0-f2bf19588d40",
                            Name: codataccounting.String("Richard Dietrich"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "eba297be-3e90-4bc4-8df8-68fd52405cb3",
                            Name: codataccounting.String("Evelyn Stracke"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2f4f127f-b0e0-4bf1-b821-7978d0acca77",
                            Name: codataccounting.String("Phil Rice"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7021a520-46b6-44e9-9fb0-e67e094fdfed",
                            Name: codataccounting.String("Dr. Yvonne Grimes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "53a34a1b-8fe9-4973-9adc-05d85ae2dfb7",
                            Name: codataccounting.String("Shawna Pouros"),
                        },
                    },
                    UnitAmount: 4837.74,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("non"),
            Note: codataccounting.String("magni"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("consequatur"),
                        Currency: codataccounting.String("illum"),
                        CurrencyRate: codataccounting.Float64(2378.75),
                        TotalAmount: codataccounting.Float64(2106.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6561eca1-6ef8-4945-9bd7-6eeeb518c4da"),
                            Name: codataccounting.String("Ollie Osinski"),
                        },
                        Currency: codataccounting.String("nemo"),
                        CurrencyRate: codataccounting.Float64(3353.52),
                        ID: codataccounting.String("12f06d4e-5b72-4f0f-9485-68a0424e00a1"),
                        Note: codataccounting.String("quibusdam"),
                        PaidOnDate: codataccounting.String("autem"),
                        Reference: codataccounting.String("voluptates"),
                        TotalAmount: codataccounting.Float64(7314.5),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("cupiditate"),
                        Currency: codataccounting.String("modi"),
                        CurrencyRate: codataccounting.Float64(1916.53),
                        TotalAmount: codataccounting.Float64(3059.65),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("645d0308-4fbb-4a5c-8eff-5cb01fe51e52"),
                            Name: codataccounting.String("Lyle Gislason"),
                        },
                        Currency: codataccounting.String("minus"),
                        CurrencyRate: codataccounting.Float64(5544.29),
                        ID: codataccounting.String("2b85f8bc-2cab-4a8d-a412-7dd597ff4711"),
                        Note: codataccounting.String("similique"),
                        PaidOnDate: codataccounting.String("id"),
                        Reference: codataccounting.String("et"),
                        TotalAmount: codataccounting.Float64(7167.79),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("porro"),
                        Currency: codataccounting.String("nihil"),
                        CurrencyRate: codataccounting.Float64(2567.42),
                        TotalAmount: codataccounting.Float64(7030.15),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("86cecc74-f77b-4484-8bd6-a6f0441d2c3b"),
                            Name: codataccounting.String("Ms. Daniel Langworth"),
                        },
                        Currency: codataccounting.String("nesciunt"),
                        CurrencyRate: codataccounting.Float64(4744.53),
                        ID: codataccounting.String("3e060459-bebb-4ad0-af25-86bcf152558d"),
                        Note: codataccounting.String("fuga"),
                        PaidOnDate: codataccounting.String("fuga"),
                        Reference: codataccounting.String("excepturi"),
                        TotalAmount: codataccounting.Float64(3583.94),
                    },
                },
            },
            RemainingCredit: 7230.31,
            SourceModifiedDate: codataccounting.String("itaque"),
            Status: shared.CreditNoteStatusEnumSubmitted,
            SubTotal: 7837.02,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "consequatur": map[string]interface{}{
                        "in": "enim",
                    },
                    "vel": map[string]interface{}{
                        "consectetur": "quis",
                        "ut": "est",
                        "fuga": "labore",
                        "adipisci": "ratione",
                    },
                    "cum": map[string]interface{}{
                        "nihil": "officiis",
                        "inventore": "esse",
                    },
                    "ex": map[string]interface{}{
                        "minus": "ad",
                    },
                },
            },
            TotalAmount: 1292.57,
            TotalDiscount: 27.7,
            TotalTaxAmount: 5026.86,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 1863.18,
                    Name: "Kate Metz Sr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 8344.99,
                    Name: "Dr. Sean Williamson",
                },
                shared.WithholdingTaxitems{
                    Amount: 3401.07,
                    Name: "Gerardo Gislason",
                },
                shared.WithholdingTaxitems{
                    Amount: 7482.66,
                    Name: "Glenda Grimes",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "ipsam",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(788995),
    }

    res, err := s.CreditNotes.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```
