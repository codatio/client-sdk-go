# InitiateSyncResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | HTTP response content type for this operation                 |
| `ErrorMessage`                                                | [*shared.ErrorMessage](../../models/shared/errormessage.md)   | :heavy_minus_sign:                                            | If model is incorrect                                         |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | HTTP response status code for this operation                  |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | Raw HTTP response; suitable for custom response parsing       |
| `SyncInitiated`                                               | [*shared.SyncInitiated](../../models/shared/syncinitiated.md) | :heavy_minus_sign:                                            | Returns the newly created SyncId                              |