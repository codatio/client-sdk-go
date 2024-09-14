// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// TaxRateRef - Reference to the tax rate to which the line item is linked.
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
