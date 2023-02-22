package shared

type CodatDataContractsDatasetsCommerceTransactionsDataset struct {
	ContractVersion *string                                         `json:"contractVersion,omitempty"`
	Transactions    []CodatDataContractsDatasetsCommerceTransaction `json:"transactions,omitempty"`
}
