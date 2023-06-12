# Payments

## Overview

Payments

### Available Operations

* [Create](#create) - Create payment
* [Get](#get) - Get payment
* [GetCreateModel](#getcreatemodel) - Get create payment model
* [List](#list) - List payments

## Create

Posts a new payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create payment model](https://docs.codat.io/accounting-api#/operations/get-create-payments-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) to see which integrations support this endpoint.

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
    res, err := s.Payments.Create(ctx, operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("306f5576-f5cd-4eb0-a86d-0bc43b18ab37"),
                Name: codataccounting.String("Darrin Crooks"),
            },
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(9819.74),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("laudantium"),
                ID: "1ddf7e08-8f74-4ef5-8c92-16e8926313bb",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("fc2c8d27-0109-46b6-aad6-e3e1d9d3b660"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 3089.27,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1192.28),
                            CurrencyRate: codataccounting.Float64(1231.37),
                            ID: codataccounting.String("aa1d5d22-47de-49b3-9461-70e768a96bb3"),
                            Type: shared.PaymentLinkTypePayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5218.9),
                            CurrencyRate: codataccounting.Float64(4416.03),
                            ID: codataccounting.String("88398eba-1bbf-4714-b356-f6349a164249"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1747.15),
                            CurrencyRate: codataccounting.Float64(1088.29),
                            ID: codataccounting.String("1ce46b95-1652-4b15-8ca9-142f052632b3"),
                            Type: shared.PaymentLinkTypeUnknown,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("harum"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("d692ffc8-7450-405e-9d3d-934e036f5c38"),
                Name: codataccounting.String("Brett Hudson"),
            },
            Reference: codataccounting.String("ea"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "corporis": map[string]interface{}{
                        "ipsum": "perferendis",
                        "est": "fugit",
                    },
                    "repudiandae": map[string]interface{}{
                        "similique": "repudiandae",
                    },
                    "assumenda": map[string]interface{}{
                        "fuga": "est",
                        "tenetur": "atque",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(4072.09),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(199180),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreatePaymentRequest](../../models/operations/createpaymentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.CreatePaymentResponse](../../models/operations/createpaymentresponse.md), error**


## Get

Get payment

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
    res, err := s.Payments.Get(ctx, operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetPaymentRequest](../../models/operations/getpaymentrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.GetPaymentResponse](../../models/operations/getpaymentresponse.md), error**


## GetCreateModel

Get create payment model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    res, err := s.Payments.GetCreateModel(ctx, operations.GetCreatePaymentsModelRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCreatePaymentsModelRequest](../../models/operations/getcreatepaymentsmodelrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.GetCreatePaymentsModelResponse](../../models/operations/getcreatepaymentsmodelresponse.md), error**


## List

Gets the latest payments for a company, with pagination

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
    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("eos"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListPaymentsRequest](../../models/operations/listpaymentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.ListPaymentsResponse](../../models/operations/listpaymentsresponse.md), error**

