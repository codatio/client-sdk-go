# CreateAPIKeyResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `APIKeyDetails`                                               | [*shared.APIKeyDetails](../../models/shared/apikeydetails.md) | :heavy_minus_sign:                                            | Success                                                       |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | HTTP response content type for this operation                 |
| `ErrorMessage`                                                | [*shared.ErrorMessage](../../models/shared/errormessage.md)   | :heavy_minus_sign:                                            | Bad Request                                                   |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | HTTP response status code for this operation                  |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | Raw HTTP response; suitable for custom response parsing       |