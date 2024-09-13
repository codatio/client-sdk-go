// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PushOperationTimedOutWebhookData struct {
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationGUID *string `json:"pushOperationGuid,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey *string `json:"pushOperationKey,omitempty"`
}

func (o *PushOperationTimedOutWebhookData) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *PushOperationTimedOutWebhookData) GetPushOperationGUID() *string {
	if o == nil {
		return nil
	}
	return o.PushOperationGUID
}

func (o *PushOperationTimedOutWebhookData) GetPushOperationKey() *string {
	if o == nil {
		return nil
	}
	return o.PushOperationKey
}
