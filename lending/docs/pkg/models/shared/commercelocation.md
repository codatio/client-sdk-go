# CommerceLocation

The Location datatype holds information on the geographic location at which stocks of products may be held, as referenced in the Products data type.

A Location also holds information on geographic locations where orders were placed, as referenced in the Orders data type.

Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-locations) for this data type.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Address`                                                                | [*shared.CommerceAddress](../../../pkg/models/shared/commerceaddress.md) | :heavy_minus_sign:                                                       | N/A                                                                      |                                                                          |
| `ID`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | A unique, persistent identifier for this record                          | 13d946f0-c5d5-42bc-b092-97ece17923ab                                     |
| `ModifiedDate`                                                           | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      | 2022-10-23 00:00:00 +0000 UTC                                            |
| `Name`                                                                   | **string*                                                                | :heavy_minus_sign:                                                       | Name of this location                                                    |                                                                          |
| `SourceModifiedDate`                                                     | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      | 2022-10-23 00:00:00 +0000 UTC                                            |