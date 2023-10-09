# AccountingPayments
(*AccountingPayments*)

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
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/types"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingPayments.CreateAccountingPayment(ctx, operations.CreateAccountingPaymentRequest{
        AccountingPayment: &shared.AccountingPayment{
            AccountRef: &shared.AccountRef{},
            Currency: syncforcommerceversion1.String("EUR"),
            CustomerRef: &shared.AccountingCustomerRef{
                ID: "<ID>",
            },
            Date: "2022-10-23T00:00:00.000Z",
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                    Amount: types.MustNewDecimalFromString("9211.94"),
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            AdditionalProperties: map[string]interface{}{
                                "Romaguera": "property",
                            },
                            Type: shared.PaymentLinkTypeInvoice,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{},
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            PaymentMethodRef: syncforcommerceversion1.String("Reduced"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "polymerize": map[string]interface{}{
                        "Terbium": "East",
                    },
                },
            },
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAccountingPaymentRequest](../../models/operations/createaccountingpaymentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.CreateAccountingPaymentResponse](../../models/operations/createaccountingpaymentresponse.md), error**

