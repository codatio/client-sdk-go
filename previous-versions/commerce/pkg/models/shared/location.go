// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The Location datatype holds information on the geographic location at which stocks of products may be held, as referenced in the Products data type.
//
// A Location also holds information on geographic locations where orders were placed, as referenced in the Orders data type.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-locations) for this data type.
type Location struct {
	Address *Address `json:"address,omitempty"`
	// A unique, persistent identifier for this record
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of this location
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

func (o *Location) GetAddress() *Address {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Location) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Location) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Location) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Location) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
