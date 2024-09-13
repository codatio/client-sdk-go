// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type LendingCustomerRef struct {
	// `customerName` from the Customer data type.
	CustomerName *string `json:"customerName,omitempty"`
	// `id` from the Customers data type.
	ID *string `json:"id,omitempty"`
}

func (o *LendingCustomerRef) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *LendingCustomerRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
