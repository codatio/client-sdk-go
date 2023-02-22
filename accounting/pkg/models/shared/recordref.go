package shared

// RecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type RecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}
