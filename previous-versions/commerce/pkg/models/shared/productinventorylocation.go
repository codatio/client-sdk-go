// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ProductInventoryLocation struct {
	// Reference to the geographic location where the order was placed.
	LocationRef *LocationRef `json:"locationRef,omitempty"`
	// The quantity of stock remaining at location.
	Quantity *decimal.Big `decimal:"number" json:"quantity,omitempty"`
}

func (p ProductInventoryLocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductInventoryLocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductInventoryLocation) GetLocationRef() *LocationRef {
	if o == nil {
		return nil
	}
	return o.LocationRef
}

func (o *ProductInventoryLocation) GetQuantity() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Quantity
}
