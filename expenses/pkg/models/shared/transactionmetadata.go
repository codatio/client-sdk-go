package shared

type TransactionMetadata struct {
	IntegrationType *string `json:"integrationType,omitempty"`
	Message         *string `json:"message,omitempty"`
	Status          *string `json:"status,omitempty"`
	TransactionID   *string `json:"transactionId,omitempty"`
}
