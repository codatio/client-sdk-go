# GenerateExcelReportResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ContentType`                                             | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `ExcelStatus`                                             | [*shared.ExcelStatus](../../models/shared/excelstatus.md) | :heavy_minus_sign:                                        | OK                                                        |
| `StatusCode`                                              | *int*                                                     | :heavy_check_mark:                                        | N/A                                                       |
| `RawResponse`                                             | [*http.Response](https://pkg.go.dev/net/http#Response)    | :heavy_minus_sign:                                        | N/A                                                       |
| `Schema`                                                  | [*shared.Schema](../../models/shared/schema.md)           | :heavy_minus_sign:                                        | Your API request was not properly authorized.             |