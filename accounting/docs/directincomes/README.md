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

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) to see which integrations support this endpoint.

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
    res, err := s.DirectIncomes.Create(ctx, operations.CreateDirectIncomeRequest{
        DirectIncome: &shared.DirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("magni"),
                ID: "cdd14fc4-3b70-4bca-88fa-70c43351c3dd",
            },
            Currency: "beatae",
            CurrencyRate: codataccounting.Float64(8962.8),
            ID: codataccounting.String("b8f7f75f-4f23-4f1c-8a58-6c3ae7d7b67f"),
            IssueDate: "officiis",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f5e142d9-5b1d-4bec-aff7-c4b156e92782"),
                        Name: codataccounting.String("Holly Toy"),
                    },
                    Description: codataccounting.String("esse"),
                    DiscountAmount: codataccounting.Float64(3851.06),
                    DiscountPercentage: codataccounting.Float64(5122.23),
                    ItemRef: &shared.ItemRef{
                        ID: "17468063-f799-4b79-96c0-b0fa0bb20a40",
                        Name: codataccounting.String("Roland Ryan"),
                    },
                    Quantity: 8828.41,
                    SubTotal: codataccounting.Float64(3996.96),
                    TaxAmount: codataccounting.Float64(3081.27),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(498.27),
                        ID: codataccounting.String("64272657-b01a-407c-88fd-3921c257930d"),
                        Name: codataccounting.String("Elisa Bailey"),
                    },
                    TotalAmount: codataccounting.Float64(6488.15),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "efa46d36-6dfa-4101-9a09-1b3ec8b53862",
                            Name: codataccounting.String("Alonzo Bins"),
                        },
                    },
                    UnitAmount: 8706.71,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("14fe72e5-21f9-4030-bdfc-338397fffa6d"),
                        Name: codataccounting.String("Meredith Dickinson V"),
                    },
                    Description: codataccounting.String("alias"),
                    DiscountAmount: codataccounting.Float64(9869.41),
                    DiscountPercentage: codataccounting.Float64(7770.83),
                    ItemRef: &shared.ItemRef{
                        ID: "157ac9fe-1961-4ce9-be41-c869dd7d9719",
                        Name: codataccounting.String("Michael Kulas"),
                    },
                    Quantity: 154.46,
                    SubTotal: codataccounting.Float64(498.92),
                    TaxAmount: codataccounting.Float64(6470.89),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3619.89),
                        ID: codataccounting.String("8ffd2967-df8f-4d88-aa8e-60be620cd9c5"),
                        Name: codataccounting.String("Darrin Schuppe II"),
                    },
                    TotalAmount: codataccounting.Float64(7561.02),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "752512be-ae1d-487e-8c5f-dcea8e7a8831",
                            Name: codataccounting.String("Lucy Howell"),
                        },
                    },
                    UnitAmount: 8726.91,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a6d77c1d-8606-4623-bd42-27866db8a749"),
                        Name: codataccounting.String("Glenn McLaughlin"),
                    },
                    Description: codataccounting.String("exercitationem"),
                    DiscountAmount: codataccounting.Float64(988.25),
                    DiscountPercentage: codataccounting.Float64(1089.18),
                    ItemRef: &shared.ItemRef{
                        ID: "cc75e4f0-c004-4b5b-b758-cc94562f0069",
                        Name: codataccounting.String("Gwendolyn Hickle"),
                    },
                    Quantity: 8194.9,
                    SubTotal: codataccounting.Float64(955.55),
                    TaxAmount: codataccounting.Float64(6650.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1149.02),
                        ID: codataccounting.String("73d84bbe-24f2-4983-8afb-0735cb6285d4"),
                        Name: codataccounting.String("Fred Mitchell"),
                    },
                    TotalAmount: codataccounting.Float64(6662.15),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e169156f-7d2e-4e20-9505-bf03a93e9448",
                            Name: codataccounting.String("Myra Pfannerstill"),
                        },
                    },
                    UnitAmount: 9771.81,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b1078903-2ac3-4331-b2e2-dd79ec74ba7e"),
                        Name: codataccounting.String("Jaime Schuster"),
                    },
                    Description: codataccounting.String("ipsum"),
                    DiscountAmount: codataccounting.Float64(4175.44),
                    DiscountPercentage: codataccounting.Float64(9669.27),
                    ItemRef: &shared.ItemRef{
                        ID: "d1ccc341-c865-4734-b4f0-a740fb4ab441",
                        Name: codataccounting.String("Ms. Dale Nolan"),
                    },
                    Quantity: 4588.95,
                    SubTotal: codataccounting.Float64(4145.21),
                    TaxAmount: codataccounting.Float64(2428.43),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5966.4),
                        ID: codataccounting.String("95d808bb-e794-4455-abc5-50a1c426b59c"),
                        Name: codataccounting.String("Rodney Jerde"),
                    },
                    TotalAmount: codataccounting.Float64(8142.27),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "c135582c-1b85-45e8-89d9-ef932e9000a1",
                            Name: codataccounting.String("Sandy Stokes Sr."),
                        },
                        shared.TrackingCategoryRef{
                            ID: "4208efd2-3411-4898-a738-79efbe8baeba",
                            Name: codataccounting.String("Geoffrey Kuhic"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "536e9035-1bb9-4763-9720-b77a5a5365a7",
                            Name: codataccounting.String("Timmy Bernier"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "71f01c0d-361f-4ed8-9c5e-ffb453e9089e",
                            Name: codataccounting.String("Roland Bode"),
                        },
                    },
                    UnitAmount: 7471.07,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ut"),
            Note: codataccounting.String("at"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("perspiciatis"),
                        Currency: codataccounting.String("molestiae"),
                        CurrencyRate: codataccounting.Float64(7361.58),
                        TotalAmount: codataccounting.Float64(8169.35),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("d9c985e4-3734-4a5d-b2d9-edd785be5e7a"),
                            Name: codataccounting.String("Sheldon Hermann"),
                        },
                        Currency: codataccounting.String("occaecati"),
                        CurrencyRate: codataccounting.Float64(4720.94),
                        ID: codataccounting.String("ba6281f4-4e3a-4233-94a6-8cc80d30ff72"),
                        Note: codataccounting.String("illo"),
                        PaidOnDate: codataccounting.String("aliquid"),
                        Reference: codataccounting.String("quaerat"),
                        TotalAmount: codataccounting.Float64(8368.03),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("alias"),
                        Currency: codataccounting.String("deserunt"),
                        CurrencyRate: codataccounting.Float64(5794.14),
                        TotalAmount: codataccounting.Float64(1167.29),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fe9d9655-3b89-4e00-89c6-692de7b35622"),
                            Name: codataccounting.String("Christine Nikolaus"),
                        },
                        Currency: codataccounting.String("error"),
                        CurrencyRate: codataccounting.Float64(7417.47),
                        ID: codataccounting.String("4ae7b1a5-b908-4d4e-b049-1aa35d4a839f"),
                        Note: codataccounting.String("consequatur"),
                        PaidOnDate: codataccounting.String("ipsum"),
                        Reference: codataccounting.String("quidem"),
                        TotalAmount: codataccounting.Float64(6785.88),
                    },
                },
            },
            Reference: codataccounting.String("quidem"),
            SourceModifiedDate: codataccounting.String("molestiae"),
            SubTotal: 4410.01,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "occaecati": map[string]interface{}{
                        "blanditiis": "a",
                    },
                    "aut": map[string]interface{}{
                        "dicta": "dolor",
                    },
                    "iste": map[string]interface{}{
                        "ut": "exercitationem",
                        "sit": "reprehenderit",
                        "officiis": "accusantium",
                    },
                },
            },
            TaxAmount: 9159.68,
            TotalAmount: 2348.84,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(577369),
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectIncomes.Get(ctx, operations.GetDirectIncomeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "impedit",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectIncomes.GetAttachment(ctx, operations.GetDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(461855),
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
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating direct incomes.

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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectIncomes.List(ctx, operations.ListDirectIncomesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("saepe"),
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectIncomes.UploadAttachment(ctx, operations.UploadDirectIncomeAttachmentRequest{
        RequestBody: &operations.UploadDirectIncomeAttachmentRequestBody{
            Content: []byte("odit"),
            RequestBody: "consectetur",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "itaque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
