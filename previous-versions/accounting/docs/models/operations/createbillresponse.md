# CreateBillResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `CreateBillResponse`                                                    | [*shared.CreateBillResponse](../../models/shared/createbillresponse.md) | :heavy_minus_sign:                                                      | Success                                                                 |
| `ErrorMessage`                                                          | [*shared.ErrorMessage](../../models/shared/errormessage.md)             | :heavy_minus_sign:                                                      | The request made is not valid.                                          |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |