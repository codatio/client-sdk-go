# WebhookConsumer

﻿A webhook consumer is an HTTP endpoint that developers can configure to subscribe to Codat's supported event types.

See our documentation for more details on [Codat's webhook service](https://docs.codat.io/using-the-api/webhooks/overview).



## Fields

| Field                                                                                                                                                | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          | Example                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CompanyID`                                                                                                                                          | **string*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                   | Unique identifier of the company to indicate company-specific events. The associated webhook consumer will receive events only for the specified ID. | 8a210b68-6988-11ed-a1eb-0242ac120002                                                                                                                 |
| `Disabled`                                                                                                                                           | **bool*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                   | Flag that enables or disables the endpoint from receiving events. Disabled when set to `true`.                                                       |                                                                                                                                                      |
| `EventTypes`                                                                                                                                         | []*string*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                   | An array of event types the webhook consumer subscribes to.                                                                                          |                                                                                                                                                      |
| `ID`                                                                                                                                                 | **string*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                   | Unique identifier for the webhook consumer.                                                                                                          | 8a210b68-6988-11ed-a1eb-0242ac120002                                                                                                                 |
| `URL`                                                                                                                                                | **string*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                   | The URL that will consume webhook events dispatched by Codat.                                                                                        |                                                                                                                                                      |