# CommercePaymentMethod

A Payment Method represents the payment method(s) used to make payments.

Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-paymentMethods) for this data type.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ID`                                                   | *string*                                               | :heavy_check_mark:                                     | A unique, persistent identifier for this record        | 13d946f0-c5d5-42bc-b092-97ece17923ab                   |
| `ModifiedDate`                                         | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    | 2022-10-23T00:00:00.000Z                               |
| `Name`                                                 | **string*                                              | :heavy_minus_sign:                                     | The name of the PaymentMethod                          | Alipay                                                 |
| `SourceModifiedDate`                                   | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    | 2022-10-23T00:00:00.000Z                               |
| `Status`                                               | [*shared.Status](../../../pkg/models/shared/status.md) | :heavy_minus_sign:                                     | Status of the Payment Method.                          |                                                        |