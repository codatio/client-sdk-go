# Expenses
(*Expenses*)

## Overview

Create and update transactions that represent your customers' spend.

### Available Operations

* [Create](#create) - Create expense transaction
* [Update](#update) - Update expense transactions

## Create

The *Create expense* endpoint creates an [expense transaction](https://docs.codat.io/sync-for-expenses-api#/schemas/ExpenseTransaction) in the accounting software for a given company's connection. 

[Expense transactions](https://docs.codat.io/sync-for-expenses-api#/schemas/ExpenseTransaction) represent transactions made with a company debit or credit card. 

### Supported Integrations

| Integration                   | Supported |
|-------------------------------|-----------|
| Dynamics 365 Business Central | Yes       |
| FreeAgent                     | Yes       |
| QuickBooks Desktop            | Yes       |
| QuickBooks Online             | Yes       |
| Oracle NetSuite               | Yes       |
| Xero                          | Yes       |

### Example Usage

```go
package main

import(
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Expenses.Create(ctx, operations.CreateExpenseTransactionRequest{
        RequestBody: []shared.ExpenseTransaction{
            shared.ExpenseTransaction{
                BankAccountRef: &shared.BankAccountReference{
                    ID: syncforexpenses.String("97"),
                },
                ContactRef: &shared.ExpenseContactRef{
                    ID: "430",
                    Type: shared.TypeSupplier.ToPointer(),
                },
                Currency: "GBP",
                CurrencyRate: types.MustNewDecimalFromString("1"),
                ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                IssueDate: "2024-05-21T00:00:00+00:00",
                Lines: []shared.ExpenseTransactionLine{
                    shared.ExpenseTransactionLine{
                        AccountRef: &shared.RecordRef{
                            ID: syncforexpenses.String("35"),
                        },
                        InvoiceTo: &shared.InvoiceTo{
                            ID: syncforexpenses.String("504"),
                            Type: shared.InvoiceToTypeCustomer.ToPointer(),
                        },
                        NetAmount: types.MustNewDecimalFromString("100"),
                        TaxAmount: types.MustNewDecimalFromString("20"),
                        TaxRateRef: &shared.RecordRef{
                            ID: syncforexpenses.String("23_Bills"),
                        },
                        TrackingRefs: []shared.TrackingRef{
                            shared.TrackingRef{
                                DataType: shared.TrackingRefDataTypeTrackingCategories.ToPointer(),
                                ID: syncforexpenses.String("DEPARTMENT_5"),
                            },
                        },
                    },
                },
                MerchantName: syncforexpenses.String("Amazon UK"),
                Notes: syncforexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                Type: shared.ExpenseTransactionTypePayment,
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateExpenseTransactionRequest](../../pkg/models/operations/createexpensetransactionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateExpenseTransactionResponse](../../pkg/models/operations/createexpensetransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

The *Update expense* endpoint updates an existing [expense transaction](https://docs.codat.io/sync-for-expenses-api#/schemas/UpdateExpenseRequest) in the accounting software for a given company's connection. 

[Expense transactions](https://docs.codat.io/sync-for-expenses-api#/schemas/UpdateExpenseRequest) represent transactions made with a company debit or credit card. 

### Supported Integrations
The following integrations are supported for the [Payment](https://docs.codat.io/expenses/sync-process/expense-transactions#transaction-types) transaction `type` only: 
| Integration           | Supported |
|-----------------------|-----------|
| FreeAgent             | Yes       |
| QuickBooks Online     | Yes       |
| Oracle NetSuite       | Yes       |
| Xero                  | Yes       |

### Example Usage

```go
package main

import(
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/types"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Expenses.Update(ctx, operations.UpdateExpenseTransactionRequest{
        UpdateExpenseRequest: &shared.UpdateExpenseRequest{
            BankAccountRef: &shared.UpdateExpenseRequestBankAccountReference{
                ID: syncforexpenses.String("97"),
            },
            ContactRef: &shared.ExpenseContactRef{
                ID: "430",
                Type: shared.TypeSupplier.ToPointer(),
            },
            Currency: "GBP",
            CurrencyRate: types.MustNewDecimalFromString("1"),
            IssueDate: "2024-05-21T00:00:00+00:00",
            Lines: []shared.ExpenseTransactionLine{
                shared.ExpenseTransactionLine{
                    AccountRef: &shared.RecordRef{
                        ID: syncforexpenses.String("35"),
                    },
                    InvoiceTo: &shared.InvoiceTo{
                        ID: syncforexpenses.String("504"),
                        Type: shared.InvoiceToTypeCustomer.ToPointer(),
                    },
                    NetAmount: types.MustNewDecimalFromString("100"),
                    TaxAmount: types.MustNewDecimalFromString("20"),
                    TaxRateRef: &shared.RecordRef{
                        ID: syncforexpenses.String("23_Bills"),
                    },
                    TrackingRefs: []shared.TrackingRef{
                        shared.TrackingRef{
                            DataType: shared.TrackingRefDataTypeTrackingCategories.ToPointer(),
                            ID: syncforexpenses.String("DEPARTMENT_5"),
                        },
                    },
                },
            },
            MerchantName: syncforexpenses.String("Amazon UK"),
            Notes: syncforexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
            Type: shared.UpdateExpenseRequestTypePayment,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateExpenseTransactionRequest](../../pkg/models/operations/updateexpensetransactionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateExpenseTransactionResponse](../../pkg/models/operations/updateexpensetransactionresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 422, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |