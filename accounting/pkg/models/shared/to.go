package shared

// To
// The details of the accounts the transfer is moving to.
type To struct {
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	Amount     *float64    `json:"amount,omitempty"`
	Currency   *string     `json:"currency,omitempty"`
}
