# CreateBillCreditNote
Available in: `BillCreditNotes`

Posts a new billCreditNote to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill credit note model](https://docs.codat.io/accounting-api#/operations/get-create-update-billCreditNotes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes) for integrations that support creating bill credit notes.

## Example Usage
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
    req := operations.CreateBillCreditNoteRequest{
        BillCreditNote: &shared.BillCreditNote{
            AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            BillCreditNoteNumber: codataccounting.String("aliquid"),
            Currency: codataccounting.String("laborum"),
            CurrencyRate: codataccounting.Float64(8811.04),
            DiscountPercentage: 2497.96,
            ID: codataccounting.String("95efb9ba-88f3-4a66-9970-74ba4469b6e2"),
            IssueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            LineItems: []shared.BillCreditNoteLineItem{
                shared.BillCreditNoteLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("41959890-afa5-463e-a516-fe4c8b711e5b"),
                        Name: codataccounting.String("Kristie Spencer"),
                    },
                    Description: codataccounting.String("pariatur"),
                    DiscountAmount: codataccounting.Float64(375.59),
                    DiscountPercentage: codataccounting.Float64(1624.93),
                    ItemRef: &shared.ItemRef{
                        ID: "8921cddc-6926-401f-b576-b0d5f0d30c5f",
                        Name: codataccounting.String("Robin D'Amore"),
                    },
                    Quantity: 4895.49,
                    SubTotal: codataccounting.Float64(543.38),
                    TaxAmount: codataccounting.Float64(3389.85),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1999.96),
                        ID: codataccounting.String("202c73d5-fe9b-490c-a890-9b3fe49a8d9c"),
                        Name: codataccounting.String("Toby Hahn"),
                    },
                    TotalAmount: codataccounting.Float64(2123.9),
                    Tracking: &shared.BillCreditNoteLineItemTracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "323f9b77-f3a4-4100-a74e-bf69280d1ba7",
                                Name: codataccounting.String("Sonya Leuschke"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("distinctio"),
                            ID: "f737ae42-03ce-45e6-a95d-8a0d446ce2af",
                        },
                        IsBilledTo: shared.BilledToTypeEnumNotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "73cf3be4-53f8-470b-b26b-5a73429cdb1a",
                            Name: codataccounting.String("Randall Cole"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "679d2322-715b-4f0c-bb1e-31b8b90f3443",
                            Name: codataccounting.String("Ms. Joe Berge"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0adcf4b9-2187-49fc-a953-f73ef7fbc7ab",
                            Name: codataccounting.String("Allan Greenholt"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "39c0f5d2-cff7-4c70-a456-26d436813f16",
                            Name: codataccounting.String("Marshall Wiza"),
                        },
                    },
                    UnitAmount: 7888.73,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("saepe"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("impedit"),
                        CurrencyRate: codataccounting.Float64(3592.71),
                        TotalAmount: codataccounting.Float64(3331.45),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("6146c3e2-50fb-4008-842e-141aac366c8d"),
                            Name: codataccounting.String("Mrs. Shane Reinger"),
                        },
                        Currency: codataccounting.String("explicabo"),
                        CurrencyRate: codataccounting.Float64(5919.35),
                        ID: codataccounting.String("07474778-a7bd-4466-928c-10ab3cdca425"),
                        Note: codataccounting.String("ab"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("cupiditate"),
                        TotalAmount: codataccounting.Float64(96.88),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("tempora"),
                        CurrencyRate: codataccounting.Float64(8920.5),
                        TotalAmount: codataccounting.Float64(3708.53),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("23c7e0bc-7178-4e47-96f2-a70c688282aa"),
                            Name: codataccounting.String("Leah Champlin"),
                        },
                        Currency: codataccounting.String("fugit"),
                        CurrencyRate: codataccounting.Float64(9564.06),
                        ID: codataccounting.String("222e9817-ee17-4cbe-a1e6-b7b95bc0ab3c"),
                        Note: codataccounting.String("consequuntur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("consequatur"),
                        TotalAmount: codataccounting.Float64(7963.92),
                    },
                },
            },
            RemainingCredit: 3082.86,
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.BillCreditNoteStatusEnumPartiallyPaid,
            SubTotal: 2328.65,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "blanditiis": map[string]interface{}{
                        "a": "nulla",
                        "quas": "esse",
                        "quasi": "a",
                    },
                    "error": map[string]interface{}{
                        "pariatur": "possimus",
                        "quia": "eveniet",
                        "asperiores": "facere",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "121aa6f1-e674-4bdb-84f1-5756082d68ea",
                SupplierName: codataccounting.String("architecto"),
            },
            TotalAmount: 6091.78,
            TotalDiscount: 9453.02,
            TotalTaxAmount: 984.78,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 920.27,
                    Name: "Mrs. Cynthia Hansen",
                },
                shared.WithholdingTaxitems{
                    Amount: 6144.65,
                    Name: "Ms. Kenneth Ledner",
                },
                shared.WithholdingTaxitems{
                    Amount: 6498.32,
                    Name: "Mrs. Priscilla Fritsch",
                },
                shared.WithholdingTaxitems{
                    Amount: 2531.91,
                    Name: "Ms. Benjamin Hirthe DVM",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(618480),
    }

    res, err := s.BillCreditNotes.CreateBillCreditNote(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```