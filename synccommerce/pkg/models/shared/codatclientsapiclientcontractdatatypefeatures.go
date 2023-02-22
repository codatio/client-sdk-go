package shared

type CodatClientsAPIClientContractDatatypeFeatures struct {
	Datatype          *string                                              `json:"datatype,omitempty"`
	SupportedFeatures []CodatClientsAPIClientContractSupportedFeatureState `json:"supportedFeatures,omitempty"`
}
