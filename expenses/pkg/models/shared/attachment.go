package shared

type Attachment struct {
	CompanyID     *string `json:"companyId,omitempty"`
	ID            *string `json:"id,omitempty"`
	TransactionID *string `json:"transactionId,omitempty"`
}
