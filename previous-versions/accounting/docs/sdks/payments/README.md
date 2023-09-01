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
    res, err := s.Payments.Create(ctx, operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("871a88ed-72a2-4d4a-b415-8ac2d0f0f58c"),
                Name: codataccounting.String("Pam Leannon"),
            },
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(4766.68),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codataccounting.String("eaque"),
                ID: "40d0d98e-9d82-4c5e-b06f-5576f5cdeb02",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("6d0bc43b-18ab-4378-b2fc-ff81ddf7e088"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 2892.95,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9630.72),
                            CurrencyRate: codataccounting.Float64(3403.92),
                            ID: codataccounting.String("4c9216e8-9263-413b-b6fc-2c8d2701096b"),
                            Type: shared.PaymentLinkTypeOther,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(4190.01),
                            CurrencyRate: codataccounting.Float64(6641.93),
                            ID: codataccounting.String("d6e3e1d9-d3b6-4603-b4a1-1aa1d5d2247d"),
                            Type: shared.PaymentLinkTypeManualJournal,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5974.19),
                            CurrencyRate: codataccounting.Float64(7299.26),
                            ID: codataccounting.String("3d46170e-768a-496b-b398-788398eba1bb"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(4939.36),
                            CurrencyRate: codataccounting.Float64(676.61),
                            ID: codataccounting.String("43356f63-49a1-4642-89b2-11ce46b95165"),
                            Type: shared.PaymentLinkTypeUnlinked,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 1068.42,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5418.42),
                            CurrencyRate: codataccounting.Float64(7604.45),
                            ID: codataccounting.String("a9142f05-2632-4b31-8ad6-92ffc8745005"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(6124.56),
                            CurrencyRate: codataccounting.Float64(8634.66),
                            ID: codataccounting.String("3d934e03-6f5c-4388-a64f-6985530a2e2a"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 4284.76,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(6672.29),
                            CurrencyRate: codataccounting.Float64(9471.82),
                            ID: codataccounting.String("863c28d0-40c6-49a3-9906-f6ebd5ad7ec7"),
                            Type: shared.PaymentLinkTypeInvoice,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5908.32),
                            CurrencyRate: codataccounting.Float64(2922.91),
                            ID: codataccounting.String("f25f634b-3730-4714-a6be-8c3e09c64d34"),
                            Type: shared.PaymentLinkTypeUnlinked,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(6784.76),
                            CurrencyRate: codataccounting.Float64(8087.97),
                            ID: codataccounting.String("299a6e5e-7aef-4134-82e9-45f53743efde"),
                            Type: shared.PaymentLinkTypeUnknown,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 5754.71,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1577.51),
                            CurrencyRate: codataccounting.Float64(1756.68),
                            ID: codataccounting.String("1f9b1f7d-9aff-4e69-a82a-ceefb04f8c51"),
                            Type: shared.PaymentLinkTypeUnlinked,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(8001.68),
                            CurrencyRate: codataccounting.Float64(6586.25),
                            ID: codataccounting.String("abea708e-d579-48d3-85d4-60599d5c3349"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(4904.31),
                            CurrencyRate: codataccounting.Float64(4326.34),
                            ID: codataccounting.String("d55209e9-a225-43b6-9765-886eeae5fd4b"),
                            Type: shared.PaymentLinkTypeInvoice,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("voluptatibus"),
            PaymentMethodRef: codataccounting.String("totam"),
            Reference: codataccounting.String("mollitia"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "provident": map[string]interface{}{
                        "aliquid": "in",
                    },
                    "quos": map[string]interface{}{
                        "sunt": "dolor",
                        "quisquam": "commodi",
                        "laudantium": "laboriosam",
                        "repellendus": "quos",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(2068.12),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(608738),
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
    res, err := s.Payments.Get(ctx, operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "earum",
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
    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("eligendi"),
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

