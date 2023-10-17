# ListAccountBalancesResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AccountBalances`                                                 | [*shared.AccountBalances](../../models/shared/accountbalances.md) | :heavy_minus_sign:                                                | Success                                                           |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | HTTP response content type for this operation                     |
| `ErrorMessage`                                                    | [*shared.ErrorMessage](../../models/shared/errormessage.md)       | :heavy_minus_sign:                                                | Your `query` parameter was not correctly formed                   |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | HTTP response status code for this operation                      |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | Raw HTTP response; suitable for custom response parsing           |