# CreateCompanyResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Company`                                              | [*shared.Company](../../models/shared/company.md)      | :heavy_minus_sign:                                     | OK                                                     |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `Schema`                                               | [*shared.Schema](../../models/shared/schema.md)        | :heavy_minus_sign:                                     | The request made is not valid.                         |