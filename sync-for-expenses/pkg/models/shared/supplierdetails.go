// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SupplierDetails struct {
	// Identifier for the supplier, unique to the company in the accounting software.
	ID *string `json:"id,omitempty"`
}

func (o *SupplierDetails) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
