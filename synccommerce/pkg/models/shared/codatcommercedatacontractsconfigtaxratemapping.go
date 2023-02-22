package shared

type CodatCommerceDataContractsConfigTaxRateMapping struct {
	SelectedAccountingTaxRateID *string  `json:"selectedAccountingTaxRateId,omitempty"`
	SelectedCommerceTaxRateIds  []string `json:"selectedCommerceTaxRateIds,omitempty"`
}
