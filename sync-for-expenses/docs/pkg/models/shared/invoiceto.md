# InvoiceTo

Unique identifier of the customer the expense is billable to. The invoiceTo object is currently only supported for QBO.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `DataType`                                                                   | [*shared.InvoiceToDataType](../../../pkg/models/shared/invoicetodatatype.md) | :heavy_minus_sign:                                                           | The type of contact.                                                         | customers                                                                    |
| `ID`                                                                         | **string*                                                                    | :heavy_minus_sign:                                                           | identifier of customer.                                                      | 80000002-1674552702                                                          |