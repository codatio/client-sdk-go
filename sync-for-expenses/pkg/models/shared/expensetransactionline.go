// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/types"
	"github.com/ericlagergren/decimal"
)

type ExpenseTransactionLine struct {
	AccountRef RecordRef `json:"accountRef"`
	// Amount of the line, exclusive of tax.
	NetAmount types.Decimal `json:"netAmount"`
	// Amount of tax for the line.
	TaxAmount    types.Decimal `json:"taxAmount"`
	TaxRateRef   *RecordRef    `json:"taxRateRef,omitempty"`
	TrackingRefs []RecordRef   `json:"trackingRefs,omitempty"`
}

func (o *ExpenseTransactionLine) GetAccountRef() RecordRef {
	if o == nil {
		return RecordRef{}
	}
	return o.AccountRef
}

func (o *ExpenseTransactionLine) GetNetAmount() types.Decimal {
	if o == nil {
		return types.Decimal{Big: *(new(decimal.Big).SetFloat64(0.0))}
	}
	return o.NetAmount
}

func (o *ExpenseTransactionLine) GetTaxAmount() types.Decimal {
	if o == nil {
		return types.Decimal{Big: *(new(decimal.Big).SetFloat64(0.0))}
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
