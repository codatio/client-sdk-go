// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Account - This data type provides a list of all the SMB's bank accounts, with rich data like balances, account numbers, and institutions holding the accounts.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
//
// Responses are paged, so you should provide `page` and `pageSize` query parameters in your request.
type Account struct {
	// Depending on the data provided by the underlying bank, not all balances are always available.
	Balance AccountBalanceAmounts `json:"balance"`
	// The currency code for the account.
	Currency string `json:"currency"`
	// The name of the person or company who holds the account.
	Holder *string `json:"holder,omitempty"`
	// The ID of the account from the provider.
	ID string `json:"id"`
	// An object containing bank account identification information.
	Identifiers AccountIdentifiers `json:"identifiers"`
	// The friendly name of the account, chosen by the holder. This may not have been set by the account holder and therefore is not always available.
	InformalName *string `json:"informalName,omitempty"`
	// The bank or other financial institution providing the account.
	Institution  AccountInstitution `json:"institution"`
	ModifiedDate *string            `json:"modifiedDate,omitempty"`
	// The name of the account according to the provider.
	Name               string  `json:"name"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The type of transactions and balances on the account.
	// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
	// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
	Type AccountType `json:"type"`
}

func (o *Account) GetBalance() AccountBalanceAmounts {
	if o == nil {
		return AccountBalanceAmounts{}
	}
	return o.Balance
}

func (o *Account) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *Account) GetHolder() *string {
	if o == nil {
		return nil
	}
	return o.Holder
}

func (o *Account) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Account) GetIdentifiers() AccountIdentifiers {
	if o == nil {
		return AccountIdentifiers{}
	}
	return o.Identifiers
}

func (o *Account) GetInformalName() *string {
	if o == nil {
		return nil
	}
	return o.InformalName
}

func (o *Account) GetInstitution() AccountInstitution {
	if o == nil {
		return AccountInstitution{}
	}
	return o.Institution
}

func (o *Account) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Account) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Account) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Account) GetType() AccountType {
	if o == nil {
		return AccountType("")
	}
	return o.Type
}
