# GetSyncStatusResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `BadRequest`                                            | *interface{}*                                           | :heavy_minus_sign:                                      | Bad Request                                             |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `NotFound`                                              | *interface{}*                                           | :heavy_minus_sign:                                      | Not Found                                               |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | N/A                                                     |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | N/A                                                     |
| `SyncStatus`                                            | [*shared.SyncStatus](../../models/shared/syncstatus.md) | :heavy_minus_sign:                                      | Success                                                 |