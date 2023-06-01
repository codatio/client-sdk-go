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
                DataType: codataccounting.String("veniam"),
                ID: "118c2cc5-7fbd-460b-9a78-ed29a9d4eea8",
            },
            Currency: "ullam",
            CurrencyRate: codataccounting.Float64(3917.15),
            ID: codataccounting.String("58c2d4f4-c88b-4e4f-a78f-d9667e46c51d"),
            IssueDate: "dolores",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("faa58dce-f234-4c95-9b9b-df2190abd9bb"),
                        Name: codataccounting.String("Delbert Cummerata"),
                    },
                    Description: codataccounting.String("ullam"),
                    DiscountAmount: codataccounting.Float64(9045.07),
                    DiscountPercentage: codataccounting.Float64(7703.76),
                    ItemRef: &shared.ItemRef{
                        ID: "2659ce02-8084-40c6-9ef6-8e45c8addfac",
                        Name: codataccounting.String("Bernice Gottlieb Jr."),
                    },
                    Quantity: 3079.36,
                    SubTotal: codataccounting.Float64(2029.21),
                    TaxAmount: codataccounting.Float64(462.58),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7597.31),
                        ID: codataccounting.String("6632b439-1fdf-401c-be91-e8f7bc69d460"),
                        Name: codataccounting.String("Tyrone Kuhn"),
                    },
                    TotalAmount: codataccounting.Float64(9053.64),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "26d10f1e-f263-41c7-80f0-f873f9d5c25f",
                            Name: codataccounting.String("Miss Alfred VonRueden"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a4a4253c-3025-4711-b42c-7e7dc548be09",
                            Name: codataccounting.String("Bradley Boyle"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "a215ca12-a4ba-49d5-9988-192cfd0c77c5",
                            Name: codataccounting.String("Meghan Koelpin"),
                        },
                    },
                    UnitAmount: 8136.42,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4ee6e8b9-0bac-4384-a239-6703fec31c50"),
                        Name: codataccounting.String("Shawn Gulgowski IV"),
                    },
                    Description: codataccounting.String("perspiciatis"),
                    DiscountAmount: codataccounting.Float64(6461.08),
                    DiscountPercentage: codataccounting.Float64(2236.36),
                    ItemRef: &shared.ItemRef{
                        ID: "6a6b2d27-eb70-47aa-a0c8-fe46e6177db9",
                        Name: codataccounting.String("Dominic Ernser"),
                    },
                    Quantity: 508.59,
                    SubTotal: codataccounting.Float64(9759.09),
                    TaxAmount: codataccounting.Float64(9767),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7372.75),
                        ID: codataccounting.String("b6970ee7-70e3-4609-bef7-c206e61b0d30"),
                        Name: codataccounting.String("Kelly Blick"),
                    },
                    TotalAmount: codataccounting.Float64(1644.32),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a3d98637-ca85-4c3f-a655-74dbaf94a7c9",
                            Name: codataccounting.String("Randal Blanda"),
                        },
                    },
                    UnitAmount: 9616.33,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("28db2cf2-bf4f-43de-9356-d7e14b21cd98"),
                        Name: codataccounting.String("Faye Jacobi"),
                    },
                    Description: codataccounting.String("quis"),
                    DiscountAmount: codataccounting.Float64(6606.91),
                    DiscountPercentage: codataccounting.Float64(9651.45),
                    ItemRef: &shared.ItemRef{
                        ID: "69a1c4b7-9ae3-4368-9c23-c39a7c0e17cb",
                        Name: codataccounting.String("Jacqueline Russel"),
                    },
                    Quantity: 6321.21,
                    SubTotal: codataccounting.Float64(5137.34),
                    TaxAmount: codataccounting.Float64(1478.83),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3588.82),
                        ID: codataccounting.String("fe22cd5c-ba6f-4bfe-8932-af6813d65bfe"),
                        Name: codataccounting.String("Elbert Schmidt"),
                    },
                    TotalAmount: codataccounting.Float64(8717.9),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "916f7fc7-dda7-40ec-a0e6-075894d61c14",
                            Name: codataccounting.String("Mr. Irving Mann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7e37c0d9-77f1-4a54-91ab-e9751b106d23",
                            Name: codataccounting.String("Kenneth Effertz"),
                        },
                    },
                    UnitAmount: 6179.29,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("815aae99-fcde-49e7-a9c9-d4f2d8a44640"),
                        Name: codataccounting.String("Dr. Wilbur Jerde"),
                    },
                    Description: codataccounting.String("nihil"),
                    DiscountAmount: codataccounting.Float64(2079.6),
                    DiscountPercentage: codataccounting.Float64(6657.15),
                    ItemRef: &shared.ItemRef{
                        ID: "2f93f467-dc0d-48da-9612-2026ab8f2774",
                        Name: codataccounting.String("Ms. Greg Sanford"),
                    },
                    Quantity: 3771.77,
                    SubTotal: codataccounting.Float64(6271.29),
                    TaxAmount: codataccounting.Float64(9494),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6235.78),
                        ID: codataccounting.String("80da7a08-9fc4-44db-a745-30e5cc7c6d0c"),
                        Name: codataccounting.String("Spencer Wintheiser"),
                    },
                    TotalAmount: codataccounting.Float64(8196.32),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "34b6f623-bcec-4ab5-8aee-5e0da8b9af6a",
                            Name: codataccounting.String("Christopher Hermiston"),
                        },
                    },
                    UnitAmount: 3905.07,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("repudiandae"),
            Note: codataccounting.String("odio"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("aliquam"),
                        Currency: codataccounting.String("quasi"),
                        CurrencyRate: codataccounting.Float64(2144.25),
                        TotalAmount: codataccounting.Float64(7771.39),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("be2d176d-c1c4-43d4-8f61-d171157cbe5e"),
                            Name: codataccounting.String("Francis Weimann"),
                        },
                        Currency: codataccounting.String("quasi"),
                        CurrencyRate: codataccounting.Float64(933.86),
                        ID: codataccounting.String("840772f3-2e3b-449d-be0f-23b7b6d9948d"),
                        Note: codataccounting.String("aliquid"),
                        PaidOnDate: codataccounting.String("saepe"),
                        Reference: codataccounting.String("facere"),
                        TotalAmount: codataccounting.Float64(8999.7),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("fugiat"),
                        Currency: codataccounting.String("eius"),
                        CurrencyRate: codataccounting.Float64(4515.93),
                        TotalAmount: codataccounting.Float64(4844.89),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("680fc7a1-7a82-4e5e-82fd-28d1040a7e91"),
                            Name: codataccounting.String("Belinda Denesik"),
                        },
                        Currency: codataccounting.String("incidunt"),
                        CurrencyRate: codataccounting.Float64(3084.29),
                        ID: codataccounting.String("cb183500-8f46-41ce-93e9-14498a9ba460"),
                        Note: codataccounting.String("similique"),
                        PaidOnDate: codataccounting.String("nulla"),
                        Reference: codataccounting.String("pariatur"),
                        TotalAmount: codataccounting.Float64(9692.94),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("temporibus"),
                        Currency: codataccounting.String("officiis"),
                        CurrencyRate: codataccounting.Float64(2522.9),
                        TotalAmount: codataccounting.Float64(826.1),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0c37daa9-182a-449d-9625-d3caffc198ee"),
                            Name: codataccounting.String("Edwin Hagenes"),
                        },
                        Currency: codataccounting.String("dignissimos"),
                        CurrencyRate: codataccounting.Float64(6043.08),
                        ID: codataccounting.String("2bcd440e-a98b-4ecc-a048-6de0d56d73b0"),
                        Note: codataccounting.String("quae"),
                        PaidOnDate: codataccounting.String("quis"),
                        Reference: codataccounting.String("nemo"),
                        TotalAmount: codataccounting.Float64(345.89),
                    },
                },
            },
            Reference: codataccounting.String("neque"),
            SourceModifiedDate: codataccounting.String("voluptates"),
            SubTotal: 5187.95,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "impedit": map[string]interface{}{
                        "explicabo": "ea",
                        "doloribus": "maiores",
                    },
                    "nihil": map[string]interface{}{
                        "impedit": "iure",
                        "ullam": "aliquid",
                    },
                    "odio": map[string]interface{}{
                        "delectus": "nostrum",
                        "harum": "reprehenderit",
                    },
                    "sit": map[string]interface{}{
                        "consectetur": "vero",
                        "eius": "optio",
                        "sapiente": "porro",
                        "impedit": "vel",
                    },
                },
            },
            TaxAmount: 6786.31,
            TotalAmount: 5828.22,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(69650),
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
        DirectIncomeID: "necessitatibus",
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
        DirectIncomeID: "maxime",
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
        DirectIncomeID: "veniam",
        TimeoutInMinutes: codataccounting.Int(181673),
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
        Query: codataccounting.String("aliquid"),
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
        DirectIncomeID: "sed",
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
            Content: []byte("modi"),
            RequestBody: "at",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "aperiam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
