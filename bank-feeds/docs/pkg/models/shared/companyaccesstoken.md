# CompanyAccessToken

Details of the access token provisioned for a company.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AccessToken`                                         | *string*                                              | :heavy_check_mark:                                    | The access token for the company.                     |                                                       |
| `ExpiresIn`                                           | *int64*                                               | :heavy_check_mark:                                    | The number of seconds until the access token expires. | 86400                                                 |
| `TokenType`                                           | *string*                                              | :heavy_check_mark:                                    | The type of token.                                    | Bearer                                                |