# Expenses
(*Expenses*)

## Overview

Create expense datasets and upload receipts.

### Available Operations

* [CreateExpenseDataset](#createexpensedataset) - Create expense-transactions
* [UpdateExpenseDataset](#updateexpensedataset) - Update expense transactions
* [UploadAttachment](#uploadattachment) - Upload attachment

## CreateExpenseDataset

Create an expense transaction

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.CreateExpenseDataset(ctx, operations.CreateExpenseDatasetRequest{
        CreateExpenseRequest: &shared.CreateExpenseRequest{
            Items: []shared.ExpenseTransaction{
                shared.ExpenseTransaction{
                    ContactRef: &shared.ContactRef{
                        ContactType: shared.ContactTypeSupplier.ToPointer(),
                        ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    Currency: "GBP",
                    ID: "4d7c6929-7770-412b-91bb-44d3bc71d111",
                    IssueDate: "2022-10-23T00:00:00.000Z",
                    Lines: []shared.ExpenseTransactionLine{
                        shared.ExpenseTransactionLine{
                            AccountRef: shared.RecordRef{
                                ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            NetAmount: types.MustNewDecimalFromString("110.42"),
                            TaxAmount: types.MustNewDecimalFromString("14.43"),
                            TaxRateRef: &shared.RecordRef{
                                ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                            },
                            TrackingRefs: []shared.RecordRef{
                                shared.RecordRef{
                                    ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                                },
                            },
                        },
                    },
                    MerchantName: syncforexpensesversion1.String("Amazon UK"),
                    Notes: syncforexpensesversion1.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
                    Type: shared.TypePayment,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateExpenseDatasetRequest](../../pkg/models/operations/createexpensedatasetrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateExpenseDatasetResponse](../../pkg/models/operations/createexpensedatasetresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## UpdateExpenseDataset

Update an expense transaction

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.UpdateExpenseDataset(ctx, operations.UpdateExpenseDatasetRequest{
        UpdateExpenseRequest: &shared.UpdateExpenseRequest{
            ContactRef: &shared.ContactRef{
                ContactType: shared.ContactTypeSupplier.ToPointer(),
                ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
            },
            Currency: syncforexpensesversion1.String("GBP"),
            IssueDate: "2022-06-28T00:00:00.000Z",
            Lines: []shared.ExpenseTransactionLine{
                shared.ExpenseTransactionLine{
                    AccountRef: shared.RecordRef{
                        ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    NetAmount: types.MustNewDecimalFromString("110.42"),
                    TaxAmount: types.MustNewDecimalFromString("14.43"),
                    TaxRateRef: &shared.RecordRef{
                        ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    TrackingRefs: []shared.RecordRef{
                        shared.RecordRef{
                            ID: syncforexpensesversion1.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                        },
                    },
                },
            },
            MerchantName: syncforexpensesversion1.String("Amazon UK"),
            Notes: syncforexpensesversion1.String("APPLE.COM/BILL - 09001077498 - Card Ending: 4590"),
            Type: "string",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TransactionID: "336694d8-2dca-4cb5-a28d-3ccb83e55eee",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateExpenseDatasetRequest](../../pkg/models/operations/updateexpensedatasetrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.UpdateExpenseDatasetResponse](../../pkg/models/operations/updateexpensedatasetresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,422,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |

## UploadAttachment

Creates an attachment in the accounting software against the given transactionId

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpensesversion1.New(
        syncforexpensesversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.UploadAttachment(ctx, operations.UploadAttachmentRequest{
        AttachmentUpload: &shared.AttachmentUpload{
            File: shared.CodatFile{
                Content: []byte("0xE3ABc1980E"),
                FileName: "elegant_producer_electric.jpeg",
            },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UploadAttachmentRequest](../../pkg/models/operations/uploadattachmentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.UploadAttachmentResponse](../../pkg/models/operations/uploadattachmentresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
