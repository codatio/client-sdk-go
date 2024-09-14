# TransactionCategory

The Banking Transaction Categories data type provides a list of hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.

Responses are paged, so you should provide `page` and `pageSize` query parameters in your request.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `HasChildren`                                                                                               | **bool*                                                                                                     | :heavy_minus_sign:                                                                                          | A Boolean indicating whether there are other bank transaction categories beneath this one in the hierarchy. |                                                                                                             |
| `ID`                                                                                                        | *string*                                                                                                    | :heavy_check_mark:                                                                                          | The unique identifier of the bank transaction category.                                                     |                                                                                                             |
| `ModifiedDate`                                                                                              | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | N/A                                                                                                         | 2022-10-23 00:00:00 +0000 UTC                                                                               |
| `Name`                                                                                                      | *string*                                                                                                    | :heavy_check_mark:                                                                                          | The name of the bank transaction category.                                                                  |                                                                                                             |
| `ParentID`                                                                                                  | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | The unique identifier of the parent bank transaction category.                                              |                                                                                                             |
| `SourceModifiedDate`                                                                                        | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | N/A                                                                                                         | 2022-10-23 00:00:00 +0000 UTC                                                                               |
| `Status`                                                                                                    | [*shared.TransactionCategoryStatus](../../../pkg/models/shared/transactioncategorystatus.md)                | :heavy_minus_sign:                                                                                          | The status of the transaction category.                                                                     |                                                                                                             |