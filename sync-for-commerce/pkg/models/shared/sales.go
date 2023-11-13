// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Sales struct {
	Accounts      map[string]ConfigAccount `json:"accounts,omitempty"`
	Grouping      *Grouping                `json:"grouping,omitempty"`
	InvoiceStatus *InvoiceStatus           `json:"invoiceStatus,omitempty"`
	NewTaxRates   *NewTaxRates             `json:"newTaxRates,omitempty"`
	SalesCustomer *Customer                `json:"salesCustomer,omitempty"`
	// Boolean indicator for syncing sales.
	SyncSales *bool                    `json:"syncSales,omitempty"`
	TaxRates  map[string]TaxRateAmount `json:"taxRates,omitempty"`
}

func (o *Sales) GetAccounts() map[string]ConfigAccount {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *Sales) GetGrouping() *Grouping {
	if o == nil {
		return nil
	}
	return o.Grouping
}

func (o *Sales) GetInvoiceStatus() *InvoiceStatus {
	if o == nil {
		return nil
	}
	return o.InvoiceStatus
}

func (o *Sales) GetNewTaxRates() *NewTaxRates {
	if o == nil {
		return nil
	}
	return o.NewTaxRates
}

func (o *Sales) GetSalesCustomer() *Customer {
	if o == nil {
		return nil
	}
	return o.SalesCustomer
}

func (o *Sales) GetSyncSales() *bool {
	if o == nil {
		return nil
	}
	return o.SyncSales
}

func (o *Sales) GetTaxRates() map[string]TaxRateAmount {
	if o == nil {
		return nil
	}
	return o.TaxRates
}
