# GetAccountsForEnhancedBalanceSheetResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `EnhancedReport`                                                | [*shared.EnhancedReport](../../models/shared/enhancedreport.md) | :heavy_minus_sign:                                              | OK                                                              |
| `ErrorMessage`                                                  | [*shared.ErrorMessage](../../models/shared/errormessage.md)     | :heavy_minus_sign:                                              | Your API request was not properly authorized.                   |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |