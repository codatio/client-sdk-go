// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Items struct {
	// The third line of the address, or city
	City *string `json:"city,omitempty"`
	// The country for the address
	Country *string `json:"country,omitempty"`
	// The first line of the address
	Line1 *string `json:"line1,omitempty"`
	// The second line of the address
	Line2 *string `json:"line2,omitempty"`
	// The postal (or zip) code for the address
	PostalCode *string `json:"postalCode,omitempty"`
	// The fourth line of the address, or region
	Region *string `json:"region,omitempty"`
	// The type of the address
	Type *CommerceAddressType `json:"type,omitempty"`
}

func (o *Items) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Items) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Items) GetLine1() *string {
	if o == nil {
		return nil
	}
	return o.Line1
}

func (o *Items) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *Items) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *Items) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *Items) GetType() *CommerceAddressType {
	if o == nil {
		return nil
	}
	return o.Type
}
