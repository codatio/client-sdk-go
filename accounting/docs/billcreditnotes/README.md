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
            AllocatedOnDate: codataccounting.String("omnis"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("molestiae"),
            CurrencyRate: codataccounting.Float64(191.93),
            DiscountPercentage: 4701.32,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("magnam"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a4469b6e-2141-4959-890a-fa563e2516fe"),
                        Name: codataccounting.String("Jasmine Lind"),
                    },
                    Description: codataccounting.String("architecto"),
                    DiscountAmount: codataccounting.Float64(995.69),
                    DiscountPercentage: codataccounting.Float64(9194.83),
                    ItemRef: &shared.ItemRef{
                        ID: "5b7fd2ed-0289-421c-9dc6-92601fb576b0",
                        Name: codataccounting.String("Dr. Herman Wolf"),
                    },
                    Quantity: 117.14,
                    SubTotal: codataccounting.Float64(7649.12),
                    TaxAmount: codataccounting.Float64(3599.78),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9441.24),
                        ID: codataccounting.String("bb258705-3202-4c73-95fe-9b90c28909b3"),
                        Name: codataccounting.String("Merle Gleichner"),
                    },
                    TotalAmount: codataccounting.Float64(5356.33),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9cbf4863-3323-4f9b-b7f3-a4100674ebf6",
                                Name: codataccounting.String("Dr. Craig Littel DDS"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a77a89eb-f737-4ae4-a03c-e5e6a95d8a0d",
                                Name: codataccounting.String("Rhonda Kautzer"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "2af7a73c-f3be-4453-b870-b326b5a73429",
                                Name: codataccounting.String("Miss Jody Rogahn"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "422bb679-d232-4271-9bf0-cbb1e31b8b90",
                                Name: codataccounting.String("Mike Greenholt"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("dolorum"),
                            ID: "1108e0ad-cf4b-4921-879f-ce953f73ef7f",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7abd74dd-39c0-4f5d-acff-7c70a45626d4",
                            Name: codataccounting.String("Mrs. Vicki Langosh"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6d9f5fce-6c55-4614-ac3e-250fb008c42e",
                            Name: codataccounting.String("Ellen Borer"),
                        },
                    },
                    UnitAmount: 8104.24,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("366c8dd6-b144-4290-b474-778a7bd466d2"),
                        Name: codataccounting.String("Miss Devin Bogan"),
                    },
                    Description: codataccounting.String("neque"),
                    DiscountAmount: codataccounting.Float64(7786.96),
                    DiscountPercentage: codataccounting.Float64(8472.76),
                    ItemRef: &shared.ItemRef{
                        ID: "ca425190-4e52-43c7-a0bc-7178e4796f2a",
                        Name: codataccounting.String("Carol Sawayn"),
                    },
                    Quantity: 5100.17,
                    SubTotal: codataccounting.Float64(1598.67),
                    TaxAmount: codataccounting.Float64(5361.78),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1438.29),
                        ID: codataccounting.String("aa482562-f222-4e98-97ee-17cbe61e6b7b"),
                        Name: codataccounting.String("Warren Rau V"),
                    },
                    TotalAmount: codataccounting.Float64(7313.98),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "c20c4f37-89fd-4871-b99d-d2efd121aa6f",
                                Name: codataccounting.String("Lila Kassulke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("libero"),
                            ID: "db04f157-5608-42d6-8ea1-9f1d17051339",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "8086a184-0394-4c26-871f-93f5f0642dac",
                            Name: codataccounting.String("Blanche Yundt II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c413aa63-aae8-4d67-864d-bb675fd5e60b",
                            Name: codataccounting.String("Arlene Heidenreich"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4f6fbee4-1f33-4317-be35-b60eb1ea4265",
                            Name: codataccounting.String("Cathy Rohan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c28744ed-53b8-48f3-a8d8-f5c0b2f2fb7b",
                            Name: codataccounting.String("Margarita Greenholt"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "76b26916-fe1f-408f-8294-e3698f447f60",
                            Name: codataccounting.String("Cecelia Lakin"),
                        },
                    },
                    UnitAmount: 2777.73,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("5e80ca55-efd2-40e4-97e1-858b6a89fbe3"),
                        Name: codataccounting.String("Wesley Nikolaus"),
                    },
                    Description: codataccounting.String("accusamus"),
                    DiscountAmount: codataccounting.Float64(2726.83),
                    DiscountPercentage: codataccounting.Float64(5436.78),
                    ItemRef: &shared.ItemRef{
                        ID: "24d0ab40-7508-48e5-9862-065e904f3b11",
                        Name: codataccounting.String("Francisco Powlowski"),
                    },
                    Quantity: 7241.48,
                    SubTotal: codataccounting.Float64(9488.61),
                    TaxAmount: codataccounting.Float64(3888.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(27.03),
                        ID: codataccounting.String("3a79f9df-e0ab-47da-8a50-ce187f86bc17"),
                        Name: codataccounting.String("Angelina Jenkins"),
                    },
                    TotalAmount: codataccounting.Float64(8872.65),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "e9526f8d-986e-4881-aad4-f0e1012563f9",
                                Name: codataccounting.String("Tricia Cronin"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "973e922a-57a1-45be-be06-0807e2b6e3ab",
                                Name: codataccounting.String("Jordan Haag"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "0597a60f-f2a5-44a3-9e94-764a3e865e79",
                                Name: codataccounting.String("Natalie Witting"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "51a5a9da-660f-4f57-bfaa-d4f9efc1b451",
                                Name: codataccounting.String("Mrs. Erma Berge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("eum"),
                            ID: "48dc2f61-5199-4ebf-90e9-fe6c632ca3ae",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "11799631-2fde-4047-b177-8ff61d017476",
                            Name: codataccounting.String("Jeanne Beer II"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b6a66065-9a1a-4dea-ab58-51d6c645b08b",
                            Name: codataccounting.String("Doris Lemke MD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "aa0fe1ad-e008-4e6f-8c5f-350d8cdb5a34",
                            Name: codataccounting.String("Cassandra Bogan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "01042181-3d52-408e-8e7e-253b668451c6",
                            Name: codataccounting.String("Brent Walter II"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e16deab3-fec9-4578-a645-84273a8418d1",
                            Name: codataccounting.String("Ms. Marilyn Feest"),
                        },
                    },
                    UnitAmount: 7468.37,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("alias"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("eos"),
                        Currency: codataccounting.String("occaecati"),
                        CurrencyRate: codataccounting.Float64(6128.67),
                        TotalAmount: codataccounting.Float64(1700.99),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1aefb9f5-8c4d-486e-a8e4-be056013f59d"),
                            Name: codataccounting.String("Claude Hickle"),
                        },
                        Currency: codataccounting.String("quis"),
                        CurrencyRate: codataccounting.Float64(5718.44),
                        ID: codataccounting.String("ecfef66e-f1ca-4a33-83c2-beb477373c8d"),
                        Note: codataccounting.String("iure"),
                        PaidOnDate: codataccounting.String("odit"),
                        Reference: codataccounting.String("voluptatibus"),
                        TotalAmount: codataccounting.Float64(4269.04),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("magnam"),
                        Currency: codataccounting.String("quibusdam"),
                        CurrencyRate: codataccounting.Float64(789.69),
                        TotalAmount: codataccounting.Float64(8180.34),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b1f2c431-0661-4e96-b49e-1cf9e06e3a43"),
                            Name: codataccounting.String("Miss Karen Batz"),
                        },
                        Currency: codataccounting.String("ea"),
                        CurrencyRate: codataccounting.Float64(6931.53),
                        ID: codataccounting.String("6bc9b8f7-59ea-4c55-a974-1d311352965b"),
                        Note: codataccounting.String("libero"),
                        PaidOnDate: codataccounting.String("rem"),
                        Reference: codataccounting.String("dolorum"),
                        TotalAmount: codataccounting.Float64(4876.76),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("fugit"),
                        Currency: codataccounting.String("alias"),
                        CurrencyRate: codataccounting.Float64(1680.42),
                        TotalAmount: codataccounting.Float64(4254.02),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("11435e13-9dbc-4225-9b1a-bda8c070e108"),
                            Name: codataccounting.String("Ms. Lynne Rau"),
                        },
                        Currency: codataccounting.String("dolores"),
                        CurrencyRate: codataccounting.Float64(8247.98),
                        ID: codataccounting.String("1ad879ee-b966-45b8-9efb-d02bae0be2d7"),
                        Note: codataccounting.String("praesentium"),
                        PaidOnDate: codataccounting.String("odit"),
                        Reference: codataccounting.String("explicabo"),
                        TotalAmount: codataccounting.Float64(3589.95),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(6214.73),
            SourceModifiedDate: codataccounting.String("earum"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "recusandae": map[string]interface{}{
                        "ut": "quidem",
                        "quis": "beatae",
                        "unde": "molestiae",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "f92443da-7ce5-42b8-95c5-37c6454efb0b",
                SupplierName: codataccounting.String("ratione"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 2899.13,
            TotalTaxAmount: 5208.75,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 3759.94,
                    Name: "Jacob Schinner",
                },
                shared.WithholdingTaxitems{
                    Amount: 6692.37,
                    Name: "Brendan Rippin",
                },
                shared.WithholdingTaxitems{
                    Amount: 9974.37,
                    Name: "Ms. Corey Kiehn",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(448369),
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
        BillCreditNoteID: "ducimus",
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
        Query: codataccounting.String("excepturi"),
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
            AllocatedOnDate: codataccounting.String("dolores"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("error"),
            CurrencyRate: codataccounting.Float64(850.76),
            DiscountPercentage: 4981.8,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("voluptate"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("eac646ec-b573-4409-a3eb-1e5a2b12eb07"),
                        Name: codataccounting.String("Harold Boyer"),
                    },
                    Description: codataccounting.String("libero"),
                    DiscountAmount: codataccounting.Float64(5665.06),
                    DiscountPercentage: codataccounting.Float64(5782.1),
                    ItemRef: &shared.ItemRef{
                        ID: "545fc95f-a889-470e-989d-bb30fcb33ea0",
                        Name: codataccounting.String("Ms. Sally Rempel"),
                    },
                    Quantity: 7566.54,
                    SubTotal: codataccounting.Float64(8200.23),
                    TaxAmount: codataccounting.Float64(2514.64),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2981.87),
                        ID: codataccounting.String("e2f52d82-d351-43bb-af48-b656bcdb35ff"),
                        Name: codataccounting.String("Raquel Greenfelder"),
                    },
                    TotalAmount: codataccounting.Float64(4407.77),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "37a8cd9e-7319-4c17-bd52-5f77b114eeb5",
                                Name: codataccounting.String("Johanna Weimann"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "5fc37814-d4c9-48e0-82bb-89eb75dad636",
                                Name: codataccounting.String("Mrs. Leslie Anderson I"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("illum"),
                            ID: "8bb31180-f739-4ae9-a057-eb809e281033",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3981d4c7-00b6-407f-bc93-c73b9da3f2ce",
                            Name: codataccounting.String("Cameron Kuhn"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f2257411-faf4-4b75-84e4-72e802857a5b",
                            Name: codataccounting.String("Karen Gleichner"),
                        },
                    },
                    UnitAmount: 6521.25,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("7d575f14-00e7-464a-9733-4ec1b781b36a"),
                        Name: codataccounting.String("Gwendolyn Anderson"),
                    },
                    Description: codataccounting.String("repellendus"),
                    DiscountAmount: codataccounting.Float64(832.91),
                    DiscountPercentage: codataccounting.Float64(607.78),
                    ItemRef: &shared.ItemRef{
                        ID: "0efada20-0ef0-4422-ab21-64cf9ab8366c",
                        Name: codataccounting.String("Kathy Frami"),
                    },
                    Quantity: 8611.23,
                    SubTotal: codataccounting.Float64(6711.16),
                    TaxAmount: codataccounting.Float64(6176.57),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8837.8),
                        ID: codataccounting.String("06bee482-5c1f-4c0e-915c-80bff918544e"),
                        Name: codataccounting.String("Joel Cruickshank"),
                    },
                    TotalAmount: codataccounting.Float64(9851.09),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ce8f1977-773e-4635-a2a7-b408f05e3d48",
                                Name: codataccounting.String("Clint Ortiz"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "13a1f5fd-9425-49c0-b36f-25ea944f3b75",
                                Name: codataccounting.String("Dr. Alexandra Bernhard"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "c37a5126-2438-435b-bc05-a23a45cefc5f",
                                Name: codataccounting.String("Miss Frankie Braun DDS"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "e2169e51-0019-4c6d-85e3-4762799bfbbe",
                                Name: codataccounting.String("Lindsey Gislason"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("cum"),
                            ID: "2bb4ecae-6c3d-45db-bade-bd5daea4c506",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "aa94c026-44cf-45e9-99a4-578adc1ac600",
                            Name: codataccounting.String("Mr. Alonzo Schoen V"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "802e2ec0-9ff8-4f0f-816f-f3477c13e902",
                            Name: codataccounting.String("Mr. Jack Gottlieb"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b0960a66-8151-4a47-aaf9-23c5949f83f3",
                            Name: codataccounting.String("Angela Schaefer"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "76ffb901-c6ec-4bb4-a243-cf789ffafeda",
                            Name: codataccounting.String("Sheila Torphy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e6e0ac18-4c2b-49c2-87c8-8373a40e1942",
                            Name: codataccounting.String("Tony Connelly"),
                        },
                    },
                    UnitAmount: 3731.06,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("055756f5-d56d-40bd-8af2-dfe13db4f62c"),
                        Name: codataccounting.String("Lorenzo Flatley"),
                    },
                    Description: codataccounting.String("error"),
                    DiscountAmount: codataccounting.Float64(2524.73),
                    DiscountPercentage: codataccounting.Float64(978.1),
                    ItemRef: &shared.ItemRef{
                        ID: "aebc0b80-a692-44d3-b2ec-fcc8f895010f",
                        Name: codataccounting.String("Lynette Senger"),
                    },
                    Quantity: 4248.53,
                    SubTotal: codataccounting.Float64(9608.13),
                    TaxAmount: codataccounting.Float64(6525.15),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(651.21),
                        ID: codataccounting.String("804e54c8-2f16-48a3-a3c8-873e484380b1"),
                        Name: codataccounting.String("Milton Powlowski"),
                    },
                    TotalAmount: codataccounting.Float64(6464.91),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "75a60a04-c495-4cc6-9917-1b51c1bdb1cf",
                                Name: codataccounting.String("Lula Lowe"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("officiis"),
                            ID: "bdfc4ccc-a99b-4c7f-80b2-dce10873e42b",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "6d678878-ba85-481a-9820-8c54fefa9c95",
                            Name: codataccounting.String("Howard Torphy"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "565d307c-fee8-4120-ae28-13fa4a41c480",
                            Name: codataccounting.String("Mike Zulauf I"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2af03102-d514-4f4c-86f1-8bf9621a6a4f",
                            Name: codataccounting.String("Tamara O'Kon"),
                        },
                    },
                    UnitAmount: 9085.39,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e3e4be75-2c65-4b34-818e-3bb91c8d975e"),
                        Name: codataccounting.String("Silvia Langworth V"),
                    },
                    Description: codataccounting.String("illum"),
                    DiscountAmount: codataccounting.Float64(5265.84),
                    DiscountPercentage: codataccounting.Float64(9493.7),
                    ItemRef: &shared.ItemRef{
                        ID: "84f144f3-e07e-4dcc-8aa5-f3cabd905a97",
                        Name: codataccounting.String("Kelley Bashirian"),
                    },
                    Quantity: 4741.85,
                    SubTotal: codataccounting.Float64(1541.3),
                    TaxAmount: codataccounting.Float64(5147.67),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1527.42),
                        ID: codataccounting.String("27b2d309-470b-4f7a-8fa8-7cf535a6fae5"),
                        Name: codataccounting.String("Gwen Reichel"),
                    },
                    TotalAmount: codataccounting.Float64(304.26),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "321f023b-75d2-4367-be1a-0cc8df79f0a3",
                                Name: codataccounting.String("Ruben Sipes DDS"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "364b7c15-dfba-4ce1-88b1-c4ee2c8c6ce6",
                                Name: codataccounting.String("Marie Wunsch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b1c7cbdb-6eec-4743-b8ba-25317747dc91",
                                Name: codataccounting.String("Janie Stanton"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "af5dd672-3dc0-4f5a-a2f3-a6b700878756",
                                Name: codataccounting.String("Gail Fay"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("est"),
                            ID: "6c98b555-5408-40d4-8bca-cc6cbd6b5f3e",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "09304f92-6bad-4255-b819-b474b0ed20e5",
                            Name: codataccounting.String("Ruby Haag"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "f639a910-abdc-4ab6-a676-696e1ec00221",
                            Name: codataccounting.String("Johnny Farrell"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "89acb3ec-fda8-4d0c-949e-f03004978a61",
                            Name: codataccounting.String("Neal Boehm"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "20688f77-c1ff-4c71-9ca1-63f2a3c80a97",
                            Name: codataccounting.String("Darrin Flatley"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cddf857a-9e61-4876-86ab-21d29dfc94d6",
                            Name: codataccounting.String("Clay Schaefer"),
                        },
                    },
                    UnitAmount: 6143.46,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("provident"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sint"),
                        Currency: codataccounting.String("aperiam"),
                        CurrencyRate: codataccounting.Float64(494.56),
                        TotalAmount: codataccounting.Float64(4312.58),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6a6d2d00-0355-4338-8ec0-86fa21e9152c"),
                            Name: codataccounting.String("Ms. Melvin Brekke III"),
                        },
                        Currency: codataccounting.String("ducimus"),
                        CurrencyRate: codataccounting.Float64(7284.74),
                        ID: codataccounting.String("8e3c8db0-3408-4d6d-b64f-fd455906d126"),
                        Note: codataccounting.String("neque"),
                        PaidOnDate: codataccounting.String("quibusdam"),
                        Reference: codataccounting.String("numquam"),
                        TotalAmount: codataccounting.Float64(5231.09),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(8846.22),
            SourceModifiedDate: codataccounting.String("omnis"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "corporis": map[string]interface{}{
                        "dolores": "placeat",
                        "excepturi": "recusandae",
                        "quos": "dicta",
                        "sapiente": "ipsum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "0be3e432-02d7-4216-9765-06641870d9d2",
                SupplierName: codataccounting.String("inventore"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 9764.03,
            TotalTaxAmount: 6012.28,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 8345.87,
                    Name: "Sherry Batz",
                },
                shared.WithholdingTaxitems{
                    Amount: 8907.65,
                    Name: "Miss Simon Bosco V",
                },
                shared.WithholdingTaxitems{
                    Amount: 1917.24,
                    Name: "Anita Dare III",
                },
            },
        },
        BillCreditNoteID: "laudantium",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(703407),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```
