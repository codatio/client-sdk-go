# InitiateSyncResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `ErrorMessage`                                                | [*shared.ErrorMessage](../../models/shared/errormessage.md)   | :heavy_minus_sign:                                            | If model is incorrect                                         |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |
| `SyncInitiated`                                               | [*shared.SyncInitiated](../../models/shared/syncinitiated.md) | :heavy_minus_sign:                                            | Returns the newly created SyncId                              |