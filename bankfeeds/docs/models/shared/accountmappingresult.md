# AccountMappingResult

The result from POSTing a Bank Account mapping.


## Fields

| Field                                  | Type                                   | Required                               | Description                            |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `Error`                                | **string*                              | :heavy_minus_sign:                     | Error returned during the post request |
| `Status`                               | **string*                              | :heavy_minus_sign:                     | Status of the POST request             |
| `SourceAccountID`                      | **string*                              | :heavy_minus_sign:                     | Unique ID for the source account       |
| `TargetAccountID`                      | **string*                              | :heavy_minus_sign:                     | Unique ID for the target account       |