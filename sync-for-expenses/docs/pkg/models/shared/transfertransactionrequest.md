# TransferTransactionRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Date`                                            | *string*                                          | :heavy_check_mark:                                | N/A                                               | 2022-10-23 00:00:00 +0000 UTC                     |
| `Description`                                     | **string*                                         | :heavy_minus_sign:                                | Any private, company notes about the transaction. | Transfer from bank account Y to bank account Z    |
| `From`                                            | [shared.From](../../../pkg/models/shared/from.md) | :heavy_check_mark:                                | N/A                                               |                                                   |
| `To`                                              | [shared.To](../../../pkg/models/shared/to.md)     | :heavy_check_mark:                                | N/A                                               |                                                   |