// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// TaxComponent - The Tax Components endpoints return tax rates data from the commerce software, including tax rate names and values. This is to support the mapping of tax rates from the commerce software to those in the accounting software.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-disputes) for this data type.
type TaxComponent struct {
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// The Boolean flag to indicate when a Tax Rate Component compounds on a sale.
	IsCompound   *bool   `json:"isCompound,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the Tax Rate Component in the source commerce software.
	Name *string `json:"name,omitempty"`
	// Rate of taxation represented as a fraction of the net price (typically in the range 0.00 - 1.00).
	Rate               *decimal.Big `decimal:"number" json:"rate,omitempty"`
	SourceModifiedDate *string      `json:"sourceModifiedDate,omitempty"`
}

func (t TaxComponent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxComponent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *TaxComponent) GetRate() *decimal.Big {
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
