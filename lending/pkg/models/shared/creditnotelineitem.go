// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/pkg/types"
	"github.com/ericlagergren/decimal"
)

type CreditNoteLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of each line item. For example, the goods or service for which credit has been issued.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *types.Decimal `json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the line item.
	DiscountPercentage *types.Decimal `json:"discountPercentage,omitempty"`
	IsDirectIncome     *bool          `json:"isDirectIncome,omitempty"`
	// Reference to the item the line is linked to.
	ItemRef *ItemRef `json:"itemRef,omitempty"`
	// Number of units of the goods or service for which credit has been issued.
	Quantity types.Decimal `json:"quantity"`
	// Amount of credit associated with the line item, including discounts but excluding tax.
	SubTotal *types.Decimal `json:"subTotal,omitempty"`
	// Amount of tax associated with the line item.
	TaxAmount *types.Decimal `json:"taxAmount,omitempty"`
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
	// Total amount of the line item, including discounts and tax.
	TotalAmount *types.Decimal `json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *AccountsReceivableTracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Unit price of the goods or service.
	UnitAmount types.Decimal `json:"unitAmount"`
}

func (o *CreditNoteLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *CreditNoteLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreditNoteLineItem) GetDiscountAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *CreditNoteLineItem) GetDiscountPercentage() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *CreditNoteLineItem) GetIsDirectIncome() *bool {
	if o == nil {
		return nil
	}
	return o.IsDirectIncome
}

func (o *CreditNoteLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *CreditNoteLineItem) GetQuantity() types.Decimal {
	if o == nil {
		return types.Decimal{Big: *(new(decimal.Big).SetFloat64(0.0))}
	}
	return o.Quantity
}

func (o *CreditNoteLineItem) GetSubTotal() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *CreditNoteLineItem) GetTaxAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *CreditNoteLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *CreditNoteLineItem) GetTotalAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *CreditNoteLineItem) GetTracking() *AccountsReceivableTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *CreditNoteLineItem) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *CreditNoteLineItem) GetUnitAmount() types.Decimal {
	if o == nil {
		return types.Decimal{Big: *(new(decimal.Big).SetFloat64(0.0))}
	}
	return o.UnitAmount
}
