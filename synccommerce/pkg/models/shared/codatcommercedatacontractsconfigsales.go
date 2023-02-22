package shared

type CodatCommerceDataContractsConfigSales struct {
	Accounts      map[string]CodatCommerceDataContractsConfigConfigAccount `json:"accounts,omitempty"`
	Grouping      *CodatCommerceDataContractsConfigGrouping                `json:"grouping,omitempty"`
	InvoiceStatus *CodatCommerceDataContractsConfigInvoiceStatus           `json:"invoiceStatus,omitempty"`
	NewTaxRates   *CodatCommerceDataContractsConfigNewTaxRates             `json:"newTaxRates,omitempty"`
	SalesCustomer *CodatCommerceDataContractsConfigSalesCustomer           `json:"salesCustomer,omitempty"`
	SyncSales     *bool                                                    `json:"syncSales,omitempty"`
	TaxRates      map[string]CodatCommerceDataContractsConfigTaxRateAmount `json:"taxRates,omitempty"`
}
