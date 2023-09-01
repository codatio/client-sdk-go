# ListBillAttachmentsResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AttachmentsDataset`                                                    | [*shared.AttachmentsDataset](../../models/shared/attachmentsdataset.md) | :heavy_minus_sign:                                                      | Success                                                                 |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Schema`                                                                | [*shared.Schema](../../models/shared/schema.md)                         | :heavy_minus_sign:                                                      | Your API request was not properly authorized.                           |