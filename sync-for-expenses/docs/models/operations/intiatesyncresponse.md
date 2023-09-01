# IntiateSyncResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CodatErrorMessage`                                                   | [*shared.CodatErrorMessage](../../models/shared/codaterrormessage.md) | :heavy_minus_sign:                                                    | If model is incorrect                                                 |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `SyncInitiated`                                                       | [*shared.SyncInitiated](../../models/shared/syncinitiated.md)         | :heavy_minus_sign:                                                    | Returns the newly created SyncId                                      |