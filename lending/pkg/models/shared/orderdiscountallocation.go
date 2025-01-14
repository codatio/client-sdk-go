// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type OrderDiscountAllocation struct {
	// Name of the discount in the commerce or point of sale platform.
	Name *string `json:"name,omitempty"`
	// Total amount of discount applied, excluding tax. This is typically positive (for discounts which decrease the amount of the order line), but can also be negative (for discounts which increase the amount of the order line).
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
}

func (o OrderDiscountAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrderDiscountAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrderDiscountAllocation) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OrderDiscountAllocation) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}
