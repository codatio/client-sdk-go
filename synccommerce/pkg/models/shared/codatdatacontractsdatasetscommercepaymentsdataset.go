package shared

type CodatDataContractsDatasetsCommercePaymentsDataset struct {
	ContractVersion *string                                     `json:"contractVersion,omitempty"`
	Payments        []CodatDataContractsDatasetsCommercePayment `json:"payments,omitempty"`
}
