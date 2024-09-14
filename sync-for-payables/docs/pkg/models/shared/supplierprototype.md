# SupplierPrototype


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `SupplierName`                                                                         | *string*                                                                               | :heavy_check_mark:                                                                     | Name of the supplier as recorded in the accounting system, typically the company name. |                                                                                        |
| `ContactName`                                                                          | **string*                                                                              | :heavy_minus_sign:                                                                     | Name of the main contact for the supplier.                                             |                                                                                        |
| `EmailAddress`                                                                         | **string*                                                                              | :heavy_minus_sign:                                                                     | Email address that the supplier may be contacted on.                                   |                                                                                        |
| `Phone`                                                                                | **string*                                                                              | :heavy_minus_sign:                                                                     | Phone number that the supplier may be contacted on.                                    | +44 25691 154789                                                                       |
| `Addresses`                                                                            | [][shared.Address](../../../pkg/models/shared/address.md)                              | :heavy_minus_sign:                                                                     | An array of Addresses.                                                                 |                                                                                        |
| `Status`                                                                               | [shared.SupplierStatus](../../../pkg/models/shared/supplierstatus.md)                  | :heavy_check_mark:                                                                     | Status of the supplier.                                                                |                                                                                        |
| `Balance`                                                                              | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big)                | :heavy_minus_sign:                                                                     | Amount outstanding against the supplier.                                               |                                                                                        |
| `DefaultCurrency`                                                                      | **string*                                                                              | :heavy_minus_sign:                                                                     | Default currency the supplier's transactional data is recorded in.                     |                                                                                        |