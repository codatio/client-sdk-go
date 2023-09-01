# ExpenseTransactionLine


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `AccountRef`                                    | [RecordRef](../../models/shared/recordref.md)   | :heavy_check_mark:                              | N/A                                             |                                                 |
| `NetAmount`                                     | *float64*                                       | :heavy_check_mark:                              | Amount of the line, exclusive of tax.           | 110.42                                          |
| `TaxAmount`                                     | *float64*                                       | :heavy_check_mark:                              | Amount of tax for the line.                     | 14.43                                           |
| `TaxRateRef`                                    | [*RecordRef](../../models/shared/recordref.md)  | :heavy_minus_sign:                              | N/A                                             |                                                 |
| `TrackingRefs`                                  | [][RecordRef](../../models/shared/recordref.md) | :heavy_minus_sign:                              | N/A                                             |                                                 |