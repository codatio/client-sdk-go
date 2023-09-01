# ListTaxComponentsResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `ErrorMessage`                                                | [*shared.ErrorMessage](../../models/shared/errormessage.md)   | :heavy_minus_sign:                                            | Your `query` parameter was not correctly formed               |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |
| `TaxComponents`                                               | [*shared.TaxComponents](../../models/shared/taxcomponents.md) | :heavy_minus_sign:                                            | OK                                                            |