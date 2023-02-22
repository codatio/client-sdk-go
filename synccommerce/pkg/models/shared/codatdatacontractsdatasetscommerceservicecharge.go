package shared

type CodatDataContractsDatasetsCommerceServiceCharge struct {
	Description   *string                                                  `json:"description,omitempty"`
	Quantity      *int                                                     `json:"quantity,omitempty"`
	TaxAmount     *float64                                                 `json:"taxAmount,omitempty"`
	TaxPercentage *float64                                                 `json:"taxPercentage,omitempty"`
	TotalAmount   *float64                                                 `json:"totalAmount,omitempty"`
	Type          *CodatDataContractsDatasetsCommerceServiceChargeTypeEnum `json:"type,omitempty"`
}
