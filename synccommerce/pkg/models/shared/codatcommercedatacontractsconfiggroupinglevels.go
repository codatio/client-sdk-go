package shared

type CodatCommerceDataContractsConfigGroupingLevels struct {
	InvoiceLevel     *CodatCommerceDataContractsConfigInvoiceLevelSelection     `json:"invoiceLevel,omitempty"`
	InvoiceLineLevel *CodatCommerceDataContractsConfigInvoiceLineLevelSelection `json:"invoiceLineLevel,omitempty"`
}
