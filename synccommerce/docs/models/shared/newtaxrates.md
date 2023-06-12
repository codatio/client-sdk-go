# NewTaxRates


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AccountingTaxRateOptions`                                | [][Option](../../models/shared/option.md)                 | :heavy_minus_sign:                                        | Array of accounting tax rate options.                     |
| `CommerceTaxRateOptions`                                  | [][Option](../../models/shared/option.md)                 | :heavy_minus_sign:                                        | Array of tax component options.                           |
| `DefaultZeroTaxRateOptions`                               | [][Option](../../models/shared/option.md)                 | :heavy_minus_sign:                                        | Default zero tax rate selected for sync.                  |
| `SelectedDefaultZeroTaxRateID`                            | **string*                                                 | :heavy_minus_sign:                                        | Default tax rate selected for sync.                       |
| `TaxRateMappings`                                         | [][TaxRateMapping](../../models/shared/taxratemapping.md) | :heavy_minus_sign:                                        | Array of tax component to rate mapppings.                 |