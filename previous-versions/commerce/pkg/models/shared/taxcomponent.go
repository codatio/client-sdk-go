// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/types"
)

// TaxComponent - The Tax Components endpoints return tax rates data from the commerce platform, including tax rate names and values. This is to support the mapping of tax rates from the commerce platform to those in the accounting platform.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-disputes) for this data type.
type TaxComponent struct {
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// The Boolean flag to indicate when a Tax Rate Component compounds on a sale.
	IsCompound   *bool   `json:"isCompound,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the Tax Rate Component in the source commerce platform.
	Name *string `json:"name,omitempty"`
	// Rate of taxation represented as a fraction of the net price (typically in the range 0.00 - 1.00).
	Rate               *types.Decimal `json:"rate,omitempty"`
	SourceModifiedDate *string        `json:"sourceModifiedDate,omitempty"`
}

func (o *TaxComponent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaxComponent) GetIsCompound() *bool {
	if o == nil {
		return nil
	}
	return o.IsCompound
}

func (o *TaxComponent) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *TaxComponent) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxComponent) GetRate() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *TaxComponent) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
