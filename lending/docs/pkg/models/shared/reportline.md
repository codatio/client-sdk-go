# ReportLine


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `AccountID`                                                                    | **string*                                                                      | :heavy_minus_sign:                                                             | Identifier for the account, unique for the company in the accounting platform. |
| `Items`                                                                        | [][shared.ReportLine](../../../pkg/models/shared/reportline.md)                | :heavy_minus_sign:                                                             | An array of ReportLine items.                                                  |
| `Name`                                                                         | **string*                                                                      | :heavy_minus_sign:                                                             | Name of the report line item.                                                  |
| `Value`                                                                        | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big)        | :heavy_check_mark:                                                             | Numerical value of the line item.                                              |