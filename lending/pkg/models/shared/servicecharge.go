// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/pkg/types"
)

type ServiceCharge struct {
	// Service charges for this order.
	Description *string `json:"description,omitempty"`
	// The number of times the charge is charged.
	Quantity *int64 `json:"quantity,omitempty"`
	// Amount of the service charge that is tax.
	TaxAmount *types.Decimal `json:"taxAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any tax applied to the service charge.
	TaxPercentage *types.Decimal `json:"taxPercentage,omitempty"`
	// Taxes breakdown as applied to service charges.
	Taxes []TaxComponentAllocation `json:"taxes,omitempty"`
	// Total service charge, including taxes.
	TotalAmount *types.Decimal `json:"totalAmount,omitempty"`
	// The type of the service charge.
	Type *ServiceChargeType `json:"type,omitempty"`
}

func (o *ServiceCharge) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ServiceCharge) GetQuantity() *int64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *ServiceCharge) GetTaxAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *ServiceCharge) GetTaxPercentage() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TaxPercentage
}

func (o *ServiceCharge) GetTaxes() []TaxComponentAllocation {
	if o == nil {
		return nil
	}
	return o.Taxes
}

func (o *ServiceCharge) GetTotalAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *ServiceCharge) GetType() *ServiceChargeType {
	if o == nil {
		return nil
	}
	return o.Type
}
