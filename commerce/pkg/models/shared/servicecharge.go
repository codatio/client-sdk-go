package shared

type ServiceChargeTypeEnum string

const (
	ServiceChargeTypeEnumGeneric     ServiceChargeTypeEnum = "Generic"
	ServiceChargeTypeEnumShipping    ServiceChargeTypeEnum = "Shipping"
	ServiceChargeTypeEnumOverpayment ServiceChargeTypeEnum = "Overpayment"
	ServiceChargeTypeEnumUnknown     ServiceChargeTypeEnum = "Unknown"
)

type ServiceCharge struct {
	Description   *string                `json:"description,omitempty"`
	Quantity      *float64               `json:"quantity,omitempty"`
	TaxAmount     *float64               `json:"taxAmount,omitempty"`
	TaxPercentage *float64               `json:"taxPercentage,omitempty"`
	Taxes         []interface{}          `json:"taxes,omitempty"`
	TotalAmount   *float64               `json:"totalAmount,omitempty"`
	Type          *ServiceChargeTypeEnum `json:"type,omitempty"`
}
