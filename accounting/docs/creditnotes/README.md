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
            AdditionalTaxAmount: codataccounting.Float64(922.79),
            AdditionalTaxPercentage: codataccounting.Float64(9391.31),
            AllocatedOnDate: codataccounting.String("fuga"),
            CreditNoteNumber: codataccounting.String("est"),
            Currency: codataccounting.String("distinctio"),
            CurrencyRate: codataccounting.Float64(8546.46),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("nulla"),
                ID: "88e71f6c-4825-42d7-b71e-7fd074009ef8",
            },
            DiscountPercentage: 8172.49,
            ID: codataccounting.String("29de1dd7-097b-45da-88c5-7fa6c78a216e"),
            IssueDate: codataccounting.String("illo"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("bafeca61-9149-4814-8b64-ff8ae170ef03"),
                        Name: codataccounting.String("Jon Wilkinson"),
                    },
                    Description: codataccounting.String("recusandae"),
                    DiscountAmount: codataccounting.Float64(2538.55),
                    DiscountPercentage: codataccounting.Float64(6520.13),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "a8685559-6673-42aa-9dcb-6682cb70f8cf",
                        Name: codataccounting.String("Dustin Wilkinson"),
                    },
                    Quantity: 9347.07,
                    SubTotal: codataccounting.Float64(5800.8),
                    TaxAmount: codataccounting.Float64(1161.7),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6935.92),
                        ID: codataccounting.String("9a9f7484-6e2c-4330-9db0-536d9e75ca00"),
                        Name: codataccounting.String("Lana Hauck"),
                    },
                    TotalAmount: codataccounting.Float64(1500.91),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "11a25a8b-f92f-4974-a8ad-9a9f8bf82211",
                                Name: codataccounting.String("Erin Frami"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "d98387f7-a79c-4d72-8d24-84da21729f2a",
                                Name: codataccounting.String("Francis Block"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5725f116-9ac1-4e41-98a2-3c23e34f2dfa",
                                Name: codataccounting.String("Blanche Cassin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f6de9221-51fe-4171-a099-853e9f543d85",
                                Name: codataccounting.String("Elaine Fadel"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptates"),
                            ID: "22446044-3bc1-4541-88c2-f56e85da7832",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "bd617c3b-0d51-4a44-bf01-bad8706d4608",
                            Name: codataccounting.String("Juana Wuckert"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "41ff5d4e-2ae4-4fb5-8b35-d17638f1edb7",
                            Name: codataccounting.String("Jeffery Hilll"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cc5cb860-f8cd-4580-ba73-810e4fe44472",
                            Name: codataccounting.String("Cory Runolfsson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b1dd3bbc-e247-4b76-84ef-f50126d71cff",
                            Name: codataccounting.String("Bryant Anderson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "74b84219-53b4-44bd-bc43-159d33e5953c",
                            Name: codataccounting.String("Mrs. Ruth Carroll"),
                        },
                    },
                    UnitAmount: 5525.12,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("63aa41e6-c31c-4c2f-9fcb-51c9a41ffbe9"),
                        Name: codataccounting.String("Willis Smitham"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(9200.57),
                    DiscountPercentage: codataccounting.Float64(9358),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "65e076cc-7abf-4616-aa5c-71641934b90f",
                        Name: codataccounting.String("Rochelle Bailey"),
                    },
                    Quantity: 683,
                    SubTotal: codataccounting.Float64(6173.25),
                    TaxAmount: codataccounting.Float64(8228.62),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1699.28),
                        ID: codataccounting.String("fc2f9e2e-1059-444b-935d-237a72f90849"),
                        Name: codataccounting.String("Rick Olson"),
                    },
                    TotalAmount: codataccounting.Float64(3105.08),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ecb7537c-d922-42c9-bf57-491aabfa2e76",
                                Name: codataccounting.String("Essie Barton"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4d456ef1-031e-4689-9f0c-2001e22cd55c",
                                Name: codataccounting.String("Robert Hickle"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a184d76d-971f-4c82-8c65-b037bb8e0cc8",
                                Name: codataccounting.String("Dan Boehm"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("officiis"),
                            ID: "4de04af2-8c5d-4ddb-86aa-1cfd6d828da0",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "19112964-6645-4c1d-81f2-9042f569b7af",
                            Name: codataccounting.String("David Volkman"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "16cbe071-bc16-43e2-b9a3-b084da99257d",
                            Name: codataccounting.String("Amber Ziemann V"),
                        },
                    },
                    UnitAmount: 2756.61,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7a742d84-496c-4bde-acf6-b99bc63562eb"),
                        Name: codataccounting.String("Lionel Wisoky"),
                    },
                    Description: codataccounting.String("porro"),
                    DiscountAmount: codataccounting.Float64(1748.97),
                    DiscountPercentage: codataccounting.Float64(6102.13),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4c060b06-a128-4776-8eef-6d0c6d6ed9c7",
                        Name: codataccounting.String("Paulette Smitham"),
                    },
                    Quantity: 2855.44,
                    SubTotal: codataccounting.Float64(3419.34),
                    TaxAmount: codataccounting.Float64(4922.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(768.18),
                        ID: codataccounting.String("509a8e87-0d3c-45a1-b9c2-42c7b66a1f30"),
                        Name: codataccounting.String("Cody Franey"),
                    },
                    TotalAmount: codataccounting.Float64(3480.56),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "6719890f-42a4-4bb4-b8d8-5b260591d745",
                                Name: codataccounting.String("Earl Rosenbaum II"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9c9c3f56-7e0e-4252-b65b-1d62fcdace1f",
                                Name: codataccounting.String("Mrs. Ashley Connelly"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e2239e8f-25cd-40d1-9d95-9f439e39266c",
                                Name: codataccounting.String("Lionel Mraz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("molestiae"),
                            ID: "aa2b2411-3695-4d1e-a698-fcc4596217c2",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "76763342-5403-48bf-b597-1e9819055738",
                            Name: codataccounting.String("Benny Ward"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c7fda395-94d6-46bc-aae4-80632b9954b6",
                            Name: codataccounting.String("Mr. Alfonso Dibbert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "36982855-3cb1-4000-abef-4921ec2053b7",
                            Name: codataccounting.String("Claire Fay"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ac8ee0f2-bf19-4588-940d-03f3deba297b",
                            Name: codataccounting.String("Nathan Watsica MD"),
                        },
                    },
                    UnitAmount: 7524.92,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("numquam"),
            Note: codataccounting.String("consequatur"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("doloribus"),
                        Currency: codataccounting.String("quos"),
                        CurrencyRate: codataccounting.Float64(4169.67),
                        TotalAmount: codataccounting.Float64(5023.8),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fd52405c-b331-4d49-af4f-127fb0e0bf1f"),
                            Name: codataccounting.String("Steve Block"),
                        },
                        Currency: codataccounting.String("dignissimos"),
                        CurrencyRate: codataccounting.Float64(5283.5),
                        ID: codataccounting.String("d0acca77-aeb7-4b70-a1a5-2046b64e99fb"),
                        Note: codataccounting.String("doloremque"),
                        PaidOnDate: codataccounting.String("officiis"),
                        Reference: codataccounting.String("nisi"),
                        TotalAmount: codataccounting.Float64(4421.29),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("necessitatibus"),
                        Currency: codataccounting.String("alias"),
                        CurrencyRate: codataccounting.Float64(5898.68),
                        TotalAmount: codataccounting.Float64(2853.36),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fdfed554-0ef5-43a3-8a1b-8fe99731adc0"),
                            Name: codataccounting.String("Hope Macejkovic"),
                        },
                        Currency: codataccounting.String("eveniet"),
                        CurrencyRate: codataccounting.Float64(1603.53),
                        ID: codataccounting.String("dfb70fb3-8742-490d-b365-61eca16ef894"),
                        Note: codataccounting.String("ad"),
                        PaidOnDate: codataccounting.String("ab"),
                        Reference: codataccounting.String("harum"),
                        TotalAmount: codataccounting.Float64(8177.91),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ducimus"),
                        Currency: codataccounting.String("nisi"),
                        CurrencyRate: codataccounting.Float64(8811.68),
                        TotalAmount: codataccounting.Float64(8848.3),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("eb518c4d-a1fa-4d35-912f-06d4e5b72f0f"),
                            Name: codataccounting.String("Valerie Little"),
                        },
                        Currency: codataccounting.String("laudantium"),
                        CurrencyRate: codataccounting.Float64(6675.59),
                        ID: codataccounting.String("0424e00a-1d6e-4b94-b464-5d03084fbba5"),
                        Note: codataccounting.String("cumque"),
                        PaidOnDate: codataccounting.String("cumque"),
                        Reference: codataccounting.String("vero"),
                        TotalAmount: codataccounting.Float64(9539.16),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("tenetur"),
                        Currency: codataccounting.String("ipsam"),
                        CurrencyRate: codataccounting.Float64(7981.22),
                        TotalAmount: codataccounting.Float64(7080.11),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("01fe51e5-28a4-45ac-82b8-5f8bc2caba8d"),
                            Name: codataccounting.String("Clifford Bogisich"),
                        },
                        Currency: codataccounting.String("illum"),
                        CurrencyRate: codataccounting.Float64(8601.44),
                        ID: codataccounting.String("597ff471-1aa1-4bc7-8b86-cecc74f77b48"),
                        Note: codataccounting.String("aliquam"),
                        PaidOnDate: codataccounting.String("totam"),
                        Reference: codataccounting.String("soluta"),
                        TotalAmount: codataccounting.Float64(8425.39),
                    },
                },
            },
            RemainingCredit: 4371.07,
            SourceModifiedDate: codataccounting.String("mollitia"),
            Status: shared.CreditNoteStatusSubmitted,
            SubTotal: 9404.72,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempora": map[string]interface{}{
                        "architecto": "nulla",
                        "qui": "maxime",
                    },
                },
            },
            TotalAmount: 1898.27,
            TotalDiscount: 7465.11,
            TotalTaxAmount: 5541.99,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 5119.21,
                    Name: "Erika Hahn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(239283),
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
        CreditNoteID: "eveniet",
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
        Query: codataccounting.String("ipsa"),
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
            AdditionalTaxAmount: codataccounting.Float64(3846.3),
            AdditionalTaxPercentage: codataccounting.Float64(253.21),
            AllocatedOnDate: codataccounting.String("labore"),
            CreditNoteNumber: codataccounting.String("ullam"),
            Currency: codataccounting.String("excepturi"),
            CurrencyRate: codataccounting.Float64(7463.15),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("voluptates"),
                ID: "bbad02f2-586b-4cf1-9255-8daa95be6cd0",
            },
            DiscountPercentage: 1829.75,
            ID: codataccounting.String("756c354a-a432-4b47-a176-3c5208c23e98"),
            IssueDate: codataccounting.String("sit"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d82f0d45-eb4a-48b6-b4ee-5cfc18edc7f7"),
                        Name: codataccounting.String("Julio Wehner"),
                    },
                    Description: codataccounting.String("recusandae"),
                    DiscountAmount: codataccounting.Float64(298.97),
                    DiscountPercentage: codataccounting.Float64(2587.5),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b3d3ed0c-5670-4ef4-abd3-c9f1cc503f6c",
                        Name: codataccounting.String("Sheryl Reichel"),
                    },
                    Quantity: 30.26,
                    SubTotal: codataccounting.Float64(6464.39),
                    TaxAmount: codataccounting.Float64(4283.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1860.92),
                        ID: codataccounting.String("90f957f3-8518-49ad-bef8-07aae03f33ca"),
                        Name: codataccounting.String("Lindsay Zemlak"),
                    },
                    TotalAmount: codataccounting.Float64(8535.32),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "4032ba26-fd36-48ba-9216-bcb415835c73",
                                Name: codataccounting.String("Anita Braun"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3133edc0-46bc-4516-bbbc-a49227c42c22",
                                Name: codataccounting.String("Warren Hills"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0495c5db-b3c5-47c1-a498-1e8aa257ddc1",
                                Name: codataccounting.String("Dennis Considine"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "de64bfcc-5469-4d40-95df-a796206bef2b",
                                Name: codataccounting.String("Lynda Dicki"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quia"),
                            ID: "c1aa010e-9aac-42e9-9355-86d18f9f97a4",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "ad2bf7d6-7ca8-44ad-99b4-1d6124353187",
                            Name: codataccounting.String("Erma Windler"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "03ad421b-d43d-41f0-8b0a-0003eb22d9b3",
                            Name: codataccounting.String("Harvey Beahan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4faa741c-57d1-4fed-8205-0d38dc3ce185",
                            Name: codataccounting.String("Marlene Cassin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ee69166a-8be3-4444-aac8-b3a2875c6c1f",
                            Name: codataccounting.String("Sam Aufderhar"),
                        },
                    },
                    UnitAmount: 18.19,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("voluptate"),
            Note: codataccounting.String("repellendus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("deserunt"),
                        Currency: codataccounting.String("error"),
                        CurrencyRate: codataccounting.Float64(7699.22),
                        TotalAmount: codataccounting.Float64(5131.41),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7ae50c16-661a-41d9-936a-7e8d53213f3f"),
                            Name: codataccounting.String("Melanie Lueilwitz"),
                        },
                        Currency: codataccounting.String("fugit"),
                        CurrencyRate: codataccounting.Float64(8256.43),
                        ID: codataccounting.String("b764c59f-0a56-4ceb-8ada-29ca79181c95"),
                        Note: codataccounting.String("laboriosam"),
                        PaidOnDate: codataccounting.String("molestiae"),
                        Reference: codataccounting.String("ab"),
                        TotalAmount: codataccounting.Float64(4014.49),
                    },
                },
            },
            RemainingCredit: 4363.67,
            SourceModifiedDate: codataccounting.String("dolorem"),
            Status: shared.CreditNoteStatusVoid,
            SubTotal: 3218.89,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ipsa": map[string]interface{}{
                        "minima": "vel",
                        "nisi": "minima",
                        "et": "autem",
                    },
                },
            },
            TotalAmount: 2204.55,
            TotalDiscount: 6359.03,
            TotalTaxAmount: 2495.41,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 2274.7,
                    Name: "Dustin Blanda",
                },
                shared.WithholdingTaxitems{
                    Amount: 7315.02,
                    Name: "Miss Bernice Cummings",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "asperiores",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(179648),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```
