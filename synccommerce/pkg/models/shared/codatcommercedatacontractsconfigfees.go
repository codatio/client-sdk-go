package shared

type CodatCommerceDataContractsConfigFees struct {
	Accounts     map[string]CodatCommerceDataContractsConfigConfigAccount `json:"accounts,omitempty"`
	FeesSupplier *CodatCommerceDataContractsConfigFeesSupplier            `json:"feesSupplier,omitempty"`
	SyncFees     *bool                                                    `json:"syncFees,omitempty"`
}
