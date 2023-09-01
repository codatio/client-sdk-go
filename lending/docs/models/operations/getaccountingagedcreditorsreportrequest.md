# GetAccountingAgedCreditorsReportRequest


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `CompanyID`                                 | *string*                                    | :heavy_check_mark:                          | N/A                                         | 8a210b68-6988-11ed-a1eb-0242ac120002        |
| `NumberOfPeriods`                           | **int*                                      | :heavy_minus_sign:                          | Number of periods to include in the report. | 12                                          |
| `PeriodLengthDays`                          | **int*                                      | :heavy_minus_sign:                          | The length of period in days.               | 30                                          |
| `ReportDate`                                | [*types.Date](../../types/date.md)          | :heavy_minus_sign:                          | Date the report is generated up to.         | 2022-12-31                                  |