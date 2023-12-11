// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type InvoiceLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Friendly name of the goods or services provided.
	Description *string `json:"description,omitempty"`
	// Numerical value of any discounts applied.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any discounts applied to the unit amount.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	// The invoice is a direct income if `True`.
	IsDirectIncome *bool `json:"isDirectIncome,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *ItemRef `json:"itemRef,omitempty"`
	// Number of units of goods or services provided.
	Quantity *decimal.Big `decimal:"number" json:"quantity"`
	// Amount of the line, inclusive of discounts but exclusive of tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Amount of tax for the line.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *TaxRateRef `json:"taxRateRef,omitempty"`
	// Total amount of the line, including tax. When pushing invoices to Xero, the total amount is exclusive of tax to allow automatic calculations if a tax rate or tax amount is not specified.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Categories, and a project and customer, against which the item is tracked.
	Tracking *Tracking `json:"tracking,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TrackingCategoryRefs []TrackingCategoryRefItems `json:"trackingCategoryRefs,omitempty"`
	// Price of each unit of goods or services.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount"`
}

func (i InvoiceLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvoiceLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InvoiceLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *InvoiceLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InvoiceLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *InvoiceLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *InvoiceLineItem) GetIsDirectIncome() *bool {
	if o == nil {
		return nil
	}
	return o.IsDirectIncome
}

func (o *InvoiceLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *InvoiceLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Quantity
}

func (o *InvoiceLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *InvoiceLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *InvoiceLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *InvoiceLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *InvoiceLineItem) GetTracking() *Tracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *InvoiceLineItem) GetTrackingCategoryRefs() []TrackingCategoryRefItems {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *InvoiceLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.UnitAmount
}
