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
                DataType: codataccounting.String("placeat"),
                ID: "9b4caa1c-fe9e-415d-b903-907f37831983",
            },
            Currency: "fugiat",
            CurrencyRate: codataccounting.Float64(2628.24),
            ID: codataccounting.String("2e54a854-6659-47c5-8233-c1471d51aaa6"),
            IssueDate: "illum",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f5abd648-7c5f-4c2b-862a-00bef69e1001"),
                        Name: codataccounting.String("Jo Homenick DDS"),
                    },
                    Description: codataccounting.String("quibusdam"),
                    DiscountAmount: codataccounting.Float64(6828.88),
                    DiscountPercentage: codataccounting.Float64(4685.8),
                    ItemRef: &shared.ItemRef{
                        ID: "afded84a-35a4-4123-8e1a-735ac26ae33b",
                        Name: codataccounting.String("Cary Medhurst MD"),
                    },
                    Quantity: 5021.5,
                    SubTotal: codataccounting.Float64(9993.15),
                    TaxAmount: codataccounting.Float64(2679.18),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4014.66),
                        ID: codataccounting.String("bca1106f-e965-4b71-9d08-cf88ec9f7b99"),
                        Name: codataccounting.String("Miss Maurice Hauck"),
                    },
                    TotalAmount: codataccounting.Float64(3393.33),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ed333bb0-ce8a-4a65-832a-986eb7e14ca5",
                            Name: codataccounting.String("Ellen Baumbach"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "15009701-9a48-4f88-ace7-bf904e01105d",
                            Name: codataccounting.String("Ms. Marsha Marks III"),
                        },
                    },
                    UnitAmount: 1337.9,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c6beb68a-0f65-47b7-903a-1480f8de30f0"),
                        Name: codataccounting.String("Shelly Steuber Jr."),
                    },
                    Description: codataccounting.String("commodi"),
                    DiscountAmount: codataccounting.Float64(1030.68),
                    DiscountPercentage: codataccounting.Float64(5433.17),
                    ItemRef: &shared.ItemRef{
                        ID: "d97e1522-9751-40da-8031-2292cc61c2a7",
                        Name: codataccounting.String("Kathy Reichel"),
                    },
                    Quantity: 4877.46,
                    SubTotal: codataccounting.Float64(9182.86),
                    TaxAmount: codataccounting.Float64(8796.44),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(935.56),
                        ID: codataccounting.String("02da2de3-5f8e-401b-b33e-aab45402ac17"),
                        Name: codataccounting.String("Amber Rogahn DDS"),
                    },
                    TotalAmount: codataccounting.Float64(8111.67),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "fc61aae5-eb5f-40c4-92b5-744d08a2267a",
                            Name: codataccounting.String("Sheldon Toy"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e3c71ad3-1bec-4b83-9237-8ae3bfc23d94",
                            Name: codataccounting.String("Sarah Okuneva"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "6a495bac-707f-406b-a8ec-c86492386f62",
                            Name: codataccounting.String("Wade Kemmer"),
                        },
                    },
                    UnitAmount: 2765.07,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("cc6b7889-0a3f-4d3c-81da-10f8c23df931"),
                        Name: codataccounting.String("Homer Franecki"),
                    },
                    Description: codataccounting.String("tempore"),
                    DiscountAmount: codataccounting.Float64(3267.12),
                    DiscountPercentage: codataccounting.Float64(1152.02),
                    ItemRef: &shared.ItemRef{
                        ID: "fad94acc-9435-4137-b26d-15321b832a56",
                        Name: codataccounting.String("Ms. Brent Marquardt DVM"),
                    },
                    Quantity: 9463.17,
                    SubTotal: codataccounting.Float64(4373.51),
                    TaxAmount: codataccounting.Float64(84.69),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8836.41),
                        ID: codataccounting.String("b9a6658e-69a4-4b84-bd38-2dbec75c68c6"),
                        Name: codataccounting.String("Alma Hartmann"),
                    },
                    TotalAmount: codataccounting.Float64(3818.37),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "ce304d88-49bf-4821-8c33-7f96bb0c69e3",
                            Name: codataccounting.String("Theresa Schumm I"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "44ba9f78-a5c0-4ed7-aab6-2e97261fb0c5",
                            Name: codataccounting.String("Lionel Connelly"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "51996b5b-4b50-4eef-b12b-7a7ab0344b17",
                            Name: codataccounting.String("Nancy Johnson"),
                        },
                    },
                    UnitAmount: 8478.05,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("eebef897-f3dd-40cc-933f-11b3e4e080aa"),
                        Name: codataccounting.String("Ms. Cynthia Gorczany"),
                    },
                    Description: codataccounting.String("accusamus"),
                    DiscountAmount: codataccounting.Float64(7590.59),
                    DiscountPercentage: codataccounting.Float64(4841.4),
                    ItemRef: &shared.ItemRef{
                        ID: "59e02f37-02c5-4c8e-ad30-ead3104fa447",
                        Name: codataccounting.String("Tanya Prohaska"),
                    },
                    Quantity: 4960.06,
                    SubTotal: codataccounting.Float64(3249.55),
                    TaxAmount: codataccounting.Float64(7261.44),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2839.36),
                        ID: codataccounting.String("4282821f-db2f-469e-9926-7c71cc8d3cd4"),
                        Name: codataccounting.String("Dolores Lehner I"),
                    },
                    TotalAmount: codataccounting.Float64(3691.81),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a82c808f-e275-41a2-847c-0449e143f961",
                            Name: codataccounting.String("Kelvin Robel"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "40d5a11f-a436-4e62-9923-3f95c9d23739",
                            Name: codataccounting.String("Leticia Kris"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b5db4f50-0183-4feb-9f67-6b7206dab750",
                            Name: codataccounting.String("Jill Davis"),
                        },
                    },
                    UnitAmount: 3951.94,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("dolore"),
            Note: codataccounting.String("in"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("fugiat"),
                        Currency: codataccounting.String("minus"),
                        CurrencyRate: codataccounting.Float64(2693.91),
                        TotalAmount: codataccounting.Float64(1928.71),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("9ed8c432-0f41-4240-9448-7ac693b94c3b"),
                            Name: codataccounting.String("Ismael Cummings"),
                        },
                        Currency: codataccounting.String("blanditiis"),
                        CurrencyRate: codataccounting.Float64(8422.41),
                        ID: codataccounting.String("795aa42f-c405-4669-b69a-006d21249450"),
                        Note: codataccounting.String("atque"),
                        PaidOnDate: codataccounting.String("quasi"),
                        Reference: codataccounting.String("natus"),
                        TotalAmount: codataccounting.Float64(8518.84),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("odio"),
                        Currency: codataccounting.String("quo"),
                        CurrencyRate: codataccounting.Float64(2153.91),
                        TotalAmount: codataccounting.Float64(7379.94),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("1b418440-60e0-4031-8d02-3dc901f5afd2"),
                            Name: codataccounting.String("Rafael Sanford"),
                        },
                        Currency: codataccounting.String("praesentium"),
                        CurrencyRate: codataccounting.Float64(2573.24),
                        ID: codataccounting.String("6ae9d892-53c8-4962-b489-6bf51e4652d3"),
                        Note: codataccounting.String("cumque"),
                        PaidOnDate: codataccounting.String("nesciunt"),
                        Reference: codataccounting.String("aliquam"),
                        TotalAmount: codataccounting.Float64(2338.04),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("at"),
                        Currency: codataccounting.String("suscipit"),
                        CurrencyRate: codataccounting.Float64(625.56),
                        TotalAmount: codataccounting.Float64(4649.44),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("78af4912-4772-45e6-a190-9e91044a5de5"),
                            Name: codataccounting.String("Shaun Schowalter"),
                        },
                        Currency: codataccounting.String("eaque"),
                        CurrencyRate: codataccounting.Float64(3973.05),
                        ID: codataccounting.String("670cf1cf-5932-4605-a51e-66bb426897d9"),
                        Note: codataccounting.String("omnis"),
                        PaidOnDate: codataccounting.String("dolorum"),
                        Reference: codataccounting.String("qui"),
                        TotalAmount: codataccounting.Float64(8610.2),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("velit"),
                        Currency: codataccounting.String("amet"),
                        CurrencyRate: codataccounting.Float64(3714.09),
                        TotalAmount: codataccounting.Float64(3815.34),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("70e93ee6-cf59-4f35-8aae-acae323a31bf"),
                            Name: codataccounting.String("Miss Luz Osinski"),
                        },
                        Currency: codataccounting.String("molestias"),
                        CurrencyRate: codataccounting.Float64(4623.08),
                        ID: codataccounting.String("716c802c-c9e0-4c7d-9d32-3f1aa63ed9cf"),
                        Note: codataccounting.String("et"),
                        PaidOnDate: codataccounting.String("nobis"),
                        Reference: codataccounting.String("quas"),
                        TotalAmount: codataccounting.Float64(3675.11),
                    },
                },
            },
            Reference: codataccounting.String("commodi"),
            SourceModifiedDate: codataccounting.String("soluta"),
            SubTotal: 7787.26,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "id": map[string]interface{}{
                        "inventore": "accusamus",
                        "maiores": "odit",
                    },
                    "numquam": map[string]interface{}{
                        "numquam": "culpa",
                        "aliquam": "iusto",
                    },
                    "voluptatibus": map[string]interface{}{
                        "maxime": "repellat",
                        "veritatis": "inventore",
                        "autem": "optio",
                    },
                },
            },
            TaxAmount: 8552.86,
            TotalAmount: 8174.25,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(349223),
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
        DirectIncomeID: "ut",
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
        DirectIncomeID: "dolore",
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
        DirectIncomeID: "numquam",
        TimeoutInMinutes: codataccounting.Int(639401),
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
        Query: codataccounting.String("reprehenderit"),
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
        DirectIncomeID: "nemo",
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
            Content: []byte("nisi"),
            RequestBody: "consequuntur",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "praesentium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
