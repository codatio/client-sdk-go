// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountMappingOption struct {
	// Identifier for the account, unique for the company.
	ID *string `json:"id,omitempty"`
	// Reference given to each nominal account for a business. It ensures money is allocated to the correct account. This code isn't a unique identifier in the Codat system.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Name of the account.
	Name *string `json:"name,omitempty"`
	// Type of account.
	Type *string `json:"type,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The current status of the account.
	Status             *AccountStatus `json:"status,omitempty"`
	SourceModifiedDate *string        `json:"sourceModifiedDate,omitempty"`
}

func (o *AccountMappingOption) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountMappingOption) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *AccountMappingOption) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountMappingOption) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountMappingOption) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountMappingOption) GetStatus() *AccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountMappingOption) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}