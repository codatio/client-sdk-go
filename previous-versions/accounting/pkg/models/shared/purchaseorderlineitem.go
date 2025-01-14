// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type PurchaseOrderLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Description of the goods / services that have been ordered.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any discounts applied to the unit amount.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	ItemRef            *ItemRef     `json:"itemRef,omitempty"`
	// The purchase order line's number.
	LineNumber *string `json:"lineNumber,omitempty"`
	// Number of units that have been ordered.
	Quantity *decimal.Big `decimal:"number" json:"quantity,omitempty"`
	// Amount of the line, inclusive of discounts but exclusive of tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Amount of tax for the line.
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
	// Total amount of the line, inclusive of discounts and tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Price of each unit.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount,omitempty"`
	// The measurement which defines a unit for this item (e.g. 'kilogram', 'litre').
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
}

func (p PurchaseOrderLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PurchaseOrderLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PurchaseOrderLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *PurchaseOrderLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PurchaseOrderLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *PurchaseOrderLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *PurchaseOrderLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *PurchaseOrderLineItem) GetLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *PurchaseOrderLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *PurchaseOrderLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *PurchaseOrderLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *PurchaseOrderLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *PurchaseOrderLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *PurchaseOrderLineItem) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *PurchaseOrderLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *PurchaseOrderLineItem) GetUnitOfMeasurement() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasurement
}
