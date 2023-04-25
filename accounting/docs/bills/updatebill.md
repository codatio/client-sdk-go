# UpdateBill
Available in: `Bills`

Posts an updated bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support updating a bill.

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
    req := operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(3642.84),
            Currency: codataccounting.String("delectus"),
            CurrencyRate: codataccounting.Float64(8484.39),
            DueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("c42c876c-2c2d-4fb4-8fc1-c76230f841fb"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("bd23fdb1-4db6-4be5-a685-998e22ae20da"),
                        Name: codataccounting.String("Lucy Wilkinson"),
                    },
                    Description: codataccounting.String("nam"),
                    DiscountAmount: codataccounting.Float64(1554.73),
                    DiscountPercentage: codataccounting.Float64(4801.08),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1a289c57-e854-4e90-839d-222465694624",
                        Name: codataccounting.String("Viola Bailey"),
                    },
                    Quantity: 9624.11,
                    SubTotal: codataccounting.Float64(4618.53),
                    TaxAmount: codataccounting.Float64(6759.55),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7262.44),
                        ID: codataccounting.String("37cef022-2519-44db-9541-0adc669af90a"),
                        Name: codataccounting.String("Stacey Satterfield"),
                    },
                    TotalAmount: codataccounting.Float64(8149.5),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "981f0689-81d6-4bb3-bcfa-a348c31bf407",
                                Name: codataccounting.String("Santiago Fritsch"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "f0c42b78-f156-4263-98a0-dc766324ccb0",
                                Name: codataccounting.String("Leticia Leannon"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "12d02529-270b-48d5-b22d-d895b8bcf24d",
                                Name: codataccounting.String("Kent Hickle"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "93352f74-5339-494d-b8de-3b6e9389f5ab",
                                Name: codataccounting.String("Harvey Wisoky"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("fugit"),
                            ID: "550a2838-2ac4-483a-bd23-15bba650164e",
                        },
                        IsBilledTo: shared.BilledToTypeEnumUnknown,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "f5bf6ae5-91bc-48bd-af36-12b63c205fda",
                            Name: codataccounting.String("Alex Bayer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a68a9a35-d086-4b6f-a6fe-f020e9f443b4",
                            Name: codataccounting.String("Erin Kozey"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "92c8dbda-6a61-4efa-a198-258fd0a9eba4",
                            Name: codataccounting.String("Shari Konopelski"),
                        },
                    },
                    UnitAmount: 9156.53,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("a"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("quaerat"),
                        CurrencyRate: codataccounting.Float64(5698.72),
                        TotalAmount: codataccounting.Float64(3999.46),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("40d6a183-1c87-4adf-996f-df1ad837ae80"),
                            Name: codataccounting.String("Ms. Terry Runolfsson"),
                        },
                        Currency: codataccounting.String("occaecati"),
                        CurrencyRate: codataccounting.Float64(3396.51),
                        ID: codataccounting.String("ba998678-fa3f-4696-991a-f388ce036144"),
                        Note: codataccounting.String("modi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("quos"),
                        TotalAmount: codataccounting.Float64(7914.54),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("977a0ef2-f536-4028-afee-f934152ed7e2"),
                    PurchaseOrderNumber: codataccounting.String("ullam"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("3f4c157d-eaa7-4170-b445-accf667aaf9b"),
                    PurchaseOrderNumber: codataccounting.String("tempore"),
                },
            },
            Reference: codataccounting.String("culpa"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.BillStatusEnumDraft,
            SubTotal: 780.74,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "ad": map[string]interface{}{
                        "voluptates": "ut",
                        "nesciunt": "ab",
                        "quibusdam": "suscipit",
                        "quidem": "delectus",
                    },
                    "nemo": map[string]interface{}{
                        "voluptatum": "sequi",
                        "atque": "maiores",
                        "expedita": "rerum",
                        "totam": "quod",
                    },
                    "aspernatur": map[string]interface{}{
                        "impedit": "nam",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "67fc4b42-5e99-4e62-b4c9-f7b79dfeb77a",
                SupplierName: codataccounting.String("ad"),
            },
            TaxAmount: 8010.59,
            TotalAmount: 1887.05,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 8736.81,
                    Name: "Candice Nienow",
                },
                shared.BillWithholdingTax{
                    Amount: 752.03,
                    Name: "Ronnie Bechtelar",
                },
                shared.BillWithholdingTax{
                    Amount: 9625.43,
                    Name: "Eduardo Armstrong",
                },
            },
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(270862),
    }

    res, err := s.Bills.UpdateBill(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateBillResponse != nil {
        // handle response
    }
}
```