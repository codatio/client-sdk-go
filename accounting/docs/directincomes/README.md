# DirectIncomes

## Overview

Direct incomes

### Available Operations

* [Create](#create) - Create direct income
* [DownloadAttachment](#downloadattachment) - Download direct income attachment
* [Get](#get) - Get direct income
* [GetAttachment](#getattachment) - Get direct income attachment
* [GetCreateModel](#getcreatemodel) - Get create direct income model
* [List](#list) - List direct incomes
* [ListAttachments](#listattachments) - List direct income attachments
* [UploadAttachment](#uploadattachment) - Create direct income attachment

## Create

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
    res, err := s.DirectIncomes.Create(ctx, operations.CreateDirectIncomeRequest{
        DirectIncome: &shared.DirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("amet"),
                ID: "d5c72795-b785-4148-96d5-49e5635b33bc",
            },
            Currency: "voluptatem",
            CurrencyRate: codataccounting.Float64(9791.61),
            ID: codataccounting.String("970c42fc-9f48-4442-a5e7-5b796065c0ef"),
            IssueDate: "culpa",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f93b90a1-b8c9-45be-9254-b739f4fe7721"),
                        Name: codataccounting.String("Muriel Carroll"),
                    },
                    Description: codataccounting.String("exercitationem"),
                    DiscountAmount: codataccounting.Float64(3381.01),
                    DiscountPercentage: codataccounting.Float64(5546.44),
                    ItemRef: &shared.ItemRef{
                        ID: "c99c722d-2bc0-4f94-887d-9caae042dd7c",
                        Name: codataccounting.String("Carlton Schowalter"),
                    },
                    Quantity: 2667.52,
                    SubTotal: codataccounting.Float64(7955.01),
                    TaxAmount: codataccounting.Float64(6633.35),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6360.7),
                        ID: codataccounting.String("1cfe9e15-df90-4390-bf37-831983d42e54"),
                        Name: codataccounting.String("Ken Herzog"),
                    },
                    TotalAmount: codataccounting.Float64(4168.84),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "97c50233-c147-41d5-9aaa-6ddf5abd6487",
                            Name: codataccounting.String("Bill Wisoky"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b862a00b-ef69-4e10-8157-630bda7afded",
                            Name: codataccounting.String("Eddie Murazik"),
                        },
                    },
                    UnitAmount: 6774.73,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("41238e1a-735a-4c26-ae33-bef971a8f46b"),
                        Name: codataccounting.String("Mr. Angelo Brakus"),
                    },
                    Description: codataccounting.String("delectus"),
                    DiscountAmount: codataccounting.Float64(8871.37),
                    DiscountPercentage: codataccounting.Float64(6076.72),
                    ItemRef: &shared.ItemRef{
                        ID: "65b711d0-8cf8-48ec-9f7b-99a550a656ed",
                        Name: codataccounting.String("Paula Frami"),
                    },
                    Quantity: 375.63,
                    SubTotal: codataccounting.Float64(7773.99),
                    TaxAmount: codataccounting.Float64(8863.66),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5147.78),
                        ID: codataccounting.String("aa65432a-986e-4b7e-94ca-564089150097"),
                        Name: codataccounting.String("Catherine Mitchell"),
                    },
                    TotalAmount: codataccounting.Float64(5264.96),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "88ece7bf-904e-4011-85d3-8908162c6beb",
                            Name: codataccounting.String("Dr. Mattie Nader"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "57b7d03a-1480-4f8d-a30f-069d810618d9",
                            Name: codataccounting.String("Alyssa Casper"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "297510da-8031-4229-acc6-1c2a702bb97e",
                            Name: codataccounting.String("Carl Batz"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a2de35f8-e01b-4f33-aaab-45402ac1704b",
                            Name: codataccounting.String("Justin Schmitt"),
                        },
                    },
                    UnitAmount: 9890.79,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("maxime"),
            Note: codataccounting.String("ex"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("deserunt"),
                        Currency: codataccounting.String("laborum"),
                        CurrencyRate: codataccounting.Float64(9299.41),
                        TotalAmount: codataccounting.Float64(3241.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("eb5f0c49-2b57-444d-88a2-267aaee79e3c"),
                            Name: codataccounting.String("Martha Orn"),
                        },
                        Currency: codataccounting.String("et"),
                        CurrencyRate: codataccounting.Float64(7321.72),
                        ID: codataccounting.String("ecb83d23-78ae-43bf-823d-9450a986a495"),
                        Note: codataccounting.String("cum"),
                        PaidOnDate: codataccounting.String("dolorum"),
                        Reference: codataccounting.String("quod"),
                        TotalAmount: codataccounting.Float64(4715.35),
                    },
                },
            },
            Reference: codataccounting.String("quae"),
            SourceModifiedDate: codataccounting.String("ducimus"),
            SubTotal: 9483.77,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ex": map[string]interface{}{
                        "magni": "laudantium",
                        "repudiandae": "minus",
                        "porro": "atque",
                    },
                },
            },
            TaxAmount: 4203.54,
            TotalAmount: 2588.07,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(599915),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

## DownloadAttachment

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
    res, err := s.DirectIncomes.DownloadAttachment(ctx, operations.DownloadDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Get

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
    res, err := s.DirectIncomes.Get(ctx, operations.GetDirectIncomeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "sunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncome != nil {
        // handle response
    }
}
```

## GetAttachment

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
    res, err := s.DirectIncomes.GetAttachment(ctx, operations.GetDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(226197),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCreateModel

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
    res, err := s.DirectIncomes.GetCreateModel(ctx, operations.GetCreateDirectIncomesModelRequest{
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

Lists the direct incomes for a given company.

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
    res, err := s.DirectIncomes.List(ctx, operations.ListDirectIncomesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("laudantium"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncomes != nil {
        // handle response
    }
}
```

## ListAttachments

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
    res, err := s.DirectIncomes.ListAttachments(ctx, operations.ListDirectIncomeAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## UploadAttachment

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
    res, err := s.DirectIncomes.UploadAttachment(ctx, operations.UploadDirectIncomeAttachmentRequest{
        RequestBody: &operations.UploadDirectIncomeAttachmentRequestBody{
            Content: []byte("commodi"),
            RequestBody: "a",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "aliquid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
