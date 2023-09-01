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
                ID: codatsynccommerce.String("7d56844e-ded8-45a9-865e-628bdfc2032b"),
                Name: codatsynccommerce.String("Angelica Langworth"),
            },
            Currency: codatsynccommerce.String("USD"),
            CurrencyRate: codatsynccommerce.Float64(1443.97),
            CustomerRef: &shared.AccountingCustomerRef{
                CompanyName: codatsynccommerce.String("dolorem"),
                ID: "b7e13584-f7ae-412c-a891-f82ce1157172",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codatsynccommerce.String("05377dcf-a89d-4f97-9e35-6686092e9c3d"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Amount: 3269.42,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(1048.34),
                            CurrencyRate: codatsynccommerce.Float64(1142.12),
                            ID: codatsynccommerce.String("1dea1026-d541-4a4d-990f-eb21780bccc0"),
                            Type: shared.PaymentLinkTypeManualJournal,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(7188.79),
                            CurrencyRate: codatsynccommerce.Float64(7148.35),
                            ID: codatsynccommerce.String("ddb48470-8fb4-4e39-9e6b-c158c4c4e545"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(5786.1),
                            CurrencyRate: codatsynccommerce.Float64(9294.43),
                            ID: codatsynccommerce.String("a342260e-9b20-40ce-b8a1-bd8fb7a0a116"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(8855.23),
                            CurrencyRate: codatsynccommerce.Float64(4909.56),
                            ID: codatsynccommerce.String("23d4097f-a30e-49af-b25b-29122030d83f"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Amount: 9319.53,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(4802.95),
                            CurrencyRate: codatsynccommerce.Float64(4938.65),
                            ID: codatsynccommerce.String("99d22e8c-1f84-4938-a5fd-c42c876c2c2d"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(6908.71),
                            CurrencyRate: codatsynccommerce.Float64(3049.64),
                            ID: codatsynccommerce.String("cfc1c762-30f8-441f-b1bd-23fdb14db6be"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(6422.79),
                            CurrencyRate: codatsynccommerce.Float64(3755.88),
                            ID: codatsynccommerce.String("85998e22-ae20-4da1-afc2-b271a289c57e"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Amount: 2505.2,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(6050.43),
                            CurrencyRate: codatsynccommerce.Float64(585.67),
                            ID: codatsynccommerce.String("439d2224-6569-4462-8070-84f7ab37cef0"),
                            Type: shared.PaymentLinkTypeUnlinked,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(1806.6),
                            CurrencyRate: codatsynccommerce.Float64(1622.51),
                            ID: codatsynccommerce.String("5194db55-410a-4dc6-a9af-90a26c7cdc98"),
                            Type: shared.PaymentLinkTypeUnknown,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(9788.57),
                            CurrencyRate: codatsynccommerce.Float64(298.53),
                            ID: codatsynccommerce.String("68981d6b-b33c-4faa-b48c-31bf407ee4fc"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(0.73),
                            CurrencyRate: codatsynccommerce.Float64(7704.67),
                            ID: codatsynccommerce.String("42b78f15-6263-498a-8dc7-66324ccb06c8"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Amount: 815.41,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codatsynccommerce.Float64(8310.37),
                            CurrencyRate: codatsynccommerce.Float64(271.97),
                            ID: codatsynccommerce.String("2529270b-8d57-422d-9895-b8bcf24db959"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Note: codatsynccommerce.String("amet"),
            PaymentMethodRef: codatsynccommerce.String("dolor"),
            Reference: codatsynccommerce.String("nostrum"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "molestiae": map[string]interface{}{
                        "ullam": "velit",
                        "adipisci": "cupiditate",
                    },
                    "occaecati": map[string]interface{}{
                        "fugiat": "molestiae",
                        "quas": "repellendus",
                    },
                    "saepe": map[string]interface{}{
                        "distinctio": "vel",
                    },
                    "necessitatibus": map[string]interface{}{
                        "nesciunt": "corrupti",
                        "cupiditate": "voluptatibus",
                        "ullam": "dolorum",
                    },
                },
            },
            TotalAmount: codatsynccommerce.Float64(7437.05),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(739946),
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

