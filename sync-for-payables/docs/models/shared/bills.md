# Bills


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Links`                                     | [Links](../../models/shared/links.md)       | :heavy_check_mark:                          | N/A                                         |
| `PageNumber`                                | *int64*                                     | :heavy_check_mark:                          | Current page number.                        |
| `PageSize`                                  | *int64*                                     | :heavy_check_mark:                          | Number of items to return in results array. |
| `Results`                                   | [][Bill](../../models/shared/bill.md)       | :heavy_minus_sign:                          | N/A                                         |
| `TotalResults`                              | *int64*                                     | :heavy_check_mark:                          | Total number of items.                      |