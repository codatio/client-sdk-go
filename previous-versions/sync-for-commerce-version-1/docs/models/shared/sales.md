# Sales


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AdditionalProperties`                                           | map[string]*interface{}*                                         | :heavy_minus_sign:                                               | N/A                                                              |
| `Accounts`                                                       | map[string][ConfigAccount](../../models/shared/configaccount.md) | :heavy_minus_sign:                                               | N/A                                                              |
| `Grouping`                                                       | [*Grouping](../../models/shared/grouping.md)                     | :heavy_minus_sign:                                               | N/A                                                              |
| `InvoiceStatus`                                                  | [*InvoiceStatus1](../../models/shared/invoicestatus1.md)         | :heavy_minus_sign:                                               | N/A                                                              |
| `NewTaxRates`                                                    | [*NewTaxRates](../../models/shared/newtaxrates.md)               | :heavy_minus_sign:                                               | N/A                                                              |
| `SalesCustomer`                                                  | [*Customer](../../models/shared/customer.md)                     | :heavy_minus_sign:                                               | N/A                                                              |
| `SyncSales`                                                      | **bool*                                                          | :heavy_minus_sign:                                               | Boolean indicator for syncing sales.                             |
| `TaxRates`                                                       | map[string][TaxRateAmount](../../models/shared/taxrateamount.md) | :heavy_minus_sign:                                               | N/A                                                              |