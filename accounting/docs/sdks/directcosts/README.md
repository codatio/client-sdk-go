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
                DataType: codataccounting.String("beatae"),
                ID: "e8aa257d-dc19-412e-bde6-4bfcc5469d40",
            },
            Currency: "GBP",
            CurrencyRate: codataccounting.Float64(3383.08),
            ID: codataccounting.String("dfa79620-6bef-42b0-a3e4-2c1aa010e9aa"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e9135586-d18f-49f9-ba4b-fad2bf7d67ca"),
                        Name: codataccounting.String("Edwin Olson"),
                    },
                    Description: codataccounting.String("unde"),
                    DiscountAmount: codataccounting.Float64(7121.86),
                    DiscountPercentage: codataccounting.Float64(2626.11),
                    ItemRef: &shared.ItemRef{
                        ID: "1d612435-3187-40cf-a8b0-3ad421bd43d1",
                        Name: codataccounting.String("Larry Sanford MD"),
                    },
                    Quantity: 393.88,
                    SubTotal: codataccounting.Float64(81.99),
                    TaxAmount: codataccounting.Float64(210.15),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2214.8),
                        ID: codataccounting.String("eb22d9b3-a70d-494f-aa74-1c57d1fedc20"),
                        Name: codataccounting.String("Michelle Sipes"),
                    },
                    TotalAmount: codataccounting.Float64(8630.37),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("quo"),
                            ID: codataccounting.String("3ce18547-2f9e-4e69-966a-8be3444eac8b"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("laborum"),
                                ID: codataccounting.String("2875c6c1-fe60-46d0-bd2a-9c87ae50c166"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1a1d9136-a7e8-4d53-a13f-3f658752db76",
                            Name: codataccounting.String("Jody Hegmann"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "0a56cebc-ada2-49ca-b918-1c95671663c5",
                            Name: codataccounting.String("Cynthia Reichert"),
                        },
                    },
                    UnitAmount: 3905.83,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("et"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(2495.41),
                        TotalAmount: codataccounting.Float64(3829.28),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("38512ab2-521b-49f2-a072-467b8a40bc05"),
                            Name: codataccounting.String("Dr. Omar Reinger"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(196.1),
                        ID: codataccounting.String("edf22a94-d20e-4c90-aa41-d1f465e85156"),
                        Note: codataccounting.String("repellat"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("a"),
                        TotalAmount: codataccounting.Float64(4778.26),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(8372.02),
                        TotalAmount: codataccounting.Float64(9459.21),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("54fdd5ea-9543-4398-9afb-42a8d63388e4"),
                            Name: codataccounting.String("Casey Anderson"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(6353.96),
                        ID: codataccounting.String("5f9b18a2-44fd-4619-839d-acd38ed0dc67"),
                        Note: codataccounting.String("et"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("optio"),
                        TotalAmount: codataccounting.Float64(4654.86),
                    },
                },
            },
            Reference: codataccounting.String("delectus"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 8783.61,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "culpa": map[string]interface{}{
                        "quasi": "veniam",
                        "provident": "consequuntur",
                        "aut": "maxime",
                        "provident": "aperiam",
                    },
                },
            },
            TaxAmount: 8337.06,
            TotalAmount: 664.08,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(741454),
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
        DirectCostID: "magnam",
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
        DirectCostID: "natus",
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
        DirectCostID: "aperiam",
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
        Query: codataccounting.String("dicta"),
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

The *List direct cost attachments* endpoint returns a list of attachments available to download for given `directCostId`.

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
        DirectCostID: "repellat",
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
            Content: []byte("dolores"),
            RequestBody: "harum",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "quibusdam",
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

