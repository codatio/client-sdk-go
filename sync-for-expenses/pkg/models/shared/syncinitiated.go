// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SyncInitiated struct {
	SyncID *string `json:"syncId,omitempty"`
}

func (o *SyncInitiated) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}
