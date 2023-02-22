package shared

// AccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type AccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
