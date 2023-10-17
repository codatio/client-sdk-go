# DatasetStatusChangedErrorWebhookData


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `DataType`                                                 | [*DataType](../../models/shared/datatype.md)               | :heavy_minus_sign:                                         | Available Data types                                       | invoices                                                   |
| `DatasetID`                                                | **string*                                                  | :heavy_minus_sign:                                         | Unique identifier for the dataset that completed its sync. |                                                            |
| `DatasetStatus`                                            | **string*                                                  | :heavy_minus_sign:                                         | The current status of the dataset's sync.                  |                                                            |