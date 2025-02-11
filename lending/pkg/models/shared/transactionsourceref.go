// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TransactionSourceRef struct {
	// The unique identitifer of the record being referenced
	ID string `json:"id"`
	// The type of source the transaction arose.
	Type TransactionSourceType `json:"type"`
}

func (o *TransactionSourceRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransactionSourceRef) GetType() TransactionSourceType {
	if o == nil {
		return TransactionSourceType("")
	}
	return o.Type
}
