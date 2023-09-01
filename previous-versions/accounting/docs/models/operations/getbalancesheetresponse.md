# GetBalanceSheetResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `BalanceSheet`                                                | [*shared.BalanceSheet1](../../models/shared/balancesheet1.md) | :heavy_minus_sign:                                            | Success                                                       |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `ErrorMessage`                                                | [*shared.ErrorMessage](../../models/shared/errormessage.md)   | :heavy_minus_sign:                                            | Your API request was not properly authorized.                 |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |