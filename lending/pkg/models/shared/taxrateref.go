// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// TaxRateRef - Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.
//
// Found on:
//
// - Bill line items
// - Bill Credit Note line items
// - Credit Note line items
// - Direct incomes line items
// - Invoice line items
// - Items
type TaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *decimal.Big `decimal:"number" json:"effectiveTaxRate,omitempty"`
	// Unique identifier for the tax rate in the accounting software.
	ID *string `json:"id,omitempty"`
	// Name of the tax rate in the accounting software.
	Name *string `json:"name,omitempty"`
}

func (t TaxRateRef) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxRateRef) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxRateRef) GetEffectiveTaxRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *TaxRateRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaxRateRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
