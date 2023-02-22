package shared

type CodatDataContractsDatasetsCommerceOrdersDataset struct {
	ContractVersion *string                                   `json:"contractVersion,omitempty"`
	Orders          []CodatDataContractsDatasetsCommerceOrder `json:"orders,omitempty"`
}
