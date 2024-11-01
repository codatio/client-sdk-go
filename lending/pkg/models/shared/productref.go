// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ProductRef - Reference that links the line item to the correct product details.
type ProductRef struct {
	// The unique identifier of the product being referenced.
	ID string `json:"id"`
	// Name of the product being referenced.
	Name *string `json:"name,omitempty"`
}

func (o *ProductRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProductRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
