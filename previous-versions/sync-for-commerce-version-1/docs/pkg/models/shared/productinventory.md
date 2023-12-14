# ProductInventory

Information about the total inventory as well as the locations inventory is in.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Locations`                                                                                 | [][shared.ProductInventoryLocation](../../../pkg/models/shared/productinventorylocation.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `TotalQuantity`                                                                             | [*decimal.Big](https://pkg.go.dev/github.com/ericlagergren/decimal#Big)                     | :heavy_minus_sign:                                                                          | The total quantity of stock remaining across locations.                                     |