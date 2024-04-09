# Attachments
(*Attachments*)

## Overview

Upload attachmens to expenses, transfers and reimbursable expense transactions.

### Available Operations

* [Upload](#upload) - Upload attachment

## Upload

The *Upload attachment* endpoint uploads an attachment in the accounting software against the given transactionId. 

[Expense transactions](https://docs.codat.io/sync-for-expenses-api#/schemas/ExpenseTransaction) represent transactions made with a company debit or credit card. 

**Integration-specific behaviour**

Each accounting software supports different file formats and sizes.

| Integration | File Size | File Extension                                                                                                      |  
|-------------|-------------|--------------------------------------------------------------------------------------------------------------|
| Xero | 4MB  | 7Z, BMP, CSV, DOC, DOCX, EML, GIF, JPEG, JPG, KEYNOTE, MSG, NUMBERS, ODF, ODS, ODT, PAGES, PDF, PNG, PPT, PPTX, RAR, RTF, TIF, TIFF, TXT, XLS, XLSX, ZIP |
| QuickBooks Online | 100MB | AI, CSV, DOC, DOCX, EPS, GIF, JPEG, JPG, ODS, PAGES, PDF, PNG, RTF, TIF, TXT, XLS, XLSX, XML  |
| NetSuite | 100MB | BMP, CSV, XLS, XLSX, JSON, PDF, PJPG, PJPEG, PNG, TXT, SVG, TIF, TIFF, DOC, DOCX, ZIP |
| Dynamics 365 Business Central | 350 MB | Dynamics do not explicitly outline which file types are supported but they do state <a className="external" href="https://learn.microsoft.com/en-gb/dynamics365/business-central/ui-how-add-link-to-record#to-attach-a-file-to-a-purchase-invoice" target="_blank">here</a> that "You can attach any type of file, such as text, image, or video files". |

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"context"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/operations"
	"log"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
    )

    ctx := context.Background()
    res, err := s.Attachments.Upload(ctx, operations.UploadExpenseAttachmentRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UploadExpenseAttachmentRequest](../../pkg/models/operations/uploadexpenseattachmentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.UploadExpenseAttachmentResponse](../../pkg/models/operations/uploadexpenseattachmentresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
