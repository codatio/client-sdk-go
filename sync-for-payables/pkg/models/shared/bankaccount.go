// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BankAccount struct {
	// Identifier for the bank account, unique for the company in the accounting software.
	ID *string `json:"id,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Name of the bank account in the accounting software.
	Name *string `json:"name,omitempty"`
	// The type of transactions and balances on the account.
	// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
	// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
	AccountType *BankAccountType `json:"accountType,omitempty"`
	// Account number for the bank account.
	//
	// Xero integrations
	// Only a UK account number shows for bank accounts with GBP currency and a combined total of sort code and account number that equals 14 digits, For non-GBP accounts, the full bank account number is populated.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Sort code for the bank account. This is relevant to UK bank accounts.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode *string `json:"sortCode,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The current status of the bank account.
	Status             *BankAccountStatus `json:"status,omitempty"`
	SourceModifiedDate *string            `json:"sourceModifiedDate,omitempty"`
}

func (o *BankAccount) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BankAccount) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *BankAccount) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *BankAccount) GetAccountType() *BankAccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *BankAccount) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *BankAccount) GetSortCode() *string {
	if o == nil {
		return nil
	}
	return o.SortCode
}

func (o *BankAccount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *BankAccount) GetStatus() *BankAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *BankAccount) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
