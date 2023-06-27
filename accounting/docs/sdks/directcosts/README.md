# DirectCosts

## Overview

Direct costs

### Available Operations

* [Create](#create) - Create direct cost
* [DownloadAttachment](#downloadattachment) - Download direct cost attachment
* [Get](#get) - Get direct cost
* [GetAttachment](#getattachment) - Get direct cost attachment
* [GetCreateModel](#getcreatemodel) - Get create direct cost model
* [List](#list) - List direct costs
* [ListAttachments](#listattachments) - List direct cost attachments
* [UploadAttachment](#uploadattachment) - Upload direct cost attachment

## Create

The *Create direct cost* endpoint creates a new [direct cost](https://docs.codat.io/accounting-api#/schemas/DirectCost) for a given company's connection.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/accounting-api#/operations/get-create-directCosts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating an account.


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
    res, err := s.DirectCosts.Create(ctx, operations.CreateDirectCostRequest{
        DirectCost: &shared.DirectCost{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("ratione"),
                ID: "f4eedbe7-8bf6-4068-a589-4ea763d5c727",
            },
            Currency: "USD",
            CurrencyRate: codataccounting.Float64(3378.33),
            ID: codataccounting.String("b785148d-6d54-49e5-a35b-33bc0f970c42"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("9f484422-5e75-4b79-a065-c0efa6f93b90"),
                        Name: codataccounting.String("Jack Rau"),
                    },
                    Description: codataccounting.String("unde"),
                    DiscountAmount: codataccounting.Float64(3331.5),
                    DiscountPercentage: codataccounting.Float64(7243.07),
                    ItemRef: &shared.ItemRef{
                        ID: "e1254b73-9f4f-4e77-a10d-1f6558c99c72",
                        Name: codataccounting.String("Mona Considine"),
                    },
                    Quantity: 349.2,
                    SubTotal: codataccounting.Float64(9387.73),
                    TaxAmount: codataccounting.Float64(5669.15),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2568.9),
                        ID: codataccounting.String("087d9caa-e042-4dd7-8aac-9b4caa1cfe9e"),
                        Name: codataccounting.String("Joann Stokes"),
                    },
                    TotalAmount: codataccounting.Float64(576.67),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("amet"),
                            ID: codataccounting.String("907f3783-1983-4d42-a54a-85466597c502"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("nesciunt"),
                                ID: codataccounting.String("c1471d51-aaa6-4ddf-9abd-6487c5fc2b86"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "a00bef69-e100-4157-a30b-da7afded84a3",
                            Name: codataccounting.String("Mr. Sonya Gutmann"),
                        },
                    },
                    UnitAmount: 5454,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e1a735ac-26ae-433b-af97-1a8f46bca110"),
                        Name: codataccounting.String("Marta Torphy"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(7436.31),
                    DiscountPercentage: codataccounting.Float64(4565.91),
                    ItemRef: &shared.ItemRef{
                        ID: "11d08cf8-8ec9-4f7b-99a5-50a656ed333b",
                        Name: codataccounting.String("Edward Sanford"),
                    },
                    Quantity: 6679.43,
                    SubTotal: codataccounting.Float64(6847.08),
                    TaxAmount: codataccounting.Float64(4224.44),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3351.76),
                        ID: codataccounting.String("432a986e-b7e1-44ca-9640-89150097019a"),
                        Name: codataccounting.String("Gwendolyn Wintheiser"),
                    },
                    TotalAmount: codataccounting.Float64(9368.8),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("impedit"),
                            ID: codataccounting.String("e7bf904e-0110-45d3-8908-162c6beb68a0"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("vel"),
                                ID: codataccounting.String("57b7d03a-1480-4f8d-a30f-069d810618d9"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("voluptate"),
                                ID: codataccounting.String("e1522975-10da-4803-9229-2cc61c2a702b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("soluta"),
                                ID: codataccounting.String("97ee102d-a2de-435f-8e01-bf33eaab4540"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dolores"),
                                ID: codataccounting.String("ac1704bf-1cc9-4fc6-9aae-5eb5f0c492b5"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "44d08a22-67aa-4ee7-9e3c-71ad31becb83",
                            Name: codataccounting.String("Martin Fahey"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ae3bfc23-d945-40a9-86a4-95bac707f06b",
                            Name: codataccounting.String("Myrtle Walker"),
                        },
                    },
                    UnitAmount: 5431.22,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("6492386f-62c9-469c-8cc6-b78890a3fd3c"),
                        Name: codataccounting.String("Jerry Spinka Jr."),
                    },
                    Description: codataccounting.String("asperiores"),
                    DiscountAmount: codataccounting.Float64(5342.04),
                    DiscountPercentage: codataccounting.Float64(7596.13),
                    ItemRef: &shared.ItemRef{
                        ID: "23df931d-a3ed-4b51-bad9-4acc94351377",
                        Name: codataccounting.String("Mrs. Eileen Spinka"),
                    },
                    Quantity: 1584.51,
                    SubTotal: codataccounting.Float64(1199.27),
                    TaxAmount: codataccounting.Float64(7214.48),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5545.08),
                        ID: codataccounting.String("32a56d69-180f-4f60-ab9a-6658e69a4b84"),
                        Name: codataccounting.String("Desiree Fisher"),
                    },
                    TotalAmount: codataccounting.Float64(8447.03),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("harum"),
                            ID: codataccounting.String("ec75c68c-6065-4946-8ce3-04d8849bf821"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("eligendi"),
                                ID: codataccounting.String("337f96bb-0c69-4e37-adb1-344ba9f78a5c"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("accusantium"),
                                ID: codataccounting.String("ed7aab62-e972-461f-b0c5-8d27b51996b5"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4b50eef7-12b7-4a7a-b034-4b1710688dee",
                            Name: codataccounting.String("Santiago Zboncak"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7f3dd0cc-d33f-411b-be4e-080aa104186e",
                            Name: codataccounting.String("Darren Herman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "02f3702c-5c8e-42d3-8ead-3104fa44707b",
                            Name: codataccounting.String("Johnny Kunze"),
                        },
                    },
                    UnitAmount: 2839.36,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4282821f-db2f-469e-9926-7c71cc8d3cd4"),
                        Name: codataccounting.String("Dolores Lehner I"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(5378.94),
                    DiscountPercentage: codataccounting.Float64(6718.73),
                    ItemRef: &shared.ItemRef{
                        ID: "82c808fe-2751-4a20-87c0-449e143f9619",
                        Name: codataccounting.String("Bennie Kirlin"),
                    },
                    Quantity: 460.36,
                    SubTotal: codataccounting.Float64(8739.01),
                    TaxAmount: codataccounting.Float64(3575.89),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6731.58),
                        ID: codataccounting.String("11fa436e-6259-4233-b95c-9d237397c785"),
                        Name: codataccounting.String("Jorge Stokes"),
                    },
                    TotalAmount: codataccounting.Float64(9792.55),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("quis"),
                            ID: codataccounting.String("00183feb-df67-46b7-a06d-ab750052a564"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("officiis"),
                                ID: codataccounting.String("dc439ed8-c432-40f4-9240-d4487ac693b9"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("quaerat"),
                                ID: codataccounting.String("c3b9d248-8d79-45aa-82fc-405669f69a00"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d2124945-0819-4d7c-bb1b-41844060e003",
                            Name: codataccounting.String("Mr. Angela Schuppe"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "dc901f5a-fd2a-46c4-8846-ae9d89253c89",
                            Name: codataccounting.String("Kathryn Windler"),
                        },
                    },
                    UnitAmount: 5683.23,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("rerum"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(9102.24),
                        TotalAmount: codataccounting.Float64(2717.82),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("652d3c34-3d61-4778-af49-1247725e6219"),
                            Name: codataccounting.String("Violet Thiel Jr."),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2893.22),
                        ID: codataccounting.String("a5de59ac-7706-4670-8f1c-f5932605251e"),
                        Note: codataccounting.String("ex"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("harum"),
                        TotalAmount: codataccounting.Float64(7414),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(3850.49),
                        TotalAmount: codataccounting.Float64(5336.81),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("97d99a2d-3356-470e-93ee-6cf59f358aae"),
                            Name: codataccounting.String("Devin Nienow"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2083.1),
                        ID: codataccounting.String("a31bf7ba-1cc9-4771-ac80-2cc9e0c7d9d3"),
                        Note: codataccounting.String("dolores"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("repellat"),
                        TotalAmount: codataccounting.Float64(647.23),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(4209.27),
                        TotalAmount: codataccounting.Float64(1945.14),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("ed9cf1c8-56bc-4ba5-9ef2-454a47facf11"),
                            Name: codataccounting.String("Marianne Steuber"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2914.14),
                        ID: codataccounting.String("4a756287-3c7d-4d9e-baf4-3dc623620f31"),
                        Note: codataccounting.String("dolor"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("doloribus"),
                        TotalAmount: codataccounting.Float64(2021.77),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(9458.52),
                        TotalAmount: codataccounting.Float64(1945.26),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("db022faa-565f-4b8f-a52e-bb9d38383879"),
                            Name: codataccounting.String("Beverly Green"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(6212.3),
                        ID: codataccounting.String("3dab30e9-17f5-40fd-a04c-8b1bb55a292b"),
                        Note: codataccounting.String("aut"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quo"),
                        TotalAmount: codataccounting.Float64(2482.32),
                    },
                },
            },
            Reference: codataccounting.String("libero"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 4850.26,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aliquam": map[string]interface{}{
                        "nisi": "labore",
                        "accusamus": "cum",
                    },
                    "sunt": map[string]interface{}{
                        "voluptatem": "non",
                        "ipsum": "laudantium",
                        "totam": "facilis",
                        "consequatur": "assumenda",
                    },
                },
            },
            TaxAmount: 1042.64,
            TotalAmount: 7272.56,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(697864),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectCostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateDirectCostRequest](../../models/operations/createdirectcostrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.CreateDirectCostResponse](../../models/operations/createdirectcostresponse.md), error**


## DownloadAttachment

The *Download direct cost attachment* endpoint downloads a specific attachment for a given `directCostId` and `attachmentId`.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support downloading a direct cost attachment.


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
    res, err := s.DirectCosts.DownloadAttachment(ctx, operations.DownloadDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "architecto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DownloadDirectCostAttachmentRequest](../../models/operations/downloaddirectcostattachmentrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.DownloadDirectCostAttachmentResponse](../../models/operations/downloaddirectcostattachmentresponse.md), error**


## Get

The *Get direct cost* endpoint returns a single direct cost for a given directCostId.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support getting a specific direct cost.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


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
    res, err := s.DirectCosts.Get(ctx, operations.GetDirectCostRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCost != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetDirectCostRequest](../../models/operations/getdirectcostrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.GetDirectCostResponse](../../models/operations/getdirectcostresponse.md), error**


## GetAttachment

The *Get direct cost attachment* endpoint returns a specific attachment for a given `directCostId` and `attachmentId`.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support getting a direct cost attachment.


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
    res, err := s.DirectCosts.GetAttachment(ctx, operations.GetDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetDirectCostAttachmentRequest](../../models/operations/getdirectcostattachmentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetDirectCostAttachmentResponse](../../models/operations/getdirectcostattachmentresponse.md), error**


## GetCreateModel

The *Get create direct cost model* endpoint returns the expected data for the request payload when creating a [direct cost](https://docs.codat.io/accounting-api#/schemas/DirectCost) for a given company and integration.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating a direct cost.


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
    res, err := s.DirectCosts.GetCreateModel(ctx, operations.GetCreateDirectCostsModelRequest{
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

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCreateDirectCostsModelRequest](../../models/operations/getcreatedirectcostsmodelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetCreateDirectCostsModelResponse](../../models/operations/getcreatedirectcostsmodelresponse.md), error**


## List

The *List direct costs* endpoint returns a list of [direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) for a given company's connection.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

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
    res, err := s.DirectCosts.List(ctx, operations.ListDirectCostsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("tenetur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCosts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListDirectCostsRequest](../../models/operations/listdirectcostsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.ListDirectCostsResponse](../../models/operations/listdirectcostsresponse.md), error**


## ListAttachments

The *List direct cost attachments* endpoint returns a list of attachments avialable to download for given `directCostId`.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support listing direct cost attachments.


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
    res, err := s.DirectCosts.ListAttachments(ctx, operations.ListDirectCostAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListDirectCostAttachmentsRequest](../../models/operations/listdirectcostattachmentsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.ListDirectCostAttachmentsResponse](../../models/operations/listdirectcostattachmentsresponse.md), error**


## UploadAttachment

The *Upload direct cost attachment* endpoint uploads an attachment and assigns it against a specific `directCostId`.

[Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) are purchases of items that are paid off at the point of the purchase.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/accounting-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support uploading a direct cost attachment.


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
    res, err := s.DirectCosts.UploadAttachment(ctx, operations.UploadDirectCostAttachmentRequest{
        RequestBody: &operations.UploadDirectCostAttachmentRequestBody{
            Content: []byte("eveniet"),
            RequestBody: "reprehenderit",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "incidunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UploadDirectCostAttachmentRequest](../../models/operations/uploaddirectcostattachmentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.UploadDirectCostAttachmentResponse](../../models/operations/uploaddirectcostattachmentresponse.md), error**

