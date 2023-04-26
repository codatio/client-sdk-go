# DirectIncomes

## Overview

Direct incomes

### Available Operations

* [CreateDirectIncome](#createdirectincome) - Create direct income
* [DownloadDirectIncomeAttachment](#downloaddirectincomeattachment) - Download direct income attachment
* [GetCreateDirectIncomesModel](#getcreatedirectincomesmodel) - Get create direct income model
* [GetDirectIncome](#getdirectincome) - Get direct income
* [GetDirectIncomeAttachment](#getdirectincomeattachment) - Get direct income attachment
* [GetDirectIncomes](#getdirectincomes) - Get direct incomes
* [ListDirectIncomeAttachments](#listdirectincomeattachments) - List direct income attachments
* [UploadDirectIncomeAttachment](#uploaddirectincomeattachment) - Create direct income attachment

## CreateDirectIncome

Posts a new direct income to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating direct incomes.

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
    req := operations.CreateDirectIncomeRequest{
        DirectIncome: &shared.DirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("rem"),
                ID: "f652ebb9-d383-4838-b902-43b293dab30e",
            },
            Currency: "excepturi",
            CurrencyRate: codataccounting.Float64(1194.72),
            ID: codataccounting.String("7f50fda0-4c8b-41bb-95a2-92b0bc3bb744"),
            IssueDate: "laboriosam",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4eb1d033-88b0-4d1b-b17a-fee74b6feb94"),
                        Name: codataccounting.String("Colleen Schmidt"),
                    },
                    Description: codataccounting.String("illum"),
                    DiscountAmount: codataccounting.Float64(6734.93),
                    DiscountPercentage: codataccounting.Float64(9682.72),
                    ItemRef: &shared.ItemRef{
                        ID: "39d16fbf-76fd-4162-b303-e3023b93e343",
                        Name: codataccounting.String("Jeanne Schowalter"),
                    },
                    Quantity: 3667.8,
                    SubTotal: codataccounting.Float64(7436.12),
                    TaxAmount: codataccounting.Float64(2976.24),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1923.84),
                        ID: codataccounting.String("13553ccf-1c20-44c4-adcc-9904c5195b86"),
                        Name: codataccounting.String("Irma Rowe"),
                    },
                    TotalAmount: codataccounting.Float64(6497.91),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "8f1e2d3b-901e-4095-abbb-4cbb19f713d9",
                            Name: codataccounting.String("Mrs. Kristi Greenholt"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c1387271-e18e-4a9e-8511-8c2cc57fbd60",
                            Name: codataccounting.String("Jack Paucek"),
                        },
                    },
                    UnitAmount: 9102.97,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("d29a9d4e-ea85-4658-82d4-f4c88be4f278"),
                        Name: codataccounting.String("Laurence Mitchell"),
                    },
                    Description: codataccounting.String("molestiae"),
                    DiscountAmount: codataccounting.Float64(8828.26),
                    DiscountPercentage: codataccounting.Float64(2892.08),
                    ItemRef: &shared.ItemRef{
                        ID: "6c51d2ff-aa58-4dce-b234-c955b9bdf219",
                        Name: codataccounting.String("Alberta Rogahn"),
                    },
                    Quantity: 7129.08,
                    SubTotal: codataccounting.Float64(6908.74),
                    TaxAmount: codataccounting.Float64(7889.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7535.25),
                        ID: codataccounting.String("2725ec26-59ce-4028-8840-c69ef68e45c8"),
                        Name: codataccounting.String("Taylor Schuster"),
                    },
                    TotalAmount: codataccounting.Float64(7591.92),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "54500430-c663-42b4-b91f-df01c3e91e8f",
                            Name: codataccounting.String("Jeannette Schmeler"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d460a77e-ceb2-46d1-8f1e-f2631c7c0f0f",
                            Name: codataccounting.String("Allan Ebert"),
                        },
                    },
                    UnitAmount: 8159.77,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("veniam"),
            Note: codataccounting.String("eligendi"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("nemo"),
                        Currency: codataccounting.String("doloribus"),
                        CurrencyRate: codataccounting.Float64(8503.86),
                        TotalAmount: codataccounting.Float64(2499.41),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e0b4a4a4-253c-4302-9711-f42c7e7dc548"),
                            Name: codataccounting.String("Bradford Balistreri"),
                        },
                        Currency: codataccounting.String("numquam"),
                        CurrencyRate: codataccounting.Float64(989.55),
                        ID: codataccounting.String("a7a215ca-12a4-4ba9-9599-88192cfd0c77"),
                        Note: codataccounting.String("eligendi"),
                        PaidOnDate: codataccounting.String("ullam"),
                        Reference: codataccounting.String("dolorem"),
                        TotalAmount: codataccounting.Float64(8782.82),
                    },
                },
            },
            Reference: codataccounting.String("esse"),
            SourceModifiedDate: codataccounting.String("vero"),
            SubTotal: 4930.61,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "eius": map[string]interface{}{
                        "vero": "nisi",
                        "recusandae": "deleniti",
                        "nobis": "excepturi",
                        "consequatur": "distinctio",
                    },
                    "similique": map[string]interface{}{
                        "consectetur": "molestias",
                        "modi": "saepe",
                        "qui": "dolor",
                        "sint": "ea",
                    },
                    "in": map[string]interface{}{
                        "sequi": "maiores",
                    },
                    "itaque": map[string]interface{}{
                        "adipisci": "sunt",
                        "quo": "veniam",
                        "sit": "deleniti",
                        "qui": "dolore",
                    },
                },
            },
            TaxAmount: 8188.66,
            TotalAmount: 1179.02,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(536427),
    }

    res, err := s.DirectIncomes.CreateDirectIncome(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

## DownloadDirectIncomeAttachment

Downloads an attachment for the specified direct income for a given company.

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
    req := operations.DownloadDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectIncomes.DownloadDirectIncomeAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetCreateDirectIncomesModel

Get create direct income model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating direct incomes.

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
    req := operations.GetCreateDirectIncomesModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.DirectIncomes.GetCreateDirectIncomesModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetDirectIncome

Gets the specified direct income for a given company and connection.

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
    req := operations.GetDirectIncomeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "perspiciatis",
    }

    res, err := s.DirectIncomes.GetDirectIncome(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncome != nil {
        // handle response
    }
}
```

## GetDirectIncomeAttachment

Gets the specified direct income attachment for a given company.

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
    req := operations.GetDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(646108),
    }

    res, err := s.DirectIncomes.GetDirectIncomeAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetDirectIncomes

Gets the direct incomes for a given company.

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
    req := operations.GetDirectIncomesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("dolor"),
    }

    res, err := s.DirectIncomes.GetDirectIncomes(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncomes != nil {
        // handle response
    }
}
```

## ListDirectIncomeAttachments

Gets all attachments for the specified direct income for a given company.

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
    req := operations.ListDirectIncomeAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectIncomes.ListDirectIncomeAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## UploadDirectIncomeAttachment

Posts a new direct income attachment for a given company.

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
    req := operations.UploadDirectIncomeAttachmentRequest{
        RequestBody: &operations.UploadDirectIncomeAttachmentRequestBody{
            Content: []byte("eum"),
            RequestBody: "culpa",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "iure",
    }

    res, err := s.DirectIncomes.UploadDirectIncomeAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
