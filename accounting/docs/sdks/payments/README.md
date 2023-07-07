# Payments

## Overview

Payments

### Available Operations

* [Create](#create) - Create payment
* [Get](#get) - Get payment
* [GetCreateModel](#getcreatemodel) - Get create payment model
* [List](#list) - List payments

## Create

The *Create payment* endpoint creates a new [payment](https://docs.codat.io/accounting-api#/schemas/Payment) for a given company's connection.

[Payments](https://docs.codat.io/accounting-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create payment model](https://docs.codat.io/accounting-api#/operations/get-create-payments-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating an account.


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
                ID: codataccounting.String("cbf18685-6a7e-482c-9f9d-0fc282c666af"),
                Name: codataccounting.String("Jacquelyn Dicki"),
            },
            Currency: codataccounting.String("USD"),
            CurrencyRate: codataccounting.Float64(5194.41),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("provident"),
                ID: "bea5d264-e41e-42ca-8482-2e513f6d9d2a",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("37c30990-77c1-40b6-8792-163e67d48860"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 2316.11,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(470),
                            CurrencyRate: codataccounting.Float64(6419.14),
                            ID: codataccounting.String("3049c3cf-6c02-476e-bb21-bad90d2743fd"),
                            Type: shared.PaymentLinkTypeOther,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7649.53),
                            CurrencyRate: codataccounting.Float64(1702.52),
                            ID: codataccounting.String("a10e6c29-78ec-4256-a5b0-9227fcc47996"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5744.03),
                            CurrencyRate: codataccounting.Float64(4585.85),
                            ID: codataccounting.String("7bbc57f3-8928-4a86-80c5-8d67d63e4aa5"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5547.96),
                            CurrencyRate: codataccounting.Float64(3030.69),
                            ID: codataccounting.String("64579cfc-6c0e-4503-b568-31f1d8ed87b2"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 5394.5,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9515.01),
                            CurrencyRate: codataccounting.Float64(6325.21),
                            ID: codataccounting.String("bc986e24-1e43-4b23-8241-7d13e3f62aa9"),
                            Type: shared.PaymentLinkTypePayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9236.58),
                            CurrencyRate: codataccounting.Float64(2889.02),
                            ID: codataccounting.String("ae8ab4a9-c492-4c5e-8ba5-d4aa4a508bd3"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(544.9),
                            CurrencyRate: codataccounting.Float64(7608.41),
                            ID: codataccounting.String("29aa8dd7-1bdd-4aa3-8b7b-91449ae69c08"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("labore"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("18bb7180-4f42-43d5-8393-5f377ac5c9b7"),
                Name: codataccounting.String("Joey Frami"),
            },
            Reference: codataccounting.String("id"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "minima": map[string]interface{}{
                        "amet": "quasi",
                    },
                    "doloremque": map[string]interface{}{
                        "recusandae": "iusto",
                        "impedit": "dolor",
                    },
                    "quaerat": map[string]interface{}{
                        "deserunt": "distinctio",
                        "alias": "voluptates",
                        "optio": "libero",
                        "voluptatum": "beatae",
                    },
                    "explicabo": map[string]interface{}{
                        "laboriosam": "ea",
                        "beatae": "eius",
                        "atque": "unde",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(2811.64),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(302190),
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

The *Get payment* endpoint returns a single payment for a given paymentId.

[Payments](https://docs.codat.io/accounting-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support getting a specific payment.

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
    res, err := s.Payments.Get(ctx, operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "fuga",
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

The *Get create payment model* endpoint returns the expected data for the request payload when creating a [payment](https://docs.codat.io/accounting-api#/schemas/Payment) for a given company and integration.

[Payments](https://docs.codat.io/accounting-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating a payment.


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

The *List payments* endpoint returns a list of [payments](https://docs.codat.io/accounting-api#/schemas/Payment) for a given company's connection.

[Payments](https://docs.codat.io/accounting-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

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
    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("voluptatum"),
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

