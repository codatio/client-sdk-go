# ServiceChargeTaxComponentAllocation


## Fields

| Field                                                                                                                            | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `Rate`                                                                                                                           | **float64*                                                                                                                       | :heavy_minus_sign:                                                                                                               | Tax amount on order line sale as available from source commerce platform.                                                        |
| `TaxComponentRef`                                                                                                                | [*ServiceChargeTaxComponentAllocationTaxComponentRef](../../models/shared/servicechargetaxcomponentallocationtaxcomponentref.md) | :heavy_minus_sign:                                                                                                               | Taxes rates reference object depending on the rates being available on source commerce package.                                  |