# From


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `AccountRef`                                                                             | [shared.AccountReference](../../../pkg/models/shared/accountreference.md)                | :heavy_check_mark:                                                                       | Reference of the account you are transferring money from.                                |
| `Amount`                                                                                 | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big)                  | :heavy_check_mark:                                                                       | Amount that has been transferred from the account in the native currency of the account. |