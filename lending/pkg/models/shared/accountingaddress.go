// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingAddress struct {
	// City of the customer address.
	City *string `json:"city,omitempty"`
	// Country of the customer address.
	Country *string `json:"country,omitempty"`
	// Line 1 of the customer address.
	Line1 *string `json:"line1,omitempty"`
	// Line 2 of the customer address.
	Line2 *string `json:"line2,omitempty"`
	// Postal code or zip code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Region of the customer address.
	Region *string `json:"region,omitempty"`
	// The type of the address
	Type AccountingAddressType `json:"type"`
}

func (o *AccountingAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *AccountingAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *AccountingAddress) GetLine1() *string {
	if o == nil {
		return nil
	}
	return o.Line1
}

func (o *AccountingAddress) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *AccountingAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AccountingAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AccountingAddress) GetType() AccountingAddressType {
	if o == nil {
		return AccountingAddressType("")
	}
	return o.Type
}
