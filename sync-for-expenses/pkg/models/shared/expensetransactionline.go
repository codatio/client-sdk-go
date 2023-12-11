// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ExpenseTransactionLine struct {
	AccountRef RecordRef `json:"accountRef"`
	// Amount of the line, exclusive of tax.
	NetAmount *decimal.Big `decimal:"number" json:"netAmount"`
	// Amount of tax for the line.
	TaxAmount    *decimal.Big `decimal:"number" json:"taxAmount"`
	TaxRateRef   *RecordRef   `json:"taxRateRef,omitempty"`
	TrackingRefs []RecordRef  `json:"trackingRefs,omitempty"`
}

func (e ExpenseTransactionLine) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExpenseTransactionLine) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExpenseTransactionLine) GetAccountRef() RecordRef {
	if o == nil {
		return RecordRef{}
	}
	return o.AccountRef
}

func (o *ExpenseTransactionLine) GetNetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetAmount
}

func (o *ExpenseTransactionLine) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *ExpenseTransactionLine) GetTaxRateRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *ExpenseTransactionLine) GetTrackingRefs() []RecordRef {
	if o == nil {
		return nil
	}
	return o.TrackingRefs
}
