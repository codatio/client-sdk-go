// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountPrototype struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Current balance in the account.
	CurrentBalance *decimal.Big `decimal:"number" json:"currentBalance,omitempty"`
	// Description for the account.
	Description *string `json:"description,omitempty"`
	// Full category of the account.
	//
	// For example, `Liability.Current` or `Income.Revenue`. To determine a list of possible categories for each integration, see our examples, follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide, or refer to the integration's own documentation.
	FullyQualifiedCategory *string `json:"fullyQualifiedCategory,omitempty"`
	// Full name of the account, for example:
	// - `Cash On Hand`
	// - `Rents Held In Trust`
	// - `Fixed Asset`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	// Confirms whether the account is a bank account or not.
	IsBankAccount *bool `json:"isBankAccount,omitempty"`
	// Name of the account.
	Name *string `json:"name,omitempty"`
	// Reference given to each nominal account for a business. It ensures money is allocated to the correct account. This code isn't a unique identifier in the Codat system.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Status of the account
	Status *AccountStatus `json:"status,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Type of account
	Type *AccountType `json:"type,omitempty"`
	// The validDatatypeLinks can be used to determine whether an account can be correctly mapped to another object; for example, accounts with a `type` of `income` might only support being used on an Invoice and Direct Income. For more information, see [Valid Data Type Links](/accounting-api#/schemas/ValidDataTypeLinks).
	ValidDatatypeLinks []ValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

func (a AccountPrototype) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountPrototype) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountPrototype) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountPrototype) GetCurrentBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrentBalance
}

func (o *AccountPrototype) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountPrototype) GetFullyQualifiedCategory() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedCategory
}

func (o *AccountPrototype) GetFullyQualifiedName() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedName
}

func (o *AccountPrototype) GetIsBankAccount() *bool {
	if o == nil {
		return nil
	}
	return o.IsBankAccount
}

func (o *AccountPrototype) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountPrototype) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *AccountPrototype) GetStatus() *AccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountPrototype) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountPrototype) GetType() *AccountType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountPrototype) GetValidDatatypeLinks() []ValidDataTypeLinks {
	if o == nil {
		return nil
	}
	return o.ValidDatatypeLinks
}