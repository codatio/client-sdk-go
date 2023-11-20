# GetCommerceCustomerResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CommerceCustomer`                                                         | [*shared.CommerceCustomer](../../../pkg/models/shared/commercecustomer.md) | :heavy_minus_sign:                                                         | OK                                                                         |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | HTTP response content type for this operation                              |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | HTTP response status code for this operation                               |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_check_mark:                                                         | Raw HTTP response; suitable for custom response parsing                    |