// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ReimbursableExpenseTransactionLine struct {
	AccountRef *RecordRef `json:"accountRef,omitempty"`
	// line description
	Description *string `json:"description,omitempty"`
	// Unique identifier for the customer billed for the transaction. The `invoiceTo` object is currently supported only for QuickBooks Online and QuickBooks Desktop.
	InvoiceTo *InvoiceTo `json:"invoiceTo,omitempty"`
	ItemRef   *ItemRef   `json:"itemRef,omitempty"`
	// Amount of the line, exclusive of tax.
	NetAmount *decimal.Big `decimal:"number" json:"netAmount"`
	// Amount of tax for the line.
	TaxAmount    *decimal.Big  `decimal:"number" json:"taxAmount,omitempty"`
	TaxRateRef   *RecordRef    `json:"taxRateRef,omitempty"`
	TrackingRefs []TrackingRef `json:"trackingRefs,omitempty"`
}

func (r ReimbursableExpenseTransactionLine) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReimbursableExpenseTransactionLine) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReimbursableExpenseTransactionLine) GetAccountRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *ReimbursableExpenseTransactionLine) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ReimbursableExpenseTransactionLine) GetInvoiceTo() *InvoiceTo {
	if o == nil {
		return nil
	}
	return o.InvoiceTo
}

func (o *ReimbursableExpenseTransactionLine) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *ReimbursableExpenseTransactionLine) GetNetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetAmount
}

func (o *ReimbursableExpenseTransactionLine) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *ReimbursableExpenseTransactionLine) GetTaxRateRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *ReimbursableExpenseTransactionLine) GetTrackingRefs() []TrackingRef {
	if o == nil {
		return nil
	}
	return o.TrackingRefs
}
