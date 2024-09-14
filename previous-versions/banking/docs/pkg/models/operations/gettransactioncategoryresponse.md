# GetTransactionCategoryResponse


## Fields

| Field                                                                                                                                                                                        | Type                                                                                                                                                                                         | Required                                                                                                                                                                                     | Description                                                                                                                                                                                  | Example                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                                                                                                | *string*                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                           | HTTP response content type for this operation                                                                                                                                                |                                                                                                                                                                                              |
| `StatusCode`                                                                                                                                                                                 | *int*                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                           | HTTP response status code for this operation                                                                                                                                                 |                                                                                                                                                                                              |
| `RawResponse`                                                                                                                                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                           | Raw HTTP response; suitable for custom response parsing                                                                                                                                      |                                                                                                                                                                                              |
| `TransactionCategory`                                                                                                                                                                        | [*shared.TransactionCategory](../../../pkg/models/shared/transactioncategory.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                                           | Success                                                                                                                                                                                      | {<br/>"id": "auto-and-transport",<br/>"name": "Auto \u0026 Transport",<br/>"hasChildren": true,<br/>"status": "Active",<br/>"modifiedDate": "2022-05-23T16:32:50",<br/>"sourceModifiedDate": "2021-04-24T07:59:10"<br/>} |