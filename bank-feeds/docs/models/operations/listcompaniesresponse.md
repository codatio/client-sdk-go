# ListCompaniesResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Companies`                                            | [*shared.Companies](../../models/shared/companies.md)  | :heavy_minus_sign:                                     | OK                                                     |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `Schema`                                               | [*shared.Schema](../../models/shared/schema.md)        | :heavy_minus_sign:                                     | Your `query` parameter was not correctly formed        |