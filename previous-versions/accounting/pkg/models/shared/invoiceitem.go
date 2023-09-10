// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
)

// InvoiceItem - Item details that are only for bills.
type InvoiceItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Short description of the product or service that has been bought by the customer.
	Description *string `json:"description,omitempty"`
	// Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.
	//
	// Found on:
	//
	// - Bill line items
	// - Bill Credit Note line items
	// - Credit Note line items
	// - Direct incomes line items
	// - Invoice line items
	// - Items
	TaxRateRef *TaxRateRef `json:"taxRateRef,omitempty"`
	// Unit price of the product or service.
	UnitPrice *types.Decimal `json:"unitPrice,omitempty"`
}

func (o *InvoiceItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *InvoiceItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InvoiceItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *InvoiceItem) GetUnitPrice() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}
