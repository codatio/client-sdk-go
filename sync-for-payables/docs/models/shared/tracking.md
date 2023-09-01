# Tracking

Categories, and a project and customer, against which the item is tracked.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CategoryRefs`                                                      | [][TrackingCategoryRef](../../models/shared/trackingcategoryref.md) | :heavy_check_mark:                                                  | N/A                                                                 |
| `CustomerRef`                                                       | [*TrackingCustomerRef](../../models/shared/trackingcustomerref.md)  | :heavy_minus_sign:                                                  | N/A                                                                 |
| `IsBilledTo`                                                        | [BilledToType](../../models/shared/billedtotype.md)                 | :heavy_check_mark:                                                  | N/A                                                                 |
| `IsRebilledTo`                                                      | [BilledToType](../../models/shared/billedtotype.md)                 | :heavy_check_mark:                                                  | N/A                                                                 |
| `ProjectRef`                                                        | [*TrackingProjectRef](../../models/shared/trackingprojectref.md)    | :heavy_minus_sign:                                                  | N/A                                                                 |