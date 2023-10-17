# CreateRule

Create an event notification to a URL or list of email addresses based on the given type or condition.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CompanyID`                                               | **string*                                                 | :heavy_minus_sign:                                        | Unique identifier for your SMB in Codat.                  | 8a210b68-6988-11ed-a1eb-0242ac120002                      |
| `Notifiers`                                               | [WebhookNotifier](../../models/shared/webhooknotifier.md) | :heavy_check_mark:                                        | N/A                                                       |                                                           |
| `Type`                                                    | *string*                                                  | :heavy_check_mark:                                        | The type of webhook.                                      |                                                           |