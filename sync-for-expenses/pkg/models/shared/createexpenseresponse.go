// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateExpenseResponse struct {
	// Unique id of sync created
	SyncID *string `json:"syncId,omitempty"`
}

func (o *CreateExpenseResponse) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}
