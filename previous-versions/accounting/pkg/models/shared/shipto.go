// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ShipToContact - Details of the named contact at the delivery address.
type ShipToContact struct {
	// Email address of the contact at the delivery address.
	Email *string `json:"email,omitempty"`
	// Name of the contact at the delivery address.
	Name *string `json:"name,omitempty"`
	// Phone number of the contact at the delivery address.
	Phone *string `json:"phone,omitempty"`
}

func (o *ShipToContact) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ShipToContact) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ShipToContact) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

// ShipTo - Delivery details for any goods that have been ordered.
type ShipTo struct {
	Address *Addressesitems `json:"address,omitempty"`
	// Details of the named contact at the delivery address.
	Contact *ShipToContact `json:"contact,omitempty"`
}

func (o *ShipTo) GetAddress() *Addressesitems {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ShipTo) GetContact() *ShipToContact {
	if o == nil {
		return nil
	}
	return o.Contact
}