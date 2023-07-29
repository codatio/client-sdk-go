# ListPurchaseOrdersResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `ErrorMessage`                                                  | [*shared.ErrorMessage](../../models/shared/errormessage.md)     | :heavy_minus_sign:                                              | Your `query` parameter was not correctly formed                 |
| `PurchaseOrders`                                                | [*shared.PurchaseOrders](../../models/shared/purchaseorders.md) | :heavy_minus_sign:                                              | Success                                                         |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |