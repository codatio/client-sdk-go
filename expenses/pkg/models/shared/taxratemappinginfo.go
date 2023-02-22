package shared

type TaxRateMappingInfo struct {
	Code                  *string  `json:"code,omitempty"`
	EffectiveTaxRate      *float64 `json:"effectiveTaxRate,omitempty"`
	ID                    *string  `json:"id,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	TotalTaxRate          *float64 `json:"totalTaxRate,omitempty"`
	ValidTransactionTypes []string `json:"validTransactionTypes,omitempty"`
}
