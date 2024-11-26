# BillPayments
(*BillPayments*)

## Overview

Get, create, and update Bill payments.

### Available Operations

* [GetPaymentOptions](#getpaymentoptions) - Get payment mapping options
* [Create](#create) - Create bill payment

## GetPaymentOptions

Use the *Get mapping options - Payments* endpoint to return a list of available mapping options for a given company's connection ID.

By default, this endpoint returns a list of active bank accounts. You can use [querying](https://docs.codat.io/using-the-api/querying) to change that.

Mapping options are a set of bank accounts used to configure the SMB's payables integration.

### Example Usage

```go
package main

import(
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillPayments.GetPaymentOptions(ctx, operations.GetMappingOptionsPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ContinuationToken: syncforpayables.String("continuationToken=eyJwYWdlIjoyLCJwYWdlU2l6ZSI6MTAwLCJwYWdlQ291bnQiOjExfQ=="),
        StatusQuery: syncforpayables.String("status=Archived"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentMappingOptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetMappingOptionsPaymentsRequest](../../pkg/models/operations/getmappingoptionspaymentsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetMappingOptionsPaymentsResponse](../../pkg/models/operations/getmappingoptionspaymentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## Create

The *Create bill payment* endpoint creates a new [bill payment](https://docs.codat.io/sync-for-payables-api#/schemas/BillPayment) for a given company's connection.

[Bill payments](https://docs.codat.io/sync-for-payables-api#/schemas/BillPayment) are an allocation of money within any Accounts Payable account.

### Example Usage

```go
package main

import(
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.BillPayments.Create(ctx, operations.CreateBillPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        BillID: "9wg4lep4ush5cxs79pl8sozmsndbaukll3ind4g7buqbm1h2",
        BillPaymentPrototype: &shared.BillPaymentPrototype{
            Amount: types.MustNewDecimalFromString("22"),
            Date: "2022-10-23T00:00:00.000Z",
            Reference: syncforpayables.String("Bill Payment against bill c13e37b6 dfaa-4894-b3be-9fe97bda9f44"),
            AccountRef: shared.BillPaymentAccountRef{
                ID: "7bda9f44sr56",
            },
            CurrencyRate: types.MustNewDecimalFromString("1"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillPayment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateBillPaymentRequest](../../pkg/models/operations/createbillpaymentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateBillPaymentResponse](../../pkg/models/operations/createbillpaymentresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |