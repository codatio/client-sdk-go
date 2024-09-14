// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PushOperationChange struct {
	// Unique identifier for the attachment created otherwise null.
	AttachmentID *string           `json:"attachmentId,omitempty"`
	RecordRef    *PushOperationRef `json:"recordRef,omitempty"`
	// Type of change being applied to record in third party platform.
	Type *PushChangeType `json:"type,omitempty"`
}

func (o *PushOperationChange) GetAttachmentID() *string {
	if o == nil {
		return nil
	}
	return o.AttachmentID
}

func (o *PushOperationChange) GetRecordRef() *PushOperationRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

func (o *PushOperationChange) GetType() *PushChangeType {
	if o == nil {
		return nil
	}
	return o.Type
}
