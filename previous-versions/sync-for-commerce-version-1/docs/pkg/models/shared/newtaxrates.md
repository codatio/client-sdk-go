# NewTaxRates


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AccountingTaxRateOptions`                                              | [][shared.Option](../../../pkg/models/shared/option.md)                 | :heavy_minus_sign:                                                      | Array of accounting tax rate options.                                   |
| `CommerceTaxRateOptions`                                                | [][shared.Option](../../../pkg/models/shared/option.md)                 | :heavy_minus_sign:                                                      | Array of tax component options.                                         |
| `DefaultZeroTaxRateOptions`                                             | [][shared.Option](../../../pkg/models/shared/option.md)                 | :heavy_minus_sign:                                                      | Default zero tax rate selected for sync.                                |
| `SelectedDefaultZeroTaxRateID`                                          | **string*                                                               | :heavy_minus_sign:                                                      | Default tax rate selected for sync.                                     |
| `TaxRateMappings`                                                       | [][shared.TaxRateMapping](../../../pkg/models/shared/taxratemapping.md) | :heavy_minus_sign:                                                      | Array of tax component to rate mapppings.                               |