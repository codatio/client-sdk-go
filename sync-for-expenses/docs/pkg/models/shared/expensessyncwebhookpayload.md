# ExpensesSyncWebhookPayload


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ReferenceCompany`                                                         | [*shared.CompanyReference](../../../pkg/models/shared/companyreference.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `SyncID`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | Unique identifier of the sync.                                             |
| `Transactions`                                                             | [][shared.Transaction](../../../pkg/models/shared/transaction.md)          | :heavy_minus_sign:                                                         | N/A                                                                        |