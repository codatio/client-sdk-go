// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProductInventory - Information about the total inventory as well as the locations inventory is in.
type ProductInventory struct {
	Locations     []ProductInventoryLocation `json:"locations,omitempty"`
	TotalQuantity *float64                   `json:"totalQuantity,omitempty"`
}
