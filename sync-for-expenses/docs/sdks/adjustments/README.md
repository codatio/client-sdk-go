# Adjustments
(*Adjustments*)

## Overview

Create transactions that represent your adjustments to your customers' spend.

### Available Operations

* [Create](#create) - Create adjustment transaction

## Create

Use the *Create adjustment expense* endpoint to create an [adjustment](https://docs.codat.io/sync-for-expenses-api#/schemas/AdjustmentTransactionRequest) in the accounting software for a given company's connection. 

Adjustments represent write-offs and transaction alterations, such as foreign exchange adjustments, in the form of a journal entry. 

### Supported Integrations

| Integration           | Supported |
|-----------------------|-----------|
| QuickBooks Desktop    | Yes       |

### Example Usage

```go
package main

import(
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Adjustments.Create(ctx, operations.CreateAdjustmentTransactionRequest{
        RequestBody: []shared.AdjustmentTransactionRequest{
            shared.AdjustmentTransactionRequest{
                Currency: "USD",
                CurrencyRate: types.MustNewDecimalFromString("1"),
                Date: "2024-05-21T00:00:00+00:00",
                ID: "3357b3df-5f2e-465d-b9ba-226519dbb8f1",
                Lines: []shared.AdjustmentTransactionLine{
                    shared.AdjustmentTransactionLine{
                        AccountRef: shared.RecordRef{
                            ID: syncforexpenses.String("80000018-1671793811"),
                        },
                        Amount: types.MustNewDecimalFromString("50"),
                        Description: syncforexpenses.String("debit line"),
                        InvoiceTo: &shared.InvoiceTo{
                            ID: syncforexpenses.String("80000002-1674552702"),
                            Type: shared.InvoiceToTypeCustomer.ToPointer(),
                        },
                        TrackingRefs: []shared.TrackingRefAdjustmentTransaction{
                            shared.TrackingRefAdjustmentTransaction{
                                DataType: shared.TrackingRefAdjustmentTransactionDataTypeTrackingCategories.ToPointer(),
                                ID: syncforexpenses.String("80000003-1674553958"),
                            },
                        },
                    },
                    shared.AdjustmentTransactionLine{
                        AccountRef: shared.RecordRef{
                            ID: syncforexpenses.String("80000028-1671794219"),
                        },
                        Amount: types.MustNewDecimalFromString("-50"),
                        Description: syncforexpenses.String("credit line"),
                        InvoiceTo: &shared.InvoiceTo{
                            ID: syncforexpenses.String("80000002-1674552702"),
                            Type: shared.InvoiceToTypeCustomer.ToPointer(),
                        },
                        TrackingRefs: []shared.TrackingRefAdjustmentTransaction{
                            shared.TrackingRefAdjustmentTransaction{
                                DataType: shared.TrackingRefAdjustmentTransactionDataTypeTrackingCategories.ToPointer(),
                                ID: syncforexpenses.String("80000003-1674553958"),
                            },
                        },
                    },
                },
                Reference: syncforexpenses.String("test reference"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdjustmentTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateAdjustmentTransactionRequest](../../pkg/models/operations/createadjustmenttransactionrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateAdjustmentTransactionResponse](../../pkg/models/operations/createadjustmenttransactionresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
