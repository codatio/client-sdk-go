# NewCompanySynchronizedWebhook

Webhook request body to notify that a new company has successfully synchronized at least one dataType for the first time.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `AlertID`                                   | **string*                                   | :heavy_minus_sign:                          | Unique identifier of the webhook event.     |                                             |
| `CompanyID`                                 | **string*                                   | :heavy_minus_sign:                          | Unique identifier for your SMB in Codat.    | 8a210b68-6988-11ed-a1eb-0242ac120002        |
| `Message`                                   | **string*                                   | :heavy_minus_sign:                          | A human readable message about the webhook. |                                             |
| `RuleID`                                    | **string*                                   | :heavy_minus_sign:                          | Unique identifier for the rule.             |                                             |
| `Type`                                      | **string*                                   | :heavy_minus_sign:                          | The type of rule.                           |                                             |