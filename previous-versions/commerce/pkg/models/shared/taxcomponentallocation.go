// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type TaxComponentAllocation struct {
	// Tax amount on order line sale as available from source commerce platform.
	Rate *decimal.Big `decimal:"number" json:"rate,omitempty"`
	// Taxes rates reference object depending on the rates being available on source commerce package.
	TaxComponentRef *TaxComponentRef `json:"taxComponentRef,omitempty"`
}

func (t TaxComponentAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxComponentAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxComponentAllocation) GetRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *TaxComponentAllocation) GetTaxComponentRef() *TaxComponentRef {
	if o == nil {
		return nil
	}
	return o.TaxComponentRef
}
