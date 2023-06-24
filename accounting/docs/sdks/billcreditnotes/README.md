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
            CurrencyRate: codataccounting.Float64(201.07),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("502a94bb-4f63-4c96-9e9a-3efa77dfb14c"),
                        Name: codataccounting.String("Nathaniel Hyatt"),
                    },
                    Description: codataccounting.String("non"),
                    DiscountAmount: codataccounting.Float64(5812.73),
                    DiscountPercentage: codataccounting.Float64(3132.18),
                    ItemRef: &shared.ItemRef{
                        ID: "efb9ba88-f3a6-4699-b074-ba4469b6e214",
                        Name: codataccounting.String("Miriam Hermann"),
                    },
                    Quantity: 5743.25,
                    SubTotal: codataccounting.Float64(336.25),
                    TaxAmount: codataccounting.Float64(6532.01),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9689.62),
                        ID: codataccounting.String("a563e251-6fe4-4c8b-b11e-5b7fd2ed0289"),
                        Name: codataccounting.String("Joan Satterfield"),
                    },
                    TotalAmount: codataccounting.Float64(8073.19),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "92601fb5-76b0-4d5f-8d30-c5fbb2587053",
                                Name: codataccounting.String("Dorothy Dach"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3d5fe9b9-0c28-4909-b3fe-49a8d9cbf486",
                                Name: codataccounting.String("Dawn Fadel"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("hic"),
                            ID: "9b77f3a4-1006-474e-bf69-280d1ba77a89",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f737ae42-03ce-45e6-a95d-8a0d446ce2af",
                            Name: codataccounting.String("Fannie Kub"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3be453f8-70b3-426b-9a73-429cdb1a8422",
                            Name: codataccounting.String("Cesar Hyatt"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d2322715-bf0c-4bb1-a31b-8b90f3443a11",
                            Name: codataccounting.String("Miss Billie Ward"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cf4b9218-79fc-4e95-bf73-ef7fbc7abd74",
                            Name: codataccounting.String("Gilberto Dickinson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0f5d2cff-7c70-4a45-a26d-436813f16d9f",
                            Name: codataccounting.String("Dixie Schamberger"),
                        },
                    },
                    UnitAmount: 7740.48,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("556146c3-e250-4fb0-88c4-2e141aac366c"),
                        Name: codataccounting.String("Mack Stoltenberg"),
                    },
                    Description: codataccounting.String("quasi"),
                    DiscountAmount: codataccounting.Float64(2703.28),
                    DiscountPercentage: codataccounting.Float64(2561.39),
                    ItemRef: &shared.ItemRef{
                        ID: "29074747-78a7-4bd4-a6d2-8c10ab3cdca4",
                        Name: codataccounting.String("Brittany Bernier II"),
                    },
                    Quantity: 8920.5,
                    SubTotal: codataccounting.Float64(3708.53),
                    TaxAmount: codataccounting.Float64(1334.65),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1970.54),
                        ID: codataccounting.String("c7e0bc71-78e4-4796-b2a7-0c688282aa48"),
                        Name: codataccounting.String("Cathy Huel"),
                    },
                    TotalAmount: codataccounting.Float64(1598.7),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "2e9817ee-17cb-4e61-a6b7-b95bc0ab3c20",
                                Name: codataccounting.String("Calvin Williamson"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("blanditiis"),
                            ID: "9fd871f9-9dd2-4efd-921a-a6f1e674bdb0",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "15756082-d68e-4a19-b1d1-7051339d0808",
                            Name: codataccounting.String("Iris Bernhard"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "394c2607-1f93-4f5f-8642-dac7af515cc4",
                            Name: codataccounting.String("Josephine Paucek"),
                        },
                    },
                    UnitAmount: 2460.63,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("aae8d678-64db-4b67-9fd5-e60b375ed4f6"),
                        Name: codataccounting.String("Rickey Ullrich"),
                    },
                    Description: codataccounting.String("sunt"),
                    DiscountAmount: codataccounting.Float64(9920.12),
                    DiscountPercentage: codataccounting.Float64(2415.45),
                    ItemRef: &shared.ItemRef{
                        ID: "3317fe35-b60e-4b1e-a426-555ba3c28744",
                        Name: codataccounting.String("Lionel Herman"),
                    },
                    Quantity: 5023.89,
                    SubTotal: codataccounting.Float64(5553.61),
                    TaxAmount: codataccounting.Float64(9425.84),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2015.17),
                        ID: codataccounting.String("a8d8f5c0-b2f2-4fb7-b194-a276b26916fe"),
                        Name: codataccounting.String("Faith Aufderhar"),
                    },
                    TotalAmount: codataccounting.Float64(2748.23),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "94e3698f-447f-4603-a8b4-45e80ca55efd",
                                Name: codataccounting.String("Deborah Turcotte"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("in"),
                            ID: "e1858b6a-89fb-4e3a-9aa8-e4824d0ab407",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "88e51862-065e-4904-b3b1-194b8abf603a",
                            Name: codataccounting.String("Lindsey Witting"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e0ab7da8-a50c-4e18-bf86-bc173d689eee",
                            Name: codataccounting.String("Darrell Collier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "8d986e88-1ead-44f0-a101-2563f94e29e9",
                            Name: codataccounting.String("Sylvia Upton"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2a57a15b-e3e0-4608-87e2-b6e3ab8845f0",
                            Name: codataccounting.String("Katrina Kovacek"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0ff2a54a-31e9-4476-8a3e-865e7956f925",
                            Name: codataccounting.String("Lynda Heathcote"),
                        },
                    },
                    UnitAmount: 8217.19,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a660ff57-bfaa-4d4f-9efc-1b4512c10326"),
                        Name: codataccounting.String("Deanna Swaniawski"),
                    },
                    Description: codataccounting.String("sapiente"),
                    DiscountAmount: codataccounting.Float64(4332.79),
                    DiscountPercentage: codataccounting.Float64(1173.2),
                    ItemRef: &shared.ItemRef{
                        ID: "5199ebfd-0e9f-4e6c-a32c-a3aed0117996",
                        Name: codataccounting.String("Joyce Cummings"),
                    },
                    Quantity: 8965.82,
                    SubTotal: codataccounting.Float64(585.34),
                    TaxAmount: codataccounting.Float64(2711.13),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4706.21),
                        ID: codataccounting.String("71778ff6-1d01-4747-a360-a15db6a66065"),
                        Name: codataccounting.String("Alfonso Bernier"),
                    },
                    TotalAmount: codataccounting.Float64(9139.92),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ab5851d6-c645-4b08-b618-91baa0fe1ade",
                                Name: codataccounting.String("Mary Leuschke"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f8c5f350-d8cd-4b5a-b418-143010421813",
                                Name: codataccounting.String("Ms. Vernon Crooks"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "ce7e253b-6684-451c-ac6e-205e16deab3f",
                                Name: codataccounting.String("Earnest McClure"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("blanditiis"),
                            ID: "a6458427-3a84-418d-9623-09fb0929921a",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeProject,
                        ProjectRef: &shared.ProjectRef{
                            ID: "b9f58c4d-86e6-48e4-be05-6013f59da757",
                            Name: codataccounting.String("Alvin Marquardt"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ef66ef1c-aa33-483c-abeb-477373c8d72f",
                            Name: codataccounting.String("Dr. Megan Spinka"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "1f2c4310-661e-4963-89e1-cf9e06e3a437",
                            Name: codataccounting.String("Sharon Adams"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6b6bc9b8-f759-4eac-95a9-741d31135296",
                            Name: codataccounting.String("Patty Reinger"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "72026114-35e1-439d-bc22-59b1abda8c07",
                            Name: codataccounting.String("Ms. Tasha Block"),
                        },
                    },
                    UnitAmount: 7551.06,
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
                        CurrencyRate: codataccounting.Float64(1729.51),
                        TotalAmount: codataccounting.Float64(8247.98),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1ad879ee-b966-45b8-9efb-d02bae0be2d7"),
                            Name: codataccounting.String("Fred Champlin"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2393.37),
                        ID: codataccounting.String("ea4b5197-f924-443d-a7ce-52b895c537c6"),
                        Note: codataccounting.String("modi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("magnam"),
                        TotalAmount: codataccounting.Float64(9149.71),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aperiam": map[string]interface{}{
                        "ratione": "labore",
                        "totam": "occaecati",
                        "voluptas": "quo",
                    },
                    "velit": map[string]interface{}{
                        "fuga": "nostrum",
                        "est": "impedit",
                        "delectus": "tempore",
                        "vero": "odit",
                    },
                    "repellat": map[string]interface{}{
                        "nemo": "reprehenderit",
                        "aperiam": "odio",
                        "minima": "in",
                        "ducimus": "excepturi",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "29177dea-c646-4ecb-9734-09e3eb1e5a2b",
                SupplierName: codataccounting.String("dicta"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 8998.67,
                    Name: "Larry Kuphal Sr.",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(386447),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateBillCreditNoteRequest](../../models/operations/createbillcreditnoterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreateBillCreditNoteResponse](../../models/operations/createbillcreditnoteresponse.md), error**


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
        BillCreditNoteID: "pariatur",
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

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetBillCreditNoteRequest](../../models/operations/getbillcreditnoterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.GetBillCreditNoteResponse](../../models/operations/getbillcreditnoteresponse.md), error**


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

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetCreateUpdateBillCreditNotesModelRequest](../../models/operations/getcreateupdatebillcreditnotesmodelrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |


### Response

**[*operations.GetCreateUpdateBillCreditNotesModelResponse](../../models/operations/getcreateupdatebillcreditnotesmodelresponse.md), error**


## List

The *List bill credit notes* endpoint returns a list of [bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) for a given company's connection.

[Bill credit notes](https://docs.codat.io/accounting-api#/schemas/BillCreditNote) are issued by a supplier for the purpose of recording credit.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

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
        Query: codataccounting.String("libero"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillCreditNotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListBillCreditNotesRequest](../../models/operations/listbillcreditnotesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.ListBillCreditNotesResponse](../../models/operations/listbillcreditnotesresponse.md), error**


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
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(3679.17),
            DiscountPercentage: 0,
            ID: codataccounting.String("1509398f-98e2-436d-8a5d-c042e0c74ffc"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("fc95fa88-970e-4189-9bb3-0fcb33ea055b"),
                        Name: codataccounting.String("Mae Krajcik"),
                    },
                    Description: codataccounting.String("non"),
                    DiscountAmount: codataccounting.Float64(2981.87),
                    DiscountPercentage: codataccounting.Float64(9322.96),
                    ItemRef: &shared.ItemRef{
                        ID: "2f52d82d-3513-4bb6-b48b-656bcdb35ff2",
                        Name: codataccounting.String("Marcus Prohaska"),
                    },
                    Quantity: 3455.06,
                    SubTotal: codataccounting.Float64(2072.96),
                    TaxAmount: codataccounting.Float64(4800.61),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6649.65),
                        ID: codataccounting.String("8cd9e731-9c17-47d5-a5f7-7b114eeb52ff"),
                        Name: codataccounting.String("Maxine Hilll"),
                    },
                    TotalAmount: codataccounting.Float64(1972.59),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "814d4c98-e0c2-4bb8-9eb7-5dad636c6005",
                                Name: codataccounting.String("Wendy Stanton"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "b31180f7-39ae-49e0-97eb-809e2810331f",
                                Name: codataccounting.String("Dr. Misty Lindgren"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("minus"),
                            ID: "700b607f-3c93-4c73-b9da-3f2ceda7e23f",
                        },
                        IsBilledTo: shared.BilledToTypeUnknown,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.ProjectRef{
                            ID: "57411faf-4b75-444e-872e-802857a5b404",
                            Name: codataccounting.String("Robin O'Hara"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "75f1400e-764a-4d73-b4ec-1b781b36a080",
                            Name: codataccounting.String("Mr. Dwayne Sipes PhD"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fada200e-f042-42eb-a164-cf9ab8366c72",
                            Name: codataccounting.String("Faith Zulauf"),
                        },
                    },
                    UnitAmount: 6176.57,
                },
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e06bee48-25c1-4fc0-a115-c80bff918544"),
                        Name: codataccounting.String("Simon Gleason"),
                    },
                    Description: codataccounting.String("vero"),
                    DiscountAmount: codataccounting.Float64(9851.09),
                    DiscountPercentage: codataccounting.Float64(7737.11),
                    ItemRef: &shared.ItemRef{
                        ID: "ce8f1977-773e-4635-a2a7-b408f05e3d48",
                        Name: codataccounting.String("Clint Ortiz"),
                    },
                    Quantity: 1105.22,
                    SubTotal: codataccounting.Float64(2010.96),
                    TaxAmount: codataccounting.Float64(6308.32),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(748.95),
                        ID: codataccounting.String("f5fd9425-9c0b-436f-a5ea-944f3b756c11"),
                        Name: codataccounting.String("Charlie Schaefer"),
                    },
                    TotalAmount: codataccounting.Float64(6837.27),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "12624383-5bbc-405a-a3a4-5cefc5fde10a",
                                Name: codataccounting.String("Della Treutel III"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "9e510019-c6dc-45e3-8762-799bfbbe6949",
                                Name: codataccounting.String("Jonathon Collins"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("incidunt"),
                            ID: "ecae6c3d-5db3-4ade-bd5d-aea4c506a8aa",
                        },
                        IsBilledTo: shared.BilledToTypeCustomer,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "c02644cf-5e9d-49a4-978a-dc1ac600dec0",
                            Name: codataccounting.String("Julie Murazik"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "2e2ec09f-f8f0-4f81-aff3-477c13e902c1",
                            Name: codataccounting.String("Cheryl Conn"),
                        },
                    },
                    UnitAmount: 46.54,
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
                        CurrencyRate: codataccounting.Float64(3881.69),
                        TotalAmount: codataccounting.Float64(4016.88),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("8151a472-af92-43c5-949f-83f350cf876f"),
                            Name: codataccounting.String("Mr. Robin Miller"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(9087.34),
                        ID: codataccounting.String("cbb4e243-cf78-49ff-afed-a53e5ae6e0ac"),
                        Note: codataccounting.String("quasi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("eius"),
                        TotalAmount: codataccounting.Float64(7861.89),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6200.54),
                        TotalAmount: codataccounting.Float64(7935.68),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("247c8837-3a40-4e19-82f3-2e55055756f5"),
                            Name: codataccounting.String("Maurice Hoppe MD"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(177.68),
                        ID: codataccounting.String("af2dfe13-db4f-462c-ba3f-8941aebc0b80"),
                        Note: codataccounting.String("deserunt"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("excepturi"),
                        TotalAmount: codataccounting.Float64(1674.35),
                    },
                },
            },
            RemainingCredit: codataccounting.Float64(0),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.BillCreditNoteStatusPaid,
            SubTotal: 805.78,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolor": map[string]interface{}{
                        "sed": "accusamus",
                        "optio": "delectus",
                        "minus": "quo",
                    },
                    "quos": map[string]interface{}{
                        "voluptatum": "iste",
                        "corporis": "accusantium",
                        "illo": "aut",
                        "doloribus": "nostrum",
                    },
                    "at": map[string]interface{}{
                        "neque": "pariatur",
                        "vel": "sapiente",
                        "mollitia": "quae",
                        "quos": "aperiam",
                    },
                    "non": map[string]interface{}{
                        "ad": "aliquam",
                        "quisquam": "quas",
                        "consequuntur": "maiores",
                        "inventore": "aliquid",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "8a363c88-73e4-4843-80b1-f6b8ca275a60",
                SupplierName: codataccounting.String("fuga"),
            },
            TotalAmount: 805.78,
            TotalDiscount: 0,
            TotalTaxAmount: 0,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 2958.92,
                    Name: "Jay Morar",
                },
            },
        },
        BillCreditNoteID: "placeat",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(378403),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateBillCreditNoteRequest](../../models/operations/updatebillcreditnoterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.UpdateBillCreditNoteResponse](../../models/operations/updatebillcreditnoteresponse.md), error**

