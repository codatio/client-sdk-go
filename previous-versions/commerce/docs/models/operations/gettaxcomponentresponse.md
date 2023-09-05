# GetTaxComponentResponse


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ContentType`                                               | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `ErrorMessage`                                              | [*shared.ErrorMessage](../../models/shared/errormessage.md) | :heavy_minus_sign:                                          | Your API request was not properly authorized.               |
| `StatusCode`                                                | *int*                                                       | :heavy_check_mark:                                          | N/A                                                         |
| `RawResponse`                                               | [*http.Response](https://pkg.go.dev/net/http#Response)      | :heavy_minus_sign:                                          | N/A                                                         |
| `TaxComponent`                                              | [*shared.TaxComponent](../../models/shared/taxcomponent.md) | :heavy_minus_sign:                                          | OK                                                          |