package shared

type CodatCommerceDataContractsConfigPayments struct {
	Accounts     map[string]CodatCommerceDataContractsConfigConfigAccount `json:"accounts,omitempty"`
	SyncPayments *bool                                                    `json:"syncPayments,omitempty"`
}
