// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PushOperationStatusChangedWebhookData struct {
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey *string `json:"pushOperationKey,omitempty"`
	// The current status of the push operation.
	Status *PushOperationStatus `json:"status,omitempty"`
}

func (o *PushOperationStatusChangedWebhookData) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *PushOperationStatusChangedWebhookData) GetPushOperationKey() *string {
	if o == nil {
		return nil
	}
	return o.PushOperationKey
}

func (o *PushOperationStatusChangedWebhookData) GetStatus() *PushOperationStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
