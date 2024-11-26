# CreateBankAccountMappingResponse


## Fields

| Field                                                                                                                                               | Type                                                                                                                                                | Required                                                                                                                                            | Description                                                                                                                                         | Example                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `BankFeedBankAccountMappingResponse`                                                                                                                | [*shared.BankFeedBankAccountMappingResponse](../../../pkg/models/shared/bankfeedbankaccountmappingresponse.md)                                      | :heavy_minus_sign:                                                                                                                                  | Success                                                                                                                                             | {<br/>"sourceAccountId": "acc-002",<br/>"targetAccountId": "account-081",<br/>"status": "Failed",<br/>"error": "A feed connection already exists to this account"<br/>} |
| `ContentType`                                                                                                                                       | *string*                                                                                                                                            | :heavy_check_mark:                                                                                                                                  | HTTP response content type for this operation                                                                                                       |                                                                                                                                                     |
| `StatusCode`                                                                                                                                        | *int*                                                                                                                                               | :heavy_check_mark:                                                                                                                                  | HTTP response status code for this operation                                                                                                        |                                                                                                                                                     |
| `RawResponse`                                                                                                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                              | :heavy_check_mark:                                                                                                                                  | Raw HTTP response; suitable for custom response parsing                                                                                             |                                                                                                                                                     |