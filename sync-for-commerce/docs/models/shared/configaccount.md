# ConfigAccount

G/L account object for configuration.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `AccountOptions`                                         | [][AccountOption](../../models/shared/accountoption.md)  | :heavy_minus_sign:                                       | Object containing account options.                       |
| `DescriptionText`                                        | **string*                                                | :heavy_minus_sign:                                       | Descriprtive text for sales configuration section.       |
| `LabelText`                                              | **string*                                                | :heavy_minus_sign:                                       | Label text for sales configuration section.              |
| `Required`                                               | **bool*                                                  | :heavy_minus_sign:                                       | Required section to be configured for sync.              |
| `SelectedAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | Selected account id from the list of available accounts. |