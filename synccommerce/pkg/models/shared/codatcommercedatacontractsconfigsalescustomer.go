package shared

type CodatCommerceDataContractsConfigSalesCustomer struct {
	CustomerOptions    []CodatCommerceDataContractsConfigOption `json:"customerOptions,omitempty"`
	SelectedCustomerID *string                                  `json:"selectedCustomerId,omitempty"`
}
