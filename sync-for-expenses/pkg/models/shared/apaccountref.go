// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ApAccountRef struct {
	// Unique identifier for the Accounts Payable account associated with the transaction. The `apAccountRef` object is currently supported only for QuickBooks Desktop.
	ID *string `json:"id,omitempty"`
}

func (o *ApAccountRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
