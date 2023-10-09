# BillCreditNoteLineItemTracking

Categories, and a project and customer, against which the item is tracked.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `CategoryRefs`                                                                     | [][TrackingCategoryRef](../../models/shared/trackingcategoryref.md)                | :heavy_check_mark:                                                                 | N/A                                                                                |
| `CustomerRef`                                                                      | [*AccountingCustomerRef](../../models/shared/accountingcustomerref.md)             | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `IsBilledTo`                                                                       | [BilledToType](../../models/shared/billedtotype.md)                                | :heavy_check_mark:                                                                 | Defines if the invoice or credit note is billed/rebilled to a project or customer. |
| `IsRebilledTo`                                                                     | [BilledToType](../../models/shared/billedtotype.md)                                | :heavy_check_mark:                                                                 | Defines if the invoice or credit note is billed/rebilled to a project or customer. |
| `ProjectRef`                                                                       | [*ProjectRef](../../models/shared/projectref.md)                                   | :heavy_minus_sign:                                                                 | N/A                                                                                |