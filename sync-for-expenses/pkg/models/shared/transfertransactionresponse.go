// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TransferTransactionResponse struct {
	// Unique id of sync created
	SyncID *string `json:"syncId,omitempty"`
}

func (o *TransferTransactionResponse) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}
