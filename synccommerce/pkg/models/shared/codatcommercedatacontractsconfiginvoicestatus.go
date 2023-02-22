package shared

type CodatCommerceDataContractsConfigInvoiceStatus struct {
	InvoiceStatusOptions  []string `json:"invoiceStatusOptions,omitempty"`
	SelectedInvoiceStatus *string  `json:"selectedInvoiceStatus,omitempty"`
}
