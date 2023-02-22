package shared

type AccountMappingInfo struct {
	AccountType           *string  `json:"accountType,omitempty"`
	Currency              *string  `json:"currency,omitempty"`
	ID                    *string  `json:"id,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	ValidTransactionTypes []string `json:"validTransactionTypes,omitempty"`
}
