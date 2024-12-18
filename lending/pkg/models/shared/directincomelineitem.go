// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type DirectIncomeLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	// Reference to the item the line is linked to.
	ItemRef *PropertieItemRef `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity *decimal.Big `decimal:"number" json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
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
	// The total amount of the line, including tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount"`
}

func (d DirectIncomeLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectIncomeLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DirectIncomeLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *DirectIncomeLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DirectIncomeLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *DirectIncomeLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *DirectIncomeLineItem) GetItemRef() *PropertieItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *DirectIncomeLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Quantity
}

func (o *DirectIncomeLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *DirectIncomeLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *DirectIncomeLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *DirectIncomeLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *DirectIncomeLineItem) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

func (o *DirectIncomeLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.UnitAmount
}
