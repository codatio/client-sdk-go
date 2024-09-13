// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type CreditNoteLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of each line item. For example, the goods or service for which credit has been issued.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the line item.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	// The credit note is a direct income if `True`.
	IsDirectIncome *bool `json:"isDirectIncome,omitempty"`
	// Reference to the item the line is linked to.
	ItemRef *PropertieItemRef `json:"itemRef,omitempty"`
	// Number of units of the goods or service for which credit has been issued.
	Quantity *decimal.Big `decimal:"number" json:"quantity"`
	// Amount of credit associated with the line item, including discounts but excluding tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Amount of tax associated with the line item.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount,omitempty"`
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
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *AccountsReceivableTracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Unit price of the goods or service.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount"`
}

func (c CreditNoteLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreditNoteLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *CreditNoteLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *CreditNoteLineItem) GetDiscountPercentage() *decimal.Big {
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

func (o *CreditNoteLineItem) GetItemRef() *PropertieItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *CreditNoteLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Quantity
}

func (o *CreditNoteLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *CreditNoteLineItem) GetTaxAmount() *decimal.Big {
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

func (o *CreditNoteLineItem) GetTotalAmount() *decimal.Big {
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

func (o *CreditNoteLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.UnitAmount
}
