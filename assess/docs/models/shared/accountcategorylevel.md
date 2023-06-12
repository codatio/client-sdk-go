# AccountCategoryLevel

An object containing an ordered list of account category levels.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Confidence`                                                                                 | **float64*                                                                                   | :heavy_minus_sign:                                                                           | Confidence level of the category. This will only be populated where `status` is `Suggested`. |
| `LevelName`                                                                                  | **string*                                                                                    | :heavy_minus_sign:                                                                           | Account category name.                                                                       |