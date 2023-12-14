# ProductInventoryLocation


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `LocationRef`                                                           | [*shared.LocationRef](../../../pkg/models/shared/locationref.md)        | :heavy_minus_sign:                                                      | Reference to the geographic location where the order was placed.        |
| `Quantity`                                                              | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big) | :heavy_minus_sign:                                                      | The quantity of stock remaining at location.                            |