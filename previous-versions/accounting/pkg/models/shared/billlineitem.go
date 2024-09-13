// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type BillLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of the goods or services received.
	Description *string `json:"description,omitempty"`
	// Numerical value of any discounts applied.
	//
	// Do not use to apply discounts in Oracle NetSuite—see Oracle NetSuite integration reference.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate of any discount applied to the bill.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	// The bill is a direct cost if `True`.
	IsDirectCost *bool    `json:"isDirectCost,omitempty"`
	ItemRef      *ItemRef `json:"itemRef,omitempty"`
	// The bill line's number.
	LineNumber           *string `json:"lineNumber,omitempty"`
	PurchaseOrderLineRef *Zero   `json:"purchaseOrderLineRef,omitempty"`
	// Number of units of goods or services received.
	Quantity *decimal.Big `decimal:"number" json:"quantity"`
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
	// Total amount of the line, including tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *PropertieTracking `json:"tracking,omitempty"`
	// Collection of categories against which this item is tracked.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// Price of each unit of goods or services.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount"`
	// The measurement which defines a unit for this item (e.g. 'kilogram', 'litre').
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
}

func (b BillLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *BillLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *BillLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *BillLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *BillLineItem) GetIsDirectCost() *bool {
	if o == nil {
		return nil
	}
	return o.IsDirectCost
}

func (o *BillLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *BillLineItem) GetLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *BillLineItem) GetPurchaseOrderLineRef() *Zero {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderLineRef
}

func (o *BillLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Quantity
}

func (o *BillLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *BillLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *BillLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *BillLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *BillLineItem) GetTracking() *PropertieTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *BillLineItem) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *BillLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.UnitAmount
}

func (o *BillLineItem) GetUnitOfMeasurement() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasurement
}
