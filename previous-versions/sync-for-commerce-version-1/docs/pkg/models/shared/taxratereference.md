# TaxRateReference

Reference to the tax rate to which the line item is linked.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `EffectiveTaxRate`                                                      | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big) | :heavy_minus_sign:                                                      | Applicable tax rate.                                                    |
| `ID`                                                                    | **string*                                                               | :heavy_minus_sign:                                                      | Unique identifier for the tax rate in the accounting software.          |
| `Name`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Name of the tax rate in the accounting software.                        |