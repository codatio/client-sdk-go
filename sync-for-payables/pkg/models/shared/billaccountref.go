// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BillAccountRef - Reference to the account to which the line item is linked.
type BillAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
}

func (o *BillAccountRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
