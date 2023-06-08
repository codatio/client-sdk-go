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
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(9182.36),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b99d488e-1e91-4e45-8ad2-abd44269802d"),
                        Name: codataccounting.String("Linda Corkery"),
                    },
                    Description: codataccounting.String("tempora"),
                    DiscountAmount: codataccounting.Float64(7037.37),
                    DiscountPercentage: codataccounting.Float64(7351.94),
                    ItemRef: &shared.ItemRef{
                        ID: "4f63c969-e9a3-4efa-b7df-b14cd66ae395",
                        Name: codataccounting.String("Toby Pouros"),
                    },
                    Quantity: 6596.69,
                    SubTotal: codataccounting.Float64(5013.24),
                    TaxAmount: codataccounting.Float64(5332.06),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9560.84),
                        ID: codataccounting.String("3a669970-74ba-4446-9b6e-2141959890af"),
                        Name: codataccounting.String("Tommy Kemmer"),
                    },
                    TotalAmount: codataccounting.Float64(1412.64),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "16fe4c8b-711e-45b7-bd2e-d028921cddc6",
                                Name: codataccounting.String("Mr. Harry Jaskolski"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b576b0d5-f0d3-40c5-bbb2-587053202c73",
                                Name: codataccounting.String("Dean Welch"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("facilis"),
                            ID: "90c28909-b3fe-449a-8d9c-bf48633323f9",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "7f3a4100-674e-4bf6-9280-d1ba77a89ebf",
                            Name: codataccounting.String("Edna Klocko"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "203ce5e6-a95d-48a0-9446-ce2af7a73cf3",
                            Name: codataccounting.String("Tomas Funk"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f870b326-b5a7-4342-9cdb-1a8422bb679d",
                            Name: codataccounting.String("Gladys Considine"),
                        },
                    },
                    UnitAmount: 1248.33,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(7653.26),
                        TotalAmount: codataccounting.Float64(7469.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b1e31b8b-90f3-4443-a110-8e0adcf4b921"),
                            Name: codataccounting.String("Darren McClure"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6064.76),
                        ID: codataccounting.String("53f73ef7-fbc7-4abd-b4dd-39c0f5d2cff7"),
                        Note: codataccounting.String("eligendi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("alias"),
                        TotalAmount: codataccounting.Float64(6394.73),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(4104.92),
                        TotalAmount: codataccounting.Float64(1369),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6d436813-f16d-49f5-bce6-c556146c3e25"),
                            Name: codataccounting.String("Mr. Elsa Reinger"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(3045.82),
                        ID: codataccounting.String("2e141aac-366c-48dd-ab14-42907474778a"),
                        Note: codataccounting.String("reprehenderit"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("fugiat"),
                        TotalAmount: codataccounting.Float64(2835.19),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8268.71),
                        TotalAmount: codataccounting.Float64(1811.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8c10ab3c-dca4-4251-904e-523c7e0bc717"),
                            Name: codataccounting.String("Sheldon Hackett"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(9594.34),
                        ID: codataccounting.String("2a70c688-282a-4a48-a562-f222e9817ee1"),
                        Note: codataccounting.String("esse"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("nam"),
                        TotalAmount: codataccounting.Float64(8771.31),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "saepe": map[string]interface{}{
                        "harum": "molestiae",
                        "rerum": "occaecati",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "5bc0ab3c-20c4-4f37-89fd-871f99dd2efd",
                SupplierName: codataccounting.String("veritatis"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 944.58,
                    Name: "Shannon Jacobi DVM",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(424032),
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
        BillCreditNoteID: "in",
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
        Query: codataccounting.String("eius"),
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
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            BillCreditNoteNumber: codataccounting.String("91fe2a83-e161-4c21-929d-c5c10c4b07e5"),
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(7422.38),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f1575608-2d68-4ea1-9f1d-17051339d080"),
                        Name: codataccounting.String("Ms. Duane O'Conner"),
                    },
                    Description: codataccounting.String("voluptatem"),
                    DiscountAmount: codataccounting.Float64(2211.61),
                    DiscountPercentage: codataccounting.Float64(5801.52),
                    ItemRef: &shared.ItemRef{
                        ID: "4c26071f-93f5-4f06-82da-c7af515cc413",
                        Name: codataccounting.String("Orlando Homenick"),
                    },
                    Quantity: 6658.59,
                    SubTotal: codataccounting.Float64(9268.8),
                    TaxAmount: codataccounting.Float64(5173.09),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8539.4),
                        ID: codataccounting.String("67864dbb-675f-4d5e-a0b3-75ed4f6fbee4"),
                        Name: codataccounting.String("Ollie Flatley"),
                    },
                    TotalAmount: codataccounting.Float64(1059.06),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "fe35b60e-b1ea-4426-955b-a3c28744ed53",
                                Name: codataccounting.String("Guy Luettgen"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a8d8f5c0-b2f2-4fb7-b194-a276b26916fe",
                                Name: codataccounting.String("Faith Aufderhar"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("incidunt"),
                            ID: "294e3698-f447-4f60-be8b-445e80ca55ef",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "0e457e18-58b6-4a89-bbe3-a5aa8e4824d0",
                            Name: codataccounting.String("Ms. Rudolph Gusikowski"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "88e51862-065e-4904-b3b1-194b8abf603a",
                            Name: codataccounting.String("Lindsey Witting"),
                        },
                    },
                    UnitAmount: 9627.71,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e0ab7da8-a50c-4e18-bf86-bc173d689eee"),
                        Name: codataccounting.String("Darrell Collier"),
                    },
                    Description: codataccounting.String("corrupti"),
                    DiscountAmount: codataccounting.Float64(8717.86),
                    DiscountPercentage: codataccounting.Float64(6216.93),
                    ItemRef: &shared.ItemRef{
                        ID: "86e881ea-d4f0-4e10-9256-3f94e29e973e",
                        Name: codataccounting.String("Carlos Considine"),
                    },
                    Quantity: 4402.64,
                    SubTotal: codataccounting.Float64(6255.28),
                    TaxAmount: codataccounting.Float64(764.86),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3613.06),
                        ID: codataccounting.String("be3e0608-07e2-4b6e-bab8-845f0597a60f"),
                        Name: codataccounting.String("Todd O'Reilly"),
                    },
                    TotalAmount: codataccounting.Float64(6803.49),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "1e94764a-3e86-45e7-956f-9251a5a9da66",
                                Name: codataccounting.String("Kristie Wyman"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("nam"),
                            ID: "faad4f9e-fc1b-4451-ac10-32648dc2f615",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "9ebfd0e9-fe6c-4632-8a3a-ed0117996312",
                            Name: codataccounting.String("Mrs. Orville Treutel"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1778ff61-d017-4476-b60a-15db6a660659",
                            Name: codataccounting.String("Raymond Muller"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "aab5851d-6c64-45b0-8b61-891baa0fe1ad",
                            Name: codataccounting.String("Donald Abbott"),
                        },
                    },
                    UnitAmount: 4042.44,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("Bill Credit Note with 1 line items, totaling 805.78"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(9920.74),
                        TotalAmount: codataccounting.Float64(1905.67),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("50d8cdb5-a341-4814-b010-421813d5208e"),
                            Name: codataccounting.String("Erick Klocko"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1995.96),
                        ID: codataccounting.String("b668451c-6c6e-4205-a16d-eab3fec9578a"),
                        Note: codataccounting.String("voluptas"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("nemo"),
                        TotalAmount: codataccounting.Float64(5510.79),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(4959.7),
                        TotalAmount: codataccounting.Float64(2005.16),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a8418d16-2309-4fb0-9299-21aefb9f58c4"),
                            Name: codataccounting.String("Byron Johns"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(9366.18),
                        ID: codataccounting.String("4be05601-3f59-4da7-97a5-9ecfef66ef1c"),
                        Note: codataccounting.String("fuga"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("consectetur"),
                        TotalAmount: codataccounting.Float64(2448.89),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(7730.35),
                        TotalAmount: codataccounting.Float64(1660.47),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("beb47737-3c8d-472f-a4d1-db1f2c431066"),
                            Name: codataccounting.String("Leigh Mante"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5964.33),
                        ID: codataccounting.String("e1cf9e06-e3a4-4370-80ae-6b6bc9b8f759"),
                        Note: codataccounting.String("necessitatibus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("impedit"),
                        TotalAmount: codataccounting.Float64(3730.4),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "error": map[string]interface{}{
                        "labore": "veritatis",
                        "vero": "consectetur",
                    },
                    "vitae": map[string]interface{}{
                        "dolorem": "ad",
                    },
                    "qui": map[string]interface{}{
                        "ex": "nemo",
                        "soluta": "libero",
                        "rem": "dolorum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "72026114-35e1-439d-bc22-59b1abda8c07",
                SupplierName: codataccounting.String("ipsa"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 800.61,
                    Name: "Lena Greenholt",
                },
                shared.WithholdingTaxitems{
                    Amount: 322.73,
                    Name: "Tamara D'Amore MD",
                },
                shared.WithholdingTaxitems{
                    Amount: 8173.39,
                    Name: "Javier McKenzie",
                },
                shared.WithholdingTaxitems{
                    Amount: 7332.89,
                    Name: "Angel Jones",
                },
            },
        },
        BillCreditNoteID: "laudantium",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(357207),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```
