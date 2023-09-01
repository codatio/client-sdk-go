# ReportLine


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `AccountID`                                                                    | **string*                                                                      | :heavy_minus_sign:                                                             | Identifier for the account, unique for the company in the accounting platform. |
| `Items`                                                                        | [][ReportLine](../../models/shared/reportline.md)                              | :heavy_minus_sign:                                                             | An array of ReportLine items.                                                  |
| `Name`                                                                         | **string*                                                                      | :heavy_minus_sign:                                                             | Name of the report line item.                                                  |
| `Value`                                                                        | *float64*                                                                      | :heavy_check_mark:                                                             | Numerical value of the line item.                                              |