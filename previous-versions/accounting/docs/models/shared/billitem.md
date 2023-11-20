# BillItem

Item details that are only for bills.


## Fields

| Field                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                | Required                                                                                                                                                                                                                                                                                            | Description                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccountRef`                                                                                                                                                                                                                                                                                        | [*AccountRef](../../models/shared/accountref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.                                                                                                                                               |
| `Description`                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Short description of the product or service that has been bought by the customer.                                                                                                                                                                                                                   |
| `TaxRateRef`                                                                                                                                                                                                                                                                                        | [*TaxRateRef](../../models/shared/taxrateref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.<br/><br/>Found on:<br/><br/>- Bill line items<br/>- Bill Credit Note line items<br/>- Credit Note line items<br/>- Direct incomes line items<br/>- Invoice line items<br/>- Items |
| `UnitPrice`                                                                                                                                                                                                                                                                                         | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big)                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Unit price of the product or service.                                                                                                                                                                                                                                                               |