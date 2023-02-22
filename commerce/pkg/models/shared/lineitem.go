package shared

type LineItemDiscountAllocations struct {
	Name        *string  `json:"name,omitempty"`
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type LineItem struct {
	DiscountAllocations []LineItemDiscountAllocations `json:"discountAllocations,omitempty"`
	ID                  string                        `json:"id"`
	ProductRef          *Zero                         `json:"productRef,omitempty"`
	ProductVariantRef   *Zero                         `json:"productVariantRef,omitempty"`
	Quantity            *float64                      `json:"quantity,omitempty"`
	TaxPercentage       *float64                      `json:"taxPercentage,omitempty"`
	Taxes               []interface{}                 `json:"taxes,omitempty"`
	TotalAmount         *float64                      `json:"totalAmount,omitempty"`
	TotalTaxAmount      *float64                      `json:"totalTaxAmount,omitempty"`
	UnitPrice           *float64                      `json:"unitPrice,omitempty"`
}
