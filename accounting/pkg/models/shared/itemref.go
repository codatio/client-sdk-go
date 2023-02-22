package shared

// ItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type ItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}
