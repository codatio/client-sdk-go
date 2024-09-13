# CodatLendingPayments
(*LoanWriteback.Payments*)

## Overview

### Available Operations

* [Create](#create) - Create payment
* [GetCreateModel](#getcreatemodel) - Get create payment model

## Create

The *Create payment* endpoint creates a new [payment](https://docs.codat.io/lending-api#/schemas/Payment) for a given company's connection.

[Payments](https://docs.codat.io/lending-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create payment model](https://docs.codat.io/lending-api#/operations/get-create-payments-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/types"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Payments.Create(ctx, operations.CreatePaymentRequest{
        AccountingPayment: &shared.AccountingPayment{
            AccountRef: &shared.AccountRef{
                ID: lending.String("8000002E-1675267199"),
                Name: lending.String("Undeposited Funds"),
            },
            Currency: lending.String("USD"),
            CurrencyRate: types.MustNewDecimalFromString("1"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: lending.String("string"),
                ID: "80000002-1674552702",
            },
            Date: "2023-02-10T11:47:04.792Z",
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: lending.String("2023-02-11T11:47:04.792Z"),
                    Amount: types.MustNewDecimalFromString("28"),
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: types.MustNewDecimalFromString("-28"),
                            CurrencyRate: types.MustNewDecimalFromString("1"),
                            ID: lending.String("181-1676374586"),
                            Type: shared.PaymentLinkTypeInvoice,
                        },
                    },
                },
            },
            ModifiedDate: lending.String("2022-10-23T00:00:00Z"),
            Note: lending.String("note 14/02 1147"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: "string",
                Name: lending.String("string"),
            },
            Reference: lending.String("ref 14/02 1147"),
            SourceModifiedDate: lending.String("2022-10-23T00:00:00Z"),
            TotalAmount: types.MustNewDecimalFromString("28"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingCreatePaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreatePaymentRequest](../../pkg/models/operations/createpaymentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreatePaymentResponse](../../pkg/models/operations/createpaymentresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## GetCreateModel

The *Get create payment model* endpoint returns the expected data for the request payload when creating a [payment](https://docs.codat.io/lending-api#/schemas/Payment) for a given company and integration.

[Payments](https://docs.codat.io/lending-api#/schemas/Payment) represent an allocation of money within any customer accounts receivable account.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating a payment.


### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/shared"
	lending "github.com/codatio/client-sdk-go/lending/v5"
	"context"
	"github.com/codatio/client-sdk-go/lending/v5/pkg/models/operations"
	"log"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.LoanWriteback.Payments.GetCreateModel(ctx, operations.GetCreatePaymentModelRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCreatePaymentModelRequest](../../pkg/models/operations/getcreatepaymentmodelrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetCreatePaymentModelResponse](../../pkg/models/operations/getcreatepaymentmodelresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
