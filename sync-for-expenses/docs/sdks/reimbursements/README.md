# Reimbursements
(*Reimbursements*)

## Overview

Create and update transactions that represent your customers' repayable spend.

### Available Operations

* [Create](#create) - Create reimbursable expense transaction
* [Update](#update) - Update reimbursable expense transaction

## Create

Use the *Create reimbursable expense* endpoint to submit an employee expense claim in the accounting platform for a given company's connection.

[Reimbursable expense requests](https://docs.codat.io/sync-for-expenses-api#/schemas/ReimbursableExpenseTransactionRequest) are reflected in the accounting software in the form of **Bills** against an employee (who exists as a supplier in the accounting platform).

### Supported Integrations
| Integration           | Supported |
|-----------------------|-----------|
| FreeAgent             | Yes       |
| QuickBooks Desktop    | Yes       |
| QuickBooks Online     | Yes       |
| Oracle NetSuite       | Yes       |

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
    res, err := s.Reimbursements.Create(ctx, operations.CreateReimbursableExpenseTransactionRequest{
        RequestBody: []shared.ReimbursableExpenseTransaction{
            shared.ReimbursableExpenseTransaction{
                ContactRef: shared.ReimbursementContactRef{
                    ID: "752",
                },
                Currency: "GBP",
                CurrencyRate: types.MustNewDecimalFromString("1"),
                DueDate: "2024-05-21",
                ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                IssueDate: "2024-05-21",
                Lines: []shared.ReimbursableExpenseTransactionLine{
                    shared.ReimbursableExpenseTransactionLine{
                        AccountRef: &shared.RecordRef{
                            ID: syncforexpenses.String("35"),
                        },
                        Description: syncforexpenses.String("Hotel"),
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
                Notes: syncforexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                Reference: syncforexpenses.String("expenses w/c 01/07"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateReimbursableExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.CreateReimbursableExpenseTransactionRequest](../../pkg/models/operations/createreimbursableexpensetransactionrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.CreateReimbursableExpenseTransactionResponse](../../pkg/models/operations/createreimbursableexpensetransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## Update

The *Update reimbursable expense* endpoint updates an existing employee expense claim in the accounting platform for a given company's connection. 

Updating an existing [reimbursable expense transaction](https://docs.codat.io/sync-for-expenses-api#/schemas/UpdateReimbursableExpenseTransactionRequest) will update the existing **bill** against an employee (who exists as a supplier in the accounting software).

### Supported Integrations
| Integration           | Supported |
|-----------------------|-----------|
| FreeAgent             | Yes       |
| QuickBooks Online     | Yes       |
| Oracle NetSuite       | Yes       |

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
    res, err := s.Reimbursements.Update(ctx, operations.UpdateReimbursableExpenseTransactionRequest{
        UpdateReimbursableExpenseTransactionRequest: &shared.UpdateReimbursableExpenseTransactionRequest{
            ContactRef: shared.ReimbursementContactRef{
                ID: "752",
            },
            Currency: "GBP",
            CurrencyRate: types.MustNewDecimalFromString("1"),
            DueDate: "2024-05-21",
            IssueDate: "2024-05-21",
            Lines: []shared.ReimbursableExpenseTransactionLine{
                shared.ReimbursableExpenseTransactionLine{
                    AccountRef: &shared.RecordRef{
                        ID: syncforexpenses.String("35"),
                    },
                    Description: syncforexpenses.String("Hotel"),
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
            Notes: syncforexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
            Reference: syncforexpenses.String("expenses w/c 01/07"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateReimbursableExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.UpdateReimbursableExpenseTransactionRequest](../../pkg/models/operations/updatereimbursableexpensetransactionrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.UpdateReimbursableExpenseTransactionResponse](../../pkg/models/operations/updatereimbursableexpensetransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |