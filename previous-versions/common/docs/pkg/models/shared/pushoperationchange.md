# PushOperationChange


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AttachmentID`                                                             | **string*                                                                  | :heavy_minus_sign:                                                         | Unique identifier for the attachment created otherwise null.               |
| `RecordRef`                                                                | [*shared.PushOperationRef](../../../pkg/models/shared/pushoperationref.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Type`                                                                     | [*shared.PushChangeType](../../../pkg/models/shared/pushchangetype.md)     | :heavy_minus_sign:                                                         | Type of change being applied to record in third party platform.            |