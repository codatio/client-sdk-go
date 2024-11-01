// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SalesConfiguration struct {
	Accounts      map[string]AccountConfiguration `json:"accounts,omitempty"`
	Grouping      *Grouping                       `json:"grouping,omitempty"`
	InvoiceStatus *InvoiceStatus                  `json:"invoiceStatus,omitempty"`
	NewTaxRates   *NewTaxRates                    `json:"newTaxRates,omitempty"`
	SalesCustomer *SalesCustomer                  `json:"salesCustomer,omitempty"`
	// Boolean indicator for syncing sales.
	SyncSales *bool                    `json:"syncSales,omitempty"`
	TaxRates  map[string]TaxRateAmount `json:"taxRates,omitempty"`
}

func (o *SalesConfiguration) GetAccounts() map[string]AccountConfiguration {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *SalesConfiguration) GetGrouping() *Grouping {
	if o == nil {
		return nil
	}
	return o.Grouping
}

func (o *SalesConfiguration) GetInvoiceStatus() *InvoiceStatus {
	if o == nil {
		return nil
	}
	return o.InvoiceStatus
}

func (o *SalesConfiguration) GetNewTaxRates() *NewTaxRates {
	if o == nil {
		return nil
	}
	return o.NewTaxRates
}

func (o *SalesConfiguration) GetSalesCustomer() *SalesCustomer {
	if o == nil {
		return nil
	}
	return o.SalesCustomer
}

func (o *SalesConfiguration) GetSyncSales() *bool {
	if o == nil {
		return nil
	}
	return o.SyncSales
}

func (o *SalesConfiguration) GetTaxRates() map[string]TaxRateAmount {
	if o == nil {
		return nil
	}
	return o.TaxRates
}
