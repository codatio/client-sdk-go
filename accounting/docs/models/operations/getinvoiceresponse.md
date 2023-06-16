# GetInvoiceResponse


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ContentType`                                                                            | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Invoice`                                                                                | [*shared.Invoice](../../models/shared/invoice.md)                                        | :heavy_minus_sign:                                                                       | Success                                                                                  |
| `StatusCode`                                                                             | *int*                                                                                    | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `RawResponse`                                                                            | [*http.Response](https://pkg.go.dev/net/http#Response)                                   | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `GetInvoice409ApplicationJSONObject`                                                     | [*GetInvoice409ApplicationJSON](../../models/operations/getinvoice409applicationjson.md) | :heavy_minus_sign:                                                                       | The data type's dataset has not been requested or is still syncing.                      |
| `Schema`                                                                                 | [*shared.Schema](../../models/shared/schema.md)                                          | :heavy_minus_sign:                                                                       | Your API request was not properly authorized.                                            |