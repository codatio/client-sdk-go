// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// ServiceChargeTaxComponentAllocationTaxComponentRef - Taxes rates reference object depending on the rates being available on source commerce package.
type ServiceChargeTaxComponentAllocationTaxComponentRef struct {
	// The unique identitifer of the tax component being referenced.
	ID string `json:"id"`
	// Name of the tax component being referenced.
	Name string `json:"name"`
}

func (o *ServiceChargeTaxComponentAllocationTaxComponentRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ServiceChargeTaxComponentAllocationTaxComponentRef) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ServiceChargeTaxComponentAllocation struct {
	// Tax amount on order line sale as available from source commerce platform.
	Rate *decimal.Big `decimal:"number" json:"rate,omitempty"`
	// Taxes rates reference object depending on the rates being available on source commerce package.
	TaxComponentRef *ServiceChargeTaxComponentAllocationTaxComponentRef `json:"taxComponentRef,omitempty"`
}

func (s ServiceChargeTaxComponentAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceChargeTaxComponentAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceChargeTaxComponentAllocation) GetRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *ServiceChargeTaxComponentAllocation) GetTaxComponentRef() *ServiceChargeTaxComponentAllocationTaxComponentRef {
	if o == nil {
		return nil
	}
	return o.TaxComponentRef
}

type ServiceCharge struct {
	// Service charges for this order.
	Description *string `json:"description,omitempty"`
	// The number of times the charge is charged.
	Quantity *int64 `json:"quantity,omitempty"`
	// Amount of the service charge that is tax.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any tax applied to the service charge.
	TaxPercentage *decimal.Big `decimal:"number" json:"taxPercentage,omitempty"`
	// Taxes breakdown as applied to service charges.
	Taxes []ServiceChargeTaxComponentAllocation `json:"taxes,omitempty"`
	// Total amount of the service charge, including tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// The type of the service charge.
	Type *ServiceChargeType `json:"type,omitempty"`
}

func (s ServiceCharge) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceCharge) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *ServiceCharge) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *ServiceCharge) GetTaxPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxPercentage
}

func (o *ServiceCharge) GetTaxes() []ServiceChargeTaxComponentAllocation {
	if o == nil {
		return nil
	}
	return o.Taxes
}

func (o *ServiceCharge) GetTotalAmount() *decimal.Big {
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
