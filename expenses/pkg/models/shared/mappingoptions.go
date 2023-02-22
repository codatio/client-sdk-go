package shared

type MappingOptions struct {
	Accounts           []AccountMappingInfo          `json:"accounts,omitempty"`
	ExpenseProvider    *string                       `json:"expenseProvider,omitempty"`
	TaxRates           []TaxRateMappingInfo          `json:"taxRates,omitempty"`
	TrackingCategories []TrackingCategoryMappingInfo `json:"trackingCategories,omitempty"`
}
