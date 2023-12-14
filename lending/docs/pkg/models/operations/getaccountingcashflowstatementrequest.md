# GetAccountingCashFlowStatementRequest


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `CompanyID`                                       | *string*                                          | :heavy_check_mark:                                | Unique identifier for a company.                  | 8a210b68-6988-11ed-a1eb-0242ac120002              |
| `PeriodLength`                                    | *int*                                             | :heavy_check_mark:                                | Number of months defining the period of interest. | 4                                                 |
| `PeriodsToCompare`                                | *int*                                             | :heavy_check_mark:                                | Number of periods with `periodLength` to compare. | 20                                                |
| `StartMonth`                                      | **string*                                         | :heavy_minus_sign:                                | The month the report starts from.                 | 2022-10-23 00:00:00 +0000 UTC                     |