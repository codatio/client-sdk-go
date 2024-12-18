// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// TransferAccount - Account details of the account sending or receiving the transfer.
type TransferAccount struct {
	// Links the current record to the underlying record or data type that created it.
	//
	// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
	AccountRef *AccountingRecordRef `json:"accountRef,omitempty"`
	// The amount transferred between accounts.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
}

func (t TransferAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransferAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransferAccount) GetAccountRef() *AccountingRecordRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *TransferAccount) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TransferAccount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}
