// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProductInventory - Information about the total inventory as well as the locations inventory is in.
type ProductInventory struct {
	Locations     []ProductInventoryLocation `json:"locations,omitempty"`
	TotalQuantity *float64                   `json:"totalQuantity,omitempty"`
}

func (o *ProductInventory) GetLocations() []ProductInventoryLocation {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *ProductInventory) GetTotalQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalQuantity
}