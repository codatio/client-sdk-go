// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SyncCompleteWebhookData struct {
	// Unique identifier for the failed sync.
	SyncID *string `json:"syncId,omitempty"`
	// The type of sync being performed.
	SyncType *string `json:"syncType,omitempty"`
}

func (o *SyncCompleteWebhookData) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}

func (o *SyncCompleteWebhookData) GetSyncType() *string {
	if o == nil {
		return nil
	}
	return o.SyncType
}
