# PushOptionChoice


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Description`                                                          | **string*                                                              | :heavy_minus_sign:                                                     | A description of the property.                                         |
| `DisplayName`                                                          | **string*                                                              | :heavy_minus_sign:                                                     | The property's display name.                                           |
| `Required`                                                             | **bool*                                                                | :heavy_minus_sign:                                                     | The property is required if `True`.                                    |
| `Type`                                                                 | [*shared.PushOptionType](../../../pkg/models/shared/pushoptiontype.md) | :heavy_minus_sign:                                                     | The option type.                                                       |
| `Value`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | Allowed value for field.                                               |