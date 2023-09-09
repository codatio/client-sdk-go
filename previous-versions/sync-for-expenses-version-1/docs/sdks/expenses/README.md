# Expenses

## Overview

Create expense datasets and upload receipts.

### Available Operations

* [CreateExpenseDataset](#createexpensedataset) - Create expense-transactions
* [UpdateExpenseDataset](#updateexpensedataset) - Update expense-transactions
* [UploadAttachment](#uploadattachment) - Upload attachment

## CreateExpenseDataset

Create an expense transaction

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.CreateExpenseDataset(ctx, operations.CreateExpenseDatasetRequest{
        CreateExpenseRequest: &shared.CreateExpenseRequest{
            Items: []shared.ExpenseTransaction{
                shared.ExpenseTransaction{
                    ContactRef: &shared.ContactRef{
                        ContactType: shared.ContactRefContactTypeSupplier.ToPointer(),
                        ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    Currency: "GBP",
                    CurrencyRate: types.MustNewDecimalFromString("6027.63"),
                    ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                    IssueDate: "2022-10-23T00:00:00.000Z",
                    Lines: []shared.ExpenseTransactionLine{
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: *types.MustNewDecimalFromString("110.42"),
                            TaxAmount: *types.MustNewDecimalFromString("14.43"),
                            TaxRateRef: &shared.RecordRef{
                                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                    },
                    MerchantName: codatsyncexpenses.String("Amazon UK"),
                    Notes: codatsyncexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                    Type: shared.ExpenseTransactionTypePayment,
                },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateExpenseDatasetRequest](../../models/operations/createexpensedatasetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreateExpenseDatasetResponse](../../models/operations/createexpensedatasetresponse.md), error**


## UpdateExpenseDataset

Update an expense transaction

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.UpdateExpenseDataset(ctx, operations.UpdateExpenseDatasetRequest{
        UpdateExpenseRequest: &shared.UpdateExpenseRequest{
            ContactRef: &shared.ContactRef{
                ContactType: shared.ContactRefContactTypeSupplier.ToPointer(),
                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
            },
            Currency: codatsyncexpenses.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("5448.83"),
            IssueDate: "2022-06-28T00:00:00.000Z",
            Lines: []shared.ExpenseTransactionLine{
                shared.ExpenseTransactionLine{
                    AccountRef: shared.RecordRef{
                        ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    NetAmount: *types.MustNewDecimalFromString("110.42"),
                    TaxAmount: *types.MustNewDecimalFromString("14.43"),
                    TaxRateRef: &shared.RecordRef{
                        ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    TrackingRefs: []shared.RecordRef{
                        shared.RecordRef{
                            ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                        },
                    },
                },
            },
            MerchantName: codatsyncexpenses.String("Amazon UK"),
            Notes: codatsyncexpenses.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
            Type: "illum",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateExpenseDataset202ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateExpenseDatasetRequest](../../models/operations/updateexpensedatasetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.UpdateExpenseDatasetResponse](../../models/operations/updateexpensedatasetresponse.md), error**


## UploadAttachment

Creates an attachment in the accounting software against the given transactionId

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.UploadAttachment(ctx, operations.UploadAttachmentRequest{
        RequestBody: &operations.UploadAttachmentRequestBody{
            Content: []byte("vel"),
            RequestBody: "error",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SyncID: "6fb40d5e-b13e-11ed-afa1-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UploadAttachmentRequest](../../models/operations/uploadattachmentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.UploadAttachmentResponse](../../models/operations/uploadattachmentresponse.md), error**

