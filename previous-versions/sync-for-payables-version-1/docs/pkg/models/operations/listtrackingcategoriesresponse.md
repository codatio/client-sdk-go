# ListTrackingCategoriesResponse


## Fields

| Field                                                                                                                                                                                                         | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   | Example                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                            | HTTP response content type for this operation                                                                                                                                                                 |                                                                                                                                                                                                               |
| `StatusCode`                                                                                                                                                                                                  | *int*                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                            | HTTP response status code for this operation                                                                                                                                                                  |                                                                                                                                                                                                               |
| `RawResponse`                                                                                                                                                                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                            | Raw HTTP response; suitable for custom response parsing                                                                                                                                                       |                                                                                                                                                                                                               |
| `TrackingCategories`                                                                                                                                                                                          | [*shared.TrackingCategories](../../../pkg/models/shared/trackingcategories.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                            | Success                                                                                                                                                                                                       | {<br/>"pageNumber": 1,<br/>"pageSize": 10,<br/>"totalResults": 1,<br/>"_links": {<br/>"self": {<br/>"href": "/companies/{id}/data/{dataType}"<br/>},<br/>"current": {<br/>"href": "/companies/{id}/data/{dataType}?page=1\u0026pageSize=10"<br/>}<br/>}<br/>} |