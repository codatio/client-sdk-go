# Propertiestracking1

Categories, and a project and customer, against which the item is tracked.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CategoryRefs`                                                      | [][TrackingCategoryRef](../../models/shared/trackingcategoryref.md) | :heavy_check_mark:                                                  | N/A                                                                 |
| `CustomerRef`                                                       | [*CustomerRef](../../models/shared/customerref.md)                  | :heavy_minus_sign:                                                  | N/A                                                                 |
| `IsBilledTo`                                                        | [BilledToType1](../../models/shared/billedtotype1.md)               | :heavy_check_mark:                                                  | N/A                                                                 |
| `IsRebilledTo`                                                      | [BilledToType1](../../models/shared/billedtotype1.md)               | :heavy_check_mark:                                                  | N/A                                                                 |
| `ProjectRef`                                                        | [*ProjectRef](../../models/shared/projectref.md)                    | :heavy_minus_sign:                                                  | N/A                                                                 |