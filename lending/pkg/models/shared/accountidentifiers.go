// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/utils"
)

// AccountIdentifiers - An object containing bank account identification information.
type AccountIdentifiers struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The local (usually national) routing number for the account.
	//
	// This is known by different names in different countries:
	// * BSB code (Australia)
	// * routing number (Canada, USA)
	// * sort code (UK)
	BankCode *string `json:"bankCode,omitempty"`
	// The ISO 9362 code (commonly called SWIFT code, SWIFT-BIC or BIC) for the account.
	Bic *string `json:"bic,omitempty"`
	// The international bank account number (IBAN) for the account, if known.
	Iban *string `json:"iban,omitempty"`
	// A portion of the actual account `number` to help account identification where number is tokenised (Plaid only)
	MaskedAccountNumber *string `json:"maskedAccountNumber,omitempty"`
	// The account number for the account. When combined with the`bankCode`, this is usually enough to uniquely identify an account within a jurisdiction.
	Number *string `json:"number,omitempty"`
	// Detailed account category
	Subtype *string `json:"subtype,omitempty"`
	// Type of account
	Type AccountIdentifierType `json:"type"`
}

func (a AccountIdentifiers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountIdentifiers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountIdentifiers) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountIdentifiers) GetBankCode() *string {
	if o == nil {
		return nil
	}
	return o.BankCode
}

func (o *AccountIdentifiers) GetBic() *string {
	if o == nil {
		return nil
	}
	return o.Bic
}

func (o *AccountIdentifiers) GetIban() *string {
	if o == nil {
		return nil
	}
	return o.Iban
}

func (o *AccountIdentifiers) GetMaskedAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.MaskedAccountNumber
}

func (o *AccountIdentifiers) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *AccountIdentifiers) GetSubtype() *string {
	if o == nil {
		return nil
	}
	return o.Subtype
}

func (o *AccountIdentifiers) GetType() AccountIdentifierType {
	if o == nil {
		return AccountIdentifierType("")
	}
	return o.Type
}
