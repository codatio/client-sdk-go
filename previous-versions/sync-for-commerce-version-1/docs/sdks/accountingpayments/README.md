# AccountingPayments

## Overview

Payments

### Available Operations

* [CreateAccountingPayment](#createaccountingpayment) - Create payment

## CreateAccountingPayment

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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingPayments.CreateAccountingPayment(ctx, operations.CreateAccountingPaymentRequest{
        AccountingPayment: &shared.AccountingPayment{
            AccountRef: &shared.AccountRef{
                ID: codatsynccommerce.String("2f64d1db-1f2c-4431-8661-e96349e1cf9e"),
                Name: codatsynccommerce.String("Alma Waters"),
            },
            Currency: codatsynccommerce.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("2244.67"),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("iusto"),
                ID: "000ae6b6-bc9b-48f7-99ea-c55a9741d311",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codatsynccommerce.String("52965bb8-a720-4261-9435-e139dbc2259b"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Amount: *types.MustNewDecimalFromString("6633.18"),
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: types.MustNewDecimalFromString("7278.88"),
                            CurrencyRate: types.MustNewDecimalFromString("8544.6"),
                            ID: codatsynccommerce.String("a8c070e1-084c-4b06-b2d1-ad879eeb9665"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("corporis"),
            PaymentMethodRef: codatsynccommerce.String("officiis"),
            Reference: codatsynccommerce.String("voluptatibus"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "at": map[string]interface{}{
                        "alias": "quia",
                    },
                },
            },
            TotalAmount: types.MustNewDecimalFromString("6941.58"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(684126),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAccountingPaymentRequest](../../models/operations/createaccountingpaymentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.CreateAccountingPaymentResponse](../../models/operations/createaccountingpaymentresponse.md), error**

