# ConfigureCustomDataTypeResponse


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                                                                                                                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                     | HTTP response content type for this operation                                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `CustomDataTypeConfiguration`                                                                                                                                                                                                                                                                                                                                                                                          | [*shared.CustomDataTypeConfiguration](../../../pkg/models/shared/customdatatypeconfiguration.md)                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                     | OK                                                                                                                                                                                                                                                                                                                                                                                                                     | {<br/>"dataSource": "api/purchaseOrders?$filter=currencyCode eq 'NOK'",<br/>"requiredData": {<br/>"currencyCode": "$[*].currencyCode",<br/>"id": "$[*].id",<br/>"number": "$[*].number",<br/>"orderDate": "$[*].orderDate",<br/>"totalAmountExcludingTax": "$[*].totalAmountExcludingTax",<br/>"totalTaxAmount": "$[*].totalTaxAmount",<br/>"vendorName": "$[*].number"<br/>},<br/>"keyBy": [<br/>"$[*].id"<br/>],<br/>"sourceModifiedDate": [<br/>"$[*].lastModifiedDateTime"<br/>]<br/>} |
| `StatusCode`                                                                                                                                                                                                                                                                                                                                                                                                           | *int*                                                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                     | HTTP response status code for this operation                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `RawResponse`                                                                                                                                                                                                                                                                                                                                                                                                          | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                     | Raw HTTP response; suitable for custom response parsing                                                                                                                                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                                                                                                                                                        |