# ListIntegrationsResponse


## Fields

| Field                                                                                                                                                                                | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          | Example                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ContentType`                                                                                                                                                                        | *string*                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                   | HTTP response content type for this operation                                                                                                                                        |                                                                                                                                                                                      |
| `Integrations`                                                                                                                                                                       | [*shared.Integrations](../../../pkg/models/shared/integrations.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                   | Success                                                                                                                                                                              | {"_links":{"pageNumber":1,"pageSize":10,"totalResults":1,"self":{"href":"/companies/{id}/data/{dataType}"},"current":{"href":"/companies/{id}/data/{dataType}?page=1&pageSize=10"}}} |
| `StatusCode`                                                                                                                                                                         | *int*                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | HTTP response status code for this operation                                                                                                                                         |                                                                                                                                                                                      |
| `RawResponse`                                                                                                                                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                                               | :heavy_check_mark:                                                                                                                                                                   | Raw HTTP response; suitable for custom response parsing                                                                                                                              |                                                                                                                                                                                      |