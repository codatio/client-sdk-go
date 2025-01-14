// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type FeesSupplier struct {
	// Selected supplier id from the list of supplier records on the accounting software.
	SelectedSupplierID *string `json:"selectedSupplierId,omitempty"`
	// List of supplier options from the list of supplier records on the accounting software.
	SupplierOptions []ConfigurationOption `json:"supplierOptions,omitempty"`
}

func (o *FeesSupplier) GetSelectedSupplierID() *string {
	if o == nil {
		return nil
	}
	return o.SelectedSupplierID
}

func (o *FeesSupplier) GetSupplierOptions() []ConfigurationOption {
	if o == nil {
		return nil
	}
	return o.SupplierOptions
}
