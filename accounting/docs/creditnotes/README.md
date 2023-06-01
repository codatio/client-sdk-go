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
            AdditionalTaxAmount: codataccounting.Float64(7609.42),
            AdditionalTaxPercentage: codataccounting.Float64(2365.88),
            AllocatedOnDate: codataccounting.String("ab"),
            CreditNoteNumber: codataccounting.String("maxime"),
            Currency: codataccounting.String("porro"),
            CurrencyRate: codataccounting.Float64(1279.53),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("reiciendis"),
                ID: "1fcb51c9-a41f-4fbe-9cbd-795ee65e076c",
            },
            DiscountPercentage: 7903.41,
            ID: codataccounting.String("7abf616e-a5c7-4164-9934-b90f2e09d19d"),
            IssueDate: codataccounting.String("magni"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c2f9e2e1-0594-44b9-b5d2-37a72f90849d"),
                        Name: codataccounting.String("Madeline Waters"),
                    },
                    Description: codataccounting.String("id"),
                    DiscountAmount: codataccounting.Float64(9081.27),
                    DiscountPercentage: codataccounting.Float64(7595.37),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b7537cd9-222c-49ff-9749-1aabfa2e761f",
                        Name: codataccounting.String("Angelica O'Reilly"),
                    },
                    Quantity: 2893.42,
                    SubTotal: codataccounting.Float64(3375.7),
                    TaxAmount: codataccounting.Float64(3985.78),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8757.66),
                        ID: codataccounting.String("f1031e68-99f0-4c20-81e2-2cd55cc0584a"),
                        Name: codataccounting.String("Lena Goyette"),
                    },
                    TotalAmount: codataccounting.Float64(3758.77),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "971fc820-c65b-4037-bb8e-0cc885187e4d",
                                Name: codataccounting.String("William Graham"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "28c5dddb-46aa-41cf-96d8-28da01319112",
                                Name: codataccounting.String("Reginald Hackett"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "45c1d81f-2904-42f5-a9b7-aff0ea2216cb",
                                Name: codataccounting.String("Miss Larry Kunde"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "163e279a-3b08-44da-9925-7d04f40847a7",
                                Name: codataccounting.String("Lori Sipes"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eius"),
                            ID: "96cbdeec-f6b9-49bc-a356-2ebfdf55c294",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "60b06a12-8776-44ee-b6d0-c6d6ed9c73dd",
                            Name: codataccounting.String("Sheila Grant"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "509a8e87-0d3c-45a1-b9c2-42c7b66a1f30",
                            Name: codataccounting.String("Cody Franey"),
                        },
                    },
                    UnitAmount: 3480.56,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b6719890-f42a-44bb-838d-85b260591d74"),
                        Name: codataccounting.String("Silvia Dietrich"),
                    },
                    Description: codataccounting.String("quae"),
                    DiscountAmount: codataccounting.Float64(3298.36),
                    DiscountPercentage: codataccounting.Float64(5726.33),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c9c3f567-e0e2-4527-a5b1-d62fcdace1f0",
                        Name: codataccounting.String("Christina Bode"),
                    },
                    Quantity: 9210.86,
                    SubTotal: codataccounting.Float64(1621.71),
                    TaxAmount: codataccounting.Float64(1333.6),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1945.74),
                        ID: codataccounting.String("9e8f25cd-0d19-4d95-9f43-9e39266cbd95"),
                        Name: codataccounting.String("Clinton Oberbrunner"),
                    },
                    TotalAmount: codataccounting.Float64(6980.88),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "4113695d-1e66-498f-8c45-96217c297767",
                                Name: codataccounting.String("Edith Frami"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("ad"),
                            ID: "4038bfb5-971e-4981-9055-7389cedbac7f",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "39594d66-bc2a-4e48-8632-b9954b6fa220",
                            Name: codataccounting.String("Crystal Johnson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8553cb10-006b-4ef4-921e-c2053b749366",
                            Name: codataccounting.String("Wilson Ledner"),
                        },
                    },
                    UnitAmount: 312.92,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f2bf1958-8d40-4d03-b3de-ba297be3e90b"),
                        Name: codataccounting.String("Jesus Abshire"),
                    },
                    Description: codataccounting.String("quos"),
                    DiscountAmount: codataccounting.Float64(4169.67),
                    DiscountPercentage: codataccounting.Float64(5023.8),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "fd52405c-b331-4d49-af4f-127fb0e0bf1f",
                        Name: codataccounting.String("Steve Block"),
                    },
                    Quantity: 4940.93,
                    SubTotal: codataccounting.Float64(5283.5),
                    TaxAmount: codataccounting.Float64(8602.82),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(65.17),
                        ID: codataccounting.String("acca77ae-b7b7-4021-a520-46b64e99fb0e"),
                        Name: codataccounting.String("Ms. Delores Tromp"),
                    },
                    TotalAmount: codataccounting.Float64(9417.76),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "fed5540e-f53a-434a-9b8f-e99731adc05d",
                                Name: codataccounting.String("Alvin Parker"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "dfb70fb3-8742-490d-b365-61eca16ef894",
                                Name: codataccounting.String("Kathleen Pollich"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "6eeeb518-c4da-41fa-9355-12f06d4e5b72",
                                Name: codataccounting.String("James Yost"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "8568a042-4e00-4a1d-aeb9-434645d03084",
                                Name: codataccounting.String("Malcolm Rempel"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("cumque"),
                            ID: "ceff5cb0-1fe5-41e5-a8a4-5ac82b85f8bc",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "aba8da41-27dd-4597-bf47-11aa1bc74b86",
                            Name: codataccounting.String("Alonzo Schaefer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f77b4848-bd6a-46f0-841d-2c3b80809437",
                            Name: codataccounting.String("Laverne Bednar II"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "59bebbad-02f2-4586-bcf1-52558daa95be",
                            Name: codataccounting.String("Mr. Robyn Stoltenberg"),
                        },
                    },
                    UnitAmount: 3159.04,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6c354aa4-32b4-47e1-b63c-5208c23e9802"),
                        Name: codataccounting.String("Willard Dibbert PhD"),
                    },
                    Description: codataccounting.String("magnam"),
                    DiscountAmount: codataccounting.Float64(3401.07),
                    DiscountPercentage: codataccounting.Float64(9204.88),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "b4a8b674-ee5c-4fc1-8edc-7f787e32e04b",
                        Name: codataccounting.String("Freda Frami"),
                    },
                    Quantity: 614.71,
                    SubTotal: codataccounting.Float64(7685.46),
                    TaxAmount: codataccounting.Float64(3532.32),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4271.07),
                        ID: codataccounting.String("70ef42bd-3c9f-41cc-903f-6c39bcd0a629"),
                        Name: codataccounting.String("Tami McClure"),
                    },
                    TotalAmount: codataccounting.Float64(9704.91),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "85189ad7-ef80-47aa-a03f-33ca79fb9de4",
                                Name: codataccounting.String("Rita Crist"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("qui"),
                            ID: "6fd368ba-9216-4bcb-8158-35c736417231",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "edc046bc-5163-4bbc-a492-27c42c22c553",
                            Name: codataccounting.String("Donna Gottlieb"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5dbb3c57-c1e4-4981-a8aa-257ddc1912eb",
                            Name: codataccounting.String("Rogelio Keebler"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fcc5469d-4015-4dfa-b962-06bef2b0a3e4",
                            Name: codataccounting.String("Robyn Braun"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "010e9aac-2e91-4355-86d1-8f9f97a4bfad",
                            Name: codataccounting.String("Juana Zboncak"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "67ca84ad-99b4-41d6-9243-531870cf68b0",
                            Name: codataccounting.String("Kristine Shields"),
                        },
                    },
                    UnitAmount: 1067.67,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("rerum"),
            Note: codataccounting.String("repellendus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nesciunt"),
                        Currency: codataccounting.String("facere"),
                        CurrencyRate: codataccounting.Float64(1047.54),
                        TotalAmount: codataccounting.Float64(9646.41),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0cb0a000-3eb2-42d9-b3a7-0d94faa741c5"),
                            Name: codataccounting.String("Mona Carroll"),
                        },
                        Currency: codataccounting.String("quibusdam"),
                        CurrencyRate: codataccounting.Float64(8045.25),
                        ID: codataccounting.String("2050d38d-c3ce-4185-872f-9ee69166a8be"),
                        Note: codataccounting.String("adipisci"),
                        PaidOnDate: codataccounting.String("dolore"),
                        Reference: codataccounting.String("tempora"),
                        TotalAmount: codataccounting.Float64(3088.64),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("debitis"),
                        Currency: codataccounting.String("similique"),
                        CurrencyRate: codataccounting.Float64(7807.03),
                        TotalAmount: codataccounting.Float64(5051.92),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b3a2875c-6c1f-4e60-ad07-d2a9c87ae50c"),
                            Name: codataccounting.String("Jessie Hirthe V"),
                        },
                        Currency: codataccounting.String("veritatis"),
                        CurrencyRate: codataccounting.Float64(8360.84),
                        ID: codataccounting.String("9136a7e8-d532-413f-bf65-8752db764c59"),
                        Note: codataccounting.String("asperiores"),
                        PaidOnDate: codataccounting.String("doloremque"),
                        Reference: codataccounting.String("id"),
                        TotalAmount: codataccounting.Float64(3349.54),
                    },
                },
            },
            RemainingCredit: 4096.77,
            SourceModifiedDate: codataccounting.String("placeat"),
            Status: shared.CreditNoteStatusPartiallyPaid,
            SubTotal: 6881.14,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "culpa": map[string]interface{}{
                        "laborum": "consequuntur",
                        "omnis": "maxime",
                        "officia": "iusto",
                        "natus": "ab",
                    },
                    "deleniti": map[string]interface{}{
                        "eligendi": "sint",
                    },
                    "ipsam": map[string]interface{}{
                        "molestiae": "ab",
                        "ex": "iure",
                    },
                    "dolorem": map[string]interface{}{
                        "ad": "ipsum",
                        "ipsa": "nam",
                        "minima": "vel",
                        "nisi": "minima",
                    },
                },
            },
            TotalAmount: 933.84,
            TotalDiscount: 4177.55,
            TotalTaxAmount: 2204.55,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 2495.41,
                    Name: "Grace Lesch Sr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 6639.81,
                    Name: "Billy Heathcote MD",
                },
                shared.WithholdingTaxitems{
                    Amount: 5936,
                    Name: "Ms. Craig Turner",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(302814),
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
        CreditNoteID: "autem",
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
        Query: codataccounting.String("reprehenderit"),
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
            AdditionalTaxAmount: codataccounting.Float64(6935.09),
            AdditionalTaxPercentage: codataccounting.Float64(5597.16),
            AllocatedOnDate: codataccounting.String("officia"),
            CreditNoteNumber: codataccounting.String("modi"),
            Currency: codataccounting.String("alias"),
            CurrencyRate: codataccounting.Float64(7216.91),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("minus"),
                ID: "05fab0d6-50ed-4f22-a94d-20ec90ea41d1",
            },
            DiscountPercentage: 9416.83,
            ID: codataccounting.String("465e8515-6fff-473f-9f54-fdd5ea954339"),
            IssueDate: codataccounting.String("voluptatum"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("afb42a8d-6338-48e4-9803-9ea5f9b18a24"),
                        Name: codataccounting.String("Jeannie Schuppe V"),
                    },
                    Description: codataccounting.String("aut"),
                    DiscountAmount: codataccounting.Float64(2158.13),
                    DiscountPercentage: codataccounting.Float64(6113.61),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "dacd38ed-0dc6-471d-87f1-e3af15920c90",
                        Name: codataccounting.String("Stephen Robel"),
                    },
                    Quantity: 474.01,
                    SubTotal: codataccounting.Float64(1205.48),
                    TaxAmount: codataccounting.Float64(9957.13),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1728.07),
                        ID: codataccounting.String("bd89c8a3-2639-4da5-b7b6-902b881a94f6"),
                        Name: codataccounting.String("Annie Huel"),
                    },
                    TotalAmount: codataccounting.Float64(6534.92),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f0af8c69-1d73-42d9-bbaf-9476a2ae8dcc",
                                Name: codataccounting.String("Angela Schroeder"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3512c737-8489-4307-90a0-0e966ec736d4",
                                Name: codataccounting.String("Stephanie Moen"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "98c783c9-2398-4ed3-93ab-7ca3c5ca8649",
                                Name: codataccounting.String("Kurt Auer"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("quibusdam"),
                            ID: "5d6989b7-2064-4510-b7d1-9ea83d492ed1",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Project,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8a2c1954-545e-4955-9cc1-85ea4901c7c4",
                            Name: codataccounting.String("Harriet Schultz"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a784aba3-d230-4edf-b381-1a115382bd7e",
                            Name: codataccounting.String("Jon Hodkiewicz IV"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "621c58f4-d739-4656-8c20-a0711a961d24",
                            Name: codataccounting.String("Lance Stokes"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8f532d89-2cf7-4812-8b51-2c878240bf54",
                            Name: codataccounting.String("Randal Lockman"),
                        },
                    },
                    UnitAmount: 5030.15,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f1bf0bc8-e1f2-406d-9d83-1d0081090f67"),
                        Name: codataccounting.String("Loretta Howe"),
                    },
                    Description: codataccounting.String("doloribus"),
                    DiscountAmount: codataccounting.Float64(1877.7),
                    DiscountPercentage: codataccounting.Float64(6638.67),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "681c5768-dce7-4424-89a2-15e08601489a",
                        Name: codataccounting.String("Ora Homenick"),
                    },
                    Quantity: 2072.02,
                    SubTotal: codataccounting.Float64(6489.85),
                    TaxAmount: codataccounting.Float64(9474.2),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2479.72),
                        ID: codataccounting.String("dd9dda33-dcd6-4348-be4a-7a98e4df37e4"),
                        Name: codataccounting.String("Pam Leannon"),
                    },
                    TotalAmount: codataccounting.Float64(3422.41),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "413e13a4-8231-4090-bbd3-54c092bd734f",
                                Name: codataccounting.String("Louise Funk"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "d86f4bb2-0fe5-4d91-9cbf-e749caf45a27",
                                Name: codataccounting.String("Ruben Muller"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c9e6d10e-9db3-4ad4-86b0-3108d9c33747",
                                Name: codataccounting.String("Jennifer Ledner"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "94f2ab1f-d567-41e9-8326-350a46714378",
                                Name: codataccounting.String("Dr. Devin Waters"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("error"),
                            ID: "1594d93a-74c0-4252-be3b-4b4db8b778eb",
                        },
                        IsBilledTo: shared.BilledToType1Project,
                        IsRebilledTo: shared.BilledToType1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "e1d2cf50-2baf-4b2c-bc46-35d5e65da028",
                            Name: codataccounting.String("Curtis Thompson"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a1e30fda-9664-489d-bb78-673e13a12a6b",
                            Name: codataccounting.String("Marshall Corkery"),
                        },
                    },
                    UnitAmount: 2959.53,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("594487f5-c843-4836-b86b-3cdf6415b044"),
                        Name: codataccounting.String("Terrell Muller"),
                    },
                    Description: codataccounting.String("dicta"),
                    DiscountAmount: codataccounting.Float64(1911.01),
                    DiscountPercentage: codataccounting.Float64(9646.4),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "4eedbe78-bf60-4682-9894-ea763d5c7279",
                        Name: codataccounting.String("Lula Koepp"),
                    },
                    Quantity: 1139.47,
                    SubTotal: codataccounting.Float64(3005.84),
                    TaxAmount: codataccounting.Float64(5032.47),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8612.36),
                        ID: codataccounting.String("6d549e56-35b3-43bc-8f97-0c42fc9f4844"),
                        Name: codataccounting.String("Andrea Harris"),
                    },
                    TotalAmount: codataccounting.Float64(3245.72),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "796065c0-efa6-4f93-b90a-1b8c95be1254",
                                Name: codataccounting.String("Jessie Fisher"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "4fe77210-d1f6-4558-899c-722d2bc0f940",
                                Name: codataccounting.String("Kurt Schuppe"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "aae042dd-7caa-4c9b-8caa-1cfe9e15df90",
                                Name: codataccounting.String("Faye Bayer"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("amet"),
                            ID: "7831983d-42e5-44a8-9466-597c50233c14",
                        },
                        IsBilledTo: shared.BilledToType1NotApplicable,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "d51aaa6d-df5a-4bd6-887c-5fc2b862a00b",
                            Name: codataccounting.String("Cary Jones"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "00157630-bda7-4afd-ad84-a35a41238e1a",
                            Name: codataccounting.String("Emily Hand"),
                        },
                    },
                    UnitAmount: 1866.88,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6ae33bef-971a-48f4-abca-1106fe965b71"),
                        Name: codataccounting.String("Meredith Barrows"),
                    },
                    Description: codataccounting.String("tenetur"),
                    DiscountAmount: codataccounting.Float64(5358.76),
                    DiscountPercentage: codataccounting.Float64(5373.21),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "ec9f7b99-a550-4a65-aed3-33bb0ce8aa65",
                        Name: codataccounting.String("Victoria Denesik"),
                    },
                    Quantity: 5250.89,
                    SubTotal: codataccounting.Float64(4102.39),
                    TaxAmount: codataccounting.Float64(8917.7),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7430.23),
                        ID: codataccounting.String("7e14ca56-4089-4150-8970-19a48f88ece7"),
                        Name: codataccounting.String("Mrs. Terrell Mayert"),
                    },
                    TotalAmount: codataccounting.Float64(81.62),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "105d3890-8162-4c6b-ab68-a0f657b7d03a",
                                Name: codataccounting.String("Dr. Carrie Lang"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("at"),
                            ID: "e30f069d-8106-418d-97e1-52297510da80",
                        },
                        IsBilledTo: shared.BilledToType1Unknown,
                        IsRebilledTo: shared.BilledToType1Unknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "2292cc61-c2a7-402b-b97e-e102da2de35f",
                            Name: codataccounting.String("Miss Frankie Bailey"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3eaab454-02ac-4170-8bf1-cc9fc61aae5e",
                            Name: codataccounting.String("Miss Jon Willms"),
                        },
                    },
                    UnitAmount: 5655.59,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("qui"),
            Note: codataccounting.String("soluta"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nihil"),
                        Currency: codataccounting.String("ut"),
                        CurrencyRate: codataccounting.Float64(2799.45),
                        TotalAmount: codataccounting.Float64(8430.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("08a2267a-aee7-49e3-871a-d31becb83d23"),
                            Name: codataccounting.String("Cassandra O'Conner"),
                        },
                        Currency: codataccounting.String("facilis"),
                        CurrencyRate: codataccounting.Float64(9571.11),
                        ID: codataccounting.String("c23d9450-a986-4a49-9bac-707f06b28ecc"),
                        Note: codataccounting.String("atque"),
                        PaidOnDate: codataccounting.String("autem"),
                        Reference: codataccounting.String("eius"),
                        TotalAmount: codataccounting.Float64(5999.15),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sunt"),
                        Currency: codataccounting.String("amet"),
                        CurrencyRate: codataccounting.Float64(5118.79),
                        TotalAmount: codataccounting.Float64(4131.3),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("f62c969c-4cc6-4b78-890a-3fd3c81da10f"),
                            Name: codataccounting.String("Earnest Crooks"),
                        },
                        Currency: codataccounting.String("voluptatibus"),
                        CurrencyRate: codataccounting.Float64(6148.15),
                        ID: codataccounting.String("31da3edb-51fa-4d94-acc9-435137726d15"),
                        Note: codataccounting.String("velit"),
                        PaidOnDate: codataccounting.String("quia"),
                        Reference: codataccounting.String("dicta"),
                        TotalAmount: codataccounting.Float64(7214.48),
                    },
                },
            },
            RemainingCredit: 5545.08,
            SourceModifiedDate: codataccounting.String("velit"),
            Status: shared.CreditNoteStatusDraft,
            SubTotal: 6851.1,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aliquid": map[string]interface{}{
                        "laboriosam": "sint",
                        "architecto": "totam",
                        "alias": "hic",
                        "tenetur": "iure",
                    },
                    "consequatur": map[string]interface{}{
                        "cum": "occaecati",
                        "fuga": "ex",
                        "autem": "nostrum",
                        "atque": "saepe",
                    },
                },
            },
            TotalAmount: 4320.55,
            TotalDiscount: 5652.45,
            TotalTaxAmount: 6841.96,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 7212.15,
                    Name: "Alexander Friesen",
                },
                shared.WithholdingTaxitems{
                    Amount: 5197.99,
                    Name: "Rosemarie Pollich",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CreditNoteID: "ducimus",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(323612),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```
