# AccountsPayableTracking

Categories, and a project and customer, against which the item is tracked.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `CategoryRefs`                                                                       | [][shared.TrackingCategoryRef](../../../pkg/models/shared/trackingcategoryref.md)    | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `CustomerRef`                                                                        | [*shared.AccountingCustomerRef](../../../pkg/models/shared/accountingcustomerref.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `IsBilledTo`                                                                         | [shared.BilledToType](../../../pkg/models/shared/billedtotype.md)                    | :heavy_check_mark:                                                                   | Defines if the invoice or credit note is billed/rebilled to a project or customer.   |
| `IsRebilledTo`                                                                       | [shared.BilledToType](../../../pkg/models/shared/billedtotype.md)                    | :heavy_check_mark:                                                                   | Defines if the invoice or credit note is billed/rebilled to a project or customer.   |
| `ProjectRef`                                                                         | [*shared.ProjectRef](../../../pkg/models/shared/projectref.md)                       | :heavy_minus_sign:                                                                   | N/A                                                                                  |