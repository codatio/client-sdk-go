// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountTransactionLine struct {
	// Amount in the bill payment currency.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// Description of the account transaction.
	Description *string `json:"description,omitempty"`
	// Links an account transaction line to the underlying record that created it.
	RecordRef *AccountTransactionLineRecordRef `json:"recordRef,omitempty"`
}

func (a AccountTransactionLine) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountTransactionLine) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountTransactionLine) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AccountTransactionLine) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountTransactionLine) GetRecordRef() *AccountTransactionLineRecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}
