package shared

type CodatCommerceDataContractsConfigTaxRateAmount struct {
	SelectedTaxRateID *string                                  `json:"selectedTaxRateId,omitempty"`
	TaxRateOptions    []CodatCommerceDataContractsConfigOption `json:"taxRateOptions,omitempty"`
}
