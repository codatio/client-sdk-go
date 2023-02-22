package shared

// ShipToContact
// Details of the named contact at the delivery address.
type ShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// ShipTo
// Delivery details for any goods that have been ordered.
type ShipTo struct {
	Address *Addressesitems `json:"address,omitempty"`
	Contact *ShipToContact  `json:"contact,omitempty"`
}
