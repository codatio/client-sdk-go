package shared

// ContactRef
// The customer or supplier for the transfer, if available.
type ContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}
