# BankStatementUploadConfiguration

Configuration settings for uploading banking data to Codat


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AccountID`                                                  | **string*                                                    | :heavy_minus_sign:                                           | The ID of the account in the third-party platform            |
| `ProviderID`                                                 | **string*                                                    | :heavy_minus_sign:                                           | TrueLayer provider ID (only required if source is TrueLayer) |
| `Source`                                                     | [*shared.Source](../../../pkg/models/shared/source.md)       | :heavy_minus_sign:                                           | The source of the banking data that determines its format    |