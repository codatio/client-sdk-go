# UpdateProfileSyncSettingsRequestBody

Include a `syncSetting` object for each data type.
`syncFromWindow`, `syncFromUTC` & `monthsToSync` only need to be included if you wish to set a value for them.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ClientID`                                                 | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | 367f7975-267b-439b-90c6-a6040ee680f3                       |
| `OverridesDefaults`                                        | **bool*                                                    | :heavy_minus_sign:                                         | N/A                                                        |                                                            |
| `Settings`                                                 | [][shared.SyncSetting](../../models/shared/syncsetting.md) | :heavy_check_mark:                                         | N/A                                                        |                                                            |