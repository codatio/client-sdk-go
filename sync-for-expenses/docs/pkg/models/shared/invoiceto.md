# InvoiceTo

Unique identifier for the customer billed for the transaction. The `invoiceTo` object is currently supported only for QuickBooks Online and QuickBooks Desktop.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ID`                                                                 | **string*                                                            | :heavy_minus_sign:                                                   | identifier of customer.                                              | 80000002-1674552702                                                  |
| `Type`                                                               | [*shared.InvoiceToType](../../../pkg/models/shared/invoicetotype.md) | :heavy_minus_sign:                                                   | The type of contact.                                                 | customer                                                             |