# Branding

Success


## Fields

| Field                                                                                                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Button`                                                                                                                                                                                                                                                                                                | [*BrandingButton](../../models/shared/brandingbutton.md)                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                      | Button branding references.                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                         |
| `Logo`                                                                                                                                                                                                                                                                                                  | [*BrandingLogo](../../models/shared/brandinglogo.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                      | Logo branding references.                                                                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                                                                         |
| `SourceID`                                                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                      | A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`. | 35b92968-9851-4095-ad60-395c95cbcba4                                                                                                                                                                                                                                                                    |