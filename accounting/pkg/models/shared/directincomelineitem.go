// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DirectIncomeLineItem struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	ItemRef            *ItemRef `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity float64 `json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
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
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount float64 `json:"unitAmount"`
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

func (o *DirectIncomeLineItem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *DirectIncomeLineItem) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *DirectIncomeLineItem) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *DirectIncomeLineItem) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *DirectIncomeLineItem) GetSubTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *DirectIncomeLineItem) GetTaxAmount() *float64 {
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

func (o *DirectIncomeLineItem) GetTotalAmount() *float64 {
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

func (o *DirectIncomeLineItem) GetUnitAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitAmount
}
