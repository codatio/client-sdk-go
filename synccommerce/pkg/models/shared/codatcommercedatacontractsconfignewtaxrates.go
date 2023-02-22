package shared

type CodatCommerceDataContractsConfigNewTaxRates struct {
	AccountingTaxRateOptions     []CodatCommerceDataContractsConfigOption         `json:"accountingTaxRateOptions,omitempty"`
	CommerceTaxRateOptions       []CodatCommerceDataContractsConfigOption         `json:"commerceTaxRateOptions,omitempty"`
	DefaultZeroTaxRateOptions    []CodatCommerceDataContractsConfigOption         `json:"defaultZeroTaxRateOptions,omitempty"`
	SelectedDefaultZeroTaxRateID *string                                          `json:"selectedDefaultZeroTaxRateId,omitempty"`
	TaxRateMappings              []CodatCommerceDataContractsConfigTaxRateMapping `json:"taxRateMappings,omitempty"`
}
