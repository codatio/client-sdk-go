# Bills

## Overview

Bills

### Available Operations

* [CreateBill](#createbill) - Create bill
* [DeleteBill](#deletebill) - Delete bill
* [DownloadBillAttachment](#downloadbillattachment) - Download bill attachment
* [GetBill](#getbill) - Get bill
* [GetBillAttachment](#getbillattachment) - Get bill attachment
* [GetBillAttachments](#getbillattachments) - List bill attachments
* [GetCreateUpdateBillsModel](#getcreateupdatebillsmodel) - Get create/update bill model
* [ListBills](#listbills) - List bills
* [UpdateBill](#updatebill) - Update bill
* [UploadBillAttachments](#uploadbillattachments) - Upload bill attachments

## CreateBill

Posts a new bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating a bill.

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
    req := operations.CreateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(941.22),
            Currency: codataccounting.String("rem"),
            CurrencyRate: codataccounting.Float64(4935.79),
            DueDate: codataccounting.String("doloremque"),
            ID: codataccounting.String("d9d21f9a-d030-4c4e-8c11-a0836429068b"),
            IssueDate: "laudantium",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("02a55e7f-73bc-4845-a320-a319f4badf94"),
                        Name: codataccounting.String("Pat Mueller"),
                    },
                    Description: codataccounting.String("aliquid"),
                    DiscountAmount: codataccounting.Float64(4693.84),
                    DiscountPercentage: codataccounting.Float64(7063.28),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "c4242666-5816-4ddc-a8ef-51fcb4c593ec",
                        Name: codataccounting.String("Beverly Satterfield"),
                    },
                    Quantity: 6615.78,
                    SubTotal: codataccounting.Float64(8409.92),
                    TaxAmount: codataccounting.Float64(590.23),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8792.08),
                        ID: codataccounting.String("c7afedbd-80df-4448-a47f-9390c5888098"),
                        Name: codataccounting.String("Eula Paucek"),
                    },
                    TotalAmount: codataccounting.Float64(6229.68),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "f3ffdd9f-7f07-49af-8d35-724cdb0f4d28",
                                Name: codataccounting.String("Alice Langosh"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "56844ede-d85a-4906-9e62-8bdfc2032b6c",
                                Name: codataccounting.String("Cody Mohr"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "3b7e1358-4f7a-4e12-8689-1f82ce115717",
                                Name: codataccounting.String("Dawn Bechtelar"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "77dcfa89-df97-45e3-9668-6092e9c3ddc5",
                                Name: codataccounting.String("Dr. Jack Buckridge"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("laborum"),
                            ID: "1026d541-a4d1-490f-ab21-780bccc0dbbd",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "484708fb-4e39-41e6-bc15-8c4c4e54599e",
                            Name: codataccounting.String("Jeffery Glover"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "0e9b200c-e78a-41bd-8fb7-a0a116ce723d",
                            Name: codataccounting.String("Susan Mraz"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a30e9af7-25b2-4912-a030-d83f5aeb7799",
                            Name: codataccounting.String("Nicholas Conroy"),
                        },
                    },
                    UnitAmount: 7607.93,
                },
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("1f849382-5fdc-442c-876c-2c2dfb4cfc1c"),
                        Name: codataccounting.String("Sue Christiansen DVM"),
                    },
                    Description: codataccounting.String("quos"),
                    DiscountAmount: codataccounting.Float64(2563.1),
                    DiscountPercentage: codataccounting.Float64(1138.94),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "fb1bd23f-db14-4db6-be5a-685998e22ae2",
                        Name: codataccounting.String("Ms. Lora Olson"),
                    },
                    Quantity: 7912.82,
                    SubTotal: codataccounting.Float64(1489.58),
                    TaxAmount: codataccounting.Float64(7211.98),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1554.73),
                        ID: codataccounting.String("71a289c5-7e85-44e9-8439-d22246569462"),
                        Name: codataccounting.String("Ms. Sharon Kuhlman"),
                    },
                    TotalAmount: codataccounting.Float64(9624.11),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "ab37cef0-2225-4194-9b55-410adc669af9",
                                Name: codataccounting.String("Sandy Considine"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "7cdc981f-0689-481d-abb3-3cfaa348c31b",
                                Name: codataccounting.String("Oscar Beatty"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("necessitatibus"),
                            ID: "4fcf0c42-b78f-4156-a639-8a0dc766324c",
                        },
                        IsBilledTo: shared.BilledToTypeEnumProject,
                        IsRebilledTo: shared.BilledToTypeEnumCustomer,
                        ProjectRef: &shared.ProjectRef{
                            ID: "06c8ca12-d025-4292-b0b8-d5722dd895b8",
                            Name: codataccounting.String("Guadalupe Wisoky"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "b9596933-52f7-4453-b994-d78de3b6e938",
                            Name: codataccounting.String("Courtney Hermann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b7f66255-0a28-4382-ac48-3afd2315bba6",
                            Name: codataccounting.String("Helen Blick"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e06f5bf6-ae59-41bc-8bde-f3612b63c205",
                            Name: codataccounting.String("Bryant Ondricka"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0774a68a-9a35-4d08-ab6f-66fef020e9f4",
                            Name: codataccounting.String("Ethel Robel"),
                        },
                    },
                    UnitAmount: 3184.03,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nihil"),
            Note: codataccounting.String("libero"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("excepturi"),
                        Currency: codataccounting.String("eos"),
                        CurrencyRate: codataccounting.Float64(7890.36),
                        TotalAmount: codataccounting.Float64(5471.84),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("dbda6a61-efa2-4198-a58f-d0a9eba47f7d"),
                            Name: codataccounting.String("Mrs. Tasha Wilkinson"),
                        },
                        Currency: codataccounting.String("aliquid"),
                        CurrencyRate: codataccounting.Float64(2946.5),
                        ID: codataccounting.String("0d6a1831-c87a-4df5-96fd-f1ad837ae80c"),
                        Note: codataccounting.String("vitae"),
                        PaidOnDate: codataccounting.String("cumque"),
                        Reference: codataccounting.String("architecto"),
                        TotalAmount: codataccounting.Float64(5754.04),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("eligendi"),
                        Currency: codataccounting.String("occaecati"),
                        CurrencyRate: codataccounting.Float64(3396.51),
                        TotalAmount: codataccounting.Float64(7343.61),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("a998678f-a3f6-4969-91af-388ce0361444"),
                            Name: codataccounting.String("Sylvester Kling"),
                        },
                        Currency: codataccounting.String("reprehenderit"),
                        CurrencyRate: codataccounting.Float64(6541.99),
                        ID: codataccounting.String("0ef2f536-028e-4fee-b934-152ed7e253f4"),
                        Note: codataccounting.String("quod"),
                        PaidOnDate: codataccounting.String("sunt"),
                        Reference: codataccounting.String("ullam"),
                        TotalAmount: codataccounting.Float64(4630.44),
                    },
                },
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("illum"),
                        Currency: codataccounting.String("voluptates"),
                        CurrencyRate: codataccounting.Float64(6410.06),
                        TotalAmount: codataccounting.Float64(6687.83),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7170f445-accf-4667-aaf9-bbad185fe431"),
                            Name: codataccounting.String("Rick Predovic"),
                        },
                        Currency: codataccounting.String("cumque"),
                        CurrencyRate: codataccounting.Float64(5291.74),
                        ID: codataccounting.String("38fbb8c2-0cb6-47fc-8b42-5e99e6234c9f"),
                        Note: codataccounting.String("voluptate"),
                        PaidOnDate: codataccounting.String("tempore"),
                        Reference: codataccounting.String("in"),
                        TotalAmount: codataccounting.Float64(6098.64),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("feb77a5c-38d4-4baf-91e5-06ef890a54b4"),
                    PurchaseOrderNumber: codataccounting.String("odio"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("5f16f56d-385a-43c4-ac63-1b99e26ced8f"),
                    PurchaseOrderNumber: codataccounting.String("natus"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("fdb9410f-63bb-4f81-b837-b01afdd78862"),
                    PurchaseOrderNumber: codataccounting.String("labore"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("189eb448-73f5-4033-b19d-bf125ce4152e"),
                    PurchaseOrderNumber: codataccounting.String("error"),
                },
            },
            Reference: codataccounting.String("expedita"),
            SourceModifiedDate: codataccounting.String("error"),
            Status: shared.BillStatusEnumVoid,
            SubTotal: 8382.27,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "earum": map[string]interface{}{
                        "odit": "odit",
                        "eius": "error",
                    },
                    "vel": map[string]interface{}{
                        "alias": "itaque",
                        "ab": "sunt",
                        "amet": "cum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "7847ec59-e1f6-47f3-84cc-e4b6d7696ff3",
                SupplierName: codataccounting.String("eligendi"),
            },
            TaxAmount: 3534.24,
            TotalAmount: 4713.15,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 4895.97,
                    Name: "Margaret Bradtke",
                },
                shared.BillWithholdingTax{
                    Amount: 4713.02,
                    Name: "Micheal Gusikowski",
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(115522),
    }

    res, err := s.Bills.CreateBill(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillResponse != nil {
        // handle response
    }
}
```

## DeleteBill

Deletes a bill from the accounting package for a given company.

> **Supported Integrations**
> 
> This functionality is currently only supported for our Oracle NetSuite and QuickBooks Online integrations. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    req := operations.DeleteBillRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.DeleteBill(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## DownloadBillAttachment

Download bill attachment

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
    req := operations.DownloadBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.DownloadBillAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetBill

Get bill

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
    req := operations.GetBillRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Bills.GetBill(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bill != nil {
        // handle response
    }
}
```

## GetBillAttachment

Get bill attachment

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
    req := operations.GetBillAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.GetBillAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetBillAttachments

Get bill attachments

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
    req := operations.GetBillAttachmentsRequest{
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.GetBillAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## GetCreateUpdateBillsModel

Get create/update bill model.

 > **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support creating and updating a bill.

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
    req := operations.GetCreateUpdateBillsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.GetCreateUpdateBillsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## ListBills

Gets the latest bills for a company, with pagination

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
    req := operations.ListBillsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("hic"),
    }

    res, err := s.Bills.ListBills(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bills != nil {
        // handle response
    }
}
```

## UpdateBill

Posts an updated bill to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update bill model](https://docs.codat.io/accounting-api#/operations/get-create-update-bills-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills) for integrations that support updating a bill.

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
    req := operations.UpdateBillRequest{
        Bill: &shared.Bill{
            AmountDue: codataccounting.Float64(5055.58),
            Currency: codataccounting.String("libero"),
            CurrencyRate: codataccounting.Float64(99.12),
            DueDate: codataccounting.String("totam"),
            ID: codataccounting.String("4c3197e1-93a2-4454-a7f9-4874c2d5cc49"),
            IssueDate: "ducimus",
            LineItems: []shared.BillLineItem{
                shared.BillLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("233e66bd-8fe5-4d00-b979-ef2038732059"),
                        Name: codataccounting.String("Pat Schmitt Jr."),
                    },
                    Description: codataccounting.String("perspiciatis"),
                    DiscountAmount: codataccounting.Float64(4069.46),
                    DiscountPercentage: codataccounting.Float64(2622.31),
                    IsDirectCost: codataccounting.Bool(false),
                    ItemRef: &shared.ItemRef{
                        ID: "00313b3e-5044-4f65-be72-dc4077d0cc3f",
                        Name: codataccounting.String("Carol Lowe"),
                    },
                    Quantity: 7738.54,
                    SubTotal: codataccounting.Float64(1199.28),
                    TaxAmount: codataccounting.Float64(3588.61),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7720.62),
                        ID: codataccounting.String("eb4d6e1e-ae0f-475a-adf2-acab58b991c9"),
                        Name: codataccounting.String("Jeanette Schultz"),
                    },
                    TotalAmount: codataccounting.Float64(3130.99),
                    Tracking: &shared.Propertiestracking{
                        CategoryRefs: []shared.TrackingCategoryRef{
                            shared.TrackingCategoryRef{
                                ID: "9461e742-1cbe-46d9-902f-0ea930b69f7a",
                                Name: codataccounting.String("Carlos Wunsch DVM"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "88500904-9116-4082-8788-8ec66183bfe9",
                                Name: codataccounting.String("Marion Mills"),
                            },
                            shared.TrackingCategoryRef{
                                ID: "40ec16fa-f75b-40b5-b2a4-da37cbaaf445",
                                Name: codataccounting.String("Bernadette Hahn"),
                            },
                        },
                        CustomerRef: &shared.CustomerRef{
                            CompanyName: codataccounting.String("explicabo"),
                            ID: "c9b2ad32-dafe-481a-88f4-444573fecd47",
                        },
                        IsBilledTo: shared.BilledToTypeEnumUnknown,
                        IsRebilledTo: shared.BilledToTypeEnumNotApplicable,
                        ProjectRef: &shared.ProjectRef{
                            ID: "3f63c820-9379-4aa6-9cd5-fbcf79da18a7",
                            Name: codataccounting.String("Martin Daugherty"),
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "5894e686-1adb-455f-9e5d-751c9fe8f750",
                            Name: codataccounting.String("Susie Wolf"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3450841f-1764-4456-b79f-3fb27e21f862",
                            Name: codataccounting.String("Ida Kihn"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6fc6b9f5-87ce-4525-8676-41a8312e5047",
                            Name: codataccounting.String("Mario Runolfsson DDS"),
                        },
                    },
                    UnitAmount: 7678.8,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("distinctio"),
            Note: codataccounting.String("numquam"),
            PaymentAllocations: []shared.BillPaymentAllocation{
                shared.BillPaymentAllocation{
                    Allocation: shared.BillPaymentAllocationAllocation{
                        AllocatedOnDate: codataccounting.String("amet"),
                        Currency: codataccounting.String("culpa"),
                        CurrencyRate: codataccounting.Float64(7057.53),
                        TotalAmount: codataccounting.Float64(7925.55),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("dc91faab-dd88-4e71-b6c4-8252d7771e7f"),
                            Name: codataccounting.String("Thomas Kirlin Jr."),
                        },
                        Currency: codataccounting.String("iste"),
                        CurrencyRate: codataccounting.Float64(8814.13),
                        ID: codataccounting.String("f8d29de1-dd70-497b-9da0-8c57fa6c78a2"),
                        Note: codataccounting.String("architecto"),
                        PaidOnDate: codataccounting.String("ea"),
                        Reference: codataccounting.String("accusamus"),
                        TotalAmount: codataccounting.Float64(768.73),
                    },
                },
            },
            PurchaseOrderRefs: []shared.PurchaseOrderRef{
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("bafeca61-9149-4814-8b64-ff8ae170ef03"),
                    PurchaseOrderNumber: codataccounting.String("distinctio"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("5f37e4aa-8685-4559-a673-2aa5dcb6682c"),
                    PurchaseOrderNumber: codataccounting.String("expedita"),
                },
                shared.PurchaseOrderRef{
                    ID: codataccounting.String("70f8cfd5-fb6e-491b-9a9f-74846e2c3309"),
                    PurchaseOrderNumber: codataccounting.String("repellendus"),
                },
            },
            Reference: codataccounting.String("soluta"),
            SourceModifiedDate: codataccounting.String("aut"),
            Status: shared.BillStatusEnumPartiallyPaid,
            SubTotal: 2271.85,
            SupplementalData: &shared.BillSupplementalData{
                Content: map[string]map[string]interface{}{
                    "quibusdam": map[string]interface{}{
                        "voluptates": "nihil",
                        "ad": "eligendi",
                        "fuga": "consequatur",
                    },
                    "sit": map[string]interface{}{
                        "earum": "quis",
                        "dolorem": "omnis",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "2c11a25a-8bf9-42f9-b428-ad9a9f8bf822",
                SupplierName: codataccounting.String("illo"),
            },
            TaxAmount: 1037.6,
            TotalAmount: 1541.17,
            WithholdingTax: []shared.BillWithholdingTax{
                shared.BillWithholdingTax{
                    Amount: 2449.9,
                    Name: "Sherri Schuster",
                },
                shared.BillWithholdingTax{
                    Amount: 2326.02,
                    Name: "Lance Zieme",
                },
            },
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(453111),
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

## UploadBillAttachments

Upload bill attachments

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
    req := operations.UploadBillAttachmentsRequest{
        RequestBody: &operations.UploadBillAttachmentsRequestBody{
            Content: []byte("cupiditate"),
            RequestBody: "maxime",
        },
        BillID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Bills.UploadBillAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
