// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServiceCharge struct {
	// Service charges for this order.
	Description *string `json:"description,omitempty"`
	// The number of times the charge is charged.
	Quantity *int64 `json:"quantity,omitempty"`
	// Amount of the service charge that is tax.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any tax applied to the service charge.
	TaxPercentage *float64 `json:"taxPercentage,omitempty"`
	// Taxes breakdown as applied to service charges.
	Taxes []TaxComponentAllocation `json:"taxes,omitempty"`
	// Total service charge, including taxes.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
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

func (o *ServiceCharge) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *ServiceCharge) GetTaxPercentage() *float64 {
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

func (o *ServiceCharge) GetTotalAmount() *float64 {
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
