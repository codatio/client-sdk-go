package shared

// PaymentMethodRef
// The Payment Method to which the payment is linked in the accounting platform.
type PaymentMethodRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}
