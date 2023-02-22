package shared

type CodatCommerceDataContractsConfigConfiguration struct {
	Fees        *CodatCommerceDataContractsConfigFees        `json:"fees,omitempty"`
	NewPayments *CodatCommerceDataContractsConfigNewPayments `json:"newPayments,omitempty"`
	Payments    *CodatCommerceDataContractsConfigPayments    `json:"payments,omitempty"`
	Sales       *CodatCommerceDataContractsConfigSales       `json:"sales,omitempty"`
}
