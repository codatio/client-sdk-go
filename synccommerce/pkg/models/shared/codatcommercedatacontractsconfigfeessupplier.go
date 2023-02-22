package shared

type CodatCommerceDataContractsConfigFeesSupplier struct {
	SelectedSupplierID *string                                  `json:"selectedSupplierId,omitempty"`
	SupplierOptions    []CodatCommerceDataContractsConfigOption `json:"supplierOptions,omitempty"`
}
