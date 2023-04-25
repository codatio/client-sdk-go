# CreateInvoice
Available in: `Invoices`

Posts a new invoice to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update invoice model](https://docs.codat.io/accounting-api#/operations/get-create-update-invoices-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices) for integrations that support creating invoices.

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
    req := operations.CreateInvoiceRequest{
        Invoice: &shared.Invoice{
            AdditionalTaxAmount: codataccounting.Float64(5328.66),
            AdditionalTaxPercentage: codataccounting.Float64(7276.55),
            AmountDue: 952.32,
            Currency: codataccounting.String("quidem"),
            CurrencyRate: codataccounting.Float64(7452.74),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("exercitationem"),
                ID: "5a292b0b-c3bb-4744-a64e-b1d03388b0d1",
            },
            DiscountPercentage: codataccounting.Float64(7272.56),
            DueDate: codataccounting.String("2022-10-23T00:00:00Z"),
            ID: codataccounting.String("b17afee7-4b6f-4eb9-857c-7edaf39d16fb"),
            InvoiceNumber: codataccounting.String("asperiores"),
            IssueDate: "2022-10-23T00:00:00Z",
            LineItems: []shared.InvoiceLineItem{
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6fd162b3-03e3-4023-b93e-34316cf55b43"),
                        Name: codataccounting.String("Edith Herman"),
                    },
                    Description: codataccounting.String("porro"),
                    DiscountAmount: codataccounting.Float64(7629.28),
                    DiscountPercentage: codataccounting.Float64(9383.57),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "1c204c4a-dcc9-4904-8519-5b8648cefa78",
                        Name: codataccounting.String("Raymond Von"),
                    },
                    Quantity: 2009.91,
                    SubTotal: codataccounting.Float64(7268.55),
                    TaxAmount: codataccounting.Float64(6231.48),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(200.26),
                        ID: codataccounting.String("1e0952bb-b4cb-4b19-b713-d95a4169c138"),
                        Name: codataccounting.String("Dr. Marilyn Kuhn V"),
                    },
                    TotalAmount: codataccounting.Float64(9110.26),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9e45118c-2cc5-47fb-960b-1a78ed29a9d4",
                                Name: codataccounting.String("Dexter O'Kon"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "658c2d4f-4c88-4be4-b278-fd9667e46c51",
                                Name: codataccounting.String("Ernest Wolf"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a58dcef2-34c9-455b-9bdf-2190abd9bbcc",
                                Name: codataccounting.String("Stella Cole"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("impedit"),
                            ID: "2659ce02-8084-40c6-9ef6-8e45c8addfac",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1NotApplicable,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "4500430c-6632-4b43-91fd-f01c3e91e8f7",
                            Name: codataccounting.String("Jermaine Jenkins"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "60a77ece-b26d-410f-9ef2-631c7c0f0f87",
                            Name: codataccounting.String("Dixie Mueller"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c25fd3e0-b4a4-4a42-93c3-025711f42c7e",
                            Name: codataccounting.String("Mercedes Sauer"),
                        },
                    },
                    UnitAmount: 5293.03,
                },
                shared.InvoiceLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("be09e41a-7a21-45ca-92a4-ba9d59988192"),
                        Name: codataccounting.String("Miss Darrin Shanahan"),
                    },
                    Description: codataccounting.String("iusto"),
                    DiscountAmount: codataccounting.Float64(7563.43),
                    DiscountPercentage: codataccounting.Float64(3550.92),
                    IsDirectIncome: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "3e7e7d4e-e6e8-4b90-bac3-84e2396703fe",
                        Name: codataccounting.String("Jeff Casper"),
                    },
                    Quantity: 246.9,
                    SubTotal: codataccounting.Float64(5343.17),
                    TaxAmount: codataccounting.Float64(1834.8),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2965.94),
                        ID: codataccounting.String("d189a36a-6b2d-427e-b707-aa60c8fe46e6"),
                        Name: codataccounting.String("Colleen Klein"),
                    },
                    TotalAmount: codataccounting.Float64(6100.01),
                    Tracking: &shared.Propertiestracking1{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "b3b70ffb-b697-40ee-b70e-36097ef7c206",
                                Name: codataccounting.String("Ricardo Brown PhD"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "308714c2-0a3d-4986-b7ca-85c3fe65574d",
                                Name: codataccounting.String("Shannon Witting"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "a7c98f13-af28-4db2-8f2b-f4f3ded356d7",
                                Name: codataccounting.String("Lawrence Gleichner"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "1cd98196-d55a-4f69-a1c4-b79ae33681c2",
                                Name: codataccounting.String("Marianne Effertz"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("voluptate"),
                            ID: "c0e17cb1-2c5b-4a82-9fe2-2cd5cba6fbfe",
                        },
                        IsBilledTo: shared.BilledToTypeEnum1Project,
                        IsRebilledTo: shared.BilledToTypeEnum1NotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "32af6813-d65b-4fec-ac2d-d6916f7fc7dd",
                            Name: codataccounting.String("Clinton Beer"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0e607589-4d61-4c14-8d90-227e37c0d977",
                            Name: codataccounting.String("Douglas Nicolas"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "91abe975-1b10-46d2-be03-e69815aae99f",
                            Name: codataccounting.String("Levi Tremblay"),
                        },
                    },
                    UnitAmount: 4467.12,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Note: codataccounting.String("odit"),
            PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("eligendi"),
                        CurrencyRate: codataccounting.Float64(6230.66),
                        TotalAmount: codataccounting.Float64(8582.27),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4f2d8a44-640c-4a60-9b73-a2f93f467dc0"),
                            Name: codataccounting.String("Max Schuppe"),
                        },
                        Currency: codataccounting.String("voluptas"),
                        CurrencyRate: codataccounting.Float64(885.63),
                        ID: codataccounting.String("22026ab8-f277-4485-8197-6af980da7a08"),
                        Note: codataccounting.String("unde"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("hic"),
                        TotalAmount: codataccounting.Float64(8024.26),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("numquam"),
                        CurrencyRate: codataccounting.Float64(2838.48),
                        TotalAmount: codataccounting.Float64(8519.16),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b274530e-5cc7-4c6d-8cbc-fdcd334b6f62"),
                            Name: codataccounting.String("Whitney Schiller"),
                        },
                        Currency: codataccounting.String("mollitia"),
                        CurrencyRate: codataccounting.Float64(7064.71),
                        ID: codataccounting.String("50aee5e0-da8b-49af-aad0-5486e7b413cb"),
                        Note: codataccounting.String("saepe"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("magni"),
                        TotalAmount: codataccounting.Float64(8253.89),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(4767.65),
                        TotalAmount: codataccounting.Float64(4100.45),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("dc1c43d4-0f61-4d17-9157-cbe5ee4f7211"),
                            Name: codataccounting.String("Miguel Adams"),
                        },
                        Currency: codataccounting.String("quia"),
                        CurrencyRate: codataccounting.Float64(9804.11),
                        ID: codataccounting.String("32e3b49d-be0f-423b-bb6d-9948d6eded47"),
                        Note: codataccounting.String("odio"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00Z"),
                        Reference: codataccounting.String("voluptas"),
                        TotalAmount: codataccounting.Float64(5378.51),
                    },
                },
            },
            SalesOrderRefs: []string{
                "reiciendis",
            },
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.InvoiceStatusEnumPaid,
            SubTotal: codataccounting.Float64(4379.2),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ab": map[string]interface{}{
                        "deserunt": "blanditiis",
                        "dolores": "necessitatibus",
                    },
                    "nemo": map[string]interface{}{
                        "totam": "eos",
                        "delectus": "illum",
                        "consequuntur": "praesentium",
                        "fugiat": "beatae",
                    },
                    "perferendis": map[string]interface{}{
                        "aperiam": "harum",
                        "iusto": "debitis",
                    },
                },
            },
            TotalAmount: 5723.68,
            TotalDiscount: codataccounting.Float64(1083.49),
            TotalTaxAmount: 1890.08,
            WithholdingTax: []shared.WithholdingTaxitems{
                shared.WithholdingTaxitems{
                    Amount: 1850.41,
                    Name: "Preston Gorczany",
                },
                shared.WithholdingTaxitems{
                    Amount: 7102.56,
                    Name: "Lena Fisher Jr.",
                },
                shared.WithholdingTaxitems{
                    Amount: 5027.38,
                    Name: "Miss Ray Hoeger",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(336692),
    }

    res, err := s.Invoices.CreateInvoice(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateInvoiceResponse != nil {
        // handle response
    }
}
```