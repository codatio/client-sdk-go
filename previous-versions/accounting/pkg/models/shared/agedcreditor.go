// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AgedCreditorAgedCurrencyOutstanding struct {
	// Array of outstanding amounts by period.
	AgedOutstandingAmounts []AgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
}

func (o *AgedCreditorAgedCurrencyOutstanding) GetAgedOutstandingAmounts() []AgedOutstandingAmount {
	if o == nil {
		return nil
	}
	return o.AgedOutstandingAmounts
}

func (o *AgedCreditorAgedCurrencyOutstanding) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

type AgedCreditor struct {
	// Array of aged creditors by currency.
	AgedCurrencyOutstanding []AgedCreditorAgedCurrencyOutstanding `json:"agedCurrencyOutstanding,omitempty"`
	// Supplier ID of the aged creditor.
	SupplierID *string `json:"supplierId,omitempty"`
	// Supplier name of the aged creditor.
	SupplierName *string `json:"supplierName,omitempty"`
}

func (o *AgedCreditor) GetAgedCurrencyOutstanding() []AgedCreditorAgedCurrencyOutstanding {
	if o == nil {
		return nil
	}
	return o.AgedCurrencyOutstanding
}

func (o *AgedCreditor) GetSupplierID() *string {
	if o == nil {
		return nil
	}
	return o.SupplierID
}

func (o *AgedCreditor) GetSupplierName() *string {
	if o == nil {
		return nil
	}
	return o.SupplierName
}