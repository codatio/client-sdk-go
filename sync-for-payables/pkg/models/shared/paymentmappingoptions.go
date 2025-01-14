// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PaymentMappingOptions - Gets the bill payments mapping options for a company's accounting software
type PaymentMappingOptions struct {
	BankAccounts []BankAccountMappingOption `json:"bankAccounts,omitempty"`
	Pagination   *Pagination                `json:"pagination,omitempty"`
}

func (o *PaymentMappingOptions) GetBankAccounts() []BankAccountMappingOption {
	if o == nil {
		return nil
	}
	return o.BankAccounts
}

func (o *PaymentMappingOptions) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
