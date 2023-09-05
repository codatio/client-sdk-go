# AccountingCreditNotes

## Overview

Credit notes

### Available Operations

* [CreateAccountingCreditNote](#createaccountingcreditnote) - Create credit note

## CreateAccountingCreditNote

The *Create credit note* endpoint creates a new [credit note](https://docs.codat.io/accounting-api#/schemas/CreditNote) for a given company's connection.

[Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote) are issued to a customer to indicate debt, typically with reference to a previously issued invoice and/or purchase.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-creditNotes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingCreditNotes.CreateAccountingCreditNote(ctx, operations.CreateAccountingCreditNoteRequest{
        AccountingCreditNote: &shared.AccountingCreditNote{
            AdditionalTaxAmount: codatsynccommerce.Float64(8700.13),
            AdditionalTaxPercentage: codatsynccommerce.Float64(8700.88),
            AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            CreditNoteNumber: codatsynccommerce.String("molestiae"),
            Currency: codatsynccommerce.String("EUR"),
            CurrencyRate: codatsynccommerce.Float64(8009.11),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("esse"),
                ID: "8ca1ba92-8fc8-4167-82cb-739205929396",
            },
            DiscountPercentage: 9437.49,
            ID: codatsynccommerce.String("ea7596eb-10fa-4aa2-b52c-5955907aff1a"),
            IssueDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            LineItems: []shared.CreditNoteLineItem{
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("2fa94677-3925-41aa-92c3-f5ad019da1ff"),
                        Name: codatsynccommerce.String("Jessie Langosh V"),
                    },
                    Description: codatsynccommerce.String("voluptate"),
                    DiscountAmount: codatsynccommerce.Float64(7392.64),
                    DiscountPercentage: codatsynccommerce.Float64(199.87),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "074f1547-1b5e-46e1-bb99-d488e1e91e45",
                        Name: codatsynccommerce.String("Monique Spinka"),
                    },
                    Quantity: 7163.27,
                    SubTotal: codatsynccommerce.Float64(8413.86),
                    TaxAmount: codatsynccommerce.Float64(2894.06),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsynccommerce.Float64(2647.3),
                        ID: codatsynccommerce.String("269802d5-02a9-44bb-8f63-c969e9a3efa7"),
                        Name: codatsynccommerce.String("Angel Wolff II"),
                    },
                    TotalAmount: codatsynccommerce.Float64(7670.24),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "66ae395e-fb9b-4a88-b3a6-6997074ba446",
                                Name: codatsynccommerce.String("Robin Keebler"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "14195989-0afa-4563-a251-6fe4c8b711e5",
                                Name: codatsynccommerce.String("Jessie Zulauf"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "ed028921-cddc-4692-a01f-b576b0d5f0d3",
                                Name: codatsynccommerce.String("Erma Hessel"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "b2587053-202c-473d-9fe9-b90c28909b3f",
                                Name: codatsynccommerce.String("Edwin Morar"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("pariatur"),
                            ID: "9cbf4863-3323-4f9b-b7f3-a4100674ebf6",
                        },
                        IsBilledTo: shared.BilledToTypeNotApplicable,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "80d1ba77-a89e-4bf7-b7ae-4203ce5e6a95",
                            Name: codatsynccommerce.String("Dr. Jimmie Murphy"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("invoice"),
                            ID: codatsynccommerce.String("6ce2af7a-73cf-43be-853f-870b326b5a73"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "29cdb1a8-422b-4b67-9d23-22715bf0cbb1",
                            Name: codatsynccommerce.String("Dale Boehm"),
                        },
                        shared.TrackingCategoryRefsitems{
                            ID: "b90f3443-a110-48e0-adcf-4b921879fce9",
                            Name: codatsynccommerce.String("Tiffany Willms"),
                        },
                    },
                    UnitAmount: 8788.7,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("f7fbc7ab-d74d-4d39-80f5-d2cff7c70a45"),
                        Name: codatsynccommerce.String("Judy Keebler"),
                    },
                    Description: codatsynccommerce.String("ratione"),
                    DiscountAmount: codatsynccommerce.Float64(4011.32),
                    DiscountPercentage: codatsynccommerce.Float64(5113.19),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "13f16d9f-5fce-46c5-9614-6c3e250fb008",
                        Name: codatsynccommerce.String("Jim Corkery I"),
                    },
                    Quantity: 896.03,
                    SubTotal: codatsynccommerce.Float64(6774.12),
                    TaxAmount: codatsynccommerce.Float64(6720.48),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsynccommerce.Float64(8104.24),
                        ID: codatsynccommerce.String("366c8dd6-b144-4290-b474-778a7bd466d2"),
                        Name: codatsynccommerce.String("Miss Devin Bogan"),
                    },
                    TotalAmount: codatsynccommerce.Float64(2065.94),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "dca42519-04e5-423c-be0b-c7178e4796f2",
                                Name: codatsynccommerce.String("Fernando Barton"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "88282aa4-8256-42f2-a2e9-817ee17cbe61",
                                Name: codatsynccommerce.String("Cecil Pollich"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "95bc0ab3-c20c-44f3-b89f-d871f99dd2ef",
                                Name: codatsynccommerce.String("Miss Peter Cronin"),
                            },
                            shared.TrackingCategoryRefsitems{
                                ID: "6f1e674b-db04-4f15-b560-82d68ea19f1d",
                                Name: codatsynccommerce.String("Allison Bednar I"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("adipisci"),
                            ID: "9d08086a-1840-4394-8260-71f93f5f0642",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeNotApplicable,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "c7af515c-c413-4aa6-baae-8d67864dbb67",
                            Name: codatsynccommerce.String("Lela Shields"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("invoice"),
                            ID: codatsynccommerce.String("0b375ed4-f6fb-4ee4-9f33-317fe35b60eb"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "ea426555-ba3c-4287-84ed-53b88f3a8d8f",
                            Name: codatsynccommerce.String("Della Bailey"),
                        },
                    },
                    UnitAmount: 9679.66,
                },
                shared.CreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codatsynccommerce.String("2fb7b194-a276-4b26-916f-e1f08f4294e3"),
                        Name: codatsynccommerce.String("Kristina Lueilwitz"),
                    },
                    Description: codatsynccommerce.String("tempora"),
                    DiscountAmount: codatsynccommerce.Float64(4554.44),
                    DiscountPercentage: codatsynccommerce.Float64(9700.76),
                    IsDirectIncome: codatsynccommerce.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "603e8b44-5e80-4ca5-9efd-20e457e1858b",
                        Name: codatsynccommerce.String("Lee Lehner"),
                    },
                    Quantity: 7105.29,
                    SubTotal: codatsynccommerce.Float64(8928.63),
                    TaxAmount: codatsynccommerce.Float64(2049.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codatsynccommerce.Float64(6771.15),
                        ID: codatsynccommerce.String("5aa8e482-4d0a-4b40-b508-8e51862065e9"),
                        Name: codatsynccommerce.String("Eva Wisozk"),
                    },
                    TotalAmount: codatsynccommerce.Float64(1157.03),
                    Tracking: &shared.CreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRefsitems{
                            shared.TrackingCategoryRefsitems{
                                ID: "94b8abf6-03a7-49f9-9fe0-ab7da8a50ce1",
                                Name: codatsynccommerce.String("Mitchell Zboncak"),
                            },
                        },
                        CustomerRef: &shared.AccountingCustomerRef{
                            CompanyName: codatsynccommerce.String("quidem"),
                            ID: "c173d689-eee9-4526-b8d9-86e881ead4f0",
                        },
                        IsBilledTo: shared.BilledToTypeProject,
                        IsRebilledTo: shared.BilledToTypeUnknown,
                        ProjectRef: &shared.CreditNoteLineItemTrackingProjectReference{
                            ID: "012563f9-4e29-4e97-be92-2a57a15be3e0",
                            Name: codatsynccommerce.String("Ms. Melissa Larson"),
                        },
                        RecordRef: &shared.RecordRef{
                            DataType: codatsynccommerce.String("journalEntry"),
                            ID: codatsynccommerce.String("b6e3ab88-45f0-4597-a60f-f2a54a31e947"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRefsitems{
                        shared.TrackingCategoryRefsitems{
                            ID: "4a3e865e-7956-4f92-91a5-a9da660ff57b",
                            Name: codatsynccommerce.String("Oliver Osinski"),
                        },
                        shared.TrackingCategoryRefsitems{
                            ID: "f9efc1b4-512c-4103-a648-dc2f615199eb",
                            Name: codatsynccommerce.String("Gilberto Bechtelar"),
                        },
                    },
                    UnitAmount: 9834.27,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("aliquid"),
            PaymentAllocations: []shared.PaymentAllocationsitems{
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(1478.08),
                        TotalAmount: codatsynccommerce.Float64(7649.95),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("a3aed011-7996-4312-bde0-4771778ff61d"),
                            Name: codatsynccommerce.String("Cheryl Kub"),
                        },
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(2352.63),
                        ID: codatsynccommerce.String("60a15db6-a660-4659-a1ad-eaab5851d6c6"),
                        Note: codatsynccommerce.String("ut"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("expedita"),
                        TotalAmount: codatsynccommerce.Float64(299.5),
                    },
                },
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(3996.6),
                        TotalAmount: codatsynccommerce.Float64(1097.84),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("891baa0f-e1ad-4e00-8e6f-8c5f350d8cdb"),
                            Name: codatsynccommerce.String("Molly Fadel IV"),
                        },
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(2745.75),
                        ID: codatsynccommerce.String("30104218-13d5-4208-ace7-e253b668451c"),
                        Note: codatsynccommerce.String("autem"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("laboriosam"),
                        TotalAmount: codatsynccommerce.Float64(9272.12),
                    },
                },
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(3502.07),
                        TotalAmount: codatsynccommerce.Float64(8956.92),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("16deab3f-ec95-478a-a458-4273a8418d16"),
                            Name: codatsynccommerce.String("Edith Beahan"),
                        },
                        Currency: codatsynccommerce.String("EUR"),
                        CurrencyRate: codatsynccommerce.Float64(38.6),
                        ID: codatsynccommerce.String("929921ae-fb9f-458c-8d86-e68e4be05601"),
                        Note: codatsynccommerce.String("non"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("enim"),
                        TotalAmount: codatsynccommerce.Float64(5752.13),
                    },
                },
                shared.PaymentAllocationsitems{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Currency: codatsynccommerce.String("USD"),
                        CurrencyRate: codatsynccommerce.Float64(4585.03),
                        TotalAmount: codatsynccommerce.Float64(3644.63),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codatsynccommerce.String("7a59ecfe-f66e-4f1c-aa33-83c2beb47737"),
                            Name: codatsynccommerce.String("Angelica Leuschke"),
                        },
                        Currency: codatsynccommerce.String("GBP"),
                        CurrencyRate: codatsynccommerce.Float64(9749.9),
                        ID: codatsynccommerce.String("64d1db1f-2c43-4106-a1e9-6349e1cf9e06"),
                        Note: codatsynccommerce.String("itaque"),
                        PaidOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                        Reference: codatsynccommerce.String("laborum"),
                        TotalAmount: codatsynccommerce.Float64(2503.98),
                    },
                },
            },
            RemainingCredit: 2244.67,
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CreditNoteStatusUnknown,
            SubTotal: 399.92,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "officia": map[string]interface{}{
                        "ea": "quidem",
                        "voluptas": "facilis",
                        "placeat": "perspiciatis",
                        "expedita": "deleniti",
                    },
                },
            },
            TotalAmount: 9543.34,
            TotalDiscount: 4555.79,
            TotalTaxAmount: 3519.36,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 8975.43,
                    Name: "Rodolfo Hintz",
                },
                shared.WithholdingTaxitems{
                    Amount: 6216.66,
                    Name: "Lucille Bogan",
                },
                shared.WithholdingTaxitems{
                    Amount: 1124.27,
                    Name: "Florence Hand",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(403026),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateAccountingCreditNoteRequest](../../models/operations/createaccountingcreditnoterequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.CreateAccountingCreditNoteResponse](../../models/operations/createaccountingcreditnoteresponse.md), error**

