// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type SalesOrderLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// Description of the goods or services that have been ordered.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *decimal.Big `decimal:"number" json:"discountAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any discounts applied to the unit amount.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
	ItemRef            *ItemRef     `json:"itemRef,omitempty"`
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
	TotalAmount *decimal.Big        `decimal:"number" json:"totalAmount,omitempty"`
	Tracking    *PropertieTracking2 `json:"tracking,omitempty"`
	// Price of each unit.
	UnitAmount *decimal.Big `decimal:"number" json:"unitAmount,omitempty"`
}

func (s SalesOrderLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SalesOrderLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SalesOrderLineItem) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *SalesOrderLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SalesOrderLineItem) GetDiscountAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *SalesOrderLineItem) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *SalesOrderLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *SalesOrderLineItem) GetQuantity() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *SalesOrderLineItem) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *SalesOrderLineItem) GetTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *SalesOrderLineItem) GetTaxRateRef() *TaxRateRef {
	if o == nil {
		return nil
	}
	return o.TaxRateRef
}

func (o *SalesOrderLineItem) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *SalesOrderLineItem) GetTracking() *PropertieTracking2 {
	if o == nil {
		return nil
	}
	return o.Tracking
}

func (o *SalesOrderLineItem) GetUnitAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}
