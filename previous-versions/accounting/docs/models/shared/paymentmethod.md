# PaymentMethod

> View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.

## Overview

A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/accounting-api#/schemas/BillPayment) and [Payments](https://docs.codat.io/accounting-api#/schemas/Payment).


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | **string*                                                          | :heavy_minus_sign:                                                 | Unique identifier for the payment method.                          |                                                                    |
| `Metadata`                                                         | [*Metadata](../../models/shared/metadata.md)                       | :heavy_minus_sign:                                                 | N/A                                                                |                                                                    |
| `ModifiedDate`                                                     | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                | 2022-10-23T00:00:00.000Z                                           |
| `Name`                                                             | **string*                                                          | :heavy_minus_sign:                                                 | Name of the payment method.                                        |                                                                    |
| `SourceModifiedDate`                                               | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                | 2022-10-23T00:00:00.000Z                                           |
| `Status`                                                           | [*PaymentMethodStatus](../../models/shared/paymentmethodstatus.md) | :heavy_minus_sign:                                                 | Status of the Payment Method.                                      |                                                                    |
| `Type`                                                             | [*PaymentMethodType](../../models/shared/paymentmethodtype.md)     | :heavy_minus_sign:                                                 | Method of payment.                                                 |                                                                    |