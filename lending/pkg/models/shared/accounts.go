// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// ReportSourceReference - A source reference containing the `sourceType` object "Banking".
type ReportSourceReference struct {
	// The data source type.
	SourceType *string `json:"sourceType,omitempty"`
}

func (o *ReportSourceReference) GetSourceType() *string {
	if o == nil {
		return nil
	}
	return o.SourceType
}

type Accounts struct {
	// The name of the account according to the provider.
	AccountName *string `json:"accountName,omitempty"`
	// The bank or other financial institution providing the account.
	AccountProvider *string `json:"accountProvider,omitempty"`
	// The type of banking account, e.g. credit or debit.
	AccountType *string `json:"accountType,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The balance of the bank account.
	CurrentBalance *decimal.Big `decimal:"number" json:"currentBalance,omitempty"`
	// Name of the banking data source, e.g. "Plaid".
	PlatformName *string `json:"platformName,omitempty"`
	// A source reference containing the `sourceType` object "Banking".
	SourceRef *ReportSourceReference `json:"sourceRef,omitempty"`
}

func (a Accounts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Accounts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Accounts) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *Accounts) GetAccountProvider() *string {
	if o == nil {
		return nil
	}
	return o.AccountProvider
}

func (o *Accounts) GetAccountType() *string {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *Accounts) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Accounts) GetCurrentBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrentBalance
}

func (o *Accounts) GetPlatformName() *string {
	if o == nil {
		return nil
	}
	return o.PlatformName
}

func (o *Accounts) GetSourceRef() *ReportSourceReference {
	if o == nil {
		return nil
	}
	return o.SourceRef
}
