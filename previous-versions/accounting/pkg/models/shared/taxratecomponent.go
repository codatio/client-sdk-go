// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

// TaxRateComponent - A tax rate can be made up of multiple sub taxes, often called components of the tax.
type TaxRateComponent struct {
	// A flag to indicate with the tax is calculated using the principle of compounding.
	IsCompound bool `json:"isCompound"`
	// Name of the tax rate component.
	Name *string `json:"name,omitempty"`
	// The rate of the tax rate component, usually a percentage.
	Rate *types.Decimal `json:"rate,omitempty"`
}

func (o *TaxRateComponent) GetIsCompound() bool {
	if o == nil {
		return false
	}
	return o.IsCompound
}

func (o *TaxRateComponent) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxRateComponent) GetRate() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Rate
}
