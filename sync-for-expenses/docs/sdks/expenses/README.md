# Expenses

## Overview

Create expense datasets and upload receipts.

### Available Operations

* [Create](#create) - Create expense transaction
* [Update](#update) - Update expense-transactions
* [UploadAttachment](#uploadattachment) - Upload attachment

## Create

Create an expense transaction

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.Create(ctx, operations.CreateExpenseTransactionRequest{
        CreateExpenseRequest: &shared.CreateExpenseRequest{
            Items: []shared.ExpenseTransaction{
                shared.ExpenseTransaction{
                    BankAccountRef: &shared.ExpenseTransactionBankAccountReference{
                        ID: codatsyncexpenses.String("787dfb37-5707-4dc0-8a86-8d74e4cc78ea"),
                    },
                    ContactRef: &shared.ContactRef{
                        ContactType: shared.ContactRefContactTypeSupplier.ToPointer(),
                        ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
                    },
                    Currency: "GBP",
                    CurrencyRate: types.MustNewDecimalFromString("9840.43"),
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateExpenseTransactionRequest](../../models/operations/createexpensetransactionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.CreateExpenseTransactionResponse](../../models/operations/createexpensetransactionresponse.md), error**


## Update

Update an expense transaction

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
	"github.com/ericlagergren/decimal"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.Update(ctx, operations.UpdateExpenseTransactionRequest{
        UpdateExpenseRequest: &shared.UpdateExpenseRequest{
            BankAccountRef: &shared.UpdateExpenseRequestBankAccountReference{
                ID: codatsyncexpenses.String("787dfb37-5707-4dc0-8a86-8d74e4cc78ea"),
            },
            ContactRef: &shared.ContactRef{
                ContactType: shared.ContactRefContactTypeSupplier.ToPointer(),
                ID: codatsyncexpenses.String("40e3e57c-2322-4898-966c-ca41adfd23fd"),
            },
            Currency: codatsyncexpenses.String("GBP"),
            CurrencyRate: types.MustNewDecimalFromString("2603.41"),
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
            Type: "maxime",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateExpenseTransactionRequest](../../models/operations/updateexpensetransactionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.UpdateExpenseTransactionResponse](../../models/operations/updateexpensetransactionresponse.md), error**


## UploadAttachment

Creates an attachment in the accounting software against the given transactionId

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Expenses.UploadAttachment(ctx, operations.UploadExpenseAttachmentRequest{
        RequestBody: &operations.UploadExpenseAttachmentRequestBody{
            Content: []byte("deleniti"),
            RequestBody: "facilis",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UploadExpenseAttachmentRequest](../../models/operations/uploadexpenseattachmentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.UploadExpenseAttachmentResponse](../../models/operations/uploadexpenseattachmentresponse.md), error**

