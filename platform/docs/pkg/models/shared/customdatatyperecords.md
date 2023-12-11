# CustomDataTypeRecords

Resulting records pulled from the source platform for a specific custom data type.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `PageNumber`                                                                        | **int64*                                                                            | :heavy_minus_sign:                                                                  | Current page number.                                                                |
| `PageSize`                                                                          | **int64*                                                                            | :heavy_minus_sign:                                                                  | Number of items to return in results array.                                         |
| `Results`                                                                           | [][shared.CustomDataTypeRecord](../../../pkg/models/shared/customdatatyperecord.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `TotalResults`                                                                      | **int64*                                                                            | :heavy_minus_sign:                                                                  | Total number of items.                                                              |