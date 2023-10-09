# PushFieldValidation


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AdditionalProperties`                                | map[string]*interface{}*                              | :heavy_minus_sign:                                    | N/A                                                   |
| `Details`                                             | *string*                                              | :heavy_check_mark:                                    | Details on the validation issue.                      |
| `Field`                                               | **string*                                             | :heavy_minus_sign:                                    | Field name that resulted in the validation issue.     |
| `Ref`                                                 | **string*                                             | :heavy_minus_sign:                                    | Unique reference identifier for the validation issue. |