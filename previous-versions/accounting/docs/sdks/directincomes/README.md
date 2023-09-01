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

The *Create direct income* endpoint creates a new [direct income](https://docs.codat.io/accounting-api#/schemas/DirectIncome) for a given company's connection.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale. For example, cash sales of items to a customer, referral commissions, and service fee refunds are considered direct incomes.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
                DataType: codataccounting.String("delectus"),
                ID: "0f873f9d-5c25-4fd3-a0b4-a4a4253c3025",
            },
            Currency: "USD",
            CurrencyRate: codataccounting.Float64(1037.45),
            ID: codataccounting.String("1f42c7e7-dc54-48be-89e4-1a7a215ca12a"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("a9d59988-192c-4fd0-877c-53e7e7d4ee6e"),
                        Name: codataccounting.String("Miss Robin Marks"),
                    },
                    Description: codataccounting.String("porro"),
                    DiscountAmount: codataccounting.Float64(2321.12),
                    DiscountPercentage: codataccounting.Float64(5621.92),
                    ItemRef: &shared.ItemRef{
                        ID: "4e239670-3fec-431c-9082-4d189a36a6b2",
                        Name: codataccounting.String("Ernest King"),
                    },
                    Quantity: 4576.32,
                    SubTotal: codataccounting.Float64(351.52),
                    TaxAmount: codataccounting.Float64(4724.55),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6309.47),
                        ID: codataccounting.String("a60c8fe4-6e61-477d-b9db-3b70ffbb6970"),
                        Name: codataccounting.String("Phil Kuhlman DVM"),
                    },
                    TotalAmount: codataccounting.Float64(2046.97),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "097ef7c2-06e6-41b0-9308-714c20a3d986",
                            Name: codataccounting.String("Georgia Rolfson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "5c3fe655-74db-4af9-8a7c-98f13af28db2",
                            Name: codataccounting.String("Timmy Cruickshank"),
                        },
                    },
                    UnitAmount: 2633.86,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("f3ded356-d7e1-44b2-9cd9-8196d55af69a"),
                        Name: codataccounting.String("Bernadette Green"),
                    },
                    Description: codataccounting.String("natus"),
                    DiscountAmount: codataccounting.Float64(6847.46),
                    DiscountPercentage: codataccounting.Float64(9064.68),
                    ItemRef: &shared.ItemRef{
                        ID: "33681c23-c39a-47c0-a17c-b12c5ba825fe",
                        Name: codataccounting.String("Julia Sanford"),
                    },
                    Quantity: 7634.76,
                    SubTotal: codataccounting.Float64(7329.06),
                    TaxAmount: codataccounting.Float64(6677.35),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4047.58),
                        ID: codataccounting.String("fbfec932-af68-413d-a5bf-ecec2dd6916f"),
                        Name: codataccounting.String("Marta Schimmel"),
                    },
                    TotalAmount: codataccounting.Float64(8284.89),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "70ec60e6-0758-494d-a1c1-4cd90227e37c",
                            Name: codataccounting.String("Mable Monahan"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f1a5491a-be97-451b-906d-23e03e69815a",
                            Name: codataccounting.String("Stewart Moen"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "cde9e729-c9d4-4f2d-8a44-640ca60db73a",
                            Name: codataccounting.String("Josefina Mills"),
                        },
                    },
                    UnitAmount: 2884.94,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("67dc0d8d-a561-4220-a6ab-8f277485c197"),
                        Name: codataccounting.String("Marguerite Wiegand"),
                    },
                    Description: codataccounting.String("aut"),
                    DiscountAmount: codataccounting.Float64(8615.52),
                    DiscountPercentage: codataccounting.Float64(6337.32),
                    ItemRef: &shared.ItemRef{
                        ID: "7a089fc4-4db2-4745-b0e5-cc7c6d0cbcfd",
                        Name: codataccounting.String("Owen Feest"),
                    },
                    Quantity: 7072.58,
                    SubTotal: codataccounting.Float64(4228.03),
                    TaxAmount: codataccounting.Float64(9964.5),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4077.32),
                        ID: codataccounting.String("23bcecab-50ae-4e5e-8da8-b9af6ad05486"),
                        Name: codataccounting.String("Lonnie Rempel I"),
                    },
                    TotalAmount: codataccounting.Float64(7771.39),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "e2d176dc-1c43-4d40-b61d-171157cbe5ee",
                            Name: codataccounting.String("Darla Koelpin Sr."),
                        },
                        shared.TrackingCategoryRef{
                            ID: "840772f3-2e3b-449d-be0f-23b7b6d9948d",
                            Name: codataccounting.String("Tricia Schumm"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "477680fc-7a17-4a82-a5e8-2fd28d1040a7",
                            Name: codataccounting.String("Perry Brekke"),
                        },
                    },
                    UnitAmount: 1850.41,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("distinctio"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(7102.56),
                        TotalAmount: codataccounting.Float64(1033.28),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("835008f4-61ce-453e-9144-98a9ba460add"),
                            Name: codataccounting.String("Lionel Torphy Jr."),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(2396.46),
                        ID: codataccounting.String("7daa9182-a49d-4962-9d3c-affc198eea44"),
                        Note: codataccounting.String("quis"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("dignissimos"),
                        TotalAmount: codataccounting.Float64(6043.08),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8037.85),
                        TotalAmount: codataccounting.Float64(8241.62),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("440ea98b-ecce-4048-ade0-d56d73b00550"),
                            Name: codataccounting.String("Tasha Leannon"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1299.24),
                        ID: codataccounting.String("6ff77c65-675f-45b7-8e3e-4cfcc6a91ec5"),
                        Note: codataccounting.String("qui"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("sed"),
                        TotalAmount: codataccounting.Float64(2650.82),
                    },
                },
            },
            Reference: codataccounting.String("at"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 50.37,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quae": map[string]interface{}{
                        "repudiandae": "hic",
                        "non": "minima",
                    },
                },
            },
            TaxAmount: 7743.39,
            TotalAmount: 8885.73,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(657173),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateDirectIncomeRequest](../../models/operations/createdirectincomerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.CreateDirectIncomeResponse](../../models/operations/createdirectincomeresponse.md), error**


## DownloadAttachment

The *Download direct income attachment* endpoint downloads a specific attachment for a given `directIncomeId` and `attachmentId`.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support downloading a direct income attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        DirectIncomeID: "dicta",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DownloadDirectIncomeAttachmentRequest](../../models/operations/downloaddirectincomeattachmentrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.DownloadDirectIncomeAttachmentResponse](../../models/operations/downloaddirectincomeattachmentresponse.md), error**


## Get

The *Get direct income* endpoint returns a single direct income for a given directIncomeId.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support getting a specific direct income.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        DirectIncomeID: "vitae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncome != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDirectIncomeRequest](../../models/operations/getdirectincomerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetDirectIncomeResponse](../../models/operations/getdirectincomeresponse.md), error**


## GetAttachment

The *Get direct income attachment* endpoint returns a specific attachment for a given `directIncomeId` and `attachmentId`.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support getting a direct income attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        DirectIncomeID: "laborum",
        TimeoutInMinutes: codataccounting.Int(794929),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetDirectIncomeAttachmentRequest](../../models/operations/getdirectincomeattachmentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.GetDirectIncomeAttachmentResponse](../../models/operations/getdirectincomeattachmentresponse.md), error**


## GetCreateModel

The *Get create direct income model* endpoint returns the expected data for the request payload when creating a [direct income](https://docs.codat.io/accounting-api#/schemas/DirectIncome) for a given company and integration.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating a direct income.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCreateDirectIncomesModelRequest](../../models/operations/getcreatedirectincomesmodelrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.GetCreateDirectIncomesModelResponse](../../models/operations/getcreatedirectincomesmodelresponse.md), error**


## List

The *List direct incomes* endpoint returns a list of [direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) for a given company's connection.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        Query: codataccounting.String("minima"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncomes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListDirectIncomesRequest](../../models/operations/listdirectincomesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ListDirectIncomesResponse](../../models/operations/listdirectincomesresponse.md), error**


## ListAttachments

The *List direct income attachments* endpoint returns a list of attachments available to download for given `directIncomeId`.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support listing direct income attachments.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
        DirectIncomeID: "neque",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListDirectIncomeAttachmentsRequest](../../models/operations/listdirectincomeattachmentsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.ListDirectIncomeAttachmentsResponse](../../models/operations/listdirectincomeattachmentsresponse.md), error**


## UploadAttachment

The *Upload direct income attachment* endpoint uploads an attachment and assigns it against a specific `directIncomeId`.

[Direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome) are sales of items directly to a customer where payment is received at the point of the sale.

**Integration-specific behaviour**

For more details on supported file types by integration see [Attachments](https://docs.codat.io/accounting-api#/schemas/Attachment).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support uploading a direct income attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
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
            Content: []byte("debitis"),
            RequestBody: "quidem",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "quidem",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UploadDirectIncomeAttachmentRequest](../../models/operations/uploaddirectincomeattachmentrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.UploadDirectIncomeAttachmentResponse](../../models/operations/uploaddirectincomeattachmentresponse.md), error**

