// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountingCustomerRef struct {
	// `customerName` from the Customer data type
	CompanyName *string `json:"companyName,omitempty"`
	// `id` from the Customers data type
	ID string `json:"id"`
}

func (o *AccountingCustomerRef) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *AccountingCustomerRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
