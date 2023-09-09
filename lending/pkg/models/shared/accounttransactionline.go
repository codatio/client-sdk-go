// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/pkg/types"
)

type AccountTransactionLine struct {
	// Amount in the bill payment currency.
	Amount *types.Decimal `json:"amount,omitempty"`
	// Description of the account transaction.
	Description *string `json:"description,omitempty"`
	// Links the current record to the underlying record or data type that created it.
	//
	// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
	RecordRef *RecordRef `json:"recordRef,omitempty"`
}

func (o *AccountTransactionLine) GetAmount() *types.Decimal {
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

func (o *AccountTransactionLine) GetRecordRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}
