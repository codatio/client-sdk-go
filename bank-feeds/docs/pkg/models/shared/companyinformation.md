# CompanyInformation

Information about the company from the underlying accounting software.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `BaseCurrency`                                                       | **string*                                                            | :heavy_minus_sign:                                                   | Currency set in the accounting software of the linked company.       |
| `CompanyName`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | Name of the linked company.                                          |
| `Currencies`                                                         | []*string*                                                           | :heavy_minus_sign:                                                   | Array of enabled currencies for the linked company.                  |
| `MulticurrencyEnabled`                                               | **bool*                                                              | :heavy_minus_sign:                                                   | Boolean showing if the organisation has multicurrency enabled.       |
| `PlanType`                                                           | **string*                                                            | :heavy_minus_sign:                                                   | Accounting software subscription type such as Trial, Demo, Standard. |