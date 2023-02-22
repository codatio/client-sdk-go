package shared

// SupplierRef
// Reference to the supplier the bill was received from.
type SupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}
