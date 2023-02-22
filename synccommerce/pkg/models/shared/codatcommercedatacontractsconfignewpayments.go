package shared

type CodatCommerceDataContractsConfigNewPayments struct {
	Accounts     map[string]CodatCommerceDataContractsConfigConfigAccount `json:"accounts,omitempty"`
	SyncPayments *bool                                                    `json:"syncPayments,omitempty"`
}
